package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEdasK8SSlbAttachment = `{
  "block": {
    "attributes": {
      "app_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "slb_configs": {
        "block": {
          "attributes": {
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "scheduler": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "slb_id": {
              "computed": true,
              "description": "if field 'slb_id' is empty, EDAS will purchase a new slb for this config",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "specification": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "port_mappings": {
              "block": {
                "attributes": {
                  "cert_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "loadbalancer_protocol": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "service_port": {
                    "block": {
                      "attributes": {
                        "port": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "protocol": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "target_port": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
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

func AlicloudEdasK8SSlbAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEdasK8SSlbAttachment), &result)
	return &result
}
