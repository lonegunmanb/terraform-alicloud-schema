package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrRestoreJobs = `{
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
        "type": [
          "list",
          "string"
        ]
      },
      "jobs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "actual_bytes": "string",
              "actual_items": "string",
              "bytes_done": "string",
              "bytes_total": "string",
              "complete_time": "string",
              "create_time": "string",
              "error_file": "string",
              "error_message": "string",
              "expire_time": "string",
              "id": "string",
              "items_done": "string",
              "items_total": "string",
              "options": "string",
              "parent_id": "string",
              "progress": "number",
              "restore_job_id": "string",
              "restore_type": "string",
              "snapshot_hash": "string",
              "snapshot_id": "string",
              "source_type": "string",
              "start_time": "string",
              "status": "string",
              "target_bucket": "string",
              "target_client_id": "string",
              "target_create_time": "string",
              "target_data_source_id": "string",
              "target_file_system_id": "string",
              "target_instance_id": "string",
              "target_path": "string",
              "target_prefix": "string",
              "updated_time": "string",
              "vault_id": "string"
            }
          ]
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_id": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "restore_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_type": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "target_file_system_id": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "target_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "vault_id": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrRestoreJobsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrRestoreJobs), &result)
	return &result
}
