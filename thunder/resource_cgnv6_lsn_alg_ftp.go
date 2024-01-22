package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgFtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_alg_ftp`: Change LSN FTP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnAlgFtpCreate,
		UpdateContext: resourceCgnv6LsnAlgFtpUpdate,
		ReadContext:   resourceCgnv6LsnAlgFtpRead,
		DeleteContext: resourceCgnv6LsnAlgFtpDelete,

		Schema: map[string]*schema.Schema{
			"ftp_value": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable FTP ALG for LSN;",
			},
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
func resourceCgnv6LsnAlgFtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgFtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgFtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgFtpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnAlgFtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgFtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgFtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgFtpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnAlgFtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgFtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgFtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnAlgFtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgFtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgFtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnAlgFtpSamplingEnable(d []interface{}) []edpt.Cgnv6LsnAlgFtpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnAlgFtpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnAlgFtpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgFtp(d *schema.ResourceData) edpt.Cgnv6LsnAlgFtp {
	var ret edpt.Cgnv6LsnAlgFtp
	ret.Inst.FtpValue = d.Get("ftp_value").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6LsnAlgFtpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
