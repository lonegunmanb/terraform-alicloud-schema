package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsDedicatedHostClusters = `{
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
              "dedicated_host_cluster_capacity": [
                "list",
                [
                  "object",
                  {
                    "available_memory": "number",
                    "available_vcpus": "number",
                    "local_storage_capacities": [
                      "list",
                      [
                        "object",
                        {
                          "available_disk": "number",
                          "data_disk_category": "string",
                          "total_disk": "number"
                        }
                      ]
                    ],
                    "total_memory": "number",
                    "total_vcpus": "number"
                  }
                ]
              ],
              "dedicated_host_cluster_id": "string",
              "dedicated_host_cluster_name": "string",
              "dedicated_host_ids": [
                "list",
                "string"
              ],
              "description": "string",
              "id": "string",
              "resource_group_id": "string",
              "tags": [
                "map",
                "string"
              ],
              "zone_id": "string"
            }
          ]
        ]
      },
      "dedicated_host_cluster_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "dedicated_host_cluster_name": {
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
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcsDedicatedHostClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsDedicatedHostClusters), &result)
	return &result
}
