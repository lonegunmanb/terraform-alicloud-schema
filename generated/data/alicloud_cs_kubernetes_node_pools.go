package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCsKubernetesNodePools = `{
  "block": {
    "attributes": {
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "nodepools": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_renew": "bool",
              "auto_renew_period": "number",
              "cis_enabled": "bool",
              "compensate_with_on_demand": "bool",
              "cpu_policy": "string",
              "data_disks": [
                "list",
                [
                  "object",
                  {
                    "auto_format": "string",
                    "auto_snapshot_policy_id": "string",
                    "bursting_enabled": "bool",
                    "category": "string",
                    "device": "string",
                    "encrypted": "string",
                    "file_system": "string",
                    "kms_key_id": "string",
                    "mount_target": "string",
                    "name": "string",
                    "performance_level": "string",
                    "provisioned_iops": "number",
                    "size": "number",
                    "snapshot_id": "string"
                  }
                ]
              ],
              "deployment_set_id": "string",
              "desired_size": "string",
              "image_id": "string",
              "image_type": "string",
              "install_cloud_monitor": "bool",
              "instance_charge_type": "string",
              "instance_types": [
                "list",
                "string"
              ],
              "internet_charge_type": "string",
              "internet_max_bandwidth_out": "number",
              "key_name": "string",
              "kubelet_configuration": [
                "list",
                [
                  "object",
                  {
                    "allowed_unsafe_sysctls": [
                      "list",
                      "string"
                    ],
                    "cluster_dns": [
                      "list",
                      "string"
                    ],
                    "container_log_max_files": "string",
                    "container_log_max_size": "string",
                    "container_log_max_workers": "string",
                    "container_log_monitor_interval": "string",
                    "cpu_cfs_quota": "string",
                    "cpu_cfs_quota_period": "string",
                    "cpu_manager_policy": "string",
                    "event_burst": "string",
                    "event_record_qps": "string",
                    "eviction_hard": [
                      "map",
                      "string"
                    ],
                    "eviction_soft": [
                      "map",
                      "string"
                    ],
                    "eviction_soft_grace_period": [
                      "map",
                      "string"
                    ],
                    "feature_gates": [
                      "map",
                      "string"
                    ],
                    "image_gc_high_threshold_percent": "string",
                    "image_gc_low_threshold_percent": "string",
                    "kube_api_burst": "string",
                    "kube_api_qps": "string",
                    "kube_reserved": [
                      "map",
                      "string"
                    ],
                    "max_pods": "string",
                    "memory_manager_policy": "string",
                    "pod_pids_limit": "string",
                    "read_only_port": "string",
                    "registry_burst": "string",
                    "registry_pull_qps": "string",
                    "reserved_memory": [
                      "list",
                      [
                        "object",
                        {
                          "limits": [
                            "map",
                            "string"
                          ],
                          "numa_node": "number"
                        }
                      ]
                    ],
                    "serialize_image_pulls": "string",
                    "system_reserved": [
                      "map",
                      "string"
                    ],
                    "topology_manager_policy": "string",
                    "tracing": [
                      "list",
                      [
                        "object",
                        {
                          "endpoint": "string",
                          "sampling_rate_per_million": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "labels": [
                "list",
                [
                  "object",
                  {
                    "key": "string",
                    "value": "string"
                  }
                ]
              ],
              "login_as_non_root": "bool",
              "management": [
                "list",
                [
                  "object",
                  {
                    "auto_repair": "bool",
                    "auto_repair_policy": [
                      "list",
                      [
                        "object",
                        {
                          "restart_node": "bool"
                        }
                      ]
                    ],
                    "auto_upgrade": "bool",
                    "auto_upgrade_policy": [
                      "list",
                      [
                        "object",
                        {
                          "auto_upgrade_kubelet": "bool"
                        }
                      ]
                    ],
                    "auto_vul_fix": "bool",
                    "auto_vul_fix_policy": [
                      "list",
                      [
                        "object",
                        {
                          "restart_node": "bool",
                          "vul_level": "string"
                        }
                      ]
                    ],
                    "enable": "bool",
                    "max_unavailable": "number",
                    "surge": "number",
                    "surge_percentage": "number"
                  }
                ]
              ],
              "multi_az_policy": "string",
              "node_name_mode": "string",
              "node_pool_id": "string",
              "node_pool_name": "string",
              "on_demand_base_capacity": "string",
              "on_demand_percentage_above_base_capacity": "string",
              "password": "string",
              "period": "number",
              "period_unit": "string",
              "platform": "string",
              "pre_user_data": "string",
              "private_pool_options": [
                "list",
                [
                  "object",
                  {
                    "private_pool_options_id": "string",
                    "private_pool_options_match_criteria": "string"
                  }
                ]
              ],
              "ram_role_name": "string",
              "rds_instances": [
                "list",
                "string"
              ],
              "resource_group_id": "string",
              "runtime_name": "string",
              "runtime_version": "string",
              "scaling_config": [
                "list",
                [
                  "object",
                  {
                    "eip_bandwidth": "number",
                    "eip_internet_charge_type": "string",
                    "enable": "bool",
                    "is_bond_eip": "bool",
                    "max_size": "number",
                    "min_size": "number",
                    "type": "string"
                  }
                ]
              ],
              "scaling_group_id": "string",
              "scaling_policy": "string",
              "security_group_id": "string",
              "security_group_ids": [
                "list",
                "string"
              ],
              "security_hardening_os": "bool",
              "soc_enabled": "bool",
              "spot_instance_pools": "number",
              "spot_instance_remedy": "bool",
              "spot_price_limit": [
                "list",
                [
                  "object",
                  {
                    "instance_type": "string",
                    "price_limit": "string"
                  }
                ]
              ],
              "spot_strategy": "string",
              "system_disk_bursting_enabled": "bool",
              "system_disk_categories": [
                "list",
                "string"
              ],
              "system_disk_category": "string",
              "system_disk_encrypt_algorithm": "string",
              "system_disk_encrypted": "bool",
              "system_disk_kms_key": "string",
              "system_disk_performance_level": "string",
              "system_disk_provisioned_iops": "number",
              "system_disk_size": "number",
              "system_disk_snapshot_policy_id": "string",
              "tags": [
                "map",
                "string"
              ],
              "taints": [
                "list",
                [
                  "object",
                  {
                    "effect": "string",
                    "key": "string",
                    "value": "string"
                  }
                ]
              ],
              "tee_config": [
                "list",
                [
                  "object",
                  {
                    "tee_enable": "bool"
                  }
                ]
              ],
              "unschedulable": "bool",
              "user_data": "string",
              "vswitch_ids": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCsKubernetesNodePoolsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCsKubernetesNodePools), &result)
	return &result
}
