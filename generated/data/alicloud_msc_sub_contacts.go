package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMscSubContacts = `{
  "block": {
    "attributes": {
      "contacts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_uid": "string",
              "contact_id": "string",
              "contact_name": "string",
              "email": "string",
              "id": "string",
              "is_account": "bool",
              "is_obsolete": "bool",
              "is_verified_email": "bool",
              "is_verified_mobile": "bool",
              "last_email_verification_time_stamp": "string",
              "last_mobile_verification_time_stamp": "string",
              "mobile": "string",
              "position": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMscSubContactsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMscSubContacts), &result)
	return &result
}
