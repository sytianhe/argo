package main

import (
	"applatix.io/axrc"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("usage: axrc http://axdb_url:port/v1 kafka-zk.axsys:9092 http://axops_url")
		os.Exit(1)
	}

	axrc.Init(os.Args[1], os.Args[2], os.Args[3])

	go UpdateRepos()
	router := GetRouterAxrc()
	router.Run(":8992")
}
