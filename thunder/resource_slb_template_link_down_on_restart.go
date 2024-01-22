package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateLinkDownOnRestart() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_link_down_on_restart`: All Link event will be treated as Link-down template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateLinkDownOnRestartCreate,
		UpdateContext: resourceSlbTemplateLinkDownOnRestartUpdate,
		ReadContext:   resourceSlbTemplateLinkDownOnRestartRead,
		DeleteContext: resourceSlbTemplateLinkDownOnRestartDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateLinkDownOnRestartCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkDownOnRestartCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkDownOnRestart(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLinkDownOnRestartRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateLinkDownOnRestartUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkDownOnRestartUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkDownOnRestart(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLinkDownOnRestartRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateLinkDownOnRestartDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkDownOnRestartDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkDownOnRestart(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateLinkDownOnRestartRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkDownOnRestartRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkDownOnRestart(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateLinkDownOnRestart(d *schema.ResourceData) edpt.SlbTemplateLinkDownOnRestart {
	var ret edpt.SlbTemplateLinkDownOnRestart
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
