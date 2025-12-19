package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEhpcQueue = `{
  "block": {
    "attributes": {
      "cluster_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_scale_in": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_scale_out": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hostname_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hostname_suffix": {
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
      "initial_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "inter_connect": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "queue_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "compute_nodes": {
        "block": {
          "attributes": {
            "auto_renew": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "auto_renew_period": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enable_ht": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "image_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_charge_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "period": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "period_unit": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "spot_price_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "spot_strategy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "system_disk": {
              "block": {
                "attributes": {
                  "category": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "level": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "size": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
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

func AlicloudEhpcQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEhpcQueue), &result)
	return &result
}
