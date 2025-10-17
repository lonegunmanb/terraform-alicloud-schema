package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudSsoUserProvisioningEvents = `{
  "block": {
    "attributes": {
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "events": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "content": "string",
              "create_time": "string",
              "deletion_strategy": "string",
              "directory_id": "string",
              "duplication_strategy": "string",
              "error_count": "number",
              "error_info": "string",
              "event_id": "string",
              "id": "string",
              "last_sync_time": "string",
              "principal_id": "string",
              "principal_name": "string",
              "principal_type": "string",
              "source_type": "string",
              "target_id": "string",
              "target_name": "string",
              "target_path": "string",
              "target_type": "string",
              "update_time": "string",
              "user_provisioning_id": "string"
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
        "optional": true,
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
      "user_provisioning_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudSsoUserProvisioningEventsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudSsoUserProvisioningEvents), &result)
	return &result
}
