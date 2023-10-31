package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudBastionhostInstances = `{
  "block": {
    "attributes": {
      "description_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "descriptions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "id": "string",
              "instance_status": "string",
              "license_code": "string",
              "private_domain": "string",
              "public_domain": "string",
              "public_network_access": "bool",
              "security_group_ids": [
                "list",
                "string"
              ],
              "tags": [
                "map",
                "string"
              ],
              "user_vswitch_id": "string"
            }
          ]
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudBastionhostInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudBastionhostInstances), &result)
	return &result
}
