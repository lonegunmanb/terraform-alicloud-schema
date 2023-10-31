package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaBasicAccelerators = `{
  "block": {
    "attributes": {
      "accelerator_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "accelerators": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bandwidth_billing_type": "string",
              "basic_accelerator_id": "string",
              "basic_accelerator_name": "string",
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
              "basic_endpoint_group_id": "string",
              "basic_ip_set_id": "string",
              "create_time": "number",
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
              "description": "string",
              "expired_time": "number",
              "id": "string",
              "instance_charge_type": "string",
              "region_id": "string",
              "status": "string"
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
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AlicloudGaBasicAcceleratorsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaBasicAccelerators), &result)
	return &result
}
