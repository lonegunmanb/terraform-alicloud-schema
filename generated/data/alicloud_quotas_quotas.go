package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudQuotasQuotas = `{
  "block": {
    "attributes": {
      "group_code": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "key_word": {
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
      "product_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "quota_action_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quota_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quotas": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "adjustable": "bool",
              "applicable_range": [
                "list",
                "string"
              ],
              "applicable_type": "string",
              "consumable": "bool",
              "id": "string",
              "quota_action_code": "string",
              "quota_description": "string",
              "quota_name": "string",
              "quota_type": "string",
              "quota_unit": "string",
              "total_quota": "number",
              "total_usage": "number",
              "unadjustable_detail": "string"
            }
          ]
        ]
      },
      "sort_field": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sort_order": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "dimensions": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudQuotasQuotasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudQuotasQuotas), &result)
	return &result
}
