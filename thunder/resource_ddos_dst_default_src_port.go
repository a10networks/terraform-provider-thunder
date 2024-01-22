package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstDefaultSrcPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_default_src_port`: DDOS Source Port & Protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstDefaultSrcPortCreate,
		UpdateContext: resourceDdosDstDefaultSrcPortUpdate,
		ReadContext:   resourceDdosDstDefaultSrcPortRead,
		DeleteContext: resourceDdosDstDefaultSrcPortDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'udp': udp; 'tcp': tcp;",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_udp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS udp src template",
						},
						"src_tcp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS tcp src template",
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
			"default_address_type": {
				Type: schema.TypeString, Required: true, Description: "DefaultAddressType",
			},
		},
	}
}
func resourceDdosDstDefaultSrcPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultSrcPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultSrcPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDefaultSrcPortRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstDefaultSrcPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultSrcPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultSrcPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDefaultSrcPortRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstDefaultSrcPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultSrcPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultSrcPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstDefaultSrcPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultSrcPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultSrcPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstDefaultSrcPortTemplate(d []interface{}) edpt.DdosDstDefaultSrcPortTemplate {

	count1 := len(d)
	var ret edpt.DdosDstDefaultSrcPortTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcUdp = in["src_udp"].(string)
		ret.SrcTcp = in["src_tcp"].(string)
	}
	return ret
}

func dataToEndpointDdosDstDefaultSrcPort(d *schema.ResourceData) edpt.DdosDstDefaultSrcPort {
	var ret edpt.DdosDstDefaultSrcPort
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Template = getObjectDdosDstDefaultSrcPortTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DefaultAddressType = d.Get("default_address_type").(string)
	return ret
}
