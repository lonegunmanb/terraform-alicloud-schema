package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMessageServiceQueues = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "queue_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "queues": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "active_messages": "number",
              "create_time": "number",
              "delay_messages": "number",
              "delay_seconds": "number",
              "id": "string",
              "inactive_messages": "number",
              "last_modify_time": "number",
              "logging_enabled": "bool",
              "maximum_message_size": "number",
              "message_retention_period": "number",
              "polling_wait_seconds": "number",
              "queue_internal_url": "string",
              "queue_name": "string",
              "queue_url": "string",
              "visibility_timeout": "number"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMessageServiceQueuesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMessageServiceQueues), &result)
	return &result
}
