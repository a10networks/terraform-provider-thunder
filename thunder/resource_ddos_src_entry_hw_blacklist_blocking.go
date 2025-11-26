package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcEntryHwBlacklistBlocking() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_entry_hw_blacklist_blocking`: Hardware blacklist blocking over threshold\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcEntryHwBlacklistBlockingCreate,
		UpdateContext: resourceDdosSrcEntryHwBlacklistBlockingUpdate,
		ReadContext:   resourceDdosSrcEntryHwBlacklistBlockingRead,
		DeleteContext: resourceDdosSrcEntryHwBlacklistBlockingDelete,

		Schema: map[string]*schema.Schema{
			"src_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Src side hardware blocking",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"src_entry_name": {
				Type: schema.TypeString, Required: true, Description: "SrcEntryName",
			},
		},
	}
}
func resourceDdosSrcEntryHwBlacklistBlockingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryHwBlacklistBlockingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntryHwBlacklistBlocking(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcEntryHwBlacklistBlockingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcEntryHwBlacklistBlockingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryHwBlacklistBlockingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntryHwBlacklistBlocking(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcEntryHwBlacklistBlockingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcEntryHwBlacklistBlockingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryHwBlacklistBlockingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntryHwBlacklistBlocking(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcEntryHwBlacklistBlockingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryHwBlacklistBlockingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntryHwBlacklistBlocking(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosSrcEntryHwBlacklistBlocking(d *schema.ResourceData) edpt.DdosSrcEntryHwBlacklistBlocking {
	var ret edpt.DdosSrcEntryHwBlacklistBlocking
	ret.Inst.SrcEnable = d.Get("src_enable").(int)
	//omit uuid
	ret.Inst.SrcEntryName = d.Get("src_entry_name").(string)
	return ret
}
