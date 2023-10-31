package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVideoSurveillanceSystemGroups = `{
  "block": {
    "attributes": {
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "app": "string",
              "callback": "string",
              "create_time": "string",
              "description": "string",
              "enabled": "bool",
              "gb_id": "string",
              "gb_ip": "string",
              "group_id": "string",
              "group_name": "string",
              "id": "string",
              "in_protocol": "string",
              "out_protocol": "string",
              "play_domain": "string",
              "push_domain": "string",
              "stats": [
                "list",
                [
                  "object",
                  {
                    "device_num": "string",
                    "ied_num": "string",
                    "ipc_num": "string",
                    "platform_num": "string"
                  }
                ]
              ]
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
      "in_protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVideoSurveillanceSystemGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVideoSurveillanceSystemGroups), &result)
	return &result
}
