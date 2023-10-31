package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallVpcFirewallControlPolicies = `{
  "block": {
    "attributes": {
      "acl_action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "acl_uuid": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination": {
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
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "member_uid": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "policies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acl_action": "string",
              "acl_uuid": "string",
              "application_id": "string",
              "application_name": "string",
              "description": "string",
              "dest_port": "string",
              "dest_port_group": "string",
              "dest_port_group_ports": [
                "list",
                "string"
              ],
              "dest_port_type": "string",
              "destination": "string",
              "destination_group_cidrs": [
                "list",
                "string"
              ],
              "destination_group_type": "string",
              "destination_type": "string",
              "hit_times": "number",
              "id": "string",
              "member_uid": "string",
              "order": "number",
              "proto": "string",
              "release": "bool",
              "source": "string",
              "source_group_cidrs": [
                "list",
                "string"
              ],
              "source_group_type": "string",
              "source_type": "string",
              "vpc_firewall_id": "string"
            }
          ]
        ]
      },
      "proto": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_firewall_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudFirewallVpcFirewallControlPoliciesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallVpcFirewallControlPolicies), &result)
	return &result
}
