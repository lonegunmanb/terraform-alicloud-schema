package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAdbDbClusters = `{
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
              "auto_renew_period": "number",
              "available_kernel_versions": [
                "list",
                [
                  "object",
                  {
                    "expire_date": "string",
                    "kernel_version": "string",
                    "release_date": "string"
                  }
                ]
              ],
              "charge_type": "string",
              "commodity_code": "string",
              "compute_resource": "string",
              "connection_string": "string",
              "create_time": "string",
              "db_cluster_category": "string",
              "db_cluster_id": "string",
              "db_cluster_network_type": "string",
              "db_cluster_type": "string",
              "db_cluster_version": "string",
              "db_node_class": "string",
              "db_node_count": "number",
              "db_node_storage": "number",
              "description": "string",
              "disk_type": "string",
              "dts_job_id": "string",
              "elastic_io_resource": "number",
              "engine": "string",
              "engine_version": "string",
              "executor_count": "string",
              "expire_time": "string",
              "expired": "string",
              "id": "string",
              "kernel_version": "string",
              "lock_mode": "string",
              "lock_reason": "string",
              "maintain_time": "string",
              "mode": "string",
              "network_type": "string",
              "payment_type": "string",
              "port": "number",
              "rds_instance_id": "string",
              "region_id": "string",
              "renewal_status": "string",
              "resource_group_id": "string",
              "security_ips": [
                "list",
                "string"
              ],
              "status": "string",
              "storage_resource": "string",
              "tags": [
                "map",
                "string"
              ],
              "vpc_cloud_instance_id": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "description": {
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_group_id": {
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
      },
      "total_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAdbDbClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAdbDbClusters), &result)
	return &result
}
