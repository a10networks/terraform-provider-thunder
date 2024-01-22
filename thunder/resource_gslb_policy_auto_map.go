package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyAutoMap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy_auto_map`: Auto Mapping Options\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyAutoMapCreate,
		UpdateContext: resourceGslbPolicyAutoMapUpdate,
		ReadContext:   resourceGslbPolicyAutoMapRead,
		DeleteContext: resourceGslbPolicyAutoMapDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All modules",
			},
			"module_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify Disable Auto Map Module",
			},
			"module_type": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"ttl": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Specify Auto Map TTL (TTL, default is 300)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceGslbPolicyAutoMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyAutoMapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyAutoMap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyAutoMapRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyAutoMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyAutoMapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyAutoMap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyAutoMapRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyAutoMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyAutoMapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyAutoMap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyAutoMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyAutoMapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyAutoMap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbPolicyAutoMap(d *schema.ResourceData) edpt.GslbPolicyAutoMap {
	var ret edpt.GslbPolicyAutoMap
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.ModuleDisable = d.Get("module_disable").(int)
	ret.Inst.ModuleType = d.Get("module_type").(string)
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
