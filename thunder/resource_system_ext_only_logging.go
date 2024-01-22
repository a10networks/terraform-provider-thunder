package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemExtOnlyLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ext_only_logging`: Configure external only logging\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemExtOnlyLoggingCreate,
		UpdateContext: resourceSystemExtOnlyLoggingUpdate,
		ReadContext:   resourceSystemExtOnlyLoggingRead,
		DeleteContext: resourceSystemExtOnlyLoggingDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable external only logging for packet driven DDOS logs",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemExtOnlyLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemExtOnlyLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemExtOnlyLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemExtOnlyLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemExtOnlyLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemExtOnlyLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemExtOnlyLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemExtOnlyLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemExtOnlyLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemExtOnlyLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemExtOnlyLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemExtOnlyLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemExtOnlyLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemExtOnlyLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemExtOnlyLogging(d *schema.ResourceData) edpt.SystemExtOnlyLogging {
	var ret edpt.SystemExtOnlyLogging
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
