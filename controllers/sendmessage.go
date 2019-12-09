package controllers

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	m "github.com/itsriyan/simple-producers-sqs/models"
)

type MsgSqs struct {
	Message int `json:"message"`
}

func SendMsg(groupId string) []*sqs.SendMessageInput {

	var exsqs []*sqs.SendMessageInput
	datas := m.GetData()

	for _, v := range datas {

		duplicationId, _ := uuid.NewUUID()

		var msgsqss MsgSqs
		msgsqss.Message = v.Data

		b, _ := json.Marshal(msgsqss)

		sp := &sqs.SendMessageInput{
			MessageBody:            aws.String(string(b)),
			QueueUrl:               aws.String(m.Url),
			DelaySeconds:           aws.Int64(0),
			MessageGroupId:         aws.String(groupId),
			MessageDeduplicationId: aws.String(duplicationId.String()),
		}
		exsqs = append(exsqs, sp)
	}

	return exsqs
}
