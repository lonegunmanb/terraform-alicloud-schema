package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCrArtifactLifecycleRules = `{
  "block": {
    "attributes": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "artifact_lifecycle_rule_id": "string",
              "auto": "bool",
              "create_time": "number",
              "enable_delete_tag": "bool",
              "enable_delete_untagged_manifest": "bool",
              "id": "string",
              "instance_id": "string",
              "modified_time": "number",
              "namespace_name": "string",
              "repo_name": "string",
              "retention_tag_count": "number",
              "schedule_time": "string",
              "scope": "string",
              "tag_regexp": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCrArtifactLifecycleRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCrArtifactLifecycleRules), &result)
	return &result
}
