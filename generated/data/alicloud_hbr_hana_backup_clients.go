package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrHanaBackupClients = `{
  "block": {
    "attributes": {
      "client_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hana_backup_clients": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alert_setting": "string",
              "client_id": "string",
              "client_name": "string",
              "client_type": "string",
              "client_version": "string",
              "cluster_id": "string",
              "id": "string",
              "instance_id": "string",
              "instance_name": "string",
              "max_version": "string",
              "network_type": "string",
              "status": "string",
              "status_message": "string",
              "use_https": "bool",
              "vault_id": "string"
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
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrHanaBackupClientsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrHanaBackupClients), &result)
	return &result
}
