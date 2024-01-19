package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudArmsSyntheticTask = `{
  "block": {
    "attributes": {
      "frequency": {
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
      "monitor_category": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "synthetic_task_name": {
        "description_kind": "plain",
        "required": true,
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
      "task_type": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "block_types": {
      "available_assertions": {
        "block": {
          "attributes": {
            "expect": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "operator": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target": {
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
        "nesting_mode": "list"
      },
      "common_setting": {
        "block": {
          "attributes": {
            "ip_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "is_open_trace": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "monitor_samples": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "trace_client_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "xtrace_region": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "custom_host": {
              "block": {
                "attributes": {
                  "select_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "hosts": {
                    "block": {
                      "attributes": {
                        "domain": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "ip_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "ips": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description_kind": "plain"
                    },
                    "min_items": 1,
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
      "custom_period": {
        "block": {
          "attributes": {
            "end_hour": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "start_hour": {
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
      "monitor_conf": {
        "block": {
          "block_types": {
            "api_http": {
              "block": {
                "attributes": {
                  "connect_timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_headers": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "target_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "request_body": {
                    "block": {
                      "attributes": {
                        "content": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
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
            "file_download": {
              "block": {
                "attributes": {
                  "connection_timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "custom_header_content": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "download_kernel": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ignore_certificate_auth_error": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ignore_certificate_canceled_error": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ignore_certificate_out_of_date_error": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ignore_certificate_status_error": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ignore_certificate_untrustworthy_error": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ignore_certificate_using_error": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ignore_invalid_host_error": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "monitor_timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "quick_protocol": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "redirection": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "transmission_size": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "validate_keywords": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "verify_way": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "white_list": {
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
            "net_dns": {
              "block": {
                "attributes": {
                  "dns_server_ip_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ns_server": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "query_method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timeout": {
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
            "net_icmp": {
              "block": {
                "attributes": {
                  "interval": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "package_num": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "package_size": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "split_package": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "target_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "tracert_enable": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tracert_num_max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "tracert_timeout": {
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
            "net_tcp": {
              "block": {
                "attributes": {
                  "connect_times": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "interval": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "tracert_enable": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tracert_num_max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "tracert_timeout": {
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
            "stream": {
              "block": {
                "attributes": {
                  "custom_header_content": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "player_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "stream_address_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "stream_monitor_timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "stream_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_url": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "white_list": {
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
            "website": {
              "block": {
                "attributes": {
                  "automatic_scrolling": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "custom_header": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "custom_header_content": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "disable_cache": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "disable_compression": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "dns_hijack_whitelist": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "element_blacklist": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "filter_invalid_ip": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "flow_hijack_jump_times": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "flow_hijack_logo": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ignore_certificate_error": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "monitor_timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "page_tamper": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "redirection": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "slow_element_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "verify_string_blacklist": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "verify_string_whitelist": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "wait_completion_time": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "monitors": {
        "block": {
          "attributes": {
            "city_code": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "client_type": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "operator_code": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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

func AlicloudArmsSyntheticTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudArmsSyntheticTask), &result)
	return &result
}
