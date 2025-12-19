package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGpdbSupabaseProject = `{
  "block": {
    "attributes": {
      "account_password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk_performance_level": {
        "computed": true,
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
      "project_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_spec": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_ip_list": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone_id": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudGpdbSupabaseProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGpdbSupabaseProject), &result)
	return &result
}
