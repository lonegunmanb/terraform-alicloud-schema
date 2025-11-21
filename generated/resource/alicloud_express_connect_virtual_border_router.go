package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectVirtualBorderRouter = `{
  "block": {
    "attributes": {
      "associated_physical_connections": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bandwidth": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "circuit_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "detect_multiplier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enable_ipv6": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_cross_account_vbr": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "local_gateway_ip": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_ipv6_gateway_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_rx_interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_tx_interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "mtu": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "peer_gateway_ip": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_ipv6_gateway_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peering_ipv6_subnet_mask": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peering_subnet_mask": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "physical_connection_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sitelink_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
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
      "vbr_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "virtual_border_router_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vlan_id": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudExpressConnectVirtualBorderRouterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectVirtualBorderRouter), &result)
	return &result
}
