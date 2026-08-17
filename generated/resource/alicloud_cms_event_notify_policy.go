package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsEventNotifyPolicy = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "uuid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "workspace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "notify_strategy": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ignore_restored_notification": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "custom_template_entries": {
              "block": {
                "attributes": {
                  "template_uuid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "grouping_setting": {
              "block": {
                "attributes": {
                  "grouping_keys": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "period_min": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "silence_sec": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
            "routes": {
              "block": {
                "attributes": {
                  "digital_employee_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_rca": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "channels": {
                    "block": {
                      "attributes": {
                        "channel_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "enabled_sub_channels": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "receivers": {
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
                  },
                  "effect_time_range": {
                    "block": {
                      "attributes": {
                        "day_in_week": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        },
                        "end_time_in_minute": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "start_time_in_minute": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "time_zone": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "filter_setting": {
                    "block": {
                      "attributes": {
                        "expression": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "relation": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "conditions": {
                          "block": {
                            "attributes": {
                              "field": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "op": {
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
      "response_plan": {
        "block": {
          "attributes": {
            "auto_recover_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "escalation_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "pushing_setting": {
              "block": {
                "attributes": {
                  "alert_action_ids": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "restore_action_ids": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "repeat_notify_setting": {
              "block": {
                "attributes": {
                  "end_incident_state": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "repeat_interval": {
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
      "subscription": {
        "block": {
          "attributes": {
            "subscribe_legacy_event": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "filter_setting": {
              "block": {
                "attributes": {
                  "expression": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "relation": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "conditions": {
                    "block": {
                      "attributes": {
                        "field": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "op": {
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
            "workspace_filter_setting": {
              "block": {
                "attributes": {
                  "workspace_uuids": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "tag_selector": {
                    "block": {
                      "attributes": {
                        "expression": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "relation": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "conditions": {
                          "block": {
                            "attributes": {
                              "field": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "op": {
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

func AlicloudCmsEventNotifyPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsEventNotifyPolicy), &result)
	return &result
}
