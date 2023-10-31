package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrBackupJobs = `{
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
              "back_job_name": "string",
              "backup_job_id": "string",
              "backup_type": "string",
              "bucket": "string",
              "bytes_done": "string",
              "bytes_total": "string",
              "complete_time": "string",
              "create_time": "string",
              "cross_account_role_name": "string",
              "cross_account_type": "string",
              "cross_account_user_id": "number",
              "error_message": "string",
              "exclude": "string",
              "file_system_id": "string",
              "id": "string",
              "include": "string",
              "instance_id": "string",
              "items_done": "string",
              "items_total": "string",
              "nas_create_time": "string",
              "ots_detail": [
                "list",
                [
                  "object",
                  {
                    "table_names": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "paths": [
                "list",
                "string"
              ],
              "plan_id": "string",
              "prefix": "string",
              "progress": "string",
              "source_type": "string",
              "start_time": "string",
              "status": "string",
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
      "sort_direction": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "operator": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "values": {
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrBackupJobsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrBackupJobs), &result)
	return &result
}
