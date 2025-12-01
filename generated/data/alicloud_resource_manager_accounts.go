package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudResourceManagerAccounts = `{
  "block": {
    "attributes": {
      "accounts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_id": "string",
              "account_name": "string",
              "deletion_status": "string",
              "display_name": "string",
              "folder_id": "string",
              "id": "string",
              "join_method": "string",
              "join_time": "string",
              "modify_time": "string",
              "payer_account_id": "string",
              "resource_directory_id": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "type": "string"
            }
          ]
        ]
      },
      "enable_details": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudResourceManagerAccountsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudResourceManagerAccounts), &result)
	return &result
}
