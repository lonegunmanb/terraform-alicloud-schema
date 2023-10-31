package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenChildInstanceRouteEntryToAttachments = `{
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
              "cen_id": "string",
              "child_instance_route_table_id": "string",
              "destination_cidr_block": "string",
              "id": "string",
              "service_type": "string",
              "status": "string",
              "transit_router_attachment_id": "string"
            }
          ]
        ]
      },
      "cen_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "child_instance_route_table_id": {
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
      "service_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_attachment_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCenChildInstanceRouteEntryToAttachmentsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenChildInstanceRouteEntryToAttachments), &result)
	return &result
}
