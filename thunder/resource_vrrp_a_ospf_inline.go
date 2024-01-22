package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAOspfInline() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_ospf_inline`: Enable OSPF under Layer 3 Inline Hot Standby Mode\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAOspfInlineCreate,
		UpdateContext: resourceVrrpAOspfInlineUpdate,
		ReadContext:   resourceVrrpAOspfInlineRead,
		DeleteContext: resourceVrrpAOspfInlineDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan": {
				Type: schema.TypeInt, Required: true, Description: "Do not filter OSPF packet on specific vlan (Vlan number)",
			},
		},
	}
}
func resourceVrrpAOspfInlineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAOspfInlineCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAOspfInline(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAOspfInlineRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAOspfInlineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAOspfInlineUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAOspfInline(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAOspfInlineRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAOspfInlineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAOspfInlineDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAOspfInline(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAOspfInlineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAOspfInlineRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAOspfInline(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVrrpAOspfInline(d *schema.ResourceData) edpt.VrrpAOspfInline {
	var ret edpt.VrrpAOspfInline
	//omit uuid
	ret.Inst.Vlan = d.Get("vlan").(int)
	return ret
}
