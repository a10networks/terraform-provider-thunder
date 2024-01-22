package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDst() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair_settings_l4_type_src_dst`: DDOS L4 type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstDelete,

		Schema: map[string]*schema.Schema{
			"apply_policy_on_overflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
			},
			"max_dynamic_entry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic src-dst entry",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"all_types": {
				Type: schema.TypeString, Required: true, Description: "AllTypes",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairSettingsL4TypeSrcDst(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairSettingsL4TypeSrcDst(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairSettingsL4TypeSrcDst(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairSettingsL4TypeSrcDstRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairSettingsL4TypeSrcDst(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstEntrySrcDstPairSettingsL4TypeSrcDst(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairSettingsL4TypeSrcDst {
	var ret edpt.DdosDstEntrySrcDstPairSettingsL4TypeSrcDst
	ret.Inst.ApplyPolicyOnOverflow = d.Get("apply_policy_on_overflow").(int)
	ret.Inst.MaxDynamicEntryCount = d.Get("max_dynamic_entry_count").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.AllTypes = d.Get("all_types").(string)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
