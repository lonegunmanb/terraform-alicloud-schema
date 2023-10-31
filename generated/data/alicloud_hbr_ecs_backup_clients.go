package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrEcsBackupClients = `{
  "block": {
    "attributes": {
      "clients": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arch_type": "string",
              "backup_status": "string",
              "client_type": "string",
              "client_version": "string",
              "create_time": "string",
              "data_network_type": "string",
              "data_proxy_setting": "string",
              "ecs_backup_client_id": "string",
              "hostname": "string",
              "id": "string",
              "instance_id": "string",
              "instance_name": "string",
              "last_heart_beat_time": "string",
              "max_client_version": "string",
              "max_cpu_core": "string",
              "max_worker": "string",
              "os_type": "string",
              "private_ipv4": "string",
              "proxy_host": "string",
              "proxy_password": "string",
              "proxy_port": "string",
              "proxy_user": "string",
              "status": "string",
              "updated_time": "string",
              "use_https": "bool",
              "zone_id": "string"
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
      "instance_ids": {
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
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrEcsBackupClientsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrEcsBackupClients), &result)
	return &result
}
