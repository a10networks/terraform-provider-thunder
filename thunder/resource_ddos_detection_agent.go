package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionAgent() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_agent`: Configure DDoS detection agent\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionAgentCreate,
		UpdateContext: resourceDdosDetectionAgentUpdate,
		ReadContext:   resourceDdosDetectionAgentRead,
		DeleteContext: resourceDdosDetectionAgentDelete,

		Schema: map[string]*schema.Schema{
			"agent_name": {
				Type: schema.TypeString, Required: true, Description: "Specify name for the agent",
			},
			"agent_type": {
				Type: schema.TypeString, Optional: true, Description: "'Cisco': Cisco; 'Juniper': Juniper;",
			},
			"agent_v4_addr": {
				Type: schema.TypeString, Optional: true, Description: "Configure agent's IPv4 address",
			},
			"agent_v6_addr": {
				Type: schema.TypeString, Optional: true, Description: "Configure agent's IPv6 address",
			},
			"netflow": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"netflow_samples_collection": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Netflow flow samples collection(default); 'disable': Disable Netflow flow samples collection;",
						},
						"netflow_sampling_rate": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure agent's netflow sampling rate",
						},
						"active_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Configure agent's flow active timeout (seconds)",
						},
						"inactive_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Configure agent's flow inactive timeout (seconds)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'sflow-packets-received': sFlow Packets Received; 'sflow-samples-received': sFlow Samples Received; 'sflow-samples-bad-len': sFlow Samples Bad Length; 'sflow-samples-non-std': sFlow Samples Non-standard; 'sflow-samples-skipped': sFlow Samples Skipped; 'sflow-sample-record-bad-len': sFlow Sample Records Bad Length; 'sflow-samples-sent-for-detection': sFlow Samples Processed For Detection; 'sflow-sample-record-invalid-layer2': sFlow Sample Records Unknown Layer-2; 'sflow-sample-ipv6-hdr-parse-fail': sFlow Sample IPv6 Record Header Parse Failures; 'sflow-disabled': sFlow Packet Samples Processing Disabled; 'netflow-disabled': Netflow Flow Samples Processing Disabled; 'netflow-v5-packets-received': Netflow v5 Packets Received; 'netflow-v5-samples-received': Netflow v5 Samples Received; 'netflow-v5-samples-sent-for-detection': Netflow v5 Samples Processed For Detection; 'netflow-v5-sample-records-bad-len': Netflow v5 Sample Records Bad Length; 'netflow-v5-max-records-exceed': Netflow v5 Sample Max Records Error; 'netflow-v9-packets-received': Netflow v9 Packets Received; 'netflow-v9-samples-received': Netflow v9 Samples Received; 'netflow-v9-samples-sent-for-detection': Netflow v9 Samples Processed For Detection; 'netflow-v9-sample-records-bad-len': Netflow v9 Sample Records Bad Length; 'netflow-v9-sample-flowset-bad-padding': Netflow v9 Sample Flowset Bad Padding; 'netflow-v9-max-records-exceed': Netflow v9 Sample Max Records Error; 'netflow-v9-template-not-found': Netflow v9 Template Not Found; 'netflow-v10-packets-received': Netflow v10 Packets Received; 'netflow-v10-samples-received': Netflow v10 Samples Received; 'netflow-v10-samples-sent-for-detection': Netflow v10 Samples Procssed For Detection; 'netflow-v10-sample-records-bad-len': Netflow v10 Sample Records Bad Length; 'netflow-v10-max-records-exceed': Netflow v10 Sample Max records Error; 'netflow-tcp-sample-received': Netflow TCP Samples Received; 'netflow-udp-sample-received': Netflow UDP Samples received; 'netflow-icmp-sample-received': Netflow ICMP Samples Received; 'netflow-other-sample-received': Netflow OTHER Samples Received; 'netflow-record-copy-oom-error': Netflow Data Record Copy Fail, Local MEM size error; 'netflow-record-rse-invalid': Netflow Data Record Reduced Size Invalid; 'netflow-sample-flow-dur-error': Netflow Sample Flow Duration Error; 'flow-dst-entry-miss': DDoS Destination Entry Lookup Failures; 'flow-ip-proto-or-port-miss': DDoS Destination Service Lookup Failures; 'flow-detection-msgq-full': Detection Message Enqueue Failures; 'flow-network-entry-miss': DDoS Destination Network-object Entry Lookup Failures;",
						},
					},
				},
			},
			"sflow": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sflow_pkt_samples_collection": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable sflow packet samples collection(default); 'disable': Disable sflow packet samples collection;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceDdosDetectionAgentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgent(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionAgentRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionAgentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgent(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionAgentRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionAgentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgent(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionAgentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgent(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDetectionAgentNetflow128(d []interface{}) edpt.DdosDetectionAgentNetflow128 {

	count1 := len(d)
	var ret edpt.DdosDetectionAgentNetflow128
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetflowSamplesCollection = in["netflow_samples_collection"].(string)
		ret.NetflowSamplingRate = in["netflow_sampling_rate"].(int)
		ret.ActiveTimeout = in["active_timeout"].(int)
		ret.InactiveTimeout = in["inactive_timeout"].(int)
		//omit uuid
	}
	return ret
}

func getSliceDdosDetectionAgentSamplingEnable(d []interface{}) []edpt.DdosDetectionAgentSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionAgentSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionAgentSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDetectionAgentSflow129(d []interface{}) edpt.DdosDetectionAgentSflow129 {

	count1 := len(d)
	var ret edpt.DdosDetectionAgentSflow129
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowPktSamplesCollection = in["sflow_pkt_samples_collection"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosDetectionAgent(d *schema.ResourceData) edpt.DdosDetectionAgent {
	var ret edpt.DdosDetectionAgent
	ret.Inst.AgentName = d.Get("agent_name").(string)
	ret.Inst.AgentType = d.Get("agent_type").(string)
	ret.Inst.AgentV4Addr = d.Get("agent_v4_addr").(string)
	ret.Inst.AgentV6Addr = d.Get("agent_v6_addr").(string)
	ret.Inst.Netflow = getObjectDdosDetectionAgentNetflow128(d.Get("netflow").([]interface{}))
	ret.Inst.SamplingEnable = getSliceDdosDetectionAgentSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Sflow = getObjectDdosDetectionAgentSflow129(d.Get("sflow").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
