package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallAlgTftp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_alg_tftp`: Configure TFTP ALG for NAT stateful firewall (default: enabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallAlgTftpCreate,
		UpdateContext: resourceCgnv6StatefulFirewallAlgTftpUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallAlgTftpRead,
		DeleteContext: resourceCgnv6StatefulFirewallAlgTftpDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'session-created': TFTP Client Sessions created; 'helper-created': TFTP Helper Sessions created; 'helper-freed': TFTP Helper Sessions freed; 'helper-freed-used': TFTP Helper Sessions freed used; 'helper-freed-unused': TFTP Helper Sessions freed unused; 'helper-already-used': TFTP Helper Session already used; 'helper-in-rml': TFTP Helper Session in Remove List;",
						},
					},
				},
			},
			"tftp_value": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6StatefulFirewallAlgTftpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgTftpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgTftp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallAlgTftpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallAlgTftpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgTftpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgTftp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallAlgTftpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallAlgTftpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgTftpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgTftp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallAlgTftpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgTftpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgTftp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6StatefulFirewallAlgTftpSamplingEnable(d []interface{}) []edpt.Cgnv6StatefulFirewallAlgTftpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6StatefulFirewallAlgTftpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6StatefulFirewallAlgTftpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallAlgTftp(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallAlgTftp {
	var ret edpt.Cgnv6StatefulFirewallAlgTftp
	ret.Inst.SamplingEnable = getSliceCgnv6StatefulFirewallAlgTftpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.TftpValue = d.Get("tftp_value").(string)
	//omit uuid
	return ret
}
