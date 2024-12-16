package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcIpamIpamPoolAllocation = `{
  "block": {
    "attributes": {
      "cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cidr_mask": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_pool_allocation_description": {
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
      "region_id": {
        "computed": true,
        "description_kind": "plain",
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

func AlicloudVpcIpamIpamPoolAllocationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcIpamIpamPoolAllocation), &result)
	return &result
}
