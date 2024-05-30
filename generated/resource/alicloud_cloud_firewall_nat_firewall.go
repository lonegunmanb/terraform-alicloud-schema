package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallNatFirewall = `{
  "block": {
    "attributes": {
      "firewall_switch": {
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
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "proxy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region_no": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "strict_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vswitch_auto": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "nat_route_entry_list": {
        "block": {
          "attributes": {
            "destination_cidr": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "nexthop_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "nexthop_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "route_table_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
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

func AlicloudCloudFirewallNatFirewallSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallNatFirewall), &result)
	return &result
}
