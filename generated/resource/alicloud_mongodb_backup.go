package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMongodbBackup = `{
  "block": {
    "attributes": {
      "backup_db_names": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_download_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_intranet_download_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_job_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_method": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_retention_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "backup_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "backup_start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_id": {
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
      "status": {
        "computed": true,
        "description_kind": "plain",
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

func AlicloudMongodbBackupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMongodbBackup), &result)
	return &result
}
