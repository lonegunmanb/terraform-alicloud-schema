package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApiGatewayInstance = `{
  "block": {
    "attributes": {
      "connect_cidr_blocks": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delete_vpc_ip_block": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "egress_ipv6_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "https_policy": {
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
      "instance_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_spec": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pricing_cycle": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_wait_switch": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "support_ipv6": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "to_connect_vpc_ip_block": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "user_vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_slb_intranet_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "zone_id": {
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
      },
      "zone_vswitch_security_group": {
        "block": {
          "attributes": {
            "cidr_block": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "security_group": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vswitch_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "zone_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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

func AlicloudApiGatewayInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApiGatewayInstance), &result)
	return &result
}
