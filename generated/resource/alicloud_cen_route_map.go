package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenRouteMap = `{
  "block": {
    "attributes": {
      "as_path_match_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cen_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cen_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cidr_match_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "community_match_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "community_operate_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_child_instance_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "destination_cidr_blocks": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "destination_instance_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "destination_instance_ids_reverse_match": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "destination_route_table_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "map_result": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "match_asns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "match_community_set": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "next_priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "operate_community_set": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "preference": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "prepend_as_path": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "priority": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "route_map_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "route_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "source_child_instance_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "source_instance_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "source_instance_ids_reverse_match": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "source_region_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "source_route_table_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_router_route_table_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transmit_direction": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudCenRouteMapSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenRouteMap), &result)
	return &result
}
