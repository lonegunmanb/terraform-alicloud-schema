package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudWafInstance = `{
  "block": {
    "attributes": {
      "big_screen": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "exclusive_ip_package": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ext_bandwidth": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ext_domain_package": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_storage": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_time": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "modify_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "prefessional_service": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "renewal_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "subscription_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "waf_log": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudWafInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudWafInstance), &result)
	return &result
}
