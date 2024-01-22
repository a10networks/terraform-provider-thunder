package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetLldp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_lldp`: Interface lldp configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetLldpCreate,
		UpdateContext: resourceInterfaceEthernetLldpUpdate,
		ReadContext:   resourceInterfaceEthernetLldpRead,
		DeleteContext: resourceInterfaceEthernetLldpDelete,

		Schema: map[string]*schema.Schema{
			"enable_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rt_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp enable/disable",
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
			"notification_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp notification configuration",
						},
						"notif_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp notification enable",
						},
					},
				},
			},
			"tx_dot1_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tx_dot1_tlvs": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp tx IEEE 802.1 Organizationally specific TLVs configuration",
						},
						"link_aggregation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface link aggregation information",
						},
						"vlan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface vlan information",
						},
					},
				},
			},
			"tx_tlvs_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tx_tlvs": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp tx TLVs configuration",
						},
						"exclude": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure which TLVs excluded. All basic TLVs will be included by default",
						},
						"management_address": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp management address",
						},
						"port_description": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp port description",
						},
						"system_capabilities": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp system capabilities",
						},
						"system_description": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp system description",
						},
						"system_name": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp system name",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceEthernetLldpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetLldpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetLldp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetLldpRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetLldpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetLldpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetLldp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetLldpRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetLldpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetLldpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetLldp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetLldpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetLldpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetLldp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceEthernetLldpEnableCfg(d []interface{}) edpt.InterfaceEthernetLldpEnableCfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetLldpEnableCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RtEnable = in["rt_enable"].(int)
		ret.Rx = in["rx"].(int)
		ret.Tx = in["tx"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetLldpNotificationCfg(d []interface{}) edpt.InterfaceEthernetLldpNotificationCfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetLldpNotificationCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Notification = in["notification"].(int)
		ret.NotifEnable = in["notif_enable"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetLldpTxDot1Cfg(d []interface{}) edpt.InterfaceEthernetLldpTxDot1Cfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetLldpTxDot1Cfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TxDot1Tlvs = in["tx_dot1_tlvs"].(int)
		ret.LinkAggregation = in["link_aggregation"].(int)
		ret.Vlan = in["vlan"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetLldpTxTlvsCfg(d []interface{}) edpt.InterfaceEthernetLldpTxTlvsCfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetLldpTxTlvsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TxTlvs = in["tx_tlvs"].(int)
		ret.Exclude = in["exclude"].(int)
		ret.ManagementAddress = in["management_address"].(int)
		ret.PortDescription = in["port_description"].(int)
		ret.SystemCapabilities = in["system_capabilities"].(int)
		ret.SystemDescription = in["system_description"].(int)
		ret.SystemName = in["system_name"].(int)
	}
	return ret
}

func dataToEndpointInterfaceEthernetLldp(d *schema.ResourceData) edpt.InterfaceEthernetLldp {
	var ret edpt.InterfaceEthernetLldp
	ret.Inst.EnableCfg = getObjectInterfaceEthernetLldpEnableCfg(d.Get("enable_cfg").([]interface{}))
	ret.Inst.NotificationCfg = getObjectInterfaceEthernetLldpNotificationCfg(d.Get("notification_cfg").([]interface{}))
	ret.Inst.TxDot1Cfg = getObjectInterfaceEthernetLldpTxDot1Cfg(d.Get("tx_dot1_cfg").([]interface{}))
	ret.Inst.TxTlvsCfg = getObjectInterfaceEthernetLldpTxTlvsCfg(d.Get("tx_tlvs_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
