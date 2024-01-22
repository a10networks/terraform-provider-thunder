package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMirrorPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_mirror_port`: Enable a port to act as Mirror Port\n\n__PLACEHOLDER__",
		CreateContext: resourceMirrorPortCreate,
		UpdateContext: resourceMirrorPortUpdate,
		ReadContext:   resourceMirrorPortRead,
		DeleteContext: resourceMirrorPortDelete,

		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet port as mirror port (Port Value)",
			},
			"mirror_dir": {
				Type: schema.TypeString, Optional: true, Default: "both", Description: "'input': Mirror incoming packets to this port; 'output': Mirror outgoing packets to this port; 'both': Mirror both incoming and outgoing packets to this port;",
			},
			"mirror_index": {
				Type: schema.TypeInt, Required: true, Description: "Mirror index",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceMirrorPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMirrorPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMirrorPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceMirrorPortRead(ctx, d, meta)
	}
	return diags
}

func resourceMirrorPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMirrorPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMirrorPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceMirrorPortRead(ctx, d, meta)
	}
	return diags
}
func resourceMirrorPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMirrorPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMirrorPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceMirrorPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMirrorPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMirrorPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointMirrorPort(d *schema.ResourceData) edpt.MirrorPort {
	var ret edpt.MirrorPort
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.MirrorDir = d.Get("mirror_dir").(string)
	ret.Inst.MirrorIndex = d.Get("mirror_index").(int)
	//omit uuid
	return ret
}
