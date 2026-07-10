package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsAlertRuleV2 = `{
  "block": {
    "attributes": {
      "alert_rule_v2_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "annotations": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "content_template": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "datasource_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "computed": true,
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
      "labels": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "notify_strategy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "observe_resource_global_scope": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "observe_resource_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "partition_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "severity_levels": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "workspace": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "action_integration_config": {
        "block": {
          "attributes": {
            "actions": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "enabled": {
              "computed": true,
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
      "arms_integration_config": {
        "block": {
          "attributes": {
            "enabled": {
              "computed": true,
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
      "condition_config": {
        "block": {
          "attributes": {
            "aggregate": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "composite_escalation": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "escalations": [
                      "list",
                      [
                        "object",
                        {
                          "comparison_operator": "string",
                          "metric_name": "string",
                          "period": "number",
                          "pre_condition": "string",
                          "statistics": "string",
                          "threshold": "string"
                        }
                      ]
                    ],
                    "relation": "string",
                    "severity": "string",
                    "times": "number"
                  }
                ]
              ]
            },
            "duration_secs": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "escalation_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "express_escalation": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "raw_expression": "string",
                    "severity": "string",
                    "times": "number"
                  }
                ]
              ]
            },
            "legacy_raw": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "legacy_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "no_data_policy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "operator": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prometheus": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "prom_ql": "string",
                    "severity": "string",
                    "times": "number"
                  }
                ]
              ]
            },
            "relation": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "severity": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "simple_escalation": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "escalations": [
                      "list",
                      [
                        "object",
                        {
                          "comparison_operator": "string",
                          "pre_condition": "string",
                          "severity": "string",
                          "statistics": "string",
                          "threshold": "string",
                          "times": "number"
                        }
                      ]
                    ],
                    "metric_name": "string",
                    "period": "number"
                  }
                ]
              ]
            },
            "threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "yoy_time_unit": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "yoy_time_value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "compare_list": {
              "block": {
                "attributes": {
                  "aggregate": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operator": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "threshold": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "yoy_time_unit": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "yoy_time_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "threshold_list": {
              "block": {
                "attributes": {
                  "severity": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "threshold": {
                    "computed": true,
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
      },
      "datasource_config": {
        "block": {
          "attributes": {
            "instance_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "legacy_raw": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "legacy_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "product_category": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "region_id": {
              "computed": true,
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
        "nesting_mode": "list"
      },
      "notify_config": {
        "block": {
          "attributes": {
            "active_days": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "active_end_time": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "active_start_time": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "notify_strategies": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "silence_time_secs": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "utc_offset": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "channels": {
              "block": {
                "attributes": {
                  "identifiers": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "type": {
                    "computed": true,
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "query_config": {
        "block": {
          "attributes": {
            "dimensions": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "map",
                  "string"
                ]
              ]
            },
            "enable_data_complete_check": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "entity_domain": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "entity_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expr": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "group_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "legacy_raw": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "legacy_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "metric": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_set": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "prom_ql": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "relation_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "service_id_list": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "entity_fields": {
              "block": {
                "attributes": {
                  "field": {
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
            "entity_filters": {
              "block": {
                "attributes": {
                  "field": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operator": {
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
            "filter_list": {
              "block": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
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
            "label_filters": {
              "block": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operator": {
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
            "measure_list": {
              "block": {
                "attributes": {
                  "group_by": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "measure_code": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "window_secs": {
                    "computed": true,
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
      },
      "schedule_config": {
        "block": {
          "attributes": {
            "interval_secs": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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

func AlicloudCmsAlertRuleV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsAlertRuleV2), &result)
	return &result
}
