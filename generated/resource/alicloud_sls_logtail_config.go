package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsLogtailConfig = `{
  "block": {
    "attributes": {
      "create_time": {
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
      "input_detail": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "input_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modify_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "log_sample": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logtail_config_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "output_detail": {
        "block": {
          "attributes": {
            "endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "logstore_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AlicloudSlsLogtailConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsLogtailConfig), &result)
	return &result
}
