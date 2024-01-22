package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLocaleTest() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_locale_test`: To test current terminal encodings for specific locale\n\n__PLACEHOLDER__",
		CreateContext: resourceLocaleTestCreate,
		UpdateContext: resourceLocaleTestUpdate,
		ReadContext:   resourceLocaleTestRead,
		DeleteContext: resourceLocaleTestDelete,

		Schema: map[string]*schema.Schema{
			"locale": {
				Type: schema.TypeString, Optional: true, Description: "'zh_CN': Chinese locale for PRC; 'zh_TW': Chinese locale for Taiwan; 'ja_JP': Japanese locale for Japan;",
			},
		},
	}
}
func resourceLocaleTestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLocaleTestCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLocaleTest(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLocaleTestRead(ctx, d, meta)
	}
	return diags
}

func resourceLocaleTestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLocaleTestUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLocaleTest(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLocaleTestRead(ctx, d, meta)
	}
	return diags
}
func resourceLocaleTestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLocaleTestDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLocaleTest(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLocaleTestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLocaleTestRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLocaleTest(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLocaleTest(d *schema.ResourceData) edpt.LocaleTest {
	var ret edpt.LocaleTest
	ret.Inst.Locale = d.Get("locale").(string)
	return ret
}
