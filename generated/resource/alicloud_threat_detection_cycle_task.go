package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionCycleTask = `{
  "block": {
    "attributes": {
      "enable": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "first_date_str": {
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
      "interval_period": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "param": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period_unit": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_end_time": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "target_start_time": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "task_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func AlicloudThreatDetectionCycleTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionCycleTask), &result)
	return &result
}
