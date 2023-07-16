package awspile

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Credentials struct {
	AccessKeyId     string
	AccessSecretKey string
}

func (c *Credentials) Retrieve(ctx context.Context) (*aws.Credentials, error) {
	return nil, nil
}
