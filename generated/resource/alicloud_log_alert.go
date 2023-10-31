package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudLogAlert = `{
  "block": {
    "attributes": {
      "alert_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alert_displayname": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "alert_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "auto_annotation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "condition": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dashboard": {
        "deprecated": true,
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
      "mute_until": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "no_data_fire": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "no_data_severity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "notify_threshold": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_interval": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule_type": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "send_resolved": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "threshold": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "throttling": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "annotations": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "group_configuration": {
        "block": {
          "attributes": {
            "fields": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "join_configurations": {
        "block": {
          "attributes": {
            "condition": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "labels": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "notification_list": {
        "block": {
          "attributes": {
            "content": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "email_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "mobile_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "service_uri": {
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
          "deprecated": true,
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "policy_configuration": {
        "block": {
          "attributes": {
            "action_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "alert_policy_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "repeat_interval": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "query_list": {
        "block": {
          "attributes": {
            "chart_title": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dashboard_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "end": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "logstore": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "power_sql_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "region": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "store": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "store_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "time_span_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "schedule": {
        "block": {
          "attributes": {
            "cron_expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "day_of_week": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "delay": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "hour": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "run_immediately": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "time_zone": {
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
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "severity_configurations": {
        "block": {
          "attributes": {
            "eval_condition": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            },
            "severity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "template_configuration": {
        "block": {
          "attributes": {
            "annotations": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "lang": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tokens": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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
  "version": 0
}`

func AlicloudLogAlertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudLogAlert), &result)
	return &result
}
