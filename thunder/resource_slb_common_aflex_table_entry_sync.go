package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommonAflexTableEntrySync() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_common_aflex_table_entry_sync`: aflex table entry sync commands\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbCommonAflexTableEntrySyncCreate,
		UpdateContext: resourceSlbCommonAflexTableEntrySyncUpdate,
		ReadContext:   resourceSlbCommonAflexTableEntrySyncRead,
		DeleteContext: resourceSlbCommonAflexTableEntrySyncDelete,

		Schema: map[string]*schema.Schema{
			"aflex_table_entry_sync_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable aflex table sync",
			},
			"aflex_table_entry_sync_max_key_len": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "aflex table entry max key length to sync",
			},
			"aflex_table_entry_sync_max_value_len": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "aflex table entry max value length to sync",
			},
			"aflex_table_entry_sync_min_lifetime": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "aflex table entry minimum lifetime to sync",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbCommonAflexTableEntrySyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonAflexTableEntrySyncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonAflexTableEntrySync(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonAflexTableEntrySyncRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCommonAflexTableEntrySyncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonAflexTableEntrySyncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonAflexTableEntrySync(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonAflexTableEntrySyncRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbCommonAflexTableEntrySyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonAflexTableEntrySyncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonAflexTableEntrySync(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCommonAflexTableEntrySyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonAflexTableEntrySyncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonAflexTableEntrySync(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbCommonAflexTableEntrySync(d *schema.ResourceData) edpt.SlbCommonAflexTableEntrySync {
	var ret edpt.SlbCommonAflexTableEntrySync
	ret.Inst.AflexTableEntrySyncEnable = d.Get("aflex_table_entry_sync_enable").(int)
	ret.Inst.AflexTableEntrySyncMaxKeyLen = d.Get("aflex_table_entry_sync_max_key_len").(int)
	ret.Inst.AflexTableEntrySyncMaxValueLen = d.Get("aflex_table_entry_sync_max_value_len").(int)
	ret.Inst.AflexTableEntrySyncMinLifetime = d.Get("aflex_table_entry_sync_min_lifetime").(int)
	//omit uuid
	return ret
}
