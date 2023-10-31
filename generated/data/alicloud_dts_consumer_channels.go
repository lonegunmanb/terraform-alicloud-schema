package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDtsConsumerChannels = `{
  "block": {
    "attributes": {
      "channels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "consumer_group_id": "string",
              "consumer_group_name": "string",
              "consumer_group_user_name": "string",
              "consumption_checkpoint": "string",
              "id": "string",
              "message_delay": "number",
              "unconsumed_data": "number"
            }
          ]
        ]
      },
      "dts_instance_id": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDtsConsumerChannelsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDtsConsumerChannels), &result)
	return &result
}
