package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSynCookie() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_syn_cookie`: Global Syn-Cookie Protection\n\n__PLACEHOLDER__",
		CreateContext: resourceSynCookieCreate,
		UpdateContext: resourceSynCookieUpdate,
		ReadContext:   resourceSynCookieRead,
		DeleteContext: resourceSynCookieDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Global Syn-Cookie Protection",
			},
			"off_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "off-threshold for Syn-cookie. (default 0) (Decimal number)",
			},
			"on_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "on-threshold for Syn-cookie. (default 0) (Decimal number, 0 for turning on Syn-cookie without threshold)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSynCookieCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSynCookieCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSynCookie(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSynCookieRead(ctx, d, meta)
	}
	return diags
}

func resourceSynCookieUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSynCookieUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSynCookie(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSynCookieRead(ctx, d, meta)
	}
	return diags
}
func resourceSynCookieDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSynCookieDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSynCookie(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSynCookieRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSynCookieRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSynCookie(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSynCookie(d *schema.ResourceData) edpt.SynCookie {
	var ret edpt.SynCookie
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.OffThreshold = d.Get("off_threshold").(int)
	ret.Inst.OnThreshold = d.Get("on_threshold").(int)
	//omit uuid
	return ret
}
