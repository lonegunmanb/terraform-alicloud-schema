package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenInstanceGrant = `{
  "block": {
    "attributes": {
      "cen_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cen_owner_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "child_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
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

func AlicloudCenInstanceGrantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenInstanceGrant), &result)
	return &result
}
