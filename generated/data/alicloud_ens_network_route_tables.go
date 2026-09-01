package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEnsNetworkRouteTables = `{
  "block": {
    "attributes": {
      "associate_type": {
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
        "type": [
          "list",
          "string"
        ]
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "network_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tables": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "associate_type": "string",
              "create_time": "string",
              "description": "string",
              "id": "string",
              "is_default_gateway_route_table": "bool",
              "network_id": "string",
              "route_table_id": "string",
              "route_table_name": "string",
              "route_table_type": "string",
              "status": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEnsNetworkRouteTablesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEnsNetworkRouteTables), &result)
	return &result
}
