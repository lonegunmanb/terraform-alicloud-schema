package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMaxComputeTunnelQuotaTimer = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nickname": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "time_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "quota_timer": {
        "block": {
          "attributes": {
            "begin_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "end_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "tunnel_quota_parameter": {
              "block": {
                "attributes": {
                  "elastic_reserved_slot_num": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "slot_num": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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

func AlicloudMaxComputeTunnelQuotaTimerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMaxComputeTunnelQuotaTimer), &result)
	return &result
}
