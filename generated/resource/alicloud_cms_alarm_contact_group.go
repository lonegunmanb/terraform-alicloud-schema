package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsAlarmContactGroup = `{
  "block": {
    "attributes": {
      "alarm_contact_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "contacts": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "describe": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_subscribed": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsAlarmContactGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsAlarmContactGroup), &result)
	return &result
}
