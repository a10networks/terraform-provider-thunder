package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateFix() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_fix`: FIX template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateFixCreate,
		UpdateContext: resourceSlbTemplateFixUpdate,
		ReadContext:   resourceSlbTemplateFixRead,
		DeleteContext: resourceSlbTemplateFixDelete,

		Schema: map[string]*schema.Schema{
			"insert_client_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert client ip to tag 11447",
			},
			"logging": {
				Type: schema.TypeString, Optional: true, Description: "'init': init only log; 'term': termination only log; 'both': both initial and termination log;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "FIX Template Name",
			},
			"tag_switching": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switching_type": {
							Type: schema.TypeString, Optional: true, Description: "'sender-comp-id': Select service group based on SenderCompID; 'target-comp-id': Select service group based on TargetCompID;",
						},
						"equals": {
							Type: schema.TypeString, Optional: true, Description: "Equals (Tag String)",
						},
						"service_group": {
							Type: schema.TypeString, Optional: true, Description: "Create a Service Group comprising Servers (Service Group Name)",
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
func resourceSlbTemplateFixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateFixCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateFix(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateFixRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateFixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateFixUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateFix(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateFixRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateFixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateFixDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateFix(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateFixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateFixRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateFix(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateFixTagSwitching(d []interface{}) []edpt.SlbTemplateFixTagSwitching {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateFixTagSwitching, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateFixTagSwitching
		oi.SwitchingType = in["switching_type"].(string)
		oi.Equals = in["equals"].(string)
		oi.ServiceGroup = in["service_group"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateFix(d *schema.ResourceData) edpt.SlbTemplateFix {
	var ret edpt.SlbTemplateFix
	ret.Inst.InsertClientIp = d.Get("insert_client_ip").(int)
	ret.Inst.Logging = d.Get("logging").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TagSwitching = getSliceSlbTemplateFixTagSwitching(d.Get("tag_switching").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
