package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpUnnumbered() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_unnumbered`: unnumbered feature\n\n__PLACEHOLDER__",
		CreateContext: resourceIpUnnumberedCreate,
		UpdateContext: resourceIpUnnumberedUpdate,
		ReadContext:   resourceIpUnnumberedRead,
		DeleteContext: resourceIpUnnumberedDelete,

		Schema: map[string]*schema.Schema{
			"use_source_ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"update_source_ip": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
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
func resourceIpUnnumberedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpUnnumberedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpUnnumbered(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpUnnumberedRead(ctx, d, meta)
	}
	return diags
}

func resourceIpUnnumberedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpUnnumberedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpUnnumbered(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpUnnumberedRead(ctx, d, meta)
	}
	return diags
}
func resourceIpUnnumberedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpUnnumberedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpUnnumbered(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpUnnumberedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpUnnumberedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpUnnumbered(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpUnnumberedUseSourceIp1043(d []interface{}) edpt.IpUnnumberedUseSourceIp1043 {

	count1 := len(d)
	var ret edpt.IpUnnumberedUseSourceIp1043
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UpdateSourceIp = in["update_source_ip"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointIpUnnumbered(d *schema.ResourceData) edpt.IpUnnumbered {
	var ret edpt.IpUnnumbered
	ret.Inst.UseSourceIp = getObjectIpUnnumberedUseSourceIp1043(d.Get("use_source_ip").([]interface{}))
	//omit uuid
	return ret
}
