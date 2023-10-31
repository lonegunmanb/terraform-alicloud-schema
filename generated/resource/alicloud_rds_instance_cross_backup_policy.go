package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRdsInstanceCrossBackupPolicy = `{
  "block": {
    "attributes": {
      "backup_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_enabled_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cross_backup_region": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cross_backup_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lock_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "log_backup_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_backup_enabled_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "retent_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "retention": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AlicloudRdsInstanceCrossBackupPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRdsInstanceCrossBackupPolicy), &result)
	return &result
}
