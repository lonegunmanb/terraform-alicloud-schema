package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEbsDiskReplicaPairs = `{
  "block": {
    "attributes": {
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
      "pairs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bandwidth": "string",
              "description": "string",
              "destination_disk_id": "string",
              "destination_region_id": "string",
              "destination_zone_id": "string",
              "disk_id": "string",
              "id": "string",
              "pair_name": "string",
              "payment_type": "string",
              "replica_pair_id": "string",
              "resource_group_id": "string",
              "rpo": "string",
              "source_zone_id": "string",
              "status": "string"
            }
          ]
        ]
      },
      "replica_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEbsDiskReplicaPairsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEbsDiskReplicaPairs), &result)
	return &result
}
