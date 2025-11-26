package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneHwBlacklistBlocking() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_hw_blacklist_blocking`: Hardware blacklist blocking over threshold\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneHwBlacklistBlockingCreate,
		UpdateContext: resourceDdosDstZoneHwBlacklistBlockingUpdate,
		ReadContext:   resourceDdosDstZoneHwBlacklistBlockingRead,
		DeleteContext: resourceDdosDstZoneHwBlacklistBlockingDelete,

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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneHwBlacklistBlockingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneHwBlacklistBlockingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneHwBlacklistBlocking(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneHwBlacklistBlockingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneHwBlacklistBlockingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneHwBlacklistBlockingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneHwBlacklistBlocking(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneHwBlacklistBlockingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneHwBlacklistBlockingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneHwBlacklistBlockingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneHwBlacklistBlocking(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneHwBlacklistBlockingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneHwBlacklistBlockingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneHwBlacklistBlocking(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneHwBlacklistBlocking(d *schema.ResourceData) edpt.DdosDstZoneHwBlacklistBlocking {
	var ret edpt.DdosDstZoneHwBlacklistBlocking
	ret.Inst.DstEnable = d.Get("dst_enable").(int)
	ret.Inst.SrcEnable = d.Get("src_enable").(int)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
