package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateFtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_ftp`: FTP template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateFtpCreate,
		UpdateContext: resourceSlbTemplateFtpUpdate,
		ReadContext:   resourceSlbTemplateFtpRead,
		DeleteContext: resourceSlbTemplateFtpDelete,

		Schema: map[string]*schema.Schema{
			"active_mode_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Non-Standard FTP Active mode port",
			},
			"active_mode_port_val": {
				Type: schema.TypeInt, Optional: true, Description: "Non-Standard FTP Active mode port",
			},
			"any": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow any FTP Active mode port",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "FTP template name",
			},
			"to": {
				Type: schema.TypeInt, Optional: true, Description: "End range of FTP Active mode port",
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
func resourceSlbTemplateFtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateFtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateFtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateFtpRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateFtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateFtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateFtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateFtpRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateFtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateFtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateFtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateFtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateFtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateFtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateFtp(d *schema.ResourceData) edpt.SlbTemplateFtp {
	var ret edpt.SlbTemplateFtp
	ret.Inst.ActiveModePort = d.Get("active_mode_port").(int)
	ret.Inst.ActiveModePortVal = d.Get("active_mode_port_val").(int)
	ret.Inst.Any = d.Get("any").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.To = d.Get("to").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
