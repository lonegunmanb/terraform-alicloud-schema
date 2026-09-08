package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigPolicy = `{
  "block": {
    "attributes": {
      "attach_resource_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "attach_resource_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environment_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_class_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_class_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_config": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AlicloudApigPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigPolicy), &result)
	return &result
}
