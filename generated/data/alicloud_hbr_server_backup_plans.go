package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrServerBackupPlans = `{
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
              "create_time": "string",
              "detail": [
                "list",
                [
                  "object",
                  {
                    "app_consistent": "bool",
                    "destination_region_id": "string",
                    "destination_retention": "number",
                    "disk_id_list": [
                      "list",
                      "string"
                    ],
                    "do_copy": "bool",
                    "enable_fs_freeze": "bool",
                    "post_script_path": "string",
                    "pre_script_path": "string",
                    "snapshot_group": "bool",
                    "timeout_in_seconds": "number"
                  }
                ]
              ],
              "disabled": "bool",
              "ecs_server_backup_plan_id": "string",
              "ecs_server_backup_plan_name": "string",
              "id": "string",
              "instance_id": "string",
              "retention": "string",
              "schedule": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "filters": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "values": {
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrServerBackupPlansSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrServerBackupPlans), &result)
	return &result
}
