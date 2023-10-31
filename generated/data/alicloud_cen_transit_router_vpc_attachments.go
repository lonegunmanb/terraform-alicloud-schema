package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenTransitRouterVpcAttachments = `{
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
              "id": "string",
              "payment_type": "string",
              "resource_type": "string",
              "status": "string",
              "transit_router_attachment_description": "string",
              "transit_router_attachment_id": "string",
              "transit_router_attachment_name": "string",
              "vpc_id": "string",
              "vpc_owner_id": "string",
              "zone_mappings": [
                "list",
                [
                  "object",
                  {
                    "vswitch_id": "string",
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
        "required": true,
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

func AlicloudCenTransitRouterVpcAttachmentsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenTransitRouterVpcAttachments), &result)
	return &result
}
