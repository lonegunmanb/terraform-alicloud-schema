package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsDedicatedHosts = `{
  "block": {
    "attributes": {
      "dedicated_host_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dedicated_host_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dedicated_host_type": {
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
              "action_on_maintenance": "string",
              "auto_placement": "string",
              "auto_release_time": "string",
              "capacity": [
                "list",
                [
                  "object",
                  {
                    "available_local_storage": "number",
                    "available_memory": "number",
                    "available_vcpus": "number",
                    "available_vgpus": "number",
                    "local_storage_category": "string",
                    "total_local_storage": "number",
                    "total_memory": "number",
                    "total_vcpus": "number",
                    "total_vgpus": "number"
                  }
                ]
              ],
              "cores": "number",
              "cpu_over_commit_ratio": "number",
              "dedicated_host_id": "string",
              "dedicated_host_name": "string",
              "dedicated_host_type": "string",
              "description": "string",
              "expired_time": "string",
              "gpu_spec": "string",
              "id": "string",
              "machine_id": "string",
              "network_attributes": [
                "list",
                [
                  "object",
                  {
                    "slb_udp_timeout": "number",
                    "udp_timeout": "number"
                  }
                ]
              ],
              "operation_locks": [
                "list",
                [
                  "object",
                  {
                    "lock_reason": "string"
                  }
                ]
              ],
              "payment_type": "string",
              "physical_gpus": "number",
              "resource_group_id": "string",
              "sale_cycle": "string",
              "sockets": "number",
              "status": "string",
              "supported_custom_instance_type_families": [
                "list",
                "string"
              ],
              "supported_instance_type_families": [
                "list",
                "string"
              ],
              "supported_instance_types_list": [
                "list",
                "string"
              ],
              "tags": [
                "map",
                "string"
              ],
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
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "operation_locks": {
        "block": {
          "attributes": {
            "lock_reason": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcsDedicatedHostsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsDedicatedHosts), &result)
	return &result
}
