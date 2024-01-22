package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpApnImsiList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_apn_imsi_list`: Configure GTP APN IMSI list\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpApnImsiListCreate,
		UpdateContext: resourceTemplateGtpApnImsiListUpdate,
		ReadContext:   resourceTemplateGtpApnImsiListRead,
		DeleteContext: resourceTemplateGtpApnImsiListDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "deny", Description: "'permit': Create a whitelist to permit the packets that match APN IMSI filters; 'deny': Create a blacklist to deny the packets that match APN IMSI filters (default);",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the GTP APN IMSI list",
			},
			"str_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apn": {
							Type: schema.TypeString, Optional: true, Description: "Specify the APN filter",
						},
						"selection_mode": {
							Type: schema.TypeString, Optional: true, Description: "'mobilestation': MS provided APN, subscription not verified; 'network': Network provided APN, subscription not verified; 'verified': MS or Network provided APN, subscription verified;",
						},
						"imsi_selection": {
							Type: schema.TypeString, Optional: true, Description: "Specify the IMSI filter",
						},
						"imsi": {
							Type: schema.TypeString, Optional: true, Description: "Specify the IMSI filter",
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
func resourceTemplateGtpApnImsiListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpApnImsiListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpApnImsiList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpApnImsiListRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpApnImsiListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpApnImsiListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpApnImsiList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpApnImsiListRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpApnImsiListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpApnImsiListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpApnImsiList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpApnImsiListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpApnImsiListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpApnImsiList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceTemplateGtpApnImsiListStrList(d []interface{}) []edpt.TemplateGtpApnImsiListStrList {

	count1 := len(d)
	ret := make([]edpt.TemplateGtpApnImsiListStrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateGtpApnImsiListStrList
		oi.Apn = in["apn"].(string)
		oi.SelectionMode = in["selection_mode"].(string)
		oi.ImsiSelection = in["imsi_selection"].(string)
		oi.Imsi = in["imsi"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTemplateGtpApnImsiList(d *schema.ResourceData) edpt.TemplateGtpApnImsiList {
	var ret edpt.TemplateGtpApnImsiList
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.StrList = getSliceTemplateGtpApnImsiListStrList(d.Get("str_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
