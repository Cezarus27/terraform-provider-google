---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/storage_bucket_iam_policy.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud Storage"
description: |-
  A datasource to retrieve the IAM policy state for Cloud Storage Bucket
---


# google_storage_bucket_iam_policy

Retrieves the current IAM policy data for bucket


## Example Usage


```hcl
data "google_storage_bucket_iam_policy" "policy" {
  bucket = google_storage_bucket.default.name
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) Used to find the parent resource to bind the IAM policy to

## Attributes Reference

The attributes are exported:

* `etag` - (Computed) The etag of the IAM policy.

* `policy_data` - (Required only by `google_storage_bucket_iam_policy`) The policy data generated by
  a `google_iam_policy` data source.
