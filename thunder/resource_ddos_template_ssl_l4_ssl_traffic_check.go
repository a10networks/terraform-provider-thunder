package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateSslL4SslTrafficCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_ssl_l4_ssl_traffic_check`: SSL Traffic Check Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateSslL4SslTrafficCheckCreate,
		UpdateContext: resourceDdosTemplateSslL4SslTrafficCheckUpdate,
		ReadContext:   resourceDdosTemplateSslL4SslTrafficCheckRead,
		DeleteContext: resourceDdosTemplateSslL4SslTrafficCheckDelete,

		Schema: map[string]*schema.Schema{
			"check_resumed_connection": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply checks to SSL connections initialized by ACK packets",
			},
			"header_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets with bad ssl header; 'ignore': Forward packets with bad ssl header;",
			},
			"header_inspection": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inspect ssl header",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ssl_l4_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "SslL4TmplName",
			},
		},
	}
}
func resourceDdosTemplateSslL4SslTrafficCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4SslTrafficCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4SslTrafficCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSslL4SslTrafficCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateSslL4SslTrafficCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4SslTrafficCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4SslTrafficCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSslL4SslTrafficCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateSslL4SslTrafficCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4SslTrafficCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4SslTrafficCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateSslL4SslTrafficCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4SslTrafficCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4SslTrafficCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateSslL4SslTrafficCheck(d *schema.ResourceData) edpt.DdosTemplateSslL4SslTrafficCheck {
	var ret edpt.DdosTemplateSslL4SslTrafficCheck
	ret.Inst.CheckResumedConnection = d.Get("check_resumed_connection").(int)
	ret.Inst.HeaderAction = d.Get("header_action").(string)
	ret.Inst.HeaderInspection = d.Get("header_inspection").(int)
	//omit uuid
	ret.Inst.SslL4TmplName = d.Get("ssl_l4_tmpl_name").(string)
	return ret
}
