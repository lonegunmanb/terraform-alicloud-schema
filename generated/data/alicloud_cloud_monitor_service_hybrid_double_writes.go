package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudMonitorServiceHybridDoubleWrites = `{
  "block": {
    "attributes": {
      "hybrid_double_writes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "namespace": "string",
              "source_namespace": "string",
              "source_user_id": "string",
              "user_id": "string"
            }
          ]
        ]
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
      "namespace": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_namespace": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_user_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudMonitorServiceHybridDoubleWritesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudMonitorServiceHybridDoubleWrites), &result)
	return &result
}
