package main

import (
	"applatix.io/common"
	"flag"
	"fmt"
	"os"
)

func main() {
	common.InitLoggers("AXRC")
	axdbAddr := flag.String("axdb", "", "URL of AXDB")
	kafkaAddr := flag.String("kafka", "", "URL of Kafka")
	notificationServiceAddr := flag.String("notification", "", "URL of notification service")
	axopsAddr := flag.String("axops", "", "URL of axops")

	flag.Parse()

	if *axdbAddr == "" || *kafkaAddr == "" || *notificationServiceAddr == "" || *axopsAddr == "" {
		fmt.Println("Usage: must provide urls for AXDB, Kafka, AXNC and AXOPS")
		os.Exit(1)
	}

	router := GetRouterAxrc()
	router.Run(":8992")
}
