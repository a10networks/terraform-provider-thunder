package thunder

import (
    "context"
    "util"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
)

func resourceIpRouteRib() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceIpRouteRibCreate,
        UpdateContext: resourceIpRouteRibUpdate,
        ReadContext: resourceIpRouteRibRead,
        DeleteContext: resourceIpRouteRibDelete,
        Schema: map[string]*schema.Schema{
            "ip_dest_addr": {
                Type: schema.TypeString, Required: true, ForceNew: true, Description: "Destination prefix",
                ValidateFunc: validation.IsIPv4Address,
            },
            "ip_mask": {
                Type: schema.TypeString, Required: true, ForceNew: true, Description: "Destination prefix mask",
                ValidateFunc: axapi.IsIpv4NetmaskBrief,
            },
            "ip_nexthop_ipv4": {
                Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "ip_next_hop": {
                            Type: schema.TypeString, Optional: true, Description: "Forwarding router's address",
                            ValidateFunc: validation.IsIPv4Address,
                        },
                        "distance_nexthop_ip": {
                            Type: schema.TypeInt, Default: 1, Optional: true, Description: "Distance value for this route",
                            ValidateFunc: validation.IntBetween(1, 255),
                        },
                        "description_nexthop_ip": {
                            Type: schema.TypeString, Optional: true, Description: "Description for static route",
                            ValidateFunc: validation.StringLenBetween(1, 63),
                        },
                    },
                },
            },
            "ip_nexthop_lif": {
                Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "lif": {
                            Type: schema.TypeString, Optional: true, Description: "LIF Interface (Logical tunnel interface name)",
                            ValidateFunc: validation.StringLenBetween(1, 15),
                        },
                        "description_nexthop_lif": {
                            Type: schema.TypeString, Optional: true, Description: "Description for static route",
                            ValidateFunc: validation.StringLenBetween(1, 63),
                        },
                    },
                },
            },
            "ip_nexthop_tunnel": {
                Type: schema.TypeList, MaxItems: 1, Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "tunnel": {
                            Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
                            ValidateFunc: validation.IntBetween(1, 128),
                        },
                        "ip_next_hop_tunnel": {
                            Type: schema.TypeString, Optional: true, Description: "Forwarding router's address",
                            ValidateFunc: validation.IsIPv4Address,
                        },
                        "distance_nexthop_tunnel": {
                            Type: schema.TypeInt, Optional: true, Default:1, Description: "Distance value for this route",
                            ValidateFunc: validation.IntBetween(1, 128),
                        },
                        "description_nexthop_tunnel": {
                            Type: schema.TypeString, Optional: true, Description: "Description for static route",
                            ValidateFunc: validation.StringLenBetween(1, 63),
                        },
                    },
                },
            },
            "ip_nexthop_partition": {
                Type: schema.TypeList, MaxItems: 1, Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "partition_name": {
                            Type: schema.TypeString, Optional: true, Description: "Name of network partition",
                            ValidateFunc: validation.StringLenBetween(1, 14),
                        },
                        "vrid_num_in_partition": {
                            Type: schema.TypeInt, Optional: true, Description: "Specify ha VRRP_A vrid",
                            ValidateFunc: validation.IntBetween(0, 31),
                        },
                        "description_nexthop_partition": {
                            Type: schema.TypeString, Optional: true, Description: "Description for static route",
                            ValidateFunc: validation.StringLenBetween(1, 63),
                        },
                        "description_partition_vrid": {
                            Type: schema.TypeString, Optional: true, Description: "Description for static route",
                            ValidateFunc: validation.StringLenBetween(1, 63),
                        },
                    },
                },
            },
            "uuid": {Type: schema.TypeString, Computed: true, Description: "uuid of the object"},
        },
    }
}

func resourceIpRouteRibCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpRouteRibCreate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointIpRouteRib(d)
        d.SetId(obj.GetId())
        err := PostEndpointIpRouteRib(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceIpRouteRibRead(ctx, d, meta)
    }
    return diags
}

func resourceIpRouteRibRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpRouteRibRead()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        instId := d.Id()
        _,err := GetEndpointIpRouteRib(client.Token, client.Host, instId)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func resourceIpRouteRibUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpRouteRibUpdate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointIpRouteRib(d)
        err := PutEndpointIpRouteRib(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceIpRouteRibRead(ctx, d, meta)
    }
    return diags
}

func resourceIpRouteRibDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpRouteRibDelete()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        instId := d.Id()
        err := DeleteEndpointIpRouteRib(client.Token, client.Host, instId)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func dataToEndpointIpRouteRib(d *schema.ResourceData) EndpointIpRouteRib{
    var ret EndpointIpRouteRib
    ret.Inst.IpDestAddr = d.Get("ip_dest_addr").(string)
    ret.Inst.IpMask = d.Get("ip_mask").(string)
    ipNexthopIpv4 := IpRouteRibIpNexthopIpv4{
        IpNextHop: d.Get("ip_nexthop_ipv4.0.ip_next_hop").(string),
        DistanceNexthopIp: d.Get("ip_nexthop_ipv4.0.distance_nexthop_ip").(int),
        DescriptionNexthopIp: d.Get("ip_nexthop_ipv4.0.description_nexthop_ip").(string),
    }
    if (ipNexthopIpv4 != IpRouteRibIpNexthopIpv4{}) {
        ret.Inst.IpNexthopIpv4 = append(ret.Inst.IpNexthopIpv4, ipNexthopIpv4)
    }

    ipNexthopLif := IpRouteRibIpNexthopLif{
        Lif: d.Get("ip_nexthop_lif.0.lif").(string),
        DescriptionNexthopLif: d.Get("ip_nexthop_lif.0.description_nexthop_lif").(string),
    }
    if (IpRouteRibIpNexthopLif{} != ipNexthopLif) {
        ret.Inst.IpNexthopLif = append(ret.Inst.IpNexthopLif, ipNexthopLif)
    }

    ipNexthopTunnel := IpRouteRibIpNexthopTunnel{
        Tunnel: d.Get("ip_nexthop_tunnel.0.tunnel").(int),
        IpNextHopTunnel: d.Get("ip_nexthop_tunnel.0.ip_next_hop_tunnel").(string),
        DistanceNexthopTunnel: d.Get("ip_nexthop_tunnel.0.distance_nexthop_tunnel").(int),
        DescriptionNexthopTunnel: d.Get("ip_nexthop_tunnel.0.description_nexthop_tunnel").(string),
    }
    if (IpRouteRibIpNexthopTunnel{} != ipNexthopTunnel) {
        ret.Inst.IpNexthopTunnel = append(ret.Inst.IpNexthopTunnel, ipNexthopTunnel)
    }

    ipNexthopPartition := IpRouteRibIpNexthopPartition{
        PartitionName: d.Get("ip_nexthop_partition.0.partition_name").(string),
        VridNumInPartition: d.Get("ip_nexthop_partition.0.vrid_num_in_partition").(int),
        DescriptionNexthopPartition: d.Get("ip_nexthop_partition.0.description_nexthop_partition").(string),
        DescriptionPartitionVrid: d.Get("ip_nexthop_partition.0.description_partition_vrid").(string),
    }
    if (IpRouteRibIpNexthopPartition{} != ipNexthopPartition) {
        ret.Inst.IpNexthopPartition = append(ret.Inst.IpNexthopPartition, ipNexthopPartition)
    }
    //omit uuid
    return ret
}
