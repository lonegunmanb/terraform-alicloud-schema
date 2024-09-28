package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGpdbDataBackups = `{
  "block": {
    "attributes": {
      "backup_mode": {
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
              "backup_end_time": "string",
              "backup_end_time_local": "string",
              "backup_method": "string",
              "backup_mode": "string",
              "backup_set_id": "string",
              "backup_size": "number",
              "backup_start_time": "string",
              "backup_start_time_local": "string",
              "bakset_name": "string",
              "consistent_time": "number",
              "data_type": "string",
              "db_instance_id": "string",
              "status": "string"
            }
          ]
        ]
      },
      "data_backup_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudGpdbDataBackupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGpdbDataBackups), &result)
	return &result
}
