package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenTransitRouterPeerAttachment = `{
  "block": {
    "attributes": {
      "auto_publish_route_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bandwidth": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "bandwidth_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cen_bandwidth_package_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cen_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_link_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "peer_transit_router_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_transit_router_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_association_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "route_table_propagation_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func AlicloudCenTransitRouterPeerAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenTransitRouterPeerAttachment), &result)
	return &result
}
