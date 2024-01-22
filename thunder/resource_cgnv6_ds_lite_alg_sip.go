package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DsLiteAlgSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_ds_lite_alg_sip`: DS-Lite SIP ALG (default: disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6DsLiteAlgSipCreate,
		UpdateContext: resourceCgnv6DsLiteAlgSipUpdate,
		ReadContext:   resourceCgnv6DsLiteAlgSipRead,
		DeleteContext: resourceCgnv6DsLiteAlgSipDelete,

		Schema: map[string]*schema.Schema{
			"sip_enable": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6DsLiteAlgSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteAlgSipRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6DsLiteAlgSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteAlgSipRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6DsLiteAlgSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6DsLiteAlgSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6DsLiteAlgSip(d *schema.ResourceData) edpt.Cgnv6DsLiteAlgSip {
	var ret edpt.Cgnv6DsLiteAlgSip
	ret.Inst.SipEnable = d.Get("sip_enable").(string)
	//omit uuid
	return ret
}
