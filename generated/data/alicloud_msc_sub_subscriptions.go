package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMscSubSubscriptions = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscriptions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "channel": "string",
              "contact_ids": [
                "list",
                "number"
              ],
              "description": "string",
              "email_status": "number",
              "id": "string",
              "item_id": "string",
              "item_name": "string",
              "pmsg_status": "number",
              "sms_status": "number",
              "tts_status": "number",
              "webhook_ids": [
                "list",
                "number"
              ],
              "webhook_status": "number"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMscSubSubscriptionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMscSubSubscriptions), &result)
	return &result
}
