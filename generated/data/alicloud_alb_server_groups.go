package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlbServerGroups = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "health_check_config": [
                "list",
                [
                  "object",
                  {
                    "health_check_codes": [
                      "list",
                      "string"
                    ],
                    "health_check_connect_port": "number",
                    "health_check_enabled": "bool",
                    "health_check_host": "string",
                    "health_check_http_version": "string",
                    "health_check_interval": "number",
                    "health_check_method": "string",
                    "health_check_path": "string",
                    "health_check_protocol": "string",
                    "health_check_timeout": "number",
                    "healthy_threshold": "number",
                    "unhealthy_threshold": "number"
                  }
                ]
              ],
              "id": "string",
              "protocol": "string",
              "scheduler": "string",
              "server_group_id": "string",
              "server_group_name": "string",
              "servers": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "port": "number",
                    "server_id": "string",
                    "server_ip": "string",
                    "server_type": "string",
                    "status": "string",
                    "weight": "number"
                  }
                ]
              ],
              "status": "string",
              "sticky_session_config": [
                "list",
                [
                  "object",
                  {
                    "cookie": "string",
                    "cookie_timeout": "number",
                    "sticky_session_enabled": "bool",
                    "sticky_session_type": "string"
                  }
                ]
              ],
              "tags": [
                "map",
                "string"
              ],
              "vpc_id": "string"
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
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "server_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlbServerGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlbServerGroups), &result)
	return &result
}
