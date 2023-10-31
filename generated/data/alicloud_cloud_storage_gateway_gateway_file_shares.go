package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudStorageGatewayGatewayFileShares = `{
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
      "shares": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_based_enumeration": "bool",
              "address": "string",
              "backend_limit": "number",
              "browsable": "bool",
              "bucket_infos": "string",
              "buckets_stub": "bool",
              "bypass_cache_read": "bool",
              "cache_mode": "string",
              "client_side_cmk": "string",
              "client_side_encryption": "bool",
              "direct_io": "bool",
              "disk_id": "string",
              "disk_type": "string",
              "download_limit": "number",
              "enabled": "bool",
              "express_sync_id": "string",
              "fast_reclaim": "bool",
              "fe_limit": "number",
              "file_num_limit": "string",
              "fs_size_limit": "string",
              "gateway_file_share_name": "string",
              "gateway_id": "string",
              "id": "string",
              "ignore_delete": "bool",
              "in_place": "bool",
              "in_rate": "string",
              "index_id": "string",
              "kms_rotate_period": "number",
              "lag_period": "string",
              "local_path": "string",
              "mns_health": "string",
              "nfs_v4_optimization": "bool",
              "obsolete_buckets": "string",
              "oss_bucket_name": "string",
              "oss_bucket_ssl": "bool",
              "oss_endpoint": "string",
              "oss_health": "string",
              "oss_used": "string",
              "out_rate": "string",
              "partial_sync_paths": "string",
              "path_prefix": "string",
              "polling_interval": "number",
              "protocol": "string",
              "remaining_meta_space": "string",
              "remote_sync": "bool",
              "remote_sync_download": "bool",
              "ro_client_list": "string",
              "ro_user_list": "string",
              "rw_client_list": "string",
              "rw_user_list": "string",
              "server_side_cmk": "string",
              "server_side_encryption": "bool",
              "size": "string",
              "squash": "string",
              "state": "string",
              "support_archive": "bool",
              "sync_progress": "number",
              "total_download": "string",
              "total_upload": "string",
              "transfer_acceleration": "bool",
              "used": "string",
              "windows_acl": "bool"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudStorageGatewayGatewayFileSharesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudStorageGatewayGatewayFileShares), &result)
	return &result
}
