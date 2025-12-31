package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudConfigAggregateRemediation = `{
  "block": {
    "attributes": {
      "aggregator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "config_rule_id": {
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
      "invoke_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remediation_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "remediation_origin_params": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remediation_source_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remediation_template_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remediation_type": {
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

func AlicloudConfigAggregateRemediationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudConfigAggregateRemediation), &result)
	return &result
}
