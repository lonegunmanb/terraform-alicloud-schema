package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMessageServiceSubscription = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "endpoint": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filter_tag": {
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
      "notify_content_format": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notify_strategy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "push_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subscription_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "topic_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "dlq_policy": {
        "block": {
          "attributes": {
            "dead_letter_target_queue": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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

func AlicloudMessageServiceSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMessageServiceSubscription), &result)
	return &result
}
