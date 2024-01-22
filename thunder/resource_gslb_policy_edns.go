package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyEdns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy_edns`: Use EDNS extension\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyEdnsCreate,
		UpdateContext: resourceGslbPolicyEdnsUpdate,
		ReadContext:   resourceGslbPolicyEdnsRead,
		DeleteContext: resourceGslbPolicyEdnsDelete,

		Schema: map[string]*schema.Schema{
			"client_subnet_geographic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use client subnet for geo-location",
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
func resourceGslbPolicyEdnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyEdnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyEdns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyEdnsRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyEdnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyEdnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyEdns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyEdnsRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyEdnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyEdnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyEdns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyEdnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyEdnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyEdns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbPolicyEdns(d *schema.ResourceData) edpt.GslbPolicyEdns {
	var ret edpt.GslbPolicyEdns
	ret.Inst.ClientSubnetGeographic = d.Get("client_subnet_geographic").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
