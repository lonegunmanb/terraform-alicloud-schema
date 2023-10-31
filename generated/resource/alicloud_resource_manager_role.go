package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudResourceManagerRole = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assume_role_policy_document": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "max_session_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "role_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudResourceManagerRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudResourceManagerRole), &result)
	return &result
}
