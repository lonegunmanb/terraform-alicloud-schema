package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSelectdbDbClusters = `{
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
              "cache_size": "number",
              "cpu": "number",
              "create_time": "string",
              "db_cluster_class": "string",
              "db_cluster_description": "string",
              "db_cluster_id": "string",
              "db_instance_id": "string",
              "engine": "string",
              "engine_version": "string",
              "id": "string",
              "memory": "number",
              "param_change_logs": [
                "list",
                [
                  "object",
                  {
                    "config_id": "number",
                    "gmt_created": "string",
                    "gmt_modified": "string",
                    "is_applied": "number",
                    "name": "string",
                    "new_value": "string",
                    "old_value": "string"
                  }
                ]
              ],
              "params": [
                "list",
                [
                  "object",
                  {
                    "comment": "string",
                    "default_value": "string",
                    "is_dynamic": "number",
                    "is_user_modifiable": "number",
                    "name": "string",
                    "optional": "string",
                    "param_category": "string",
                    "value": "string"
                  }
                ]
              ],
              "payment_type": "string",
              "region_id": "string",
              "status": "string",
              "vpc_id": "string",
              "zone_id": "string"
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSelectdbDbClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSelectdbDbClusters), &result)
	return &result
}
