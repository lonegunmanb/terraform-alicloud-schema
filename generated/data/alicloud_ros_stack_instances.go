package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRosStackInstances = `{
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "parameter_overrides": [
                "list",
                [
                  "object",
                  {
                    "parameter_key": "string",
                    "parameter_value": "string"
                  }
                ]
              ],
              "stack_group_id": "string",
              "stack_group_name": "string",
              "stack_id": "string",
              "stack_instance_account_id": "string",
              "stack_instance_region_id": "string",
              "status": "string",
              "status_reason": "string"
            }
          ]
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stack_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stack_instance_account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stack_instance_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudRosStackInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRosStackInstances), &result)
	return &result
}
