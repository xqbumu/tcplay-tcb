package main

import (
	"context"
	"encoding/json"

	"github.com/tencentyun/scf-go-lib/cloudevents/scf"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

type DefineEvent struct {
	// test event define
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func hello(ctx context.Context, event scf.APIGatewayProxyRequest) (interface{}, error) {
	ret, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	return scf.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":     "application/json",
			"Content-Encoding": "gzip",
			"X-ORZLAB-FROM":    "hello",
		},
		Body: string(ret),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(hello)
}
