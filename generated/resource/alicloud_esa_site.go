package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaSite = `{
  "block": {
    "attributes": {
      "access_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "add_client_geolocation_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "add_real_client_ip_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ai_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ai_template": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "automatic_frequency_control_action_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "automatic_frequency_control_enable": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "automatic_frequency_control_level": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_architecture_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_reserve_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_reserve_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "case_insensitive": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "coverage": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cross_border_optimization": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "development_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flatten_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_mode": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipv6_enable": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_region": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "paused": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "performance_data_collection_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "real_client_ip_header_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "seo_bypass": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_name_exclusive": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tag_name": {
        "description_kind": "plain",
        "optional": true,
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
      "version_management": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "site_waf_settings": {
        "block": {
          "block_types": {
            "add_bot_protection_headers": {
              "block": {
                "attributes": {
                  "enable": {
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
            "add_security_headers": {
              "block": {
                "attributes": {
                  "enable": {
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
            "bandwidth_abuse_protection": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "bot_management": {
              "block": {
                "block_types": {
                  "definite_bots": {
                    "block": {
                      "attributes": {
                        "action": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "effect_on_static": {
                    "block": {
                      "attributes": {
                        "enable": {
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
                  "js_detection": {
                    "block": {
                      "attributes": {
                        "enable": {
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
                  "likely_bots": {
                    "block": {
                      "attributes": {
                        "action": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "verified_bots": {
                    "block": {
                      "attributes": {
                        "action": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description_kind": "plain",
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
            "client_ip_identifier": {
              "block": {
                "attributes": {
                  "headers": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "mode": {
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
            "security_level": {
              "block": {
                "attributes": {
                  "value": {
                    "computed": true,
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

func AlicloudEsaSiteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaSite), &result)
	return &result
}
