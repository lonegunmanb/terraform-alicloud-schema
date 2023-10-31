package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudConfigRule = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "compliance": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "compliance_type": "string",
              "count": "number"
            }
          ]
        ]
      },
      "compliance_pack_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "config_rule_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "config_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "config_rule_trigger_types": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_source": {
        "computed": true,
        "description_kind": "plain",
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
      "modified_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope_compliance_resource_types": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_detail_message_type": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_maximum_execution_frequency": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
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

func AlicloudConfigRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudConfigRule), &result)
	return &result
}
