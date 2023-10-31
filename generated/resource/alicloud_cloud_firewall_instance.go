package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallInstance = `{
  "block": {
    "attributes": {
      "account_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "band_width": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "cfw_account": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cfw_log": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "cfw_log_storage": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cfw_service": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fw_vpc_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ip_number": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "logistics": {
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
        "required": true,
        "type": "number"
      },
      "release_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "renew_period": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "renewal_duration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "renewal_duration_unit": {
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
      "spec": {
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudFirewallInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallInstance), &result)
	return &result
}
