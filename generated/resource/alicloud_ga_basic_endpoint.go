package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaBasicEndpoint = `{
  "block": {
    "attributes": {
      "accelerator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "basic_endpoint_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_address": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_sub_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_sub_address_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_zone_id": {
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

func AlicloudGaBasicEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaBasicEndpoint), &result)
	return &result
}
