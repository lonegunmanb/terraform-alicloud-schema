package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudKvstoreInstances = `{
  "block": {
    "attributes": {
      "architecture_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edition_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "engine_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "expired": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_instance": {
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
              "architecture_type": "string",
              "auto_renew": "bool",
              "auto_renew_period": "number",
              "availability_zone": "string",
              "bandwidth": "number",
              "capacity": "number",
              "charge_type": "string",
              "config": [
                "map",
                "string"
              ],
              "connection_domain": "string",
              "connection_mode": "string",
              "connections": "number",
              "create_time": "string",
              "db_instance_id": "string",
              "db_instance_name": "string",
              "destroy_time": "string",
              "end_time": "string",
              "engine_version": "string",
              "expire_time": "string",
              "has_renew_change_order": "bool",
              "id": "string",
              "instance_class": "string",
              "instance_release_protection": "bool",
              "instance_type": "string",
              "is_rds": "bool",
              "maintain_end_time": "string",
              "maintain_start_time": "string",
              "max_connections": "number",
              "name": "string",
              "network_type": "string",
              "node_type": "string",
              "package_type": "string",
              "payment_type": "string",
              "port": "number",
              "private_ip": "string",
              "qps": "number",
              "region_id": "string",
              "replacate_id": "string",
              "resource_group_id": "string",
              "search_key": "string",
              "secondary_zone_id": "string",
              "security_group_id": "string",
              "security_ip_group_attribute": "string",
              "security_ip_group_name": "string",
              "security_ips": [
                "list",
                "string"
              ],
              "ssl_enable": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "user_name": "string",
              "vpc_auth_mode": "string",
              "vpc_cloud_instance_id": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
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
      "network_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "search_key": {
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
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AlicloudKvstoreInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudKvstoreInstances), &result)
	return &result
}
