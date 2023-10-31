package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectVirtualBorderRouters = `{
  "block": {
    "attributes": {
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
      "routers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_point_id": "string",
              "activation_time": "string",
              "circuit_code": "string",
              "cloud_box_instance_id": "string",
              "create_time": "string",
              "description": "string",
              "detect_multiplier": "number",
              "ecc_id": "string",
              "enable_ipv6": "bool",
              "id": "string",
              "local_gateway_ip": "string",
              "local_ipv6_gateway_ip": "string",
              "min_rx_interval": "number",
              "min_tx_interval": "number",
              "payment_vbr_expire_time": "string",
              "peer_gateway_ip": "string",
              "peer_ipv6_gateway_ip": "string",
              "peering_ipv6_subnet_mask": "string",
              "peering_subnet_mask": "string",
              "physical_connection_business_status": "string",
              "physical_connection_id": "string",
              "physical_connection_owner_uid": "string",
              "physical_connection_status": "string",
              "recovery_time": "string",
              "route_table_id": "string",
              "status": "string",
              "termination_time": "string",
              "type": "string",
              "virtual_border_router_id": "string",
              "virtual_border_router_name": "string",
              "vlan_id": "number",
              "vlan_interface_id": "string"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudExpressConnectVirtualBorderRoutersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectVirtualBorderRouters), &result)
	return &result
}
