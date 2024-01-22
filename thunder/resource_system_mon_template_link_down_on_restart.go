package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMonTemplateLinkDownOnRestart() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_mon_template_link_down_on_restart`: All Link event will be treated as Link-down template\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMonTemplateLinkDownOnRestartCreate,
		UpdateContext: resourceSystemMonTemplateLinkDownOnRestartUpdate,
		ReadContext:   resourceSystemMonTemplateLinkDownOnRestartRead,
		DeleteContext: resourceSystemMonTemplateLinkDownOnRestartDelete,

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
func resourceSystemMonTemplateLinkDownOnRestartCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateLinkDownOnRestartCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateLinkDownOnRestart(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMonTemplateLinkDownOnRestartRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMonTemplateLinkDownOnRestartUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateLinkDownOnRestartUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateLinkDownOnRestart(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMonTemplateLinkDownOnRestartRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMonTemplateLinkDownOnRestartDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateLinkDownOnRestartDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateLinkDownOnRestart(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMonTemplateLinkDownOnRestartRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateLinkDownOnRestartRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateLinkDownOnRestart(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemMonTemplateLinkDownOnRestart(d *schema.ResourceData) edpt.SystemMonTemplateLinkDownOnRestart {
	var ret edpt.SystemMonTemplateLinkDownOnRestart
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
