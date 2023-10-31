package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudQuickBiUser = `{
  "block": {
    "attributes": {
      "account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "admin_user": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "auth_admin_user": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nick_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudQuickBiUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudQuickBiUser), &result)
	return &result
}
