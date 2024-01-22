package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceManagementLldp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_management_lldp`: Interface lldp configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceManagementLldpCreate,
		UpdateContext: resourceInterfaceManagementLldpUpdate,
		ReadContext:   resourceInterfaceManagementLldpRead,
		DeleteContext: resourceInterfaceManagementLldpDelete,

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
		},
	}
}
func resourceInterfaceManagementLldpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagementLldpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagementLldp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceManagementLldpRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceManagementLldpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagementLldpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagementLldp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceManagementLldpRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceManagementLldpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagementLldpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagementLldp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceManagementLldpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagementLldpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagementLldp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceManagementLldpEnableCfg(d []interface{}) edpt.InterfaceManagementLldpEnableCfg {

	count1 := len(d)
	var ret edpt.InterfaceManagementLldpEnableCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RtEnable = in["rt_enable"].(int)
		ret.Rx = in["rx"].(int)
		ret.Tx = in["tx"].(int)
	}
	return ret
}

func getObjectInterfaceManagementLldpNotificationCfg(d []interface{}) edpt.InterfaceManagementLldpNotificationCfg {

	count1 := len(d)
	var ret edpt.InterfaceManagementLldpNotificationCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Notification = in["notification"].(int)
		ret.NotifEnable = in["notif_enable"].(int)
	}
	return ret
}

func getObjectInterfaceManagementLldpTxDot1Cfg(d []interface{}) edpt.InterfaceManagementLldpTxDot1Cfg {

	count1 := len(d)
	var ret edpt.InterfaceManagementLldpTxDot1Cfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TxDot1Tlvs = in["tx_dot1_tlvs"].(int)
		ret.LinkAggregation = in["link_aggregation"].(int)
		ret.Vlan = in["vlan"].(int)
	}
	return ret
}

func getObjectInterfaceManagementLldpTxTlvsCfg(d []interface{}) edpt.InterfaceManagementLldpTxTlvsCfg {

	count1 := len(d)
	var ret edpt.InterfaceManagementLldpTxTlvsCfg
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

func dataToEndpointInterfaceManagementLldp(d *schema.ResourceData) edpt.InterfaceManagementLldp {
	var ret edpt.InterfaceManagementLldp
	ret.Inst.EnableCfg = getObjectInterfaceManagementLldpEnableCfg(d.Get("enable_cfg").([]interface{}))
	ret.Inst.NotificationCfg = getObjectInterfaceManagementLldpNotificationCfg(d.Get("notification_cfg").([]interface{}))
	ret.Inst.TxDot1Cfg = getObjectInterfaceManagementLldpTxDot1Cfg(d.Get("tx_dot1_cfg").([]interface{}))
	ret.Inst.TxTlvsCfg = getObjectInterfaceManagementLldpTxTlvsCfg(d.Get("tx_tlvs_cfg").([]interface{}))
	//omit uuid
	return ret
}
