package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDfsFileSystem = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_redundancy_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dedicated_cluster_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "file_system_name": {
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
      "partition_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "protocol_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provisioned_throughput_in_mi_bps": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "space_capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "storage_set_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "throughput_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AlicloudDfsFileSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDfsFileSystem), &result)
	return &result
}
