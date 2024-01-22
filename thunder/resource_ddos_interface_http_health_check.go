package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosInterfaceHttpHealthCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_interface_http_health_check`: Use DST Interface-ip entry as HTTP health-check target\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosInterfaceHttpHealthCheckCreate,
		UpdateContext: resourceDdosInterfaceHttpHealthCheckUpdate,
		ReadContext:   resourceDdosInterfaceHttpHealthCheckRead,
		DeleteContext: resourceDdosInterfaceHttpHealthCheckDelete,

		Schema: map[string]*schema.Schema{
			"challenge_method": {
				Type: schema.TypeString, Optional: true, Description: "'http-redirect': http-redirect; 'javascript': javascript;",
			},
			"challenge_redirect_code": {
				Type: schema.TypeString, Optional: true, Default: "302", Description: "'302': 302 Found; '307': 307 Temporary Redirect;",
			},
			"challenge_uri_encode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Encode the challenge phrase in uri instead of in http cookie. Default encoded in http cookie",
			},
			"enable": {
				Type: schema.TypeString, Required: true, Description: "'enable': enable;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosInterfaceHttpHealthCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosInterfaceHttpHealthCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosInterfaceHttpHealthCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosInterfaceHttpHealthCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosInterfaceHttpHealthCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosInterfaceHttpHealthCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosInterfaceHttpHealthCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosInterfaceHttpHealthCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosInterfaceHttpHealthCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosInterfaceHttpHealthCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosInterfaceHttpHealthCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosInterfaceHttpHealthCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosInterfaceHttpHealthCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosInterfaceHttpHealthCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosInterfaceHttpHealthCheck(d *schema.ResourceData) edpt.DdosInterfaceHttpHealthCheck {
	var ret edpt.DdosInterfaceHttpHealthCheck
	ret.Inst.ChallengeMethod = d.Get("challenge_method").(string)
	ret.Inst.ChallengeRedirectCode = d.Get("challenge_redirect_code").(string)
	ret.Inst.ChallengeUriEncode = d.Get("challenge_uri_encode").(int)
	ret.Inst.Enable = d.Get("enable").(string)
	//omit uuid
	return ret
}
