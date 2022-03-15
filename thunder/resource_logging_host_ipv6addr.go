package thunder

import (
    "context"
    "util"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLoggingHostIpv6addr() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceLoggingHostIpv6addrCreate,
        UpdateContext: resourceLoggingHostIpv6addrUpdate,
        ReadContext: resourceLoggingHostIpv6addrRead,
        DeleteContext: resourceLoggingHostIpv6addrDelete,
        Schema: map[string]*schema.Schema {
            "host_ipv6":{
                Type: schema.TypeString, Required: true, ForceNew: true, Description: "Set syslog host ipv6 address",
                ValidateFunc: validation.IsIPv6Address,
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

func resourceLoggingHostIpv6addrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostIpv6addrCreate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointLoggingHostIpv6addr(d)
        d.SetId(obj.GetId())
        err := PostEndpointLoggingHostIpv6addr(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceLoggingHostIpv6addrRead(ctx, d, meta)
    }
    return diags
}

func resourceLoggingHostIpv6addrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostIpv6addrRead()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        _,err := GetEndpointLoggingHostIpv6addr(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func resourceLoggingHostIpv6addrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostIpv6addrUpdate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointLoggingHostIpv6addr(d)
        err := PutEndpointLoggingHostIpv6addr(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceLoggingHostIpv6addrRead(ctx, d, meta)
    }
    return diags
}

func resourceLoggingHostIpv6addrDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceLoggingHostIpv6addrDelete()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        err := DeleteEndpointLoggingHostIpv6addr(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func dataToEndpointLoggingHostIpv6addr(d *schema.ResourceData) EndpointLoggingHostIpv6addr{
    var ret EndpointLoggingHostIpv6addr
    ret.Inst.HostIpv6 = d.Get("host_ipv6").(string)
    ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
    ret.Inst.Port = d.Get("port").(int)
    ret.Inst.Tcp = d.Get("tcp").(int)
    ret.Inst.OverTls = d.Get("over_tls").(int)

    //omit uuid
    return ret
}
