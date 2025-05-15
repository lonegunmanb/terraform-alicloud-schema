package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAmqpExchange = `{
  "block": {
    "attributes": {
      "alternate_exchange": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_delete_state": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "exchange_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "exchange_type": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "internal": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "virtual_host_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "x_delayed_type": {
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

func AlicloudAmqpExchangeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAmqpExchange), &result)
	return &result
}
