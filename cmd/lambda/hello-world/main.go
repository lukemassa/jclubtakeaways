package main

import (
  "context"

  "github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) (string, error) {
  return "hello from lambda, but with CI!", nil
}

func main() {
  lambda.Start(handler)
}
