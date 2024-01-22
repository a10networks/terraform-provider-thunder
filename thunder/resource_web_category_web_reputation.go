package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryWebReputation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_category_web_reputation`: Used for Web-Reputation Show Commands\n\n__PLACEHOLDER__",
		CreateContext: resourceWebCategoryWebReputationCreate,
		UpdateContext: resourceWebCategoryWebReputationUpdate,
		ReadContext:   resourceWebCategoryWebReputationRead,
		DeleteContext: resourceWebCategoryWebReputationDelete,

		Schema: map[string]*schema.Schema{
			"bypassed_urls": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"intercepted_urls": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"url": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceWebCategoryWebReputationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryWebReputationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryWebReputation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryWebReputationRead(ctx, d, meta)
	}
	return diags
}

func resourceWebCategoryWebReputationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryWebReputationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryWebReputation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryWebReputationRead(ctx, d, meta)
	}
	return diags
}
func resourceWebCategoryWebReputationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryWebReputationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryWebReputation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebCategoryWebReputationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryWebReputationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryWebReputation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectWebCategoryWebReputationBypassedUrls3658(d []interface{}) edpt.WebCategoryWebReputationBypassedUrls3658 {

	var ret edpt.WebCategoryWebReputationBypassedUrls3658
	return ret
}

func getObjectWebCategoryWebReputationInterceptedUrls3659(d []interface{}) edpt.WebCategoryWebReputationInterceptedUrls3659 {

	var ret edpt.WebCategoryWebReputationInterceptedUrls3659
	return ret
}

func getObjectWebCategoryWebReputationUrl3660(d []interface{}) edpt.WebCategoryWebReputationUrl3660 {

	var ret edpt.WebCategoryWebReputationUrl3660
	return ret
}

func dataToEndpointWebCategoryWebReputation(d *schema.ResourceData) edpt.WebCategoryWebReputation {
	var ret edpt.WebCategoryWebReputation
	ret.Inst.BypassedUrls = getObjectWebCategoryWebReputationBypassedUrls3658(d.Get("bypassed_urls").([]interface{}))
	ret.Inst.InterceptedUrls = getObjectWebCategoryWebReputationInterceptedUrls3659(d.Get("intercepted_urls").([]interface{}))
	ret.Inst.Url = getObjectWebCategoryWebReputationUrl3660(d.Get("url").([]interface{}))
	//omit uuid
	return ret
}
