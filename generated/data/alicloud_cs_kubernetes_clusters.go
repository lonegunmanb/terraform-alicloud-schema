package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCsKubernetesClusters = `{
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
              "availability_zone": "string",
              "cluster_network_type": "string",
              "connections": [
                "map",
                "string"
              ],
              "id": "string",
              "image_id": "string",
              "key_name": "string",
              "log_config": [
                "list",
                [
                  "object",
                  {
                    "project": "string",
                    "type": "string"
                  }
                ]
              ],
              "master_auto_renew": "bool",
              "master_auto_renew_period": "number",
              "master_disk_category": "string",
              "master_disk_size": "number",
              "master_instance_charge_type": "string",
              "master_instance_types": [
                "list",
                "string"
              ],
              "master_nodes": [
                "list",
                [
                  "object",
                  {
                    "id": "string",
                    "name": "string",
                    "private_ip": "string"
                  }
                ]
              ],
              "master_period": "number",
              "master_period_unit": "string",
              "name": "string",
              "nat_gateway_id": "string",
              "node_cidr_mask": "number",
              "pod_cidr": "string",
              "security_group_id": "string",
              "service_cidr": "string",
              "slb_internet_enabled": "bool",
              "vpc_id": "string",
              "vswitch_ids": [
                "list",
                "string"
              ],
              "worker_auto_renew": "bool",
              "worker_auto_renew_period": "number",
              "worker_data_disk_category": "string",
              "worker_data_disk_size": "number",
              "worker_disk_category": "string",
              "worker_disk_size": "number",
              "worker_instance_charge_type": "string",
              "worker_instance_types": [
                "list",
                "string"
              ],
              "worker_nodes": [
                "list",
                [
                  "object",
                  {
                    "id": "string",
                    "name": "string",
                    "private_ip": "string"
                  }
                ]
              ],
              "worker_numbers": [
                "list",
                "number"
              ],
              "worker_period": "number",
              "worker_period_unit": "string"
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
      "kube_config_file_prefix": {
        "description_kind": "plain",
        "optional": true,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCsKubernetesClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCsKubernetesClusters), &result)
	return &result
}
