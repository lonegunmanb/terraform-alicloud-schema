package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectRouterInterfaces = `{
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
      "include_reservation_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interfaces": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_point_id": "string",
              "bandwidth": "number",
              "business_status": "string",
              "connected_time": "string",
              "create_time": "string",
              "cross_border": "bool",
              "description": "string",
              "end_time": "string",
              "has_reservation_data": "string",
              "hc_rate": "number",
              "hc_threshold": "string",
              "health_check_source_ip": "string",
              "health_check_target_ip": "string",
              "id": "string",
              "opposite_access_point_id": "string",
              "opposite_bandwidth": "number",
              "opposite_interface_business_status": "string",
              "opposite_interface_id": "string",
              "opposite_interface_owner_id": "string",
              "opposite_interface_spec": "string",
              "opposite_interface_status": "string",
              "opposite_region_id": "string",
              "opposite_router_id": "string",
              "opposite_router_type": "string",
              "opposite_vpc_instance_id": "string",
              "payment_type": "string",
              "reservation_active_time": "string",
              "reservation_bandwidth": "string",
              "reservation_internet_charge_type": "string",
              "reservation_order_type": "string",
              "role": "string",
              "router_id": "string",
              "router_interface_id": "string",
              "router_interface_name": "string",
              "router_type": "string",
              "spec": "string",
              "status": "string",
              "vpc_instance_id": "string"
            }
          ]
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
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
            "value": {
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

func AlicloudExpressConnectRouterInterfacesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectRouterInterfaces), &result)
	return &result
}
