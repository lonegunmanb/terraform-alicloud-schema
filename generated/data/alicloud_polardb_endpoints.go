package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbEndpoints = `{
  "block": {
    "attributes": {
      "db_cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_endpoint_id": {
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
              "address_items": [
                "list",
                [
                  "object",
                  {
                    "connection_string": "string",
                    "ip_address": "string",
                    "net_type": "string",
                    "port": "string",
                    "vpc_id": "string",
                    "vswitch_id": "string"
                  }
                ]
              ],
              "auto_add_new_nodes": "string",
              "db_endpoint_id": "string",
              "endpoint_config": "string",
              "endpoint_type": "string",
              "nodes": "string",
              "read_write_mode": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudPolardbEndpointsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbEndpoints), &result)
	return &result
}
