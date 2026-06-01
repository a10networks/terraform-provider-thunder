package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallAlgFtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_alg_ftp`: Configure FTP ALG for NAT stateful firewall (default: enabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallAlgFtpCreate,
		UpdateContext: resourceCgnv6StatefulFirewallAlgFtpUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallAlgFtpRead,
		DeleteContext: resourceCgnv6StatefulFirewallAlgFtpDelete,

		Schema: map[string]*schema.Schema{
			"ftp_value": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable ALG;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'client-port-request': PORT Requests From Client; 'client-eprt-request': EPRT Requests From Client; 'server-pasv-reply': PASV Replies From Server; 'server-epsv-reply': EPSV Replies From Server; 'port-retransmits': PORT Retransmits; 'pasv-retransmits': PASV Retransmits; 'smp-app-type-mismatch': SMP App Type Mismatch; 'retransmit-sanity-check-failure': Retransmit Sanity Check Failure; 'smp-conn-alloc-failure': SMP Helper Conn Alloc Failure; 'port-helper-created': PORT Helper Created; 'pasv-helper-created': PASV Helper Created; 'port-helper-acquire-in-del-q': PORT Helper Acquire In Del Queue; 'port-helper-acquire-already-used': PORT Helper Acquire Already Used; 'pasv-helper-acquire-in-del-q': PASV Helper Acquire In Del Queue; 'pasv-helper-acquire-already-used': PASV Helper Acquire Already Used; 'port-helper-freed-used': PORT Helper Freed Used; 'port-helper-freed-unused': PORT Helper Freed Unused; 'pasv-helper-freed-used': PASV Helper Freed Used; 'pasv-helper-freed-unused': PASV Helper Freed Unused;",
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
func resourceCgnv6StatefulFirewallAlgFtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgFtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgFtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallAlgFtpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallAlgFtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgFtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgFtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallAlgFtpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallAlgFtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgFtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgFtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallAlgFtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgFtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgFtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6StatefulFirewallAlgFtpSamplingEnable(d []interface{}) []edpt.Cgnv6StatefulFirewallAlgFtpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6StatefulFirewallAlgFtpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6StatefulFirewallAlgFtpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallAlgFtp(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallAlgFtp {
	var ret edpt.Cgnv6StatefulFirewallAlgFtp
	ret.Inst.FtpValue = d.Get("ftp_value").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6StatefulFirewallAlgFtpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
