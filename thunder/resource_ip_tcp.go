package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_tcp`: Global TCP parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceIpTcpCreate,
		UpdateContext: resourceIpTcpUpdate,
		ReadContext:   resourceIpTcpRead,
		DeleteContext: resourceIpTcpDelete,

		Schema: map[string]*schema.Schema{
			"syn_cookie": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold": {
							Type: schema.TypeInt, Optional: true, Default: 4, Description: "SYN cookie expire threshold (seconds (default is 4))",
						},
						"sack_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable HW Syn-Cookie SACK support",
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
func resourceIpTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceIpTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceIpTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpTcpSynCookie(d []interface{}) edpt.IpTcpSynCookie {

	count1 := len(d)
	var ret edpt.IpTcpSynCookie
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Threshold = in["threshold"].(int)
		ret.SackEnable = in["sack_enable"].(int)
	}
	return ret
}

func dataToEndpointIpTcp(d *schema.ResourceData) edpt.IpTcp {
	var ret edpt.IpTcp
	ret.Inst.SynCookie = getObjectIpTcpSynCookie(d.Get("syn_cookie").([]interface{}))
	//omit uuid
	return ret
}
