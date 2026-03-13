package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcNatIpCidr = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dry_run": {
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
      "nat_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_ip_cidr": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_ip_cidr_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_ip_cidr_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
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

func AlicloudVpcNatIpCidrSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcNatIpCidr), &result)
	return &result
}
