package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDbInstanceClassInfos = `{
  "block": {
    "attributes": {
      "commodity_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_id": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "order_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "infos": {
        "block": {
          "attributes": {
            "class_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "class_group": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cpu": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instruction_set_arch": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_connections": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_iombps": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_iops": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "memory_class": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reference_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDbInstanceClassInfosSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDbInstanceClassInfos), &result)
	return &result
}
