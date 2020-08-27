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

var NAME_VRID = "1"

var TEST_VRRP_VRID_RESOURCE = `
	resource "vthunder_vrrp_vrid" "vrrp_vrid" {
	  vrid_val = 1
      floating_ip {
        ip_address_cfg {
        ip_address = "1.1.1.10"
       }
      ip_address_cfg {
      ip_address = "1.1.1.11"
     }
    }
      blade_parameters {
      priority = 250
   }
}
`

//Acceptance test
func TestAccVthunderVrrpVrid_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_VRRP_VRID_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_vrrp_vrid.vrrp_vrid", "vrid_val", "1"),
					resource.TestCheckResourceAttr("vthunder_vrrp_vrid.vrrp_vrid", "floating_ip.0.ip_address_cfg.0.ip_address", "1.1.1.10"),
					resource.TestCheckResourceAttr("vthunder_vrrp_vrid.vrrp_vrid", "floating_ip.0.ip_address_cfg.1.ip_address", "1.1.1.11"),
					resource.TestCheckResourceAttr("vthunder_vrrp_vrid.vrrp_vrid", "blade_parameters.0.priority", "250"),
				),
			},
		},
	})
}

func TestAccVthunderVrrpVrid_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckVrrpVridDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_VRRP_VRID_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckVrrpVridExists(NAME_VRID, true),
				),
				ResourceName:      NAME_VRID,
				ImportState:       false,
				ImportStateVerify: true,
			},
		},
	})
}

func TestDataToVrrpVrid(t *testing.T) {

	VridVal := 1

	resourceDataMap := map[string]interface{}{
		"vrid_val": VridVal,

		"floating_ip": map[string]interface{}{},

		"blade_parameters": map[string]interface{}{},
	}

	resourceSchema := map[string]*schema.Schema{
		"vrid_val": {
			Type: schema.TypeInt,
		},

		"blade_parameters": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"fail_over_policy_template": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"tracking_options": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"route": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"ip_destination_cfg": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"protocol": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
														"ip_destination": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
														"distance": {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "",
														},
														"priority_cost": {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "",
														},
														"gateway": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
														"mask": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
													},
												},
											},
											"ipv6_destination_cfg": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"protocol": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
														"distance": {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "",
														},
														"gatewayv6": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
														"priority_cost": {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "",
														},
														"ipv6_destination": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
													},
												},
											},
										},
									},
								},
								"bgp": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"bgp_ipv6_address_cfg": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"bgp_ipv6_address": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
														"priority_cost": {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "",
														},
													},
												},
											},
											"bgp_ipv4_address_cfg": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"priority_cost": {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "",
														},
														"bgp_ipv4_address": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
													},
												},
											},
										},
									},
								},
								"trunk_cfg": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"priority_cost": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "",
											},
											"trunk": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "",
											},
											"per_port_pri": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "",
											},
										},
									},
								},
								"interface": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"ethernet": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "",
											},
											"priority_cost": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "",
											},
										},
									},
								},
								"uuid": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "",
								},
								"gateway": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"ipv4_gateway_list": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"ip_address": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
														"priority_cost": {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "",
														},
														"uuid": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
													},
												},
											},
											"ipv6_gateway_list": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"priority_cost": {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "",
														},
														"uuid": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
														"ipv6_address": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "",
														},
													},
												},
											},
										},
									},
								},
								"vlan_cfg": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"vlan": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "",
											},
											"priority_cost": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "",
											},
											"timeout": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "",
											},
										},
									},
								},
							},
						},
					},
					"priority": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"uuid": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
				},
			},
		},

		"floating_ip": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ipv6_address_cfg": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"ethernet": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "",
								},
								"trunk": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "",
								},
								"ipv6_address": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "",
								},
								"ve": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "",
								},
							},
						},
					},
					"ip_address_part_cfg": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"ip_address_partition": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "",
								},
							},
						},
					},
					"ipv6_address_part_cfg": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"ethernet": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "",
								},
								"ipv6_address_partition": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "",
								},
								"trunk": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "",
								},
								"ve": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "",
								},
							},
						},
					},
					"ip_address_cfg": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"ip_address": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "",
								},
							},
						},
					},
				},
			},
		},
	}

	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)

	var sInstance go_vthunder.Vrid
	var s go_vthunder.VridInstance

	var floating_ip go_vthunder.FloatingIP

	floating_ip.IPAddress = make([]go_vthunder.IPAddressCfg, 0, 1)

	var blade_parameters go_vthunder.BladeParameters
	blade_parameters.Priority = 0

	sInstance.VridVal = VridVal
	sInstance.Ipv6Address = floating_ip
	sInstance.UUID_BladeParams = blade_parameters

	s.UUID = sInstance

	cases := []struct {
		input  *schema.ResourceData
		output go_vthunder.VridInstance
	}{
		{
			input:  resourceLocalData,
			output: s,
		},
	}

	for _, tc := range cases {
		output := dataToVrrpVrid(resourceLocalData)
		if !reflect.DeepEqual(output, tc.output) {
			t.Fatalf("Got:\n\n%#v\n\nExpected:\n\n%#v", output, tc.output)
		}
	}

}

func testCheckVrrpVridExists(name string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		client := testAccProvider.Meta().(vThunder)
		vs, err := go_vthunder.GetVrrpVrid(client.Token, name, client.Host)

		if err != nil {
			return err
		}
		if exists && vs == nil {
			return fmt.Errorf(" vrrp vrid %s was not created.", name)
		}
		if !exists && vs != nil {
			return fmt.Errorf(" vrrp vrid %s still exists.", name)
		}
		return nil
	}
}

func testCheckVrrpVridDestroyed(s *terraform.State) error {
	client := testAccProvider.Meta().(vThunder)
	for _, rs := range s.RootModule().Resources {

		name := rs.Primary.ID
		sr, err := go_vthunder.GetVrrpVrid(client.Token, name, client.Host)
		if err != nil {
			return err
		}
		if sr == nil {
			return fmt.Errorf(" vrrp vrid %s not destroyed.", name)
		}
	}
	return nil
}
