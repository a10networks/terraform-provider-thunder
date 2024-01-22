package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTrunkXauiHwHash() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_trunk_xaui_hw_hash`: Set Trunk XAUI HW Hash Mode\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemTrunkXauiHwHashCreate,
		UpdateContext: resourceSystemTrunkXauiHwHashUpdate,
		ReadContext:   resourceSystemTrunkXauiHwHashRead,
		DeleteContext: resourceSystemTrunkXauiHwHashDelete,

		Schema: map[string]*schema.Schema{
			"mode": {
				Type: schema.TypeInt, Optional: true, Default: 6, Description: "Set HW hash mode, default is 6 (1:dst-mac 2:src-mac 3:src-dst-mac 4:src-ip 5:dst-ip 6:rtag6 7:rtag7)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemTrunkXauiHwHashCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkXauiHwHashCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkXauiHwHash(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTrunkXauiHwHashRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemTrunkXauiHwHashUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkXauiHwHashUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkXauiHwHash(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTrunkXauiHwHashRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemTrunkXauiHwHashDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkXauiHwHashDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkXauiHwHash(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemTrunkXauiHwHashRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkXauiHwHashRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkXauiHwHash(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemTrunkXauiHwHash(d *schema.ResourceData) edpt.SystemTrunkXauiHwHash {
	var ret edpt.SystemTrunkXauiHwHash
	ret.Inst.Mode = d.Get("mode").(int)
	//omit uuid
	return ret
}
