package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudMonitorServiceMetricAlarmRule = `{
  "block": {
    "attributes": {
      "contact_groups": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dimensions": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "effective_interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "email_subject": {
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
      "interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metric_alarm_rule_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metric_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "no_data_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "no_effective_interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resources": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "send_ok": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "silence_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "webhook": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "composite_expression": {
        "block": {
          "attributes": {
            "expression_list_join": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expression_raw": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "level": {
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
          "block_types": {
            "expression_list": {
              "block": {
                "attributes": {
                  "comparison_operator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "metric_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "period": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      "escalations": {
        "block": {
          "block_types": {
            "critical": {
              "block": {
                "attributes": {
                  "comparison_operator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pre_condition": {
                    "computed": true,
                    "description_kind": "plain",
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
            "info": {
              "block": {
                "attributes": {
                  "comparison_operator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pre_condition": {
                    "computed": true,
                    "description_kind": "plain",
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
            "warn": {
              "block": {
                "attributes": {
                  "comparison_operator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pre_condition": {
                    "computed": true,
                    "description_kind": "plain",
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "labels": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "prometheus": {
        "block": {
          "attributes": {
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
          "block_types": {
            "annotations": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
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

func AlicloudCloudMonitorServiceMetricAlarmRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudMonitorServiceMetricAlarmRule), &result)
	return &result
}
