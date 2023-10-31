package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudStorageGatewayGatewayCacheDisks = `{
  "block": {
    "attributes": {
      "disks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cache_disk_category": "string",
              "cache_disk_size_in_gb": "number",
              "cache_id": "string",
              "expired_time": "number",
              "gateway_id": "string",
              "id": "string",
              "iops": "number",
              "is_used": "bool",
              "local_file_path": "string",
              "renew_url": "string",
              "status": "number"
            }
          ]
        ]
      },
      "gateway_id": {
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
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudStorageGatewayGatewayCacheDisksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudStorageGatewayGatewayCacheDisks), &result)
	return &result
}
