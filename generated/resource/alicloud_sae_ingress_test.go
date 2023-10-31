package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-alicloud-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAlicloudSaeIngressSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AlicloudSaeIngressSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
