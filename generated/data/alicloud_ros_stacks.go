package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRosStacks = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent_stack_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "show_nested_stack": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "stack_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stacks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "deletion_protection": "string",
              "description": "string",
              "disable_rollback": "bool",
              "drift_detection_time": "string",
              "id": "string",
              "parameters": [
                "list",
                [
                  "object",
                  {
                    "parameter_key": "string",
                    "parameter_value": "string"
                  }
                ]
              ],
              "parent_stack_id": "string",
              "ram_role_name": "string",
              "root_stack_id": "string",
              "stack_drift_status": "string",
              "stack_id": "string",
              "stack_name": "string",
              "stack_policy_body": "string",
              "status": "string",
              "status_reason": "string",
              "tags": [
                "map",
                "string"
              ],
              "template_description": "string",
              "timeout_in_minutes": "number"
            }
          ]
        ]
      },
      "status": {
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

func AlicloudRosStacksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRosStacks), &result)
	return &result
}
