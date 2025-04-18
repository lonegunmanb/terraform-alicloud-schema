package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMongodbPublicNetworkAddress = `{
  "block": {
    "attributes": {
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
      "replica_sets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_domain": "string",
              "connection_port": "string",
              "connection_type": "string",
              "network_type": "string",
              "replica_set_role": "string",
              "role_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMongodbPublicNetworkAddressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMongodbPublicNetworkAddress), &result)
	return &result
}
