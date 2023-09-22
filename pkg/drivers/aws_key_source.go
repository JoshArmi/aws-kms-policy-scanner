package drivers

import (
	"context"
	"encoding/json"
	"log"

	"aws-kms-policy-scanner/pkg/entities"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func GetKeys() map[string]entities.Key {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := kms.NewFromConfig(cfg)

	paginator := kms.NewListKeysPaginator(client, &kms.ListKeysInput{})

	key_map := map[string]entities.Key{}

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

			policy := entities.Policy{}

			json.Unmarshal([]byte(*output.Policy), &policy)

			key_map[*key.KeyId] = entities.Key{Arn: *key.KeyArn, Policy: policy}
		}
	}

	return key_map
}
