package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudStorageGatewayGatewayFileShare = `{
  "block": {
    "attributes": {
      "access_based_enumeration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "backend_limit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "browsable": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bypass_cache_read": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cache_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "direct_io": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "download_limit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "fast_reclaim": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fe_limit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "gateway_file_share_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "ignore_delete": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "in_place": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "index_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lag_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "local_path": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nfs_v4_optimization": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "oss_bucket_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "oss_bucket_ssl": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "oss_endpoint": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "partial_sync_paths": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "polling_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remote_sync": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "remote_sync_download": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ro_client_list": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ro_user_list": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rw_client_list": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rw_user_list": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "squash": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "support_archive": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "transfer_acceleration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "windows_acl": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudStorageGatewayGatewayFileShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudStorageGatewayGatewayFileShare), &result)
	return &result
}
