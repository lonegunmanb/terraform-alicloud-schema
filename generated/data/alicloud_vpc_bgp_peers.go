package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcBgpPeers = `{
  "block": {
    "attributes": {
      "bgp_group_id": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auth_key": "string",
              "bfd_multi_hop": "number",
              "bgp_group_id": "string",
              "bgp_peer_id": "string",
              "bgp_peer_name": "string",
              "bgp_status": "string",
              "description": "string",
              "enable_bfd": "bool",
              "hold": "string",
              "id": "string",
              "ip_version": "string",
              "is_fake": "bool",
              "keepalive": "string",
              "local_asn": "string",
              "peer_asn": "string",
              "peer_ip_address": "string",
              "route_limit": "string",
              "router_id": "string",
              "status": "string"
            }
          ]
        ]
      },
      "router_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVpcBgpPeersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcBgpPeers), &result)
	return &result
}
