package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallControlPolicies = `{
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
      "direction": {
        "description_kind": "plain",
        "required": true,
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
        "type": [
          "list",
          "string"
        ]
      },
      "ip_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
              "direction": "string",
              "dns_result": "string",
              "dns_result_time": "string",
              "hit_times": "string",
              "id": "string",
              "order": "number",
              "proto": "string",
              "release": "bool",
              "source": "string",
              "source_group_cidrs": [
                "list",
                "string"
              ],
              "source_group_type": "string",
              "source_type": "string"
            }
          ]
        ]
      },
      "proto": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudFirewallControlPoliciesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallControlPolicies), &result)
	return &result
}
