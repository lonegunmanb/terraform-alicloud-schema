package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCdnIpInfo = `{
  "block": {
    "attributes": {
      "cdn_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "isp": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "isp_ename": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region_ename": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCdnIpInfoSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCdnIpInfo), &result)
	return &result
}
