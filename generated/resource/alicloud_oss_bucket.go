package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOssBucket = `{
  "block": {
    "attributes": {
      "acl": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "extranet_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "force_destroy": {
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
      "intranet_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lifecycle_rule_allow_same_action_overlap": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "logging_isenable": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redundancy_type": {
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
      "storage_class": {
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
      }
    },
    "block_types": {
      "access_monitor": {
        "block": {
          "attributes": {
            "status": {
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
      },
      "cors_rule": {
        "block": {
          "attributes": {
            "allowed_headers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allowed_methods": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allowed_origins": {
              "description_kind": "plain",
              "required": true,
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
            "max_age_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 10,
        "nesting_mode": "list"
      },
      "lifecycle_rule": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prefix": {
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
            }
          },
          "block_types": {
            "abort_multipart_upload": {
              "block": {
                "attributes": {
                  "created_before_date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "expiration": {
              "block": {
                "attributes": {
                  "created_before_date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "expired_object_delete_marker": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "filter": {
              "block": {
                "attributes": {
                  "object_size_greater_than": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "object_size_less_than": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "not": {
                    "block": {
                      "attributes": {
                        "prefix": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "tag": {
                          "block": {
                            "attributes": {
                              "key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "noncurrent_version_expiration": {
              "block": {
                "attributes": {
                  "days": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "noncurrent_version_transition": {
              "block": {
                "attributes": {
                  "days": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "is_access_time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "return_to_std_when_visit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "storage_class": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "transitions": {
              "block": {
                "attributes": {
                  "created_before_date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "is_access_time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "return_to_std_when_visit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "storage_class": {
                    "description_kind": "plain",
                    "required": true,
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
        "max_items": 1000,
        "nesting_mode": "list"
      },
      "logging": {
        "block": {
          "attributes": {
            "target_bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_prefix": {
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
      "referer_config": {
        "block": {
          "attributes": {
            "allow_empty": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "referers": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "server_side_encryption_rule": {
        "block": {
          "attributes": {
            "kms_master_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sse_algorithm": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "transfer_acceleration": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "versioning": {
        "block": {
          "attributes": {
            "status": {
              "description_kind": "plain",
              "required": true,
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
            "error_document": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "index_document": {
              "description_kind": "plain",
              "required": true,
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
  "version": 0
}`

func AlicloudOssBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOssBucket), &result)
	return &result
}
