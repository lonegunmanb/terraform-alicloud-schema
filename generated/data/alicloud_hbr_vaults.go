package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrVaults = `{
  "block": {
    "attributes": {
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
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
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
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vault_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vaults": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bucket_name": "string",
              "bytes_done": "string",
              "created_time": "string",
              "dedup": "bool",
              "description": "string",
              "id": "string",
              "index_available": "bool",
              "index_level": "string",
              "index_update_time": "string",
              "latest_replication_time": "string",
              "payment_type": "string",
              "replication": "bool",
              "replication_source_region_id": "string",
              "replication_source_vault_id": "string",
              "retention": "string",
              "search_enabled": "bool",
              "source_types": [
                "list",
                "string"
              ],
              "status": "string",
              "storage_size": "string",
              "updated_time": "string",
              "vault_id": "string",
              "vault_name": "string",
              "vault_status_message": "string",
              "vault_storage_class": "string",
              "vault_type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrVaultsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrVaults), &result)
	return &result
}
