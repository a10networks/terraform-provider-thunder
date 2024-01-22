package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterLogFile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_log_file`: Logging to file\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterLogFileCreate,
		UpdateContext: resourceRouterLogFileUpdate,
		ReadContext:   resourceRouterLogFileRead,
		DeleteContext: resourceRouterLogFileDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Logging filename (File name)",
			},
			"per_protocol": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Per protocol",
			},
			"rotate": {
				Type: schema.TypeInt, Optional: true, Description: "Log file rotation (Number of backup files)",
			},
			"size": {
				Type: schema.TypeInt, Optional: true, Description: "Log file maximum size (File size in MBytes)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRouterLogFileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterLogFileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterLogFile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterLogFileRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterLogFileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterLogFileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterLogFile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterLogFileRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterLogFileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterLogFileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterLogFile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterLogFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterLogFileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterLogFile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterLogFile(d *schema.ResourceData) edpt.RouterLogFile {
	var ret edpt.RouterLogFile
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PerProtocol = d.Get("per_protocol").(int)
	ret.Inst.Rotate = d.Get("rotate").(int)
	ret.Inst.Size = d.Get("size").(int)
	//omit uuid
	return ret
}
