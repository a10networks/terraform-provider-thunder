package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkLldp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_lldp`: Configure LLDP\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkLldpCreate,
		UpdateContext: resourceNetworkLldpUpdate,
		ReadContext:   resourceNetworkLldpRead,
		DeleteContext: resourceNetworkLldpDelete,

		Schema: map[string]*schema.Schema{
			"enable_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lldp",
						},
						"rx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lldp rx",
						},
						"tx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lldp tx",
						},
					},
				},
			},
			"management_address": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns": {
										Type: schema.TypeString, Required: true, Description: "Configure lldp management-address, subtype is dns (lldp management-address dns address)",
									},
									"interface": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ethernet": {
													Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface ethernet (lldp management-address interface port number)",
												},
												"ve": {
													Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface management (lldp management-address interface port number)",
												},
												"management": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "configure lldp management-address interface management",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ipv4_addr_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4": {
										Type: schema.TypeString, Required: true, Description: "Configure lldp management-address, subtype is ipv4 (lldp management-address ipv4 address)",
									},
									"interface_ipv4": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_eth": {
													Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface ethernet (lldp management-address interface port number)",
												},
												"ipv4_ve": {
													Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface ve (lldp management-address interface port number)",
												},
												"ipv4_mgmt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "configure lldp management-address interface management",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ipv6_addr_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6": {
										Type: schema.TypeString, Required: true, Description: "Configure lldp management-address, subtype is ipv6 (lldp management-address ipv6 address)",
									},
									"interface_ipv6": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv6_eth": {
													Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface ethernet (lldp management-address interface port number)",
												},
												"ipv6_ve": {
													Type: schema.TypeInt, Optional: true, Description: "configure lldp management-address interface ve (lldp management-address interface port number)",
												},
												"ipv6_mgmt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "configure lldp management-address interface management",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"notification_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lldp notification",
						},
						"interval": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Configure lldp notification interval, default is 30 (The lldp notification interval value, default is 30)",
						},
					},
				},
			},
			"system_description": {
				Type: schema.TypeString, Optional: true, Description: "Configure lldp system description",
			},
			"system_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure lldp system name",
			},
			"tx_set": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fast_count": {
							Type: schema.TypeInt, Optional: true, Default: 4, Description: "Configure lldp tx fast count value (The lldp tx fast count value, default is 4)",
						},
						"fast_interval": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure lldp tx fast interval value (The lldp tx fast interval value, default is 1)",
						},
						"hold": {
							Type: schema.TypeInt, Optional: true, Default: 4, Description: "Configure lldp tx hold multiplier (The lldp tx hold value, default is 4)",
						},
						"tx_interval": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Configure lldp tx interval (The lldp tx interval value, default is 30)",
						},
						"reinit_delay": {
							Type: schema.TypeInt, Optional: true, Default: 2, Description: "Configure lldp tx reinit delay (The lldp tx reinit_delay value, default is 2)",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkLldpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkLldpRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkLldpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkLldpRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkLldpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkLldpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLldpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLldp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectNetworkLldpEnableCfg(d []interface{}) edpt.NetworkLldpEnableCfg {

	count1 := len(d)
	var ret edpt.NetworkLldpEnableCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Rx = in["rx"].(int)
		ret.Tx = in["tx"].(int)
	}
	return ret
}

func getObjectNetworkLldpManagementAddress1073(d []interface{}) edpt.NetworkLldpManagementAddress1073 {

	count1 := len(d)
	var ret edpt.NetworkLldpManagementAddress1073
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsList = getSliceNetworkLldpManagementAddressDnsList(in["dns_list"].([]interface{}))
		ret.Ipv4AddrList = getSliceNetworkLldpManagementAddressIpv4AddrList(in["ipv4_addr_list"].([]interface{}))
		ret.Ipv6AddrList = getSliceNetworkLldpManagementAddressIpv6AddrList(in["ipv6_addr_list"].([]interface{}))
	}
	return ret
}

func getSliceNetworkLldpManagementAddressDnsList(d []interface{}) []edpt.NetworkLldpManagementAddressDnsList {

	count1 := len(d)
	ret := make([]edpt.NetworkLldpManagementAddressDnsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkLldpManagementAddressDnsList
		oi.Dns = in["dns"].(string)
		oi.Interface = getObjectNetworkLldpManagementAddressDnsListInterface(in["interface"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNetworkLldpManagementAddressDnsListInterface(d []interface{}) edpt.NetworkLldpManagementAddressDnsListInterface {

	count1 := len(d)
	var ret edpt.NetworkLldpManagementAddressDnsListInterface
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethernet = in["ethernet"].(int)
		ret.Ve = in["ve"].(int)
		ret.Management = in["management"].(int)
	}
	return ret
}

func getSliceNetworkLldpManagementAddressIpv4AddrList(d []interface{}) []edpt.NetworkLldpManagementAddressIpv4AddrList {

	count1 := len(d)
	ret := make([]edpt.NetworkLldpManagementAddressIpv4AddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkLldpManagementAddressIpv4AddrList
		oi.Ipv4 = in["ipv4"].(string)
		oi.InterfaceIpv4 = getObjectNetworkLldpManagementAddressIpv4AddrListInterfaceIpv4(in["interface_ipv4"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNetworkLldpManagementAddressIpv4AddrListInterfaceIpv4(d []interface{}) edpt.NetworkLldpManagementAddressIpv4AddrListInterfaceIpv4 {

	count1 := len(d)
	var ret edpt.NetworkLldpManagementAddressIpv4AddrListInterfaceIpv4
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Eth = in["ipv4_eth"].(int)
		ret.Ipv4Ve = in["ipv4_ve"].(int)
		ret.Ipv4Mgmt = in["ipv4_mgmt"].(int)
	}
	return ret
}

func getSliceNetworkLldpManagementAddressIpv6AddrList(d []interface{}) []edpt.NetworkLldpManagementAddressIpv6AddrList {

	count1 := len(d)
	ret := make([]edpt.NetworkLldpManagementAddressIpv6AddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkLldpManagementAddressIpv6AddrList
		oi.Ipv6 = in["ipv6"].(string)
		oi.InterfaceIpv6 = getObjectNetworkLldpManagementAddressIpv6AddrListInterfaceIpv6(in["interface_ipv6"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNetworkLldpManagementAddressIpv6AddrListInterfaceIpv6(d []interface{}) edpt.NetworkLldpManagementAddressIpv6AddrListInterfaceIpv6 {

	count1 := len(d)
	var ret edpt.NetworkLldpManagementAddressIpv6AddrListInterfaceIpv6
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6Eth = in["ipv6_eth"].(int)
		ret.Ipv6Ve = in["ipv6_ve"].(int)
		ret.Ipv6Mgmt = in["ipv6_mgmt"].(int)
	}
	return ret
}

func getObjectNetworkLldpNotificationCfg(d []interface{}) edpt.NetworkLldpNotificationCfg {

	count1 := len(d)
	var ret edpt.NetworkLldpNotificationCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Notification = in["notification"].(int)
		ret.Interval = in["interval"].(int)
	}
	return ret
}

func getObjectNetworkLldpTxSet(d []interface{}) edpt.NetworkLldpTxSet {

	count1 := len(d)
	var ret edpt.NetworkLldpTxSet
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FastCount = in["fast_count"].(int)
		ret.FastInterval = in["fast_interval"].(int)
		ret.Hold = in["hold"].(int)
		ret.TxInterval = in["tx_interval"].(int)
		ret.ReinitDelay = in["reinit_delay"].(int)
	}
	return ret
}

func dataToEndpointNetworkLldp(d *schema.ResourceData) edpt.NetworkLldp {
	var ret edpt.NetworkLldp
	ret.Inst.EnableCfg = getObjectNetworkLldpEnableCfg(d.Get("enable_cfg").([]interface{}))
	ret.Inst.ManagementAddress = getObjectNetworkLldpManagementAddress1073(d.Get("management_address").([]interface{}))
	ret.Inst.NotificationCfg = getObjectNetworkLldpNotificationCfg(d.Get("notification_cfg").([]interface{}))
	ret.Inst.SystemDescription = d.Get("system_description").(string)
	ret.Inst.SystemName = d.Get("system_name").(string)
	ret.Inst.TxSet = getObjectNetworkLldpTxSet(d.Get("tx_set").([]interface{}))
	//omit uuid
	return ret
}
