package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAmqpQueue = `{
  "block": {
    "attributes": {
      "auto_delete_state": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_expire_state": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dead_letter_exchange": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dead_letter_routing_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclusive_state": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maximum_priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "message_ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "queue_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func AlicloudAmqpQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAmqpQueue), &result)
	return &result
}
