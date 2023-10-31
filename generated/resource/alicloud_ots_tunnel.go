package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOtsTunnel = `{
  "block": {
    "attributes": {
      "channels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "channel_id": "string",
              "channel_rpo": "number",
              "channel_status": "string",
              "channel_type": "string",
              "client_id": "string"
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "expired": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "table_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tunnel_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tunnel_rpo": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tunnel_stage": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel_type": {
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

func AlicloudOtsTunnelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOtsTunnel), &result)
	return &result
}
