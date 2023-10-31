package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMongodbShardingNetworkPrivateAddresses = `{
  "block": {
    "attributes": {
      "addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "db_instance_id": "string",
              "expired_time": "string",
              "ip_address": "string",
              "network_address": "string",
              "network_type": "string",
              "node_id": "string",
              "node_type": "string",
              "port": "string",
              "role": "string",
              "vpc_id": "string",
              "vswitch_id": "string"
            }
          ]
        ]
      },
      "db_instance_id": {
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
      "node_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMongodbShardingNetworkPrivateAddressesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMongodbShardingNetworkPrivateAddresses), &result)
	return &result
}
