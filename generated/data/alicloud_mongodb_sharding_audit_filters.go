package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMongodbShardingAuditFilters = `{
  "block": {
    "attributes": {
      "db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "audit_status": "string",
              "db_instance_id": "string",
              "filter": "string",
              "hot_storage_period": "number",
              "id": "string",
              "region_id": "string",
              "role_type": "string",
              "service_type": "string",
              "storage_period": "number"
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMongodbShardingAuditFiltersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMongodbShardingAuditFilters), &result)
	return &result
}
