package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64AlgRtsp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat64_alg_rtsp`: NAT64 RTSP ALG (default: disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat64AlgRtspCreate,
		UpdateContext: resourceCgnv6Nat64AlgRtspUpdate,
		ReadContext:   resourceCgnv6Nat64AlgRtspRead,
		DeleteContext: resourceCgnv6Nat64AlgRtspDelete,

		Schema: map[string]*schema.Schema{
			"rtsp_enable": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable NAT64 RTSP ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Nat64AlgRtspCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgRtspCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgRtsp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgRtspRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat64AlgRtspUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgRtspUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgRtsp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64AlgRtspRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat64AlgRtspDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgRtspDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgRtsp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat64AlgRtspRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgRtspRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgRtsp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat64AlgRtsp(d *schema.ResourceData) edpt.Cgnv6Nat64AlgRtsp {
	var ret edpt.Cgnv6Nat64AlgRtsp
	ret.Inst.RtspEnable = d.Get("rtsp_enable").(string)
	//omit uuid
	return ret
}
