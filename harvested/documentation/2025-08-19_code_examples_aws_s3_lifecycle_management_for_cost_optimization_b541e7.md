---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T20:33:35.371186'
profile: code_examples
source: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html
topic: AWS S3 Lifecycle Management for Cost Optimization
---

# AWS S3 Lifecycle Management for Cost Optimization

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
[English](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html "Language Selector. Currently set to: English")
[Preferences ](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html)
[Contact Us](https://aws.amazon.com/contact-us/?cmpid=docs_headercta_contactus)
[Feedback](https://docs.aws.amazon.com/forms/aws-doc-feedback?hidden_service_name=S3&topic_url=https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html)
[![AWS Documentation](https://docs.aws.amazon.com/assets/r/images/aws_logo_light.svg)](https://docs.aws.amazon.com)
[Get started](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html)
[Service guides](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html)
[Developer tools](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html)
[AI resources](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html)
[Create an AWS Account](https://portal.aws.amazon.com)
## Amazon Simple Storage Service
### User Guide
  1. [Documentation](https://docs.aws.amazon.com/index.html)
  2. ...
  3. [Amazon Simple Storage Service (S3)](https://docs.aws.amazon.com/s3/index.html)
  4. User Guide


  1. [Documentation](https://docs.aws.amazon.com/index.html)
  2. [Amazon Simple Storage Service (S3)](https://docs.aws.amazon.com/s3/index.html)
  3. User Guide


# 
Managing the lifecycle of objects
[ PDF](https://docs.aws.amazon.com/pdfs/AmazonS3/latest/userguide/s3-userguide.pdf#object-lifecycle-mgmt)
[ RSS](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-userguide-rss-updates.rss)
Focus mode
Managing the lifecycle of objects - Amazon Simple Storage Service
[](https://docs.aws.amazon.com/pdfs/AmazonS3/latest/userguide/s3-userguide.pdf#object-lifecycle-mgmt "Open PDF")
[Documentation](https://docs.aws.amazon.com/index.html)[Amazon Simple Storage Service (S3)](https://docs.aws.amazon.com/s3/index.html)[User Guide](https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html)
[Managing the complete lifecycle of objects](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html#lifecycle-config-overview-what)
S3 Lifecycle helps you store objects cost effectively throughout their lifecycle by transitioning them to lower-cost storage classes, or, deleting expired objects on your behalf. To manage the lifecycle of your objects, create an _S3 Lifecycle configuration_ for your bucket. An S3 Lifecycle configuration is a set of rules that define actions that Amazon S3 applies to a group of objects. There are two types of actions:
  * **Transition actions** – These actions define when objects transition to another storage class. For example, you might choose to transition objects to the S3 Standard-IA storage class 30 days after creating them, or archive objects to the S3 Glacier Flexible Retrieval storage class one year after creating them. For more information, see [Understanding and managing Amazon S3 storage classes](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html). 
There are costs associated with lifecycle transition requests. For pricing information, see [Amazon S3 pricing](https://aws.amazon.com/s3/pricing/).
  * **Expiration actions** – These actions define when objects expire. Amazon S3 deletes expired objects on your behalf. For example, you might to choose to expire objects after they have been stored for a regulatory compliance period. For more information, see [Expiring objects](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-expire-general-considerations.html).
There are potential costs associated with lifecycle expiration only when you expire objects in a storage class with a minimum storage duration. For more information, see [Minimum storage duration charge](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-expire-general-considerations.html#lifecycle-expire-minimum-storage).


###### Important
**General purpose buckets** — You can't use a bucket policy to prevent deletions or transitions by an S3 Lifecycle rule. For example, even if your bucket policy denies all actions for all principals, your S3 Lifecycle configuration still functions as normal.
###### Existing and new objects
When you add a Lifecycle configuration to a bucket, the configuration rules apply to both existing objects and objects that you add later. For example, if you add a Lifecycle configuration rule today with an expiration action that causes objects to expire 30 days after creation, Amazon S3 will queue for removal any existing objects that are more than 30 days old.
###### Changes in billing
If there is any delay between when an object becomes eligible for a lifecycle action and when Amazon S3 transfers or expires your object, billing changes are applied as soon as the object becomes eligible for the lifecycle action. For example, if an object is scheduled to expire and Amazon S3 doesn't immediately expire the object, you won't be charged for storage after the expiration time. 
The one exception to this behavior is if you have a lifecycle rule to transition to the S3 Intelligent-Tiering storage class. In that case, billing changes don't occur until the object has transitioned to S3 Intelligent-Tiering. For more information about S3 Lifecycle rules, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html). 
###### Note
There are no data retrieval charges for lifecycle transitions. However, there are per-request ingestion charges when using `PUT`, `COPY`, or lifecycle rules to move data into any S3 storage class. Consider the ingestion or transition cost before moving objects into any storage class. For more information about cost considerations, see [Amazon S3 pricing](https://aws.amazon.com/s3/pricing/).
###### Monitoring the effect of lifecycle rules
To monitor the effect of updates made by active lifecycle rules, see [How do I monitor the actions taken by my lifecycle rules?](https://docs.aws.amazon.com/AmazonS3/latest/userguide/troubleshoot-lifecycle.html#troubleshoot-lifecycle-2).
## Managing the complete lifecycle of objects
With S3 Lifecycle configuration rules you can tell Amazon S3 to transition objects to less-expensive storage classes, archive or delete them. For example: 
  * If you upload periodic logs to a bucket, your application might need them for a week or a month. After that, you might want to delete them.
  * Some documents are frequently accessed for a limited period of time. After that, they are infrequently accessed. At some point, you might not need real-time access to them, but your organization or regulations might require you to archive them for a specific period. After that, you can delete them. 
  * You might upload some types of data to Amazon S3 primarily for archival purposes. For example, you might archive digital media, financial, and healthcare records, raw genomics sequence data, long-term database backups, and data that must be retained for regulatory compliance.


By combining S3 Lifecycle actions to manage an object's complete lifecycle. For example, suppose that the objects you create have a well-defined lifecycle. Initially, the objects are frequently accessed for a period of 30 days. Then, objects are infrequently accessed for up to 90 days. After that, the objects are no longer needed, so you might choose to archive or delete them. 
In this scenario, you can create an S3 Lifecycle rule in which you specify the initial transition action to S3 Intelligent-Tiering, S3 Standard-IA, or S3 One Zone-IA storage, another transition action to S3 Glacier Flexible Retrieval storage for archiving, and an expiration action. As you move the objects from one storage class to another, you save on storage costs. For more information about cost considerations, see [Amazon S3 pricing](https://aws.amazon.com/s3/pricing/).
###### Topics
  * [Transitioning objects using Amazon S3 Lifecycle](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-transition-general-considerations.html)
  * [Expiring objects](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-expire-general-considerations.html)
  * [Setting an S3 Lifecycle configuration on a bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/how-to-set-lifecycle-configuration-intro.html)
  * [How S3 Lifecycle interacts with other bucket configurations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-and-other-bucket-config.html)
  * [Configuring S3 Lifecycle event notifications](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-configure-notification.html)
  * [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html)
  * [How Amazon S3 handles conflicts in lifecycle configurations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-conflicts.html)
  * [Examples of S3 Lifecycle configurations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-configuration-examples.html)
  * [Troubleshooting Amazon S3 Lifecycle issues](https://docs.aws.amazon.com/AmazonS3/latest/userguide/troubleshoot-lifecycle.html)


[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)
Restoring an archived object
Transitioning objects
Did this page help you? - Yes
Thanks for letting us know we're doing a good job!
If you've got a moment, please tell us what we did right so we can do more of it.
Did this page help you? - No
Thanks for letting us know this page needs work. We're sorry we let you down.
If you've got a moment, please tell us how we can make the documentation better.
  * ### On this page
    1. [Managing the complete lifecycle of objects](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html#lifecycle-config-overview-what)
  * ### Related resources
[Amazon S3 API Reference](https://docs.aws.amazon.com/AmazonS3/latest/API/Welcome.html)
[AWS CLI commands for Amazon S3](https://docs.aws.amazon.com/cli/latest/reference/s3/)
[SDKs & Tools](https://aws.amazon.com/tools/)
  * #### Did this page help you?
Yes
No
[Provide feedback](https://docs.aws.amazon.com/forms/aws-doc-feedback?hidden_service_name=S3&topic_url=https://docs.aws.amazon.com/en_us/AmazonS3/latest/userguide/object-lifecycle-mgmt.html)


#### Next topic:
[Transitioning objects](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-transition-general-considerations.html)
#### Previous topic:
[Restoring an archived object](https://docs.aws.amazon.com/AmazonS3/latest/userguide/restoring-objects.html)
#### Need help?
  * [Try AWS re:Post](https://repost.aws/tags/TADSTjraA0Q4-a1dxk6eUYaw)
  * [Connect with an AWS IQ expert](https://iq.aws.amazon.com/get-started/?utm=docs&service=Amazon%20Simple%20Storage%20Service)


[Privacy](https://aws.amazon.com/privacy)[Site terms](https://aws.amazon.com/terms)[Cookie preferences](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html)
© 2025, Amazon Web Services, Inc. or its affiliates. All rights reserved.


## Source Information
- URL: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html
- Harvested: 2025-08-19T20:33:35.371186
- Profile: code_examples
- Agent: architect
