package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateSmpp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_smpp`: SMPP template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateSmppCreate,
		UpdateContext: resourceSlbTemplateSmppUpdate,
		ReadContext:   resourceSlbTemplateSmppRead,
		DeleteContext: resourceSlbTemplateSmppDelete,

		Schema: map[string]*schema.Schema{
			"client_enquire_link": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond client ENQUIRE_LINK packet directly instead of forwarding to server",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "SMPP Template Name",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "Configure the password used to bind",
			},
			"server_enquire_link": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send server ENQUIRE_LINK packet for every persist connection when enable conn-reuse",
			},
			"server_enquire_link_val": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Set interval of keep-alive packet for each persistent connection (second, default is 30)",
			},
			"server_selection_per_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force server selection on every SMPP request when enable conn-reuse",
			},
			"user": {
				Type: schema.TypeString, Optional: true, Description: "Configure the user to bind (The name used to bind)",
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
func resourceSlbTemplateSmppCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSmppCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSmpp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateSmppRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateSmppUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSmppUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSmpp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateSmppRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateSmppDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSmppDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSmpp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateSmppRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSmppRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSmpp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateSmpp(d *schema.ResourceData) edpt.SlbTemplateSmpp {
	var ret edpt.SlbTemplateSmpp
	ret.Inst.ClientEnquireLink = d.Get("client_enquire_link").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.ServerEnquireLink = d.Get("server_enquire_link").(int)
	ret.Inst.ServerEnquireLinkVal = d.Get("server_enquire_link_val").(int)
	ret.Inst.ServerSelectionPerRequest = d.Get("server_selection_per_request").(int)
	ret.Inst.User = d.Get("user").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
