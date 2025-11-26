package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorAgent() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor_agent`: Configure xflow agent\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorAgentCreate,
		UpdateContext: resourceVisibilityMonitorAgentUpdate,
		ReadContext:   resourceVisibilityMonitorAgentRead,
		DeleteContext: resourceVisibilityMonitorAgentDelete,

		Schema: map[string]*schema.Schema{
			"agent_name": {
				Type: schema.TypeString, Required: true, Description: "Specify name for the agent",
			},
			"agent_v4_addr": {
				Type: schema.TypeString, Optional: true, Description: "Configure agent's IPv4 address",
			},
			"agent_v6_addr": {
				Type: schema.TypeString, Optional: true, Description: "Configure agent's IPv6 address",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'sflow-packets-received': sFlow Packets Received; 'sflow-samples-received': sFlow Samples Received; 'sflow-samples-bad-len': sFlow Samples Bad Length; 'sflow-samples-non-std': sFlow Samples Non-standard; 'sflow-samples-skipped': sFlow Samples Skipped; 'sflow-sample-record-bad-len': sFlow Sample Records Bad Length; 'sflow-samples-sent-for-detection': sFlow Samples Processed For Detection; 'sflow-sample-record-invalid-layer2': sFlow Sample Records Unknown Layer-2; 'sflow-sample-ipv6-hdr-parse-fail': sFlow Sample IPv6 Record Header Parse Failures; 'sflow-disabled': sFlow Packet Samples Processing Disabled; 'netflow-disabled': Netflow Flow Samples Processing Disabled; 'netflow-v5-packets-received': Netflow v5 Packets Received; 'netflow-v5-samples-received': Netflow v5 Samples Received; 'netflow-v5-samples-sent-for-detection': Netflow v5 Samples Processed For Detection; 'netflow-v5-sample-records-bad-len': Netflow v5 Sample Records Bad Length; 'netflow-v5-max-records-exceed': Netflow v5 Sample Max Records Error; 'netflow-v9-packets-received': Netflow v9 Packets Received; 'netflow-v9-samples-received': Netflow v9 Samples Received; 'netflow-v9-samples-sent-for-detection': Netflow v9 Samples Processed For Detection; 'netflow-v9-sample-records-bad-len': Netflow v9 Sample Records Bad Length; 'netflow-v9-max-records-exceed': Netflow v9 Sample Max Records Error; 'netflow-v10-packets-received': Netflow v10 Packets Received; 'netflow-v10-samples-received': Netflow v10 Samples Received; 'netflow-v10-samples-sent-for-detection': Netflow v10 Samples Procssed For Detection; 'netflow-v10-sample-records-bad-len': Netflow v10 Sample Records Bad Length; 'netflow-v10-max-records-exceed': Netflow v10 Sample Max records Error; 'netflow-tcp-sample-received': Netflow TCP Samples Received; 'netflow-udp-sample-received': Netflow UDP Samples received; 'netflow-icmp-sample-received': Netflow ICMP Samples Received; 'netflow-other-sample-received': Netflow OTHER Samples Received; 'netflow-record-copy-oom-error': Netflow Data Record Copy Fail OOM; 'netflow-record-rse-invalid': Netflow Data Record Reduced Size Invalid; 'netflow-sample-flow-dur-error': Netflow Sample Flow Duration Error;",
						},
					},
				},
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
func resourceVisibilityMonitorAgentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorAgentCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorAgent(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorAgentRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorAgentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorAgentUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorAgent(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorAgentRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorAgentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorAgentDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorAgent(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorAgentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorAgentRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorAgent(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVisibilityMonitorAgentSamplingEnable(d []interface{}) []edpt.VisibilityMonitorAgentSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitorAgentSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitorAgentSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonitorAgent(d *schema.ResourceData) edpt.VisibilityMonitorAgent {
	var ret edpt.VisibilityMonitorAgent
	ret.Inst.AgentName = d.Get("agent_name").(string)
	ret.Inst.AgentV4Addr = d.Get("agent_v4_addr").(string)
	ret.Inst.AgentV6Addr = d.Get("agent_v6_addr").(string)
	ret.Inst.SamplingEnable = getSliceVisibilityMonitorAgentSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
