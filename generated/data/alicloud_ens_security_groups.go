package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEnsSecurityGroups = `{
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
              "create_time": "string",
              "description": "string",
              "id": "string",
              "instance_count": "number",
              "permissions": [
                "set",
                [
                  "object",
                  {
                    "creation_time": "string",
                    "description": "string",
                    "dest_cidr_ip": "string",
                    "direction": "string",
                    "ip_protocol": "string",
                    "ipv6_dest_cidr_ip": "string",
                    "ipv6_source_cidr_ip": "string",
                    "policy": "string",
                    "port_range": "string",
                    "priority": "number",
                    "source_cidr_ip": "string",
                    "source_port_range": "string"
                  }
                ]
              ],
              "security_group_id": "string",
              "security_group_name": "string"
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
      "security_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEnsSecurityGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEnsSecurityGroups), &result)
	return &result
}
