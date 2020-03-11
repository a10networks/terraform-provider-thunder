package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_CACHE_RESOURCE = `
resource "vthunder_slb_template_cache" "testname" {
	name = "testcache"
	user_tag = "test_tag"
	accept_reload_req = 0
	default_policy_nocache = 0
	age = 4550
	disable_insert_via = 0
	replacement_policy = "LFU"
	disable_insert_age = 0
	max_content_size = 0
	max_cache_size = 700
	remove_cookies = 0
	verify_host = 0
	min_content_size = 0
	sampling_enable {
        counters1 = "all"
      }
}
`

//Acceptance test
func TestSlbTemplateCache_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_CACHE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "name", "testcache"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "accept_reload_req", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "default_policy_nocache", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "age", "4550"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "disable_insert_via", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "replacement_policy", "LFU"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "disable_insert_age", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "max_content_size", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "max_cache_size", "700"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "remove_cookies", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "verify_host", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "min_content_size", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_cache.testname", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
