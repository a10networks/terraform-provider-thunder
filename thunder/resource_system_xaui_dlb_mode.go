package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemXauiDlbMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_xaui_dlb_mode`: Enable/Disable Dynamic Load Balancing traffic distribution support\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemXauiDlbModeCreate,
		UpdateContext: resourceSystemXauiDlbModeUpdate,
		ReadContext:   resourceSystemXauiDlbModeRead,
		DeleteContext: resourceSystemXauiDlbModeDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable Dynamic Load Balancing traffic distribution support",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemXauiDlbModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemXauiDlbModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemXauiDlbMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemXauiDlbModeRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemXauiDlbModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemXauiDlbModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemXauiDlbMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemXauiDlbModeRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemXauiDlbModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemXauiDlbModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemXauiDlbMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemXauiDlbModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemXauiDlbModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemXauiDlbMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemXauiDlbMode(d *schema.ResourceData) edpt.SystemXauiDlbMode {
	var ret edpt.SystemXauiDlbMode
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
