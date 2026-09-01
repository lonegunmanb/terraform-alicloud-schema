package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectRouterExpressConnectRouters = `{
  "block": {
    "attributes": {
      "ecr_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecr_name": {
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
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alibaba_side_asn": "string",
              "biz_status": "string",
              "create_time": "string",
              "description": "string",
              "ecr_id": "string",
              "ecr_name": "string",
              "id": "string",
              "modify_time": "string",
              "owner_id": "string",
              "resource_group_id": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
        ]
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudExpressConnectRouterExpressConnectRoutersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectRouterExpressConnectRouters), &result)
	return &result
}
