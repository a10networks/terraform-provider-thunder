package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDynamicClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dynamic_class_list`: Configure class-list for automatic population by BGP\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDynamicClassListCreate,
		UpdateContext: resourceDdosDynamicClassListUpdate,
		ReadContext:   resourceDdosDynamicClassListRead,
		DeleteContext: resourceDdosDynamicClassListDelete,

		Schema: map[string]*schema.Schema{
			"class_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify name of the class list",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDynamicClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDynamicClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDynamicClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDynamicClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDynamicClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDynamicClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDynamicClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDynamicClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDynamicClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDynamicClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDynamicClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDynamicClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDynamicClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDynamicClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDynamicClassList(d *schema.ResourceData) edpt.DdosDynamicClassList {
	var ret edpt.DdosDynamicClassList
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	//omit uuid
	return ret
}
