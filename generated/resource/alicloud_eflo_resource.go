package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEfloResource = `{
  "block": {
    "attributes": {
      "cluster_desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_name": {
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
      "resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "machine_types": {
        "block": {
          "attributes": {
            "bond_num": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "cpu_info": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "disk_info": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gpu_info": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "memory_info": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network_info": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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
      "user_access_param": {
        "block": {
          "attributes": {
            "access_id": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "access_key": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "endpoint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "workspace_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEfloResourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEfloResource), &result)
	return &result
}
