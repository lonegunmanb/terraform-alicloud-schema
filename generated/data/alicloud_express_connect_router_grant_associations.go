package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectRouterGrantAssociations = `{
  "block": {
    "attributes": {
      "associations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "ecr_id": "string",
              "grant_id": "string",
              "id": "string",
              "instance_id": "string",
              "instance_owner_bid": "string",
              "instance_owner_id": "string",
              "instance_region_id": "string",
              "instance_type": "string",
              "modify_time": "string",
              "owner_id": "string",
              "status": "string"
            }
          ]
        ]
      },
      "caller_type": {
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
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudExpressConnectRouterGrantAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectRouterGrantAssociations), &result)
	return &result
}
