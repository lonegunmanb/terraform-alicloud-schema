package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaAccelerators = `{
  "block": {
    "attributes": {
      "accelerators": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accelerator_id": "string",
              "accelerator_name": "string",
              "basic_bandwidth_package": [
                "list",
                [
                  "object",
                  {
                    "bandwidth": "number",
                    "bandwidth_type": "string",
                    "instance_id": "string"
                  }
                ]
              ],
              "cen_id": "string",
              "cross_domain_bandwidth_package": [
                "list",
                [
                  "object",
                  {
                    "bandwidth": "number",
                    "instance_id": "string"
                  }
                ]
              ],
              "ddos_id": "string",
              "description": "string",
              "dns_name": "string",
              "expired_time": "number",
              "id": "string",
              "payment_type": "string",
              "second_dns_name": "string",
              "spec": "string",
              "status": "string"
            }
          ]
        ]
      },
      "bandwidth_billing_type": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudGaAcceleratorsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaAccelerators), &result)
	return &result
}
