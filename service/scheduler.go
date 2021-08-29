package service

import (
	"awesomeProject/client"
	"awesomeProject/model"
	"encoding/csv"
	"errors"
	"fmt"
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

func (s SchedulerService) CreateInstantJob(contactsPath string) error {
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
		s.Client.SendMessage("Here is the message", contact.Number)
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

	for _, order := range orders {
		dataTime, err := time.Parse(time.RFC3339, order.Date)
		if err != nil {
			return err
		}

		isBefore := dataTime.Before(time.Now())

		if !order.Completed && !isBefore {
			s.Client.SendMessage(order.Message, order.Number)
		}
	}
	return nil
}

func readCsvFromFile(path string) ([][]string, error) {
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
	fmt.Printf("succesfully parsed contacts")
	return contacts
}

func parseOrders(data [][]string) []model.Order {
	var orders []model.Order
	for _, line := range data {
		order := model.Order{
			Number:    line[0],
			Message:   line[1],
			Date:      line[2],
			Completed: false,
		}
		orders = append(orders, order)
	}

	fmt.Printf("succesfully parsed orders")
	return orders
}
