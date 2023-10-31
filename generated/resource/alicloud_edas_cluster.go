package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEdasCluster = `{
  "block": {
    "attributes": {
      "cluster_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_type": {
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
      "logical_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEdasClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEdasCluster), &result)
	return &result
}
