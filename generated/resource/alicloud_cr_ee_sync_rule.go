package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCrEeSyncRule = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "repo_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sync_direction": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sync_scope": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tag_filter": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_namespace_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_repo_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCrEeSyncRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCrEeSyncRule), &result)
	return &result
}
