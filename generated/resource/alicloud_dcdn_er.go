package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDcdnEr = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "er_name": {
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
      "env_conf": {
        "block": {
          "block_types": {
            "preset_canary_anhui": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_beijing": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_chongqing": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_fujian": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_gansu": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_guangdong": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_guangxi": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_guizhou": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_hainan": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_hebei": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_heilongjiang": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_henan": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_hong_kong": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_hubei": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_hunan": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_jiangsu": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_jiangxi": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_jilin": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_liaoning": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_macau": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_neimenggu": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_ningxia": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_overseas": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_qinghai": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_shaanxi": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_shandong": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_shanghai": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_shanxi": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_sichuan": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_taiwan": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_tianjin": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_xinjiang": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_xizang": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_yunnan": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "preset_canary_zhejiang": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "production": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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
            "staging": {
              "block": {
                "attributes": {
                  "allowed_hosts": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "code_rev": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spec_name": {
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

func AlicloudDcdnErSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDcdnEr), &result)
	return &result
}
