package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionAntiBruteForceRule = `{
  "block": {
    "attributes": {
      "anti_brute_force_rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_rule": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fail_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "forbidden_time": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "span": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "uuid_list": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "protocol_type": {
        "block": {
          "attributes": {
            "rdp": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sql_server": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssh": {
              "computed": true,
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
            },
            "update": {
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

func AlicloudThreatDetectionAntiBruteForceRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionAntiBruteForceRule), &result)
	return &result
}
