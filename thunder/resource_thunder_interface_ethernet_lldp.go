package thunder

//Thunder resource InterfaceEthernetLLDP

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetLLDP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceEthernetLLDPCreate,
		UpdateContext: resourceInterfaceEthernetLLDPUpdate,
		ReadContext:   resourceInterfaceEthernetLLDPRead,
		DeleteContext: resourceInterfaceEthernetLLDPDelete,
		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"tx_tlvs_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_description": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"system_name": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"system_capabilities": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tx_tlvs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"port_description": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"exclude": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"management_address": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"tx_dot1_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_aggregation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vlan": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tx_dot1_tlvs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"enable_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rt_enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tx": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"rx": {
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
			"notification_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"notif_enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceEthernetLLDPCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceEthernetLLDP (Inside resourceInterfaceEthernetLLDPCreate) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceEthernetLLDP(d)
		logger.Println("[INFO] received V from method data to InterfaceEthernetLLDP --")
		d.SetId(strconv.Itoa(name))
		err := go_thunder.PostInterfaceEthernetLLDP(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceEthernetLLDPRead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceEthernetLLDPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading InterfaceEthernetLLDP (Inside resourceInterfaceEthernetLLDPRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetInterfaceEthernetLLDP(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceInterfaceEthernetLLDPUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceEthernetLLDPRead(ctx, d, meta)
}

func resourceInterfaceEthernetLLDPDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceEthernetLLDPRead(ctx, d, meta)
}
func dataToInterfaceEthernetLLDP(d *schema.ResourceData) go_thunder.EthernetLLDP {
	var vc go_thunder.EthernetLLDP
	var c go_thunder.EthernetLLDPInstance

	var obj1 go_thunder.LLDPTxDot1Cfg
	prefix := "tx_dot1_cfg.0."
	obj1.LinkAggregation = d.Get(prefix + "link_aggregation").(int)
	obj1.Vlan = d.Get(prefix + "vlan").(int)
	obj1.TxDot1Tlvs = d.Get(prefix + "tx_dot1_tlvs").(int)
	c.LinkAggregation = obj1

	var obj2 go_thunder.LLDPNotificationCfg
	prefix = "notification_cfg.0."
	obj2.Notification = d.Get(prefix + "notification").(int)
	obj2.NotifEnable = d.Get(prefix + "notif_enable").(int)
	c.Notification = obj2

	var obj3 go_thunder.LLDPEnableCfg
	prefix = "enable_cfg.0."
	obj3.Rx = d.Get(prefix + "rx").(int)
	obj3.Tx = d.Get(prefix + "tx").(int)
	obj3.RtEnable = d.Get(prefix + "rt_enable").(int)
	c.Rx = obj3

	var obj4 go_thunder.LLDPTxTlvsCfg
	prefix = "tx_tlvs_cfg.0."
	obj4.SystemCapabilities = d.Get(prefix + "system_capabilities").(int)
	obj4.SystemDescription = d.Get(prefix + "system_description").(int)
	obj4.ManagementAddress = d.Get(prefix + "management_address").(int)
	obj4.TxTlvs = d.Get(prefix + "tx_tlvs").(int)
	obj4.Exclude = d.Get(prefix + "exclude").(int)
	obj4.PortDescription = d.Get(prefix + "port_description").(int)
	obj4.SystemName = d.Get(prefix + "system_name").(int)
	c.SystemCapabilities = obj4

	vc.UUID = c
	return vc
}
