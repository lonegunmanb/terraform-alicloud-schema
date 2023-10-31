package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrOtsSnapshots = `{
  "block": {
    "attributes": {
      "end_time": {
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
      "output_file": {
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
              "backup_type": "string",
              "bytes_total": "string",
              "complete_time": "string",
              "create_time": "string",
              "created_time": "string",
              "id": "string",
              "instance_name": "string",
              "job_id": "string",
              "parent_snapshot_hash": "string",
              "range_end": "string",
              "range_start": "string",
              "retention": "string",
              "snapshot_hash": "string",
              "snapshot_id": "string",
              "source_type": "string",
              "start_time": "string",
              "status": "string",
              "table_name": "string",
              "updated_time": "string",
              "vault_id": "string"
            }
          ]
        ]
      },
      "start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrOtsSnapshotsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrOtsSnapshots), &result)
	return &result
}
