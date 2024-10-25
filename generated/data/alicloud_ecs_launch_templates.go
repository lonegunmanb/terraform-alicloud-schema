package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsLaunchTemplates = `{
  "block": {
    "attributes": {
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
      "launch_template_name": {
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
      },
      "template_resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "templates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_release_time": "string",
              "created_by": "string",
              "data_disks": [
                "list",
                [
                  "object",
                  {
                    "category": "string",
                    "delete_with_instance": "bool",
                    "description": "string",
                    "encrypted": "bool",
                    "name": "string",
                    "performance_level": "string",
                    "size": "number",
                    "snapshot_id": "string"
                  }
                ]
              ],
              "default_version_number": "number",
              "deployment_set_id": "string",
              "description": "string",
              "enable_vm_os_config": "bool",
              "host_name": "string",
              "http_endpoint": "string",
              "http_put_response_hop_limit": "number",
              "http_tokens": "string",
              "id": "string",
              "image_id": "string",
              "image_owner_alias": "string",
              "instance_charge_type": "string",
              "instance_name": "string",
              "instance_type": "string",
              "internet_charge_type": "string",
              "internet_max_bandwidth_in": "number",
              "internet_max_bandwidth_out": "number",
              "io_optimized": "string",
              "key_pair_name": "string",
              "latest_version_number": "number",
              "launch_template_id": "string",
              "launch_template_name": "string",
              "modified_time": "string",
              "network_interfaces": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "name": "string",
                    "primary_ip": "string",
                    "security_group_id": "string",
                    "vswitch_id": "string"
                  }
                ]
              ],
              "network_type": "string",
              "password_inherit": "bool",
              "period": "number",
              "private_ip_address": "string",
              "ram_role_name": "string",
              "resource_group_id": "string",
              "security_enhancement_strategy": "string",
              "security_group_id": "string",
              "security_group_ids": [
                "list",
                "string"
              ],
              "spot_duration": "string",
              "spot_price_limit": "number",
              "spot_strategy": "string",
              "system_disk": [
                "list",
                [
                  "object",
                  {
                    "category": "string",
                    "delete_with_instance": "bool",
                    "description": "string",
                    "iops": "string",
                    "name": "string",
                    "performance_level": "string",
                    "size": "number"
                  }
                ]
              ],
              "template_tags": [
                "map",
                "string"
              ],
              "user_data": "string",
              "version_description": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcsLaunchTemplatesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsLaunchTemplates), &result)
	return &result
}
