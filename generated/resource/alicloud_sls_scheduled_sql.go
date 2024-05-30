package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsScheduledSql = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
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
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scheduled_sql_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "schedule": {
        "block": {
          "attributes": {
            "cron_expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delay": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "run_immediately": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "time_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "scheduled_sql_configuration": {
        "block": {
          "attributes": {
            "data_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dest_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dest_logstore": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dest_project": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dest_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "from_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "from_time_expr": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_retries": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_run_time_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "resource_pool": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "script": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_logstore": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sql_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "to_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "to_time_expr": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
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

func AlicloudSlsScheduledSqlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsScheduledSql), &result)
	return &result
}
