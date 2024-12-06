package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSaeApplication = `{
  "block": {
    "attributes": {
      "acr_assume_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "acr_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "app_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "app_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "auto_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_enable_application_scaling_rule": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "batch_wait_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "change_order_desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "command": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "command_args": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "command_args_v2": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "config_map_mount_desc": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cpu": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "custom_host_alias": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deploy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "edas_container_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_ahas": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_grey_tag_route": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "envs": {
        "computed": true,
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
      "image_pull_secrets": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "jar_start_args": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "jar_start_options": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "jdk": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "liveness": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "memory": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "micro_registration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_ready_instance_ratio": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_ready_instances": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "mount_desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mount_host": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nas_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oss_ak_id": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "oss_ak_secret": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "oss_mount_descs": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "package_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "php": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "php_arms_config_location": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "php_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "php_config_location": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "post_start": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pre_stop": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "programming_language": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "readiness": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replicas": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sls_configs": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "timezone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tomcat_config": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_strategy": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "war_start_options": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "web_container": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "config_map_mount_desc_v2": {
        "block": {
          "attributes": {
            "config_map_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "custom_host_alias_v2": {
        "block": {
          "attributes": {
            "host_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "kafka_configs": {
        "block": {
          "attributes": {
            "kafka_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kafka_instance_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "kafka_configs": {
              "block": {
                "attributes": {
                  "kafka_topic": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_dir": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_type": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "liveness_v2": {
        "block": {
          "attributes": {
            "initial_delay_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "period_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "exec": {
              "block": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "http_get": {
              "block": {
                "attributes": {
                  "is_contain_key_word": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "key_word": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "scheme": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "tcp_socket": {
              "block": {
                "attributes": {
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "nas_configs": {
        "block": {
          "attributes": {
            "mount_domain": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nas_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nas_path": {
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
        "nesting_mode": "list"
      },
      "oss_mount_descs_v2": {
        "block": {
          "attributes": {
            "bucket_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bucket_path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read_only": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "post_start_v2": {
        "block": {
          "block_types": {
            "exec": {
              "block": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "pre_stop_v2": {
        "block": {
          "block_types": {
            "exec": {
              "block": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "pvtz_discovery_svc": {
        "block": {
          "attributes": {
            "enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "namespace_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "port_protocols": {
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "readiness_v2": {
        "block": {
          "attributes": {
            "initial_delay_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "period_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "exec": {
              "block": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "http_get": {
              "block": {
                "attributes": {
                  "is_contain_key_word": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "key_word": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "scheme": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "tcp_socket": {
              "block": {
                "attributes": {
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "tomcat_config_v2": {
        "block": {
          "attributes": {
            "context_path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_threads": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "uri_encoding": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_body_encoding_for_uri": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "update_strategy_v2": {
        "block": {
          "attributes": {
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "batch_update": {
              "block": {
                "attributes": {
                  "batch": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "batch_wait_time": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "release_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSaeApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSaeApplication), &result)
	return &result
}
