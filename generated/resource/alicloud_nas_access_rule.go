package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNasAccessRule = `{
  "block": {
    "attributes": {
      "access_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "access_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_type": {
        "computed": true,
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
      "ipv6_source_cidr_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rw_access_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_cidr_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_access_type": {
        "computed": true,
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

func AlicloudNasAccessRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNasAccessRule), &result)
	return &result
}
