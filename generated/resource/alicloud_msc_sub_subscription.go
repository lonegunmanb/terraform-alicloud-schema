package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMscSubSubscription = `{
  "block": {
    "attributes": {
      "channel": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "contact_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "email_status": {
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
      "item_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pmsg_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sms_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tts_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "webhook_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "webhook_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMscSubSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMscSubSubscription), &result)
	return &result
}
