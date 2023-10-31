package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudConfigAggregateConfigRule = `{
  "block": {
    "attributes": {
      "aggregate_config_rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "aggregator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "config_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "config_rule_trigger_types": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclude_resource_ids_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "input_parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "maximum_execution_frequency": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_ids_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_ids_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_types_scope": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "risk_level": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "source_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_owner": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tag_key_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tag_value_scope": {
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

func AlicloudConfigAggregateConfigRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudConfigAggregateConfigRule), &result)
	return &result
}
