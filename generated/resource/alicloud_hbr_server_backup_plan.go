package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrServerBackupPlan = `{
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
      "disabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ecs_server_backup_plan_name": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retention": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "schedule": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "detail": {
        "block": {
          "attributes": {
            "app_consistent": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "destination_region_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_retention": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_id_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "do_copy": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_fs_freeze": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "post_script_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pre_script_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "snapshot_group": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrServerBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrServerBackupPlan), &result)
	return &result
}
