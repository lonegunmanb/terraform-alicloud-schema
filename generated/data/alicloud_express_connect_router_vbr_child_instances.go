package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectRouterVbrChildInstances = `{
  "block": {
    "attributes": {
      "child_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "child_instance_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "child_instance_type": {
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "child_instance_id": "string",
              "child_instance_owner_id": "string",
              "child_instance_region_id": "string",
              "child_instance_type": "string",
              "create_time": "string",
              "description": "string",
              "ecr_id": "string",
              "id": "string",
              "modify_time": "string",
              "status": "string"
            }
          ]
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudExpressConnectRouterVbrChildInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectRouterVbrChildInstances), &result)
	return &result
}
