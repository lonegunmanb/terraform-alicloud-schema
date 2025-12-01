package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcv3Functions = `{
  "block": {
    "attributes": {
      "functions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code_size": "number",
              "cpu": "number",
              "create_time": "string",
              "custom_container_config": [
                "list",
                [
                  "object",
                  {
                    "acceleration_info": [
                      "list",
                      [
                        "object",
                        {
                          "status": "string"
                        }
                      ]
                    ],
                    "acceleration_type": "string",
                    "acr_instance_id": "string",
                    "command": [
                      "list",
                      "string"
                    ],
                    "entrypoint": [
                      "list",
                      "string"
                    ],
                    "health_check_config": [
                      "list",
                      [
                        "object",
                        {
                          "failure_threshold": "number",
                          "http_get_url": "string",
                          "initial_delay_seconds": "number",
                          "period_seconds": "number",
                          "success_threshold": "number",
                          "timeout_seconds": "number"
                        }
                      ]
                    ],
                    "image": "string",
                    "port": "number",
                    "resolved_image_uri": "string"
                  }
                ]
              ],
              "custom_dns": [
                "list",
                [
                  "object",
                  {
                    "dns_options": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "name_servers": [
                      "list",
                      "string"
                    ],
                    "searches": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "custom_runtime_config": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "command": [
                      "list",
                      "string"
                    ],
                    "health_check_config": [
                      "list",
                      [
                        "object",
                        {
                          "failure_threshold": "number",
                          "http_get_url": "string",
                          "initial_delay_seconds": "number",
                          "period_seconds": "number",
                          "success_threshold": "number",
                          "timeout_seconds": "number"
                        }
                      ]
                    ],
                    "port": "number"
                  }
                ]
              ],
              "description": "string",
              "disk_size": "number",
              "environment_variables": [
                "map",
                "string"
              ],
              "function_arn": "string",
              "function_id": "string",
              "function_name": "string",
              "gpu_config": [
                "list",
                [
                  "object",
                  {
                    "gpu_memory_size": "number",
                    "gpu_type": "string"
                  }
                ]
              ],
              "handler": "string",
              "id": "string",
              "idle_timeout": "number",
              "instance_concurrency": "number",
              "instance_isolation_mode": "string",
              "instance_lifecycle_config": [
                "list",
                [
                  "object",
                  {
                    "initializer": [
                      "list",
                      [
                        "object",
                        {
                          "command": [
                            "list",
                            "string"
                          ],
                          "handler": "string",
                          "timeout": "number"
                        }
                      ]
                    ],
                    "pre_stop": [
                      "list",
                      [
                        "object",
                        {
                          "handler": "string",
                          "timeout": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "internet_access": "bool",
              "invocation_restriction": [
                "list",
                [
                  "object",
                  {
                    "disable": "bool",
                    "last_modified_time": "string",
                    "reason": "string"
                  }
                ]
              ],
              "last_modified_time": "string",
              "last_update_status": "string",
              "last_update_status_reason": "string",
              "last_update_status_reason_code": "string",
              "layers": [
                "list",
                "string"
              ],
              "log_config": [
                "list",
                [
                  "object",
                  {
                    "enable_instance_metrics": "bool",
                    "enable_request_metrics": "bool",
                    "log_begin_rule": "string",
                    "logstore": "string",
                    "project": "string"
                  }
                ]
              ],
              "memory_size": "number",
              "nas_config": [
                "list",
                [
                  "object",
                  {
                    "group_id": "number",
                    "mount_points": [
                      "list",
                      [
                        "object",
                        {
                          "enable_tls": "bool",
                          "mount_dir": "string",
                          "server_addr": "string"
                        }
                      ]
                    ],
                    "user_id": "number"
                  }
                ]
              ],
              "oss_mount_config": [
                "list",
                [
                  "object",
                  {
                    "mount_points": [
                      "list",
                      [
                        "object",
                        {
                          "bucket_name": "string",
                          "bucket_path": "string",
                          "endpoint": "string",
                          "mount_dir": "string",
                          "read_only": "bool"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "resource_group_id": "string",
              "role": "string",
              "runtime": "string",
              "session_affinity": "string",
              "session_affinity_config": "string",
              "state": "string",
              "state_reason": "string",
              "state_reason_code": "string",
              "tags": [
                "map",
                "string"
              ],
              "timeout": "number",
              "tracing_config": [
                "list",
                [
                  "object",
                  {
                    "params": [
                      "map",
                      "string"
                    ],
                    "type": "string"
                  }
                ]
              ],
              "vpc_config": [
                "list",
                [
                  "object",
                  {
                    "security_group_id": "string",
                    "vpc_id": "string",
                    "vswitch_ids": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
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
      "prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudFcv3FunctionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcv3Functions), &result)
	return &result
}
