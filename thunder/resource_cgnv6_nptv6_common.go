package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nptv6Common() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nptv6_common`: NPTv6 common settings\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nptv6CommonCreate,
		UpdateContext: resourceCgnv6Nptv6CommonUpdate,
		ReadContext:   resourceCgnv6Nptv6CommonRead,
		DeleteContext: resourceCgnv6Nptv6CommonDelete,

		Schema: map[string]*schema.Schema{
			"send_icmpv6_on_error": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable to send ICMPv6 when error discovered;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Nptv6CommonCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nptv6CommonCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nptv6Common(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nptv6CommonRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nptv6CommonUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nptv6CommonUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nptv6Common(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nptv6CommonRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nptv6CommonDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nptv6CommonDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nptv6Common(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nptv6CommonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nptv6CommonRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nptv6Common(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nptv6Common(d *schema.ResourceData) edpt.Cgnv6Nptv6Common {
	var ret edpt.Cgnv6Nptv6Common
	ret.Inst.SendIcmpv6OnError = d.Get("send_icmpv6_on_error").(string)
	//omit uuid
	return ret
}
