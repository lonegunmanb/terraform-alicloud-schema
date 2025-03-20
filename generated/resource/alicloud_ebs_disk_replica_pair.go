package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEbsDiskReplicaPair = `{
  "block": {
    "attributes": {
      "bandwidth": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_disk_id": {
        "description_kind": "plain",
        "required": true,
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
      "disk_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disk_replica_pair_name": {
        "computed": true,
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
      "pair_name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "period_unit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "source_zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
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

func AlicloudEbsDiskReplicaPairSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEbsDiskReplicaPair), &result)
	return &result
}
