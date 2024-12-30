package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSchedulerxJob = `{
  "block": {
    "attributes": {
      "attempt_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "class_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execute_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fail_times": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "group_id": {
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
      "job_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "job_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "job_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_attempt": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_concurrency": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
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
      "success_notice_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "task_dispatch_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timezone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "x_attrs": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "job_monitor_info": {
        "block": {
          "block_types": {
            "contact_info": {
              "block": {
                "attributes": {
                  "ding": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_mail": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_phone": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "monitor_config": {
              "block": {
                "attributes": {
                  "fail_enable": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "miss_worker_enable": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "send_channel": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "timeout_enable": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "timeout_kill_enable": {
                    "computed": true,
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "map_task_xattrs": {
        "block": {
          "attributes": {
            "consumer_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dispatcher_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "page_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "queue_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "task_attempt_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "task_max_attempt": {
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
      "time_config": {
        "block": {
          "attributes": {
            "calendar": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_offset": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "time_expression": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "time_type": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
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

func AlicloudSchedulerxJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSchedulerxJob), &result)
	return &result
}
