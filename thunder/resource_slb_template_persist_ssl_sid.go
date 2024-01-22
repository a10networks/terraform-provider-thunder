package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePersistSslSid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_persist_ssl_sid`: SSL session ID  persistence\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePersistSslSidCreate,
		UpdateContext: resourceSlbTemplatePersistSslSidUpdate,
		ReadContext:   resourceSlbTemplatePersistSslSidRead,
		DeleteContext: resourceSlbTemplatePersistSslSidDelete,

		Schema: map[string]*schema.Schema{
			"dont_honor_conn_rules": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not observe connection rate rules",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "SSL session ID persistence template name",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Persistence timeout (in minutes)",
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
func resourceSlbTemplatePersistSslSidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistSslSidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistSslSid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePersistSslSidRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePersistSslSidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistSslSidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistSslSid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePersistSslSidRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePersistSslSidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistSslSidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistSslSid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePersistSslSidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePersistSslSidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePersistSslSid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplatePersistSslSid(d *schema.ResourceData) edpt.SlbTemplatePersistSslSid {
	var ret edpt.SlbTemplatePersistSslSid
	ret.Inst.DontHonorConnRules = d.Get("dont_honor_conn_rules").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
