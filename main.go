package main

import (
	"awesomeProject/client"
	"awesomeProject/service"
	"flag"
	"os"
)

func main() {
	wpClient := client.WpClient{}
	schedulerService := service.SchedulerService{Client: wpClient}

	strategyType := flag.String("type", "", "specify program run strategy")
	ordersPath := flag.String("ordersPath", "", "order list for sending message")
	contactsPath := flag.String("contactsPath", "", "contacts list for sending message")
	flag.Parse()

	if *strategyType == "instant" || *strategyType == "schedule" {
		err := wpClient.CreateConnection()
		if err != nil {
			return
		}
	}
	if *strategyType == "instant" {
		err := schedulerService.CreateInstantJob(*contactsPath)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	} else if *strategyType == "schedule" {
		err := schedulerService.CreateScheduleJob(*ordersPath)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
}
