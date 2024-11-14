package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcv3LayerVersion = `{
  "block": {
    "attributes": {
      "acl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "code_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compatible_runtime": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
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
      "layer_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "layer_version_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "license": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "code": {
        "block": {
          "attributes": {
            "checksum": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "oss_bucket_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "oss_object_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zip_file": {
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

func AlicloudFcv3LayerVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcv3LayerVersion), &result)
	return &result
}
