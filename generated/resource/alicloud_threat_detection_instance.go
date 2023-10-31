package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionInstance = `{
  "block": {
    "attributes": {
      "buy_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "container_image_scan": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "honeypot": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "honeypot_switch": {
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
      "instance_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "modify_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "renewal_period_unit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "renewal_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sas_anti_ransomware": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sas_sc": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "sas_sdk": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sas_sdk_switch": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sas_sls_storage": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sas_webguard_boolean": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sas_webguard_order_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "threat_analysis": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "threat_analysis_switch": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "v_core": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionInstance), &result)
	return &result
}
