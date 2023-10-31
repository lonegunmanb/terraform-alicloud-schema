package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMongodbInstances = `{
  "block": {
    "attributes": {
      "availability_zone": {
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
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "instance_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "availability_zone": "string",
              "charge_type": "string",
              "creation_time": "string",
              "engine": "string",
              "engine_version": "string",
              "expiration_time": "string",
              "id": "string",
              "instance_class": "string",
              "instance_type": "string",
              "lock_mode": "string",
              "mongos": [
                "list",
                [
                  "object",
                  {
                    "class": "string",
                    "description": "string",
                    "node_id": "string"
                  }
                ]
              ],
              "name": "string",
              "network_type": "string",
              "region_id": "string",
              "replication": "string",
              "shards": [
                "list",
                [
                  "object",
                  {
                    "class": "string",
                    "description": "string",
                    "node_id": "string",
                    "storage": "number"
                  }
                ]
              ],
              "status": "string",
              "storage": "number",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
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

func AlicloudMongodbInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMongodbInstances), &result)
	return &result
}
