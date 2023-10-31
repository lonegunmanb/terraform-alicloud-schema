package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudArmsAlertContacts = `{
  "block": {
    "attributes": {
      "alert_contact_name": {
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
              "alert_contact_id": "string",
              "alert_contact_name": "string",
              "create_time": "string",
              "ding_robot_webhook_url": "string",
              "email": "string",
              "id": "string",
              "phone_num": "string",
              "system_noc": "bool",
              "webhook": "string"
            }
          ]
        ]
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
      "phone_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudArmsAlertContactsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudArmsAlertContacts), &result)
	return &result
}
