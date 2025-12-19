package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudConfigReportTemplate = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "report_file_formats": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "report_granularity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "report_language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "report_template_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "report_template_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subscription_frequency": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "report_scope": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "match_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
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

func AlicloudConfigReportTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudConfigReportTemplate), &result)
	return &result
}
