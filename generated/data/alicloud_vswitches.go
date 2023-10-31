package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVswitches = `{
  "block": {
    "attributes": {
      "cidr_block": {
        "description_kind": "plain",
        "optional": true,
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
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_id": {
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
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "vswitches": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "available_ip_address_count": "number",
              "cidr_block": "string",
              "creation_time": "string",
              "description": "string",
              "id": "string",
              "ipv6_cidr_block": "string",
              "is_default": "bool",
              "name": "string",
              "resource_group_id": "string",
              "route_table_id": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "vpc_id": "string",
              "vswitch_id": "string",
              "vswitch_name": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVswitchesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVswitches), &result)
	return &result
}
