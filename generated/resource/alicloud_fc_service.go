package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcService = `{
  "block": {
    "attributes": {
      "description": {
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
      "internet_access": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "last_modified": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publish": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_id": {
        "computed": true,
        "description_kind": "plain",
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
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "log_config": {
        "block": {
          "attributes": {
            "enable_instance_metrics": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_request_metrics": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "logstore": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "project": {
              "description_kind": "plain",
              "required": true,
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
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "user_id": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "mount_points": {
              "block": {
                "attributes": {
                  "mount_dir": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "server_addr": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "tracing_config": {
        "block": {
          "attributes": {
            "params": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "vpc_config": {
        "block": {
          "attributes": {
            "security_group_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vpc_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vswitch_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
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

func AlicloudFcServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcService), &result)
	return &result
}
