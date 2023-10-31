package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcs = `{
  "block": {
    "attributes": {
      "cidr_block": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dhcp_options_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_details": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "is_default": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
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
      "total_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "vpc_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "vpcs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cidr_block": "string",
              "creation_time": "string",
              "description": "string",
              "id": "string",
              "ipv6_cidr_block": "string",
              "is_default": "bool",
              "region_id": "string",
              "resource_group_id": "string",
              "route_table_id": "string",
              "router_id": "string",
              "secondary_cidr_blocks": [
                "list",
                "string"
              ],
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "user_cidrs": [
                "list",
                "string"
              ],
              "vpc_id": "string",
              "vpc_name": "string",
              "vrouter_id": "string",
              "vswitch_ids": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVpcsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcs), &result)
	return &result
}
