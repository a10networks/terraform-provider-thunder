package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6PortList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_port_list`: Configure port list\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6PortListCreate,
		UpdateContext: resourceCgnv6PortListUpdate,
		ReadContext:   resourceCgnv6PortListRead,
		DeleteContext: resourceCgnv6PortListDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the port list",
			},
			"port_config": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"original_port": {
							Type: schema.TypeInt, Optional: true, Description: "Original port to be translated",
						},
						"translated_port": {
							Type: schema.TypeInt, Optional: true, Description: "Port after translation",
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
		},
	}
}
func resourceCgnv6PortListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PortListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6PortList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6PortListRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6PortListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PortListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6PortList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6PortListRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6PortListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PortListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6PortList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6PortListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PortListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6PortList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6PortListPortConfig(d []interface{}) []edpt.Cgnv6PortListPortConfig {

	count1 := len(d)
	ret := make([]edpt.Cgnv6PortListPortConfig, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6PortListPortConfig
		oi.OriginalPort = in["original_port"].(int)
		oi.TranslatedPort = in["translated_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6PortList(d *schema.ResourceData) edpt.Cgnv6PortList {
	var ret edpt.Cgnv6PortList
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PortConfig = getSliceCgnv6PortListPortConfig(d.Get("port_config").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
