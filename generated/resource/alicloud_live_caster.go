package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudLiveCaster = `{
  "block": {
    "attributes": {
      "auto_switch_urgent_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_switch_urgent_on": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "callback_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "caster_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delay": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "domain_name": {
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
      "norm_type": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "program_effect": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "program_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "record_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "side_output_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "side_output_url_list": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sync_groups_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "transcode_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "urgent_image_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "urgent_image_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "urgent_live_stream_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "urgent_material_id": {
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

func AlicloudLiveCasterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudLiveCaster), &result)
	return &result
}
