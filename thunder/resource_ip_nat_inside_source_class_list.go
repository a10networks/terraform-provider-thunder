package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatInsideSourceClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_inside_source_class_list`: Class List\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatInsideSourceClassListCreate,
		UpdateContext: resourceIpNatInsideSourceClassListUpdate,
		ReadContext:   resourceIpNatInsideSourceClassListRead,
		DeleteContext: resourceIpNatInsideSourceClassListDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Class List Name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpNatInsideSourceClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatInsideSourceClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatInsideSourceClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatInsideSourceClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatInsideSourceClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatInsideSourceClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpNatInsideSourceClassList(d *schema.ResourceData) edpt.IpNatInsideSourceClassList {
	var ret edpt.IpNatInsideSourceClassList
	ret.Inst.Name = d.Get("name").(string)
	//omit uuid
	return ret
}
