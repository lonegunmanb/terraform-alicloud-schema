package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenTransitRouterCidr = `{
  "block": {
    "attributes": {
      "cidr": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "publish_cidr_route": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "transit_router_cidr_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_router_cidr_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_id": {
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

func AlicloudCenTransitRouterCidrSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenTransitRouterCidr), &result)
	return &result
}
