package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_alg_sip`: Change Firewall SIP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceFwAlgSipCreate,
		UpdateContext: resourceFwAlgSipUpdate,
		ReadContext:   resourceFwAlgSipRead,
		DeleteContext: resourceFwAlgSipDelete,

		Schema: map[string]*schema.Schema{
			"default_port_disable": {
				Type: schema.TypeString, Optional: true, Description: "'default-port-disable': Disable SIP ALG default port 5060;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'stat-request': Request Received; 'stat-response': Response Received; 'method-register': Method REGISTER; 'method-invite': Method INVITE; 'method-ack': Method ACK; 'method-cancel': Method CANCEL; 'method-bye': Method BYE; 'method-options': Method OPTIONS; 'method-prack': Method PRACK; 'method-subscribe': Method SUBSCRIBE; 'method-notify': Method NOTIFY; 'method-publish': Method PUBLISH; 'method-info': Method INFO; 'method-refer': Method REFER; 'method-message': Method MESSAGE; 'method-update': Method UPDATE; 'method-unknown': Method Unknown; 'parse-error': Message Parse Error; 'keep-alive': Keep Alive; 'contact-error': Contact Process Error; 'sdp-error': SDP Process Error; 'rtp-port-no-op': RTP Port No Op; 'rtp-rtcp-port-success': RTP RTCP Port Success; 'rtp-port-failure': RTP Port Failure; 'rtcp-port-failure': RTCP Port Failure; 'contact-port-no-op': Contact Port No Op; 'contact-port-success': Contact Port Success; 'contact-port-failure': Contact Port Failure; 'contact-new': Contact Alloc; 'contact-alloc-failure': Contact Alloc Failure; 'contact-eim': Contact EIM; 'contact-eim-set': Contact EIM Set; 'rtp-new': RTP Alloc; 'rtp-alloc-failure': RTP Alloc Failure; 'rtp-eim': RTP EIM; 'helper-found': SMP Helper Conn Found; 'helper-created': SMP Helper Conn Created; 'helper-deleted': SMP Helper Conn Already Deleted; 'helper-freed': SMP Helper Conn Freed; 'helper-failure': SMP Helper Failure;",
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
func resourceFwAlgSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgSipRead(ctx, d, meta)
	}
	return diags
}

func resourceFwAlgSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAlgSipRead(ctx, d, meta)
	}
	return diags
}
func resourceFwAlgSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwAlgSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwAlgSipSamplingEnable(d []interface{}) []edpt.FwAlgSipSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwAlgSipSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAlgSipSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwAlgSip(d *schema.ResourceData) edpt.FwAlgSip {
	var ret edpt.FwAlgSip
	ret.Inst.DefaultPortDisable = d.Get("default_port_disable").(string)
	ret.Inst.SamplingEnable = getSliceFwAlgSipSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
