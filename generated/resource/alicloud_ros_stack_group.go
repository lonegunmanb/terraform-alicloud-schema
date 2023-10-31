package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRosStackGroup = `{
  "block": {
    "attributes": {
      "account_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "administration_role_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role_name": {
        "computed": true,
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
      "operation_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "operation_preferences": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stack_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_body": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "parameters": {
        "block": {
          "attributes": {
            "parameter_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parameter_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
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

func AlicloudRosStackGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRosStackGroup), &result)
	return &result
}
