package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectTrafficQosQueue = `{
  "block": {
    "attributes": {
      "bandwidth_percent": {
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
      "qos_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "queue_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "queue_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "queue_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "queue_type": {
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

func AlicloudExpressConnectTrafficQosQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectTrafficQosQueue), &result)
	return &result
}
