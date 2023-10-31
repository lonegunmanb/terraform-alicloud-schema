package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMnsQueue = `{
  "block": {
    "attributes": {
      "delay_seconds": {
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
      "maximum_message_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "message_retention_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "polling_wait_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "visibility_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMnsQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMnsQueue), &result)
	return &result
}
