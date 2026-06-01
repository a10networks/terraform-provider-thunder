package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcIpFilteringFilterClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_ip_filtering_filter_class_list`: Configure class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcIpFilteringFilterClassListCreate,
		UpdateContext: resourceDdosSrcIpFilteringFilterClassListUpdate,
		ReadContext:   resourceDdosSrcIpFilteringFilterClassListRead,
		DeleteContext: resourceDdosSrcIpFilteringFilterClassListDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'deny': Deny incoming packets by Source IP; 'bypass': Bypass incoming packets by Source IP;",
			},
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "Class-list name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"src_ip_filtering_name": {
				Type: schema.TypeString, Required: true, Description: "Src_ip_filtering_name",
			},
		},
	}
}
func resourceDdosSrcIpFilteringFilterClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcIpFilteringFilterClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcIpFilteringFilterClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcIpFilteringFilterClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcIpFilteringFilterClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcIpFilteringFilterClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcIpFilteringFilterClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcIpFilteringFilterClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcIpFilteringFilterClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcIpFilteringFilterClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcIpFilteringFilterClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcIpFilteringFilterClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcIpFilteringFilterClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcIpFilteringFilterClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosSrcIpFilteringFilterClassList(d *schema.ResourceData) edpt.DdosSrcIpFilteringFilterClassList {
	var ret edpt.DdosSrcIpFilteringFilterClassList
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	//omit uuid
	ret.Inst.Src_ip_filtering_name = d.Get("src_ip_filtering_name").(string)
	return ret
}
