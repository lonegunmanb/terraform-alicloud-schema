package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudResourceManagerAccountDeletionCheckTask = `{
  "block": {
    "attributes": {
      "abandon_able_checks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "check_id": "string",
              "check_name": "string",
              "description": "string"
            }
          ]
        ]
      },
      "account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "allow_delete": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "not_allow_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "check_id": "string",
              "check_name": "string",
              "description": "string"
            }
          ]
        ]
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

func AlicloudResourceManagerAccountDeletionCheckTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudResourceManagerAccountDeletionCheckTask), &result)
	return &result
}
