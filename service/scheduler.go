package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"gowp-scheduler/client"
	"gowp-scheduler/model"
	"os"
	"time"
)

type SchedulerInterface interface {
	CreateInstantJob(contactsPath string) error
	CreateScheduleJob(ordersPath string) error
}

type SchedulerService struct {
	Client client.ChatClient
}

func (s SchedulerService) CreateInstantJob(contactsPath, message string) error {
	file, err := readCsvFromFile(contactsPath)
	if err != nil {
		return err
	}
	contacts := parseContacts(file)

	if len(contacts) == 0 {
		return errors.New("contact csv length is zero")
	}

	for _, contact := range contacts {
		fmt.Println(contact.Name, contact.Number)
		err = s.Client.SendMessage(message, contact.Number)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s SchedulerService) CreateScheduleJob(ordersPath string) error {
	file, err := readCsvFromFile(ordersPath)
	if err != nil {
		return err
	}
	orders := parseOrders(file)

	if len(orders) == 0 {
		return errors.New("contact csv length is zero")
	}

	doneOrders := 0
	const loopTime = 10
	// nolint:staticcheck
	for range time.Tick(time.Second * loopTime) {
		for _, order := range orders {
			if doneOrders == len(orders) {
				return nil
			}
			dataTime, err := time.Parse(time.RFC3339, order.Date)
			if err != nil {
				return err
			}

			isBefore := dataTime.Before(time.Now())

			if !order.Completed && !isBefore {
				err = s.Client.SendMessage(order.Message, order.Number)
				if err != nil {
					return err
				}
				order.Completed = true
				doneOrders++
			}
		}
	}

	return nil
}

func readCsvFromFile(path string) ([][]string, error) {
	if path == "" {
		return nil, errors.New("path is missing")
	}
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fmt.Printf("successfully opened %s\n", path)
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	return csvLines, nil
}

func parseContacts(data [][]string) []model.Contact {
	var contacts []model.Contact
	for _, line := range data {
		contact := model.Contact{
			Name:   line[0],
			Number: line[1],
		}
		contacts = append(contacts, contact)
	}
	fmt.Printf("successfully parsed contacts")
	return contacts
}

func parseOrders(data [][]string) []*model.Order {
	orders := make([]*model.Order, 0)

	for _, line := range data {
		order := model.Order{
			Number:    line[0],
			Message:   line[1],
			Date:      line[2],
			Completed: false,
		}
		orders = append(orders, &order)
	}

	fmt.Printf("successfully parsed orders")
	return orders
}
