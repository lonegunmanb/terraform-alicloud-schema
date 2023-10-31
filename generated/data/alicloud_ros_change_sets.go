package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRosChangeSets = `{
  "block": {
    "attributes": {
      "change_set_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
      "sets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "change_set_id": "string",
              "change_set_name": "string",
              "change_set_type": "string",
              "description": "string",
              "disable_rollback": "bool",
              "execution_status": "string",
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
              "stack_id": "string",
              "stack_name": "string",
              "status": "string",
              "template_body": "string",
              "timeout_in_minutes": "number"
            }
          ]
        ]
      },
      "stack_id": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudRosChangeSetsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRosChangeSets), &result)
	return &result
}
