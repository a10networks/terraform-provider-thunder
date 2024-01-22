package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayMgmtInfo() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_mgmt_info`: Overlay management specific data\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayMgmtInfoCreate,
		UpdateContext: resourceOverlayMgmtInfoUpdate,
		ReadContext:   resourceOverlayMgmtInfoRead,
		DeleteContext: resourceOverlayMgmtInfoDelete,

		Schema: map[string]*schema.Schema{
			"appstring": {
				Type: schema.TypeString, Required: true, Description: "Application string",
			},
			"plugin_name": {
				Type: schema.TypeString, Required: true, Description: "Plugin string",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceOverlayMgmtInfoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayMgmtInfoCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayMgmtInfo(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayMgmtInfoRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayMgmtInfoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayMgmtInfoUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayMgmtInfo(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayMgmtInfoRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayMgmtInfoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayMgmtInfoDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayMgmtInfo(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayMgmtInfoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayMgmtInfoRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayMgmtInfo(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayMgmtInfo(d *schema.ResourceData) edpt.OverlayMgmtInfo {
	var ret edpt.OverlayMgmtInfo
	ret.Inst.Appstring = d.Get("appstring").(string)
	ret.Inst.Plugin_name = d.Get("plugin_name").(string)
	//omit uuid
	return ret
}
