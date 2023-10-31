package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudConfigAggregateDelivery = `{
  "block": {
    "attributes": {
      "aggregator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "configuration_item_change_notification": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "configuration_snapshot": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delivery_channel_condition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delivery_channel_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_channel_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delivery_channel_target_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "delivery_channel_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "non_compliant_notification": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "oversized_data_oss_target_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AlicloudConfigAggregateDeliverySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudConfigAggregateDelivery), &result)
	return &result
}
