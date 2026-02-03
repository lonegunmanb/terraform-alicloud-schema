package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcv3CustomDomain = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "api_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain_name": {
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
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subdomain_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "auth_config": {
        "block": {
          "attributes": {
            "auth_info": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "auth_type": {
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
      "cert_config": {
        "block": {
          "attributes": {
            "cert_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "certificate": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cors_config": {
        "block": {
          "attributes": {
            "allow_credentials": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_headers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allow_methods": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allow_origins": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "expose_headers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "max_age": {
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
      "route_config": {
        "block": {
          "block_types": {
            "routes": {
              "block": {
                "attributes": {
                  "function_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "methods": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "qualifier": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "rewrite_config": {
                    "block": {
                      "block_types": {
                        "equal_rules": {
                          "block": {
                            "attributes": {
                              "match": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "replacement": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "regex_rules": {
                          "block": {
                            "attributes": {
                              "match": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "replacement": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "wildcard_rules": {
                          "block": {
                            "attributes": {
                              "match": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "replacement": {
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
      },
      "tls_config": {
        "block": {
          "attributes": {
            "cipher_suites": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "max_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_version": {
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
      "waf_config": {
        "block": {
          "attributes": {
            "enable_waf": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
  "version": 0
}`

func AlicloudFcv3CustomDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcv3CustomDomain), &result)
	return &result
}
