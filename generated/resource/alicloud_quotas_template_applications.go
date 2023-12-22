package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudQuotasTemplateApplications = `{
  "block": {
    "attributes": {
      "aliyun_uids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
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
      "quota_application_details": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "aliyun_uid": "string",
              "application_id": "string",
              "approve_value": "number",
              "audit_reason": "string",
              "dimensions": [
                "map",
                "string"
              ],
              "env_language": "string",
              "notice_type": "number",
              "period": [
                "list",
                [
                  "object",
                  {
                    "period_unit": "string",
                    "period_value": "number"
                  }
                ]
              ],
              "quota_arn": "string",
              "quota_description": "string",
              "quota_name": "string",
              "quota_unit": "string",
              "reason": "string",
              "status": "string"
            }
          ]
        ]
      },
      "quota_category": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reason": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudQuotasTemplateApplicationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudQuotasTemplateApplications), &result)
	return &result
}
