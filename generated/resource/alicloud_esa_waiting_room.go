package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaWaitingRoom = `{
  "block": {
    "attributes": {
      "cookie_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "custom_page_html": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_session_renewal_enable": {
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
      "json_response_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "new_users_per_minute": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "queue_all_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "queuing_method": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "queuing_status_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "session_duration": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "total_active_users": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "waiting_room_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "waiting_room_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "waiting_room_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "host_name_and_path": {
        "block": {
          "attributes": {
            "domain": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "subdomain": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
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

func AlicloudEsaWaitingRoomSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaWaitingRoom), &result)
	return &result
}
