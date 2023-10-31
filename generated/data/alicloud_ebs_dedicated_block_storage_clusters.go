package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEbsDedicatedBlockStorageClusters = `{
  "block": {
    "attributes": {
      "clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "available_capacity": "string",
              "category": "string",
              "create_time": "string",
              "dedicated_block_storage_cluster_id": "string",
              "dedicated_block_storage_cluster_name": "string",
              "delivery_capacity": "string",
              "description": "string",
              "expired_time": "string",
              "id": "string",
              "performance_level": "string",
              "resource_group_id": "string",
              "status": "string",
              "supported_category": "string",
              "total_capacity": "string",
              "type": "string",
              "used_capacity": "string",
              "zone_id": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEbsDedicatedBlockStorageClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEbsDedicatedBlockStorageClusters), &result)
	return &result
}
