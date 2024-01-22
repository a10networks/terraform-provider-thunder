package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBanner() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_banner`: Define a login banner\n\n__PLACEHOLDER__",
		CreateContext: resourceBannerCreate,
		UpdateContext: resourceBannerUpdate,
		ReadContext:   resourceBannerRead,
		DeleteContext: resourceBannerDelete,

		Schema: map[string]*schema.Schema{
			"exec_banner_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set EXEC process creation banner",
						},
						"exec_banner": {
							Type: schema.TypeString, Optional: true, Description: "Banner text, string -n is taken as line break of multi-line banner text, use --n for -n, -077 for ? and -011 for tab",
						},
					},
				},
			},
			"login_banner_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"login": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set login banner",
						},
						"login_banner": {
							Type: schema.TypeString, Optional: true, Description: "Banner text, string -n is taken as line break of multi-line banner text, use --n to indicate -n",
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
func resourceBannerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBannerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBanner(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBannerRead(ctx, d, meta)
	}
	return diags
}

func resourceBannerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBannerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBanner(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBannerRead(ctx, d, meta)
	}
	return diags
}
func resourceBannerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBannerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBanner(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceBannerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBannerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBanner(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectBannerExecBannerCfg(d []interface{}) edpt.BannerExecBannerCfg {

	count1 := len(d)
	var ret edpt.BannerExecBannerCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Exec = in["exec"].(int)
		ret.ExecBanner = in["exec_banner"].(string)
	}
	return ret
}

func getObjectBannerLoginBannerCfg(d []interface{}) edpt.BannerLoginBannerCfg {

	count1 := len(d)
	var ret edpt.BannerLoginBannerCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Login = in["login"].(int)
		ret.LoginBanner = in["login_banner"].(string)
	}
	return ret
}

func dataToEndpointBanner(d *schema.ResourceData) edpt.Banner {
	var ret edpt.Banner
	ret.Inst.ExecBannerCfg = getObjectBannerExecBannerCfg(d.Get("exec_banner_cfg").([]interface{}))
	ret.Inst.LoginBannerCfg = getObjectBannerLoginBannerCfg(d.Get("login_banner_cfg").([]interface{}))
	//omit uuid
	return ret
}
