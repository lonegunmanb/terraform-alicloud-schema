package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsLogtailPipelineConfig = `{
  "block": {
    "attributes": {
      "aggregators": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "map",
            "string"
          ]
        ]
      },
      "config_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "flushers": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          [
            "map",
            "string"
          ]
        ]
      },
      "globals": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "inputs": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          [
            "map",
            "string"
          ]
        ]
      },
      "log_sample": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "processors": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "map",
            "string"
          ]
        ]
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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

func AlicloudSlsLogtailPipelineConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsLogtailPipelineConfig), &result)
	return &result
}
