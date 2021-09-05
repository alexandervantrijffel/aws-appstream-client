package awsclient

import (
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
		logging.Debugf("request: %s/%+v, Params: %s",
			r.ClientInfo.ServiceName, r.Operation, r.Params)
	})

	return appstream.New(sess), nil
}

type AppStreamContext struct {
	Service   *appstream.AppStream
	StackName string
	FleetName string
}

func CreateStreamingURL(context AppStreamContext, userId, appId string, validity int64) (string, error) {
	logging.Infof("appId %s", appId)
	inp := appstream.CreateStreamingURLInput{
		StackName: &context.StackName,
		FleetName: &context.FleetName,
		UserId:    &userId,
		Validity:  &validity,
		// use SessionContext to pass data to the app instance as environment variable
		// %APPSTREAM_SESSION_CONTEXT%
		//SessionContext: "some value",
	}
	if len(appId) > 0 {
		inp.ApplicationId = &appId
	}

	streamingUrl, err := context.Service.CreateStreamingURL(&inp)
	if err != nil {
		return "", err
	}
	return *streamingUrl.StreamingURL, nil
}

func DescribeSessions(context AppStreamContext) ([]*appstream.Session, error) {
	input := appstream.DescribeSessionsInput{
		StackName: &context.StackName,
		FleetName: &context.FleetName,
	}
	activeSessions, err := context.Service.DescribeSessions(&input)
	if err != nil {
		return activeSessions.Sessions, err
	}
	if len(activeSessions.Sessions) == 0 {
		return activeSessions.Sessions, nil
	}
	return activeSessions.Sessions, nil
}
