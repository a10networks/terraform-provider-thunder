package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateImapPop3() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_imap_pop3`: IMAP\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateImapPop3Create,
		UpdateContext: resourceSlbTemplateImapPop3Update,
		ReadContext:   resourceSlbTemplateImapPop3Read,
		DeleteContext: resourceSlbTemplateImapPop3Delete,

		Schema: map[string]*schema.Schema{
			"logindisabled": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Login before STARTTLS.Works only for imap",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "IMAP-POP3 Template Name",
			},
			"starttls": {
				Type: schema.TypeString, Optional: true, Default: "disabled", Description: "'disabled': Disable STARTTLS; 'optional': STARTTLS is optional requirement; 'enforced': Must issue STARTTLS command before imap transaction;",
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
func resourceSlbTemplateImapPop3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateImapPop3Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateImapPop3(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateImapPop3Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateImapPop3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateImapPop3Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateImapPop3(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateImapPop3Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateImapPop3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateImapPop3Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateImapPop3(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateImapPop3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateImapPop3Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateImapPop3(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateImapPop3(d *schema.ResourceData) edpt.SlbTemplateImapPop3 {
	var ret edpt.SlbTemplateImapPop3
	ret.Inst.Logindisabled = d.Get("logindisabled").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Starttls = d.Get("starttls").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
