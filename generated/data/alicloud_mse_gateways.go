package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMseGateways = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gateway_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateways": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_vswitch_id": "string",
              "gateway_name": "string",
              "gateway_unique_id": "string",
              "id": "string",
              "payment_type": "string",
              "replica": "string",
              "slb_list": [
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
              ],
              "spec": "string",
              "status": "string",
              "vpc_id": "string",
              "vswitch_id": "string"
            }
          ]
        ]
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
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
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
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMseGatewaysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMseGateways), &result)
	return &result
}
