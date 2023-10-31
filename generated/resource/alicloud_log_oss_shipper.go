package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudLogOssShipper = `{
  "block": {
    "attributes": {
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
        "description_kind": "plain",
        "optional": true,
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
      "csv_config_nullidentifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "csv_config_quote": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "format": {
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
      "json_enable_tag": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "logstore_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "oss_bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "oss_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path_format": {
        "description_kind": "plain",
        "required": true,
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
      "shipper_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "parquet_config": {
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

func AlicloudLogOssShipperSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudLogOssShipper), &result)
	return &result
}
