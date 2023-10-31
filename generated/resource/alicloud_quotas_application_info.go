package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudQuotasApplicationInfo = `{
  "block": {
    "attributes": {
      "approve_value": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "audit_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "audit_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "desire_value": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "effective_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "env_language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "expire_time": {
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
      "notice_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "product_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "quota_action_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "quota_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quota_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "quota_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "quota_unit": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reason": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "dimensions": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
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

func AlicloudQuotasApplicationInfoSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudQuotasApplicationInfo), &result)
	return &result
}
