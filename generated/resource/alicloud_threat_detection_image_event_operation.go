package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionImageEventOperation = `{
  "block": {
    "attributes": {
      "conditions": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_type": {
        "computed": true,
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
      "operation_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scenarios": {
        "description_kind": "plain",
        "optional": true,
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

func AlicloudThreatDetectionImageEventOperationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionImageEventOperation), &result)
	return &result
}
