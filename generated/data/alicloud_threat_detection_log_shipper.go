package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionLogShipper = `{
  "block": {
    "attributes": {
      "auth_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "buy_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable": {
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
      "open_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sls_project_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sls_service_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionLogShipperSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionLogShipper), &result)
	return &result
}
