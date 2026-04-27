package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlidnsCloudGtmInstanceConfig = `{
  "block": {
    "attributes": {
      "address_pool_lb_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_status": {
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
      "instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "remark": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule_host_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule_rr_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_zone_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_zone_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sequence_lb_strategy_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ttl": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
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

func AlicloudAlidnsCloudGtmInstanceConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlidnsCloudGtmInstanceConfig), &result)
	return &result
}
