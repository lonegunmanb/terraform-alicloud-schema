package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcIpamIpamPool = `{
  "block": {
    "attributes": {
      "allocation_default_cidr_mask": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_max_cidr_mask": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_min_cidr_mask": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "auto_import": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "clear_allocation_default_cidr_mask": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "ip_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_pool_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_pool_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_scope_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pool_region_id": {
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
      "source_ipam_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func AlicloudVpcIpamIpamPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcIpamIpamPool), &result)
	return &result
}
