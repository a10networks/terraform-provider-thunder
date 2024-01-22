package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAccessListStandard() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_access_list_standard`: Configure Standard Access List\n\n__PLACEHOLDER__",
		CreateContext: resourceAccessListStandardCreate,
		UpdateContext: resourceAccessListStandardUpdate,
		ReadContext:   resourceAccessListStandardRead,
		DeleteContext: resourceAccessListStandardDelete,

		Schema: map[string]*schema.Schema{
			"std": {
				Type: schema.TypeInt, Required: true, Description: "IP standard access list",
			},
			"stdrules": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number",
						},
						"std_remark": {
							Type: schema.TypeString, Optional: true, Description: "Access list entry comment (Notes for this ACL)",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Deny; 'permit': Permit; 'l3-vlan-fwd-disable': Disable L3 forwarding between VLANs;",
						},
						"any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any source host",
						},
						"host": {
							Type: schema.TypeString, Optional: true, Description: "A single source host (Host address)",
						},
						"subnet": {
							Type: schema.TypeString, Optional: true, Description: "Source Address",
						},
						"rev_subnet_mask": {
							Type: schema.TypeString, Optional: true, Description: "Network Mask 0=apply 255=ignore",
						},
						"log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log matches against this entry",
						},
						"transparent_session_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only log transparent sessions",
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
func resourceAccessListStandardCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccessListStandardCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccessListStandard(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAccessListStandardRead(ctx, d, meta)
	}
	return diags
}

func resourceAccessListStandardUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccessListStandardUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccessListStandard(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAccessListStandardRead(ctx, d, meta)
	}
	return diags
}
func resourceAccessListStandardDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccessListStandardDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccessListStandard(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAccessListStandardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccessListStandardRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccessListStandard(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAccessListStandardStdrules(d []interface{}) []edpt.AccessListStandardStdrules {

	count1 := len(d)
	ret := make([]edpt.AccessListStandardStdrules, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AccessListStandardStdrules
		oi.SeqNum = in["seq_num"].(int)
		oi.StdRemark = in["std_remark"].(string)
		oi.Action = in["action"].(string)
		oi.Any = in["any"].(int)
		oi.Host = in["host"].(string)
		oi.Subnet = in["subnet"].(string)
		oi.RevSubnetMask = in["rev_subnet_mask"].(string)
		oi.Log = in["log"].(int)
		oi.TransparentSessionOnly = in["transparent_session_only"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAccessListStandard(d *schema.ResourceData) edpt.AccessListStandard {
	var ret edpt.AccessListStandard
	ret.Inst.Std = d.Get("std").(int)
	ret.Inst.Stdrules = getSliceAccessListStandardStdrules(d.Get("stdrules").([]interface{}))
	//omit uuid
	return ret
}
