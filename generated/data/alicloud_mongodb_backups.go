package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMongodbBackups = `{
  "block": {
    "attributes": {
      "backup_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_db_names": "string",
              "backup_download_url": "string",
              "backup_end_time": "string",
              "backup_id": "string",
              "backup_intranet_download_url": "string",
              "backup_job_id": "string",
              "backup_method": "string",
              "backup_mode": "string",
              "backup_size": "number",
              "backup_start_time": "string",
              "backup_type": "string",
              "id": "string",
              "status": "string"
            }
          ]
        ]
      },
      "db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
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

func AlicloudMongodbBackupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMongodbBackups), &result)
	return &result
}
