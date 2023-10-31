package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNlbServerGroups = `{
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
              "address_ip_version": "string",
              "connection_drain": "bool",
              "connection_drain_timeout": "number",
              "health_check": [
                "list",
                [
                  "object",
                  {
                    "health_check_connect_port": "number",
                    "health_check_connect_timeout": "number",
                    "health_check_domain": "string",
                    "health_check_enabled": "bool",
                    "health_check_http_code": [
                      "list",
                      "string"
                    ],
                    "health_check_interval": "number",
                    "health_check_type": "string",
                    "health_check_url": "string",
                    "healthy_threshold": "number",
                    "http_check_method": "string",
                    "unhealthy_threshold": "number"
                  }
                ]
              ],
              "id": "string",
              "preserve_client_ip_enabled": "bool",
              "protocol": "string",
              "related_load_balancer_ids": [
                "list",
                "string"
              ],
              "resource_group_id": "string",
              "scheduler": "string",
              "server_count": "number",
              "server_group_name": "string",
              "server_group_type": "string",
              "status": "string",
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
      "server_group_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "server_group_type": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudNlbServerGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNlbServerGroups), &result)
	return &result
}
