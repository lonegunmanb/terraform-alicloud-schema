package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMongodbServerlessInstances = `{
  "block": {
    "attributes": {
      "db_instance_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
              "capacity_unit": "number",
              "db_instance_class": "string",
              "db_instance_description": "string",
              "db_instance_id": "string",
              "db_instance_release_protection": "bool",
              "db_instance_storage": "number",
              "engine": "string",
              "engine_version": "string",
              "expire_time": "string",
              "id": "string",
              "kind_code": "string",
              "lock_mode": "string",
              "maintain_end_time": "string",
              "maintain_start_time": "string",
              "max_connections": "number",
              "max_iops": "number",
              "network_type": "string",
              "payment_type": "string",
              "protocol_type": "string",
              "resource_group_id": "string",
              "security_ip_groups": [
                "list",
                [
                  "object",
                  {
                    "security_ip_group_attribute": "string",
                    "security_ip_group_name": "string",
                    "security_ip_list": "string"
                  }
                ]
              ],
              "status": "string",
              "storage_engine": "string",
              "tags": [
                "map",
                "string"
              ],
              "vpc_auth_mode": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
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

func AlicloudMongodbServerlessInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMongodbServerlessInstances), &result)
	return &result
}
