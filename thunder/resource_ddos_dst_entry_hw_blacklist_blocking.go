package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryHwBlacklistBlocking() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_hw_blacklist_blocking`: Hardware blacklist blocking over threshold\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntryHwBlacklistBlockingCreate,
		UpdateContext: resourceDdosDstEntryHwBlacklistBlockingUpdate,
		ReadContext:   resourceDdosDstEntryHwBlacklistBlockingRead,
		DeleteContext: resourceDdosDstEntryHwBlacklistBlockingDelete,

		Schema: map[string]*schema.Schema{
			"dst_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Dst side hardware blocking",
			},
			"src_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Src side hardware blocking",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntryHwBlacklistBlockingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryHwBlacklistBlockingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryHwBlacklistBlocking(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryHwBlacklistBlockingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntryHwBlacklistBlockingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryHwBlacklistBlockingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryHwBlacklistBlocking(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryHwBlacklistBlockingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntryHwBlacklistBlockingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryHwBlacklistBlockingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryHwBlacklistBlocking(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntryHwBlacklistBlockingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryHwBlacklistBlockingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryHwBlacklistBlocking(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstEntryHwBlacklistBlocking(d *schema.ResourceData) edpt.DdosDstEntryHwBlacklistBlocking {
	var ret edpt.DdosDstEntryHwBlacklistBlocking
	ret.Inst.DstEnable = d.Get("dst_enable").(int)
	ret.Inst.SrcEnable = d.Get("src_enable").(int)
	//omit uuid
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
