package thunder

//Thunder resource - Ethernet

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEthernet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEthernetCreate,
		UpdateContext: resourceEthernetUpdate,
		ReadContext:   resourceEthernetRead,
		DeleteContext: resourceEthernetDelete,

		Schema: map[string]*schema.Schema{
			"ethernet_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"l3_vlan_fwd_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"trap_source": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"flow_control": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ifnum": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"speed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"mtu": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ip": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dhcp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"address_list": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"ipv4_netmask": {
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
						"duplexity": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceEthernetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating ip (Inside resourceEthernetCreate)")
		name := d.Get("ethernet_list.0.ifnum").(int)
		ethernet := dataToEthernet(d)
		d.SetId(strconv.Itoa(name))
		logger.Println("[INFO] received V from method data to ip --")
		err := go_thunder.PutEthernet(client.Token, ethernet, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEthernetRead(ctx, d, meta)
	}
	return diags
}

func resourceEthernetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading vrrp common (Inside resourceEthernetRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		logger.Println("[INFO] Fetching ethernet" + name)

		vc, err := go_thunder.GetEthernet(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No ethernet found")
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceEthernetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceEthernetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

// Utility method for Ethernet structure instantiation
func dataToEthernet(d *schema.ResourceData) map[int]go_thunder.Ethernet {
	var m = make(map[int]go_thunder.Ethernet)
	ethCount := d.Get("ethernet_list.#").(int)

	for i := 0; i < ethCount; i++ {
		prefix := fmt.Sprintf("ethernet_list.%d", i)
		var ethernet go_thunder.Ethernet
		var ethernetInst go_thunder.EthernetInstance

		var ip go_thunder.IP
		ip.UUID = d.Get(prefix + ".ip.0.uuid").(string)

		ip.Ipv4Address = make([]go_thunder.Address, 0, 1)
		var ad go_thunder.Address
		ad.Ipv4Address = d.Get(prefix + ".ip.0.address_list.0.ipv4_address").(string)
		ad.Ipv4Netmask = d.Get(prefix + ".ip.0.address_list.0.ipv4_netmask").(string)
		ip.Ipv4Address = append(ip.Ipv4Address, ad)

		ethernetInst.Dhcp = ip

		ethernetInst.LoadInterval = d.Get(prefix + ".load_interval").(int)
		ethernetInst.L3VlanFwdDisable = d.Get(prefix + ".l3_vlan_fwd_disable").(int)
		ethernetInst.FlowControl = d.Get(prefix + ".flow_control").(int)
		ethernetInst.Action = d.Get(prefix + ".action").(string)

		ethernetInst.Ifnum = d.Get(prefix + ".ifnum").(int)
		ethernetInst.Speed = d.Get(prefix + ".speed").(string)
		ethernetInst.Mtu = d.Get(prefix + ".mtu").(int)
		ethernetInst.Duplexity = d.Get(prefix + ".duplexity").(string)

		ethernet.UUID = ethernetInst

		m[ethernetInst.Ifnum] = ethernet

	}

	return m
}
