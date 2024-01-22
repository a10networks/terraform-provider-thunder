package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugTableSyncEvent() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_table_sync_event`: Debug table sync events\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugTableSyncEventCreate,
		UpdateContext: resourceDebugTableSyncEventUpdate,
		ReadContext:   resourceDebugTableSyncEventRead,
		DeleteContext: resourceDebugTableSyncEventDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "all Event Information",
			},
			"arp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ARP Event Information",
			},
			"fibv4": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPv4 FIB Event Information",
			},
			"fibv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPv6 FIB Event Information",
			},
			"mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "MAC Event Information",
			},
			"nd6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ND6 Event Information",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugTableSyncEventCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTableSyncEventCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTableSyncEvent(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugTableSyncEventRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugTableSyncEventUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTableSyncEventUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTableSyncEvent(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugTableSyncEventRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugTableSyncEventDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTableSyncEventDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTableSyncEvent(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugTableSyncEventRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTableSyncEventRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTableSyncEvent(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugTableSyncEvent(d *schema.ResourceData) edpt.DebugTableSyncEvent {
	var ret edpt.DebugTableSyncEvent
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Arp = d.Get("arp").(int)
	ret.Inst.Fibv4 = d.Get("fibv4").(int)
	ret.Inst.Fibv6 = d.Get("fibv6").(int)
	ret.Inst.Mac = d.Get("mac").(int)
	ret.Inst.Nd6 = d.Get("nd6").(int)
	//omit uuid
	return ret
}
