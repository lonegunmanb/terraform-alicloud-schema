package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGpdbInstances = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_categories": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_modes": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
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
      "instance_network_type": {
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
              "connection_string": "string",
              "cpu_cores": "string",
              "create_time": "string",
              "creation_time": "string",
              "db_instance_category": "string",
              "db_instance_class": "string",
              "db_instance_id": "string",
              "db_instance_mode": "string",
              "description": "string",
              "engine": "string",
              "engine_version": "string",
              "id": "string",
              "instance_network_type": "string",
              "ip_whitelist": [
                "set",
                [
                  "object",
                  {
                    "ip_group_attribute": "string",
                    "ip_group_name": "string",
                    "security_ip_list": "string"
                  }
                ]
              ],
              "maintain_end_time": "string",
              "maintain_start_time": "string",
              "master_node_num": "string",
              "memory_size": "string",
              "payment_type": "string",
              "region_id": "string",
              "seg_node_num": "string",
              "status": "string",
              "storage_size": "number",
              "storage_type": "string",
              "tags": [
                "map",
                "string"
              ],
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
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudGpdbInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGpdbInstances), &result)
	return &result
}
