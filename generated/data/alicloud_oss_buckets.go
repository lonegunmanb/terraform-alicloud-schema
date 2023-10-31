package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOssBuckets = `{
  "block": {
    "attributes": {
      "buckets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acl": "string",
              "cors_rules": [
                "list",
                [
                  "object",
                  {
                    "allowed_headers": [
                      "list",
                      "string"
                    ],
                    "allowed_methods": [
                      "list",
                      "string"
                    ],
                    "allowed_origins": [
                      "list",
                      "string"
                    ],
                    "expose_headers": [
                      "list",
                      "string"
                    ],
                    "max_age_seconds": "number"
                  }
                ]
              ],
              "creation_date": "string",
              "extranet_endpoint": "string",
              "intranet_endpoint": "string",
              "lifecycle_rule": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool",
                    "expiration": [
                      "list",
                      [
                        "object",
                        {
                          "date": "string",
                          "days": "number"
                        }
                      ]
                    ],
                    "id": "string",
                    "prefix": "string"
                  }
                ]
              ],
              "location": "string",
              "logging": [
                "list",
                [
                  "object",
                  {
                    "target_bucket": "string",
                    "target_prefix": "string"
                  }
                ]
              ],
              "name": "string",
              "owner": "string",
              "policy": "string",
              "redundancy_type": "string",
              "referer_config": [
                "list",
                [
                  "object",
                  {
                    "allow_empty": "bool",
                    "referers": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "server_side_encryption_rule": [
                "list",
                [
                  "object",
                  {
                    "kms_master_key_id": "string",
                    "sse_algorithm": "string"
                  }
                ]
              ],
              "storage_class": "string",
              "tags": [
                "map",
                "string"
              ],
              "versioning": [
                "list",
                [
                  "object",
                  {
                    "status": "string"
                  }
                ]
              ],
              "website": [
                "list",
                [
                  "object",
                  {
                    "error_document": "string",
                    "index_document": "string"
                  }
                ]
              ]
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOssBucketsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOssBuckets), &result)
	return &result
}
