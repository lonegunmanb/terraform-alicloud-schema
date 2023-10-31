package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSecurityGroupRules = `{
  "block": {
    "attributes": {
      "direction": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_desc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nic_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "dest_cidr_ip": "string",
              "dest_group_id": "string",
              "dest_group_owner_account": "string",
              "direction": "string",
              "ip_protocol": "string",
              "nic_type": "string",
              "policy": "string",
              "port_range": "string",
              "priority": "number",
              "source_cidr_ip": "string",
              "source_group_id": "string",
              "source_group_owner_account": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSecurityGroupRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSecurityGroupRules), &result)
	return &result
}
