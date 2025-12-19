package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMongodbNode = `{
  "block": {
    "attributes": {
      "account_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "account_password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_pay": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "business_info": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "effective_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "from_app": {
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
      "node_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "node_storage": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "node_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "order_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "readonly_replicas": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "shard_direct": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "switch_time": {
        "description_kind": "plain",
        "optional": true,
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

func AlicloudMongodbNodeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMongodbNode), &result)
	return &result
}
