package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectRouterInterface = `{
  "block": {
    "attributes": {
      "access_point_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_pay": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "business_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connected_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cross_border": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "delete_health_check_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "has_reservation_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hc_rate": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "hc_threshold": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_source_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_target_ip": {
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
      "opposite_access_point_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "opposite_bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "opposite_interface_business_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "opposite_interface_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "opposite_interface_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "opposite_interface_spec": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "opposite_interface_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "opposite_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "opposite_router_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "opposite_router_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "opposite_vpc_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "pricing_cycle": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reservation_active_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reservation_bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reservation_internet_charge_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reservation_order_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "router_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "router_interface_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router_interface_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "spec": {
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
      "vpc_instance_id": {
        "computed": true,
        "description_kind": "plain",
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

func AlicloudExpressConnectRouterInterfaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectRouterInterface), &result)
	return &result
}
