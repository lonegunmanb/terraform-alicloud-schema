package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDtsConsumerChannel = `{
  "block": {
    "attributes": {
      "consumer_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "consumer_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "consumer_group_password": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "consumer_group_user_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDtsConsumerChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDtsConsumerChannel), &result)
	return &result
}
