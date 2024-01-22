package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbForwardProxyNotifyPage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_forward_proxy_notify_page`: reply a notify page for forward-proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbForwardProxyNotifyPageCreate,
		UpdateContext: resourceSlbForwardProxyNotifyPageUpdate,
		ReadContext:   resourceSlbForwardProxyNotifyPageRead,
		DeleteContext: resourceSlbForwardProxyNotifyPageDelete,

		Schema: map[string]*schema.Schema{
			"message": {
				Type: schema.TypeString, Optional: true, Description: "Message in notify page",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name of this object",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbForwardProxyNotifyPageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbForwardProxyNotifyPageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbForwardProxyNotifyPage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbForwardProxyNotifyPageRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbForwardProxyNotifyPageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbForwardProxyNotifyPageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbForwardProxyNotifyPage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbForwardProxyNotifyPageRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbForwardProxyNotifyPageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbForwardProxyNotifyPageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbForwardProxyNotifyPage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbForwardProxyNotifyPageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbForwardProxyNotifyPageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbForwardProxyNotifyPage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbForwardProxyNotifyPage(d *schema.ResourceData) edpt.SlbForwardProxyNotifyPage {
	var ret edpt.SlbForwardProxyNotifyPage
	ret.Inst.Message = d.Get("message").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
