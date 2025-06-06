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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/eventarc/Channel.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package eventarc

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceEventarcChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventarcChannelCreate,
		Read:   resourceEventarcChannelRead,
		Update: resourceEventarcChannelUpdate,
		Delete: resourceEventarcChannelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceEventarcChannelImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the resource`,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The resource name of the channel. Must be unique within the location on the project.`,
			},
			"crypto_key_name": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern 'projects/*/locations/*/keyRings/*/cryptoKeys/*'.`,
			},
			"third_party_provider": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: 'projects/{project}/locations/{location}/providers/{provider_id}'.`,
			},
			"activation_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The activation token for the channel. The token must be used by the provider to register the channel for publishing.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The creation time.`,
			},
			"pubsub_topic": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: 'projects/{project}/topics/{topic_id}'.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The state of a Channel.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The last-modified time.`,
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

func resourceEventarcChannelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandEventarcChannelName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	providerProp, err := expandEventarcChannelThirdPartyProvider(d.Get("third_party_provider"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("third_party_provider"); !tpgresource.IsEmptyValue(reflect.ValueOf(providerProp)) && (ok || !reflect.DeepEqual(v, providerProp)) {
		obj["provider"] = providerProp
	}
	cryptoKeyNameProp, err := expandEventarcChannelCryptoKeyName(d.Get("crypto_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("crypto_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(cryptoKeyNameProp)) && (ok || !reflect.DeepEqual(v, cryptoKeyNameProp)) {
		obj["cryptoKeyName"] = cryptoKeyNameProp
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{EventarcBasePath}}projects/{{project}}/locations/{{location}}/channels?channelId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Channel: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Channel: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "POST",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutCreate),
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.EventarcChannel403Retry},
	})
	if err != nil {
		return fmt.Errorf("Error creating Channel: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/channels/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = EventarcOperationWaitTime(
		config, res, tpgresource.GetResourceNameFromSelfLink(project), "Creating Channel", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Channel: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Channel %q: %#v", d.Id(), res)

	return resourceEventarcChannelRead(d, meta)
}

func resourceEventarcChannelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{EventarcBasePath}}projects/{{project}}/locations/{{location}}/channels/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Channel: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "GET",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.EventarcChannel403Retry},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("EventarcChannel %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}

	if err := d.Set("name", flattenEventarcChannelName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("uid", flattenEventarcChannelUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("create_time", flattenEventarcChannelCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("update_time", flattenEventarcChannelUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("third_party_provider", flattenEventarcChannelThirdPartyProvider(res["provider"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("pubsub_topic", flattenEventarcChannelPubsubTopic(res["pubsubTopic"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("state", flattenEventarcChannelState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("activation_token", flattenEventarcChannelActivationToken(res["activationToken"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("crypto_key_name", flattenEventarcChannelCryptoKeyName(res["cryptoKeyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}

	return nil
}

func resourceEventarcChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Channel: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	obj := make(map[string]interface{})
	cryptoKeyNameProp, err := expandEventarcChannelCryptoKeyName(d.Get("crypto_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("crypto_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, cryptoKeyNameProp)) {
		obj["cryptoKeyName"] = cryptoKeyNameProp
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{EventarcBasePath}}projects/{{project}}/locations/{{location}}/channels/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Channel %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("crypto_key_name") {
		updateMask = append(updateMask, "cryptoKeyName")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:               config,
			Method:               "PATCH",
			Project:              billingProject,
			RawURL:               url,
			UserAgent:            userAgent,
			Body:                 obj,
			Timeout:              d.Timeout(schema.TimeoutUpdate),
			Headers:              headers,
			ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.EventarcChannel403Retry},
		})

		if err != nil {
			return fmt.Errorf("Error updating Channel %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Channel %q: %#v", d.Id(), res)
		}

		err = EventarcOperationWaitTime(
			config, res, tpgresource.GetResourceNameFromSelfLink(project), "Updating Channel", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceEventarcChannelRead(d, meta)
}

func resourceEventarcChannelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Channel: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{EventarcBasePath}}projects/{{project}}/locations/{{location}}/channels/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Channel %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "DELETE",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutDelete),
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.EventarcChannel403Retry},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Channel")
	}

	err = EventarcOperationWaitTime(
		config, res, tpgresource.GetResourceNameFromSelfLink(project), "Deleting Channel", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Channel %q: %#v", d.Id(), res)
	return nil
}

func resourceEventarcChannelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/channels/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/channels/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenEventarcChannelName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("name")
}

func flattenEventarcChannelUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEventarcChannelCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEventarcChannelUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEventarcChannelThirdPartyProvider(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEventarcChannelPubsubTopic(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEventarcChannelState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEventarcChannelActivationToken(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEventarcChannelCryptoKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandEventarcChannelName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return expandToRegionalLongForm("projects/%s/locations/%s/channels/%s", v, d, config)
}

func expandEventarcChannelThirdPartyProvider(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEventarcChannelCryptoKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
