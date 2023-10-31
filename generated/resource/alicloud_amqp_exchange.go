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
