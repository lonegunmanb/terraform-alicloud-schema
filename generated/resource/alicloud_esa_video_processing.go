package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaVideoProcessing = `{
  "block": {
    "attributes": {
      "config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "flv_seek_end": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flv_seek_start": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flv_video_seek_mode": {
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
      "mp4_seek_end": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mp4_seek_start": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sequence": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "video_seek_enable": {
        "description_kind": "plain",
        "optional": true,
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

func AlicloudEsaVideoProcessingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaVideoProcessing), &result)
	return &result
}
