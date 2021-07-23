package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_SERVER_RESOURCE = `
resource "thunder_slb_template_server" "server" {
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
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "name", "testserver"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "stats_data_action", "stats-data-enable"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "slow_start", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "weight", "3"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "bw_rate_limit", "2"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "spoofing_cache", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "conn_limit", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "resume", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "max_dynamic_server", "3"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "rate_interval", "second"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "min_ttl_ratio", "2"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "bw_rate_limit_no_logging", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "dynamic_server_prefix", "DRS"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "conn_limit_no_logging", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "extended_stats", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "conn_rate_limit_no_logging", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "bw_rate_limit_duration", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "bw_rate_limit_resume", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "bw_rate_limit_acct", "to-server-only"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "log_selection_failure", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "conn_rate_limit", "5"),
					resource.TestCheckResourceAttr("thunder_slb_template_server.server", "dns_query_interval", "2"),
				),
			},
		},
	})
}
