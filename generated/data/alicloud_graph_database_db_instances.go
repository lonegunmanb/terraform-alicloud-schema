package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGraphDatabaseDbInstances = `{
  "block": {
    "attributes": {
      "db_instance_description": {
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_string": "string",
              "create_time": "string",
              "current_minor_version": "string",
              "db_instance_category": "string",
              "db_instance_cpu": "string",
              "db_instance_description": "string",
              "db_instance_id": "string",
              "db_instance_ip_array": [
                "set",
                [
                  "object",
                  {
                    "db_instance_ip_array_attribute": "string",
                    "db_instance_ip_array_name": "string",
                    "security_ips": "string"
                  }
                ]
              ],
              "db_instance_memory": "string",
              "db_instance_network_type": "string",
              "db_instance_storage_type": "string",
              "db_instance_type": "string",
              "db_node_class": "string",
              "db_node_count": "string",
              "db_node_storage": "string",
              "db_version": "string",
              "expire_time": "string",
              "expired": "string",
              "id": "string",
              "latest_minor_version": "string",
              "lock_mode": "string",
              "lock_reason": "string",
              "maintain_time": "string",
              "master_db_instance_id": "string",
              "payment_type": "string",
              "port": "number",
              "public_connection_string": "string",
              "public_port": "number",
              "read_only_db_instance_ids": [
                "list",
                "string"
              ],
              "status": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
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

func AlicloudGraphDatabaseDbInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGraphDatabaseDbInstances), &result)
	return &result
}
