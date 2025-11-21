package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudWafv3DefenseTemplate = `{
  "block": {
    "attributes": {
      "defense_scene": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "defense_template_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "defense_template_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_groups": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "resource_manager_resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resources": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "status": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_origin": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_type": {
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

func AlicloudWafv3DefenseTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudWafv3DefenseTemplate), &result)
	return &result
}
