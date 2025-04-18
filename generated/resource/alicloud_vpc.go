package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpc = `{
  "block": {
    "attributes": {
      "cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "classic_link_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "dns_hostname_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_ipv6": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_delete": {
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
      "ipv4_cidr_mask": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ipv4_ipam_pool_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_cidr_blocks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ipv6_cidr_block": "string",
              "ipv6_isp": "string"
            }
          ]
        ]
      },
      "ipv6_isp": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_default": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
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
      "router_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "router_table_id": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_cidr_blocks": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "secondary_cidr_mask": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "system_route_table_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_route_table_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_route_table_route_propagation_enable": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "user_cidrs": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_name": {
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

func AlicloudVpcSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpc), &result)
	return &result
}
