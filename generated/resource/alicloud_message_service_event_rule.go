package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMessageServiceEventRule = `{
  "block": {
    "attributes": {
      "delivery_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_types": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "match_rules": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "list",
            [
              "object",
              {
                "match_state": "string",
                "name": "string",
                "prefix": "string",
                "suffix": "string"
              }
            ]
          ]
        ]
      },
      "rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "endpoint": {
        "block": {
          "attributes": {
            "endpoint_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "endpoint_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
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

func AlicloudMessageServiceEventRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMessageServiceEventRule), &result)
	return &result
}
