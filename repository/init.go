package repository

import (
	"context"

	"github.com/Yuki-TU/dynamodb-sample/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Client struct {
	dbclient *dynamodb.Client
}

func New(ctx context.Context) (*Client, error) {

	cfg, err := awsconfig.LoadDefaultConfig(ctx)
	if config.Get().AWSEndpoint != "" {
		// ローカルへの接続
		cfg.BaseEndpoint = aws.String(config.Get().AWSEndpoint)
	}
	if err != nil {
		return nil, err
	}
	dynamoClient := dynamodb.NewFromConfig(cfg)
	return &Client{
		dbclient: dynamoClient,
	}, nil
}
