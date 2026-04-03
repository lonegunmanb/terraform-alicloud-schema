package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaWafRulesets = `{
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
      "phase": {
        "description_kind": "plain",
        "required": true,
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
              "fields": [
                "list",
                "string"
              ],
              "id": "string",
              "name": "string",
              "phase": "string",
              "ruleset_id": "string",
              "status": "string",
              "target": "string",
              "types": [
                "list",
                "string"
              ],
              "update_time": "string"
            }
          ]
        ]
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_version": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "query_args": {
        "block": {
          "attributes": {
            "any_like": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "desc": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name_like": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "order_by": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEsaWafRulesetsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaWafRulesets), &result)
	return &result
}
