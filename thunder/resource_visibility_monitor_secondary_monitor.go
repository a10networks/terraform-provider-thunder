package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorSecondaryMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor_secondary_monitor`: Configure secondary monitoring key\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorSecondaryMonitorCreate,
		UpdateContext: resourceVisibilityMonitorSecondaryMonitorUpdate,
		ReadContext:   resourceVisibilityMonitorSecondaryMonitorRead,
		DeleteContext: resourceVisibilityMonitorSecondaryMonitorDelete,

		Schema: map[string]*schema.Schema{
			"debug_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"debug_ip_addr": {
							Type: schema.TypeString, Required: true, Description: "Specify source/dest ip addr",
						},
						"debug_port": {
							Type: schema.TypeInt, Required: true, Description: "Specify port",
						},
						"debug_protocol": {
							Type: schema.TypeString, Required: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"delete_debug_file": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"debug_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify source/dest ip addr",
						},
						"debug_port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify port",
						},
						"debug_protocol": {
							Type: schema.TypeString, Optional: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
						},
					},
				},
			},
			"mon_entity_topk": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable topk for secondary entities",
			},
			"replay_debug_file": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"debug_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify source/dest ip addr",
						},
						"debug_port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify port",
						},
						"debug_protocol": {
							Type: schema.TypeString, Optional: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
						},
					},
				},
			},
			"secondary_monitoring_key": {
				Type: schema.TypeString, Required: true, Description: "'service': Monitor traffic to any service;",
			},
			"source_entity_topk": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable topk for sources to secondary-entities",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityMonitorSecondaryMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorSecondaryMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorSecondaryMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorSecondaryMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorSecondaryMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorSecondaryMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVisibilityMonitorSecondaryMonitorDebugList(d []interface{}) []edpt.VisibilityMonitorSecondaryMonitorDebugList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitorSecondaryMonitorDebugList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitorSecondaryMonitorDebugList
		oi.DebugIpAddr = in["debug_ip_addr"].(string)
		oi.DebugPort = in["debug_port"].(int)
		oi.DebugProtocol = in["debug_protocol"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitorSecondaryMonitorDeleteDebugFile1914(d []interface{}) edpt.VisibilityMonitorSecondaryMonitorDeleteDebugFile1914 {

	count1 := len(d)
	var ret edpt.VisibilityMonitorSecondaryMonitorDeleteDebugFile1914
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DebugIpAddr = in["debug_ip_addr"].(string)
		ret.DebugPort = in["debug_port"].(int)
		ret.DebugProtocol = in["debug_protocol"].(string)
	}
	return ret
}

func getObjectVisibilityMonitorSecondaryMonitorReplayDebugFile1915(d []interface{}) edpt.VisibilityMonitorSecondaryMonitorReplayDebugFile1915 {

	count1 := len(d)
	var ret edpt.VisibilityMonitorSecondaryMonitorReplayDebugFile1915
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DebugIpAddr = in["debug_ip_addr"].(string)
		ret.DebugPort = in["debug_port"].(int)
		ret.DebugProtocol = in["debug_protocol"].(string)
	}
	return ret
}

func dataToEndpointVisibilityMonitorSecondaryMonitor(d *schema.ResourceData) edpt.VisibilityMonitorSecondaryMonitor {
	var ret edpt.VisibilityMonitorSecondaryMonitor
	ret.Inst.DebugList = getSliceVisibilityMonitorSecondaryMonitorDebugList(d.Get("debug_list").([]interface{}))
	ret.Inst.DeleteDebugFile = getObjectVisibilityMonitorSecondaryMonitorDeleteDebugFile1914(d.Get("delete_debug_file").([]interface{}))
	ret.Inst.MonEntityTopk = d.Get("mon_entity_topk").(int)
	ret.Inst.ReplayDebugFile = getObjectVisibilityMonitorSecondaryMonitorReplayDebugFile1915(d.Get("replay_debug_file").([]interface{}))
	ret.Inst.SecondaryMonitoringKey = d.Get("secondary_monitoring_key").(string)
	ret.Inst.SourceEntityTopk = d.Get("source_entity_topk").(int)
	//omit uuid
	return ret
}
