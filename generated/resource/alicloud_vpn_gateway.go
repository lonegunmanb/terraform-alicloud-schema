package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpnGateway = `{
  "block": {
    "attributes": {
      "auto_pay": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_propagate": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bandwidth": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "business_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disaster_recovery_internet_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disaster_recovery_vswitch_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_ipsec": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_ssl": {
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
      "instance_charge_type": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_connections": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ssl_vpn_internet_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
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
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpn_gateway_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpn_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func AlicloudVpnGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpnGateway), &result)
	return &result
}
