package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpMapList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_map_list`: Configure IP Map List name\n\n__PLACEHOLDER__",
		CreateContext: resourceIpMapListCreate,
		UpdateContext: resourceIpMapListUpdate,
		ReadContext:   resourceIpMapListRead,
		DeleteContext: resourceIpMapListDelete,

		Schema: map[string]*schema.Schema{
			"file": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create/Edit a IP Map List stored as a file",
			},
			"mapping_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_start_ip": {
							Type: schema.TypeString, Optional: true, Description: "Local Start IPv4 Address of this list",
						},
						"global_start_ip": {
							Type: schema.TypeString, Optional: true, Description: "Global Start IPv4 Address of this list",
						},
						"count1": {
							Type: schema.TypeInt, Optional: true, Description: "Number of addresses to be translated in this range",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the IP Map List",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpMapListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpMapListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpMapList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpMapListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpMapListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpMapListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpMapList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpMapListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpMapListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpMapListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpMapList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpMapListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpMapListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpMapList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpMapListMappingList(d []interface{}) []edpt.IpMapListMappingList {

	count1 := len(d)
	ret := make([]edpt.IpMapListMappingList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpMapListMappingList
		oi.LocalStartIp = in["local_start_ip"].(string)
		oi.GlobalStartIp = in["global_start_ip"].(string)
		oi.Count1 = in["count1"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpMapList(d *schema.ResourceData) edpt.IpMapList {
	var ret edpt.IpMapList
	ret.Inst.File = d.Get("file").(int)
	ret.Inst.MappingList = getSliceIpMapListMappingList(d.Get("mapping_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
