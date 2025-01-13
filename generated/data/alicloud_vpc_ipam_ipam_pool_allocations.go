package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcIpamIpamPoolAllocations = `{
  "block": {
    "attributes": {
      "allocations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cidr": "string",
              "create_time": "string",
              "id": "string",
              "ipam_pool_allocation_description": "string",
              "ipam_pool_allocation_id": "string",
              "ipam_pool_allocation_name": "string",
              "ipam_pool_id": "string",
              "region_id": "string",
              "resource_id": "string",
              "resource_owner_id": "number",
              "resource_region_id": "string",
              "resource_type": "string",
              "source_cidr": "string",
              "status": "string",
              "total_count": "number"
            }
          ]
        ]
      },
      "cidr": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ipam_pool_allocation_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_pool_allocation_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_pool_id": {
        "description_kind": "plain",
        "required": true,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVpcIpamIpamPoolAllocationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcIpamIpamPoolAllocations), &result)
	return &result
}
