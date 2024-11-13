package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenTransitRouterVbrAttachment = `{
  "block": {
    "attributes": {
      "auto_publish_route_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cen_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
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
      "resource_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_association_enabled": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "route_table_propagation_enabled": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      },
      "transit_router_attachment_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_attachment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_router_attachment_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vbr_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vbr_owner_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func AlicloudCenTransitRouterVbrAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenTransitRouterVbrAttachment), &result)
	return &result
}
