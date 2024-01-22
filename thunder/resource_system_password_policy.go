package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_password_policy`: Configure Password Complexity, Passsword Aging, Password history under password policy\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemPasswordPolicyCreate,
		UpdateContext: resourceSystemPasswordPolicyUpdate,
		ReadContext:   resourceSystemPasswordPolicyRead,
		DeleteContext: resourceSystemPasswordPolicyDelete,

		Schema: map[string]*schema.Schema{
			"aging": {
				Type: schema.TypeString, Optional: true, Description: "'Strict': Strict: Max Age-60 Days; 'Medium': Medium: Max Age- 90 Days; 'Simple': Simple: Max Age-120 Days;",
			},
			"complexity": {
				Type: schema.TypeString, Optional: true, Description: "'Strict': Strict: Min length:8, Min Lower Case:2, Min Upper Case:2, Min Numbers:2, Min Special Character:1, CHANGE Min 8 Characters; 'Medium': Medium: Min length:6, Min Lower Case:2, Min Upper Case:2, Min Numbers:1, Min Special Character:1, CHANGE Min 6 Characters; 'Default': Default: Min length:9, Min Lower Case:1, Min Upper Case:1, Min Numbers:1, Min Special Character:1, CHANGE Min 1 Characters; 'Simple': Simple: Min length:4, Min Lower Case:1, Min Upper Case:1, Min Numbers:1, Min Special Character:0, CHANGE Min 4 Characters;",
			},
			"forbid_consecutive_character": {
				Type: schema.TypeString, Optional: true, Default: "0", Description: "'0': Will disable the check; '3': Three consecutive characters on keyboard will not be allowed.; '4': Four consecutive characters on keyboard will not be allowed.; '5': Five consecutive characters on keyboard will not be allowed.;",
			},
			"history": {
				Type: schema.TypeString, Optional: true, Description: "'Strict': Strict: Does not allow upto 5 old passwords; 'Medium': Medium: Does not allow upto 4 old passwords; 'Simple': Simple: Does not allow upto 3 old passwords;",
			},
			"min_pswd_len": {
				Type: schema.TypeInt, Optional: true, Description: "Configure custom password length",
			},
			"repeat_character_check": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Prohibition of consecutive repeated input of the same letter/number, case sensitive; 'disable': Will not check if the password contains repeat characters;",
			},
			"username_check": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Prohibition to set password contains user account, case sensitive; 'disable': Will not check if the password contains user account;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemPasswordPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPasswordPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPasswordPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPasswordPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemPasswordPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPasswordPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPasswordPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPasswordPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemPasswordPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPasswordPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPasswordPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemPasswordPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPasswordPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPasswordPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemPasswordPolicy(d *schema.ResourceData) edpt.SystemPasswordPolicy {
	var ret edpt.SystemPasswordPolicy
	ret.Inst.Aging = d.Get("aging").(string)
	ret.Inst.Complexity = d.Get("complexity").(string)
	ret.Inst.ForbidConsecutiveCharacter = d.Get("forbid_consecutive_character").(string)
	ret.Inst.History = d.Get("history").(string)
	ret.Inst.MinPswdLen = d.Get("min_pswd_len").(int)
	ret.Inst.RepeatCharacterCheck = d.Get("repeat_character_check").(string)
	ret.Inst.UsernameCheck = d.Get("username_check").(string)
	//omit uuid
	return ret
}
