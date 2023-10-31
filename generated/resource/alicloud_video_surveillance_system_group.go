package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVideoSurveillanceSystemGroup = `{
  "block": {
    "attributes": {
      "callback": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "capture_image": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "capture_interval": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "capture_oss_bucket": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capture_oss_path": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capture_video": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "in_protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lazy_pull": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "out_protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "play_domain": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "push_domain": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVideoSurveillanceSystemGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVideoSurveillanceSystemGroup), &result)
	return &result
}
