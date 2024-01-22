package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMultiQueueSupport() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_multi_queue_support`: To Enable/disable Multi-Queue-Support\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMultiQueueSupportCreate,
		UpdateContext: resourceSystemMultiQueueSupportUpdate,
		ReadContext:   resourceSystemMultiQueueSupportRead,
		DeleteContext: resourceSystemMultiQueueSupportDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Multi-Queue-Support",
			},
		},
	}
}
func resourceSystemMultiQueueSupportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMultiQueueSupportCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMultiQueueSupport(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMultiQueueSupportRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMultiQueueSupportUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMultiQueueSupportUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMultiQueueSupport(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMultiQueueSupportRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMultiQueueSupportDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMultiQueueSupportDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMultiQueueSupport(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMultiQueueSupportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMultiQueueSupportRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMultiQueueSupport(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemMultiQueueSupport(d *schema.ResourceData) edpt.SystemMultiQueueSupport {
	var ret edpt.SystemMultiQueueSupport
	ret.Inst.Enable = d.Get("enable").(int)
	return ret
}
