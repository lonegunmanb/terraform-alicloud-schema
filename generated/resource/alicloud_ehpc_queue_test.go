package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-alicloud-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAlicloudEhpcQueueSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AlicloudEhpcQueueSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
