package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type User struct {
	UserId        int64
	FirstName     string
	FirstNameKana string
	LastName      string
	LastNameKana  string
	TTL           int64
}

// GetUsers はユーザー情報を全件取得する
func (c Client) GetUsers(ctx context.Context) ([]User, error) {
	// UserIdが高い順に取得
	input := &dynamodb.ScanInput{
		TableName: aws.String("Users"),
		IndexName: aws.String("LastNameIndex"),
	}
	res, err := c.dbclient.Scan(ctx, input)
	if err != nil {
		return nil, err
	}

	u := make([]User, 0, len(res.Items))
	for _, item := range res.Items {
		u = append(u, User{
			UserId: func() int64 {
				v, _ := strconv.ParseInt(item["UserId"].(*types.AttributeValueMemberN).Value, 10, 64)
				return v
			}(),
			FirstName:     item["FirstName"].(*types.AttributeValueMemberS).Value,
			FirstNameKana: item["FirstNameKana"].(*types.AttributeValueMemberS).Value,
			LastName:      item["LastName"].(*types.AttributeValueMemberS).Value,
			LastNameKana:  item["LastNameKana"].(*types.AttributeValueMemberS).Value,
		})
	}
	return u, nil
}

// GetUserByID はユーザーIDを指定してユーザー情報を取得する
func (c Client) GetUserByID(ctx context.Context, userID int64) (User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"UserId": &types.AttributeValueMemberN{Value: strconv.FormatInt(userID, 10)},
		},
	}
	res, err := c.dbclient.GetItem(ctx, input)
	if err != nil {
		return User{}, fmt.Errorf("failed to get user: %w", err)
	}
	var u User
	err = attributevalue.UnmarshalMap(res.Item, &u)
	if err != nil {
		return User{}, fmt.Errorf("failed to unmarshal user: %w", err)
	}
	return u, nil
}

// CreateUser はユーザー情報を新規作成する
func (c Client) CreateUser(ctx context.Context, u User) error {
	_, err := c.dbclient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("Users"),
		Item: map[string]types.AttributeValue{
			"UserId":        &types.AttributeValueMemberN{Value: strconv.FormatInt(u.UserId, 10)},
			"FirstName":     &types.AttributeValueMemberS{Value: u.FirstName},
			"FirstNameKana": &types.AttributeValueMemberS{Value: u.FirstNameKana},
			"LastName":      &types.AttributeValueMemberS{Value: u.LastName},
			"LastNameKana":  &types.AttributeValueMemberS{Value: u.LastNameKana},
		},
	})
	return err
}

// UpdateUser はユーザー情報を更新する
func (c Client) UpdateUser(ctx context.Context, u User) error {
	_, err := c.dbclient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"UserId": &types.AttributeValueMemberN{Value: strconv.FormatInt(u.UserId, 10)},
		},
		AttributeUpdates: map[string]types.AttributeValueUpdate{
			"FirstName":     {Value: &types.AttributeValueMemberS{Value: u.FirstName}},
			"FirstNameKana": {Value: &types.AttributeValueMemberS{Value: u.FirstNameKana}},
			"LastName":      {Value: &types.AttributeValueMemberS{Value: u.LastName}},
			"LastNameKana":  {Value: &types.AttributeValueMemberS{Value: u.LastNameKana}},
		},
	})
	return err
}
