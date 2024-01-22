package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorHeaderInsert() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_header_insert`: Insert a header into request\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorHeaderInsertCreate,
		UpdateContext: resourceHealthMonitorHeaderInsertUpdate,
		ReadContext:   resourceHealthMonitorHeaderInsertRead,
		DeleteContext: resourceHealthMonitorHeaderInsertDelete,

		Schema: map[string]*schema.Schema{
			"insert_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"insert_content": {
							Type: schema.TypeString, Optional: true, Description: "Header Content (Format: \"[name]:[value]\")",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorHeaderInsertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorHeaderInsertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorHeaderInsert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorHeaderInsertRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorHeaderInsertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorHeaderInsertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorHeaderInsert(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorHeaderInsertRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorHeaderInsertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorHeaderInsertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorHeaderInsert(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorHeaderInsertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorHeaderInsertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorHeaderInsert(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceHealthMonitorHeaderInsertInsertList(d []interface{}) []edpt.HealthMonitorHeaderInsertInsertList {

	count1 := len(d)
	ret := make([]edpt.HealthMonitorHeaderInsertInsertList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.HealthMonitorHeaderInsertInsertList
		oi.InsertContent = in["insert_content"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointHealthMonitorHeaderInsert(d *schema.ResourceData) edpt.HealthMonitorHeaderInsert {
	var ret edpt.HealthMonitorHeaderInsert
	ret.Inst.InsertList = getSliceHealthMonitorHeaderInsertInsertList(d.Get("insert_list").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
