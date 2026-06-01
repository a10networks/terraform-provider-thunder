package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcIpFiltering() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_src_ip_filtering`: Configure src-ip-filtering\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneSrcIpFilteringCreate,
		UpdateContext: resourceDdosDstZoneSrcIpFilteringUpdate,
		ReadContext:   resourceDdosDstZoneSrcIpFilteringRead,
		DeleteContext: resourceDdosDstZoneSrcIpFilteringDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Apply src IP filtering",
			},
			"per_class_list_hit_tracking": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable per class-list hit count tracking",
			},
			"per_service_tracking": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable per service drop and bypass count tracking",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'class-list-1-match': Packets Match Class-List 1; 'class-list-2-match': Packets Match Class-List 2; 'class-list-3-match': Packets Match Class-List 3; 'class-list-4-match': Packets Match Class-List 4; 'class-list-5-match': Packets Match Class-List 5; 'class-list-6-match': Packets Match Class-List 6; 'class-list-7-match': Packets Match Class-List 7; 'class-list-8-match': Packets Match Class-List 8;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneSrcIpFilteringCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcIpFilteringCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcIpFiltering(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcIpFilteringRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneSrcIpFilteringUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcIpFilteringUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcIpFiltering(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcIpFilteringRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneSrcIpFilteringDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcIpFilteringDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcIpFiltering(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneSrcIpFilteringRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcIpFilteringRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcIpFiltering(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZoneSrcIpFilteringSamplingEnable(d []interface{}) []edpt.DdosDstZoneSrcIpFilteringSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcIpFilteringSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcIpFilteringSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcIpFiltering(d *schema.ResourceData) edpt.DdosDstZoneSrcIpFiltering {
	var ret edpt.DdosDstZoneSrcIpFiltering
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PerClassListHitTracking = d.Get("per_class_list_hit_tracking").(int)
	ret.Inst.PerServiceTracking = d.Get("per_service_tracking").(int)
	ret.Inst.SamplingEnable = getSliceDdosDstZoneSrcIpFilteringSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
