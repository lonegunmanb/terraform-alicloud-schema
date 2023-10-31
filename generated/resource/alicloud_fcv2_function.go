package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcv2Function = `{
  "block": {
    "attributes": {
      "ca_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "code_checksum": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cpu": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gpu_memory_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "initialization_timeout": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "initializer": {
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
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "layers": {
        "description_kind": "plain",
        "optional": true,
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
      "runtime": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_name": {
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
            "oss_bucket_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "oss_object_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zip_file": {
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
      "custom_container_config": {
        "block": {
          "attributes": {
            "acceleration_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "args": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "command": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "web_server_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
      "custom_health_check_config": {
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
            "pre_freeze": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudFcv2FunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcv2Function), &result)
	return &result
}
