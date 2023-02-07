package thunder

import (
    "context"
    "util"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLoggingHostIpv4addr() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceLoggingHostIpv4addrCreate,
        UpdateContext: resourceLoggingHostIpv4addrUpdate,
        ReadContext: resourceLoggingHostIpv4addrRead,
        DeleteContext: resourceLoggingHostIpv4addrDelete,
        Schema: map[string]*schema.Schema {
            "host_ipv4": {
                Type: schema.TypeString, Required: true, ForceNew: true, Description: "Set syslog host ip address",
            },
            "use_mgmt_port": {
                Type: schema.TypeInt, Optional: true, Description: "Use management port for connections",
                ValidateFunc: validation.IntBetween(0, 1),
            },
            "port": {
                Type: schema.TypeInt, Optional: true, Default: 514, Description: "Set remote syslog port number",
                ValidateFunc: validation.IntBetween(1, 32767),
            },
            "tcp": {
                Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use TCP as transport protocol",
                ValidateFunc: validation.IntBetween(0, 1),
            },
            "over_tls": {
                Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable remote logging over TLS session",
                ValidateFunc: validation.IntBetween(0, 1),
            },
            "uuid": {
                Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
            },
        },
    }
}

func resourceLoggingHostIpv4addrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostIpv4addrCreate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointLoggingHostIpv4addr(d)
        d.SetId(obj.GetId())
        err := PostEndpointLoggingHostIpv4addr(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceLoggingHostIpv4addrRead(ctx, d, meta)
    }
    return diags
}

func resourceLoggingHostIpv4addrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostIpv4addrRead()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        _,err := GetEndpointLoggingHostIpv4addr(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func resourceLoggingHostIpv4addrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostIpv4addrUpdate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointLoggingHostIpv4addr(d)
        err := PutEndpointLoggingHostIpv4addr(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceLoggingHostIpv4addrRead(ctx, d, meta)
    }
    return diags
}

func resourceLoggingHostIpv4addrDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostIpv4addrDelete()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        err := DeleteEndpointLoggingHostIpv4addr(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func dataToEndpointLoggingHostIpv4addr(d *schema.ResourceData) EndpointLoggingHostIpv4addr{
    var ret EndpointLoggingHostIpv4addr
    ret.Inst.HostIpv4 = d.Get("host_ipv4").(string)
    ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
    ret.Inst.Port = d.Get("port").(int)
    ret.Inst.Tcp = d.Get("tcp").(int)
    ret.Inst.OverTls = d.Get("over_tls").(int)
    //omit uuid
    return ret
}
