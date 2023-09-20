package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := kms.NewFromConfig(cfg)

	paginator := kms.NewListKeysPaginator(client, &kms.ListKeysInput{})

	for paginator.HasMorePages() {
		keys, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		policyName := "default"

		for _, key := range keys.Keys {
			output, err := client.GetKeyPolicy(context.TODO(), &kms.GetKeyPolicyInput{KeyId: key.KeyId, PolicyName: &policyName})
			if err != nil {
				log.Fatal(err)
			}

			var policy map[string]interface{}

			json.Unmarshal([]byte(*output.Policy), &policy)

			fmt.Println("Struct is:", policy)
		}
	}
}
