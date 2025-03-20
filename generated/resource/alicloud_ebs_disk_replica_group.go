package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEbsDiskReplicaGroup = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disk_replica_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_name": {
        "computed": true,
        "deprecated": true,
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
      "one_shot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "pair_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reverse_replicate": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "rpo": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "source_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
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
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEbsDiskReplicaGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEbsDiskReplicaGroup), &result)
	return &result
}
