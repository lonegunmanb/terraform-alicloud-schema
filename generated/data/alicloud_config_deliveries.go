package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudConfigDeliveries = `{
  "block": {
    "attributes": {
      "deliveries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_id": "string",
              "configuration_item_change_notification": "bool",
              "configuration_snapshot": "bool",
              "delivery_channel_assume_role_arn": "string",
              "delivery_channel_condition": "string",
              "delivery_channel_id": "string",
              "delivery_channel_name": "string",
              "delivery_channel_target_arn": "string",
              "delivery_channel_type": "string",
              "description": "string",
              "id": "string",
              "non_compliant_notification": "bool",
              "oversized_data_oss_target_arn": "string",
              "status": "number"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudConfigDeliveriesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudConfigDeliveries), &result)
	return &result
}
