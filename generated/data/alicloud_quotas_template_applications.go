package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudQuotasTemplateApplications = `{
  "block": {
    "attributes": {
      "applications": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "aliyun_uids": [
                "list",
                "string"
              ],
              "apply_time": "string",
              "audit_status_vos": [
                "list",
                [
                  "object",
                  {
                    "count": "number",
                    "status": "string"
                  }
                ]
              ],
              "batch_quota_application_id": "string",
              "desire_value": "number",
              "dimensions": [
                "set",
                [
                  "object",
                  {
                    "key": "string",
                    "value": "string"
                  }
                ]
              ],
              "effective_time": "string",
              "expire_time": "string",
              "id": "string",
              "product_code": "string",
              "quota_action_code": "string",
              "quota_category": "string",
              "reason": "string"
            }
          ]
        ]
      },
      "batch_quota_application_id": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "product_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quota_action_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quota_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
