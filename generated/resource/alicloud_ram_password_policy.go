package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRamPasswordPolicy = `{
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
      "max_login_attemps": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_password_age": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "minimum_password_different_character": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "minimum_password_length": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "password_not_contain_user_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudRamPasswordPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRamPasswordPolicy), &result)
	return &result
}
