package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallNatFirewallControlPolicy = `{
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
      "application_name_list": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "destination_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "direction": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_resolve_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "end_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "new_order": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
        "type": "string"
      },
      "repeat_days": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "number"
        ]
      },
      "repeat_end_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repeat_start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repeat_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AlicloudCloudFirewallNatFirewallControlPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallNatFirewallControlPolicy), &result)
	return &result
}
