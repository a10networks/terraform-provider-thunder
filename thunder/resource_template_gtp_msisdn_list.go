package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpMsisdnList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_msisdn_list`: Configure GTP MSISDN list\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpMsisdnListCreate,
		UpdateContext: resourceTemplateGtpMsisdnListUpdate,
		ReadContext:   resourceTemplateGtpMsisdnListRead,
		DeleteContext: resourceTemplateGtpMsisdnListDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "deny", Description: "'permit': Create a whitelist to permit the packets that match MSISDN filters; 'deny': Create a blacklist to deny the packets that match MSISDN filters;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the GTP MSISDN list",
			},
			"str_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msisdn": {
							Type: schema.TypeString, Optional: true, Description: "Specify the MSISDN filter",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceTemplateGtpMsisdnListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMsisdnListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMsisdnList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpMsisdnListRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpMsisdnListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMsisdnListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMsisdnList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpMsisdnListRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpMsisdnListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMsisdnListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMsisdnList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpMsisdnListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMsisdnListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMsisdnList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceTemplateGtpMsisdnListStrList(d []interface{}) []edpt.TemplateGtpMsisdnListStrList {

	count1 := len(d)
	ret := make([]edpt.TemplateGtpMsisdnListStrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateGtpMsisdnListStrList
		oi.Msisdn = in["msisdn"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTemplateGtpMsisdnList(d *schema.ResourceData) edpt.TemplateGtpMsisdnList {
	var ret edpt.TemplateGtpMsisdnList
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.StrList = getSliceTemplateGtpMsisdnListStrList(d.Get("str_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
