package main

import (
	"applatix.io/axerror"
	"applatix.io/axops/tool"
	"applatix.io/axrc/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type Repo struct {
	Username   string   `json:"username,omitempty"`
	Protocol   string   `json:"protocol,omitemtpy"`
	Repos      []string `json:"repos,omitempty"`
	UseWebhook *bool    `json:"use_webhook,omitempty"`
}

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
			continue
		}

		listRepos := []Repo{}
		for _, t := range tools {
			configType := t.GetType()
			utils.InfoLog.Println(configType)

			if configType == tool.TypeGIT {
				configObj, ok := t.(*tool.GitConfig)
				if !ok {
					utils.ErrorLog.Printf("Failed to convert to git config, %v", t.GetID())
					continue
				}
				r := Repo {
					Username: configObj.Username,
					Protocol: configObj.Protocol,
					Repos: configObj.Repos,
					UseWebhook: configObj.UseWebhook,
				}
			}

			githubT, ok := t.(*tool.GitHubConfig)
			if !ok {
				utils.ErrorLog.Printf("Failed to convert to Github config, %v", t.GetID())
				continue
			}
			utils.InfoLog.Println(githubT.Repos)
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
