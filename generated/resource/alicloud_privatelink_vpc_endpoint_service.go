package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPrivatelinkVpcEndpointService = `{
  "block": {
    "attributes": {
      "auto_accept_connection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "connect_bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "dry_run": {
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
      "payer": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_business_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_domain": {
        "computed": true,
        "description_kind": "plain",
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

func AlicloudPrivatelinkVpcEndpointServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPrivatelinkVpcEndpointService), &result)
	return &result
}
