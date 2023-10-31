package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSddpRules = `{
  "block": {
    "attributes": {
      "category": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "content_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "product_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "risk_level_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "category": "number",
              "category_name": "string",
              "content": "string",
              "content_category": "string",
              "create_time": "string",
              "custom_type": "number",
              "description": "string",
              "display_name": "string",
              "gmt_modified": "string",
              "id": "string",
              "login_name": "string",
              "major_key": "string",
              "name": "string",
              "product_code": "string",
              "product_id": "string",
              "risk_level_id": "string",
              "risk_level_name": "string",
              "rule_id": "string",
              "stat_express": "string",
              "status": "number",
              "target": "string",
              "user_id": "string",
              "warn_level": "number"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "warn_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSddpRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSddpRules), &result)
	return &result
}
