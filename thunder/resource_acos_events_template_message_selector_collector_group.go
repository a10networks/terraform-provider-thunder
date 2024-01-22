package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsTemplateMessageSelectorCollectorGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_template_message_selector_collector_group`: Specify the log server group for receiving log messages\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsTemplateMessageSelectorCollectorGroupCreate,
		UpdateContext: resourceAcosEventsTemplateMessageSelectorCollectorGroupUpdate,
		ReadContext:   resourceAcosEventsTemplateMessageSelectorCollectorGroupRead,
		DeleteContext: resourceAcosEventsTemplateMessageSelectorCollectorGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify the log server group for receiving log messages",
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
func resourceAcosEventsTemplateMessageSelectorCollectorGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsTemplateMessageSelectorCollectorGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsTemplateMessageSelectorCollectorGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsTemplateMessageSelectorCollectorGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsTemplateMessageSelectorCollectorGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsTemplateMessageSelectorCollectorGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsTemplateMessageSelectorCollectorGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsTemplateMessageSelectorCollectorGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsTemplateMessageSelectorCollectorGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsTemplateMessageSelectorCollectorGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsTemplateMessageSelectorCollectorGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsTemplateMessageSelectorCollectorGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsTemplateMessageSelectorCollectorGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsTemplateMessageSelectorCollectorGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsTemplateMessageSelectorCollectorGroup(d *schema.ResourceData) edpt.AcosEventsTemplateMessageSelectorCollectorGroup {
	var ret edpt.AcosEventsTemplateMessageSelectorCollectorGroup
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
