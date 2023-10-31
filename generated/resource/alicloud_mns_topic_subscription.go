package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMnsTopicSubscription = `{
  "block": {
    "attributes": {
      "endpoint": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filter_tag": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notify_content_format": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notify_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "topic_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMnsTopicSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMnsTopicSubscription), &result)
	return &result
}
