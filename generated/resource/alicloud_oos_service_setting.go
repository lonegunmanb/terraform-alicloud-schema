package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOosServiceSetting = `{
  "block": {
    "attributes": {
      "delivery_oss_bucket_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delivery_oss_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delivery_oss_key_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delivery_sls_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delivery_sls_project_name": {
        "description_kind": "plain",
        "optional": true,
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

func AlicloudOosServiceSettingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOosServiceSetting), &result)
	return &result
}
