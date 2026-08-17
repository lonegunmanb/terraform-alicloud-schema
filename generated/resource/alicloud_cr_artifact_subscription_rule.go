package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCrArtifactSubscriptionRule = `{
  "block": {
    "attributes": {
      "accelerate": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "artifact_subscription_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
        "type": "string"
      },
      "namespace_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "override": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "platform": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "repo_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_namespace_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_provider": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_repo_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tag_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "tag_regexp": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudCrArtifactSubscriptionRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCrArtifactSubscriptionRule), &result)
	return &result
}
