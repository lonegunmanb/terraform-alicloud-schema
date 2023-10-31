package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMseGateway = `{
  "block": {
    "attributes": {
      "backup_vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_slb": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enterprise_security_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gateway_name": {
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
      "internet_slb_spec": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replica": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "slb_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "associate_id": "string",
              "gateway_slb_mode": "string",
              "gateway_slb_status": "string",
              "gmt_create": "string",
              "slb_id": "string",
              "slb_ip": "string",
              "slb_port": "string",
              "type": "string"
            }
          ]
        ]
      },
      "slb_spec": {
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
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudMseGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMseGateway), &result)
	return &result
}
