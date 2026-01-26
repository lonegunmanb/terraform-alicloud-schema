package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlikafkaScheduledScalingRule = `{
  "block": {
    "attributes": {
      "duration_minutes": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "first_scheduled_time": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "repeat_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserved_pub_flow": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "reserved_sub_flow": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "time_zone": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "weekly_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
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

func AlicloudAlikafkaScheduledScalingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlikafkaScheduledScalingRule), &result)
	return &result
}
