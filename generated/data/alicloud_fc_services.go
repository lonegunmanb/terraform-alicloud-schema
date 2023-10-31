package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcServices = `{
  "block": {
    "attributes": {
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
      "services": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "creation_time": "string",
              "description": "string",
              "id": "string",
              "internet_access": "bool",
              "last_modification_time": "string",
              "log_config": [
                "list",
                [
                  "object",
                  {
                    "logstore": "string",
                    "project": "string"
                  }
                ]
              ],
              "name": "string",
              "nas_config": [
                "list",
                [
                  "object",
                  {
                    "group_id": "number",
                    "mount_points": [
                      "list",
                      [
                        "object",
                        {
                          "mount_dir": "string",
                          "server_addr": "string"
                        }
                      ]
                    ],
                    "user_id": "number"
                  }
                ]
              ],
              "role": "string",
              "vpc_config": [
                "list",
                [
                  "object",
                  {
                    "security_group_id": "string",
                    "vpc_id": "string",
                    "vswitch_ids": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudFcServicesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcServices), &result)
	return &result
}
