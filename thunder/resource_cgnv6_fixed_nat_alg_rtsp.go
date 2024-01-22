package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgRtsp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_fixed_nat_alg_rtsp`: Change Fixed NAT RTSP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6FixedNatAlgRtspCreate,
		UpdateContext: resourceCgnv6FixedNatAlgRtspUpdate,
		ReadContext:   resourceCgnv6FixedNatAlgRtspRead,
		DeleteContext: resourceCgnv6FixedNatAlgRtspDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'streams-created': Streams Created; 'streams-freed': Streams Freed; 'stream-creation-failure': Stream Creation Failures; 'ports-allocated': Stream Client Ports Allocated; 'ports-freed': Stream Client Ports Freed; 'port-allocation-failure': Stream Client Port Allocation Failures; 'unknown-client-port-from-server': Server Replies With Unknown Client Ports; 'data-session-created': Data Session Created; 'data-session-freed': Data Session Freed; 'no-session-mem': Data Session Creation Failures; 'smp-inserted': SMP Session Inserted; 'smp-removed': SMP Session Removed; 'smp-reused': SMP Session Reused; 'fixed-nat-lid-standby': New Session Fixed NAT LID Standby; 'smp-deleted': New Session SMP Already Deleted; 'control-closed': New Session Closed; 'data-session-exists': New Session Already Exists; 'data-session-creation-failure': New Data Session Creation Failure; 'rtp-reversed': RTP Reverse Creation; 'rtcp-reversed': RTCP Reverse Creation; 'cross-cpu-sent': Cross CPU Sent; 'cross-cpu-rcv': Cross CPU Received; 'cross-cpu-no-session': Cross CPU No Session Found; 'cross-cpu-created': Cross CPU Creation; 'cross-cpu-rcv-failure': Cross CPU Receive Failure; 'data-free-smp-retry-lookup': Data Session Free SMP Retry Lookup; 'data-free-smp-not-found': Data Session Free SMP Not Found; 'ha-streams-sent': HA Streams Sent; 'ha-streams-rcv': HA Streams Received; 'ha-stream-incompatible': HA Incompatible Streams Received; 'ha-stream-exists': HA Stream Already Exists; 'ha-port-allocation-failure': HA Stream Port Allocation Failure; 'ha-data-session-sent': HA Data Session Sent; 'ha-data-session-rcv': HA Data Session Received; 'ha-data-no-smp': HA Data Session SMP Not Found; 'ha-control-closed': HA New Data Control Session Closed; 'ha-data-exists': HA New Data Session Already Exists; 'ha-extension-failure': HA Conn Extension Failure; 'ha-stream-smp-reused': HA SMP Session Reused; 'ha-stream-smp-acquire-failure': HA SMP Session Acquire Failure; 'smp-app-type-mismatch': SMP ALG App Type Mismatch;",
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
func resourceCgnv6FixedNatAlgRtspCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgRtspCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgRtsp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatAlgRtspRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6FixedNatAlgRtspUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgRtspUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgRtsp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatAlgRtspRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6FixedNatAlgRtspDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgRtspDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgRtsp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6FixedNatAlgRtspRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgRtspRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgRtsp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6FixedNatAlgRtspSamplingEnable(d []interface{}) []edpt.Cgnv6FixedNatAlgRtspSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatAlgRtspSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatAlgRtspSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatAlgRtsp(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgRtsp {
	var ret edpt.Cgnv6FixedNatAlgRtsp
	ret.Inst.SamplingEnable = getSliceCgnv6FixedNatAlgRtspSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
