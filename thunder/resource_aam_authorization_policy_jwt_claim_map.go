package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthorizationPolicyJwtClaimMap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authorization_policy_jwt_claim_map`: Map attributes to JWT claims\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthorizationPolicyJwtClaimMapCreate,
		UpdateContext: resourceAamAuthorizationPolicyJwtClaimMapUpdate,
		ReadContext:   resourceAamAuthorizationPolicyJwtClaimMapRead,
		DeleteContext: resourceAamAuthorizationPolicyJwtClaimMapDelete,

		Schema: map[string]*schema.Schema{
			"attr_num": {
				Type: schema.TypeInt, Required: true, Description: "Spcify attribute ID for claim mapping",
			},
			"bool_val": {
				Type: schema.TypeString, Optional: true, Description: "'true': True; 'false': False;",
			},
			"boolean_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Claim type is boolean",
			},
			"claim": {
				Type: schema.TypeString, Optional: true, Description: "Specify JWT claim name to map to.",
			},
			"num_val": {
				Type: schema.TypeInt, Optional: true, Description: "Specify JWT claim value.",
			},
			"number_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Claim type is number",
			},
			"str_val": {
				Type: schema.TypeString, Optional: true, Description: "Specify JWT claim value.",
			},
			"string_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Claim type is string",
			},
			"type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify claim type",
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
func resourceAamAuthorizationPolicyJwtClaimMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyJwtClaimMapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicyJwtClaimMap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthorizationPolicyJwtClaimMapRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthorizationPolicyJwtClaimMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyJwtClaimMapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicyJwtClaimMap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthorizationPolicyJwtClaimMapRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthorizationPolicyJwtClaimMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyJwtClaimMapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicyJwtClaimMap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthorizationPolicyJwtClaimMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyJwtClaimMapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicyJwtClaimMap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAamAuthorizationPolicyJwtClaimMap(d *schema.ResourceData) edpt.AamAuthorizationPolicyJwtClaimMap {
	var ret edpt.AamAuthorizationPolicyJwtClaimMap
	ret.Inst.AttrNum = d.Get("attr_num").(int)
	ret.Inst.BoolVal = d.Get("bool_val").(string)
	ret.Inst.BooleanType = d.Get("boolean_type").(int)
	ret.Inst.Claim = d.Get("claim").(string)
	ret.Inst.NumVal = d.Get("num_val").(int)
	ret.Inst.NumberType = d.Get("number_type").(int)
	ret.Inst.StrVal = d.Get("str_val").(string)
	ret.Inst.StringType = d.Get("string_type").(int)
	ret.Inst.Type = d.Get("type").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
