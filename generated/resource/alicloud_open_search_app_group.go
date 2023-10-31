package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOpenSearchAppGroup = `{
  "block": {
    "attributes": {
      "app_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "charge_way": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "current_version": {
        "description_kind": "plain",
        "optional": true,
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "order_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "order": {
        "block": {
          "attributes": {
            "auto_renew": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "pricing_cycle": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "quota": {
        "block": {
          "attributes": {
            "compute_resource": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "doc_size": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "qps": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "spec": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOpenSearchAppGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOpenSearchAppGroup), &result)
	return &result
}
