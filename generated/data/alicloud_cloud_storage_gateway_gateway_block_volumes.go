package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudStorageGatewayGatewayBlockVolumes = `{
  "block": {
    "attributes": {
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
        "type": "number"
      },
      "volumes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "cache_mode": "string",
              "chap_enabled": "bool",
              "chap_in_user": "string",
              "chunk_size": "number",
              "disk_id": "string",
              "disk_type": "string",
              "enabled": "bool",
              "gateway_block_volume_name": "string",
              "gateway_id": "string",
              "id": "string",
              "index_id": "string",
              "local_path": "string",
              "lun_id": "number",
              "oss_bucket_name": "string",
              "oss_bucket_ssl": "bool",
              "oss_endpoint": "string",
              "port": "number",
              "protocol": "string",
              "size": "number",
              "state": "string",
              "status": "number",
              "target": "string",
              "total_download": "number",
              "total_upload": "number",
              "volume_state": "number"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudStorageGatewayGatewayBlockVolumesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudStorageGatewayGatewayBlockVolumes), &result)
	return &result
}
