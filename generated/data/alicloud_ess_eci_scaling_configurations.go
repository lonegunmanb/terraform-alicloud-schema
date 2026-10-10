package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssEciScalingConfigurations = `{
  "block": {
    "attributes": {
      "configurations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acr_registry_infos": [
                "list",
                [
                  "object",
                  {
                    "domains": [
                      "list",
                      "string"
                    ],
                    "instance_id": "string",
                    "instance_name": "string",
                    "region_id": "string"
                  }
                ]
              ],
              "active_deadline_seconds": "number",
              "auto_create_eip": "bool",
              "auto_match_image_cache": "bool",
              "container_group_name": "string",
              "containers": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "commands": [
                      "list",
                      "string"
                    ],
                    "cpu": "number",
                    "environment_vars": [
                      "list",
                      [
                        "object",
                        {
                          "field_ref_field_path": "string",
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "gpu": "number",
                    "image": "string",
                    "image_pull_policy": "string",
                    "lifecycle_pre_stop_handler_execs": [
                      "list",
                      "string"
                    ],
                    "liveness_probe_exec_commands": [
                      "list",
                      "string"
                    ],
                    "liveness_probe_failure_threshold": "number",
                    "liveness_probe_http_get_path": "string",
                    "liveness_probe_http_get_port": "number",
                    "liveness_probe_http_get_scheme": "string",
                    "liveness_probe_initial_delay_seconds": "number",
                    "liveness_probe_period_seconds": "number",
                    "liveness_probe_success_threshold": "number",
                    "liveness_probe_tcp_socket_port": "number",
                    "liveness_probe_timeout_seconds": "number",
                    "memory": "number",
                    "name": "string",
                    "ports": [
                      "list",
                      [
                        "object",
                        {
                          "port": "number",
                          "protocol": "string"
                        }
                      ]
                    ],
                    "readiness_probe_exec_commands": [
                      "list",
                      "string"
                    ],
                    "readiness_probe_failure_threshold": "number",
                    "readiness_probe_http_get_path": "string",
                    "readiness_probe_http_get_port": "number",
                    "readiness_probe_http_get_scheme": "string",
                    "readiness_probe_initial_delay_seconds": "number",
                    "readiness_probe_period_seconds": "number",
                    "readiness_probe_success_threshold": "number",
                    "readiness_probe_tcp_socket_port": "number",
                    "readiness_probe_timeout_seconds": "number",
                    "security_context_capability_adds": [
                      "list",
                      "string"
                    ],
                    "security_context_read_only_root_file_system": "bool",
                    "security_context_run_as_user": "number",
                    "stdin": "bool",
                    "tty": "bool",
                    "volume_mounts": [
                      "list",
                      [
                        "object",
                        {
                          "mount_path": "string",
                          "mount_propagation": "string",
                          "name": "string",
                          "read_only": "bool",
                          "sub_path": "string"
                        }
                      ]
                    ],
                    "working_dir": "string"
                  }
                ]
              ],
              "cost_optimization": "bool",
              "cpu": "number",
              "cpu_options_core": "number",
              "cpu_options_threads_per_core": "number",
              "creation_time": "string",
              "description": "string",
              "dns_config_options": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "value": "string"
                  }
                ]
              ],
              "dns_policy": "string",
              "egress_bandwidth": "number",
              "eip_bandwidth": "number",
              "enable_sls": "bool",
              "ephemeral_storage": "number",
              "host_aliases": [
                "list",
                [
                  "object",
                  {
                    "hostnames": [
                      "list",
                      "string"
                    ],
                    "ip": "string"
                  }
                ]
              ],
              "host_name": "string",
              "id": "string",
              "image_registry_credentials": [
                "list",
                [
                  "object",
                  {
                    "password": "string",
                    "server": "string",
                    "username": "string"
                  }
                ]
              ],
              "image_snapshot_id": "string",
              "ingress_bandwidth": "number",
              "init_containers": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "commands": [
                      "list",
                      "string"
                    ],
                    "cpu": "number",
                    "environment_vars": [
                      "list",
                      [
                        "object",
                        {
                          "field_ref_field_path": "string",
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "gpu": "number",
                    "image": "string",
                    "image_pull_policy": "string",
                    "memory": "number",
                    "name": "string",
                    "ports": [
                      "list",
                      [
                        "object",
                        {
                          "port": "number",
                          "protocol": "string"
                        }
                      ]
                    ],
                    "security_context_capability_adds": [
                      "list",
                      "string"
                    ],
                    "security_context_read_only_root_file_system": "bool",
                    "security_context_run_as_user": "number",
                    "volume_mounts": [
                      "list",
                      [
                        "object",
                        {
                          "mount_path": "string",
                          "mount_propagation": "string",
                          "name": "string",
                          "read_only": "bool",
                          "sub_path": "string"
                        }
                      ]
                    ],
                    "working_dir": "string"
                  }
                ]
              ],
              "instance_family_level": "string",
              "instance_types": [
                "list",
                "string"
              ],
              "ipv6_address_count": "number",
              "lifecycle_state": "string",
              "load_balancer_weight": "number",
              "memory": "number",
              "ram_role_name": "string",
              "resource_group_id": "string",
              "restart_policy": "string",
              "scaling_configuration_name": "string",
              "scaling_group_id": "string",
              "security_context_sysctls": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "value": "string"
                  }
                ]
              ],
              "security_group_id": "string",
              "spot_price_limit": "number",
              "spot_strategy": "string",
              "tags": [
                "map",
                "string"
              ],
              "termination_grace_period_seconds": "number",
              "volumes": [
                "list",
                [
                  "object",
                  {
                    "config_file_volume_config_file_to_paths": [
                      "list",
                      [
                        "object",
                        {
                          "content": "string",
                          "mode": "number",
                          "path": "string"
                        }
                      ]
                    ],
                    "config_file_volume_default_mode": "number",
                    "disk_volume_disk_id": "string",
                    "disk_volume_disk_size": "number",
                    "disk_volume_fs_type": "string",
                    "empty_dir_volume_medium": "string",
                    "empty_dir_volume_size_limit": "string",
                    "flex_volume_driver": "string",
                    "flex_volume_fs_type": "string",
                    "flex_volume_options": "string",
                    "host_path_volume_path": "string",
                    "host_path_volume_type": "string",
                    "name": "string",
                    "nfs_volume_path": "string",
                    "nfs_volume_read_only": "bool",
                    "nfs_volume_server": "string",
                    "type": "string"
                  }
                ]
              ]
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
      "scaling_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssEciScalingConfigurationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssEciScalingConfigurations), &result)
	return &result
}
