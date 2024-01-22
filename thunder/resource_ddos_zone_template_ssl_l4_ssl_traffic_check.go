package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateSslL4SslTrafficCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_ssl_l4_ssl_traffic_check`: SSL Traffic Check Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateSslL4SslTrafficCheckCreate,
		UpdateContext: resourceDdosZoneTemplateSslL4SslTrafficCheckUpdate,
		ReadContext:   resourceDdosZoneTemplateSslL4SslTrafficCheckRead,
		DeleteContext: resourceDdosZoneTemplateSslL4SslTrafficCheckDelete,

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
func resourceDdosZoneTemplateSslL4SslTrafficCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSslL4SslTrafficCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSslL4SslTrafficCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateSslL4SslTrafficCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateSslL4SslTrafficCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSslL4SslTrafficCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSslL4SslTrafficCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateSslL4SslTrafficCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateSslL4SslTrafficCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSslL4SslTrafficCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSslL4SslTrafficCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateSslL4SslTrafficCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSslL4SslTrafficCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSslL4SslTrafficCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateSslL4SslTrafficCheck(d *schema.ResourceData) edpt.DdosZoneTemplateSslL4SslTrafficCheck {
	var ret edpt.DdosZoneTemplateSslL4SslTrafficCheck
	ret.Inst.CheckResumedConnection = d.Get("check_resumed_connection").(int)
	ret.Inst.HeaderAction = d.Get("header_action").(string)
	ret.Inst.HeaderInspection = d.Get("header_inspection").(int)
	//omit uuid
	ret.Inst.SslL4TmplName = d.Get("ssl_l4_tmpl_name").(string)
	return ret
}
