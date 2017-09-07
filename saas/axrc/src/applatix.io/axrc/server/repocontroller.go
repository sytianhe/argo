package main

import (
	"applatix.io/axerror"
	"applatix.io/axops/tool"
	"applatix.io/axrc/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func PingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(axerror.REST_STATUS_OK, "pong")
		return
	}
}

func UpdateRepos() {
	for {
		time.Sleep(time.Second * 30)
		tools, axErr := tool.GetToolsByCategory("scm")
		if axErr != nil {
			utils.ErrorLog.Printf("Failed to retrieve tools (err: %v)", axErr)
		}

		for _, t := range tools {
			utils.InfoLog.Println(t)
			//if t.GetType() == "BitbucketConfig" {
			//	t.(tool.GitHubConfig).Repos
			//} else if t.GetType() == "Bitbucket" {
			//	t.(tool.BitbucketConfig)
			//}
			//

		}
	}
}

func ParseRepoUrl() {

}
