package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOosTemplate = `{
  "block": {
    "attributes": {
      "auto_delete_executions": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "content": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "has_trigger": {
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
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "share_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "template_format": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOosTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOosTemplate), &result)
	return &result
}
