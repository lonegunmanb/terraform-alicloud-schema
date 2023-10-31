package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbClusters = `{
  "block": {
    "attributes": {
      "clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "charge_type": "string",
              "connection_string": "string",
              "create_time": "string",
              "db_node_class": "string",
              "db_node_number": "number",
              "db_nodes": [
                "list",
                [
                  "object",
                  {
                    "create_time": "string",
                    "db_node_class": "string",
                    "db_node_id": "string",
                    "db_node_role": "string",
                    "db_node_status": "string",
                    "max_connections": "number",
                    "max_iops": "number",
                    "region_id": "string",
                    "zone_id": "string"
                  }
                ]
              ],
              "db_type": "string",
              "db_version": "string",
              "delete_lock": "number",
              "description": "string",
              "engine": "string",
              "expire_time": "string",
              "expired": "string",
              "id": "string",
              "lock_mode": "string",
              "network_type": "string",
              "port": "string",
              "region_id": "string",
              "status": "string",
              "storage_used": "number",
              "vpc_id": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "db_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "descriptions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudPolardbClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbClusters), &result)
	return &result
}
