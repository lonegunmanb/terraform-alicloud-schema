package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudClickHouseDbClusters = `{
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
              "ali_uid": "string",
              "bid": "string",
              "category": "string",
              "commodity_code": "string",
              "connection_string": "string",
              "control_version": "string",
              "create_time": "string",
              "db_cluster_access_white_list": [
                "list",
                [
                  "object",
                  {
                    "db_cluster_ip_array_attribute": "string",
                    "db_cluster_ip_array_name": "string",
                    "security_ip_list": "string"
                  }
                ]
              ],
              "db_cluster_description": "string",
              "db_cluster_id": "string",
              "db_cluster_network_type": "string",
              "db_cluster_type": "string",
              "db_node_class": "string",
              "db_node_count": "string",
              "db_node_storage": "string",
              "encryption_key": "string",
              "encryption_type": "string",
              "engine": "string",
              "engine_version": "string",
              "expire_time": "string",
              "id": "string",
              "is_expired": "string",
              "lock_mode": "string",
              "lock_reason": "string",
              "maintain_time": "string",
              "payment_type": "string",
              "port": "number",
              "public_connection_string": "string",
              "public_port": "string",
              "scale_out_status": [
                "list",
                [
                  "object",
                  {
                    "progress": "string",
                    "ratio": "string"
                  }
                ]
              ],
              "status": "string",
              "storage_type": "string",
              "support_backup": "number",
              "support_https_port": "bool",
              "support_mysql_port": "bool",
              "vpc_cloud_instance_id": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "db_cluster_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudClickHouseDbClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudClickHouseDbClusters), &result)
	return &result
}
