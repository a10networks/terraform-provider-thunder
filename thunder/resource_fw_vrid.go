package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwVrid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_vrid`: Vrrp group ID\n\n__PLACEHOLDER__",
		CreateContext: resourceFwVridCreate,
		UpdateContext: resourceFwVridUpdate,
		ReadContext:   resourceFwVridRead,
		DeleteContext: resourceFwVridDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Vrrp group (VRRP-A vrid)",
			},
		},
	}
}
func resourceFwVridCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwVridCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwVrid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwVridRead(ctx, d, meta)
	}
	return diags
}

func resourceFwVridUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwVridUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwVrid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwVridRead(ctx, d, meta)
	}
	return diags
}
func resourceFwVridDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwVridDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwVrid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwVridRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwVridRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwVrid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwVrid(d *schema.ResourceData) edpt.FwVrid {
	var ret edpt.FwVrid
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
