package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudPhoneInstanceGroup = `{
  "block": {
    "attributes": {
      "amount": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "auto_pay": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gpu_acceleration": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_group_spec": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "number_of_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "office_site_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "period_unit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
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

func AlicloudCloudPhoneInstanceGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudPhoneInstanceGroup), &result)
	return &result
}
