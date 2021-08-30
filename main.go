package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"gowp-scheduler/client"
	"gowp-scheduler/service"
)

const INSTANT = "instant"
const SCHEDULE = "schedule"

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

	if *strategyType == INSTANT {
		messageText := getFromCli()

		err := schedulerService.CreateInstantJob(*contactsPath, messageText)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	} else if *strategyType == SCHEDULE {
		err := schedulerService.CreateScheduleJob(*ordersPath)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
}

func getFromCli() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Message: ")

	messageText, _ := reader.ReadString('\n')
	messageText = strings.Replace(messageText, "\n", "", -1)

	fmt.Println(messageText)
	fmt.Print("Are You Sure? [y/N] :")

	yesOrNo, _ := reader.ReadString('\n')
	yesOrNo = strings.Replace(yesOrNo, "\n", "", -1)

	if !strings.EqualFold(yesOrNo, "y") {
		os.Exit(1)
	}

	return messageText
}
