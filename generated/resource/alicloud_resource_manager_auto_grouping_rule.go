package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudResourceManagerAutoGroupingRule = `{
  "block": {
    "attributes": {
      "exclude_region_ids_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclude_resource_group_ids_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclude_resource_ids_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclude_resource_types_scope": {
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
      "resource_ids_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_types_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "rule_contents": {
        "block": {
          "attributes": {
            "auto_grouping_scope_condition": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_resource_group_condition": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
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

func AlicloudResourceManagerAutoGroupingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudResourceManagerAutoGroupingRule), &result)
	return &result
}
