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
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "container_image_scan_new": {
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
      "post_paid_flag": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "post_pay_module_switch": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rasp_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "renew_period": {
        "computed": true,
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
      "sas_cspm": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sas_cspm_switch": {
        "computed": true,
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
        "computed": true,
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
        "computed": true,
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
      "subscription_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "threat_analysis": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "threat_analysis_flow": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "threat_analysis_sls_storage": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "threat_analysis_switch": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "threat_analysis_switch1": {
        "computed": true,
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
        "optional": true,
        "type": "string"
      },
      "vul_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vul_switch": {
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

func AlicloudThreatDetectionInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionInstance), &result)
	return &result
}
