package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApiGatewayInstance = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
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
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "support_ipv6": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
