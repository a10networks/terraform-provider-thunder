package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcIpFiltering() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_ip_filtering`: Config src IP filtering\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcIpFilteringCreate,
		UpdateContext: resourceDdosSrcIpFilteringUpdate,
		ReadContext:   resourceDdosSrcIpFilteringRead,
		DeleteContext: resourceDdosSrcIpFilteringDelete,

		Schema: map[string]*schema.Schema{
			"filter_class_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_name": {
							Type: schema.TypeString, Required: true, Description: "Class-list name",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Deny incoming packets by Source IP; 'bypass': Bypass incoming packets by Source IP;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the Filter",
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
func resourceDdosSrcIpFilteringCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcIpFilteringCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcIpFiltering(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcIpFilteringRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcIpFilteringUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcIpFilteringUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcIpFiltering(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcIpFilteringRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcIpFilteringDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcIpFilteringDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcIpFiltering(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcIpFilteringRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcIpFilteringRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcIpFiltering(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosSrcIpFilteringFilterClassListList(d []interface{}) []edpt.DdosSrcIpFilteringFilterClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcIpFilteringFilterClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcIpFilteringFilterClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Action = in["action"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosSrcIpFiltering(d *schema.ResourceData) edpt.DdosSrcIpFiltering {
	var ret edpt.DdosSrcIpFiltering
	ret.Inst.FilterClassListList = getSliceDdosSrcIpFilteringFilterClassListList(d.Get("filter_class_list_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
