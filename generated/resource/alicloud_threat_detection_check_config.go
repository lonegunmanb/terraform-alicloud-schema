package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionCheckConfig = `{
  "block": {
    "attributes": {
      "configure": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cycle_days": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "number"
        ]
      },
      "enable_add_check": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_auto_check": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "end_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "system_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vendors": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "selected_checks": {
        "block": {
          "attributes": {
            "check_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "section_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
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

func AlicloudThreatDetectionCheckConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionCheckConfig), &result)
	return &result
}
