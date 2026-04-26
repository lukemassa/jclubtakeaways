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

func getSecretKey(ctx context.Context) (string, error) {
	ssmParameterARN := os.Getenv("WEB_CLIENT_KEY_ARN")
	if ssmParameterARN == "" {
		return "", errors.New("WEB_CLIENT_KEY_ARN not set")
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := ssm.NewFromConfig(cfg)

	out, err := client.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(ssmParameterARN),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", fmt.Errorf("failed to fetch secret from SSM: %w", err)
	}

	if out.Parameter == nil || out.Parameter.Value == nil {
		return "", fmt.Errorf("SSM parameter returned empty value")
	}

	return *out.Parameter.Value, nil
}

func getTokener(ctx context.Context) (*token.Tokener, error) {
	if globalTokener != nil {
		return globalTokener, nil
	}

	secretKey, err := getSecretKey(ctx)
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

func handler(ctx context.Context, req events.LambdaFunctionURLRequest) (Response, error) {

	tokener, err := getTokener(ctx)
	if err != nil {
		return Response{
			Error: err.Error(),
		}, nil
	}
	token, err := tokener.Get()
	if err != nil {
		return Response{
			Error: err.Error(),
		}, nil
	}
	return Response{
		Token: token,
	}, nil
}

func main() {
	lambda.Start(handler)
}
