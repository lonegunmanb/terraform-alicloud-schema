package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSaeIngress = `{
  "block": {
    "attributes": {
      "cert_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cert_ids": {
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
      "listener_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "listener_protocol": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balance_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "slb_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "default_rule": {
        "block": {
          "attributes": {
            "app_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "app_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "rules": {
        "block": {
          "attributes": {
            "app_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "app_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "backend_protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "domain": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rewrite_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSaeIngressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSaeIngress), &result)
	return &result
}
