package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_monitor`: Monitor template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateMonitorCreate,
		UpdateContext: resourceSlbTemplateMonitorUpdate,
		ReadContext:   resourceSlbTemplateMonitorRead,
		DeleteContext: resourceSlbTemplateMonitorDelete,

		Schema: map[string]*schema.Schema{
			"clear_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sessions": {
							Type: schema.TypeString, Optional: true, Description: "'all': Clear all sessions; 'sequence': Sequence number;",
						},
						"clear_all_sequence": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the port physical port number)",
						},
						"clear_sequence": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number",
						},
					},
				},
			},
			"id1": {
				Type: schema.TypeInt, Required: true, Description: "Monitor template ID Number",
			},
			"link_disable_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"diseth": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the physical port number (Ethernet interface number)",
						},
						"dis_sequence": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the physical port number)",
						},
					},
				},
			},
			"link_down_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"linkdown_ethernet1": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
						},
						"link_down_sequence1": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the port physical port number)",
						},
						"linkdown_ethernet2": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
						},
						"link_down_sequence2": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the port physical port number)",
						},
						"linkdown_ethernet3": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
						},
						"link_down_sequence3": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the port physical port number)",
						},
					},
				},
			},
			"link_enable_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enaeth": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the physical port number (Ethernet interface number)",
						},
						"ena_sequence": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the physical port number)",
						},
					},
				},
			},
			"link_up_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"linkup_ethernet1": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
						},
						"link_up_sequence1": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the port physical port number)",
						},
						"linkup_ethernet2": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
						},
						"link_up_sequence2": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the port physical port number)",
						},
						"linkup_ethernet3": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
						},
						"link_up_sequence3": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the port physical port number)",
						},
					},
				},
			},
			"monitor_relation": {
				Type: schema.TypeString, Optional: true, Default: "monitor-and", Description: "'monitor-and': Configures the monitors in current template to work with AND logic; 'monitor-or': Configures the monitors in current template to work with OR logic;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateMonitorClearCfg(d []interface{}) []edpt.SlbTemplateMonitorClearCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateMonitorClearCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateMonitorClearCfg
		oi.Sessions = in["sessions"].(string)
		oi.ClearAllSequence = in["clear_all_sequence"].(int)
		oi.ClearSequence = in["clear_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateMonitorLinkDisableCfg(d []interface{}) []edpt.SlbTemplateMonitorLinkDisableCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateMonitorLinkDisableCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateMonitorLinkDisableCfg
		oi.Diseth = in["diseth"].(int)
		oi.DisSequence = in["dis_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateMonitorLinkDownCfg(d []interface{}) []edpt.SlbTemplateMonitorLinkDownCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateMonitorLinkDownCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateMonitorLinkDownCfg
		oi.LinkdownEthernet1 = in["linkdown_ethernet1"].(int)
		oi.LinkDownSequence1 = in["link_down_sequence1"].(int)
		oi.LinkdownEthernet2 = in["linkdown_ethernet2"].(int)
		oi.LinkDownSequence2 = in["link_down_sequence2"].(int)
		oi.LinkdownEthernet3 = in["linkdown_ethernet3"].(int)
		oi.LinkDownSequence3 = in["link_down_sequence3"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateMonitorLinkEnableCfg(d []interface{}) []edpt.SlbTemplateMonitorLinkEnableCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateMonitorLinkEnableCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateMonitorLinkEnableCfg
		oi.Enaeth = in["enaeth"].(int)
		oi.EnaSequence = in["ena_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateMonitorLinkUpCfg(d []interface{}) []edpt.SlbTemplateMonitorLinkUpCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateMonitorLinkUpCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateMonitorLinkUpCfg
		oi.LinkupEthernet1 = in["linkup_ethernet1"].(int)
		oi.LinkUpSequence1 = in["link_up_sequence1"].(int)
		oi.LinkupEthernet2 = in["linkup_ethernet2"].(int)
		oi.LinkUpSequence2 = in["link_up_sequence2"].(int)
		oi.LinkupEthernet3 = in["linkup_ethernet3"].(int)
		oi.LinkUpSequence3 = in["link_up_sequence3"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateMonitor(d *schema.ResourceData) edpt.SlbTemplateMonitor {
	var ret edpt.SlbTemplateMonitor
	ret.Inst.ClearCfg = getSliceSlbTemplateMonitorClearCfg(d.Get("clear_cfg").([]interface{}))
	ret.Inst.Id1 = d.Get("id1").(int)
	ret.Inst.LinkDisableCfg = getSliceSlbTemplateMonitorLinkDisableCfg(d.Get("link_disable_cfg").([]interface{}))
	ret.Inst.LinkDownCfg = getSliceSlbTemplateMonitorLinkDownCfg(d.Get("link_down_cfg").([]interface{}))
	ret.Inst.LinkEnableCfg = getSliceSlbTemplateMonitorLinkEnableCfg(d.Get("link_enable_cfg").([]interface{}))
	ret.Inst.LinkUpCfg = getSliceSlbTemplateMonitorLinkUpCfg(d.Get("link_up_cfg").([]interface{}))
	ret.Inst.MonitorRelation = d.Get("monitor_relation").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
