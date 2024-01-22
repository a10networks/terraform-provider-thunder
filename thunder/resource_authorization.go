package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAuthorization() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_authorization`: Configuration for command authorization\n\n__PLACEHOLDER__",
		CreateContext: resourceAuthorizationCreate,
		UpdateContext: resourceAuthorizationUpdate,
		ReadContext:   resourceAuthorizationRead,
		DeleteContext: resourceAuthorizationDelete,

		Schema: map[string]*schema.Schema{
			"commands": {
				Type: schema.TypeInt, Optional: true, Description: "Commands level for authorization",
			},
			"debug": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the debug level for authorization (Debug level for command authorization. bitwise OR of the following: 1(common), 2(packet),4(packet detail), 8(md5))",
			},
			"method": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tacplus": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Using TACACS+ protocol",
						},
						"none": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "No command authorization",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAuthorizationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthorizationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthorization(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAuthorizationRead(ctx, d, meta)
	}
	return diags
}

func resourceAuthorizationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthorizationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthorization(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAuthorizationRead(ctx, d, meta)
	}
	return diags
}
func resourceAuthorizationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthorizationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthorization(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAuthorizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthorizationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthorization(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAuthorizationMethod(d []interface{}) edpt.AuthorizationMethod {

	count1 := len(d)
	var ret edpt.AuthorizationMethod
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tacplus = in["tacplus"].(int)
		ret.None = in["none"].(int)
	}
	return ret
}

func dataToEndpointAuthorization(d *schema.ResourceData) edpt.Authorization {
	var ret edpt.Authorization
	ret.Inst.Commands = d.Get("commands").(int)
	ret.Inst.Debug = d.Get("debug").(int)
	ret.Inst.Method = getObjectAuthorizationMethod(d.Get("method").([]interface{}))
	//omit uuid
	return ret
}
