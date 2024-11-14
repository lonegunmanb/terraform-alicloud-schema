package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGwlbServerGroup = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dry_run": {
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
      "protocol": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduler": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_group_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
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
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "connection_drain_config": {
        "block": {
          "attributes": {
            "connection_drain_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "connection_drain_timeout": {
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
      },
      "health_check_config": {
        "block": {
          "attributes": {
            "health_check_connect_port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "health_check_connect_timeout": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "health_check_domain": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "health_check_http_code": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "health_check_interval": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "health_check_path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_protocol": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "healthy_threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "unhealthy_threshold": {
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
      },
      "servers": {
        "block": {
          "attributes": {
            "port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "server_group_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "server_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "server_ip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func AlicloudGwlbServerGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGwlbServerGroup), &result)
	return &result
}
