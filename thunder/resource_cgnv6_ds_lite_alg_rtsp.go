package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DsLiteAlgRtsp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_ds_lite_alg_rtsp`: DS-Lite RTSP ALG (default: disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6DsLiteAlgRtspCreate,
		UpdateContext: resourceCgnv6DsLiteAlgRtspUpdate,
		ReadContext:   resourceCgnv6DsLiteAlgRtspRead,
		DeleteContext: resourceCgnv6DsLiteAlgRtspDelete,

		Schema: map[string]*schema.Schema{
			"rtsp_enable": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable ALG;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6DsLiteAlgRtspCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgRtspCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgRtsp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteAlgRtspRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6DsLiteAlgRtspUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgRtspUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgRtsp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteAlgRtspRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6DsLiteAlgRtspDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgRtspDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgRtsp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6DsLiteAlgRtspRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteAlgRtspRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteAlgRtsp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6DsLiteAlgRtsp(d *schema.ResourceData) edpt.Cgnv6DsLiteAlgRtsp {
	var ret edpt.Cgnv6DsLiteAlgRtsp
	ret.Inst.RtspEnable = d.Get("rtsp_enable").(string)
	//omit uuid
	return ret
}
