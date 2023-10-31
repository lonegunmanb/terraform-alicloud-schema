package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudauthFaceConfig = `{
  "block": {
    "attributes": {
      "biz_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "biz_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gmt_modified": {
        "computed": true,
        "description_kind": "plain",
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

func AlicloudCloudauthFaceConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudauthFaceConfig), &result)
	return &result
}
