package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {
	ctx := context.Background()

	// Load the AWS configuration (from environment, shared config files, etc.)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	// Create an STS client
	client := sts.NewFromConfig(cfg)

	// Call GetCallerIdentity
	output, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("failed to get caller identity, %v", err)
	}

	// Print the result
	fmt.Printf("UserId: %s\n", *output.UserId)
	fmt.Printf("Account: %s\n", *output.Account)
	fmt.Printf("ARN: %s\n", *output.Arn)
}
