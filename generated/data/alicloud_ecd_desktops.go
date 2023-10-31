package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdDesktops = `{
  "block": {
    "attributes": {
      "desktop_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desktops": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cpu": "number",
              "create_time": "string",
              "desktop_id": "string",
              "desktop_name": "string",
              "desktop_type": "string",
              "directory_id": "string",
              "end_user_ids": [
                "list",
                "string"
              ],
              "expired_time": "string",
              "id": "string",
              "image_id": "string",
              "memory": "string",
              "network_interface_id": "string",
              "payment_type": "string",
              "policy_group_id": "string",
              "status": "string",
              "system_disk_size": "number"
            }
          ]
        ]
      },
      "end_user_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
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
      "office_site_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcdDesktopsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdDesktops), &result)
	return &result
}
