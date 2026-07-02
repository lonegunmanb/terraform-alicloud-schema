package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOssBucketInventory = `{
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
      },
      "included_object_versions": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "inventory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "destination": {
        "block": {
          "block_types": {
            "oss_bucket_destination": {
              "block": {
                "attributes": {
                  "account_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bucket": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "format": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "encryption": {
                    "block": {
                      "attributes": {
                        "sseoss": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "ssekms": {
                          "block": {
                            "attributes": {
                              "key_id": {
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
      "filter": {
        "block": {
          "attributes": {
            "last_modify_begin_time_stamp": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "last_modify_end_time_stamp": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "lower_size_bound": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_class": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "upper_size_bound": {
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
      "incremental_inventory": {
        "block": {
          "attributes": {
            "is_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "optional_fields": {
              "block": {
                "attributes": {
                  "field": {
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
            "schedule": {
              "block": {
                "attributes": {
                  "frequency": {
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
      "optional_fields": {
        "block": {
          "attributes": {
            "field": {
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
      "schedule": {
        "block": {
          "attributes": {
            "frequency": {
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

func AlicloudOssBucketInventorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOssBucketInventory), &result)
	return &result
}
