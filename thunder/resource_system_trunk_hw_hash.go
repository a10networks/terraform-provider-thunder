package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTrunkHwHash() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_trunk_hw_hash`: Set Trunk HW Hash Mode\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemTrunkHwHashCreate,
		UpdateContext: resourceSystemTrunkHwHashUpdate,
		ReadContext:   resourceSystemTrunkHwHashRead,
		DeleteContext: resourceSystemTrunkHwHashDelete,

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
func resourceSystemTrunkHwHashCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkHwHashCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkHwHash(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTrunkHwHashRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemTrunkHwHashUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkHwHashUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkHwHash(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTrunkHwHashRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemTrunkHwHashDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkHwHashDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkHwHash(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemTrunkHwHashRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTrunkHwHashRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTrunkHwHash(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemTrunkHwHash(d *schema.ResourceData) edpt.SystemTrunkHwHash {
	var ret edpt.SystemTrunkHwHash
	ret.Inst.Mode = d.Get("mode").(int)
	//omit uuid
	return ret
}
