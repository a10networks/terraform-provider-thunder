package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyConnectionCountBySite() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy_connection_count_by_site`: Select Site Service-IPs with the lowest aggregated connection-count\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyConnectionCountBySiteCreate,
		UpdateContext: resourceGslbPolicyConnectionCountBySiteUpdate,
		ReadContext:   resourceGslbPolicyConnectionCountBySiteRead,
		DeleteContext: resourceGslbPolicyConnectionCountBySiteDelete,

		Schema: map[string]*schema.Schema{
			"connection_count_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable connection-count-by-site",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"policy_name": {
				Type: schema.TypeString, Required: true, Description: "Policy_name",
			},
		},
	}
}
func resourceGslbPolicyConnectionCountBySiteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyConnectionCountBySiteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyConnectionCountBySite(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyConnectionCountBySiteRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyConnectionCountBySiteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyConnectionCountBySiteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyConnectionCountBySite(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyConnectionCountBySiteRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyConnectionCountBySiteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyConnectionCountBySiteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyConnectionCountBySite(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyConnectionCountBySiteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyConnectionCountBySiteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyConnectionCountBySite(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbPolicyConnectionCountBySite(d *schema.ResourceData) edpt.GslbPolicyConnectionCountBySite {
	var ret edpt.GslbPolicyConnectionCountBySite
	ret.Inst.ConnectionCountEnable = d.Get("connection_count_enable").(int)
	//omit uuid
	ret.Inst.Policy_name = d.Get("policy_name").(string)
	return ret
}
