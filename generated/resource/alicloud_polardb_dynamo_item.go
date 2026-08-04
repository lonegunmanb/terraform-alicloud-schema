package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbDynamoItem = `{
  "block": {
    "attributes": {
      "account_auth": {
        "computed": true,
        "description": "The authentication password for PolarDB DynamoDB. If not set, it is resolved from the cluster's DynamoDB-type account automatically.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "account_name": {
        "computed": true,
        "description": "The account name for PolarDB DynamoDB authentication. If not set, it is resolved from the cluster's DynamoDB-type account automatically.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "db_cluster_id": {
        "description": "The ID of the PolarDB cluster where the DynamoDB table resides.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint": {
        "description": "The PolarDB DynamoDB-compatible endpoint URL.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hash_key": {
        "description": "The partition key (hash key) attribute name of the item.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "item": {
        "description": "JSON representation of the DynamoDB item attributes.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "range_key": {
        "description": "The sort key (range key) attribute name of the item.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table_name": {
        "description": "The name of the DynamoDB-compatible table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudPolardbDynamoItemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbDynamoItem), &result)
	return &result
}
