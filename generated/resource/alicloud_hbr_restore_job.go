package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrRestoreJob = `{
  "block": {
    "attributes": {
      "cross_account_role_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cross_account_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cross_account_user_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "exclude": {
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
      "include": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "options": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_job_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "snapshot_hash": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "snapshot_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_client_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_create_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_data_source_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_file_system_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_instance_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_table_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "udm_detail": {
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
    "block_types": {
      "ots_detail": {
        "block": {
          "attributes": {
            "overwrite_existing": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
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

func AlicloudHbrRestoreJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrRestoreJob), &result)
	return &result
}
