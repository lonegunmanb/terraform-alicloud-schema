package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectRouterVpcAssociations = `{
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
              "create_time": "string",
              "ecr_id": "string",
              "id": "string",
              "modify_time": "string",
              "status": "string",
              "vpc_id": "string",
              "vpc_owner_id": "string"
            }
          ]
        ]
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
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudExpressConnectRouterVpcAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectRouterVpcAssociations), &result)
	return &result
}
