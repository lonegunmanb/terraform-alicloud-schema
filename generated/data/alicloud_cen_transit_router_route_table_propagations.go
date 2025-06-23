package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenTransitRouterRouteTablePropagations = `{
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "propagations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "resource_id": "string",
              "resource_type": "string",
              "status": "string",
              "transit_router_attachment_id": "string",
              "transit_router_route_table_id": "string"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_attachment_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_route_table_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCenTransitRouterRouteTablePropagationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenTransitRouterRouteTablePropagations), &result)
	return &result
}
