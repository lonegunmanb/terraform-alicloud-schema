package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionBaselineStrategy = `{
  "block": {
    "attributes": {
      "baseline_strategy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "baseline_strategy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "custom_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cycle_days": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "cycle_start_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "end_time": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "risk_sub_type_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_time": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionBaselineStrategySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionBaselineStrategy), &result)
	return &result
}
