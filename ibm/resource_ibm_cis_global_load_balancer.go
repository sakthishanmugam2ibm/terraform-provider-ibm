package ibm

import (
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
)

func resourceIBMCISGlb() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			"domain_id": {
				Type:        schema.TypeString,
				Description: "Associated CIS domain",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "name",
				Required:    true,
			},
			"fallback_pool_id": {
				Type:        schema.TypeString,
				Description: "fallback pool ID",
				Required:    true,
			},
			"default_pool_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of default Pool IDs",
				//ValidateFunc: validation.StringLenBetween(1, 32),
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description for the load balancer instance",
				//ValidateFunc: validation.StringLenBetween(0, 1024),
			},
			"ttl": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"proxied"},
				Description:   "TTL value", // this is set to zero regardless of config when proxied=true

			},
			"proxied": {
				Type:          schema.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"ttl"},
				Description:   "set to true if proxy needs to be enabled",
			},
			"session_affinity": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "none",
				// Set to cookie when proxy=true
				ValidateFunc: validateAllowedStringValue([]string{"none", "cookie"}),
				Description:  "Session affinity info",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "set to true of LB needs to enabled",
			},
			// "region_pools": &schema.Schema{
			// 	Type:     schema.TypeMap,
			// 	Optional: true,
			// 	Computed: true,
			// 	Elem:     &schema.Schema{Type: schema.TypeString},
			// },
			// "pop_pools": &schema.Schema{
			// 	Type:     schema.TypeMap,
			// 	Optional: true,
			// 	Computed: true,
			// 	Elem:     &schema.Schema{Type: schema.TypeString},
			// },
			"created_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load balancer creation date",
			},
			"modified_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load balancer modified date",
			},
		},

		Create: resourceCISGlbCreate,
		Read:   resourceCISGlbRead,
		Update: resourceCISGlbUpdate,
		Delete: resourceCISGlbDelete,
		// No Exists due to errors in CIS API returning incorrect return codes on 404
		Importer: &schema.ResourceImporter{},
	}
}

func resourceCISGlbCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}

	cisId := d.Get("cis_id").(string)
	name := d.Get("name").(string)
	zoneId, _, err := convertTftoCisTwoVar(d.Get("domain_id").(string))
	if err != nil {
		return err
	}

	var glb *v1.Glb
	var glbObj v1.Glb
	glbNew := v1.GlbBody{}
	glbNew.Name = name

	tfDefaultPools := expandStringList(d.Get("default_pool_ids").(*schema.Set).List())
	defaultPoolIds, _, err := convertTfToCisTwoVarSlice(tfDefaultPools)
	glbNew.DefaultPools = defaultPoolIds
	// glbNew.RegionPools
	// glbNew.PopPools

	fbPoolId := d.Get("fallback_pool_id").(string)
	glbNew.FallbackPool, _, err = convertTftoCisTwoVar(fbPoolId)
	glbNew.Proxied = d.Get("proxied").(bool)
	glbNew.SessionAffinity = d.Get("session_affinity").(string)

	if description, ok := d.GetOk("description"); ok {
		glbNew.Desc = description.(string)
	}
	if ttl, ok := d.GetOk("ttl"); ok {
		glbNew.Ttl = ttl.(int)
	}
	glb, err = cisClient.Glbs().CreateGlb(cisId, zoneId, glbNew)
	if err != nil {
		log.Printf("CreateGlbs Failed %s\n", err)
		return err
	}
	glbObj = *glb
	d.SetId(convertCisToTfThreeVar(glbObj.Id, zoneId, cisId))

	return resourceCISGlbUpdate(d, meta)
}

func resourceCISGlbRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	// Extract CIS Ids from TF Id
	glbId, zoneId, cisId, err := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}
	var glb *v1.Glb

	glb, err = cisClient.Glbs().GetGlb(cisId, zoneId, glbId)
	if err != nil {
		if checkCisGlbDeleted(d, meta, err, glb) {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during GlbRead %v\n", err)
		return err
	}
	glbObj := *glb
	d.Set("cis_id", cisId)
	d.Set("domain_id", convertCisToTfTwoVar(zoneId, cisId))
	d.Set("name", glbObj.Name)
	d.Set("default_pool_ids", convertCisToTfTwoVarSlice(glbObj.DefaultPools, cisId))
	d.Set("description", glbObj.Desc)
	d.Set("fallback_pool_id", convertCisToTfTwoVar(glbObj.FallbackPool, cisId))
	d.Set("ttl", glbObj.Ttl)
	d.Set("proxied", glbObj.Proxied)
	d.Set("enabled", glbObj.Enabled)
	d.Set("session_affinity", glbObj.SessionAffinity)

	return nil
}

func resourceCISGlbUpdate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	// Extract CIS Ids from TF Id
	glbId, zoneId, cisId, err := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}
	glbUpdate := v1.GlbBody{}

	if d.HasChange("name") || d.HasChange("default_pool_ids") || d.HasChange("fallback_pool_id") || d.HasChange("proxied") || d.HasChange("session_affinity") || d.HasChange("description") || d.HasChange("ttl") || d.HasChange("enabled") {

		name := d.Get("name").(string)
		glbUpdate.Name = name
		tfDefaultPools := expandStringList(d.Get("default_pool_ids").(*schema.Set).List())
		defaultPoolIds, _, err := convertTfToCisTwoVarSlice(tfDefaultPools)
		glbUpdate.DefaultPools = defaultPoolIds
		fbPoolId := d.Get("fallback_pool_id").(string)
		glbUpdate.FallbackPool, _, err = convertTftoCisTwoVar(fbPoolId)
		glbUpdate.Proxied = d.Get("proxied").(bool)
		glbUpdate.SessionAffinity = d.Get("session_affinity").(string)

		if description, ok := d.GetOk("description"); ok {
			glbUpdate.Desc = description.(string)
		}
		if ttl, ok := d.GetOk("ttl"); ok {
			glbUpdate.Ttl = ttl.(int)
		}
		if enabled, ok := d.GetOk("enabled"); ok {
			glbUpdate.Enabled = enabled.(bool)
		}
		_, err = cisClient.Glbs().UpdateGlb(cisId, zoneId, glbId, glbUpdate)
		if err != nil {
			log.Printf("[WARN] Error getting zone during GlbUpdate %v\n", err)
			return err
		}
	}

	return resourceCISGlbRead(d, meta)
}

func resourceCISGlbDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	glbId, zoneId, cisId, _ := convertTfToCisThreeVar(d.Id())
	var glb *v1.Glb
	emptyGlb := new(v1.Glb)

	glb, err = cisClient.Glbs().GetGlb(cisId, zoneId, glbId)
	if err != nil {
		if checkCisGlbDeleted(d, meta, err, glb) {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during GlbRead %v\n", err)
		return err
	}

	glbObj := *glb
	if !reflect.DeepEqual(emptyGlb, glbObj) {
		err = cisClient.Glbs().DeleteGlb(cisId, zoneId, glbId)
		if err != nil {
			log.Printf("[WARN] DeleteGlb Failed %s\n", err)
			return err
		}
	}

	d.SetId("")
	return nil
}

func checkCisGlbDeleted(d *schema.ResourceData, meta interface{}, errCheck error, glb *v1.Glb) bool {
	// Check if error is due to removal of Cis resource and hence all subresources
	if strings.Contains(errCheck.Error(), "Object not found") ||
		strings.Contains(errCheck.Error(), "status code: 404") ||
		strings.Contains(errCheck.Error(), "Invalid zone identifier") { //code 400
		log.Printf("[WARN] Removing resource from state because it's not found via the CIS API")
		return true
	}
	_, _, cisId, _ := convertTfToCisThreeVar(d.Id())
	exists, errNew := rcInstanceExists(cisId, "ibm_cis", meta)
	if errNew != nil {
		log.Printf("resourceCISglbRead - Failure validating service exists %s\n", errNew)
		return false
	}
	if !exists {
		log.Printf("[WARN] Removing glb from state because parent cis instance is in removed state")
		return true
	}
	return false
}
