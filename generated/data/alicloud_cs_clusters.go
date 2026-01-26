package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCsClusters = `{
  "block": {
    "attributes": {
      "cluster_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_spec": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_mode": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "cluster_domain": "string",
              "cluster_id": "string",
              "cluster_name": "string",
              "cluster_spec": "string",
              "cluster_type": "string",
              "current_version": "string",
              "deletion_protection": "bool",
              "id": "string",
              "ip_stack": "string",
              "maintenance_window": [
                "list",
                [
                  "object",
                  {
                    "duration": "string",
                    "enable": "bool",
                    "maintenance_time": "string",
                    "recurrence": "string",
                    "weekly_period": "string"
                  }
                ]
              ],
              "node_cidr_mask": "string",
              "operation_policy": [
                "list",
                [
                  "object",
                  {
                    "cluster_auto_upgrade": [
                      "list",
                      [
                        "object",
                        {
                          "channel": "string",
                          "enabled": "bool"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "pod_cidr": "string",
              "profile": "string",
              "proxy_mode": "string",
              "region_id": "string",
              "resource_group_id": "string",
              "security_group_id": "string",
              "service_cidr": "string",
              "state": "string",
              "tags": [
                "map",
                "string"
              ],
              "timezone": "string",
              "vpc_id": "string",
              "vswitch_ids": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "profile": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCsClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCsClusters), &result)
	return &result
}
