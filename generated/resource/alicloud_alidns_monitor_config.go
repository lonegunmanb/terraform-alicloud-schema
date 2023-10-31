package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlidnsMonitorConfig = `{
  "block": {
    "attributes": {
      "addr_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "evaluation_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interval": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "monitor_extend_info": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "protocol_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "timeout": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "block_types": {
      "isp_city_node": {
        "block": {
          "attributes": {
            "city_code": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "isp_code": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlidnsMonitorConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlidnsMonitorConfig), &result)
	return &result
}
