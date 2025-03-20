package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcRouteEntry = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_cidr_block": {
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
      "nexthop_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nexthop_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_entry_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_id": {
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
      "next_hops": {
        "block": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "next_hop_region_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "next_hop_related_info": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "instance_id": "string",
                    "instance_type": "string",
                    "region_id": "string"
                  }
                ]
              ]
            },
            "nexthop_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nexthop_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "route_publish_targets": {
        "block": {
          "attributes": {
            "publish_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "target_instance_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
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

func AlicloudVpcRouteEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcRouteEntry), &result)
	return &result
}
