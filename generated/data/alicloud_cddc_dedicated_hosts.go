package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCddcDedicatedHosts = `{
  "block": {
    "attributes": {
      "allocation_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dedicated_host_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "host_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hosts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allocation_status": "string",
              "bastion_instance_id": "string",
              "cpu_allocation_ratio": "string",
              "cpu_used": "string",
              "create_time": "string",
              "dedicated_host_group_id": "string",
              "dedicated_host_id": "string",
              "disk_allocation_ratio": "string",
              "ecs_class_code": "string",
              "end_time": "string",
              "engine": "string",
              "expired_time": "string",
              "host_class": "string",
              "host_cpu": "string",
              "host_mem": "string",
              "host_name": "string",
              "host_storage": "string",
              "host_type": "string",
              "id": "string",
              "image_category": "string",
              "ip_address": "string",
              "mem_allocation_ratio": "string",
              "memory_used": "string",
              "open_permission": "string",
              "status": "string",
              "storage_used": "string",
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
      "order_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AlicloudCddcDedicatedHostsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCddcDedicatedHosts), &result)
	return &result
}
