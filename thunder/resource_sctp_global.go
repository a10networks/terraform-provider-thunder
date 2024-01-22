package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSctpGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sctp_global`: SCTP Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSctpGlobalCreate,
		UpdateContext: resourceSctpGlobalUpdate,
		ReadContext:   resourceSctpGlobalRead,
		DeleteContext: resourceSctpGlobalDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'sctp-static-nat-session-created': SCTP Static NAT Session Created; 'sctp-static-nat-session-deleted': SCTP Static NAT Session Deleted; 'sctp-fw-session-created': SCTP Firewall Session Created; 'sctp-fw-session-deleted': SCTP Firewall Session Deleted; 'pkt-err-drop': Packet Error Drop; 'bad-csum': Bad Checksum; 'bad-payload-drop': Bad Payload Drop; 'bad-alignment-drop': Bad Alignment Drop; 'oos-pkt-drop': Out-of-state Packet Drop; 'max-multi-home-drop': Maximum Multi-homing IP Addresses Drop; 'multi-home-remove-ip-skip': Multi-homing Remove IP Parameter Skip; 'multi-home-addr-not-found-drop': Multi-homing IP Address Not Found Drop; 'static-nat-cfg-not-found': Static NAT Config Not Found Drop; 'cfg-err-drop': Configuration Error Drop; 'vrrp-standby-drop': NAT Resource VRRP-A Standby Drop; 'invalid-frag-chunk-drop': Invalid Fragmented Chunks Drop; 'disallowed-chunk-filtered': Disallowed Chunk Filtered; 'disallowed-pkt-drop': Disallowed Packet Drop; 'rate-limit-drop': Rate-limit Drop; 'sby-session-created': Standby Session Created; 'sby-session-create-fail': Standby Session Create Failed; 'sby-session-updated': Standby Session Updated; 'sby-session-update-fail': Standby Session Update Failed; 'sby-static-nat-cfg-not-found': Static NAT Config Not Found on Standby; 'sctp-out-of-system-memory': Out of System Memory; 'conn_ext_size_max': Max Conn Extension Size; 'bad-csum-shadow': Bad Checksum Shadow; 'bad-payload-drop-shadow': Bad Packet Payload Drop Shadow; 'bad-alignment-drop-shadow': Bad Packet Alignment Drop Shadow; 'sctp-chunk-type-init': SCTP Chunk Type INIT; 'sctp-chunk-type-init-ack': SCTP Chunk Type INIT-ACK; 'sctp-chunk-type-cookie-echo': SCTP Chunk Type COOKIE-ECHO; 'sctp-chunk-type-cookie-ack': SCTP Chunk Type COOKIE-ACK; 'sctp-chunk-type-sack': SCTP Chunk Type SACK; 'sctp-chunk-type-asconf': SCTP Chunk Type ASCONF; 'sctp-chunk-type-asconf-ack': SCTP Chunk Type ASCONF-ACK; 'sctp-chunk-type-data': SCTP Chunk Type DATA; 'sctp-chunk-type-abort': SCTP Chunk Type ABORT; 'sctp-chunk-type-shutdown': SCTP Chunk Type SHUTDOWN; 'sctp-chunk-type-shutdown-ack': SCTP Chunk Type SHUTDOWN-ACK; 'sctp-chunk-type-shutdown-complete': SCTP Chunk Type SHUTDOWN-COMPLETE; 'sctp-chunk-type-error-op': SCTP Chunk Type ERROR-OP; 'sctp-chunk-type-heartbeat': SCTP Chunk Type HEARTBEAT; 'sctp-chunk-type-heartbeat-ack': SCTP Chunk Type HEARTBEAT-ACK; 'sctp-chunk-type-other': SCTP Chunk Type OTHER; 'sctp-heartbeat-no-tuple': SCTP HEARTBEAT/ACK no tuple found; 'sctp-data-no-tuple': SCTP DATA chunk no tuple found; 'sctp-data-no-ext-match': SCTP DATA no extended match found; 'sctp-chunk-type-init-drop': SCTP Chunk Type INIT drop; 'sctp-chunk-type-init-ack-drop': SCTP Chunk Type INIT-ACK drop; 'sctp-chunk-type-shutdown-complete-drop': SCTP Chunk Type SHUTDOWN-COMPLETE drop; 'sctp-chunk-type-abort-data-drop': SCTP Chunk Type with ABORT and DATA drop; 'sctp-chunk-heart-beat-clubbed': SCTP HEARTBEAT chunk with other chunk; 'sctp-retx-init-ack-drop': SCTP Chunk Type INIT_ACK with retx mismatched vtag drop; 'sctp-route-err-heartbeat-drop': SCTP HEARTBEAT ROUTE lookup failed drop; 'sctp-reroute-failover': SCTP REROUTE lookup for chunks other than HEARTBEAT; 'sctp-route-err-drop': SCTP ROUTE lookup failed for chunks other than HEARTBEAT drop; 'sctp-no-ext-match': SCTP no extended match found; 'sctp-retx-init-ack': SCTP Chunk Type INIT_ACK retransmitted; 'sctp-retx-init-drop': SCTP Retransmitted INIT drop; 'sctp-retx-init': SCTP Retransmitted INIT; 'sctp-asconf-process-drop': SCTP ASCONF process drop; 'sctp-init-vtag-zero-drop': SCTP INIT VTAG ZERO drop; 'pkt-len-err-drop': Invalid Packet Length Drop; 'pkt-chunk-len-err-drop': Invalid Chunk Length Drop; 'pkt-asconf-param-len-err-drop': Invalid Parameter Length Drop;",
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
func resourceSctpGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSctpGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSctpGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSctpGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceSctpGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSctpGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSctpGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSctpGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceSctpGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSctpGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSctpGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSctpGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSctpGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSctpGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSctpGlobalSamplingEnable(d []interface{}) []edpt.SctpGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SctpGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SctpGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSctpGlobal(d *schema.ResourceData) edpt.SctpGlobal {
	var ret edpt.SctpGlobal
	ret.Inst.SamplingEnable = getSliceSctpGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
