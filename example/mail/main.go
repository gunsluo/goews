package main

import (
	"fmt"
	"log"

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

	filename := "a.txt"
	// content, err := os.ReadFile(filename)
	// if err != nil {
	// 	log.Fatal("read file: ", err.Error())
	// }
	content := []byte("content")

	htmlBody := `<!DOCTYPE html>
		<html lang="en">
		<head>
		  <title>Simple HTML document</title>
		</head>
		<body>
		  <h1>The email body, as html!</h1>
		</body>
		</html>`

	err = c.SendEmail(
		goews.SendEmailParams{
			From:     "email@exchangedomain",
			To:       []string{"ji.luo@target-energysolutions.com", "ji.luo1@target-energysolutions.com"},
			Cc:       []string{"junkun.ren@target-energysolutions.com", "junkun.ren1@target-energysolutions.com"},
			Bcc:      []string{"Dongsheng.liu@target-energysolutions.com", "Dongsheng.liu1@target-energysolutions.com"},
			Subject:  "An email subject",
			Body:     htmlBody,
			BodyType: schema.BodyTypeHTML,
			FileAttachments: []goews.FileAttachment{
				{
					Name:        filename,
					ContentType: "",
					Size:        int64(len(content)),
					Content:     content,
				},
				{
					Name:        "b.txt",
					ContentType: "",
					Size:        int64(len(content)),
					Content:     content,
				},
			},
		})
	if err != nil {
		log.Fatal("err>: ", err.Error())
	}

	fmt.Println("--- success ---")
}
