package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenTransitRouterVpnAttachments = `{
  "block": {
    "attributes": {
      "attachments": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_publish_route_enabled": "bool",
              "cen_id": "string",
              "charge_type": "string",
              "create_time": "string",
              "id": "string",
              "resource_type": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "transit_router_attachment_description": "string",
              "transit_router_attachment_id": "string",
              "transit_router_attachment_name": "string",
              "transit_router_id": "string",
              "vpn_id": "string",
              "vpn_owner_id": "number",
              "zone": [
                "set",
                [
                  "object",
                  {
                    "zone_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "cen_id": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
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
      "status": {
        "description_kind": "plain",
        "optional": true,
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
      "transit_router_attachment_id": {
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCenTransitRouterVpnAttachmentsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenTransitRouterVpnAttachments), &result)
	return &result
}
