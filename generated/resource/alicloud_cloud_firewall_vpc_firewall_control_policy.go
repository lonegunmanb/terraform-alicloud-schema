package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallVpcFirewallControlPolicy = `{
  "block": {
    "attributes": {
      "acl_action": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "acl_uuid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "application_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "application_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dest_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dest_port_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dest_port_group_ports": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "dest_port_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_group_cidrs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "destination_group_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hit_times": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "member_uid": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "proto": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "release": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_group_cidrs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "source_group_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_firewall_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudFirewallVpcFirewallControlPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallVpcFirewallControlPolicy), &result)
	return &result
}
