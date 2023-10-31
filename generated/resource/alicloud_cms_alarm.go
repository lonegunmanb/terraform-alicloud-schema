package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsAlarm = `{
  "block": {
    "attributes": {
      "contact_groups": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "dimensions": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "effective_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "end_time": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metric": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metric_dimensions": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notify_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "operator": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "silence_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "start_time": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "statistics": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "threshold": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "triggered_count": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "webhook": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "escalations_critical": {
        "block": {
          "attributes": {
            "comparison_operator": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "statistics": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "times": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "escalations_info": {
        "block": {
          "attributes": {
            "comparison_operator": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "statistics": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "times": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "escalations_warn": {
        "block": {
          "attributes": {
            "comparison_operator": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "statistics": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "times": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "prometheus": {
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
            "level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prom_ql": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "times": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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

func AlicloudCmsAlarmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsAlarm), &result)
	return &result
}
