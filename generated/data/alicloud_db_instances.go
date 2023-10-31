package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDbInstances = `{
  "block": {
    "attributes": {
      "connection_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "engine": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acl": "string",
              "availability_zone": "string",
              "ca_type": "string",
              "charge_type": "string",
              "client_ca_cert": "string",
              "client_ca_cert_expire_time": "string",
              "client_cert_revocation_list": "string",
              "connection_mode": "string",
              "connection_string": "string",
              "create_time": "string",
              "creator": "string",
              "db_instance_storage_type": "string",
              "db_instance_type": "string",
              "db_type": "string",
              "delete_date": "string",
              "deletion_protection": "bool",
              "description": "string",
              "encryption_key": "string",
              "encryption_key_status": "string",
              "engine": "string",
              "engine_version": "string",
              "expire_time": "string",
              "guard_instance_id": "string",
              "ha_mode": "string",
              "host_instance_infos": [
                "list",
                [
                  "object",
                  {
                    "data_sync_time": "string",
                    "log_sync_time": "string",
                    "node_id": "string",
                    "node_type": "string",
                    "region_id": "string",
                    "sync_status": "string",
                    "zone_id": "string"
                  }
                ]
              ],
              "id": "string",
              "instance_storage": "number",
              "instance_type": "string",
              "key_usage": "string",
              "last_modify_status": "string",
              "master_instance_id": "string",
              "master_zone": "string",
              "material_expire_time": "string",
              "modify_status_reason": "string",
              "name": "string",
              "net_type": "string",
              "origin": "string",
              "parameters": [
                "list",
                [
                  "object",
                  {
                    "checking_code": "string",
                    "force_modify": "string",
                    "force_restart": "string",
                    "parameter_description": "string",
                    "parameter_name": "string",
                    "parameter_value": "string"
                  }
                ]
              ],
              "port": "string",
              "readonly_instance_ids": [
                "list",
                "string"
              ],
              "region_id": "string",
              "replication_acl": "string",
              "require_update": "string",
              "require_update_item": "string",
              "require_update_reason": "string",
              "server_ca_url": "string",
              "server_cert": "string",
              "server_key": "string",
              "ssl_create_time": "string",
              "ssl_enabled": "string",
              "ssl_expire_time": "string",
              "status": "string",
              "sync_mode": "string",
              "temp_instance_id": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id_slave_a": "string",
              "zone_id_slave_b": "string"
            }
          ]
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
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "total_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDbInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDbInstances), &result)
	return &result
}
