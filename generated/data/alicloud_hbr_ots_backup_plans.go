package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrOtsBackupPlans = `{
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
      "plan_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "plan_name": {
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
              "created_time": "string",
              "disabled": "bool",
              "id": "string",
              "ots_backup_plan_id": "string",
              "ots_backup_plan_name": "string",
              "ots_detail": [
                "set",
                [
                  "object",
                  {
                    "table_names": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "retention": "string",
              "schedule": "string",
              "source_type": "string",
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

func AlicloudHbrOtsBackupPlansSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrOtsBackupPlans), &result)
	return &result
}
