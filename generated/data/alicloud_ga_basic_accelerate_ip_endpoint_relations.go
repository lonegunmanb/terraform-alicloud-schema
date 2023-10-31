package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaBasicAccelerateIpEndpointRelations = `{
  "block": {
    "attributes": {
      "accelerate_ip_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "accelerator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_id": {
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "relations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accelerate_ip_id": "string",
              "accelerator_id": "string",
              "basic_endpoint_name": "string",
              "endpoint_address": "string",
              "endpoint_id": "string",
              "endpoint_sub_address": "string",
              "endpoint_sub_address_type": "string",
              "endpoint_type": "string",
              "endpoint_zone_id": "string",
              "id": "string",
              "ip_address": "string",
              "status": "string"
            }
          ]
        ]
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

func AlicloudGaBasicAccelerateIpEndpointRelationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaBasicAccelerateIpEndpointRelations), &result)
	return &result
}
