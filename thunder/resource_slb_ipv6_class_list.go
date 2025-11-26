package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbIpv6ClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_ipv6_class_list`: IPv6 subnet add remove config\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbIpv6ClassListCreate,
		UpdateContext: resourceSlbIpv6ClassListUpdate,
		ReadContext:   resourceSlbIpv6ClassListRead,
		DeleteContext: resourceSlbIpv6ClassListDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'add': Add the entry; 'delete': Delete the entry;",
						},
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 host or subnet",
						},
						"lid": {
							Type: schema.TypeInt, Optional: true, Description: "Use Limit ID defined in template (Specify LID index)",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Use global Limit ID (Specify global LID index)",
						},
						"lsn_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LSN Limit ID (LID index)",
						},
						"lsn_radius_profile": {
							Type: schema.TypeInt, Optional: true, Description: "LSN RADIUS Profile Index",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the class list",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
		},
	}
}
func resourceSlbIpv6ClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbIpv6ClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbIpv6ClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbIpv6ClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbIpv6ClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbIpv6ClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbIpv6ClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbIpv6ClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbIpv6ClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbIpv6ClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbIpv6ClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbIpv6ClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbIpv6ClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbIpv6ClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbIpv6ClassListIpv6List(d []interface{}) []edpt.SlbIpv6ClassListIpv6List {

	count1 := len(d)
	ret := make([]edpt.SlbIpv6ClassListIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbIpv6ClassListIpv6List
		oi.Action = in["action"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.Lid = in["lid"].(int)
		oi.Glid = in["glid"].(string)
		oi.LsnLid = in["lsn_lid"].(int)
		oi.LsnRadiusProfile = in["lsn_radius_profile"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbIpv6ClassList(d *schema.ResourceData) edpt.SlbIpv6ClassList {
	var ret edpt.SlbIpv6ClassList
	ret.Inst.Ipv6List = getSliceSlbIpv6ClassListIpv6List(d.Get("ipv6_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	return ret
}
