package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOssBucketObjects = `{
  "block": {
    "attributes": {
      "bucket_name": {
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
      "key_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "objects": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acl": "string",
              "cache_control": "string",
              "content_disposition": "string",
              "content_encoding": "string",
              "content_length": "string",
              "content_md5": "string",
              "content_type": "string",
              "etag": "string",
              "expires": "string",
              "key": "string",
              "last_modification_time": "string",
              "server_side_encryption": "string",
              "sse_kms_key_id": "string",
              "storage_class": "string"
            }
          ]
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOssBucketObjectsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOssBucketObjects), &result)
	return &result
}
