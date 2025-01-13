package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcIpamIpamPools = `{
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
      "ipam_pool_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_pool_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_scope_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "pool_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pools": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allocation_default_cidr_mask": "number",
              "allocation_max_cidr_mask": "number",
              "allocation_min_cidr_mask": "number",
              "auto_import": "bool",
              "create_time": "string",
              "has_sub_pool": "bool",
              "id": "string",
              "ip_version": "string",
              "ipam_id": "string",
              "ipam_pool_description": "string",
              "ipam_pool_id": "string",
              "ipam_pool_name": "string",
              "ipam_scope_id": "string",
              "pool_depth": "number",
              "pool_region_id": "string",
              "region_id": "string",
              "resource_group_id": "string",
              "source_ipam_pool_id": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_ipam_pool_id": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVpcIpamIpamPoolsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcIpamIpamPools), &result)
	return &result
}
