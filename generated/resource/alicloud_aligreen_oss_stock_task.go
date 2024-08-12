package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAligreenOssStockTask = `{
  "block": {
    "attributes": {
      "audio_antispam_freeze_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "audio_auto_freeze_opened": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "audio_max_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "audio_opened": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "audio_scan_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "audio_scenes": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_freeze_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "biz_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "buckets": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "callback_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "end_date": {
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
      "image_ad_freeze_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_auto_freeze_opened": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "image_live_freeze_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_opened": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "image_porn_freeze_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_scan_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "image_scenes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "image_terrorism_freeze_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scan_image_no_file_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "start_date": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "video_ad_freeze_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "video_auto_freeze_opened": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "video_frame_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "video_live_freeze_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "video_max_frames": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "video_max_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "video_opened": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "video_porn_freeze_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "video_scan_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "video_scenes": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "video_terrorism_freeze_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "video_voice_antispam_freeze_config": {
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

func AlicloudAligreenOssStockTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAligreenOssStockTask), &result)
	return &result
}
