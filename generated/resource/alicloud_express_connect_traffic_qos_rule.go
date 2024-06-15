package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectTrafficQosRule = `{
  "block": {
    "attributes": {
      "dst_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dst_ipv6_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dst_port_range": {
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
      "match_dscp": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "priority": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "qos_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "queue_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remarking_dscp": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "rule_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rule_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "src_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "src_ipv6_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "src_port_range": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func AlicloudExpressConnectTrafficQosRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectTrafficQosRule), &result)
	return &result
}
