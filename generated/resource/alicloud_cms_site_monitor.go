package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsSiteMonitor = `{
  "block": {
    "attributes": {
      "address": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "alert_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "options_json": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "task_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "task_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "isp_cities": {
        "block": {
          "attributes": {
            "city": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "isp": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func AlicloudCmsSiteMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsSiteMonitor), &result)
	return &result
}
