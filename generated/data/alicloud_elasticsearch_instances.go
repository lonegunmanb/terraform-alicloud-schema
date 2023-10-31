package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudElasticsearchInstances = `{
  "block": {
    "attributes": {
      "description_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "descriptions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "created_at": "string",
              "data_node_amount": "number",
              "data_node_disk_size": "number",
              "data_node_disk_type": "string",
              "data_node_spec": "string",
              "description": "string",
              "id": "string",
              "instance_charge_type": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "updated_at": "string",
              "version": "string",
              "vswitch_id": "string"
            }
          ]
        ]
      },
      "output_file": {
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
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudElasticsearchInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudElasticsearchInstances), &result)
	return &result
}
