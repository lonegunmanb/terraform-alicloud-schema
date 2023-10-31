package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudConfigAggregateConfigRules = `{
  "block": {
    "attributes": {
      "aggregate_config_rule_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "aggregator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
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
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "risk_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_id": "string",
              "aggregate_config_rule_name": "string",
              "aggregator_id": "string",
              "compliance": [
                "list",
                [
                  "object",
                  {
                    "compliance_type": "string",
                    "count": "number"
                  }
                ]
              ],
              "compliance_pack_id": "string",
              "config_rule_arn": "string",
              "config_rule_id": "string",
              "config_rule_trigger_types": "string",
              "description": "string",
              "event_source": "string",
              "exclude_resource_ids_scope": "string",
              "id": "string",
              "input_parameters": [
                "map",
                "string"
              ],
              "maximum_execution_frequency": "string",
              "modified_timestamp": "string",
              "region_ids_scope": "string",
              "resource_group_ids_scope": "string",
              "resource_types_scope": [
                "list",
                "string"
              ],
              "risk_level": "number",
              "source_identifier": "string",
              "source_owner": "string",
              "status": "string",
              "tag_key_scope": "string",
              "tag_value_scope": "string"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudConfigAggregateConfigRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudConfigAggregateConfigRules), &result)
	return &result
}
