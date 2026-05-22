package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRosStackInstances = `{
  "block": {
    "attributes": {
      "account_ids": {
        "description": "List of target account IDs for self-managed permissions. Maximum 50 accounts.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "deployment_options": {
        "description": "Deployment options for service-managed permissions. Only supports 'IgnoreExisting'.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "disable_rollback": {
        "description": "When creating stack instances fails, whether to disable the rollback policy.",
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
      "operation_description": {
        "description": "Description of the stack instances operation. Length: 1-256 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_ids": {
        "description": "List of target region IDs. Maximum 20 regions.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "stack_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stack_instances": {
        "computed": true,
        "description": "List of stack instances with their latest operation tracking ID.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_id": "string",
              "drift_detection_time": "string",
              "last_operation_id": "string",
              "rd_folder_id": "string",
              "region_id": "string",
              "stack_drift_status": "string",
              "stack_group_id": "string",
              "stack_group_name": "string",
              "stack_id": "string",
              "status": "string",
              "status_reason": "string"
            }
          ]
        ]
      },
      "timeout_in_minutes": {
        "description": "The amount of time that can elapse before the stack operation status is set to TIMED_OUT.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "deployment_targets": {
        "block": {
          "attributes": {
            "account_ids": {
              "description": "List of target account IDs for service-managed permissions.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "rd_folder_ids": {
              "description": "List of Resource Directory folder IDs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Configuration block defining deployment targets. Conflicts with account_ids.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "operation_preferences": {
        "block": {
          "attributes": {
            "failure_tolerance_count": {
              "description": "Number of failures tolerated per region. Range: 0-20.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "failure_tolerance_percentage": {
              "description": "Percentage of failures tolerated per region. Range: 0-100.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrent_count": {
              "description": "Maximum number of concurrent operations per region. Range: 1-20.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrent_percentage": {
              "description": "Maximum percentage of concurrent targets per region. Range: 1-100.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "region_concurrency_type": {
              "description": "Concurrency type for regions. Valid values: SEQUENTIAL, PARALLEL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Preferences for how the operation is performed across multiple accounts and regions.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "parameter_overrides": {
        "block": {
          "attributes": {
            "parameter_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "parameter_value": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "Parameters to override in the stack instances.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudRosStackInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRosStackInstances), &result)
	return &result
}
