package thunder

//Thunder resource SnmpServerCommunityRead

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSnmpServerCommunityRead() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerCommunityReadCreate,
		Update: resourceSnmpServerCommunityReadUpdate,
		Read:   resourceSnmpServerCommunityReadRead,
		Delete: resourceSnmpServerCommunityReadDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"remote": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_host": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ipv4_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ipv4_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_host": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ipv4_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ipv6_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_host": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ipv6_mask": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"oid_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"remote": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"host_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_host": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"ipv4_mask": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"ipv4_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_host": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"ipv4_mask": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"ipv6_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv6_host": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"ipv6_mask": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
								},
							},
						},
						"oid_val": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
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
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerCommunityReadCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerCommunityRead (Inside resourceSnmpServerCommunityReadCreate) ")
		name1 := d.Get("user").(string)
		data := dataToSnmpServerCommunityRead(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerCommunityRead --")
		d.SetId(name1)
		go_thunder.PostSnmpServerCommunityRead(client.Token, data, client.Host)

		return resourceSnmpServerCommunityReadRead(d, meta)

	}
	return nil
}

func resourceSnmpServerCommunityReadRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerCommunityRead (Inside resourceSnmpServerCommunityReadRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerCommunityRead(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerCommunityReadUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SnmpServerCommunityRead   (Inside resourceSnmpServerCommunityReadUpdate) ")
		data := dataToSnmpServerCommunityRead(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerCommunityRead ")
		go_thunder.PutSnmpServerCommunityRead(client.Token, name1, data, client.Host)

		return resourceSnmpServerCommunityReadRead(d, meta)

	}
	return nil
}

func resourceSnmpServerCommunityReadDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerCommunityReadDelete) " + name1)
		err := go_thunder.DeleteSnmpServerCommunityRead(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSnmpServerCommunityRead(d *schema.ResourceData) go_thunder.SnmpServerCommunityRead {
	var vc go_thunder.SnmpServerCommunityRead
	var c go_thunder.SnmpServerCommunityReadInstance

	var obj1 go_thunder.SnmpServerCommunityRemote
	prefix1 := "remote.0."

	HostListCount := d.Get(prefix1 + "host_list.#").(int)
	obj1.DNSHost = make([]go_thunder.SnmpServerCommunityHostList, 0, HostListCount)

	for i := 0; i < HostListCount; i++ {
		var obj1_1 go_thunder.SnmpServerCommunityHostList
		prefix1_1 := prefix1 + fmt.Sprintf("host_list.%d.", i)
		obj1_1.DNSHost = d.Get(prefix1_1 + "dns_host").(string)
		obj1_1.Ipv4Mask = d.Get(prefix1_1 + "ipv4_mask").(string)
		obj1.DNSHost = append(obj1.DNSHost, obj1_1)
	}

	Ipv4ListCount := d.Get("ipv4_list.#").(int)
	obj1.Ipv4Host = make([]go_thunder.SnmpServerCommunityIpv4List, 0, Ipv4ListCount)

	for i := 0; i < Ipv4ListCount; i++ {
		var obj1_2 go_thunder.SnmpServerCommunityIpv4List
		prefix1_2 := prefix1 + fmt.Sprintf("ipv4_list.%d.", i)
		obj1_2.Ipv4Host = d.Get(prefix1_2 + "ipv4_host").(string)
		obj1_2.Ipv4Mask = d.Get(prefix1_2 + "ipv4_mask").(string)
		obj1.Ipv4Host = append(obj1.Ipv4Host, obj1_2)
	}

	Ipv6ListCount := d.Get("ipv6_list.#").(int)
	obj1.Ipv6Host = make([]go_thunder.SnmpServerCommunityIpv6List, 0, Ipv6ListCount)

	for i := 0; i < Ipv6ListCount; i++ {
		var obj1_3 go_thunder.SnmpServerCommunityIpv6List
		prefix1_3 := prefix1 + fmt.Sprintf("ipv6_list.%d.", i)
		obj1_3.Ipv6Host = d.Get(prefix1_3 + "ipv6_host").(string)
		obj1_3.Ipv6Mask = d.Get(prefix1_3 + "ipv6_mask").(int)
		obj1.Ipv6Host = append(obj1.Ipv6Host, obj1_3)
	}

	c.HostList = obj1

	OidListCount := d.Get("oid_list.#").(int)
	c.OidVal = make([]go_thunder.SnmpServerCommunityOidList, 0, OidListCount)

	for i := 0; i < OidListCount; i++ {
		var obj2 go_thunder.SnmpServerCommunityOidList
		prefix2 := fmt.Sprintf("oid_list.%d.", i)

		var obj2_1 go_thunder.SnmpServerCommunityRemote
		prefix2_1 := "remote.0."

		HostListCount := d.Get(prefix2_1 + "host_list.#").(int)
		obj1.DNSHost = make([]go_thunder.SnmpServerCommunityHostList, 0, HostListCount)

		for i := 0; i < HostListCount; i++ {
			var obj2_1_1 go_thunder.SnmpServerCommunityHostList
			prefix2_1_1 := prefix2_1 + fmt.Sprintf("host_list.%d.", i)
			obj2_1_1.DNSHost = d.Get(prefix2_1_1 + "dns_host").(string)
			obj2_1_1.Ipv4Mask = d.Get(prefix2_1_1 + "ipv4_mask").(string)
			obj2_1.DNSHost = append(obj2_1.DNSHost, obj2_1_1)
		}

		Ipv4ListCount := d.Get(prefix2_1 + "ipv4_list.#").(int)
		obj2_1.Ipv4Host = make([]go_thunder.SnmpServerCommunityIpv4List, 0, Ipv4ListCount)

		for i := 0; i < Ipv4ListCount; i++ {
			var obj2_1_2 go_thunder.SnmpServerCommunityIpv4List
			prefix2_1_2 := prefix2_1 + fmt.Sprintf("ipv4_list.%d.", i)
			obj2_1_2.Ipv4Host = d.Get(prefix2_1_2 + "ipv4_host").(string)
			obj2_1_2.Ipv4Mask = d.Get(prefix2_1_2 + "ipv4_mask").(string)
			obj2_1.Ipv4Host = append(obj2_1.Ipv4Host, obj2_1_2)
		}

		Ipv6ListCount := d.Get(prefix2_1 + "ipv6_list.#").(int)
		obj2_1.Ipv6Host = make([]go_thunder.SnmpServerCommunityIpv6List, 0, Ipv6ListCount)

		for i := 0; i < Ipv6ListCount; i++ {
			var obj2_1_3 go_thunder.SnmpServerCommunityIpv6List
			prefix2_1_3 := prefix2_1 + fmt.Sprintf("ipv6_list.%d.", i)
			obj2_1_3.Ipv6Host = d.Get(prefix2_1_3 + "ipv6_host").(string)
			obj2_1_3.Ipv6Mask = d.Get(prefix2_1_3 + "ipv6_mask").(int)
			obj2_1.Ipv6Host = append(obj2_1.Ipv6Host, obj2_1_3)
		}

		obj2.HostList = obj2_1

		obj2.OidVal = d.Get(prefix2 + "oid_val").(string)
		obj2.UserTag = d.Get(prefix2 + "user_tag").(string)
		c.OidVal = append(c.OidVal, obj2)
	}

	c.UserTag = d.Get("user_tag").(string)
	c.User = d.Get("user").(string)

	vc.UUID = c
	return vc
}
