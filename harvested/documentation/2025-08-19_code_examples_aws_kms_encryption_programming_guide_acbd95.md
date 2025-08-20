---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T18:48:24.678522'
profile: code_examples
source: https://docs.aws.amazon.com/kms/latest/developerguide/programming-encryption.html
topic: AWS KMS Encryption Programming Guide
---

# AWS KMS Encryption Programming Guide

## Select your cookie preferences
We use essential cookies and similar tools that are necessary to provide our site and services. We use performance cookies to collect anonymous statistics, so we can understand how customers use our site and make improvements. Essential cookies cannot be deactivated, but you can choose “Customize” or “Decline” to decline performance cookies. If you agree, AWS and approved third parties will also use cookies to provide useful site features, remember your preferences, and display relevant content, including relevant advertising. To accept or decline all non-essential cookies, choose “Accept” or “Decline.” To make more detailed choices, choose “Customize.”
AcceptDeclineCustomize
## Customize cookie preferences
We use cookies and similar tools (collectively, "cookies") for the following purposes.
### Essential
Essential cookies are necessary to provide our site and services and cannot be deactivated. They are usually set in response to your actions on the site, such as setting your privacy preferences, signing in, or filling in forms. 
### Performance
Performance cookies provide anonymous statistics about how customers navigate our site so we can improve site experience and performance. Approved third parties may perform analytics on our behalf, but they cannot use the data for their own purposes.
Allowed
### Functional
Functional cookies help us provide useful site features, remember your preferences, and display relevant content. Approved third parties may set these cookies to provide certain site features. If you do not allow these cookies, then some or all of these services may not function properly.
Allowed
### Advertising
Advertising cookies may be set through our site by us or our advertising partners and help us deliver relevant marketing content. If you do not allow these cookies, you will experience less relevant advertising.
Allowed
Blocking some types of cookies may impact your experience of our sites. You may review and change your choices at any time by selecting Cookie preferences in the footer of this site. We and selected third-parties use cookies or similar technologies as specified in the [AWS Cookie Notice](https://aws.amazon.com/legal/cookies/).
CancelSave preferences
## Unable to save cookie preferences
We will only store essential cookies at this time, because we were unable to save your cookie preferences.If you want to change your cookie preferences, try again later using the link in the AWS console footer, or contact support if the problem persists.
Dismiss
[English](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html "Language Selector. Currently set to: English")
[Preferences ](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html)
[Contact Us](https://aws.amazon.com/contact-us/?cmpid=docs_headercta_contactus)
[Feedback](https://docs.aws.amazon.com/forms/aws-doc-feedback?hidden_service_name=Key%20Management%20Service%20\(KMS\)&topic_url=https://docs.aws.amazon.com/kms/latest/developerguide/overview.html)
[![AWS Documentation](https://docs.aws.amazon.com/assets/r/images/aws_logo_light.svg)](https://docs.aws.amazon.com)
[Get started](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html)
[Service guides](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html)
[Developer tools](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html)
[AI resources](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html)
[Create an AWS Account](https://portal.aws.amazon.com)
## AWS Key Management Service
### Developer Guide
  * [AWS Key Management Service](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html)
  * [Accessing AWS Key Management Service](https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html)
    * [Working with AWS SDKs](https://docs.aws.amazon.com/kms/latest/developerguide/sdk-general-information-section.html)
    * [Hybrid post-quantum TLS](https://docs.aws.amazon.com/kms/latest/developerguide/pqtls.html)
      * [Configure hybrid post-quantum TLS](https://docs.aws.amazon.com/kms/latest/developerguide/pqtls-how-to.html)
    * [Connect to AWS KMS through a VPC endpoint](https://docs.aws.amazon.com/kms/latest/developerguide/kms-vpc-endpoint.html)
      * [Create a VPC endpoint for AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/vpce-create-endpoint.html)
      * [Connect to a VPC endpoint](https://docs.aws.amazon.com/kms/latest/developerguide/vpce-connect.html)
      * [Use VPC endpoints to control access to AWS KMS resources](https://docs.aws.amazon.com/kms/latest/developerguide/vpce-policy-condition.html)
      * [Logging AWS KMS requests that use a VPC endpoint](https://docs.aws.amazon.com/kms/latest/developerguide/vpce-logging.html)
    * [Dual-stack endpoints](https://docs.aws.amazon.com/kms/latest/developerguide/ipv6-kms.html)
  * [Concepts](https://docs.aws.amazon.com/kms/latest/developerguide/concepts-intro.html)
    * [AWS KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html)
      * [Asymmetric keys](https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
      * [HMAC keys](https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html)
      * [ML-DSA keys](https://docs.aws.amazon.com/kms/latest/developerguide/mldsa.html)
      * [Multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)
        * [Security considerations for multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/mrk-when-to-use.html)
        * [How multi-Region keys work](https://docs.aws.amazon.com/kms/latest/developerguide/mrk-how-it-works.html)
      * [Imported key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html)
        * [Special considerations for imported key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-considerations.html)
        * [Protecting imported key material](https://docs.aws.amazon.com/kms/latest/developerguide/import-keys-protect.html)
      * [KMS keys in a CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/manage-cmk-keystore.html)
      * [KMS keys in external key stores](https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external-key-manage.html)
    * [AWS KMS cryptography essentials](https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html)
    * [Supported algorithms](https://docs.aws.amazon.com/kms/latest/developerguide/supported-algorithms.html)
  * [KMS key access and permissions](https://docs.aws.amazon.com/kms/latest/developerguide/control-access.html)
    * [Key policies](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)
      * [Creating a key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-overview.html)
      * [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html)
      * [View a key policies](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-viewing.html)
      * [Change a key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying.html)
      * [Permissions for AWS services](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-services.html)
    * [IAM policies](https://docs.aws.amazon.com/kms/latest/developerguide/iam-policies.html)
      * [Best practices for IAM policies](https://docs.aws.amazon.com/kms/latest/developerguide/iam-policies-best-practices.html)
      * [Specifying KMS keys in IAM policy statements](https://docs.aws.amazon.com/kms/latest/developerguide/cmks-in-iam-policies.html)
      * [Examples](https://docs.aws.amazon.com/kms/latest/developerguide/customer-managed-policies.html)
    * [Resource control policies](https://docs.aws.amazon.com/kms/latest/developerguide/resource-control-policies.html)
    * [Grants](https://docs.aws.amazon.com/kms/latest/developerguide/grants.html)
      * [Best practices](https://docs.aws.amazon.com/kms/latest/developerguide/grant-best-practices.html)
      * [Controlling access to grants](https://docs.aws.amazon.com/kms/latest/developerguide/grant-authorization.html)
      * [Creating grants](https://docs.aws.amazon.com/kms/latest/developerguide/create-grant-overview.html)
      * [Viewing grants](https://docs.aws.amazon.com/kms/latest/developerguide/grant-view.html)
      * [Using a grant token](https://docs.aws.amazon.com/kms/latest/developerguide/using-grant-token.html)
      * [Retiring and revoking grants](https://docs.aws.amazon.com/kms/latest/developerguide/grant-delete.html)
    * [Condition keys](https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html)
      * [AWS global condition keys](https://docs.aws.amazon.com/kms/latest/developerguide/conditions-aws.html)
      * [AWS KMS condition keys](https://docs.aws.amazon.com/kms/latest/developerguide/conditions-kms.html)
      * [AWS KMS condition keys for AWS Nitro Enclaves](https://docs.aws.amazon.com/kms/latest/developerguide/conditions-nitro-enclaves.html)
    * [Least-privilege permissions](https://docs.aws.amazon.com/kms/latest/developerguide/least-privilege.html)
    * [Attribute-based access control (ABAC)](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html)
      * [Troubleshooting ABAC for AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/troubleshooting-tags-aliases.html)
    * [Role-based access control (RBAC)](https://docs.aws.amazon.com/kms/latest/developerguide/rbac.html)
    * [Cross-account access](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html)
    * [Control access to multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-auth.html)
    * [Determining access](https://docs.aws.amazon.com/kms/latest/developerguide/determining-access.html)
      * [Examining the key policy](https://docs.aws.amazon.com/kms/latest/developerguide/determining-access-key-policy.html)
      * [Examining IAM policies](https://docs.aws.amazon.com/kms/latest/developerguide/determining-access-iam-policies.html)
      * [Examining grants](https://docs.aws.amazon.com/kms/latest/developerguide/determining-access-grants.html)
    * [Encryption context](https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html)
    * [Testing your permissions](https://docs.aws.amazon.com/kms/latest/developerguide/testing-permissions.html)
    * [Troubleshooting AWS KMS permissions](https://docs.aws.amazon.com/kms/latest/developerguide/policy-evaluation.html)
    * [Glossary](https://docs.aws.amazon.com/kms/latest/developerguide/access-glossary.html)
  * [Create a KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html)
    * [Create a symmetric encryption KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/create-symmetric-cmk.html)
    * [Create an asymmetric KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/asymm-create-key.html)
    * [Create an HMAC KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/hmac-create-key.html)
    * [Create multi-Region primary keys](https://docs.aws.amazon.com/kms/latest/developerguide/create-primary-keys.html)
    * [Create multi-Region replica keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-replicate.html)
    * [Create a KMS key with imported key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-conceptual.html)
      * [Step 1: Create an AWS KMS key without key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-create-cmk.html)
      * [Step 2: Download the wrapping public key and import token](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-get-public-key-and-token.html)
      * [Step 3: Encrypt the key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-encrypt-key-material.html)
      * [Step 4: Import the key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html)
    * [Create a KMS key in an AWS CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html)
    * [Create a KMS key in external key stores](https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keys.html)
  * [Identify and view keys](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html)
    * [Find the key ID and key ARN](https://docs.aws.amazon.com/kms/latest/developerguide/find-cmk-id-arn.html)
    * [Access and list KMS key details](https://docs.aws.amazon.com/kms/latest/developerguide/finding-keys.html)
    * [Identify different key types](https://docs.aws.amazon.com/kms/latest/developerguide/identify-key-types.html)
    * [Customize your console view](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-console-customize.html)
    * [Find KMS keys and key material in an AWS CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/find-key-material.html)
      * [Find the KMS keys in an AWS CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/find-cmk-in-keystore.html)
      * [Find all keys for an AWS CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/find-all-kmsuser-keys.html)
      * [Find the KMS key for an AWS CloudHSM key](https://docs.aws.amazon.com/kms/latest/developerguide/find-label-for-key-handle.html)
      * [Find the AWS CloudHSM key for a KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/find-handle-for-cmk-id.html)
  * [Enable and disable keys](https://docs.aws.amazon.com/kms/latest/developerguide/enabling-keys.html)
  * [Rotate keys](https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
    * [Enable automatic key rotation](https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable.html)
    * [Disable automatic key rotation](https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-disable.html)
    * [Perform on-demand key rotation](https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-on-demand.html)
    * [List rotations and key materials](https://docs.aws.amazon.com/kms/latest/developerguide/list-rotations.html)
    * [Rotate keys manually](https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys-manually.html)
    * [Change the primary key in a set of multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-update.html)
  * [Delete keys](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html)
    * [Control access to key deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys-adding-permission.html)
    * [Schedule key deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys-scheduling-key-deletion.html)
    * [Cancel key deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys-cancelling-key-deletion.html)
    * [Create an alarm](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys-creating-cloudwatch-alarm.html)
    * [Determine past usage of a KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys-determining-usage.html)
    * [Delete imported key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-delete-key-material.html)
  * [Generate data keys](https://docs.aws.amazon.com/kms/latest/developerguide/data-keys.html)
    * [How unusable KMS keys affect data keys](https://docs.aws.amazon.com/kms/latest/developerguide/unusable-kms-keys.html)
  * [Generate data key pairs](https://docs.aws.amazon.com/kms/latest/developerguide/data-key-pairs.html)
  * [Perform offline operations with public keys](https://docs.aws.amazon.com/kms/latest/developerguide/offline-public-key.html)
    * [Download public key](https://docs.aws.amazon.com/kms/latest/developerguide/download-public-key.html)
    * [Example offline operations](https://docs.aws.amazon.com/kms/latest/developerguide/offline-operations.html)
  * [Monitor keys](https://docs.aws.amazon.com/kms/latest/developerguide/monitoring-overview.html)
    * [Logging with AWS CloudTrail](https://docs.aws.amazon.com/kms/latest/developerguide/logging-using-cloudtrail.html)
      * [Examples of AWS KMS log entries](https://docs.aws.amazon.com/kms/latest/developerguide/understanding-kms-entries.html)
        * [CancelKeyDeletion](https://docs.aws.amazon.com/kms/latest/developerguide/ct-cancel-key-deletion.html)
        * [ConnectCustomKeyStore](https://docs.aws.amazon.com/kms/latest/developerguide/ct-connect-keystore.html)
        * [CreateAlias](https://docs.aws.amazon.com/kms/latest/developerguide/ct-createalias.html)
        * [CreateCustomKeyStore](https://docs.aws.amazon.com/kms/latest/developerguide/ct-create-keystore.html)
        * [CreateGrant](https://docs.aws.amazon.com/kms/latest/developerguide/ct-creategrant.html)
        * [CreateKey](https://docs.aws.amazon.com/kms/latest/developerguide/ct-createkey.html)
        * [Decrypt](https://docs.aws.amazon.com/kms/latest/developerguide/ct-decrypt.html)
        * [DeleteAlias](https://docs.aws.amazon.com/kms/latest/developerguide/ct-deletealias.html)
        * [DeleteCustomKeyStore](https://docs.aws.amazon.com/kms/latest/developerguide/ct-delete-keystore.html)
        * [DeleteExpiredKeyMaterial](https://docs.aws.amazon.com/kms/latest/developerguide/ct-deleteexpiredkeymaterial.html)
        * [DeleteImportedKeyMaterial](https://docs.aws.amazon.com/kms/latest/developerguide/ct-deleteimportedkeymaterial.html)
        * [DeleteKey](https://docs.aws.amazon.com/kms/latest/developerguide/ct-delete-key.html)
        * [DescribeCustomKeyStores](https://docs.aws.amazon.com/kms/latest/developerguide/ct-describe-keystores.html)
        * [DescribeKey](https://docs.aws.amazon.com/kms/latest/developerguide/ct-describekey.html)
        * [DisableKey](https://docs.aws.amazon.com/kms/latest/developerguide/ct-disablekey.html)
        * [DisableKeyRotation](https://docs.aws.amazon.com/kms/latest/developerguide/ct-disable-key-rotation.html)
        * [DisconnectCustomKeyStore](https://docs.aws.amazon.com/kms/latest/developerguide/ct-disconnect-keystore.html)
        * [EnableKey](https://docs.aws.amazon.com/kms/latest/developerguide/ct-enablekey.html)
        * [EnableKeyRotation](https://docs.aws.amazon.com/kms/latest/developerguide/ct-enablekeyrotation.html)
        * [Encrypt](https://docs.aws.amazon.com/kms/latest/developerguide/ct-encrypt.html)
        * [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/developerguide/ct-generatedatakey.html)
        * [GenerateDataKeyPair](https://docs.aws.amazon.com/kms/latest/developerguide/ct-generatedatakeypair.html)
        * [GenerateDataKeyPairWithoutPlaintext](https://docs.aws.amazon.com/kms/latest/developerguide/ct-generatedatakeypairwithoutplaintext.html)
        * [GenerateDataKeyWithoutPlaintext](https://docs.aws.amazon.com/kms/latest/developerguide/ct-generatedatakeyplaintext.html)
        * [GenerateMac](https://docs.aws.amazon.com/kms/latest/developerguide/ct-generatemac.html)
        * [GenerateRandom](https://docs.aws.amazon.com/kms/latest/developerguide/ct-generaterandom.html)
        * [GetKeyPolicy](https://docs.aws.amazon.com/kms/latest/developerguide/ct-getkeypolicy.html)
        * [GetKeyRotationStatus](https://docs.aws.amazon.com/kms/latest/developerguide/ct-getkeyrotationstatus.html)
        * [GetParametersForImport](https://docs.aws.amazon.com/kms/latest/developerguide/ct-getparametersforimport.html)
        * [ImportKeyMaterial](https://docs.aws.amazon.com/kms/latest/developerguide/ct-importkeymaterial.html)
        * [ListAliases](https://docs.aws.amazon.com/kms/latest/developerguide/ct-listaliases.html)
        * [ListGrants](https://docs.aws.amazon.com/kms/latest/developerguide/ct-listgrants.html)
        * [ListKeyRotations](https://docs.aws.amazon.com/kms/latest/developerguide/ct-listkeyrotations.html)
        * [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/developerguide/ct-put-key-policy.html)
        * [ReEncrypt](https://docs.aws.amazon.com/kms/latest/developerguide/ct-reencrypt.html)
        * [ReplicateKey](https://docs.aws.amazon.com/kms/latest/developerguide/ct-replicate-key.html)
        * [RetireGrant](https://docs.aws.amazon.com/kms/latest/developerguide/ct-retire-grant.html)
        * [RevokeGrant](https://docs.aws.amazon.com/kms/latest/developerguide/ct-revoke-grant.html)
        * [RotateKey](https://docs.aws.amazon.com/kms/latest/developerguide/ct-rotatekey.html)
        * [RotateKeyOnDemand](https://docs.aws.amazon.com/kms/latest/developerguide/ct-rotatekeyondemand.html)
        * [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/developerguide/ct-schedule-key-deletion.html)
        * [Sign](https://docs.aws.amazon.com/kms/latest/developerguide/ct-sign.html)
        * [SynchronizeMultiRegionKey](https://docs.aws.amazon.com/kms/latest/developerguide/ct-synchronize-multi-region-key.html)
        * [TagResource](https://docs.aws.amazon.com/kms/latest/developerguide/ct-tagresource.html)
        * [UntagResource](https://docs.aws.amazon.com/kms/latest/developerguide/ct-untagresource.html)
        * [UpdateAlias](https://docs.aws.amazon.com/kms/latest/developerguide/ct-updatealias.html)
        * [UpdateCustomKeyStore](https://docs.aws.amazon.com/kms/latest/developerguide/ct-update-keystore.html)
        * [UpdateKeyDescription](https://docs.aws.amazon.com/kms/latest/developerguide/ct-update-key-description.html)
        * [UpdatePrimaryRegion](https://docs.aws.amazon.com/kms/latest/developerguide/ct-update-primary-region.html)
        * [VerifyMac](https://docs.aws.amazon.com/kms/latest/developerguide/ct-verifymac.html)
        * [Verify](https://docs.aws.amazon.com/kms/latest/developerguide/ct-verify.html)
        * [Amazon EC2 example one](https://docs.aws.amazon.com/kms/latest/developerguide/ct-ec2one.html)
        * [Amazon EC2 example two](https://docs.aws.amazon.com/kms/latest/developerguide/ct-ec2two.html)
    * [Monitor keys with CloudWatch](https://docs.aws.amazon.com/kms/latest/developerguide/monitoring-cloudwatch.html)
      * [Create a CloudWatch alarm for expiration of imported key material](https://docs.aws.amazon.com/kms/latest/developerguide/imported-key-material-expiration-alarm.html)
      * [Create CloudWatch alarms for external key stores](https://docs.aws.amazon.com/kms/latest/developerguide/xks-alarms.html)
    * [Monitor keys with Amazon EventBridge](https://docs.aws.amazon.com/kms/latest/developerguide/kms-events.html)
  * [Aliases](https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html)
    * [Controlling access to aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-access.html)
    * [Create aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-create.html)
    * [Find the alias name and alias ARN](https://docs.aws.amazon.com/kms/latest/developerguide/alias-view.html)
    * [Update aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-update.html)
    * [Delete an alias](https://docs.aws.amazon.com/kms/latest/developerguide/alias-delete.html)
    * [Use aliases to control access to KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/alias-authorization.html)
    * [Learn how to use aliases in your applications](https://docs.aws.amazon.com/kms/latest/developerguide/alias-using.html)
    * [Find aliases in AWS CloudTrail logs](https://docs.aws.amazon.com/kms/latest/developerguide/alias-ct.html)
  * [Tags](https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html)
    * [Controlling access to tags](https://docs.aws.amazon.com/kms/latest/developerguide/tag-permissions.html)
    * [Add tags](https://docs.aws.amazon.com/kms/latest/developerguide/add-tags.html)
    * [Edit tags](https://docs.aws.amazon.com/kms/latest/developerguide/edit-tags.html)
    * [Remove tags](https://docs.aws.amazon.com/kms/latest/developerguide/remove-tags.html)
    * [View tags](https://docs.aws.amazon.com/kms/latest/developerguide/view-tags.html)
    * [Use tags to control access to KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/tag-authorization.html)
  * [Key stores](https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html)
    * [AWS CloudHSM key stores](https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html)
      * [Control access to your AWS CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/authorize-key-store.html)
      * [Create an AWS CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/create-keystore.html)
      * [View an AWS CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/view-keystore.html)
      * [Edit AWS CloudHSM key store settings](https://docs.aws.amazon.com/kms/latest/developerguide/update-keystore.html)
      * [Connect an AWS CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/connect-keystore.html)
      * [Disconnect an AWS CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/disconnect-keystore.html)
      * [Delete an AWS CloudHSM key store](https://docs.aws.amazon.com/kms/latest/developerguide/delete-keystore.html)
      * [Troubleshooting a custom key store](https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html)
    * [External key stores](https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html)
      * [Control access to your external key store](https://docs.aws.amazon.com/kms/latest/developerguide/authorize-xks-key-store.html)
      * [Choose a proxy connectivity option](https://docs.aws.amazon.com/kms/latest/developerguide/choose-xks-connectivity.html)
        * [Configure VPC endpoint service connectivity](https://docs.aws.amazon.com/kms/latest/developerguide/vpc-connectivity.html)
      * [Create an external key store](https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keystore.html)
      * [Edit external key store properties](https://docs.aws.amazon.com/kms/latest/developerguide/update-xks-keystore.html)
      * [View external key stores](https://docs.aws.amazon.com/kms/latest/developerguide/view-xks-keystore.html)
      * [Monitor external key stores](https://docs.aws.amazon.com/kms/latest/developerguide/xks-monitoring.html)
      * [Connect and disconnect external key stores](https://docs.aws.amazon.com/kms/latest/developerguide/xks-connect-disconnect.html)
        * [Connect an external key store](https://docs.aws.amazon.com/kms/latest/developerguide/about-xks-connecting.html)
        * [Disconnect an external key store](https://docs.aws.amazon.com/kms/latest/developerguide/about-xks-disconnecting.html)
      * [Delete an external key store](https://docs.aws.amazon.com/kms/latest/developerguide/delete-xks.html)
      * [Troubleshooting external key stores](https://docs.aws.amazon.com/kms/latest/developerguide/xks-troubleshooting.html)
  * [Security](https://docs.aws.amazon.com/kms/latest/developerguide/kms-security.html)
    * [Data protection](https://docs.aws.amazon.com/kms/latest/developerguide/data-protection.html)
    * [Identity and access management](https://docs.aws.amazon.com/kms/latest/developerguide/security-iam.html)
      * [AWS managed policies](https://docs.aws.amazon.com/kms/latest/developerguide/security-iam-awsmanpol.html)
      * [Service-linked roles](https://docs.aws.amazon.com/kms/latest/developerguide/using-service-linked-roles.html)
        * [Authorizing AWS KMS to manage AWS CloudHSM and Amazon EC2 resources](https://docs.aws.amazon.com/kms/latest/developerguide/authorize-kms.html)
        * [Authorizing AWS KMS to synchronize multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-auth-slr.html)
    * [Logging and monitoring](https://docs.aws.amazon.com/kms/latest/developerguide/security-logging-monitoring.html)
    * [Compliance validation](https://docs.aws.amazon.com/kms/latest/developerguide/kms-compliance.html)
    * [Resilience](https://docs.aws.amazon.com/kms/latest/developerguide/disaster-recovery-resiliency.html)
    * [Infrastructure security](https://docs.aws.amazon.com/kms/latest/developerguide/infrastructure-security.html)
  * [Quotas](https://docs.aws.amazon.com/kms/latest/developerguide/limits.html)
    * [Resource quotas](https://docs.aws.amazon.com/kms/latest/developerguide/resource-limits.html)
    * [Request quotas](https://docs.aws.amazon.com/kms/latest/developerguide/requests-per-second.html)
    * [Throttling requests](https://docs.aws.amazon.com/kms/latest/developerguide/throttling.html)
  * [Code examples](https://docs.aws.amazon.com/kms/latest/developerguide/service_code_examples.html)
    * [Basics](https://docs.aws.amazon.com/kms/latest/developerguide/service_code_examples_basics.html)
      * [Hello AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_Hello_section.html)
      * [Learn the basics](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_Scenario_Basics_section.html)
      * [Actions](https://docs.aws.amazon.com/kms/latest/developerguide/service_code_examples_actions.html)
        * [CreateAlias](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateAlias_section.html)
        * [CreateGrant](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html)
        * [CreateKey](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateKey_section.html)
        * [Decrypt](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_Decrypt_section.html)
        * [DeleteAlias](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_DeleteAlias_section.html)
        * [DescribeKey](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_DescribeKey_section.html)
        * [DisableKey](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_DisableKey_section.html)
        * [EnableKey](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_EnableKey_section.html)
        * [EnableKeyRotation](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_EnableKeyRotation_section.html)
        * [Encrypt](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_Encrypt_section.html)
        * [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_GenerateDataKey_section.html)
        * [GenerateDataKeyWithoutPlaintext](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_GenerateDataKeyWithoutPlaintext_section.html)
        * [GenerateRandom](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_GenerateRandom_section.html)
        * [GetKeyPolicy](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_GetKeyPolicy_section.html)
        * [ListAliases](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_ListAliases_section.html)
        * [ListGrants](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_ListGrants_section.html)
        * [ListKeyPolicies](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_ListKeyPolicies_section.html)
        * [ListKeys](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_ListKeys_section.html)
        * [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_PutKeyPolicy_section.html)
        * [ReEncrypt](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_ReEncrypt_section.html)
        * [RetireGrant](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_RetireGrant_section.html)
        * [RevokeGrant](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_RevokeGrant_section.html)
        * [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_ScheduleKeyDeletion_section.html)
        * [Sign](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_Sign_section.html)
        * [TagResource](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_TagResource_section.html)
        * [UpdateAlias](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_UpdateAlias_section.html)
        * [Verify](https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_Verify_section.html)
    * [Scenarios](https://docs.aws.amazon.com/kms/latest/developerguide/service_code_examples_scenarios.html)
      * [Work with table encryption](https://docs.aws.amazon.com/kms/latest/developerguide/example_dynamodb_Scenario_EncryptionExamples_section.html)
  * [Cryptographic attestation for AWS Nitro Enclaves](https://docs.aws.amazon.com/kms/latest/developerguide/services-nitro-enclaves.html)
    * [How to call AWS KMS APIs for a Nitro enclave](https://docs.aws.amazon.com/kms/latest/developerguide/how-nitro-enclaves.html)
    * [Monitoring requests for Nitro enclaves](https://docs.aws.amazon.com/kms/latest/developerguide/ct-nitro-enclave.html)
  * [Encrypting AWS services](https://docs.aws.amazon.com/kms/latest/developerguide/service-integration.html)
    * [Amazon Elastic Block Store (Amazon EBS)](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html)
    * [Amazon EMR](https://docs.aws.amazon.com/kms/latest/developerguide/services-emr.html)
    * [Amazon Redshift](https://docs.aws.amazon.com/kms/latest/developerguide/services-redshift.html)
  * [Reference](https://docs.aws.amazon.com/kms/latest/developerguide/references.html)
    * [Key state reference](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
    * [Key type reference](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-compare.html)
    * [Key spec reference](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose-key-spec.html)
    * [Permissions reference](https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
    * [AWS KMS internal operations](https://docs.aws.amazon.com/kms/latest/developerguide/kms-internals.html)
      * [Domains and domain state](https://docs.aws.amazon.com/kms/latest/developerguide/domains-and-domain-state.html)
      * [Internal communication security](https://docs.aws.amazon.com/kms/latest/developerguide/internal-communication-security.html)
      * [Replication process for multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/replicate-key-details.html)
      * [Durability protection](https://docs.aws.amazon.com/kms/latest/developerguide/durability-protection.html)
  * [Document history](https://docs.aws.amazon.com/kms/latest/developerguide/dochistory.html)


  1. [Documentation](https://docs.aws.amazon.com/index.html)
  2. ...
  3. [AWS KMS](https://docs.aws.amazon.com/kms/index.html)
  4. Developer Guide


  1. [Documentation](https://docs.aws.amazon.com/index.html)
  2. [AWS KMS](https://docs.aws.amazon.com/kms/index.html)
  3. Developer Guide


# 
AWS Key Management Service
[ PDF](https://docs.aws.amazon.com/pdfs/kms/latest/developerguide/kms-dg.pdf#overview)
[ RSS](https://docs.aws.amazon.com/kms/latest/developerguide/aws-kms-document-history.rss)
Focus mode
AWS Key Management Service - AWS Key Management Service
[](https://docs.aws.amazon.com/pdfs/kms/latest/developerguide/kms-dg.pdf#overview "Open PDF")
[Documentation](https://docs.aws.amazon.com/index.html)[AWS KMS](https://docs.aws.amazon.com/kms/index.html)[Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html)
[Why use AWS KMS?](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html#service-kms-why)[AWS KMS in AWS Regions](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html#kms_regions)[AWS KMS pricing](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html#pricing)[AWS KMS service level agreement](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html#kms_service_levels)
AWS Key Management Service (AWS KMS) is an AWS managed service that makes it easy for you to create and control the encryption keys that are used to encrypt your data. The AWS KMS keys that you create in AWS KMS are protected by [FIPS 140-3 Security Level 3 validated hardware security modules (HSM)](https://csrc.nist.gov/projects/cryptographic-module-validation-program/certificate/4884). They never leave AWS KMS unencrypted. To use or manage your KMS keys, you interact with AWS KMS.
## Why use AWS KMS?
When you encrypt data, you need to protect your encryption key. If you encrypt your key, you need to protect its encryption key. Eventually, you must protect the highest level encryption key (known as a _root key_) in the hierarchy that protects your data. That's where AWS KMS comes in.
![Root key protect the data keys that protect your data](https://docs.aws.amazon.com/images/kms/latest/developerguide/images/key-hierarchy-root.png)
AWS KMS protects your root keys. KMS keys are created, managed, used, and deleted entirely within AWS KMS. They never leave the service unencrypted. To use or manage your KMS keys, you call AWS KMS.
![AWS KMS protects your root keys](https://docs.aws.amazon.com/images/kms/latest/developerguide/images/key-hierarchy-kms-key.png)
Additionally, you can create and manage [key policies](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html) in AWS KMS, ensuring that only trusted users have access to KMS keys.
## AWS KMS in AWS Regions
The AWS Regions in which AWS KMS is supported are listed in [AWS Key Management Service Endpoints and Quotas](https://docs.aws.amazon.com/general/latest/gr/kms.html). If an AWS KMS feature is not supported in an AWS Region that AWS KMS supports, the regional difference is described in the topic about the feature. 
## AWS KMS pricing
As with other AWS products, using AWS KMS does not require contracts or minimum purchases. For more information about AWS KMS pricing, see [AWS Key Management Service Pricing](https://aws.amazon.com/kms/pricing/).
## AWS KMS service level agreement
AWS Key Management Service is backed by a [service level agreement](https://aws.amazon.com/kms/sla/) that defines our service availability policy.
[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)
Accessing AWS Key Management Service
Did this page help you? - Yes
Thanks for letting us know we're doing a good job!
If you've got a moment, please tell us what we did right so we can do more of it.
Did this page help you? - No
Thanks for letting us know this page needs work. We're sorry we let you down.
If you've got a moment, please tell us how we can make the documentation better.
  * ### On this page
    1. [Why use AWS KMS?](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html#service-kms-why)
    2. [AWS KMS in AWS Regions](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html#kms_regions)
    3. [AWS KMS pricing](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html#pricing)
    4. [AWS KMS service level agreement](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html#kms_service_levels)
  * #### Did this page help you?
Yes
No
[Provide feedback](https://docs.aws.amazon.com/forms/aws-doc-feedback?hidden_service_name=Key%20Management%20Service%20%28KMS%29&topic_url=https://docs.aws.amazon.com/en_us/kms/latest/developerguide/overview.html)


#### Next topic:
[Accessing AWS Key Management Service](https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html)
#### Need help?
  * [Try AWS re:Post](https://repost.aws/tags/TAMC3vcPOPTF-rPAHZVRj1PQ)
  * [Connect with an AWS IQ expert](https://iq.aws.amazon.com/get-started/?utm=docs&service=AWS%20Key%20Management%20Service)


[Privacy](https://aws.amazon.com/privacy)[Site terms](https://aws.amazon.com/terms)[Cookie preferences](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html)
© 2025, Amazon Web Services, Inc. or its affiliates. All rights reserved.
## Root key protect the data keys that protect your data
![Root key protect the data keys that protect your data](https://docs.aws.amazon.com/images/kms/latest/developerguide/images/key-hierarchy-root.png)
Close
## AWS KMS protects your root keys
![AWS KMS protects your root keys](https://docs.aws.amazon.com/images/kms/latest/developerguide/images/key-hierarchy-kms-key.png)
Close


## Source Information
- URL: https://docs.aws.amazon.com/kms/latest/developerguide/programming-encryption.html
- Harvested: 2025-08-19T18:48:24.678522
- Profile: code_examples
- Agent: architect
