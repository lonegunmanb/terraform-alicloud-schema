package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrNasBackupPlan = `{
  "block": {
    "attributes": {
      "backup_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
      "file_system_id": {
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
      "nas_backup_plan_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "options": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "retention": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrNasBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrNasBackupPlan), &result)
	return &result
}
