package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsCollectorGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_collector_group`: Configure log servers group\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsCollectorGroupCreate,
		UpdateContext: resourceAcosEventsCollectorGroupUpdate,
		ReadContext:   resourceAcosEventsCollectorGroupRead,
		DeleteContext: resourceAcosEventsCollectorGroupDelete,

		Schema: map[string]*schema.Schema{
			"facility": {
				Type: schema.TypeString, Optional: true, Default: "local0", Description: "'local0': Local use(Default); 'local1': Local use; 'local2': Local use; 'local3': Local use; 'local4': Local use; 'local5': Local use; 'local6': Local use; 'local7': Local use;  (Facility parameter for syslog messages)",
			},
			"format": {
				Type: schema.TypeString, Optional: true, Default: "syslog", Description: "'syslog': log message format is syslog (Default); 'cef': log message format is cef; 'leef': log message format is leef;",
			},
			"health_check": {
				Type: schema.TypeString, Optional: true, Description: "Health Check (Monitor Name)",
			},
			"log_distribution": {
				Type: schema.TypeString, Optional: true, Default: "round-robin", Description: "'round-robin': Log server selection will be based on round-robin (Default); 'hashing': Log messages at this node only;",
			},
			"log_server_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Member name",
						},
						"port": {
							Type: schema.TypeInt, Required: true, Description: "Port number",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify log server group name",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Description: "'udp': use udp syslog protocol to send messages to log collector; 'tcp': use tcp syslog protocol to send messages to log collector;",
			},
			"rate": {
				Type: schema.TypeInt, Optional: true, Default: 500, Description: "Specify the log message rate per second(Default 500)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'msgs_sent': Number of log messages sent; 'msgs_rate_limited': Number of rate limited log messages; 'msgs_dropped': Number of messages dropped for other reasons;",
						},
					},
				},
			},
			"server_distribution_hash": {
				Type: schema.TypeString, Optional: true, Default: "name", Description: "'name': Hashing will be based on log server name (Default); 'ip-tuple': Hashing will be based on ip-tuple;",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use managament port to connect to the log servers",
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
func resourceAcosEventsCollectorGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsCollectorGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsCollectorGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsCollectorGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsCollectorGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsCollectorGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsCollectorGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsCollectorGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsCollectorGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsCollectorGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsCollectorGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsCollectorGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsCollectorGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsCollectorGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAcosEventsCollectorGroupLogServerList(d []interface{}) []edpt.AcosEventsCollectorGroupLogServerList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsCollectorGroupLogServerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsCollectorGroupLogServerList
		oi.Name = in["name"].(string)
		oi.Port = in["port"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAcosEventsCollectorGroupSamplingEnable(d []interface{}) []edpt.AcosEventsCollectorGroupSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AcosEventsCollectorGroupSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsCollectorGroupSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosEventsCollectorGroup(d *schema.ResourceData) edpt.AcosEventsCollectorGroup {
	var ret edpt.AcosEventsCollectorGroup
	ret.Inst.Facility = d.Get("facility").(string)
	ret.Inst.Format = d.Get("format").(string)
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.LogDistribution = d.Get("log_distribution").(string)
	ret.Inst.LogServerList = getSliceAcosEventsCollectorGroupLogServerList(d.Get("log_server_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Rate = d.Get("rate").(int)
	ret.Inst.SamplingEnable = getSliceAcosEventsCollectorGroupSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServerDistributionHash = d.Get("server_distribution_hash").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
