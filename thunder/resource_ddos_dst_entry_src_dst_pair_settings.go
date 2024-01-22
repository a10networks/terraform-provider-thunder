package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairSettings() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair_settings`: Configure src-dst-pair settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairSettingsCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairSettingsUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairSettingsRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairSettingsDelete,

		Schema: map[string]*schema.Schema{
			"age": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle age for ip entry",
			},
			"all_types": {
				Type: schema.TypeString, Required: true, Description: "'all-types': Settings for all types (default or class-list);",
			},
			"apply_policy_on_overflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
			},
			"enable_class_list_overflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply class-list overflow policy upon exceeding dynamic entry count specified for DST entry or each class-list",
			},
			"l4_type_src_dst_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
						},
						"max_dynamic_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic src-dst entry",
						},
						"apply_policy_on_overflow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
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
			"max_dynamic_entry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic src-dst entry",
			},
			"src_prefix_len": {
				Type: schema.TypeInt, Optional: true, Description: "Specify src prefix length for IPv6 (default: not set)",
			},
			"unlimited_dynamic_entry_count": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "No limit for maximum dynamic src entry count",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntrySrcDstPairSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairSettingsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairSettings(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairSettingsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairSettingsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairSettings(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairSettingsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairSettingsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairSettings(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairSettingsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairSettings(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairSettingsL4TypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairSettingsL4TypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairSettingsL4TypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ApplyPolicyOnOverflow = in["apply_policy_on_overflow"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstEntrySrcDstPairSettings(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairSettings {
	var ret edpt.DdosDstEntrySrcDstPairSettings
	ret.Inst.Age = d.Get("age").(int)
	ret.Inst.AllTypes = d.Get("all_types").(string)
	ret.Inst.ApplyPolicyOnOverflow = d.Get("apply_policy_on_overflow").(int)
	ret.Inst.EnableClassListOverflow = d.Get("enable_class_list_overflow").(int)
	ret.Inst.L4TypeSrcDstList = getSliceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstList(d.Get("l4_type_src_dst_list").([]interface{}))
	ret.Inst.MaxDynamicEntryCount = d.Get("max_dynamic_entry_count").(int)
	ret.Inst.SrcPrefixLen = d.Get("src_prefix_len").(int)
	ret.Inst.UnlimitedDynamicEntryCount = d.Get("unlimited_dynamic_entry_count").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
