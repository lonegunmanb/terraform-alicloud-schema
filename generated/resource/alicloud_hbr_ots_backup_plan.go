package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrOtsBackupPlan = `{
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
      "disabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ots_backup_plan_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retention": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vault_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "ots_detail": {
        "block": {
          "attributes": {
            "table_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "rules": {
        "block": {
          "attributes": {
            "backup_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "retention": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrOtsBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrOtsBackupPlan), &result)
	return &result
}
