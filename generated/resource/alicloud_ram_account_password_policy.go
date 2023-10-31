package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRamAccountPasswordPolicy = `{
  "block": {
    "attributes": {
      "hard_expiry": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_login_attempts": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_password_age": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "minimum_password_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "password_reuse_prevention": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "require_lowercase_characters": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "require_numbers": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "require_symbols": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "require_uppercase_characters": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudRamAccountPasswordPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRamAccountPasswordPolicy), &result)
	return &result
}
