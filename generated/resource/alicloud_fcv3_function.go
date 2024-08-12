package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcv3Function = `{
  "block": {
    "attributes": {
      "cpu": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "environment_variables": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "function_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "handler": {
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
      "instance_concurrency": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "internet_access": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "layers": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": [
          "list",
          "string"
        ]
      },
      "memory_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "timeout": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "code": {
        "block": {
          "attributes": {
            "checksum": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "oss_bucket_name": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "oss_object_name": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "zip_file": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "custom_container_config": {
        "block": {
          "attributes": {
            "acceleration_info": {
              "computed": true,
              "deprecated": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "status": "string"
                  }
                ]
              ]
            },
            "acceleration_type": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "acr_instance_id": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "command": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "entrypoint": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "image": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "resolved_image_uri": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "health_check_config": {
              "block": {
                "attributes": {
                  "failure_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "http_get_url": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "initial_delay_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "period_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "success_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "timeout_seconds": {
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
      "custom_dns": {
        "block": {
          "attributes": {
            "name_servers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "searches": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "dns_options": {
              "block": {
                "attributes": {
                  "name": {
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "custom_runtime_config": {
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
            "command": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "health_check_config": {
              "block": {
                "attributes": {
                  "failure_threshold": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "http_get_url": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "initial_delay_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "period_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "success_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "timeout_seconds": {
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
      "gpu_config": {
        "block": {
          "attributes": {
            "gpu_memory_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "gpu_type": {
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
      "instance_lifecycle_config": {
        "block": {
          "block_types": {
            "initializer": {
              "block": {
                "attributes": {
                  "handler": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "pre_stop": {
              "block": {
                "attributes": {
                  "handler": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timeout": {
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
      "log_config": {
        "block": {
          "attributes": {
            "enable_instance_metrics": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_request_metrics": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "log_begin_rule": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "logstore": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project": {
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
      "nas_config": {
        "block": {
          "attributes": {
            "group_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "user_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "mount_points": {
              "block": {
                "attributes": {
                  "enable_tls": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "mount_dir": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "server_addr": {
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
      "oss_mount_config": {
        "block": {
          "block_types": {
            "mount_points": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bucket_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "endpoint": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mount_dir": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      },
      "vpc_config": {
        "block": {
          "attributes": {
            "security_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vswitch_ids": {
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
  "version": 0
}`

func AlicloudFcv3FunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcv3Function), &result)
	return &result
}
