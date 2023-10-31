package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSaeApplicationScalingRule = `{
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
      },
      "min_ready_instance_ratio": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_ready_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scaling_rule_enable": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "scaling_rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scaling_rule_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "scaling_rule_metric": {
        "block": {
          "attributes": {
            "max_replicas": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_replicas": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "metrics": {
              "block": {
                "attributes": {
                  "metric_target_average_utilization": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "metric_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "slb_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "slb_log_store": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "slb_project": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vport": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "scale_down_rules": {
              "block": {
                "attributes": {
                  "disabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "stabilization_window_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "step": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "set"
            },
            "scale_up_rules": {
              "block": {
                "attributes": {
                  "disabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "stabilization_window_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "step": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "scaling_rule_timer": {
        "block": {
          "attributes": {
            "begin_date": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "end_date": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "period": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "schedules": {
              "block": {
                "attributes": {
                  "at_time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_replicas": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_replicas": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_replicas": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSaeApplicationScalingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSaeApplicationScalingRule), &result)
	return &result
}
