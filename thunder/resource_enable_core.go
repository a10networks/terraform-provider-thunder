package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableCore() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_core`: Enable system coredump switch\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableCoreCreate,
		UpdateContext: resourceEnableCoreUpdate,
		ReadContext:   resourceEnableCoreRead,
		DeleteContext: resourceEnableCoreDelete,

		Schema: map[string]*schema.Schema{
			"core_level": {
				Type: schema.TypeString, Optional: true, Default: "a10", Description: "'a10': Enable A10 core dump, by default; 'system': Enable system coredump;",
			},
			"full": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable full system core dump",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceEnableCoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableCoreCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableCore(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableCoreRead(ctx, d, meta)
	}
	return diags
}

func resourceEnableCoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableCoreUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableCore(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableCoreRead(ctx, d, meta)
	}
	return diags
}
func resourceEnableCoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableCoreDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableCore(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableCoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableCoreRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableCore(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointEnableCore(d *schema.ResourceData) edpt.EnableCore {
	var ret edpt.EnableCore
	ret.Inst.CoreLevel = d.Get("core_level").(string)
	ret.Inst.Full = d.Get("full").(int)
	//omit uuid
	return ret
}
