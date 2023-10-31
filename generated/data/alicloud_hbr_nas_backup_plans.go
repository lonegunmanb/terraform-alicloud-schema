package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrNasBackupPlans = `{
  "block": {
    "attributes": {
      "file_system_id": {
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
              "disabled": "bool",
              "file_system_id": "string",
              "id": "string",
              "nas_backup_plan_id": "string",
              "nas_backup_plan_name": "string",
              "options": "string",
              "path": [
                "list",
                "string"
              ],
              "retention": "string",
              "schedule": "string",
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

func AlicloudHbrNasBackupPlansSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrNasBackupPlans), &result)
	return &result
}
