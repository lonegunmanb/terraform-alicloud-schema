package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEciImageCaches = `{
  "block": {
    "attributes": {
      "caches": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "container_group_id": "string",
              "events": [
                "list",
                [
                  "object",
                  {
                    "count": "number",
                    "first_timestamp": "string",
                    "last_timestamp": "string",
                    "message": "string",
                    "name": "string",
                    "type": "string"
                  }
                ]
              ],
              "expire_date_time": "string",
              "id": "string",
              "image_cache_id": "string",
              "image_cache_name": "string",
              "images": [
                "list",
                "string"
              ],
              "progress": "string",
              "snapshot_id": "string",
              "status": "string"
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
      "image": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_cache_name": {
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
      },
      "snapshot_id": {
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

func AlicloudEciImageCachesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEciImageCaches), &result)
	return &result
}
