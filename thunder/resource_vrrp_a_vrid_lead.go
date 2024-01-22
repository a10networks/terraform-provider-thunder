package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAVridLead() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_vrid_lead`: Define a vrid-lead\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAVridLeadCreate,
		UpdateContext: resourceVrrpAVridLeadUpdate,
		ReadContext:   resourceVrrpAVridLeadRead,
		DeleteContext: resourceVrrpAVridLeadDelete,

		Schema: map[string]*schema.Schema{
			"partition": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partition": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Partition name",
						},
						"shared_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Shared partition",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "VRRP-A id",
									},
									"vrid_value": {
										Type: schema.TypeInt, Optional: true, Description: "Specify ha VRRP-A vrid",
									},
								},
							},
						},
						"name_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "Partition name",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "VRRP-A id",
									},
									"vrid_value": {
										Type: schema.TypeInt, Optional: true, Description: "Specify ha VRRP-A vrid",
									},
								},
							},
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
			"vrid_lead_str": {
				Type: schema.TypeString, Required: true, Description: "VRRP-A VRID leader name",
			},
		},
	}
}
func resourceVrrpAVridLeadCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridLeadCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridLead(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridLeadRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAVridLeadUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridLeadUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridLead(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridLeadRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAVridLeadDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridLeadDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridLead(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAVridLeadRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridLeadRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridLead(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVrrpAVridLeadPartition(d []interface{}) edpt.VrrpAVridLeadPartition {

	count1 := len(d)
	var ret edpt.VrrpAVridLeadPartition
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Partition = in["partition"].(int)
		ret.SharedCfg = getObjectVrrpAVridLeadPartitionSharedCfg(in["shared_cfg"].([]interface{}))
		ret.NameCfg = getObjectVrrpAVridLeadPartitionNameCfg(in["name_cfg"].([]interface{}))
	}
	return ret
}

func getObjectVrrpAVridLeadPartitionSharedCfg(d []interface{}) edpt.VrrpAVridLeadPartitionSharedCfg {

	count1 := len(d)
	var ret edpt.VrrpAVridLeadPartitionSharedCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Shared = in["shared"].(int)
		ret.Vrid = in["vrid"].(int)
		ret.VridValue = in["vrid_value"].(int)
	}
	return ret
}

func getObjectVrrpAVridLeadPartitionNameCfg(d []interface{}) edpt.VrrpAVridLeadPartitionNameCfg {

	count1 := len(d)
	var ret edpt.VrrpAVridLeadPartitionNameCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.Vrid = in["vrid"].(int)
		ret.VridValue = in["vrid_value"].(int)
	}
	return ret
}

func dataToEndpointVrrpAVridLead(d *schema.ResourceData) edpt.VrrpAVridLead {
	var ret edpt.VrrpAVridLead
	ret.Inst.Partition = getObjectVrrpAVridLeadPartition(d.Get("partition").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VridLeadStr = d.Get("vrid_lead_str").(string)
	return ret
}
