package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsTemplateMessageSelector() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_template_message_selector`: Specify the message selector\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsTemplateMessageSelectorCreate,
		UpdateContext: resourceAcosEventsTemplateMessageSelectorUpdate,
		ReadContext:   resourceAcosEventsTemplateMessageSelectorRead,
		DeleteContext: resourceAcosEventsTemplateMessageSelectorDelete,

		Schema: map[string]*schema.Schema{
			"collector_group_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify the log server group for receiving log messages",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify the message selector name",
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
func resourceAcosEventsTemplateMessageSelectorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsTemplateMessageSelectorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsTemplateMessageSelector(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsTemplateMessageSelectorRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsTemplateMessageSelectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsTemplateMessageSelectorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsTemplateMessageSelector(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsTemplateMessageSelectorRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsTemplateMessageSelectorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsTemplateMessageSelectorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsTemplateMessageSelector(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsTemplateMessageSelectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsTemplateMessageSelectorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsTemplateMessageSelector(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAcosEventsTemplateMessageSelectorCollectorGroupList(d []interface{}) []edpt.AcosEventsTemplateMessageSelectorCollectorGroupList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsTemplateMessageSelectorCollectorGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsTemplateMessageSelectorCollectorGroupList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosEventsTemplateMessageSelector(d *schema.ResourceData) edpt.AcosEventsTemplateMessageSelector {
	var ret edpt.AcosEventsTemplateMessageSelector
	ret.Inst.CollectorGroupList = getSliceAcosEventsTemplateMessageSelectorCollectorGroupList(d.Get("collector_group_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
