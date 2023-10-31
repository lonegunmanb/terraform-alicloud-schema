package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrEcsBackupPlan = `{
  "block": {
    "attributes": {
      "backup_type": {
        "description_kind": "plain",
        "required": true,
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
      "detail": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ecs_backup_plan_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "instance_id": {
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
        "optional": true,
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
      "speed_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_paths": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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

func AlicloudHbrEcsBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrEcsBackupPlan), &result)
	return &result
}
