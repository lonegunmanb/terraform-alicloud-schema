package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlidnsGtmInstances = `{
  "block": {
    "attributes": {
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alert_config": [
                "list",
                [
                  "object",
                  {
                    "dingtalk_notice": "bool",
                    "email_notice": "bool",
                    "notice_type": "string",
                    "sms_notice": "bool"
                  }
                ]
              ],
              "alert_group": [
                "list",
                "string"
              ],
              "cname_type": "string",
              "create_time": "string",
              "expire_time": "string",
              "id": "string",
              "instance_id": "string",
              "instance_name": "string",
              "package_edition": "string",
              "payment_type": "string",
              "public_cname_mode": "string",
              "public_rr": "string",
              "public_user_domain_name": "string",
              "public_zone_name": "string",
              "resource_group_id": "string",
              "strategy_mode": "string",
              "ttl": "number"
            }
          ]
        ]
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlidnsGtmInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlidnsGtmInstances), &result)
	return &result
}
