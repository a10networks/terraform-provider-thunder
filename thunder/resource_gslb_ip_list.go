package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbIpList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_ip_list`: Specify a IP List\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbIpListCreate,
		UpdateContext: resourceGslbIpListUpdate,
		ReadContext:   resourceGslbIpListRead,
		DeleteContext: resourceGslbIpListDelete,

		Schema: map[string]*schema.Schema{
			"gslb_ip_list_addr_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP address",
						},
						"ip_mask": {
							Type: schema.TypeString, Optional: true, Description: "IP mask",
						},
						"id1": {
							Type: schema.TypeInt, Optional: true, Description: "ID Number",
						},
					},
				},
			},
			"gslb_ip_list_filename": {
				Type: schema.TypeString, Optional: true, Description: "Load IP List file (IP List filename)",
			},
			"gslb_ip_list_obj_name": {
				Type: schema.TypeString, Required: true, Description: "Specify IP List name",
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
func resourceGslbIpListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbIpListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbIpList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbIpListRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbIpListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbIpListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbIpList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbIpListRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbIpListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbIpListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbIpList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbIpListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbIpListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbIpList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbIpListGslbIpListAddrList(d []interface{}) []edpt.GslbIpListGslbIpListAddrList {

	count1 := len(d)
	ret := make([]edpt.GslbIpListGslbIpListAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbIpListGslbIpListAddrList
		oi.Ip = in["ip"].(string)
		oi.IpMask = in["ip_mask"].(string)
		oi.Id1 = in["id1"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbIpList(d *schema.ResourceData) edpt.GslbIpList {
	var ret edpt.GslbIpList
	ret.Inst.GslbIpListAddrList = getSliceGslbIpListGslbIpListAddrList(d.Get("gslb_ip_list_addr_list").([]interface{}))
	ret.Inst.GslbIpListFilename = d.Get("gslb_ip_list_filename").(string)
	ret.Inst.GslbIpListObjName = d.Get("gslb_ip_list_obj_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
