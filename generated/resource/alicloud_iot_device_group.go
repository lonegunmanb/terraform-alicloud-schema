package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudIotDeviceGroup = `{
  "block": {
    "attributes": {
      "group_desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_name": {
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
      "iot_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "super_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudIotDeviceGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudIotDeviceGroup), &result)
	return &result
}
