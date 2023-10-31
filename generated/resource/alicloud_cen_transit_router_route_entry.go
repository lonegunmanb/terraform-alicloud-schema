package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenTransitRouterRouteEntry = `{
  "block": {
    "attributes": {
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
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_router_route_entry_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_route_entry_destination_cidr_block": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_router_route_entry_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_router_route_entry_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_route_entry_next_hop_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_route_entry_next_hop_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_router_route_table_id": {
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

func AlicloudCenTransitRouterRouteEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenTransitRouterRouteEntry), &result)
	return &result
}
