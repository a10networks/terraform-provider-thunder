package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DsLiteAlgPptp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_ds_lite_alg_pptp`: DS-Lite PPTP ALG (default: disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6DsLiteAlgPptpCreate,
		UpdateContext: resourceCgnv6DsLiteAlgPptpUpdate,
		ReadContext:   resourceCgnv6DsLiteAlgPptpRead,
		DeleteContext: resourceCgnv6DsLiteAlgPptpDelete,

		Schema: map[string]*schema.Schema{
			"pptp_enable": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6DsLiteAlgPptpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgPptpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgPptp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteAlgPptpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6DsLiteAlgPptpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgPptpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgPptp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteAlgPptpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6DsLiteAlgPptpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgPptpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgPptp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6DsLiteAlgPptpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgPptpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgPptp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6DsLiteAlgPptp(d *schema.ResourceData) edpt.Cgnv6DsLiteAlgPptp {
	var ret edpt.Cgnv6DsLiteAlgPptp
	ret.Inst.PptpEnable = d.Get("pptp_enable").(string)
	//omit uuid
	return ret
}
