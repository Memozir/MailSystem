package handlers

import (
	"encoding/json"
	"log"
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
						/*
							err = handler.Db.AddPackageToClient(
								r.Context(),
								senderReceiver.Val.(db.SenderReceiverRes).Sender,
								packageId,
							)
						*/
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
								err = handler.Db.AddPackageToStorehouse(
									r.Context(),
									departmentReceiver,
									packageId,
									false,
								)

								if err != nil {
									log.Printf("ADD PACKAGE TO STOREHOUSE ERROR: %s", err.Error())
									rw.WriteHeader(http.StatusCreated)
								} else {
									log.Println("PACKAGE WAS CREATED")
								}
							}
						}
					}
				}
			}
		}
	}
}
