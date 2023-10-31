package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNasAccessRules = `{
  "block": {
    "attributes": {
      "access_group_name": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_rule_id": "string",
              "priority": "number",
              "rw_access": "string",
              "source_cidr_ip": "string",
              "user_access": "string"
            }
          ]
        ]
      },
      "rw_access": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_cidr_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_access": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudNasAccessRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNasAccessRules), &result)
	return &result
}
