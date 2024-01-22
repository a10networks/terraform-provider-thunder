package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLocale() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_locale`: Set locale for the CLI startup\n\n__PLACEHOLDER__",
		CreateContext: resourceLocaleCreate,
		UpdateContext: resourceLocaleUpdate,
		ReadContext:   resourceLocaleRead,
		DeleteContext: resourceLocaleDelete,

		Schema: map[string]*schema.Schema{
			"test": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"locale": {
							Type: schema.TypeString, Optional: true, Description: "'zh_CN': Chinese locale for PRC; 'zh_TW': Chinese locale for Taiwan; 'ja_JP': Japanese locale for Japan;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeString, Optional: true, Default: "en_US.UTF-8", Description: "'en_US.UTF-8': English locale for the USA, encoding with UTF-8 (default); 'zh_CN.UTF-8': Chinese locale for PRC, encoding with UTF-8; 'zh_CN.GB18030': Chinese locale for PRC, encoding with GB18030; 'zh_CN.GBK': Chinese locale for PRC, encoding with GBK; 'zh_CN.GB2312': Chinese locale for PRC, encoding with GB2312; 'zh_TW.UTF-8': Chinese locale for Taiwan, encoding with UTF-8; 'zh_TW.BIG5': Chinese locale for Taiwan, encoding with BIG5; 'zh_TW.EUCTW': Chinese locale for Taiwan, encoding with EUC-TW; 'ja_JP.UTF-8': Japanese locale for Japan, encoding with UTF-8; 'ja_JP.EUC-JP': Japanese locale for Japan, encoding with EUC-JP;",
			},
		},
	}
}
func resourceLocaleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLocaleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLocale(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLocaleRead(ctx, d, meta)
	}
	return diags
}

func resourceLocaleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLocaleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLocale(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLocaleRead(ctx, d, meta)
	}
	return diags
}
func resourceLocaleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLocaleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLocale(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLocaleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLocaleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLocale(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectLocaleTest1049(d []interface{}) edpt.LocaleTest1049 {

	count1 := len(d)
	var ret edpt.LocaleTest1049
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Locale = in["locale"].(string)
	}
	return ret
}

func dataToEndpointLocale(d *schema.ResourceData) edpt.Locale {
	var ret edpt.Locale
	ret.Inst.Test = getObjectLocaleTest1049(d.Get("test").([]interface{}))
	//omit uuid
	ret.Inst.Value = d.Get("value").(string)
	return ret
}
