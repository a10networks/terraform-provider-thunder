package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DsLiteAlgH323() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_ds_lite_alg_h323`: DS-Lite H323 ALG (default: disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6DsLiteAlgH323Create,
		UpdateContext: resourceCgnv6DsLiteAlgH323Update,
		ReadContext:   resourceCgnv6DsLiteAlgH323Read,
		DeleteContext: resourceCgnv6DsLiteAlgH323Delete,

		Schema: map[string]*schema.Schema{
			"h323_enable": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6DsLiteAlgH323Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgH323Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgH323(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteAlgH323Read(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6DsLiteAlgH323Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgH323Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgH323(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteAlgH323Read(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6DsLiteAlgH323Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgH323Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgH323(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6DsLiteAlgH323Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgH323Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgH323(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6DsLiteAlgH323(d *schema.ResourceData) edpt.Cgnv6DsLiteAlgH323 {
	var ret edpt.Cgnv6DsLiteAlgH323
	ret.Inst.H323Enable = d.Get("h323_enable").(string)
	//omit uuid
	return ret
}
