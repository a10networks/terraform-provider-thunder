package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTechreport() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_techreport`: Configure show tech\n\n__PLACEHOLDER__",
		CreateContext: resourceTechreportCreate,
		UpdateContext: resourceTechreportUpdate,
		ReadContext:   resourceTechreportRead,
		DeleteContext: resourceTechreportDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the polling techreport",
			},
			"enable_full_history": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable 31 day poll techreports (default seven) on platforms with less than 32GB disk (no op otherwise)",
			},
			"interval": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeInt, Optional: true, Default: 15, Description: "Showtech interval in minutes (default is 15)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"max_logfile_size": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Log file size for periodic techsupport (default is 1)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"max_partitions": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Maximum partions to show in per periodic techsupport (default is 30)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"priority_partition_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"part_name": {
							Type: schema.TypeString, Required: true, Description: "Name of partition always want to show in showtech (shared is always shown by default)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceTechreportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreport(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTechreportRead(ctx, d, meta)
	}
	return diags
}

func resourceTechreportUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreport(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTechreportRead(ctx, d, meta)
	}
	return diags
}
func resourceTechreportDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreport(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTechreportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTechreportRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTechreport(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTechreportInterval1897(d []interface{}) edpt.TechreportInterval1897 {

	count1 := len(d)
	var ret edpt.TechreportInterval1897
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectTechreportMaxLogfileSize1898(d []interface{}) edpt.TechreportMaxLogfileSize1898 {

	count1 := len(d)
	var ret edpt.TechreportMaxLogfileSize1898
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectTechreportMaxPartitions1899(d []interface{}) edpt.TechreportMaxPartitions1899 {

	count1 := len(d)
	var ret edpt.TechreportMaxPartitions1899
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getSliceTechreportPriorityPartitionList(d []interface{}) []edpt.TechreportPriorityPartitionList {

	count1 := len(d)
	ret := make([]edpt.TechreportPriorityPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TechreportPriorityPartitionList
		oi.PartName = in["part_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTechreport(d *schema.ResourceData) edpt.Techreport {
	var ret edpt.Techreport
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.EnableFullHistory = d.Get("enable_full_history").(int)
	ret.Inst.Interval = getObjectTechreportInterval1897(d.Get("interval").([]interface{}))
	ret.Inst.MaxLogfileSize = getObjectTechreportMaxLogfileSize1898(d.Get("max_logfile_size").([]interface{}))
	ret.Inst.MaxPartitions = getObjectTechreportMaxPartitions1899(d.Get("max_partitions").([]interface{}))
	ret.Inst.PriorityPartitionList = getSliceTechreportPriorityPartitionList(d.Get("priority_partition_list").([]interface{}))
	//omit uuid
	return ret
}
