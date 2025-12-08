package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaWafRule = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "phase": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ruleset_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "waf_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "config": {
        "block": {
          "attributes": {
            "action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "managed_group_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "managed_list": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "notes": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sigchl": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "status": {
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
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "actions": {
              "block": {
                "block_types": {
                  "bypass": {
                    "block": {
                      "attributes": {
                        "custom_rules": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "number"
                          ]
                        },
                        "regular_rules": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "number"
                          ]
                        },
                        "regular_types": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "skip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "tags": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "response": {
                    "block": {
                      "attributes": {
                        "code": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "id": {
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
            "app_package": {
              "block": {
                "block_types": {
                  "package_signs": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sign": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
            },
            "app_sdk": {
              "block": {
                "attributes": {
                  "custom_sign_status": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "feature_abnormal": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "custom_sign": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "managed_rulesets": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "attack_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "number_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "number_total": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "protection_level": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "managed_rules": {
                    "block": {
                      "attributes": {
                        "action": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "status": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
            },
            "rate_limit": {
              "block": {
                "attributes": {
                  "interval": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "on_hit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "ttl": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "characteristics": {
                    "block": {
                      "attributes": {
                        "logic": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "criteria": {
                          "block": {
                            "attributes": {
                              "logic": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "match_type": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "criteria": {
                                "block": {
                                  "attributes": {
                                    "logic": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "match_type": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "criteria": {
                                      "block": {
                                        "attributes": {
                                          "match_type": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
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
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "threshold": {
                    "block": {
                      "attributes": {
                        "distinct_managed_rules": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "managed_rules_blocked": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "request": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "traffic": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "response_status": {
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "security_level": {
              "block": {
                "attributes": {
                  "value": {
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
            "timer": {
              "block": {
                "attributes": {
                  "scopes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "zone": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "periods": {
                    "block": {
                      "attributes": {
                        "end": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "start": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "weekly_periods": {
                    "block": {
                      "attributes": {
                        "days": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "daily_periods": {
                          "block": {
                            "attributes": {
                              "end": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "start": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
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
        "nesting_mode": "list"
      },
      "shared": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cross_site_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "actions": {
              "block": {
                "block_types": {
                  "response": {
                    "block": {
                      "attributes": {
                        "code": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "id": {
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
            "match": {
              "block": {
                "attributes": {
                  "logic": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "match_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "criteria": {
                    "block": {
                      "attributes": {
                        "logic": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "match_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "criteria": {
                          "block": {
                            "attributes": {
                              "logic": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "match_type": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "criteria": {
                                "block": {
                                  "attributes": {
                                    "match_type": {
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

func AlicloudEsaWafRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaWafRule), &result)
	return &result
}
