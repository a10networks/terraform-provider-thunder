package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_CACHE_RESOURCE = `
resource "thunder_slb_template_cache" "template_cache1" {
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
func TestAccSlbTemplateCache_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_CACHE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "name", "testcache"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "accept_reload_req", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "default_policy_nocache", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "age", "4550"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "disable_insert_via", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "replacement_policy", "LFU"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "disable_insert_age", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "max_content_size", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "max_cache_size", "700"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "remove_cookies", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "verify_host", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "min_content_size", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_cache.template_cache1", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
