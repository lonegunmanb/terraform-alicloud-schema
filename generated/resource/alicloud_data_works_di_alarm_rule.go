package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDataWorksDiAlarmRule = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "di_alarm_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "di_alarm_rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "di_job_id": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metric_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "notification_settings": {
        "block": {
          "attributes": {
            "inhibition_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "notification_channels": {
              "block": {
                "attributes": {
                  "channels": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "severity": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "notification_receivers": {
              "block": {
                "attributes": {
                  "receiver_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "receiver_values": {
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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
      },
      "trigger_conditions": {
        "block": {
          "attributes": {
            "ddl_report_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "severity": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDataWorksDiAlarmRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDataWorksDiAlarmRule), &result)
	return &result
}
