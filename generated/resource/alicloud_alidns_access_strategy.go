package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlidnsAccessStrategy = `{
  "block": {
    "attributes": {
      "access_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_addr_pool_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_latency_optimization": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_lba_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_max_return_addr_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "default_min_available_addr_num": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "failover_addr_pool_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failover_latency_optimization": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failover_lba_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failover_max_return_addr_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "failover_min_available_addr_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "strategy_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "strategy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "default_addr_pools": {
        "block": {
          "attributes": {
            "addr_pool_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "lba_weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "failover_addr_pools": {
        "block": {
          "attributes": {
            "addr_pool_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lba_weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "lines": {
        "block": {
          "attributes": {
            "line_code": {
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
  "version": 0
}`

func AlicloudAlidnsAccessStrategySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlidnsAccessStrategy), &result)
	return &result
}
