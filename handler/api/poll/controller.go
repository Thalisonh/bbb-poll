package poll

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
)

type IPollController interface {
	Vote(ctx *gin.Context)
}

type PollController struct {
}

func NewPollController() IPollController {
	return &PollController{}
}

func (c *PollController) Vote(ctx *gin.Context) {
	id := ""
	secret := ""

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(id, secret, ""),
	})

	queue := ""
	message := "{teste}"

	svc := sqs.New(sess)
	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(1),
		QueueUrl:     &queue,
		MessageBody:  &message,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Not working"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "It's working"})
}
