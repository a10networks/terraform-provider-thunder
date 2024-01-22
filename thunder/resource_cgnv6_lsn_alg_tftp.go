package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgTftp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_alg_tftp`: Change LSN TFTP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnAlgTftpCreate,
		UpdateContext: resourceCgnv6LsnAlgTftpUpdate,
		ReadContext:   resourceCgnv6LsnAlgTftpRead,
		DeleteContext: resourceCgnv6LsnAlgTftpDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'session-created': TFTP Client Sessions Created; 'helper-created': TFTP Helper Sessions created; 'helper-freed': TFTP Helper Sessions freed; 'helper-freed-used': TFTP Helper Sessions freed used; 'helper-freed-unused': TFTP Helper Sessions freed unused; 'helper-already-used': TFTP Helper Session already used; 'helper-in-rml': TFTP Helper Session in Remove List;",
						},
					},
				},
			},
			"tftp_value": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable TFTP ALG for LSN;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnAlgTftpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgTftpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgTftp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgTftpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnAlgTftpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgTftpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgTftp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnAlgTftpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnAlgTftpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgTftpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgTftp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnAlgTftpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgTftpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgTftp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnAlgTftpSamplingEnable(d []interface{}) []edpt.Cgnv6LsnAlgTftpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnAlgTftpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnAlgTftpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgTftp(d *schema.ResourceData) edpt.Cgnv6LsnAlgTftp {
	var ret edpt.Cgnv6LsnAlgTftp
	ret.Inst.SamplingEnable = getSliceCgnv6LsnAlgTftpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.TftpValue = d.Get("tftp_value").(string)
	//omit uuid
	return ret
}
