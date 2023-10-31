package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDtsMigrationInstance = `{
  "block": {
    "attributes": {
      "compute_unit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "database_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "destination_endpoint_engine_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_endpoint_region": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dts_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_class": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_endpoint_engine_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_endpoint_region": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sync_architecture": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDtsMigrationInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDtsMigrationInstance), &result)
	return &result
}
