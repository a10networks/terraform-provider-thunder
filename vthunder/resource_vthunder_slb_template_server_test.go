package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_SERVER_RESOURCE = `
resource "vthunder_slb_template_server" "testname" {
	name = "testserver"
	user_tag = "test_tag"
	stats_data_action = "stats-data-enable"
	slow_start = 1
	weight = 3
	bw_rate_limit = 2
	spoofing_cache = 1
	conn_limit = 1
	resume = 1
	max_dynamic_server = 3
	rate_interval = "second"
	min_ttl_ratio = 2
	bw_rate_limit_no_logging = 1
	dynamic_server_prefix = "DRS"
	conn_limit_no_logging = 1
	extended_stats = 1
	conn_rate_limit_no_logging = 1
	bw_rate_limit_duration = 1
	bw_rate_limit_resume = 1
	bw_rate_limit_acct = "to-server-only"
	log_selection_failure = 1
	conn_rate_limit = 5
	dns_query_interval = 2
}
`

//Acceptance test
func TestAccSlbTemplateServer_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_SERVER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "name", "testserver"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "stats_data_action", "stats-data-enable"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "slow_start", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "weight", "3"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "bw_rate_limit", "2"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "spoofing_cache", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "conn_limit", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "resume", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "max_dynamic_server", "3"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "rate_interval", "second"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "min_ttl_ratio", "2"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "bw_rate_limit_no_logging", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "dynamic_server_prefix", "DRS"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "conn_limit_no_logging", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "extended_stats", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "conn_rate_limit_no_logging", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "bw_rate_limit_duration", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "bw_rate_limit_resume", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "bw_rate_limit_acct", "to-server-only"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "log_selection_failure", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "conn_rate_limit", "5"),
					resource.TestCheckResourceAttr("vthunder_slb_template_server.testname", "dns_query_interval", "2"),
				),
			},
		},
	})
}
