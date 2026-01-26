package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEventBridgeEventSourceV2 = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_bus_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_source_name": {
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
      "linked_external_source": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "source_http_event_parameters": {
        "block": {
          "attributes": {
            "ip": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "method": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "public_web_hook_url": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "referer": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "security_config": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_web_hook_url": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_kafka_parameters": {
        "block": {
          "attributes": {
            "consumer_group": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "offset_reset": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vswitch_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_mns_parameters": {
        "block": {
          "attributes": {
            "is_base64_decode": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "queue_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_oss_event_parameters": {
        "block": {
          "attributes": {
            "event_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "match_rules": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "list",
                  [
                    "object",
                    {
                      "match_state": "string",
                      "name": "string",
                      "prefix": "string",
                      "suffix": "string"
                    }
                  ]
                ]
              ]
            },
            "sts_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_rabbit_mq_parameters": {
        "block": {
          "attributes": {
            "instance_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "queue_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "virtual_host_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_rocketmq_parameters": {
        "block": {
          "attributes": {
            "auth_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_network": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_password": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_security_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_username": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_vpc_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_vswitch_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "offset": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timestamp": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "topic": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_scheduled_event_parameters": {
        "block": {
          "attributes": {
            "schedule": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "time_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_data": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_sls_parameters": {
        "block": {
          "attributes": {
            "consume_position": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "log_store": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_name": {
              "description_kind": "plain",
              "optional": true,
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

func AlicloudEventBridgeEventSourceV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEventBridgeEventSourceV2), &result)
	return &result
}
