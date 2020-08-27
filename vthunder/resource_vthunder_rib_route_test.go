package vthunder

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"reflect"
	"testing"
)

var DEST_ADDR = "0.0.0.0"

var MASK = "/0"

var TEST_ROUTE_RESOURCE = `
resource "vthunder_rib_route" "rib" {
ip_dest_addr="` + DEST_ADDR + `"
ip_mask="` + MASK + `"
ip_nexthop_ipv4 {
ip_next_hop="10.0.2.9"
distance_nexthop_ip=1
} 
}
`

//Acceptance test
func TestAccVthunderRibRoute_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckRibRouteDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckRibRouteExists(DEST_ADDR+"+/"+MASK, true),
					resource.TestCheckResourceAttr("vthunder_rib_route.rib", "ip_dest_addr", DEST_ADDR),
					resource.TestCheckResourceAttr("vthunder_rib_route.rib", "ip_mask", MASK),
					resource.TestCheckResourceAttr("vthunder_rib_route.rib", "ip_nexthop_ipv4.0.ip_next_hop", "10.0.2.9"),
					resource.TestCheckResourceAttr("vthunder_rib_route.rib", "ip_nexthop_ipv4.0.distance_nexthop_ip", "1"),
				),
			},
		},
	})
}

func TestAccVthunderRibRoute_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckRibRouteDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckRibRouteExists(NAME, true),
				),
				ResourceName:      NAME,
				ImportState:       false,
				ImportStateVerify: true,
			},
		},
	})
}

//Unit test for utility method to create Route structure
func TestDataToRibRoute(t *testing.T) {

	ip_dest_addr := "0.0.0.0"
	ip_mask := "/0"
	ip_next_hop := "10.0.2.9"
	distance_nexthop_ip := 1

	resourceSchema := map[string]*schema.Schema{
		"ip_dest_addr": {
			Type: schema.TypeString,
		},
		"ip_mask": {
			Type: schema.TypeString,
		},
		"ip_nexthop_ipv4": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"distance_nexthop_ip": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"description_nexthop_ip": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"ip_next_hop": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
				},
			},
		},
	}
	resourceDataMap := map[string]interface{}{
		"ip_dest_addr": ip_dest_addr,
		"ip_mask":      ip_mask,
		"ip_nexthop_ipv4": map[string]interface{}{
			"distance_nexthop_ip": distance_nexthop_ip,
			"ip_next_hop":         ip_next_hop,
		},
	}
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)

	var rib go_vthunder.Rib
	var rInst go_vthunder.RibInst

	rInst.IPDestAddr = "0.0.0.0"
	rInst.IPMask = "/0"
	rInst.Instance = "0.0.0.0+//0"

	rInst.IPNextHop = make([]go_vthunder.IPNexthopIpv4, 0, 1)
	/*var hop IPNexthopIpv4
	hop.DistanceNexthopIP = distance_nexthop_ip
	hop.IPNextHop = ip_next_hop
	rInst.IPNextHop = append(rInst.IPNextHop, hop)*/

	rib.UUID = rInst

	cases := []struct {
		input  *schema.ResourceData
		output go_vthunder.Rib
	}{
		{
			input:  resourceLocalData,
			output: rib,
		},
	}

	for _, tc := range cases {
		output := dataToRibRoute(resourceLocalData)
		if !reflect.DeepEqual(output, tc.output) {
			t.Fatalf("Got:\n\n%#v\n\nExpected:\n\n%#v", output, tc.output)
		}
	}

}

func testCheckRibRouteExists(name string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		client := testAccProvider.Meta().(vThunder)
		route, err := go_vthunder.GetRibRoute(client.Token, client.Host, name)

		if err != nil {
			return err
		}
		if exists && route == nil {
			return fmt.Errorf(" route %s was not created.", name)
		}
		if !exists && route != nil {
			return fmt.Errorf(" route %s still exists.", name)
		}
		return nil
	}
}

func testCheckRibRouteDestroyed(s *terraform.State) error {
	client := testAccProvider.Meta().(vThunder)
	for _, rs := range s.RootModule().Resources {

		name := rs.Primary.ID
		route, err := go_vthunder.GetRibRoute(client.Token, client.Host, name)
		if err != nil {
			return err
		}
		if route == nil {
			return fmt.Errorf(" route %s not destroyed.", name)
		}
	}
	return nil
}
