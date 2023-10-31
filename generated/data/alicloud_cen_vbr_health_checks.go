package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenVbrHealthChecks = `{
  "block": {
    "attributes": {
      "cen_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "checks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cen_id": "string",
              "health_check_interval": "number",
              "health_check_source_ip": "string",
              "health_check_target_ip": "string",
              "healthy_threshold": "number",
              "id": "string",
              "vbr_instance_id": "string",
              "vbr_instance_region_id": "string"
            }
          ]
        ]
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
      "vbr_instance_id": {
        "description_kind": "plain",
        "optional": true,
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCenVbrHealthChecksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenVbrHealthChecks), &result)
	return &result
}
