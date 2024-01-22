package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateOther() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_other`: OTHER template configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateOtherCreate,
		UpdateContext: resourceDdosTemplateOtherUpdate,
		ReadContext:   resourceDdosTemplateOtherRead,
		DeleteContext: resourceDdosTemplateOtherDelete,

		Schema: map[string]*schema.Schema{
			"filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"other_filter_seq": {
							Type: schema.TypeInt, Required: true, Description: "Sequence number",
						},
						"other_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "Filter Expression using Berkeley Packet Filter syntax",
						},
						"other_filter_unmatched": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "action taken when it does not match",
						},
						"other_filter_action": {
							Type: schema.TypeString, Optional: true, Description: "'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;",
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
				Type: schema.TypeString, Required: true, Description: "DDOS OTHER Template Name",
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
func resourceDdosTemplateOtherCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateOtherCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateOther(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateOtherRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateOtherUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateOtherUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateOther(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateOtherRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateOtherDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateOtherDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateOther(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateOtherRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateOtherRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateOther(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosTemplateOtherFilterList(d []interface{}) []edpt.DdosTemplateOtherFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateOtherFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateOtherFilterList
		oi.OtherFilterSeq = in["other_filter_seq"].(int)
		oi.OtherFilterRegex = in["other_filter_regex"].(string)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.OtherFilterUnmatched = in["other_filter_unmatched"].(int)
		oi.OtherFilterAction = in["other_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosTemplateOther(d *schema.ResourceData) edpt.DdosTemplateOther {
	var ret edpt.DdosTemplateOther
	ret.Inst.FilterList = getSliceDdosTemplateOtherFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
