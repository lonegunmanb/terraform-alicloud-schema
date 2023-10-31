package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOtsSearchIndex = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "current_sync_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "index_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "index_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sync_phase": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "table_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "time_to_live": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "schema": {
        "block": {
          "block_types": {
            "field_schema": {
              "block": {
                "attributes": {
                  "analyzer": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_sort_and_agg": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "field_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "field_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "index": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "is_array": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "store": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "index_setting": {
              "block": {
                "attributes": {
                  "routing_fields": {
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
              "nesting_mode": "set"
            },
            "index_sort": {
              "block": {
                "block_types": {
                  "sorter": {
                    "block": {
                      "attributes": {
                        "field_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mode": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "order": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sorter_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOtsSearchIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOtsSearchIndex), &result)
	return &result
}
