package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgTftp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_fixed_nat_alg_tftp`: Change Fixed NAT TFTP ALG Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6FixedNatAlgTftpCreate,
		UpdateContext: resourceCgnv6FixedNatAlgTftpUpdate,
		ReadContext:   resourceCgnv6FixedNatAlgTftpRead,
		DeleteContext: resourceCgnv6FixedNatAlgTftpDelete,

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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6FixedNatAlgTftpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgTftpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgTftp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatAlgTftpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6FixedNatAlgTftpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgTftpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgTftp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatAlgTftpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6FixedNatAlgTftpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgTftpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgTftp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6FixedNatAlgTftpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgTftpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgTftp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6FixedNatAlgTftpSamplingEnable(d []interface{}) []edpt.Cgnv6FixedNatAlgTftpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6FixedNatAlgTftpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6FixedNatAlgTftpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatAlgTftp(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgTftp {
	var ret edpt.Cgnv6FixedNatAlgTftp
	ret.Inst.SamplingEnable = getSliceCgnv6FixedNatAlgTftpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
