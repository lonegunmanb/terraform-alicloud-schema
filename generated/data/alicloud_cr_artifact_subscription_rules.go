package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCrArtifactSubscriptionRules = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
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
      "namespace_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repo_name": {
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
              "accelerate": "bool",
              "artifact_subscription_rule_id": "string",
              "create_time": "string",
              "id": "string",
              "instance_id": "string",
              "modified_time": "string",
              "namespace_name": "string",
              "override": "bool",
              "platform": [
                "list",
                "string"
              ],
              "region_id": "string",
              "repo_name": "string",
              "source_domain": "string",
              "source_namespace_name": "string",
              "source_provider": "string",
              "source_repo_name": "string",
              "tag_count": "number",
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

func AlicloudCrArtifactSubscriptionRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCrArtifactSubscriptionRules), &result)
	return &result
}
