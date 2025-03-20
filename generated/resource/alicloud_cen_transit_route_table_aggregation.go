package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenTransitRouteTableAggregation = `{
  "block": {
    "attributes": {
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
      "transit_route_table_aggregation_cidr": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_route_table_aggregation_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_route_table_aggregation_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_route_table_aggregation_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_route_table_aggregation_scope_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "transit_route_table_id": {
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

func AlicloudCenTransitRouteTableAggregationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenTransitRouteTableAggregation), &result)
	return &result
}
