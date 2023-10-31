package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRosStackInstance = `{
  "block": {
    "attributes": {
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
      "retain_stacks": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "stack_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stack_instance_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stack_instance_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "timeout_in_minutes": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "parameter_overrides": {
        "block": {
          "attributes": {
            "parameter_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "parameter_value": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
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

func AlicloudRosStackInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRosStackInstance), &result)
	return &result
}
