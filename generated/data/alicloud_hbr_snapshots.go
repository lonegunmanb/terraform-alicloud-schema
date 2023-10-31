package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrSnapshots = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "complete_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "complete_time_checker": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "file_system_id": {
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
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshots": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "actual_bytes": "string",
              "actual_items": "string",
              "backup_type": "string",
              "bucket": "string",
              "bytes_done": "string",
              "bytes_total": "string",
              "client_id": "string",
              "complete_time": "string",
              "create_time": "string",
              "created_time": "string",
              "error_file": "string",
              "file_system_id": "string",
              "id": "string",
              "instance_id": "string",
              "items_done": "string",
              "items_total": "string",
              "job_id": "string",
              "parent_snapshot_hash": "string",
              "path": "string",
              "prefix": "string",
              "retention": "string",
              "snapshot_hash": "string",
              "snapshot_id": "string",
              "source_type": "string",
              "start_time": "string",
              "status": "string",
              "updated_time": "string"
            }
          ]
        ]
      },
      "source_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrSnapshotsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrSnapshots), &result)
	return &result
}
