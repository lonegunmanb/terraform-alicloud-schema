package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRdsCrossRegionBackups = `{
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
              "backup_end_time": "string",
              "backup_method": "string",
              "backup_set_scale": "string",
              "backup_set_status": "number",
              "backup_start_time": "string",
              "backup_type": "string",
              "category": "string",
              "consistent_time": "string",
              "cross_backup_download_link": "string",
              "cross_backup_id": "string",
              "cross_backup_region": "string",
              "cross_backup_set_file": "string",
              "cross_backup_set_location": "string",
              "cross_backup_set_size": "number",
              "db_instance_storage_type": "string",
              "engine": "string",
              "engine_version": "string",
              "id": "string",
              "instance_id": "number",
              "recovery_begin_time": "string",
              "recovery_end_time": "string",
              "restore_regions": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "cross_backup_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cross_backup_region": {
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
      "resource_group_id": {
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

func AlicloudRdsCrossRegionBackupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRdsCrossRegionBackups), &result)
	return &result
}
