package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOssBucketInventories = `{
  "block": {
    "attributes": {
      "bucket": {
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
      "inventories": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "destination": [
                "list",
                [
                  "object",
                  {
                    "oss_bucket_destination": [
                      "list",
                      [
                        "object",
                        {
                          "account_id": "string",
                          "bucket": "string",
                          "encryption": [
                            "list",
                            [
                              "object",
                              {
                                "ssekms": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "key_id": "string"
                                    }
                                  ]
                                ],
                                "sseoss": "string"
                              }
                            ]
                          ],
                          "format": "string",
                          "prefix": "string",
                          "role_arn": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "filter": [
                "list",
                [
                  "object",
                  {
                    "last_modify_begin_time_stamp": "number",
                    "last_modify_end_time_stamp": "number",
                    "lower_size_bound": "number",
                    "prefix": "string",
                    "storage_class": "string",
                    "upper_size_bound": "number"
                  }
                ]
              ],
              "id": "string",
              "included_object_versions": "string",
              "incremental_inventory": [
                "list",
                [
                  "object",
                  {
                    "is_enabled": "bool",
                    "optional_fields": [
                      "list",
                      [
                        "object",
                        {
                          "field": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "schedule": [
                      "list",
                      [
                        "object",
                        {
                          "frequency": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "inventory_id": "string",
              "is_enabled": "bool",
              "optional_fields": [
                "list",
                [
                  "object",
                  {
                    "field": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "schedule": [
                "list",
                [
                  "object",
                  {
                    "frequency": "string"
                  }
                ]
              ]
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

func AlicloudOssBucketInventoriesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOssBucketInventories), &result)
	return &result
}
