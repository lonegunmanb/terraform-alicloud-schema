package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEdasApplication = `{
  "block": {
    "attributes": {
      "application_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "build_pack_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "descriotion": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecu_info": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_url": {
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
      "logical_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "package_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "war_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEdasApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEdasApplication), &result)
	return &result
}
