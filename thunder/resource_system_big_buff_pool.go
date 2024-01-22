package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemBigBuffPool() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_big_buff_pool`: System big buff pool config\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemBigBuffPoolCreate,
		UpdateContext: resourceSystemBigBuffPoolUpdate,
		ReadContext:   resourceSystemBigBuffPoolRead,
		DeleteContext: resourceSystemBigBuffPoolDelete,

		Schema: map[string]*schema.Schema{
			"big_buff_pool": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure big I/O buffer pool",
			},
		},
	}
}
func resourceSystemBigBuffPoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBigBuffPoolCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBigBuffPool(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemBigBuffPoolRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemBigBuffPoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBigBuffPoolUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBigBuffPool(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemBigBuffPoolRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemBigBuffPoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBigBuffPoolDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBigBuffPool(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemBigBuffPoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBigBuffPoolRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBigBuffPool(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemBigBuffPool(d *schema.ResourceData) edpt.SystemBigBuffPool {
	var ret edpt.SystemBigBuffPool
	ret.Inst.BigBuffPool = d.Get("big_buff_pool").(int)
	return ret
}
