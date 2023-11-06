package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-alicloud-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAlicloudThreatDetectionClientFileProtectSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AlicloudThreatDetectionClientFileProtectSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
