package handlers

import (
	"encoding/json"
	"log"
	db "mail_system/internal/db/postgres"
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
				createData := time.Now()
				deliverDate := createData.AddDate(0, 0, 3).Format("2006-11-01")

				packageId, err := handler.Db.CreatePackage(
					r.Context(),
					packageCreate.Weight,
					packageCreate.Type,
					senderReceiver.Val.(db.SenderReceiverRes).Sender,
					senderReceiver.Val.(db.SenderReceiverRes).Receiver,
					departmentReceiver,
					createData.Format("2006-11-01"),
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
						workerId, err := handler.Db.GetEmployeeByLogin(
							r.Context(),
							packageCreate.WorkerLogin,
						)

						if err != nil {
							log.Printf("GET EMPLOYEE BY LOGIN (CREATE PACKAGE) ERROR: %s", err.Error())
							rw.WriteHeader(http.StatusBadRequest)
						} else {
							err = handler.Db.AddEmployeeToPackageResponsibleList(
								r.Context(),
								workerId,
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
									log.Println("PACKAGE WAS CREATED")
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
