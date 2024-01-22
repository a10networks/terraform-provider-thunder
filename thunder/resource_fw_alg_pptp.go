package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgPptp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_alg_pptp`: Change Firewall PPTP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceFwAlgPptpCreate,
		UpdateContext: resourceFwAlgPptpUpdate,
		ReadContext:   resourceFwAlgPptpRead,
		DeleteContext: resourceFwAlgPptpDelete,

		Schema: map[string]*schema.Schema{
			"default_port_disable": {
				Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable PPTP ALG default port 1723;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'calls-established': Calls Established; 'call-req-pns-call-id-mismatch': Call ID Mismatch on Call Request; 'call-reply-pns-call-id-mismatch': Call ID Mismatch on Call Reply; 'gre-session-created': GRE Session Created; 'gre-session-freed': GRE Session Freed; 'call-req-retransmit': Call Request Retransmit; 'call-req-new': Call Request New; 'call-req-ext-alloc-failure': Call Request Ext Alloc Failure; 'call-reply-call-id-unknown': Call Reply Unknown Client Call ID; 'call-reply-retransmit': Call Reply Retransmit; 'call-reply-ext-ext-alloc-failure': Call Request Ext Alloc Failure; 'smp-app-type-mismatch': SMP App Type Mismatch; 'smp-client-call-id-mismatch': SMP Client Call ID Mismatch; 'smp-sessions-created': SMP Session Created; 'smp-sessions-freed': SMP Session Freed; 'smp-alloc-failure': SMP Session Alloc Failure; 'gre-conn-creation-failure': GRE Conn Alloc Failure; 'gre-conn-ext-creation-failure': GRE Conn Ext Alloc Failure; 'gre-no-fwd-route': GRE No Fwd Route; 'gre-no-rev-route': GRE No Rev Route; 'gre-no-control-conn': GRE No Control Conn; 'gre-conn-already-exists': GRE Conn Already Exists; 'gre-free-no-ext': GRE Free No Ext; 'gre-free-no-smp': GRE Free No SMP; 'gre-free-smp-app-type-mismatch': GRE Free SMP App Type Mismatch; 'control-freed': Control Session Freed; 'control-free-no-ext': Control Free No Ext; 'control-free-no-smp': Control Free No SMP; 'control-free-smp-app-type-mismatch': Control Free SMP App Type Mismatch;",
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
func resourceFwAlgPptpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgPptpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgPptp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgPptpRead(ctx, d, meta)
	}
	return diags
}

func resourceFwAlgPptpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgPptpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgPptp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgPptpRead(ctx, d, meta)
	}
	return diags
}
func resourceFwAlgPptpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgPptpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgPptp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwAlgPptpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgPptpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgPptp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwAlgPptpSamplingEnable(d []interface{}) []edpt.FwAlgPptpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwAlgPptpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgPptpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwAlgPptp(d *schema.ResourceData) edpt.FwAlgPptp {
	var ret edpt.FwAlgPptp
	ret.Inst.DefaultPortDisable = d.Get("default_port_disable").(string)
	ret.Inst.SamplingEnable = getSliceFwAlgPptpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
