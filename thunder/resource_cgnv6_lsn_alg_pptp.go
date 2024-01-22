package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgPptp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_alg_pptp`: Change LSN PPTP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnAlgPptpCreate,
		UpdateContext: resourceCgnv6LsnAlgPptpUpdate,
		ReadContext:   resourceCgnv6LsnAlgPptpRead,
		DeleteContext: resourceCgnv6LsnAlgPptpDelete,

		Schema: map[string]*schema.Schema{
			"pptp_value": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable PPTP ALG for LSN;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'calls-established': Calls Established; 'mismatched-pns-call-id': Mismatched PNS Call ID; 'gre-sessions-created': GRE Sessions Created; 'gre-sessions-freed': GRE Sessions Freed; 'no-gre-session-match': No Matching GRE Session; 'smp-sessions-created': SMP Sessions Created; 'smp-sessions-freed': SMP Sessions Freed; 'smp-session-creation-failure': SMP Session Creation Failures; 'extension-creation-failure': Extension Creation Failures; 'ha-sent': HA Info Sent; 'ha-rcv': HA Info Received; 'ha-no-mem': HA Memory Allocation Failure; 'ha-conflict': HA Call ID Conflicts; 'ha-overwrite': HA Call ID Overwrites; 'ha-call-sent': HA Call Sent; 'ha-call-rcv': HA Call Received; 'ha-smp-conflict': HA SMP Conflicts; 'ha-smp-in-del-q': HA SMP Deleted; 'smp-app-type-mismatch': SMP ALG App Type Mismatch; 'quota-inc': Quota Incremented; 'quota-dec': Quota Decremented; 'quota-inc-not-found': Quota Not Found on Increment; 'quota-dec-not-found': Quota Not Found on Decrement; 'call-req-pns-call-id-mismatch': Call ID Mismatch on Call Request; 'call-reply-pns-call-id-mismatch': Call ID Mismatch on Call Reply; 'call-req-retransmit': Call Request Retransmit; 'call-req-new': Call Request New; 'call-req-ext-alloc-failure': Call Request Ext Alloc Failure; 'call-reply-call-id-unknown': Call Reply Unknown Client Call ID; 'call-reply-retransmit': Call Reply Retransmit; 'call-reply-retransmit-wrong-control': Call Reply Retransmit Wrong Control; 'call-reply-retransmit-acquired': Call Reply Retransmit Acquired; 'call-reply-ext-alloc-failure': Call Request Ext Alloc Failure; 'smp-client-call-id-mismatch': SMP Client Call ID Mismatch; 'smp-alloc-failure': SMP Session Alloc Failure; 'gre-conn-creation-failure': GRE Conn Alloc Failure; 'gre-conn-ext-creation-failure': GRE Conn Ext Alloc Failure; 'gre-no-fwd-route': GRE No Fwd Route; 'gre-no-rev-route': GRE No Rev Route; 'gre-no-control-conn': GRE No Control Conn; 'gre-conn-already-exists': GRE Conn Already Exists; 'gre-free-no-ext': GRE Free No Ext; 'gre-free-no-smp': GRE Free No SMP; 'gre-free-smp-app-type-mismatch': GRE Free SMP App Type Mismatch; 'control-freed': Control Session Freed; 'control-free-no-ext': Control Free No Ext; 'control-free-no-smp': Control Free No SMP; 'control-free-smp-app-type-mismatch': Control Free SMP App Type Mismatch;",
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
func resourceCgnv6LsnAlgPptpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgPptpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgPptp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgPptpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnAlgPptpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgPptpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgPptp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgPptpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnAlgPptpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgPptpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgPptp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnAlgPptpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgPptpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgPptp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnAlgPptpSamplingEnable(d []interface{}) []edpt.Cgnv6LsnAlgPptpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnAlgPptpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnAlgPptpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgPptp(d *schema.ResourceData) edpt.Cgnv6LsnAlgPptp {
	var ret edpt.Cgnv6LsnAlgPptp
	ret.Inst.PptpValue = d.Get("pptp_value").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6LsnAlgPptpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
