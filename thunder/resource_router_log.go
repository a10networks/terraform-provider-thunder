package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_log`: Router log options\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterLogCreate,
		UpdateContext: resourceRouterLogUpdate,
		ReadContext:   resourceRouterLogRead,
		DeleteContext: resourceRouterLogDelete,

		Schema: map[string]*schema.Schema{
			"file": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"per_protocol": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Per protocol",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Logging filename (File name)",
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
				},
			},
			"log_buffer": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Logging goes to log-buffer",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRouterLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterLogRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterLogRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouterLogFile1296(d []interface{}) edpt.RouterLogFile1296 {

	count1 := len(d)
	var ret edpt.RouterLogFile1296
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PerProtocol = in["per_protocol"].(int)
		ret.Name = in["name"].(string)
		ret.Rotate = in["rotate"].(int)
		ret.Size = in["size"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointRouterLog(d *schema.ResourceData) edpt.RouterLog {
	var ret edpt.RouterLog
	ret.Inst.File = getObjectRouterLogFile1296(d.Get("file").([]interface{}))
	ret.Inst.LogBuffer = d.Get("log_buffer").(int)
	//omit uuid
	return ret
}
