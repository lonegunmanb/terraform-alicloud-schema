package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsAlertRulesV2 = `{
  "block": {
    "attributes": {
      "filter_datasource_type_eq": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_display_name_contains": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_display_name_not_contains": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_enabled_eq": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "filter_labels_all_of_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_labels_all_of_value": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_labels_any_of_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_labels_any_of_value": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_notify_strategy_id_eq": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_observe_resource_global_scope_eq": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "filter_observe_resource_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_observe_resource_list_contains": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_observe_resource_type_eq": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_partition_key_eq": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_severity_levels_contains": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_status_eq": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_uuid_eq": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_uuid_in": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action_integration_config": [
                "list",
                [
                  "object",
                  {
                    "actions": [
                      "list",
                      "string"
                    ],
                    "enabled": "bool"
                  }
                ]
              ],
              "alert_rule_v2_id": "string",
              "annotations": [
                "map",
                "string"
              ],
              "arms_integration_config": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "condition_config": [
                "list",
                [
                  "object",
                  {
                    "aggregate": "string",
                    "compare_list": [
                      "list",
                      [
                        "object",
                        {
                          "aggregate": "string",
                          "operator": "string",
                          "threshold": "number",
                          "yoy_time_unit": "string",
                          "yoy_time_value": "number"
                        }
                      ]
                    ],
                    "composite_escalation": [
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
                    ],
                    "duration_secs": "number",
                    "escalation_type": "string",
                    "express_escalation": [
                      "list",
                      [
                        "object",
                        {
                          "raw_expression": "string",
                          "severity": "string",
                          "times": "number"
                        }
                      ]
                    ],
                    "legacy_raw": "string",
                    "legacy_type": "string",
                    "no_data_policy": "string",
                    "operator": "string",
                    "prometheus": [
                      "list",
                      [
                        "object",
                        {
                          "prom_ql": "string",
                          "severity": "string",
                          "times": "number"
                        }
                      ]
                    ],
                    "relation": "string",
                    "severity": "string",
                    "simple_escalation": [
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
                    ],
                    "threshold": "number",
                    "threshold_list": [
                      "list",
                      [
                        "object",
                        {
                          "severity": "string",
                          "threshold": "number"
                        }
                      ]
                    ],
                    "type": "string",
                    "yoy_time_unit": "string",
                    "yoy_time_value": "number"
                  }
                ]
              ],
              "content_template": "string",
              "created_at": "string",
              "datasource_config": [
                "list",
                [
                  "object",
                  {
                    "instance_id": "string",
                    "legacy_raw": "string",
                    "legacy_type": "string",
                    "product_category": "string",
                    "region_id": "string",
                    "type": "string"
                  }
                ]
              ],
              "datasource_type": "string",
              "display_name": "string",
              "enabled": "bool",
              "id": "string",
              "labels": [
                "map",
                "string"
              ],
              "notify_config": [
                "list",
                [
                  "object",
                  {
                    "active_days": [
                      "list",
                      "number"
                    ],
                    "active_end_time": "string",
                    "active_start_time": "string",
                    "channels": [
                      "list",
                      [
                        "object",
                        {
                          "identifiers": [
                            "list",
                            "string"
                          ],
                          "type": "string"
                        }
                      ]
                    ],
                    "notify_strategies": [
                      "list",
                      "string"
                    ],
                    "silence_time_secs": "number",
                    "type": "string",
                    "utc_offset": "string"
                  }
                ]
              ],
              "notify_strategy_id": "string",
              "observe_resource_global_scope": "bool",
              "observe_resource_type": "string",
              "partition_key": "string",
              "query_config": [
                "list",
                [
                  "object",
                  {
                    "dimensions": [
                      "list",
                      [
                        "map",
                        "string"
                      ]
                    ],
                    "enable_data_complete_check": "bool",
                    "entity_domain": "string",
                    "entity_fields": [
                      "list",
                      [
                        "object",
                        {
                          "field": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "entity_filters": [
                      "list",
                      [
                        "object",
                        {
                          "field": "string",
                          "operator": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "entity_type": "string",
                    "expr": "string",
                    "filter_list": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "type": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "group_id": "string",
                    "label_filters": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "operator": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "legacy_raw": "string",
                    "legacy_type": "string",
                    "measure_list": [
                      "list",
                      [
                        "object",
                        {
                          "group_by": [
                            "list",
                            "string"
                          ],
                          "measure_code": "string",
                          "window_secs": "number"
                        }
                      ]
                    ],
                    "metric": "string",
                    "metric_set": "string",
                    "namespace": "string",
                    "prom_ql": "string",
                    "relation_type": "string",
                    "service_id_list": [
                      "list",
                      "string"
                    ],
                    "type": "string"
                  }
                ]
              ],
              "schedule_config": [
                "list",
                [
                  "object",
                  {
                    "interval_secs": "number",
                    "type": "string"
                  }
                ]
              ],
              "severity_levels": "string",
              "status": "string",
              "updated_at": "string",
              "workspace": "string"
            }
          ]
        ]
      },
      "workspace": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsAlertRulesV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsAlertRulesV2), &result)
	return &result
}
