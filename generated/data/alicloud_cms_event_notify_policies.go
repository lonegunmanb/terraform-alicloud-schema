package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsEventNotifyPolicies = `{
  "block": {
    "attributes": {
      "enable_details": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order_by": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order_desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "description": "string",
              "enabled": "bool",
              "id": "string",
              "name": "string",
              "notify_strategy": [
                "list",
                [
                  "object",
                  {
                    "custom_template_entries": [
                      "list",
                      [
                        "object",
                        {
                          "template_uuid": "string"
                        }
                      ]
                    ],
                    "description": "string",
                    "grouping_setting": [
                      "list",
                      [
                        "object",
                        {
                          "grouping_keys": [
                            "list",
                            "string"
                          ],
                          "period_min": "number",
                          "silence_sec": "number",
                          "times": "number"
                        }
                      ]
                    ],
                    "ignore_restored_notification": "bool",
                    "routes": [
                      "list",
                      [
                        "object",
                        {
                          "channels": [
                            "list",
                            [
                              "object",
                              {
                                "channel_type": "string",
                                "enabled_sub_channels": [
                                  "list",
                                  "string"
                                ],
                                "receivers": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "digital_employee_name": "string",
                          "effect_time_range": [
                            "list",
                            [
                              "object",
                              {
                                "day_in_week": [
                                  "list",
                                  "number"
                                ],
                                "end_time_in_minute": "number",
                                "start_time_in_minute": "number",
                                "time_zone": "string"
                              }
                            ]
                          ],
                          "enable_rca": "bool",
                          "filter_setting": [
                            "list",
                            [
                              "object",
                              {
                                "conditions": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "field": "string",
                                      "op": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "expression": "string",
                                "relation": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "response_plan": [
                "list",
                [
                  "object",
                  {
                    "auto_recover_seconds": "number",
                    "escalation_id": [
                      "list",
                      "string"
                    ],
                    "pushing_setting": [
                      "list",
                      [
                        "object",
                        {
                          "alert_action_ids": [
                            "list",
                            "string"
                          ],
                          "restore_action_ids": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "repeat_notify_setting": [
                      "list",
                      [
                        "object",
                        {
                          "end_incident_state": "string",
                          "repeat_interval": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "subscription": [
                "list",
                [
                  "object",
                  {
                    "filter_setting": [
                      "list",
                      [
                        "object",
                        {
                          "conditions": [
                            "list",
                            [
                              "object",
                              {
                                "field": "string",
                                "op": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "expression": "string",
                          "relation": "string"
                        }
                      ]
                    ],
                    "subscribe_legacy_event": "bool",
                    "workspace_filter_setting": [
                      "list",
                      [
                        "object",
                        {
                          "tag_selector": [
                            "list",
                            [
                              "object",
                              {
                                "conditions": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "field": "string",
                                      "op": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "expression": "string",
                                "relation": "string"
                              }
                            ]
                          ],
                          "workspace_uuids": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "update_time": "string",
              "user_id": "string",
              "uuid": "string",
              "version": "number",
              "workspace": "string"
            }
          ]
        ]
      },
      "workspace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsEventNotifyPoliciesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsEventNotifyPolicies), &result)
	return &result
}
