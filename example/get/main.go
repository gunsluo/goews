package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gunsluo/goews/v3"
	"github.com/gunsluo/goews/v3/schema"
)

func main() {
	c, err := goews.NewClient(
		goews.SetAddress("https://outlook.office365.com/EWS/Exchange.asmx"),
		goews.SetCredentials("email@exchangedomain", "password"),
		goews.EnabledNTLM(),
		goews.SkipTLS(),
		goews.Debug(),
	)
	if err != nil {
		log.Fatal("->: ", err.Error())
	}

	messages, err := c.QueryMessage(goews.QueryMessageParams{
		FolderId:  schema.DistinguishedFolderIdinbox,
		StartTime: time.Now().Add(-1 * time.Hour),
		Offset:    0,
		Limit:     20,
		BodyType:  schema.BodyTypeText,
	})
	if err != nil {
		log.Fatal("err>: ", err.Error())
	}

	for _, message := range messages {
		fmt.Println("->", message.Sender, message.Subject, message.Body)
	}

	fmt.Println("--- success ---")
}
