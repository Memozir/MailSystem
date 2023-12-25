package handlers

import (
	"encoding/json"
	"log"
	"mail_system/internal/config"
	db "mail_system/internal/db/postgres"
	"mail_system/internal/model"
	"net/http"
	"time"
)

type PackageCreateJSON struct {
	Type        int    `json:"type"`
	Sender      string `json:"sender"`
	Receiver    string `json:"receiver"`
	Weight      int    `json:"weight"`
	WorkerLogin string `json:"worker_login"`
}

func (handler *MailHandlers) CreateDepartmentPackageHandler(rw http.ResponseWriter, r *http.Request) {
	var packageCreate PackageCreateJSON

	err := json.NewDecoder(r.Body).Decode(&packageCreate)
	if err != nil {
		log.Printf("CREATE PACKAGE DECODE ERROR: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		senderReceiver, err := handler.Db.GetSenderReceiverIdByLogin(
			r.Context(),
			packageCreate.Sender,
			packageCreate.Receiver)

		if err != nil {
			log.Printf("CREATE PACKAGE GET SENDER AND RECEIVER ERROR: %s", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			departmentReceiver, err := handler.Db.GetDepartmentByReceiver(r.Context(),
				senderReceiver.Val.(db.SenderReceiverRes).Receiver)

			if err != nil {
				log.Printf("CREATE PACKAGE GET DEPARTMENT BY RECEIVER ERROR: %s", err.Error())
				rw.WriteHeader(http.StatusBadRequest)
			} else {
				createDate := time.Now()
				deliverDate := createDate.AddDate(0, 0, 3).Format("2006-01-02")

				tran, err := handler.Db.BeginTran(r.Context())
				defer func() {
					err = tran.Rollback(r.Context())
					log.Printf("TRANSACTION ROLLBACK ERROR: %s", err.Error())
					rw.WriteHeader(http.StatusBadRequest)
				}()

				if err != nil {
					log.Printf("TRANSACTION BEGIN ERROR: %s", err.Error())
					rw.WriteHeader(http.StatusBadRequest)
				} else {
					packageId, err := handler.Db.CreatePackage(
						r.Context(),
						packageCreate.Weight,
						packageCreate.Type,
						senderReceiver.Val.(db.SenderReceiverRes).Sender,
						senderReceiver.Val.(db.SenderReceiverRes).Receiver,
						departmentReceiver,
						createDate.Format("2006-01-02"),
						deliverDate,
					)

					if err != nil {
						log.Printf("CREATE PACKAGE: %s", err.Error())
						rw.WriteHeader(http.StatusBadRequest)
					} else {
						err = handler.Db.ProducePaymentInfo(
							r.Context(),
							packageId,
							packageCreate.Type,
							packageCreate.Weight,
						)

						if err != nil {
							log.Printf("PRODUCE PAYMENT INFO ERROR: %s", err.Error())
							rw.WriteHeader(http.StatusBadRequest)
						} else {

							err = handler.Db.AddPackageToClient(
								r.Context(),
								senderReceiver.Val.(db.SenderReceiverRes).Sender,
								packageId,
							)

							if err != nil {
								log.Printf("ADD PACKAGE TO CLIENT ERROR: %s", err.Error())
								rw.WriteHeader(http.StatusBadRequest)
							} else {
								worker := handler.Db.GetEmployeeByLogin(
									r.Context(),
									packageCreate.WorkerLogin,
								)

								if worker.Err != nil {
									log.Printf("GET EMPLOYEE BY LOGIN (CREATE PACKAGE) ERROR: %s", err.Error())
									rw.WriteHeader(http.StatusBadRequest)
								} else {
									err = handler.Db.AddEmployeeToPackageResponsibleList(
										r.Context(),
										worker.Val.(model.Employee).EmployeeId,
										packageId,
									)

									if err != nil {
										log.Printf("ADD EMPLOYEE TO RESPONSIBLE LIST ERROR: %s", err.Error())
										rw.WriteHeader(http.StatusBadRequest)
									} else {
										currentDepartmentId := handler.Db.GetEmployeeDepartment(r.Context(),
											packageCreate.WorkerLogin)

										if currentDepartmentId.Err != nil {
											log.Printf("GET CURRENT DEPARTMENT ERROR: %s", err.Error())
											rw.WriteHeader(http.StatusBadRequest)
										} else {
											adminId, err := handler.Db.GetEmployeeDepartmentByRole(r.Context(),
												currentDepartmentId.Val.(uint64), config.AdminRole)

											if err != nil {
												log.Printf("GET ADMIN OF CURRENT DEPARTMENT ERROR: %s", err.Error())
												rw.WriteHeader(http.StatusBadRequest)
											} else {

												err = handler.Db.AddEmployeeToPackageResponsibleList(
													r.Context(),
													adminId,
													packageId,
												)

												if err != nil {
													log.Printf("ADD ADMIN TO RESPONSIBLE LIST ERROR: %s", err.Error())
													rw.WriteHeader(http.StatusBadRequest)
												} else {
													err = handler.Db.AddPackageToStorehouse(
														r.Context(),
														departmentReceiver,
														packageId,
														false,
													)

													if err != nil {
														log.Printf("ADD PACKAGE TO STOREHOUSE ERROR: %s", err.Error())
														rw.WriteHeader(http.StatusBadRequest)
													} else {
														log.Println("PACKAGE WAS CREATED")
														err = tran.Commit(r.Context())

														if err != nil {
															log.Printf("COMMIT TRANSACTION ERROR: %s", err.Error())
															rw.WriteHeader(http.StatusBadRequest)
														} else {
															rw.WriteHeader(http.StatusCreated)
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

type EmployeePackagesResponseJSON struct {
	Packages []model.Package `json:"packages"`
}

func (handler *MailHandlers) GetEmployeePackages(rw http.ResponseWriter, r *http.Request) {

	var employeeInfo UserAuthRequest
	err := json.NewDecoder(r.Body).Decode(&employeeInfo)

	if err != nil {
		log.Printf("GET EMPLOYEE PACKAGES ERROR: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	}

	employee := handler.Db.GetEmployeeByLogin(r.Context(), employeeInfo.Login)

	if employee.Err != nil {
		log.Printf("GET EMPLOYEE ERROR: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		res, err := handler.Db.GetEmployeePackages(r.Context(), 17)
		//res, err := handler.Db.GetEmployeePackages(r.Context(), 1, int(employee.Val.(model.Employee).RoleCode))

		if err != nil {
			log.Printf("GET EMPLOYEE PACKAGES ERROR: %s", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			response := EmployeePackagesResponseJSON{Packages: res}
			err = json.NewEncoder(rw).Encode(&response)

			if err != nil {
				log.Printf("GET EMPLOYEE PACKAGES MARSHAL PACKAGES ERROR: %s", err.Error())
				rw.WriteHeader(http.StatusBadRequest)
			} else {
				log.Printf("GET EMPLOYEE PACKAGES SUCCESS")
				rw.WriteHeader(http.StatusOK)
			}
		}
	}
}
