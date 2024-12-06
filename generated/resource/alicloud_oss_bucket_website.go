package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOssBucketWebsite = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "error_document": {
        "block": {
          "attributes": {
            "http_status": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "key": {
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
      "index_document": {
        "block": {
          "attributes": {
            "suffix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "support_sub_dir": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
      },
      "routing_rules": {
        "block": {
          "block_types": {
            "routing_rule": {
              "block": {
                "attributes": {
                  "rule_number": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "condition": {
                    "block": {
                      "attributes": {
                        "http_error_code_returned_equals": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key_prefix_equals": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key_suffix_equals": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "include_headers": {
                          "block": {
                            "attributes": {
                              "ends_with": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "equals": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "starts_with": {
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
                  "lua_config": {
                    "block": {
                      "attributes": {
                        "script": {
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
                  "redirect": {
                    "block": {
                      "attributes": {
                        "enable_replace_prefix": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "host_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "http_redirect_code": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mirror_allow_get_image_info": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_allow_head_object": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_allow_video_snapshot": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_async_status": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "mirror_check_md5": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_dst_region": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mirror_dst_slave_vpc_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mirror_dst_vpc_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mirror_follow_redirect": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_is_express_tunnel": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_pass_original_slashes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_pass_query_string": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_proxy_pass": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_role": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mirror_save_oss_meta": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_sni": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_switch_all_errors": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_tunnel_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mirror_url": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mirror_url_probe": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mirror_url_slave": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mirror_user_last_modified": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mirror_using_role": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "pass_query_string": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "protocol": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "redirect_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "replace_key_prefix_with": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "replace_key_with": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "transparent_mirror_response_codes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "mirror_auth": {
                          "block": {
                            "attributes": {
                              "access_key_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "access_key_secret": {
                                "description_kind": "plain",
                                "optional": true,
                                "sensitive": true,
                                "type": "string"
                              },
                              "auth_type": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "region": {
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
                        "mirror_headers": {
                          "block": {
                            "attributes": {
                              "pass": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "pass_all": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "remove": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "set": {
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
                        "mirror_multi_alternates": {
                          "block": {
                            "block_types": {
                              "mirror_multi_alternate": {
                                "block": {
                                  "attributes": {
                                    "mirror_multi_alternate_dst_region": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "mirror_multi_alternate_number": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "mirror_multi_alternate_url": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "mirror_multi_alternate_vpc_id": {
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
                        "mirror_return_headers": {
                          "block": {
                            "block_types": {
                              "return_header": {
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
                        "mirror_taggings": {
                          "block": {
                            "block_types": {
                              "taggings": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOssBucketWebsiteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOssBucketWebsite), &result)
	return &result
}
