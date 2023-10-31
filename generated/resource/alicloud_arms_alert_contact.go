package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudArmsAlertContact = `{
  "block": {
    "attributes": {
      "alert_contact_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ding_robot_webhook_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "email": {
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
      "phone_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_noc": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudArmsAlertContactSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudArmsAlertContact), &result)
	return &result
}
