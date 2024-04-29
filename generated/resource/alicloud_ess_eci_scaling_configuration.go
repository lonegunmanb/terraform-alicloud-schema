package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssEciScalingConfiguration = `{
  "block": {
    "attributes": {
      "active": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "active_deadline_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "auto_create_eip": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_match_image_cache": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "container_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cpu": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "egress_bandwidth": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "eip_bandwidth": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enable_sls": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ephemeral_storage": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "force_delete": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "host_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_snapshot_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ingress_bandwidth": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "instance_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ipv6_address_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "load_balancer_weight": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "memory": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ram_role_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restart_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_configuration_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_price_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "spot_strategy": {
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
      "termination_grace_period_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "acr_registry_infos": {
        "block": {
          "attributes": {
            "domains": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "instance_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "containers": {
        "block": {
          "attributes": {
            "args": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "commands": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "cpu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "gpu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "image": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image_pull_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lifecycle_pre_stop_handler_execs": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "liveness_probe_exec_commands": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "liveness_probe_failure_threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "liveness_probe_http_get_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "liveness_probe_http_get_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "liveness_probe_http_get_scheme": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "liveness_probe_initial_delay_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "liveness_probe_period_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "liveness_probe_success_threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "liveness_probe_tcp_socket_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "liveness_probe_timeout_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "memory": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "readiness_probe_exec_commands": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "readiness_probe_failure_threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "readiness_probe_http_get_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "readiness_probe_http_get_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "readiness_probe_http_get_scheme": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "readiness_probe_initial_delay_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "readiness_probe_period_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "readiness_probe_success_threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "readiness_probe_tcp_socket_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "readiness_probe_timeout_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "security_context_capability_adds": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "security_context_read_only_root_file_system": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "security_context_run_as_user": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "working_dir": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "environment_vars": {
              "block": {
                "attributes": {
                  "field_ref_field_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "ports": {
              "block": {
                "attributes": {
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "volume_mounts": {
              "block": {
                "attributes": {
                  "mount_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "read_only": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "host_aliases": {
        "block": {
          "attributes": {
            "hostnames": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "image_registry_credentials": {
        "block": {
          "attributes": {
            "password": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "username": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "init_containers": {
        "block": {
          "attributes": {
            "args": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "commands": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "cpu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "gpu": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "image": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image_pull_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "memory": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_context_capability_adds": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "security_context_read_only_root_file_system": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "security_context_run_as_user": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "working_dir": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "environment_vars": {
              "block": {
                "attributes": {
                  "field_ref_field_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "ports": {
              "block": {
                "attributes": {
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "volume_mounts": {
              "block": {
                "attributes": {
                  "mount_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "read_only": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      },
      "volumes": {
        "block": {
          "attributes": {
            "disk_volume_disk_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_volume_disk_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_volume_fs_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "flex_volume_driver": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "flex_volume_fs_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "flex_volume_options": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nfs_volume_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nfs_volume_read_only": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "nfs_volume_server": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "config_file_volume_config_file_to_paths": {
              "block": {
                "attributes": {
                  "content": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssEciScalingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssEciScalingConfiguration), &result)
	return &result
}
