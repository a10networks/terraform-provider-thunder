package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugIpv6OspfLsa() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ipv6_ospf_lsa`: OSPFv3 Link State Advertisement\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugIpv6OspfLsaCreate,
		UpdateContext: resourceDebugIpv6OspfLsaUpdate,
		ReadContext:   resourceDebugIpv6OspfLsaRead,
		DeleteContext: resourceDebugIpv6OspfLsaDelete,

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
func resourceDebugIpv6OspfLsaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfLsaCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfLsa(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfLsaRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugIpv6OspfLsaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfLsaUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfLsa(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfLsaRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugIpv6OspfLsaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfLsaDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfLsa(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugIpv6OspfLsaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfLsaRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfLsa(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugIpv6OspfLsa(d *schema.ResourceData) edpt.DebugIpv6OspfLsa {
	var ret edpt.DebugIpv6OspfLsa
	ret.Inst.Flooding = d.Get("flooding").(int)
	ret.Inst.Gererate = d.Get("gererate").(int)
	ret.Inst.Install = d.Get("install").(int)
	ret.Inst.Maxage = d.Get("maxage").(int)
	ret.Inst.Refresh = d.Get("refresh").(int)
	//omit uuid
	return ret
}
