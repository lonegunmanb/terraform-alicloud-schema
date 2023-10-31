package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCassandraDataCenters = `{
  "block": {
    "attributes": {
      "centers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster_id": "string",
              "commodity_instance": "string",
              "created_time": "string",
              "data_center_id": "string",
              "data_center_name": "string",
              "disk_size": "number",
              "disk_type": "string",
              "expire_time": "string",
              "instance_type": "string",
              "lock_mode": "string",
              "node_count": "number",
              "pay_type": "string",
              "status": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCassandraDataCentersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCassandraDataCenters), &result)
	return &result
}
