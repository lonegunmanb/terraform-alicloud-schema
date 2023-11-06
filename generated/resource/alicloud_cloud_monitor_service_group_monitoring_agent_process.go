package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudMonitorServiceGroupMonitoringAgentProcess = `{
  "block": {
    "attributes": {
      "group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_monitoring_agent_process_id": {
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
      "match_express_filter_relation": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "process_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "alert_config": {
        "block": {
          "attributes": {
            "comparison_operator": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "effective_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "escalations_level": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "silence_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "statistics": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "threshold": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "times": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "webhook": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "target_list": {
              "block": {
                "attributes": {
                  "arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "json_params": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "level": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_list_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "match_express": {
        "block": {
          "attributes": {
            "function": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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

func AlicloudCloudMonitorServiceGroupMonitoringAgentProcessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudMonitorServiceGroupMonitoringAgentProcess), &result)
	return &result
}
