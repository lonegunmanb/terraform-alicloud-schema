package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsHybridMonitorSlsTasks = `{
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
      "keyword": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tasks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attach_labels": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "value": "string"
                  }
                ]
              ],
              "collect_interval": "number",
              "collect_target_endpoint": "string",
              "collect_target_path": "string",
              "collect_target_type": "string",
              "collect_timout": "number",
              "create_time": "string",
              "description": "string",
              "extra_info": "string",
              "group_id": "string",
              "hybrid_monitor_sls_task_id": "string",
              "id": "string",
              "instances": [
                "list",
                "string"
              ],
              "log_file_path": "string",
              "log_process": "string",
              "log_sample": "string",
              "log_split": "string",
              "match_express": [
                "list",
                [
                  "object",
                  {
                    "function": "string",
                    "name": "string",
                    "value": "string"
                  }
                ]
              ],
              "match_express_relation": "string",
              "namespace": "string",
              "network_type": "string",
              "sls_process": "string",
              "sls_process_config": [
                "list",
                [
                  "object",
                  {
                    "express": [
                      "list",
                      [
                        "object",
                        {
                          "alias": "string",
                          "express": "string"
                        }
                      ]
                    ],
                    "filter": [
                      "list",
                      [
                        "object",
                        {
                          "filters": [
                            "list",
                            [
                              "object",
                              {
                                "operator": "string",
                                "sls_key_name": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "relation": "string"
                        }
                      ]
                    ],
                    "group_by": [
                      "list",
                      [
                        "object",
                        {
                          "alias": "string",
                          "sls_key_name": "string"
                        }
                      ]
                    ],
                    "statistics": [
                      "list",
                      [
                        "object",
                        {
                          "alias": "string",
                          "function": "string",
                          "parameter_one": "string",
                          "parameter_two": "string",
                          "sls_key_name": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "task_name": "string",
              "task_type": "string",
              "upload_region": "string",
              "yarm_config": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsHybridMonitorSlsTasksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsHybridMonitorSlsTasks), &result)
	return &result
}
