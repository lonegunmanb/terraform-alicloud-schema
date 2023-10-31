package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNasAccessRule = `{
  "block": {
    "attributes": {
      "access_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "access_rule_id": {
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
      "priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "rw_access_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_cidr_ip": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_access_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudNasAccessRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNasAccessRule), &result)
	return &result
}
