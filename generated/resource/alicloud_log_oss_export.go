package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudLogOssExport = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "buffer_interval": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "buffer_size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "compress_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "csv_config_columns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "csv_config_delimiter": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "csv_config_escape": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "csv_config_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "csv_config_linefeed": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "csv_config_null": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "csv_config_quote": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "export_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "from_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "json_enable_tag": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "log_read_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logstore_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "path_format": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "suffix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_zone": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "config_columns": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func AlicloudLogOssExportSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudLogOssExport), &result)
	return &result
}
