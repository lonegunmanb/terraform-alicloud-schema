package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDmsEnterpriseUser = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_execute_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_result_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "mobile": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nick_name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tid": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "uid": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDmsEnterpriseUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDmsEnterpriseUser), &result)
	return &result
}
