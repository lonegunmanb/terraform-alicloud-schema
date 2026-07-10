package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectRouterTrAssociations = `{
  "block": {
    "attributes": {
      "association_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "association_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "associations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allowed_prefixes": [
                "list",
                "string"
              ],
              "allowed_prefixes_mode": "string",
              "association_id": "string",
              "association_node_type": "string",
              "cen_id": "string",
              "create_time": "string",
              "ecr_id": "string",
              "id": "string",
              "modify_time": "string",
              "status": "string",
              "transit_router_id": "string",
              "transit_router_owner_id": "string"
            }
          ]
        ]
      },
      "cen_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecr_id": {
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

func AlicloudExpressConnectRouterTrAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectRouterTrAssociations), &result)
	return &result
}
