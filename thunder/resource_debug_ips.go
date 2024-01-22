





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceDebugIps() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_debug_ips`: Debug IPS\n\n__PLACEHOLDER__",
        CreateContext: resourceDebugIpsCreate,
        UpdateContext: resourceDebugIpsUpdate,
        ReadContext: resourceDebugIpsRead,
        DeleteContext: resourceDebugIpsDelete,

        Schema: map[string]*schema.Schema{
        "level": {
        Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-4)",
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
    }
}
func resourceDebugIpsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDebugIpsCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDebugIps(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceDebugIpsRead(ctx, d, meta)
    }
    return diags
}

func resourceDebugIpsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDebugIpsUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDebugIps(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceDebugIpsRead(ctx, d, meta)
    }
    return diags
}
func resourceDebugIpsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDebugIpsDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDebugIps(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceDebugIpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDebugIpsRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDebugIps(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointDebugIps(d *schema.ResourceData) edpt.DebugIps {
    var ret edpt.DebugIps
	ret.Inst.Level = d.Get("level").(int)
	//omit uuid
    return ret
}

