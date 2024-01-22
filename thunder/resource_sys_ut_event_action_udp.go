package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtEventActionUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_event_action_udp`: UDP header\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtEventActionUdpCreate,
		UpdateContext: resourceSysUtEventActionUdpUpdate,
		ReadContext:   resourceSysUtEventActionUdpRead,
		DeleteContext: resourceSysUtEventActionUdpDelete,

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
			"src_port": {
				Type: schema.TypeInt, Optional: true, Description: "Source port value",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"event_number": {
				Type: schema.TypeString, Required: true, Description: "EventNumber",
			},
			"direction": {
				Type: schema.TypeString, Required: true, Description: "Direction",
			},
		},
	}
}
func resourceSysUtEventActionUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtEventActionUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtEventActionUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtEventActionUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtEventActionUdp(d *schema.ResourceData) edpt.SysUtEventActionUdp {
	var ret edpt.SysUtEventActionUdp
	ret.Inst.Checksum = d.Get("checksum").(string)
	ret.Inst.DestPort = d.Get("dest_port").(int)
	ret.Inst.DestPortValue = d.Get("dest_port_value").(int)
	ret.Inst.Length = d.Get("length").(int)
	ret.Inst.NatPool = d.Get("nat_pool").(string)
	ret.Inst.SrcPort = d.Get("src_port").(int)
	//omit uuid
	ret.Inst.EventNumber = d.Get("event_number").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
