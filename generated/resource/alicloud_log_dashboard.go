package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudLogDashboard = `{
  "block": {
    "attributes": {
      "attribute": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "char_list": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dashboard_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
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
      "project_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudLogDashboardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudLogDashboard), &result)
	return &result
}
