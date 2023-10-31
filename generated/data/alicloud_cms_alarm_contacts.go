package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsAlarmContacts = `{
  "block": {
    "attributes": {
      "chanel_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "chanel_value": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "contacts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alarm_contact_name": "string",
              "channels_aliim": "string",
              "channels_ding_web_hook": "string",
              "channels_mail": "string",
              "channels_sms": "string",
              "channels_state_aliim": "string",
              "channels_state_ding_web_hook": "string",
              "channels_state_mail": "string",
              "channels_status_sms": "string",
              "contact_groups": [
                "list",
                "string"
              ],
              "describe": "string",
              "id": "string",
              "lang": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsAlarmContactsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsAlarmContacts), &result)
	return &result
}
