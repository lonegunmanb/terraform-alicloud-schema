package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudActiontrails = `{
  "block": {
    "attributes": {
      "actiontrails": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "event_rw": "string",
              "id": "string",
              "is_organization_trail": "bool",
              "oss_bucket_name": "string",
              "oss_key_prefix": "string",
              "oss_write_role_arn": "string",
              "sls_project_arn": "string",
              "sls_write_role_arn": "string",
              "status": "string",
              "trail_name": "string",
              "trail_region": "string"
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
      "include_organization_trail": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "include_shadow_trails": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
        "type": "string"
      },
      "trails": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "event_rw": "string",
              "id": "string",
              "is_organization_trail": "bool",
              "oss_bucket_name": "string",
              "oss_key_prefix": "string",
              "oss_write_role_arn": "string",
              "sls_project_arn": "string",
              "sls_write_role_arn": "string",
              "status": "string",
              "trail_name": "string",
              "trail_region": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudActiontrailsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudActiontrails), &result)
	return &result
}
