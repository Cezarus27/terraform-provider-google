---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/accesscontextmanager/ServicePerimeter.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Access Context Manager (VPC Service Controls)"
description: |-
  ServicePerimeter describes a set of GCP resources which can freely import
  and export data amongst themselves, but not export outside of the
  ServicePerimeter.
---

# google_access_context_manager_service_perimeter

ServicePerimeter describes a set of GCP resources which can freely import
and export data amongst themselves, but not export outside of the
ServicePerimeter. If a request with a source within this ServicePerimeter
has a target outside of the ServicePerimeter, the request will be blocked.
Otherwise the request is allowed. There are two types of Service Perimeter
- Regular and Bridge. Regular Service Perimeters cannot overlap, a single
GCP project can only belong to a single regular Service Perimeter. Service
Perimeter Bridges can contain only GCP projects as members, a single GCP
project may belong to multiple Service Perimeter Bridges.


To get more information about ServicePerimeter, see:

* [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
* How-to Guides
    * [Guide to Ingress and Egress Rules](https://cloud.google.com/vpc-service-controls/docs/ingress-egress-rules)
    * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)

~> **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
you must specify a `billing_project` and set `user_project_override` to true
in the provider configuration. Otherwise the ACM API will return a 403 error.
Your account must have the `serviceusage.services.use` permission on the
`billing_project` you defined.

## Example Usage - Access Context Manager Service Perimeter Basic


```hcl
resource "google_access_context_manager_service_perimeter" "service-perimeter" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/restrict_storage"
  title  = "restrict_storage"
  status {
    restricted_services = ["storage.googleapis.com"]
  }
}

resource "google_access_context_manager_access_level" "access-level" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/accessLevels/chromeos_no_lock"
  title  = "chromeos_no_lock"
  basic {
    conditions {
      device_policy {
        require_screen_lock = false
        os_constraints {
          os_type = "DESKTOP_CHROME_OS"
        }
      }
      regions = [
        "CH",
        "IT",
        "US",
      ]
    }
  }
}

resource "google_access_context_manager_access_policy" "access-policy" {
  parent = "organizations/123456789"
  title  = "my policy"
}
```
## Example Usage - Access Context Manager Service Perimeter Secure Data Exchange


```hcl
resource "google_access_context_manager_service_perimeters" "secure-data-exchange" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"

  service_perimeters {
    name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/"
    title  = ""
    status {
      restricted_services = ["storage.googleapis.com"]
    }
  }

  service_perimeters {
    name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/"
    title  = ""
    status {
      restricted_services = ["bigtable.googleapis.com"]
      		vpcAccessibleServices = {
			enableRestriction = true
			allowedServices = ["bigquery.googleapis.com"]
		}
    }
  }
}

resource "google_access_context_manager_access_level" "access-level" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/accessLevels/secure_data_exchange"
  title  = "secure_data_exchange"
  basic {
    conditions {
      device_policy {
        require_screen_lock = false
        os_constraints {
          os_type = "DESKTOP_CHROME_OS"
        }
      }
      regions = [
        "CH",
        "IT",
        "US",
      ]
    }
  }
}

resource "google_access_context_manager_access_policy" "access-policy" {
  parent = "organizations/123456789"
  title  = "my policy"
}

resource "google_access_context_manager_service_perimeter" "test-access" {
  parent         = "accessPolicies/${google_access_context_manager_access_policy.test-access.name}"
  name           = "accessPolicies/${google_access_context_manager_access_policy.test-access.name}/servicePerimeters/%s"
  title          = "%s"
  perimeter_type = "PERIMETER_TYPE_REGULAR"
  status {
    restricted_services = ["bigquery.googleapis.com", "storage.googleapis.com"]
		access_levels       = [google_access_context_manager_access_level.access-level.name]

		vpc_accessible_services {
			enable_restriction = true
			allowed_services   = ["bigquery.googleapis.com", "storage.googleapis.com"]
		}

		ingress_policies {
			ingress_from {
				sources {
					access_level = google_access_context_manager_access_level.test-access.name
				}
				identity_type = "ANY_IDENTITY"
			}

			ingress_to {
				resources = [ "*" ]
				operations {
					service_name = "bigquery.googleapis.com"

					method_selectors {
						method = "BigQueryStorage.ReadRows"
					}

					method_selectors {
						method = "TableService.ListTables"
					}

					method_selectors {
						permission = "bigquery.jobs.get"
					}
				}

				operations {
					service_name = "storage.googleapis.com"

					method_selectors {
						method = "google.storage.objects.create"
					}
				}
			}
		}

		egress_policies {
			egress_from {
				identity_type = "ANY_USER_ACCOUNT"
			}
		}
  }
}
```
## Example Usage - Access Context Manager Service Perimeter Dry-Run


```hcl
resource "google_access_context_manager_service_perimeter" "service-perimeter" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/restrict_bigquery_dryrun_storage"
  title  = "restrict_bigquery_dryrun_storage"

  # Service 'bigquery.googleapis.com' will be restricted.
  status {
    restricted_services = ["bigquery.googleapis.com"]
  }

  # Service 'storage.googleapis.com' will be in dry-run mode.
  spec {
    restricted_services = ["storage.googleapis.com"]
  }

  use_explicit_dry_run_spec = true

}

resource "google_access_context_manager_access_policy" "access-policy" {
  parent = "organizations/123456789"
  title  = "my policy"
}
```
## Example Usage - Access Context Manager Service Perimeter Granular Controls


```hcl
resource "google_access_context_manager_access_policy" "access-policy" {
  parent = "organizations/123456789"
  title  = "Policy with Granular Controls Support"
}

resource "google_access_context_manager_service_perimeter" "granular-controls-perimeter" {
  parent         = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name           = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/%s"
  title          = "%s"
  perimeter_type = "PERIMETER_TYPE_REGULAR"
  status {
      restricted_services = ["bigquery.googleapis.com"]

      vpc_accessible_services {
          enable_restriction = true
          allowed_services   = ["bigquery.googleapis.com"]
      }

      ingress_policies {
          ingress_from {
              sources {
                 resource = "projects/1234" 
              }
              identities = ["group:database-admins@google.com"]
              identities = ["principal://iam.googleapis.com/locations/global/workforcePools/1234/subject/janedoe"]
              identities = ["principalSet://iam.googleapis.com/locations/global/workforcePools/1234/*"]
          }
          ingress_to {
              resources = [ "*" ]
              roles = ["roles/bigquery.admin", "organizations/1234/roles/bigquery_custom_role"]
          }
      }

      egress_policies {
          egress_from {
              identities = ["group:database-admins@google.com"]
              identities = ["principal://iam.googleapis.com/locations/global/workforcePools/1234/subject/janedoe"]
              identities = ["principalSet://iam.googleapis.com/locations/global/workforcePools/1234/*"]
          }
          egress_to {
              resources = [ "*" ]
              roles = ["roles/bigquery.admin", "organizations/1234/roles/bigquery_custom_role"]
          }
      }
   }
}
```

## Argument Reference

The following arguments are supported:


* `title` -
  (Required)
  Human readable title. Must be unique within the Policy.

* `parent` -
  (Required)
  The AccessPolicy this ServicePerimeter lives in.
  Format: accessPolicies/{policy_id}

* `name` -
  (Required)
  Resource name for the ServicePerimeter. The short_name component must
  begin with a letter and only include alphanumeric and '_'.
  Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}


