package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenVbrHealthCheck = `{
  "block": {
    "attributes": {
      "cen_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "health_check_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_source_ip": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_target_ip": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "healthy_threshold": {
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
      "vbr_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vbr_instance_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "vbr_instance_region_id": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudCenVbrHealthCheckSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenVbrHealthCheck), &result)
	return &result
}
