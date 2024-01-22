package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugOspfLsa() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ospf_lsa`: OSPFv3 Link State Advertisement\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugOspfLsaCreate,
		UpdateContext: resourceDebugOspfLsaUpdate,
		ReadContext:   resourceDebugOspfLsaRead,
		DeleteContext: resourceDebugOspfLsaDelete,

		Schema: map[string]*schema.Schema{
			"flooding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "LSA Flooding",
			},
			"gererate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "LSA Generation",
			},
			"install": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "LSA Installation",
			},
			"maxage": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "LSA MaxAge processing",
			},
			"refresh": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "LSA Refreshment",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugOspfLsaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfLsaCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfLsa(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfLsaRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugOspfLsaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfLsaUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfLsa(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfLsaRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugOspfLsaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfLsaDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfLsa(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugOspfLsaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfLsaRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfLsa(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugOspfLsa(d *schema.ResourceData) edpt.DebugOspfLsa {
	var ret edpt.DebugOspfLsa
	ret.Inst.Flooding = d.Get("flooding").(int)
	ret.Inst.Gererate = d.Get("gererate").(int)
	ret.Inst.Install = d.Get("install").(int)
	ret.Inst.Maxage = d.Get("maxage").(int)
	ret.Inst.Refresh = d.Get("refresh").(int)
	//omit uuid
	return ret
}
