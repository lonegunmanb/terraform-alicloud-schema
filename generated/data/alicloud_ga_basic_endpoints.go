package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaBasicEndpoints = `{
  "block": {
    "attributes": {
      "endpoint_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accelerator_id": "string",
              "basic_endpoint_name": "string",
              "endpoint_address": "string",
              "endpoint_group_id": "string",
              "endpoint_id": "string",
              "endpoint_sub_address": "string",
              "endpoint_sub_address_type": "string",
              "endpoint_type": "string",
              "endpoint_zone_id": "string",
              "id": "string",
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
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AlicloudGaBasicEndpointsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaBasicEndpoints), &result)
	return &result
}
