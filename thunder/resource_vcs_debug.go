package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsDebug() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_debug`: VCS debug\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsDebugCreate,
		UpdateContext: resourceVcsDebugUpdate,
		ReadContext:   resourceVcsDebugRead,
		DeleteContext: resourceVcsDebugDelete,

		Schema: map[string]*schema.Schema{
			"daemon": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Daemon component",
			},
			"daemon_msg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Daemon message component",
			},
			"election": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Election component",
			},
			"election_pdu": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Election pdu component",
			},
			"encoder": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Encoder component",
			},
			"handler": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Handler component",
			},
			"info": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Information component",
			},
			"lib": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Lib component",
			},
			"net": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Net component",
			},
			"ssl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SSL component",
			},
			"util": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Utility component",
			},
			"vblade": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "vBlade component",
			},
			"vblade_msg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "vBlade Message component",
			},
			"vmaster": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "vMaster component",
			},
			"vmaster_msg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "vMaster Message component",
			},
		},
	}
}
func resourceVcsDebugCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDebugCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDebug(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsDebugRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsDebugUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDebugUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDebug(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsDebugRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsDebugDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDebugDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDebug(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsDebugRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDebugRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDebug(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVcsDebug(d *schema.ResourceData) edpt.VcsDebug {
	var ret edpt.VcsDebug
	ret.Inst.Daemon = d.Get("daemon").(int)
	ret.Inst.Daemon_msg = d.Get("daemon_msg").(int)
	ret.Inst.Election = d.Get("election").(int)
	ret.Inst.Election_pdu = d.Get("election_pdu").(int)
	ret.Inst.Encoder = d.Get("encoder").(int)
	ret.Inst.Handler = d.Get("handler").(int)
	ret.Inst.Info = d.Get("info").(int)
	ret.Inst.Lib = d.Get("lib").(int)
	ret.Inst.Net = d.Get("net").(int)
	ret.Inst.Ssl = d.Get("ssl").(int)
	ret.Inst.Util = d.Get("util").(int)
	ret.Inst.Vblade = d.Get("vblade").(int)
	ret.Inst.Vblade_msg = d.Get("vblade_msg").(int)
	ret.Inst.Vmaster = d.Get("vmaster").(int)
	ret.Inst.Vmaster_msg = d.Get("vmaster_msg").(int)
	return ret
}
