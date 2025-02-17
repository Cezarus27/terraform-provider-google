---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "Apigee"
description: |-
  An `Environment Reference` in Apigee.
---

# google_apigee_env_references

An `Environment Reference` in Apigee.


To get more information about EnvReferences, see:

* [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.references/create)
* How-to Guides
    * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)


## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  Required. The resource id of this reference. Values must match the regular expression [\w\s-.]+.

* `resource_type` -
  (Required)
  The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.

* `refers` -
  (Required)
  Required. The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resourceType.

* `env_id` -
  (Required)
  The Apigee environment group associated with the Apigee environment,
  in the format `organizations/{{org_name}}/environments/{{env_name}}`.


- - -


* `description` -
  (Optional)
  Optional. A human-readable description of this reference.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{env_id}}/references/{{name}}`


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 1 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 1 minutes.

## Import


EnvReferences can be imported using any of these accepted formats:

* `{{env_id}}/references/{{name}}`
* `{{env_id}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import EnvReferences using one of the formats above. For example:

```tf
import {
  id = "{{env_id}}/references/{{name}}"
  to = google_apigee_env_references.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), EnvReferences can be imported using one of the formats above. For example:

```
$ terraform import google_apigee_env_references.default {{env_id}}/references/{{name}}
$ terraform import google_apigee_env_references.default {{env_id}}/{{name}}
```
