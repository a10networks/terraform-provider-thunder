package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgRtsp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_alg_rtsp`: Change LSN RTSP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnAlgRtspCreate,
		UpdateContext: resourceCgnv6LsnAlgRtspUpdate,
		ReadContext:   resourceCgnv6LsnAlgRtspRead,
		DeleteContext: resourceCgnv6LsnAlgRtspDelete,

		Schema: map[string]*schema.Schema{
			"rtsp_value": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable RTSP ALG for LSN;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'streams-created': Streams Created; 'streams-freed': Streams Freed; 'stream-creation-failure': Stream Creation Failures; 'ports-allocated': Stream Client Ports Allocated; 'ports-freed': Stream Client Ports Freed; 'port-allocation-failure': Stream Client Port Allocation Failures; 'unknown-client-port-from-server': Server Replies With Unknown Client Ports; 'data-session-created': Data Session Created; 'data-session-freed': Data Session Freed; 'no-session-mem': Data Session Creation Failures; 'ha-sent': HA Sent; 'ha-rcv': HA RCV; 'smp-inserted': SMP Session Inserted; 'smp-removed': SMP Session Removed; 'smp-reused': SMP Session Reused; 'nat-pool-standby': New Session NAT Pool Standby; 'smp-deleted': New Session SMP Already Deleted; 'control-closed': New Session Closed; 'data-session-exists': New Session Already Exists; 'data-session-creation-failure': New Data Session Creation Failure; 'rtp-reversed': RTP Reverse Creation; 'rtcp-reversed': RTCP Reverse Creation; 'cross-cpu-sent': Cross CPU Sent; 'cross-cpu-rcv': Cross CPU Received; 'cross-cpu-no-session': Cross CPU No Session Found; 'cross-cpu-created': Cross CPU Creation; 'cross-cpu-rcv-failure': Cross CPU Receive Failure; 'data-free-smp-retry-lookup': Data Session Free SMP Retry Lookup; 'data-free-smp-not-found': Data Session Free SMP Not Found; 'ha-streams-sent': HA Streams Sent; 'ha-streams-rcv': HA Streams Received; 'ha-stream-incompatible': HA Incompatible Streams Received; 'ha-stream-exists': HA Stream Already Exists; 'ha-port-allocation-failure': HA Stream Port Allocation Failure; 'ha-data-session-sent': HA Data Session Sent; 'ha-data-session-rcv': HA Data Session Received; 'ha-data-no-smp': HA Data Session SMP Not Found; 'ha-control-closed': HA New Data Control Session Closed; 'ha-data-exists': HA New Data Session Already Exists; 'ha-extension-failure': HA Conn Extension Failure; 'ha-stream-smp-reused': HA SMP Session Reused; 'ha-stream-smp-acquire-failure': HA SMP Session Acquire Failure; 'smp-app-type-mismatch': SMP ALG App Type Mismatch;",
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
func resourceCgnv6LsnAlgRtspCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgRtspCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgRtsp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgRtspRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnAlgRtspUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgRtspUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgRtsp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgRtspRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnAlgRtspDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgRtspDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgRtsp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnAlgRtspRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgRtspRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgRtsp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnAlgRtspSamplingEnable(d []interface{}) []edpt.Cgnv6LsnAlgRtspSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnAlgRtspSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnAlgRtspSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgRtsp(d *schema.ResourceData) edpt.Cgnv6LsnAlgRtsp {
	var ret edpt.Cgnv6LsnAlgRtsp
	ret.Inst.RtspValue = d.Get("rtsp_value").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6LsnAlgRtspSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
