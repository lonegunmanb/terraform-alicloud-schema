package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCrArtifactLifecycleRule = `{
  "block": {
    "attributes": {
      "artifact_lifecycle_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
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
      "modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "namespace_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repo_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retention_tag_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "schedule_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tag_regexp": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCrArtifactLifecycleRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCrArtifactLifecycleRule), &result)
	return &result
}
