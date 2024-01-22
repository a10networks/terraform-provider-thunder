package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateConnectionReuse() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_connection_reuse`: Connection Reuse\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateConnectionReuseCreate,
		UpdateContext: resourceSlbTemplateConnectionReuseUpdate,
		ReadContext:   resourceSlbTemplateConnectionReuseRead,
		DeleteContext: resourceSlbTemplateConnectionReuseDelete,

		Schema: map[string]*schema.Schema{
			"add_header": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert HTTP Connection: keep-alive header",
			},
			"keep_alive_conn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Keep a number of server connections open",
			},
			"limit_per_server": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "Max Server Connections allowed (Connections per Server Port (default 1000))",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Connection Reuse Template Name",
			},
			"num_conn_per_port": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Connections per Server Port (default 100)",
			},
			"preopen": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Preopen server connection",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 2400, Description: "Timeout in seconds. Multiple of 60 (default 2400)",
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
func resourceSlbTemplateConnectionReuseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateConnectionReuseCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateConnectionReuse(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateConnectionReuseRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateConnectionReuseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateConnectionReuseUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateConnectionReuse(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateConnectionReuseRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateConnectionReuseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateConnectionReuseDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateConnectionReuse(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateConnectionReuseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateConnectionReuseRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateConnectionReuse(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateConnectionReuse(d *schema.ResourceData) edpt.SlbTemplateConnectionReuse {
	var ret edpt.SlbTemplateConnectionReuse
	ret.Inst.AddHeader = d.Get("add_header").(int)
	ret.Inst.KeepAliveConn = d.Get("keep_alive_conn").(int)
	ret.Inst.LimitPerServer = d.Get("limit_per_server").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NumConnPerPort = d.Get("num_conn_per_port").(int)
	ret.Inst.Preopen = d.Get("preopen").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
