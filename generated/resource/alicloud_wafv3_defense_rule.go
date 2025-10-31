package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudWafv3DefenseRule = `{
  "block": {
    "attributes": {
      "defense_origin": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "defense_scene": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "defense_type": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "rule_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "template_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "config": {
        "block": {
          "attributes": {
            "abroad_regions": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bypass_regular_rules": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "bypass_regular_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "bypass_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "cc_effect": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cc_status": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "cn_regions": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gray_status": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "remote_addr": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "rule_action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "throttle_threhold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "throttle_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ua": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "account_identifiers": {
              "block": {
                "attributes": {
                  "decode_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "position": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "sub_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "conditions": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "op_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sub_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "gray_config": {
              "block": {
                "attributes": {
                  "gray_rate": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gray_sub_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gray_target": {
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
            "rate_limit": {
              "block": {
                "attributes": {
                  "interval": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "sub_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ttl": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "status": {
                    "block": {
                      "attributes": {
                        "code": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "count": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "ratio": {
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
            "time_config": {
              "block": {
                "attributes": {
                  "time_scope": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "time_zone": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "time_periods": {
                    "block": {
                      "attributes": {
                        "end": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "start": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "week_time_periods": {
                    "block": {
                      "attributes": {
                        "day": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "day_periods": {
                          "block": {
                            "attributes": {
                              "end": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "start": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
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
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func AlicloudWafv3DefenseRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudWafv3DefenseRule), &result)
	return &result
}
