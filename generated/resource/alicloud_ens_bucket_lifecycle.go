package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEnsBucketLifecycle = `{
  "block": {
    "attributes": {
      "allow_same_action_overlap": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bucket_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "expiration_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
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

func AlicloudEnsBucketLifecycleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEnsBucketLifecycle), &result)
	return &result
}
