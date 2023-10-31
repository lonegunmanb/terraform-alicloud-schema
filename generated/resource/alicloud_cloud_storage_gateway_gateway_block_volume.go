package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudStorageGatewayGatewayBlockVolume = `{
  "block": {
    "attributes": {
      "cache_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "chap_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "chap_in_password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "chap_in_user": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "chunk_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "gateway_block_volume_name": {
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
      "index_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "is_source_deletion": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "local_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recovery": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func AlicloudCloudStorageGatewayGatewayBlockVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudStorageGatewayGatewayBlockVolume), &result)
	return &result
}
