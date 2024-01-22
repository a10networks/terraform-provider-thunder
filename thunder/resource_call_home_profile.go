package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCallHomeProfile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_call_home_profile`: Call home profile\n\n__PLACEHOLDER__",
		CreateContext: resourceCallHomeProfileCreate,
		UpdateContext: resourceCallHomeProfileUpdate,
		ReadContext:   resourceCallHomeProfileRead,
		DeleteContext: resourceCallHomeProfileDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "register", Description: "'register': Register the device to the portal; 'deregister': Deregister the device from the portal;",
			},
			"export_policy": {
				Type: schema.TypeString, Optional: true, Default: "restrictive", Description: "'restrictive': Export minimal information about system and config, default policy; 'permissive': Export as much as possible information about system, config and solution data;",
			},
			"ipv4": {
				Type: schema.TypeString, Optional: true, Description: "set IPV4 address for the host",
			},
			"ipv6": {
				Type: schema.TypeString, Optional: true, Description: "set IPV6 address for the host",
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "set hostname url for the host",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "Set port for the call home portal",
			},
			"time": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set export time of the data in minutes. default 0 (12 AM). exported between 12-01 AM",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for connections",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCallHomeProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCallHomeProfileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCallHomeProfile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCallHomeProfileRead(ctx, d, meta)
	}
	return diags
}

func resourceCallHomeProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCallHomeProfileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCallHomeProfile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCallHomeProfileRead(ctx, d, meta)
	}
	return diags
}
func resourceCallHomeProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCallHomeProfileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCallHomeProfile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCallHomeProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCallHomeProfileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCallHomeProfile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCallHomeProfile(d *schema.ResourceData) edpt.CallHomeProfile {
	var ret edpt.CallHomeProfile
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ExportPolicy = d.Get("export_policy").(string)
	ret.Inst.Ipv4 = d.Get("ipv4").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.Time = d.Get("time").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
