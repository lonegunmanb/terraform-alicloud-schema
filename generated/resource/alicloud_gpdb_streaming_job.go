package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGpdbStreamingJob = `{
  "block": {
    "attributes": {
      "account": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "consistency": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dest_columns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "dest_database": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dest_schema": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dest_table": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "error_limit_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "fallback_offset": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_name": {
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
      "job_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "match_columns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "src_columns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "try_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "update_columns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "write_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AlicloudGpdbStreamingJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGpdbStreamingJob), &result)
	return &result
}
