package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApiGatewayGroup = `{
  "block": {
    "attributes": {
      "base_path": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
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
      "instance_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sub_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_intranet_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "user_log_config": {
        "block": {
          "attributes": {
            "jwt_claims": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query_string": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "request_body": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "request_headers": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response_body": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "response_headers": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudApiGatewayGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApiGatewayGroup), &result)
	return &result
}
