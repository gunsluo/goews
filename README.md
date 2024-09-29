## EWS Exchange Web Service
Exchange Web Service client for golang

### Usage:

**Here's a reference** [example](./example)
```go
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
```
> Note: if you are using an on-premises Exchange server (or even if you manage your servers at the cloud), you need to pass the username as `AD_DOMAINNAME\username` instead, for examle `MYCOMANY\mhewedy`.

### Supported Feature matrix:

| Category                         	| Operation            	| Supported*       	|
|----------------------------------	|----------------------	|------------------	|
| eDiscovery operations            	|                      	|                  	|
| Exchange mailbox data operations 	|                      	|                  	|
|                                  	| CreateItem operation 	| ✔️ (Email & Calendar)|
|                                  	| GetUserPhoto      	|                   |
| Availability operations          	|                      	|                  	|
|                                  	| GetUserAvailability  	| ✔️             	|
|                                  	| GetRoomLists      	|               	|
| Bulk transfer operations         	|                      	|                  	|
| Delegate management operations   	|                      	|                  	|
| Inbox rules operations           	|                      	|                  	|
| Mail app management operations   	|                      	|                  	|
| Mail tips operation              	|                      	|                  	|
| Message tracking operations      	|                      	|                  	|
| Notification operations          	|                      	|                  	|
| Persona operations               	|                      	|                  	|
|                                   | FindItem              |  ✔️             	|
|                                   | GetItem               |  ✔️            	|
|                                   | FindPeople            |               	|
|                                   | GetPersona            |               	|
| Retention policy operation       	|                      	|                  	|
| Service configuration operation  	|                      	|                  	|
| Sharing operations               	|                      	|                  	|
| Synchronization operations       	|                      	|                  	|
| Time zone operation              	|                      	|                  	|
| Unified Messaging operations     	|                      	|                  	|
| Unified Contact Store operations 	|                      	|                  	|
| User configuration operations    	|                      	|                  	|

* Not always 100% of fields are mapped.

### Extras
Besides the operations supported above, few new operations under the namespace `ewsutil` has been introduced:
* `client.SendEmail` 
* `client.QueryMessage`
* `client.CreateItem`
* `client.FindItem`
* `client.GetItem`

NTLM is supported as well as Basic authentication

#### Reference:

* https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/ews-operations-in-exchange
