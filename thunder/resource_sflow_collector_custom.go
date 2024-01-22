package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowCollectorCustom() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_collector_custom`: Configure Custom sFlow collector\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowCollectorCustomCreate,
		UpdateContext: resourceSflowCollectorCustomUpdate,
		ReadContext:   resourceSflowCollectorCustomRead,
		DeleteContext: resourceSflowCollectorCustomDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_addr": {
				Type: schema.TypeString, Optional: true, Description: "Configure IP address of the sFlow Receiver",
			},
			"ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "Configure IPv6 address of the sFlow Receiver",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Configure name of sFlow Receiver",
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
func resourceSflowCollectorCustomCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorCustomCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorCustom(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorCustomRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowCollectorCustomUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorCustomUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorCustom(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorCustomRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowCollectorCustomDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorCustomDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorCustom(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowCollectorCustomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorCustomRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorCustom(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowCollectorCustom(d *schema.ResourceData) edpt.SflowCollectorCustom {
	var ret edpt.SflowCollectorCustom
	ret.Inst.Ipv4_addr = d.Get("ipv4_addr").(string)
	ret.Inst.Ipv6_addr = d.Get("ipv6_addr").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
