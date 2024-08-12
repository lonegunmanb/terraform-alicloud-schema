package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsMetricRuleTemplate = `{
  "block": {
    "attributes": {
      "apply_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_end_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_id": {
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
      "metric_rule_template_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notify_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rest_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "silence_time": {
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
      "alert_templates": {
        "block": {
          "attributes": {
            "category": {
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
            "rule_name": {
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
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "set"
                  },
                  "info": {
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
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "set"
                  },
                  "warn": {
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
                          "type": "string"
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsMetricRuleTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsMetricRuleTemplate), &result)
	return &result
}
