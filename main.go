package main

import (
	"log"
	"os"
	// "time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"github.com/itsriyan/simple-producers-sqs/controllers"
	"github.com/itsriyan/simple-producers-sqs/models"
)

func main() {
	groupId, _ := uuid.NewUUID()

	for i := 0; i < 1; i++ {

		sess := session.New(&aws.Config{
			Region:      aws.String(models.Region),
			Credentials: credentials.NewSharedCredentials(os.Getenv("AWS"), models.CredProfile),
			MaxRetries:  aws.Int(5),
		})

		svc := sqs.New(sess)

		log.Println("Start Send Data")
		sendparams := controllers.SendMsg(groupId.String())
		if len(sendparams) >= 1 {

			for _, v := range sendparams {

				send_resp, err := svc.SendMessage(v)
				if err != nil {
					log.Println(err)
				}

				log.Printf("[Data] \n%v \n\n", send_resp)
			}
		}
	}
}
