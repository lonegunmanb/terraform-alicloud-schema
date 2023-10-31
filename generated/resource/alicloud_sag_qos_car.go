package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSagQosCar = `{
  "block": {
    "attributes": {
      "description": {
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
      "limit_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_bandwidth_abs": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_bandwidth_percent": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_bandwidth_abs": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_bandwidth_percent": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "percent_source_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "qos_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSagQosCarSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSagQosCar), &result)
	return &result
}
