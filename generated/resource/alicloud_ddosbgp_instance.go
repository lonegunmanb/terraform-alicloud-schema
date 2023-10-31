package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDdosbgpInstance = `{
  "block": {
    "attributes": {
      "bandwidth": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "base_bandwidth": {
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
      "ip_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "ip_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "normal_bandwidth": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDdosbgpInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDdosbgpInstance), &result)
	return &result
}
