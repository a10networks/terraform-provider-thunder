//package vthunder
//
//import (
//	"fmt"
//	"github.com/go_vthunder/vthunder"
//	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
//	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
//	"github.com/hashicorp/terraform-plugin-sdk/terraform"
//	"reflect"
//	"testing"
//)
//
//
//var TEST_VTEP_RESOURCE = `
//resource "vthunder_overlay_tunnel_vtep" "vtep" {
//id=1
//user_tag="Tag"
//encap="Enable"
//}
//`
//
////Acceptance test
//func TestAccVthunderOverlayTunnelVtep_create(t *testing.T) {
//	resource.Test(t, resource.TestCase{
//		PreCheck: func() {
//			testAcctPreCheck(t)
//		},
//		Providers:    testAccProviders,
//		Steps: []resource.TestStep{
//			{
//				Config: TEST_SR_RESOURCE,
//				Check: resource.ComposeTestCheckFunc(
//					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "id", "1"),
//					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "user_tag", "Tag"),
//					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "encap", "Enable"),
//				),
//			},
//		},
//	})
//}
