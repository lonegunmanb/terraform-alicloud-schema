package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOtsTable = `{
  "block": {
    "attributes": {
      "allow_update": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deviation_cell_version_in_sec": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_sse": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_version": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "sse_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sse_key_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sse_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "time_to_live": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "block_types": {
      "defined_column": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 32,
        "nesting_mode": "list"
      },
      "primary_key": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 4,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOtsTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOtsTable), &result)
	return &result
}
