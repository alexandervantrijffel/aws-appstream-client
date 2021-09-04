package awsclient

import (
	"encoding/json"

	"github.com/alexandervantrijffel/goutil/logging"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appstream"
)

func NewService() (*appstream.AppStream, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		// Profile: "myProfile",

		// Provide SDK Config options, such as Region.
		// Config: aws.Config{
		// 	Region: aws.String("eu-west-1"),
		// },

		// Force enable Shared Config support
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}
	sess.Handlers.Send.PushFront(func(r *request.Request) {
		// Log every request made and its payload
		logging.Debugf("Request: %s/%s, Params: %s",
			r.ClientInfo.ServiceName, r.Operation, r.Params)
	})

	return appstream.New(sess),nil
}

func DescribeSessions(service *appstream.AppStream, stackName, fleetName string) error {
	logging.Info("Sessions command started")
	input := appstream.DescribeSessionsInput{
		StackName: &stackName,
		FleetName: &fleetName,
	}
	activeSessions, err := service.DescribeSessions(&input)
	if err != nil {
		return err
	}
	if len(activeSessions.Sessions) == 0 {
		logging.Info("No active AppStream sessions found for the current region")
		return nil
	}

	logging.Info("Sessions", prettyPrint(activeSessions.Sessions))

	return nil
}

func prettyPrint(object interface{}) string {
	r, _ := json.MarshalIndent(object, "", "    ")
	return string(r)
}
