package thunder

//Thunder resource vrrp vrid

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpVrid() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVrrpVridCreate,
		UpdateContext: resourceVrrpVridUpdate,
		ReadContext:   resourceVrrpVridRead,
		DeleteContext: resourceVrrpVridDelete,

		Schema: map[string]*schema.Schema{
			"preempt_mode": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"follow": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrid_lead": {
							Type:        schema.TypeString,
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
			"vrid_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceVrrpVridCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating vrrp vrid (Inside resourceVrrpVridCreate)")

	if client.Host != "" {
		name := d.Get("vrid_val").(int)
		vc := dataToVrrpVrid(d)
		d.SetId(strconv.Itoa(name))
		err := go_thunder.PostVrrpVrid(client.Token, vc, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpVridRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpVridRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading vrrp common (Inside resourceVrrpVridRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		logger.Println("[INFO] Fetching vrrp vrid" + name)

		vc, err := go_thunder.GetVrrpVrid(client.Token, name, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No vrrp vrid found")
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceVrrpVridUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying vrrp vrid   (Inside resourceVrrpVridUpdate " + name)
		v := dataToVrrpVrid(d)
		d.SetId(name)
		err := go_thunder.PutVrrpVrid(client.Token, name, v, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpVridRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpVridDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger := util.GetLoggerInstance()

	if client.Host != "" {

		name := d.Id()
		logger.Println("[INFO] Deleting vrrp vrid (Inside resourceVrrpVridDelete) " + name)

		err := go_thunder.DeleteVrrpVrid(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete vrrp vrid  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

//Utility method to instantiate Vrrp Vrid Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToVrrpVrid(d *schema.ResourceData) go_thunder.VridInstance {
	var vc go_thunder.VridInstance

	var v go_thunder.Vrid

	v.VridVal = d.Get("vrid_val").(int)

	var floating_ip go_thunder.FloatingIP

	ip_address_count := d.Get("floating_ip.0.ip_address_cfg.#").(int)

	floating_ip.IPAddress = make([]go_thunder.IPAddressCfg, 0, ip_address_count)

	for i := 0; i < ip_address_count; i++ {
		var ip_cfg go_thunder.IPAddressCfg
		prefix := fmt.Sprintf("floating_ip.0.ip_address_cfg.%d", i)
		ip_cfg.IPAddress = d.Get(prefix + ".ip_address").(string)
		floating_ip.IPAddress = append(floating_ip.IPAddress, ip_cfg)
	}

	v.Ipv6Address = floating_ip

	var blade go_thunder.BladeParameters
	blade.Priority = d.Get("blade_parameters.0.priority").(int)

	v.UUID_BladeParams = blade

	vc.UUID = v

	return vc
}
