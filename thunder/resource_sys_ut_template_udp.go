package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtTemplateUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_template_udp`: UDP header\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtTemplateUdpCreate,
		UpdateContext: resourceSysUtTemplateUdpUpdate,
		ReadContext:   resourceSysUtTemplateUdpRead,
		DeleteContext: resourceSysUtTemplateUdpDelete,

		Schema: map[string]*schema.Schema{
			"checksum": {
				Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
			},
			"dest_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
			},
			"dest_port_value": {
				Type: schema.TypeInt, Optional: true, Description: "Dest port value",
			},
			"length": {
				Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at UDP header",
			},
			"nat_pool": {
				Type: schema.TypeString, Optional: true, Description: "Nat pool port",
			},
			"src_port_range": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_port_start": {
							Type: schema.TypeInt, Optional: true, Description: "Source port value",
						},
						"src_port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Src port end value",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSysUtTemplateUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtTemplateUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtTemplateUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtTemplateUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSysUtTemplateUdpSrcPortRange(d []interface{}) []edpt.SysUtTemplateUdpSrcPortRange {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateUdpSrcPortRange, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateUdpSrcPortRange
		oi.SrcPortStart = in["src_port_start"].(int)
		oi.SrcPortEnd = in["src_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysUtTemplateUdp(d *schema.ResourceData) edpt.SysUtTemplateUdp {
	var ret edpt.SysUtTemplateUdp
	ret.Inst.Checksum = d.Get("checksum").(string)
	ret.Inst.DestPort = d.Get("dest_port").(int)
	ret.Inst.DestPortValue = d.Get("dest_port_value").(int)
	ret.Inst.Length = d.Get("length").(int)
	ret.Inst.NatPool = d.Get("nat_pool").(string)
	ret.Inst.SrcPortRange = getSliceSysUtTemplateUdpSrcPortRange(d.Get("src_port_range").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
