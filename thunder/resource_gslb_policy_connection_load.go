package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyConnectionLoad() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy_connection_load`: Select Service-IP with the lowest connection-load\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyConnectionLoadCreate,
		UpdateContext: resourceGslbPolicyConnectionLoadUpdate,
		ReadContext:   resourceGslbPolicyConnectionLoadRead,
		DeleteContext: resourceGslbPolicyConnectionLoadDelete,

		Schema: map[string]*schema.Schema{
			"connection_load_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable connection-load",
			},
			"connection_load_fail_break": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Break when exceed limit",
			},
			"connection_load_interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Interval between two samples, Unit: second (Interval value,default is 5)",
			},
			"connection_load_limit": {
				Type: schema.TypeInt, Optional: true, Description: "The value of the connection-load limit, default is unlimited",
			},
			"connection_load_samples": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify samples for connection-load (Number of samples used to calculate the connection load, default is 5)",
			},
			"limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit of maxinum connection load, default is unlimited",
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
func resourceGslbPolicyConnectionLoadCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyConnectionLoadCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyConnectionLoad(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyConnectionLoadRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyConnectionLoadUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyConnectionLoadUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyConnectionLoad(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyConnectionLoadRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyConnectionLoadDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyConnectionLoadDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyConnectionLoad(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyConnectionLoadRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyConnectionLoadRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyConnectionLoad(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbPolicyConnectionLoad(d *schema.ResourceData) edpt.GslbPolicyConnectionLoad {
	var ret edpt.GslbPolicyConnectionLoad
	ret.Inst.ConnectionLoadEnable = d.Get("connection_load_enable").(int)
	ret.Inst.ConnectionLoadFailBreak = d.Get("connection_load_fail_break").(int)
	ret.Inst.ConnectionLoadInterval = d.Get("connection_load_interval").(int)
	ret.Inst.ConnectionLoadLimit = d.Get("connection_load_limit").(int)
	ret.Inst.ConnectionLoadSamples = d.Get("connection_load_samples").(int)
	ret.Inst.Limit = d.Get("limit").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
