package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDdoscooInstance = `{
  "block": {
    "attributes": {
      "address_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bandwidth_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "base_bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "domain_count": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "edition_sale": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "function_version": {
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
      "ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "modify_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "normal_bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "normal_qps": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "port_count": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_plan": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "product_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
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

func AlicloudDdoscooInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDdoscooInstance), &result)
	return &result
}
