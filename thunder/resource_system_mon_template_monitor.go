package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMonTemplateMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_mon_template_monitor`: Monitor template\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMonTemplateMonitorCreate,
		UpdateContext: resourceSystemMonTemplateMonitorUpdate,
		ReadContext:   resourceSystemMonTemplateMonitorRead,
		DeleteContext: resourceSystemMonTemplateMonitorDelete,

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
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
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
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
						},
						"linkdown_ethernet2": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
						},
						"link_down_sequence2": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the seqeuence number)",
						},
						"linkdown_ethernet3": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
						},
						"link_down_sequence3": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
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
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
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
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
						},
						"linkup_ethernet2": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
						},
						"link_up_sequence2": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequence number)",
						},
						"linkup_ethernet3": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port physical port number (Ethernet interface number)",
						},
						"link_up_sequence3": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number (Specify the sequece number)",
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
func resourceSystemMonTemplateMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMonTemplateMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMonTemplateMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMonTemplateMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMonTemplateMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMonTemplateMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemMonTemplateMonitorClearCfg(d []interface{}) []edpt.SystemMonTemplateMonitorClearCfg {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorClearCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorClearCfg
		oi.Sessions = in["sessions"].(string)
		oi.ClearAllSequence = in["clear_all_sequence"].(int)
		oi.ClearSequence = in["clear_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorLinkDisableCfg(d []interface{}) []edpt.SystemMonTemplateMonitorLinkDisableCfg {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorLinkDisableCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorLinkDisableCfg
		oi.Diseth = in["diseth"].(int)
		oi.DisSequence = in["dis_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorLinkDownCfg(d []interface{}) []edpt.SystemMonTemplateMonitorLinkDownCfg {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorLinkDownCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorLinkDownCfg
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

func getSliceSystemMonTemplateMonitorLinkEnableCfg(d []interface{}) []edpt.SystemMonTemplateMonitorLinkEnableCfg {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorLinkEnableCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorLinkEnableCfg
		oi.Enaeth = in["enaeth"].(int)
		oi.EnaSequence = in["ena_sequence"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMonTemplateMonitorLinkUpCfg(d []interface{}) []edpt.SystemMonTemplateMonitorLinkUpCfg {

	count1 := len(d)
	ret := make([]edpt.SystemMonTemplateMonitorLinkUpCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMonTemplateMonitorLinkUpCfg
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

func dataToEndpointSystemMonTemplateMonitor(d *schema.ResourceData) edpt.SystemMonTemplateMonitor {
	var ret edpt.SystemMonTemplateMonitor
	ret.Inst.ClearCfg = getSliceSystemMonTemplateMonitorClearCfg(d.Get("clear_cfg").([]interface{}))
	ret.Inst.Id1 = d.Get("id1").(int)
	ret.Inst.LinkDisableCfg = getSliceSystemMonTemplateMonitorLinkDisableCfg(d.Get("link_disable_cfg").([]interface{}))
	ret.Inst.LinkDownCfg = getSliceSystemMonTemplateMonitorLinkDownCfg(d.Get("link_down_cfg").([]interface{}))
	ret.Inst.LinkEnableCfg = getSliceSystemMonTemplateMonitorLinkEnableCfg(d.Get("link_enable_cfg").([]interface{}))
	ret.Inst.LinkUpCfg = getSliceSystemMonTemplateMonitorLinkUpCfg(d.Get("link_up_cfg").([]interface{}))
	ret.Inst.MonitorRelation = d.Get("monitor_relation").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
