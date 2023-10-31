package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlidnsAccessStrategies = `{
  "block": {
    "attributes": {
      "enable_details": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "strategies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_mode": "string",
              "access_strategy_id": "string",
              "create_time": "string",
              "create_timestamp": "string",
              "default_addr_pool_type": "string",
              "default_addr_pools": [
                "list",
                [
                  "object",
                  {
                    "addr_count": "number",
                    "addr_pool_id": "string",
                    "lba_weight": "number",
                    "name": "string"
                  }
                ]
              ],
              "default_available_addr_num": "number",
              "default_latency_optimization": "string",
              "default_lba_strategy": "string",
              "default_max_return_addr_num": "number",
              "default_min_available_addr_num": "number",
              "effective_addr_pool_group_type": "string",
              "failover_addr_pool_type": "string",
              "failover_addr_pools": [
                "list",
                [
                  "object",
                  {
                    "addr_count": "number",
                    "addr_pool_id": "string",
                    "lba_weight": "number",
                    "name": "string"
                  }
                ]
              ],
              "failover_available_addr_num": "number",
              "failover_latency_optimization": "string",
              "failover_lba_strategy": "string",
              "failover_max_return_addr_num": "number",
              "failover_min_available_addr_num": "number",
              "id": "string",
              "instance_id": "string",
              "lines": [
                "list",
                [
                  "object",
                  {
                    "group_code": "string",
                    "group_name": "string",
                    "line_code": "string",
                    "line_name": "string"
                  }
                ]
              ],
              "strategy_mode": "string",
              "strategy_name": "string"
            }
          ]
        ]
      },
      "strategy_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlidnsAccessStrategiesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlidnsAccessStrategies), &result)
	return &result
}
