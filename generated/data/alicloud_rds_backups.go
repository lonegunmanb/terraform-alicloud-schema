package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRdsBackups = `{
  "block": {
    "attributes": {
      "backup_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_status": {
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
              "backup_download_url": "string",
              "backup_end_time": "string",
              "backup_id": "string",
              "backup_initiator": "string",
              "backup_intranet_download_url": "string",
              "backup_method": "string",
              "backup_mode": "string",
              "backup_size": "string",
              "backup_start_time": "string",
              "backup_status": "string",
              "backup_type": "string",
              "consistent_time": "string",
              "copy_only_backup": "string",
              "db_instance_id": "string",
              "encryption": "string",
              "host_instance_id": "string",
              "id": "string",
              "is_avail": "number",
              "meta_status": "string",
              "storage_class": "string",
              "store_status": "string"
            }
          ]
        ]
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

func AlicloudRdsBackupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRdsBackups), &result)
	return &result
}
