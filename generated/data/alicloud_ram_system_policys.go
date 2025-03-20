package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRamSystemPolicys = `{
  "block": {
    "attributes": {
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
      "policys": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attachment_count": "number",
              "create_time": "string",
              "description": "string",
              "id": "string",
              "policy_name": "string",
              "policy_type": "string",
              "update_date": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudRamSystemPolicysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRamSystemPolicys), &result)
	return &result
}
