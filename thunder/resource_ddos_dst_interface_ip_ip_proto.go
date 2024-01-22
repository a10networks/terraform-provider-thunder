package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstInterfaceIpIpProto() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_interface_ip_ip_proto`: Configure L4 on local interface IPs\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstInterfaceIpIpProtoCreate,
		UpdateContext: resourceDdosDstInterfaceIpIpProtoUpdate,
		ReadContext:   resourceDdosDstInterfaceIpIpProtoRead,
		DeleteContext: resourceDdosDstInterfaceIpIpProtoDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "IP protocol number",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"addr": {
				Type: schema.TypeString, Required: true, Description: "Addr",
			},
		},
	}
}
func resourceDdosDstInterfaceIpIpProtoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpIpProtoCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpIpProto(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstInterfaceIpIpProtoRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstInterfaceIpIpProtoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpIpProtoUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpIpProto(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstInterfaceIpIpProtoRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstInterfaceIpIpProtoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpIpProtoDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpIpProto(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstInterfaceIpIpProtoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpIpProtoRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpIpProto(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstInterfaceIpIpProto(d *schema.ResourceData) edpt.DdosDstInterfaceIpIpProto {
	var ret edpt.DdosDstInterfaceIpIpProto
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Addr = d.Get("addr").(string)
	return ret
}
