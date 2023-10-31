package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdPolicyGroup = `{
  "block": {
    "attributes": {
      "camera_redirect": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "clipboard": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_list": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "html_access": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "html_file_transfer": {
        "computed": true,
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
      "local_drive": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recording": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recording_end_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recording_expires": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "recording_fps": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "recording_start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "usb_redirect": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "visual_quality": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "watermark": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "watermark_transparency": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "watermark_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "authorize_access_policy_rules": {
        "block": {
          "attributes": {
            "cidr_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "authorize_security_policy_rules": {
        "block": {
          "attributes": {
            "cidr_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port_range": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "priority": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcdPolicyGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdPolicyGroup), &result)
	return &result
}
