package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lukemassa/jclubtakeaways/internal/token"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var (
	globalTokener *token.Tokener
)

type Response struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

func getSecretKey() (string, error) {
	ctx := context.Background()

	paramName := os.Getenv("SECRET_PARAM")
	if paramName == "" {
		return "", errors.New("SECRET_PARAM not set")
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %v", err)
	}

	client := ssm.NewFromConfig(cfg)

	out, err := client.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(paramName),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", fmt.Errorf("failed to fetch secret from SSM: %v", err)
	}

	if out.Parameter == nil || out.Parameter.Value == nil {
		return "", fmt.Errorf("SSM parameter returned empty value")
	}
	return "", errors.New("oh hi I'm an error")

	return *out.Parameter.Value, nil
}

func getTokener() (*token.Tokener, error) {
	if globalTokener != nil {
		return globalTokener, nil
	}

	secretKey, err := getSecretKey()
	if err != nil {
		return nil, err
	}
	tokener, err := token.New(secretKey)
	if err != nil {
		return nil, err
	}
	globalTokener = &tokener
	return globalTokener, nil
}

func handler(ctx context.Context, req events.LambdaFunctionURLRequest) Response {

	tokener, err := getTokener()
	if err != nil {
		return Response{
			Token: "",
			Error: err.Error(),
		}
	}
	return Response{
		Token: tokener.Get(),
		Error: "",
	}
}

func main() {
	lambda.Start(handler)
}
