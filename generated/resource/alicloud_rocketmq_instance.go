package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRocketmqInstance = `{
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
      "auto_renew_period_unit": {
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "period_unit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remark": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "series_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sub_series_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "network_info": {
        "block": {
          "attributes": {
            "endpoints": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "endpoint_type": "string",
                    "endpoint_url": "string",
                    "ip_white_list": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          },
          "block_types": {
            "internet_info": {
              "block": {
                "attributes": {
                  "flow_out_bandwidth": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "flow_out_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "internet_spec": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ip_whitelist": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "vpc_info": {
              "block": {
                "attributes": {
                  "vpc_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "vswitch_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "product_info": {
        "block": {
          "attributes": {
            "auto_scaling": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "message_retention_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "msg_process_spec": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "send_receive_ratio": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "support_auto_scaling": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "software": {
        "block": {
          "attributes": {
            "maintain_time": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "software_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "upgrade_method": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AlicloudRocketmqInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRocketmqInstance), &result)
	return &result
}
