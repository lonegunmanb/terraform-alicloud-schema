package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGpdbDbInstanceIpArray = `{
  "block": {
    "attributes": {
      "db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_ip_array_attribute": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_ip_array_name": {
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
      "modify_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_ip_list": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudGpdbDbInstanceIpArraySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGpdbDbInstanceIpArray), &result)
	return &result
}
