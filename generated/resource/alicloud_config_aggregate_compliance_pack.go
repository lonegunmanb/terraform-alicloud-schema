package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudConfigAggregateCompliancePack = `{
  "block": {
    "attributes": {
      "aggregate_compliance_pack_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "aggregator_compliance_pack_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aggregator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compliance_pack_template_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
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
      "risk_level": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "config_rule_ids": {
        "block": {
          "attributes": {
            "config_rule_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "config_rules": {
        "block": {
          "attributes": {
            "managed_rule_identifier": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "config_rule_parameters": {
              "block": {
                "attributes": {
                  "parameter_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parameter_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
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

func AlicloudConfigAggregateCompliancePackSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudConfigAggregateCompliancePack), &result)
	return &result
}
