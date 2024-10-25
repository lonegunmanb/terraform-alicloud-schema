package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsCollectionPolicy = `{
  "block": {
    "attributes": {
      "centralize_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "data_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "centralize_config": {
        "block": {
          "attributes": {
            "dest_logstore": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dest_project": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dest_region": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dest_ttl": {
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
      "data_config": {
        "block": {
          "attributes": {
            "data_project": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "data_region": {
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
      "policy_config": {
        "block": {
          "attributes": {
            "instance_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "regions": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resource_mode": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "resource_directory": {
        "block": {
          "attributes": {
            "account_group_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "members": {
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

func AlicloudSlsCollectionPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsCollectionPolicy), &result)
	return &result
}