* `description` -
  (Optional)
  Description of the ServicePerimeter and its use. Does not affect
  behavior.

* `perimeter_type` -
  (Optional)
  Specifies the type of the Perimeter. There are two types: regular and
  bridge. Regular Service Perimeter contains resources, access levels,
  and restricted services. Every resource can be in at most
  ONE regular Service Perimeter.
  In addition to being in a regular service perimeter, a resource can also
  be in zero or more perimeter bridges. A perimeter bridge only contains
  resources. Cross project operations are permitted if all effected
  resources share some perimeter (whether bridge or regular). Perimeter
  Bridge does not contain access levels or services: those are governed
  entirely by the regular perimeter that resource is in.
  Perimeter Bridges are typically useful when building more complex
  topologies with many independent perimeters that need to share some data
  with a common perimeter, but should not be able to share data among
  themselves.
  Default value is `PERIMETER_TYPE_REGULAR`.
  Possible values are: `PERIMETER_TYPE_REGULAR`, `PERIMETER_TYPE_BRIDGE`.

* `status` -
  (Optional)
  ServicePerimeter configuration. Specifies sets of resources,
  restricted services and access levels that determine
  perimeter content and boundaries.
  Structure is [documented below](#nested_status).

* `spec` -
  (Optional)
  Proposed (or dry run) ServicePerimeter configuration.
  This configuration allows to specify and test ServicePerimeter configuration
  without enforcing actual access restrictions. Only allowed to be set when
  the `useExplicitDryRunSpec` flag is set.
  Structure is [documented below](#nested_spec).

* `use_explicit_dry_run_spec` -
  (Optional)
  Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
  for all Service Perimeters, and that spec is identical to the status for those
  Service Perimeters. When this flag is set, it inhibits the generation of the
  implicit spec, thereby allowing the user to explicitly provide a
  configuration ("spec") to use in a dry-run version of the Service Perimeter.
  This allows the user to test changes to the enforced config ("status") without
  actually enforcing them. This testing is done through analyzing the differences
  between currently enforced and suggested restrictions. useExplicitDryRunSpec must
  bet set to True if any of the fields in the spec are set to non-default values.



<a name="nested_status"></a>The `status` block supports:

* `resources` -
  (Optional)
  A list of GCP resources that are inside of the service perimeter.
  Currently only projects are allowed.
  Format: projects/{project_number}

* `access_levels` -
  (Optional)
  A list of AccessLevel resource names that allow resources within
  the ServicePerimeter to be accessed from the internet.
  AccessLevels listed must be in the same policy as this
  ServicePerimeter. Referencing a nonexistent AccessLevel is a
  syntax error. If no AccessLevel names are listed, resources within
  the perimeter can only be accessed via GCP calls with request
  origins within the perimeter. For Service Perimeter Bridge, must
  be empty.
  Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}

* `restricted_services` -
  (Optional)
  GCP services that are subject to the Service Perimeter
  restrictions. Must contain a list of services. For example, if
  `storage.googleapis.com` is specified, access to the storage
  buckets inside the perimeter must meet the perimeter's access
  restrictions.

* `vpc_accessible_services` -
  (Optional)
  Specifies how APIs are allowed to communicate within the Service
  Perimeter.
  Structure is [documented below](#nested_status_vpc_accessible_services).

* `ingress_policies` -
  (Optional)
  List of `IngressPolicies` to apply to the perimeter. A perimeter may
  have multiple `IngressPolicies`, each of which is evaluated
  separately. Access is granted if any `Ingress Policy` grants it.
  Must be empty for a perimeter bridge.
  Structure is [documented below](#nested_status_ingress_policies).

* `egress_policies` -
  (Optional)
  List of EgressPolicies to apply to the perimeter. A perimeter may
  have multiple EgressPolicies, each of which is evaluated separately.
  Access is granted if any EgressPolicy grants it. Must be empty for
  a perimeter bridge.
  Structure is [documented below](#nested_status_egress_policies).


<a name="nested_status_vpc_accessible_services"></a>The `vpc_accessible_services` block supports:

* `enable_restriction` -
  (Optional)
  Whether to restrict API calls within the Service Perimeter to the
  list of APIs specified in 'allowedServices'.

* `allowed_services` -
  (Optional)
  The list of APIs usable within the Service Perimeter.
  Must be empty unless `enableRestriction` is True.

<a name="nested_status_ingress_policies"></a>The `ingress_policies` block supports:

* `ingress_from` -
  (Optional)
  Defines the conditions on the source of a request causing this `IngressPolicy`
  to apply.
  Structure is [documented below](#nested_status_ingress_policies_ingress_policies_ingress_from).

* `ingress_to` -
  (Optional)
  Defines the conditions on the `ApiOperation` and request destination that cause
  this `IngressPolicy` to apply.
  Structure is [documented below](#nested_status_ingress_policies_ingress_policies_ingress_to).

* `title` -
  (Optional)
  Human readable title. Must be unique within the perimeter. Does not affect behavior.


<a name="nested_status_ingress_policies_ingress_policies_ingress_from"></a>The `ingress_from` block supports:

* `identity_type` -
  (Optional)
  Specifies the type of identities that are allowed access from outside the
  perimeter. If left unspecified, then members of `identities` field will be
  allowed access.
  Possible values are: `IDENTITY_TYPE_UNSPECIFIED`, `ANY_IDENTITY`, `ANY_USER_ACCOUNT`, `ANY_SERVICE_ACCOUNT`.

* `identities` -
  (Optional)
  Identities can be an individual user, service account, Google group,
  or third-party identity. For third-party identity, only single identities
  are supported and other identity types are not supported.The v1 identities
  that have the prefix user, group and serviceAccount in
  https://cloud.google.com/iam/docs/principal-identifiers#v1 are supported.

* `sources` -
  (Optional)
  Sources that this `IngressPolicy` authorizes access from.
  Structure is [documented below](#nested_status_ingress_policies_ingress_policies_ingress_from_sources).


<a name="nested_status_ingress_policies_ingress_policies_ingress_from_sources"></a>The `sources` block supports:

* `access_level` -
  (Optional)
  An `AccessLevel` resource name that allow resources within the
  `ServicePerimeters` to be accessed from the internet. `AccessLevels` listed
  must be in the same policy as this `ServicePerimeter`. Referencing a nonexistent
  `AccessLevel` will cause an error. If no `AccessLevel` names are listed,
  resources within the perimeter can only be accessed via Google Cloud calls
  with request origins within the perimeter.
  Example `accessPolicies/MY_POLICY/accessLevels/MY_LEVEL.`
  If * is specified, then all IngressSources will be allowed.

* `resource` -
  (Optional)
  A Google Cloud resource that is allowed to ingress the perimeter.
  Requests from these resources will be allowed to access perimeter data.
  Currently only projects and VPCs are allowed.
  Project format: `projects/{projectNumber}`
  VPC network format:
  `//compute.googleapis.com/projects/{PROJECT_ID}/global/networks/{NAME}`.
  The project may be in any Google Cloud organization, not just the
  organization that the perimeter is defined in. `*` is not allowed, the case
  of allowing all Google Cloud resources only is not supported.

<a name="nested_status_ingress_policies_ingress_policies_ingress_to"></a>The `ingress_to` block supports:

* `resources` -
  (Optional)
  A list of resources, currently only projects in the form
  `projects/<projectnumber>`, protected by this `ServicePerimeter`
  that are allowed to be accessed by sources defined in the
  corresponding `IngressFrom`. A request matches if it contains
  a resource in this list. If `*` is specified for resources,
  then this `IngressTo` rule will authorize access to all
  resources inside the perimeter, provided that the request
  also matches the `operations` field.

* `roles` -
  (Optional)
  A list of IAM roles that represent the set of operations that the sources
  specified in the corresponding `IngressFrom`
  are allowed to perform.

* `operations` -
  (Optional)
  A list of `ApiOperations` the sources specified in corresponding `IngressFrom`
  are allowed to perform in this `ServicePerimeter`.
  Structure is [documented below](#nested_status_ingress_policies_ingress_policies_ingress_to_operations).


<a name="nested_status_ingress_policies_ingress_policies_ingress_to_operations"></a>The `operations` block supports:

* `service_name` -
  (Optional)
  The name of the API whose methods or permissions the `IngressPolicy` or
  `EgressPolicy` want to allow. A single `ApiOperation` with `serviceName`
  field set to `*` will allow all methods AND permissions for all services.

* `method_selectors` -
  (Optional)
  API methods or permissions to allow. Method or permission must belong to
  the service specified by serviceName field. A single `MethodSelector` entry
  with `*` specified for the method field will allow all methods AND
  permissions for the service specified in `serviceName`.
  Structure is [documented below](#nested_status_ingress_policies_ingress_policies_ingress_to_operations_operations_method_selectors).


<a name="nested_status_ingress_policies_ingress_policies_ingress_to_operations_operations_method_selectors"></a>The `method_selectors` block supports:

* `method` -
  (Optional)
  Value for method should be a valid method name for the corresponding
  serviceName in `ApiOperation`. If `*` used as value for `method`, then
  ALL methods and permissions are allowed.

* `permission` -
  (Optional)
  Value for permission should be a valid Cloud IAM permission for the
  corresponding `serviceName` in `ApiOperation`.

<a name="nested_status_egress_policies"></a>The `egress_policies` block supports:

* `egress_from` -
  (Optional)
  Defines conditions on the source of a request causing this `EgressPolicy` to apply.
  Structure is [documented below](#nested_status_egress_policies_egress_policies_egress_from).

* `egress_to` -
  (Optional)
  Defines the conditions on the `ApiOperation` and destination resources that
  cause this `EgressPolicy` to apply.
  Structure is [documented below](#nested_status_egress_policies_egress_policies_egress_to).

* `title` -
  (Optional)
  Human readable title. Must be unique within the perimeter. Does not affect behavior.


<a name="nested_status_egress_policies_egress_policies_egress_from"></a>The `egress_from` block supports:

* `identity_type` -
  (Optional)
  Specifies the type of identities that are allowed access to outside the
  perimeter. If left unspecified, then members of `identities` field will
  be allowed access.
  Possible values are: `IDENTITY_TYPE_UNSPECIFIED`, `ANY_IDENTITY`, `ANY_USER_ACCOUNT`, `ANY_SERVICE_ACCOUNT`.

* `sources` -
  (Optional)
  Sources that this EgressPolicy authorizes access from.
  Structure is [documented below](#nested_status_egress_policies_egress_policies_egress_from_sources).

* `source_restriction` -
  (Optional)
  Whether to enforce traffic restrictions based on `sources` field. If the `sources` field is non-empty, then this field must be set to `SOURCE_RESTRICTION_ENABLED`.
  Possible values are: `SOURCE_RESTRICTION_UNSPECIFIED`, `SOURCE_RESTRICTION_ENABLED`, `SOURCE_RESTRICTION_DISABLED`.

* `identities` -
  (Optional)
  Identities can be an individual user, service account, Google group,
  or third-party identity. For third-party identity, only single identities
  are supported and other identity types are not supported.The v1 identities
  that have the prefix user, group and serviceAccount in
  https://cloud.google.com/iam/docs/principal-identifiers#v1 are supported.


<a name="nested_status_egress_policies_egress_policies_egress_from_sources"></a>The `sources` block supports:

* `access_level` -
  (Optional)
  An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.

* `resource` -
  (Optional)
  A Google Cloud resource that is allowed to egress the perimeter.
  Requests from these resources are allowed to access data outside the perimeter.
  Currently only projects are allowed. Project format: `projects/{project_number}`.
  The resource may be in any Google Cloud organization, not just the
  organization that the perimeter is defined in. `*` is not allowed, the
  case of allowing all Google Cloud resources only is not supported.

<a name="nested_status_egress_policies_egress_policies_egress_to"></a>The `egress_to` block supports:

* `resources` -
  (Optional)
  A list of resources, currently only projects in the form
  `projects/<projectnumber>`, that match this to stanza. A request matches
  if it contains a resource in this list. If * is specified for resources,
  then this `EgressTo` rule will authorize access to all resources outside
  the perimeter.

* `external_resources` -
  (Optional)
  A list of external resources that are allowed to be accessed. A request
  matches if it contains an external resource in this list (Example:
  s3://bucket/path). Currently '*' is not allowed.

* `roles` -
  (Optional)
  A list of IAM roles that represent the set of operations that the sources
  specified in the corresponding `EgressFrom`
  are allowed to perform.

* `operations` -
  (Optional)
  A list of `ApiOperations` that this egress rule applies to. A request matches
  if it contains an operation/service in this list.
  Structure is [documented below](#nested_status_egress_policies_egress_policies_egress_to_operations).


<a name="nested_status_egress_policies_egress_policies_egress_to_operations"></a>The `operations` block supports:

* `service_name` -
  (Optional)
  The name of the API whose methods or permissions the `IngressPolicy` or
  `EgressPolicy` want to allow. A single `ApiOperation` with serviceName
  field set to `*` will allow all methods AND permissions for all services.

* `method_selectors` -
  (Optional)
  API methods or permissions to allow. Method or permission must belong
  to the service specified by `serviceName` field. A single MethodSelector
  entry with `*` specified for the `method` field will allow all methods
  AND permissions for the service specified in `serviceName`.
  Structure is [documented below](#nested_status_egress_policies_egress_policies_egress_to_operations_operations_method_selectors).


<a name="nested_status_egress_policies_egress_policies_egress_to_operations_operations_method_selectors"></a>The `method_selectors` block supports:

* `method` -
  (Optional)
  Value for `method` should be a valid method name for the corresponding
  `serviceName` in `ApiOperation`. If `*` used as value for method,
  then ALL methods and permissions are allowed.

* `permission` -
  (Optional)
  Value for permission should be a valid Cloud IAM permission for the
  corresponding `serviceName` in `ApiOperation`.

<a name="nested_spec"></a>The `spec` block supports:

* `resources` -
  (Optional)
  A list of GCP resources that are inside of the service perimeter.
  Currently only projects are allowed.
  Format: projects/{project_number}

* `access_levels` -
  (Optional)
  A list of AccessLevel resource names that allow resources within
  the ServicePerimeter to be accessed from the internet.
  AccessLevels listed must be in the same policy as this
  ServicePerimeter. Referencing a nonexistent AccessLevel is a
  syntax error. If no AccessLevel names are listed, resources within
  the perimeter can only be accessed via GCP calls with request
  origins within the perimeter. For Service Perimeter Bridge, must
  be empty.
  Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}

* `restricted_services` -
  (Optional)
  GCP services that are subject to the Service Perimeter
  restrictions. Must contain a list of services. For example, if
  `storage.googleapis.com` is specified, access to the storage
  buckets inside the perimeter must meet the perimeter's access
  restrictions.

* `vpc_accessible_services` -
  (Optional)
  Specifies how APIs are allowed to communicate within the Service
  Perimeter.
  Structure is [documented below](#nested_spec_vpc_accessible_services).

* `ingress_policies` -
  (Optional)
  List of `IngressPolicies` to apply to the perimeter. A perimeter may
  have multiple `IngressPolicies`, each of which is evaluated
  separately. Access is granted if any `Ingress Policy` grants it.
  Must be empty for a perimeter bridge.
  Structure is [documented below](#nested_spec_ingress_policies).

* `egress_policies` -
  (Optional)
  List of EgressPolicies to apply to the perimeter. A perimeter may
  have multiple EgressPolicies, each of which is evaluated separately.
  Access is granted if any EgressPolicy grants it. Must be empty for
  a perimeter bridge.
  Structure is [documented below](#nested_spec_egress_policies).


<a name="nested_spec_vpc_accessible_services"></a>The `vpc_accessible_services` block supports:

* `enable_restriction` -
  (Optional)
  Whether to restrict API calls within the Service Perimeter to the
  list of APIs specified in 'allowedServices'.

* `allowed_services` -
  (Optional)
  The list of APIs usable within the Service Perimeter.
  Must be empty unless `enableRestriction` is True.

<a name="nested_spec_ingress_policies"></a>The `ingress_policies` block supports:

* `ingress_from` -
  (Optional)
  Defines the conditions on the source of a request causing this `IngressPolicy`
  to apply.
  Structure is [documented below](#nested_spec_ingress_policies_ingress_policies_ingress_from).

* `ingress_to` -
  (Optional)
  Defines the conditions on the `ApiOperation` and request destination that cause
  this `IngressPolicy` to apply.
  Structure is [documented below](#nested_spec_ingress_policies_ingress_policies_ingress_to).

* `title` -
  (Optional)
  Human readable title. Must be unique within the perimeter. Does not affect behavior.


<a name="nested_spec_ingress_policies_ingress_policies_ingress_from"></a>The `ingress_from` block supports:

* `identity_type` -
  (Optional)
  Specifies the type of identities that are allowed access from outside the
  perimeter. If left unspecified, then members of `identities` field will be
  allowed access.
  Possible values are: `IDENTITY_TYPE_UNSPECIFIED`, `ANY_IDENTITY`, `ANY_USER_ACCOUNT`, `ANY_SERVICE_ACCOUNT`.

* `identities` -
  (Optional)
  A list of identities that are allowed access through this ingress policy.
  Should be in the format of email address. The email address should represent
  individual user or service account only.

* `sources` -
  (Optional)
  Sources that this `IngressPolicy` authorizes access from.
  Structure is [documented below](#nested_spec_ingress_policies_ingress_policies_ingress_from_sources).


<a name="nested_spec_ingress_policies_ingress_policies_ingress_from_sources"></a>The `sources` block supports:

* `access_level` -
  (Optional)
  An `AccessLevel` resource name that allow resources within the
  `ServicePerimeters` to be accessed from the internet. `AccessLevels` listed
  must be in the same policy as this `ServicePerimeter`. Referencing a nonexistent
  `AccessLevel` will cause an error. If no `AccessLevel` names are listed,
  resources within the perimeter can only be accessed via Google Cloud calls
  with request origins within the perimeter.
  Example `accessPolicies/MY_POLICY/accessLevels/MY_LEVEL.`
  If * is specified, then all IngressSources will be allowed.

* `resource` -
  (Optional)
  A Google Cloud resource that is allowed to ingress the perimeter.
  Requests from these resources will be allowed to access perimeter data.
  Currently only projects are allowed. Format `projects/{project_number}`
  The project may be in any Google Cloud organization, not just the
  organization that the perimeter is defined in. `*` is not allowed, the case
  of allowing all Google Cloud resources only is not supported.

<a name="nested_spec_ingress_policies_ingress_policies_ingress_to"></a>The `ingress_to` block supports:

* `resources` -
  (Optional)
  A list of resources, currently only projects in the form
  `projects/<projectnumber>`, protected by this `ServicePerimeter`
  that are allowed to be accessed by sources defined in the
  corresponding `IngressFrom`. A request matches if it contains
  a resource in this list. If `*` is specified for resources,
  then this `IngressTo` rule will authorize access to all
  resources inside the perimeter, provided that the request
  also matches the `operations` field.

* `roles` -
  (Optional)
  A list of IAM roles that represent the set of operations that the sources
  specified in the corresponding `IngressFrom`
  are allowed to perform.

* `operations` -
  (Optional)
  A list of `ApiOperations` the sources specified in corresponding `IngressFrom`
  are allowed to perform in this `ServicePerimeter`.
  Structure is [documented below](#nested_spec_ingress_policies_ingress_policies_ingress_to_operations).


<a name="nested_spec_ingress_policies_ingress_policies_ingress_to_operations"></a>The `operations` block supports:

* `service_name` -
  (Optional)
  The name of the API whose methods or permissions the `IngressPolicy` or
  `EgressPolicy` want to allow. A single `ApiOperation` with `serviceName`
  field set to `*` will allow all methods AND permissions for all services.

* `method_selectors` -
  (Optional)
  API methods or permissions to allow. Method or permission must belong to
  the service specified by serviceName field. A single `MethodSelector` entry
  with `*` specified for the method field will allow all methods AND
  permissions for the service specified in `serviceName`.
  Structure is [documented below](#nested_spec_ingress_policies_ingress_policies_ingress_to_operations_operations_method_selectors).


<a name="nested_spec_ingress_policies_ingress_policies_ingress_to_operations_operations_method_selectors"></a>The `method_selectors` block supports:

* `method` -
  (Optional)
  Value for method should be a valid method name for the corresponding
  serviceName in `ApiOperation`. If `*` used as value for `method`, then
  ALL methods and permissions are allowed.

* `permission` -
  (Optional)
  Value for permission should be a valid Cloud IAM permission for the
  corresponding `serviceName` in `ApiOperation`.

<a name="nested_spec_egress_policies"></a>The `egress_policies` block supports:

* `egress_from` -
  (Optional)
  Defines conditions on the source of a request causing this `EgressPolicy` to apply.
  Structure is [documented below](#nested_spec_egress_policies_egress_policies_egress_from).

* `egress_to` -
  (Optional)
  Defines the conditions on the `ApiOperation` and destination resources that
  cause this `EgressPolicy` to apply.
  Structure is [documented below](#nested_spec_egress_policies_egress_policies_egress_to).

* `title` -
  (Optional)
  Human readable title. Must be unique within the perimeter. Does not affect behavior.


<a name="nested_spec_egress_policies_egress_policies_egress_from"></a>The `egress_from` block supports:

* `identity_type` -
  (Optional)
  Specifies the type of identities that are allowed access to outside the
  perimeter. If left unspecified, then members of `identities` field will
  be allowed access.
  Possible values are: `IDENTITY_TYPE_UNSPECIFIED`, `ANY_IDENTITY`, `ANY_USER_ACCOUNT`, `ANY_SERVICE_ACCOUNT`.

* `sources` -
  (Optional)
  Sources that this EgressPolicy authorizes access from.
  Structure is [documented below](#nested_spec_egress_policies_egress_policies_egress_from_sources).

* `source_restriction` -
  (Optional)
  Whether to enforce traffic restrictions based on `sources` field. If the `sources` field is non-empty, then this field must be set to `SOURCE_RESTRICTION_ENABLED`.
  Possible values are: `SOURCE_RESTRICTION_UNSPECIFIED`, `SOURCE_RESTRICTION_ENABLED`, `SOURCE_RESTRICTION_DISABLED`.

* `identities` -
  (Optional)
  A list of identities that are allowed access through this `EgressPolicy`.
  Should be in the format of email address. The email address should
  represent individual user or service account only.


<a name="nested_spec_egress_policies_egress_policies_egress_from_sources"></a>The `sources` block supports:

* `access_level` -
  (Optional)
  An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.

* `resource` -
  (Optional)
  A Google Cloud resource that is allowed to egress the perimeter.
  Requests from these resources are allowed to access data outside the perimeter.
  Currently only projects are allowed. Project format: `projects/{project_number}`.
  The resource may be in any Google Cloud organization, not just the
  organization that the perimeter is defined in. `*` is not allowed, the
  case of allowing all Google Cloud resources only is not supported.

<a name="nested_spec_egress_policies_egress_policies_egress_to"></a>The `egress_to` block supports:

* `resources` -
  (Optional)
  A list of resources, currently only projects in the form
  `projects/<projectnumber>`, that match this to stanza. A request matches
  if it contains a resource in this list. If * is specified for resources,
  then this `EgressTo` rule will authorize access to all resources outside
  the perimeter.

* `external_resources` -
  (Optional)
  A list of external resources that are allowed to be accessed. A request
  matches if it contains an external resource in this list (Example:
  s3://bucket/path). Currently '*' is not allowed.

* `roles` -
  (Optional)
  A list of IAM roles that represent the set of operations that the sources
  specified in the corresponding `EgressFrom`
  are allowed to perform.

* `operations` -
  (Optional)
  A list of `ApiOperations` that this egress rule applies to. A request matches
  if it contains an operation/service in this list.
  Structure is [documented below](#nested_spec_egress_policies_egress_policies_egress_to_operations).


<a name="nested_spec_egress_policies_egress_policies_egress_to_operations"></a>The `operations` block supports:

* `service_name` -
  (Optional)
  The name of the API whose methods or permissions the `IngressPolicy` or
  `EgressPolicy` want to allow. A single `ApiOperation` with serviceName
  field set to `*` will allow all methods AND permissions for all services.

* `method_selectors` -
  (Optional)
  API methods or permissions to allow. Method or permission must belong
  to the service specified by `serviceName` field. A single MethodSelector
  entry with `*` specified for the `method` field will allow all methods
  AND permissions for the service specified in `serviceName`.
  Structure is [documented below](#nested_spec_egress_policies_egress_policies_egress_to_operations_operations_method_selectors).


<a name="nested_spec_egress_policies_egress_policies_egress_to_operations_operations_method_selectors"></a>The `method_selectors` block supports:

* `method` -
  (Optional)
  Value for `method` should be a valid method name for the corresponding
  `serviceName` in `ApiOperation`. If `*` used as value for method,
  then ALL methods and permissions are allowed.

* `permission` -
  (Optional)
  Value for permission should be a valid Cloud IAM permission for the
  corresponding `serviceName` in `ApiOperation`.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{name}}`

* `create_time` -
  Time the AccessPolicy was created in UTC.

* `update_time` -
  Time the AccessPolicy was updated in UTC.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


ServicePerimeter can be imported using any of these accepted formats:

* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ServicePerimeter using one of the formats above. For example:

```tf
import {
  id = "{{name}}"
  to = google_access_context_manager_service_perimeter.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), ServicePerimeter can be imported using one of the formats above. For example:

```
$ terraform import google_access_context_manager_service_perimeter.default {{name}}
```
