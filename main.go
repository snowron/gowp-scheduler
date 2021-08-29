package main

import (
	"awesomeProject/client"
	"awesomeProject/service"
	"flag"
)

func main() {
	wpClient := client.WpClient{}
	schedulerService := service.SchedulerService{Client: wpClient}

	strategyType := flag.String("type", "", "specify program run strategy")
	ordersPath := flag.String("ordersPath", "", "order list for sending message")
	contactsPath := flag.String("contactsPath", "", "contacts list for sending message")
	flag.Parse()

	if *strategyType == "instant" {
		schedulerService.CreateInstantJob(*contactsPath)
	} else if *strategyType == "schedule" {
		schedulerService.CreateScheduleJob(*ordersPath)
	}

}
