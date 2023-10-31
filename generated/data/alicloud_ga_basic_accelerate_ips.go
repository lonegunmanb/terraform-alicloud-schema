package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaBasicAccelerateIps = `{
  "block": {
    "attributes": {
      "accelerate_ip_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "accelerate_ip_id": {
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
      "ip_set_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ips": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accelerate_ip_address": "string",
              "accelerate_ip_id": "string",
              "accelerator_id": "string",
              "id": "string",
              "ip_set_id": "string",
              "status": "string"
            }
          ]
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

func AlicloudGaBasicAccelerateIpsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaBasicAccelerateIps), &result)
	return &result
}
