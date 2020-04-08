package vthunder

//vThunder resource overlay tunnel vtep

import (
	"fmt"
	"log"
	"strconv"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOverlayTunnelVtep() *schema.Resource {
	return &schema.Resource{
		Create: resourceVtepCreate,
		Update: resourceVtepUpdate,
		Read:   resourceVtepRead,
		Delete: resourceVtepDelete,

		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vni": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"destination_vtep": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
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
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"destination_ip_address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
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
					},
				},
			},
			"source_ip_address": {
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
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"encap": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"id2": {
				Type:        schema.TypeInt,
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceVtepCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := strconv.Itoa(d.Get("id2").(int))
		logger.Println("[INFO] Creating vtep (Inside resourceVtepCreate    " + name)
		v := dataToVtep(d)
		d.SetId(name)
		go_vthunder.PostVtep(client.Token, v, client.Host)

		return resourceVtepRead(d, meta)
	}
	return nil
}

func resourceVtepRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading Vtep (Inside resourceVtepRead)")

	if client.Host != "" {
		client := meta.(vThunder)

		name := d.Id()

		logger.Println("[INFO] Fetching Vtep Read" + name)

		vtep, err := go_vthunder.GetVtep(client.Token, name, client.Host)

		if vtep == nil {
			logger.Println("[INFO] No vtep found " + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceVtepUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := strconv.Itoa(d.Get("id2").(int))
		logger.Println("[INFO] Modifying Vtep (Inside resourceVtepUpdate    " + name)
		v := dataToVtep(d)
		d.SetId(name)
		go_vthunder.PutVtep(client.Token, name, v, client.Host)

		return resourceVtepRead(d, meta)
	}
	return nil
}

func resourceVtepDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting vtep (Inside resourceVtepDelete) " + name)

		err := go_vthunder.DeleteVtep(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete vtep  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate vtep structure
func dataToVtep(d *schema.ResourceData) go_vthunder.Vtep {
	var s go_vthunder.Vtep

	var sInstance go_vthunder.VtepInstance

	sInstance.UUID = d.Get("uuid").(string)
	sInstance.UserTag = d.Get("user_tag").(string)
	sInstance.Encap = d.Get("encap").(string)
	sInstance.ID = d.Get("id2").(int)

	samplingCount := d.Get("sampling_enable.#").(int)
	sInstance.Counters1 = make([]go_vthunder.SamplingEnableOT, 0, samplingCount)
	for i := 0; i < samplingCount; i++ {
		var samp go_vthunder.SamplingEnableOT
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		samp.Counters1 = d.Get(prefix + "counters1").(string)
		sInstance.Counters1 = append(sInstance.Counters1, samp)
	}

	hostListCount := d.Get("host_list.#").(int)
	sInstance.IPAddr = make([]go_vthunder.HostList, 0, hostListCount)

	for i := 0; i < hostListCount; i++ {
		var host go_vthunder.HostList
		prefix := fmt.Sprintf("host_list.%d.", i)
		host.DestinationVtep = d.Get(prefix + "destination_vtep").(string)
		host.IPAddr = d.Get(prefix + "ip_addr").(string)
		host.OverlayMacAddr = d.Get(prefix + "overlay_mac_addr").(string)
		host.Vni = d.Get(prefix + "vni").(int)
		host.UUID = d.Get(prefix + "uuid").(string)
		sInstance.IPAddr = append(sInstance.IPAddr, host)
	}

	destAddrCnt := d.Get("destination_ip_address_list.#").(int)
	sInstance.Segment = make([]go_vthunder.DestinationIPAddressList, 0, destAddrCnt)
	for i := 0; i < destAddrCnt; i++ {
		var destAddrr go_vthunder.DestinationIPAddressList
		prefix := fmt.Sprintf("destination_ip_address_list.%d.", i)
		destAddrr.UUID = d.Get(prefix + "uuid").(string)
		destAddrr.IPAddress = d.Get(prefix + "ip_address").(string)
		destAddrr.UserTag = d.Get(prefix + "user_tag").(string)
		destAddrr.Encap = d.Get(prefix + "encap").(string)
		sInstance.Segment = append(sInstance.Segment, destAddrr)
	}

	s.UUID = sInstance

	return s
}
