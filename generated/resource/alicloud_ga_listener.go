package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaListener = `{
  "block": {
    "attributes": {
      "accelerator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_affinity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_version": {
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
      "idle_timeout": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "listener_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "proxy_protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "request_timeout": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "security_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "certificates": {
        "block": {
          "attributes": {
            "id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "forwarded_for_config": {
        "block": {
          "attributes": {
            "forwarded_for_ga_ap_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "forwarded_for_ga_id_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "forwarded_for_port_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "forwarded_for_proto_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "real_ip_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "port_ranges": {
        "block": {
          "attributes": {
            "from_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "to_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
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

func AlicloudGaListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaListener), &result)
	return &result
}
