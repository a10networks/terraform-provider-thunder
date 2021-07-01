package thunder

//Thunder resource OverlayTunnelVtep

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceOverlayTunnelVtep() *schema.Resource {
	return &schema.Resource{
		Create: resourceOverlayTunnelVtepCreate,
		Update: resourceOverlayTunnelVtepUpdate,
		Read:   resourceOverlayTunnelVtepRead,
		Delete: resourceOverlayTunnelVtepDelete,
		Schema: map[string]*schema.Schema{
			"id1": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"encap": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"local_ip_address": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"vni_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"segment": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"local_ipv6_address": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"remote_ip_address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"encap": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"use_lif": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"partition": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"lif": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"gre_keepalive": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retry_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"retry_count": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"use_gre_key": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gre_key": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"vni_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"segment": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"remote_ipv6_address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"use_lif": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"partition": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"lif": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"host_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"overlay_mac_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"vni": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"remote_vtep": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceOverlayTunnelVtepCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating OverlayTunnelVtep (Inside resourceOverlayTunnelVtepCreate) ")
		name1 := d.Get("id1").(int)
		data := dataToOverlayTunnelVtep(d)
		logger.Println(d)
		logger.Println("[INFO] received formatted data from method data to OverlayTunnelVtep --")
		d.SetId(strconv.Itoa(name1))
		go_thunder.PostOverlayTunnelVtep(client.Token, data, client.Host)

		return resourceOverlayTunnelVtepRead(d, meta)

	}
	return nil
}

func resourceOverlayTunnelVtepRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading OverlayTunnelVtep (Inside resourceOverlayTunnelVtepRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetOverlayTunnelVtep(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceOverlayTunnelVtepUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying OverlayTunnelVtep   (Inside resourceOverlayTunnelVtepUpdate) ")
		data := dataToOverlayTunnelVtep(d)
		logger.Println("[INFO] received formatted data from method data to OverlayTunnelVtep ")
		go_thunder.PutOverlayTunnelVtep(client.Token, name1, data, client.Host)

		return resourceOverlayTunnelVtepRead(d, meta)

	}
	return nil
}

func resourceOverlayTunnelVtepDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceOverlayTunnelVtepDelete) " + name1)
		err := go_thunder.DeleteOverlayTunnelVtep(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToOverlayTunnelVtep(d *schema.ResourceData) go_thunder.OverlayTunnelVtep {
	var vc go_thunder.OverlayTunnelVtep
	var c go_thunder.OverlayTunnelVtepInstance
	c.ID = d.Get("id1").(int)
	c.Encap = d.Get("encap").(string)
	c.UserTag = d.Get("user_tag").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.OverlayTunnelSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.OverlayTunnelSamplingEnable
		prefix1 := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix1 + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	var obj2 go_thunder.OverlayTunnelLocalIPAddress

	//obj2 := &obj2_val
	prefix2 := "local_ip_address.0."
	obj2.IPAddress = d.Get(prefix2 + "ip_address").(string)

	VniListCount := d.Get(prefix2 + "vni_list.#").(int)
	obj2.Segment = make([]go_thunder.OverlayTunnelVniList, 0, VniListCount)

	for i := 0; i < VniListCount; i++ {
		var obj2_1 go_thunder.OverlayTunnelVniList
		prefix2_1 := fmt.Sprintf(prefix2+"vni_list.%d.", i)
		obj2_1.Segment = d.Get(prefix2_1 + "segment").(int)
		obj2.Segment = append(obj2.Segment, obj2_1)
	}

	c.IPAddress = obj2

	var obj3 go_thunder.OverlayTunnelLocalIpv6Address

	//obj3 := &obj3_val
	prefix3 := "local_ipv6_address.0."
	obj3.Ipv6Address = d.Get(prefix3 + "ipv6_address").(string)

	c.Ipv6Address = obj3

	RemoteIpAddressListCount := d.Get("remote_ip_address_list.#").(int)
	c.GreKey = make([]go_thunder.OverlayTunnelRemoteIPAddressList, 0, RemoteIpAddressListCount)

	for i := 0; i < RemoteIpAddressListCount; i++ {
		var obj4 go_thunder.OverlayTunnelRemoteIPAddressList
		prefix4 := fmt.Sprintf("remote_ip_address_list.%d.", i)
		obj4.IPAddress = d.Get(prefix4 + "ip_address").(string)
		obj4.Encap = d.Get(prefix4 + "encap").(string)
		obj4.UserTag = d.Get(prefix4 + "user_tag").(string)

		var obj4_1 go_thunder.OverlayTunnelUseLif
		prefix4_1 := prefix4 + "use_lif.0."
		obj4_1.Partition = d.Get(prefix4_1 + "partition").(string)
		obj4_1.Lif = d.Get(prefix4_1 + "lif").(string)

		obj4.Partition = obj4_1

		var obj4_2 go_thunder.OverlayTunnelGreKeepalive
		prefix4_2 := prefix4 + "gre_keepalive.0."
		obj4_2.RetryTime = d.Get(prefix4_2 + "retry_time").(int)
		obj4_2.RetryCount = d.Get(prefix4_2 + "retry_count").(int)

		obj4.RetryTime = obj4_2

		var obj4_3 go_thunder.OverlayTunnelUseGreKey
		prefix4_3 := prefix4 + "use_gre_key.0."
		obj4_3.GreKey = d.Get(prefix4_3 + "gre_key").(int)

		obj4.GreKey = obj4_3

		VniListCount := d.Get(prefix4 + "vni_list.#").(int)
		obj4.Segment = make([]go_thunder.OverlayTunnelVniList, 0, VniListCount)

		for i := 0; i < VniListCount; i++ {
			var obj4_4 go_thunder.OverlayTunnelVniList
			prefix4_4 := fmt.Sprintf(prefix4+"vni_list.%d.", i)
			obj4_4.Segment = d.Get(prefix4_4 + "segment").(int)
			obj4.Segment = append(obj4.Segment, obj4_4)
		}

		c.GreKey = append(c.GreKey, obj4)
	}

	RemoteIpv6AddressListCount := d.Get("remote_ipv6_address_list.#").(int)
	c.Partition = make([]go_thunder.OverlayTunnelRemoteIpv6AddressList, 0, RemoteIpv6AddressListCount)

	for i := 0; i < RemoteIpv6AddressListCount; i++ {
		var obj5 go_thunder.OverlayTunnelRemoteIpv6AddressList
		prefix5 := fmt.Sprintf("remote_ipv6_address_list.%d.", i)
		obj5.Ipv6Address = d.Get(prefix5 + "ipv6_address").(string)
		obj5.UserTag = d.Get(prefix5 + "user_tag").(string)

		var obj5_1 go_thunder.OverlayTunnelUseLif
		prefix5_1 := prefix5 + "use_lif.0."
		obj5_1.Partition = d.Get(prefix5_1 + "partition").(string)
		obj5_1.Lif = d.Get(prefix5_1 + "lif").(string)

		obj5.Partition = obj5_1

		c.Partition = append(c.Partition, obj5)
	}

	HostListCount := d.Get("host_list.#").(int)
	c.IPAddr = make([]go_thunder.OverlayTunnelHostList, 0, HostListCount)

	for i := 0; i < HostListCount; i++ {
		var obj6 go_thunder.OverlayTunnelHostList
		prefix6 := fmt.Sprintf("host_list.%d.", i)
		obj6.IPAddr = d.Get(prefix6 + "ip_addr").(string)
		obj6.OverlayMacAddr = d.Get(prefix6 + "overlay_mac_addr").(string)
		obj6.Vni = d.Get(prefix6 + "vni").(int)
		obj6.RemoteVtep = d.Get(prefix6 + "remote_vtep").(string)
		c.IPAddr = append(c.IPAddr, obj6)
	}

	vc.ID = c
	return vc
}
