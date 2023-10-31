package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOtsTunnels = `{
  "block": {
    "attributes": {
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
      "instance_name": {
        "description_kind": "plain",
        "required": true,
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
      "table_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tunnels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "channels": [
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
              ],
              "create_time": "number",
              "expired": "bool",
              "id": "string",
              "instance_name": "string",
              "table_name": "string",
              "tunnel_id": "string",
              "tunnel_name": "string",
              "tunnel_rpo": "number",
              "tunnel_stage": "string",
              "tunnel_type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOtsTunnelsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOtsTunnels), &result)
	return &result
}
