package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_RESOURCE_USAGE_RESOURCE = `
resource "thunder_slb_resource_usage" "resource_usage1" {
	real_server_count = 128
	stream_template_count = 128
	proxy_template_count = 128
	http_template_count = 128
	persist_srcip_template_count = 128
	server_ssl_template_count = 128
	service_group_count = 128
	persist_cookie_template_count = 128
	virtual_server_count = 64
	fast_udp_template_count = 128
	fast_tcp_template_count = 128
	virtual_port_count = 128
	conn_reuse_template_count = 128
	real_port_count = 256
	client_ssl_template_count = 128
	nat_pool_addr_count = 10
	health_monitor_count = 1023
	pbslb_subnet_count = 65536 
}
`

//Acceptance test
func TestAccSlbResourceUsage_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_RESOURCE_USAGE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "real_server_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "stream_template_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "proxy_template_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "http_template_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "persist_srcip_template_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "server_ssl_template_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "service_group_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "persist_cookie_template_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "virtual_server_count", "64"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "fast_udp_template_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "fast_tcp_template_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "virtual_port_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "conn_reuse_template_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "real_port_count", "256"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "client_ssl_template_count", "128"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "nat_pool_addr_count", "10"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "health_monitor_count", "1023"),
					resource.TestCheckResourceAttr("thunder_slb_resource_usage.resource_usage1", "pbslb_subnet_count", "65536"),
				),
			},
		},
	})
}
