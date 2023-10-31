package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGpdbDbInstancePlans = `{
  "block": {
    "attributes": {
      "db_instance_id": {
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
      "plan_schedule_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "plan_type": {
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
              "db_instance_plan_name": "string",
              "id": "string",
              "plan_config": [
                "list",
                [
                  "object",
                  {
                    "pause": [
                      "list",
                      [
                        "object",
                        {
                          "execute_time": "string",
                          "plan_cron_time": "string",
                          "plan_task_status": "string"
                        }
                      ]
                    ],
                    "resume": [
                      "list",
                      [
                        "object",
                        {
                          "execute_time": "string",
                          "plan_cron_time": "string",
                          "plan_task_status": "string"
                        }
                      ]
                    ],
                    "scale_in": [
                      "list",
                      [
                        "object",
                        {
                          "execute_time": "string",
                          "plan_cron_time": "string",
                          "plan_task_status": "string",
                          "segment_node_num": "string"
                        }
                      ]
                    ],
                    "scale_out": [
                      "list",
                      [
                        "object",
                        {
                          "execute_time": "string",
                          "plan_cron_time": "string",
                          "plan_task_status": "string",
                          "segment_node_num": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "plan_desc": "string",
              "plan_end_date": "string",
              "plan_id": "string",
              "plan_schedule_type": "string",
              "plan_start_date": "string",
              "plan_type": "string",
              "status": "string"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudGpdbDbInstancePlansSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGpdbDbInstancePlans), &result)
	return &result
}
