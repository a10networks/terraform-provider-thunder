package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgFtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_fixed_nat_alg_ftp`: Change Fixed NAT FTP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6FixedNatAlgFtpCreate,
		UpdateContext: resourceCgnv6FixedNatAlgFtpUpdate,
		ReadContext:   resourceCgnv6FixedNatAlgFtpRead,
		DeleteContext: resourceCgnv6FixedNatAlgFtpDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'port-requests': PORT Requests From Client; 'eprt-requests': EPRT Requests From Client; 'lprt-requests': LPRT Requests From Client; 'pasv-replies': PASV Replies From Server; 'epsv-replies': EPSV Replies From Server; 'lpsv-replies': LPSV Replies From Server; 'port-retransmits': Port Mode Request Retransmits; 'pasv-retransmits': Passive Mode Reply Retransmits; 'port-helper-created': Port Mode Helper Created; 'pasv-helper-created': Passive Mode Helper Created; 'port-helper-freed': Port Mode Helper Freed; 'pasv-helper-freed': Passive Mode Helper Freed; 'port-helper-unused': Port Mode Helper Unused; 'pasv-helper-unused': Passive Mode Helper Unused; 'port-helper-creation-failure': Port Helper Creation Failure; 'pasv-helper-creation-failure': Passive Helper Creation Failure; 'get-conn-ext-failure': Get Conn Extension Failure; 'smp-app-type-mismatch': SMP ALG App Type Mismatch;",
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
func resourceCgnv6FixedNatAlgFtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgFtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgFtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatAlgFtpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6FixedNatAlgFtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgFtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgFtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatAlgFtpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6FixedNatAlgFtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgFtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgFtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6FixedNatAlgFtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgFtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgFtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6FixedNatAlgFtpSamplingEnable(d []interface{}) []edpt.Cgnv6FixedNatAlgFtpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatAlgFtpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatAlgFtpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatAlgFtp(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgFtp {
	var ret edpt.Cgnv6FixedNatAlgFtp
	ret.Inst.SamplingEnable = getSliceCgnv6FixedNatAlgFtpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
