package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSchedulerxAppGroup = `{
  "block": {
    "attributes": {
      "app_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "app_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "app_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_jobs": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_log": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "group_id": {
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
      "max_concurrency": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_jobs": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "monitor_config_json": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "monitor_contacts_json": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule_busy_workers": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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

func AlicloudSchedulerxAppGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSchedulerxAppGroup), &result)
	return &result
}
