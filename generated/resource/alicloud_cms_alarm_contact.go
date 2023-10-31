package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsAlarmContact = `{
  "block": {
    "attributes": {
      "alarm_contact_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "channels_aliim": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "channels_ding_web_hook": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "channels_mail": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "channels_sms": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "describe": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsAlarmContactSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsAlarmContact), &result)
	return &result
}
