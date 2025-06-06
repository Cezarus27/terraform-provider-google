// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apihub/ApiHubInstance.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apihub

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceApihubApiHubInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceApihubApiHubInstanceCreate,
		Read:   resourceApihubApiHubInstanceRead,
		Update: resourceApihubApiHubInstanceUpdate,
		Delete: resourceApihubApiHubInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApihubApiHubInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `Available configurations to provision an ApiHub Instance.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cmek_key_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Optional. The Customer Managed Encryption Key (CMEK) used for data encryption.
The CMEK name should follow the format of
'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)',
where the location must match the instance location.
If the CMEK is not provided, a GMEK will be created for the instance.`,
						},
						"disable_search": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
							Description: `Optional. If true, the search will be disabled for the instance. The default value
is false.`,
						},
						"encryption_type": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							ForceNew: true,
							Description: `Optional. Encryption type for the region. If the encryption type is CMEK, the
cmek_key_name must be provided. If no encryption type is provided,
GMEK will be used.
Possible values:
ENCRYPTION_TYPE_UNSPECIFIED
GMEK
CMEK`,
						},
						"vertex_location": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Optional. The name of the Vertex AI location where the data store is stored.`,
						},
					},
				},
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.`,
			},
			"api_hub_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Optional. Identifier to assign to the Api Hub instance. Must be unique within
scope of the parent resource. If the field is not provided,
system generated id will be used.

This value should be 4-40 characters, and valid characters
are '/a-z[0-9]-_/'.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Optional. Description of the ApiHub instance.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Optional. Instance labels to represent user-provided metadata.
Refer to cloud documentation on labels for more details.
https://cloud.google.com/compute/docs/labeling-resources

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Creation timestamp.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				ForceNew:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Identifier. Format:
'projects/{project}/locations/{location}/apiHubInstances/{apiHubInstance}'.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The current state of the ApiHub instance.
Possible values:
STATE_UNSPECIFIED
INACTIVE
CREATING
ACTIVE
UPDATING
DELETING
FAILED`,
			},
			"state_message": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Extra information about ApiHub instance state. Currently the message
would be populated when state is 'FAILED'.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Last update timestamp.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApihubApiHubInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandApihubApiHubInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	configProp, err := expandApihubApiHubInstanceConfig(d.Get("config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	labelsProp, err := expandApihubApiHubInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApihubBasePath}}projects/{{project}}/locations/{{location}}/apiHubInstances?apiHubInstanceId={{api_hub_instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ApiHubInstance: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApiHubInstance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating ApiHubInstance: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ApihubOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating ApiHubInstance", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create ApiHubInstance: %s", err)
	}

	if err := d.Set("name", flattenApihubApiHubInstanceName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ApiHubInstance %q: %#v", d.Id(), res)

	return resourceApihubApiHubInstanceRead(d, meta)
}

func resourceApihubApiHubInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApihubBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApiHubInstance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApihubApiHubInstance %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}

	if err := d.Set("description", flattenApihubApiHubInstanceDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}
	if err := d.Set("name", flattenApihubApiHubInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}
	if err := d.Set("create_time", flattenApihubApiHubInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}
	if err := d.Set("update_time", flattenApihubApiHubInstanceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}
	if err := d.Set("state", flattenApihubApiHubInstanceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}
	if err := d.Set("state_message", flattenApihubApiHubInstanceStateMessage(res["stateMessage"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}
	if err := d.Set("config", flattenApihubApiHubInstanceConfig(res["config"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}
	if err := d.Set("labels", flattenApihubApiHubInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}
	if err := d.Set("terraform_labels", flattenApihubApiHubInstanceTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}
	if err := d.Set("effective_labels", flattenApihubApiHubInstanceEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApiHubInstance: %s", err)
	}

	return nil
}

func resourceApihubApiHubInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	// Only the root field "labels" and "terraform_labels" are mutable
	return resourceApihubApiHubInstanceRead(d, meta)
}

func resourceApihubApiHubInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Apihub ApiHubInstance resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceApihubApiHubInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/apiHubInstances/(?P<api_hub_instance_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<api_hub_instance_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<api_hub_instance_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Set name based on the components
	if err := d.Set("name", "projects/{{project}}/locations/{{location}}/apiHubInstances/{{api_hub_instance_id}}"); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, d.Get("name").(string))
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApihubApiHubInstanceDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApihubApiHubInstanceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApihubApiHubInstanceCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApihubApiHubInstanceUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApihubApiHubInstanceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApihubApiHubInstanceStateMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApihubApiHubInstanceConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["encryption_type"] =
		flattenApihubApiHubInstanceConfigEncryptionType(original["encryptionType"], d, config)
	transformed["cmek_key_name"] =
		flattenApihubApiHubInstanceConfigCmekKeyName(original["cmekKeyName"], d, config)
	transformed["disable_search"] =
		flattenApihubApiHubInstanceConfigDisableSearch(original["disableSearch"], d, config)
	transformed["vertex_location"] =
		flattenApihubApiHubInstanceConfigVertexLocation(original["vertexLocation"], d, config)
	return []interface{}{transformed}
}
func flattenApihubApiHubInstanceConfigEncryptionType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApihubApiHubInstanceConfigCmekKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApihubApiHubInstanceConfigDisableSearch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApihubApiHubInstanceConfigVertexLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApihubApiHubInstanceLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenApihubApiHubInstanceTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenApihubApiHubInstanceEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApihubApiHubInstanceDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApihubApiHubInstanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEncryptionType, err := expandApihubApiHubInstanceConfigEncryptionType(original["encryption_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEncryptionType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["encryptionType"] = transformedEncryptionType
	}

	transformedCmekKeyName, err := expandApihubApiHubInstanceConfigCmekKeyName(original["cmek_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCmekKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cmekKeyName"] = transformedCmekKeyName
	}

	transformedDisableSearch, err := expandApihubApiHubInstanceConfigDisableSearch(original["disable_search"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisableSearch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disableSearch"] = transformedDisableSearch
	}

	transformedVertexLocation, err := expandApihubApiHubInstanceConfigVertexLocation(original["vertex_location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVertexLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vertexLocation"] = transformedVertexLocation
	}

	return transformed, nil
}

func expandApihubApiHubInstanceConfigEncryptionType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApihubApiHubInstanceConfigCmekKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApihubApiHubInstanceConfigDisableSearch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApihubApiHubInstanceConfigVertexLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApihubApiHubInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
