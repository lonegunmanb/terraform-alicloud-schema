package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlbHealthCheckTemplates = `{
  "block": {
    "attributes": {
      "health_check_template_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "health_check_template_name": {
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
      "templates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "health_check_codes": [
                "list",
                "string"
              ],
              "health_check_connect_port": "number",
              "health_check_host": "string",
              "health_check_http_version": "string",
              "health_check_interval": "number",
              "health_check_method": "string",
              "health_check_path": "string",
              "health_check_protocol": "string",
              "health_check_template_id": "string",
              "health_check_template_name": "string",
              "health_check_timeout": "number",
              "healthy_threshold": "number",
              "id": "string",
              "unhealthy_threshold": "number"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlbHealthCheckTemplatesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlbHealthCheckTemplates), &result)
	return &result
}
