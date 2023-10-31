package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSimpleApplicationServerPlans = `{
  "block": {
    "attributes": {
      "bandwidth": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "core": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "disk_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "flow": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "memory": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "plans": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bandwidth": "number",
              "core": "number",
              "disk_size": "number",
              "flow": "number",
              "id": "string",
              "memory": "number",
              "plan_id": "string",
              "support_platform": "string"
            }
          ]
        ]
      },
      "platform": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSimpleApplicationServerPlansSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSimpleApplicationServerPlans), &result)
	return &result
}
