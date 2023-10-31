package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionAntiBruteForceRule = `{
  "block": {
    "attributes": {
      "anti_brute_force_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "anti_brute_force_rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_rule": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fail_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "forbidden_time": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "span": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "uuid_list": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionAntiBruteForceRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionAntiBruteForceRule), &result)
	return &result
}
