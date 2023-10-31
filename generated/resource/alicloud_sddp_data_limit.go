package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSddpDataLimit = `{
  "block": {
    "attributes": {
      "audit_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "engine_type": {
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
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_store_day": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "parent_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSddpDataLimitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSddpDataLimit), &result)
	return &result
}
