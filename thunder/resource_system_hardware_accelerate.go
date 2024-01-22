package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHardwareAccelerate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_hardware_accelerate`: Enable hardware acceleration\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemHardwareAccelerateCreate,
		UpdateContext: resourceSystemHardwareAccelerateUpdate,
		ReadContext:   resourceSystemHardwareAccelerateRead,
		DeleteContext: resourceSystemHardwareAccelerateDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hit-counts': Total packets hit counts; 'hit-index': HW Fwd hit index; 'ipv4-forward-counts': Total IPv4 hardware forwarded packets; 'ipv6-forward-counts': Total IPv6 hardware forwarded packets; 'hw-fwd-module-status': Hardware forwarder status flags; 'hw-fwd-prog-reqs': Hardware forward programming requests; 'hw-fwd-prog-errors': Hardware forward programming Errors; 'hw-fwd-flow-singlebit-errors': Hardware forward singlebit Errors; 'hw-fwd-flow-tag-mismatch': Hardware forward tag mismatch errors; 'hw-fwd-flow-seq-mismatch': Hardware forward sequence mismatch errors; 'hw-fwd-ageout-drop-count': Hardware forward ageout drop count; 'hw-fwd-invalidation-drop': Hardware forward invalid drop count; 'hw-fwd-flow-hit-index': Hardware forward flow hit index; 'hw-fwd-flow-reason-flags': Hardware forward flow reason flags; 'hw-fwd-flow-drop-count': Hardware forward flow drop count; 'hw-fwd-flow-error-count': Hardware forward flow error count; 'hw-fwd-flow-unalign-count': Hardware forward flow unalign count; 'hw-fwd-flow-underflow-count': Hardware forward flow underflow count; 'hw-fwd-flow-tx-full-drop': Hardware forward flow tx full drop count; 'hw-fwd-flow-qdr-full-drop': Hardware forward flow qdr full drop count; 'hw-fwd-phyport-mismatch-drop': Hardware forward phyport mismatch count; 'hw-fwd-vlanid-mismatch-drop': Hardware forward vlanid mismatch count; 'hw-fwd-vmid-drop': Hardware forward vmid mismatch count; 'hw-fwd-protocol-mismatch-drop': Hardware forward protocol mismatch count; 'hw-fwd-avail-ipv4-entry': Hardware forward available ipv4 entries count; 'hw-fwd-avail-ipv6-entry': Hardware forward available ipv6 entries count; 'hw-fwd-rate-drop-count': Hardware forward rate drop count; 'hw-fwd-normal-ageout-rcvd': Hardware forward normal ageout received count; 'hw-fwd-tcp-fin-ageout-rcvd': Hardware forward tcp FIN ageout received count; 'hw-fwd-tcp-rst-ageout-rcvd': Hardware forward tcp RST ageout received count; 'hw-fwd-lookup-fail-rcvd': Hardware forward entry lookup fail count; 'hw-fwd-stats-update-rcvd': Hardware forward entry stats update count; 'hw-fwd-flow-sflow-count': hardware forward rate drop count;",
						},
					},
				},
			},
			"session_forwarding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure session-forwarding in Hardware (default:off)",
			},
			"slb": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'entry-create': Hardware Entries Created; 'entry-create-failure': Hardware Entries Created; 'entry-create-fail-server-down': Hardware Entries Created; 'entry-create-fail-max-entry': Hardware Entries Created; 'entry-free': Hardware Entries Freed; 'entry-free-opp-entry': Hardware Entries Free due to opposite tuple entry ageout event; 'entry-free-no-hw-prog': Hardware Entry Free no hw prog; 'entry-free-no-conn': Hardware Entry Free no matched conn; 'entry-free-no-sw-entry': Hardware Entry Free no software entry; 'entry-counter': Hardware Entry Count; 'entry-age-out': Hardware Entries Aged Out; 'entry-age-out-idle': Hardware Entries Aged Out Idle; 'entry-age-out-tcp-fin': Hardware Entries Aged Out TCP FIN; 'entry-age-out-tcp-rst': Hardware Entries Aged Out TCP RST; 'entry-age-out-invalid-dst': Hardware Entries Aged Out invalid dst; 'entry-force-hw-invalidate': Hardware Entries Force HW Invalidate; 'entry-invalidate-server-down': Hardware Entries Invalidate due to server down; 'tcam-create': TCAM Entries Created; 'tcam-free': TCAM Entries Freed; 'tcam-counter': TCAM Entry Count;",
									},
								},
							},
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
func resourceSystemHardwareAccelerateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareAccelerateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareAccelerate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemHardwareAccelerateRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemHardwareAccelerateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareAccelerateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareAccelerate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemHardwareAccelerateRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemHardwareAccelerateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareAccelerateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareAccelerate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemHardwareAccelerateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareAccelerateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareAccelerate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemHardwareAccelerateSamplingEnable(d []interface{}) []edpt.SystemHardwareAccelerateSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemHardwareAccelerateSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareAccelerateSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemHardwareAccelerateSlb1574(d []interface{}) edpt.SystemHardwareAccelerateSlb1574 {

	count1 := len(d)
	var ret edpt.SystemHardwareAccelerateSlb1574
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceSystemHardwareAccelerateSlbSamplingEnable1575(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSystemHardwareAccelerateSlbSamplingEnable1575(d []interface{}) []edpt.SystemHardwareAccelerateSlbSamplingEnable1575 {

	count1 := len(d)
	ret := make([]edpt.SystemHardwareAccelerateSlbSamplingEnable1575, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareAccelerateSlbSamplingEnable1575
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemHardwareAccelerate(d *schema.ResourceData) edpt.SystemHardwareAccelerate {
	var ret edpt.SystemHardwareAccelerate
	ret.Inst.SamplingEnable = getSliceSystemHardwareAccelerateSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SessionForwarding = d.Get("session_forwarding").(int)
	ret.Inst.Slb = getObjectSystemHardwareAccelerateSlb1574(d.Get("slb").([]interface{}))
	//omit uuid
	return ret
}
