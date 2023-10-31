package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcFunction = `{
  "block": {
    "attributes": {
      "ca_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "code_checksum": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_variables": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "filename": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "function_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "handler": {
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
      "initialization_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "initializer": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_concurrency": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "layers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "memory_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oss_bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oss_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "custom_container_config": {
        "block": {
          "attributes": {
            "args": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "command": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudFcFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcFunction), &result)
	return &result
}
