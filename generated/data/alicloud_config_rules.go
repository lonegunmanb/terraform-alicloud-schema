package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudConfigRules = `{
  "block": {
    "attributes": {
      "config_rule_state": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
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
      "member_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "message_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_account": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "rule_name": {
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
              "account_id": "string",
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
              "config_rule_state": "string",
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
              "rule_name": "string",
              "scope_compliance_resource_types": [
                "list",
                "string"
              ],
              "source_detail_message_type": "string",
              "source_identifier": "string",
              "source_maximum_execution_frequency": "string",
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

func AlicloudConfigRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudConfigRules), &result)
	return &result
}
