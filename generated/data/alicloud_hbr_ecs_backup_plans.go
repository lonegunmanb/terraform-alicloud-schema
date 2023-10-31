package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrEcsBackupPlans = `{
  "block": {
    "attributes": {
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
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
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
      "plans": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_type": "string",
              "create_time": "string",
              "created_time": "string",
              "detail": "string",
              "disabled": "bool",
              "ecs_backup_plan_id": "string",
              "ecs_backup_plan_name": "string",
              "exclude": "string",
              "id": "string",
              "include": "string",
              "instance_id": "string",
              "options": "string",
              "path": [
                "list",
                "string"
              ],
              "retention": "string",
              "schedule": "string",
              "source_type": "string",
              "speed_limit": "string",
              "updated_time": "string",
              "vault_id": "string"
            }
          ]
        ]
      },
      "vault_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrEcsBackupPlansSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrEcsBackupPlans), &result)
	return &result
}
