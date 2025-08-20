---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T22:09:00.395028'
profile: deep_research
source: https://min.io/docs/minio/linux/administration/identity-access-management/policy-based-access-control.html
topic: MinIO IAM and Access Control Policies
---

# MinIO IAM and Access Control Policies

[ ](https://min.io) Search documentation _Ctrl K_
Close search
[Docs Home](https://docs.min.io/)
  * [ Installation ](https://docs.min.io/enterprise/aistor-object-store/installation/)
    * [ Kubernetes ](https://docs.min.io/enterprise/aistor-object-store/installation/kubernetes/)
      * [ Install AIStor ](https://docs.min.io/enterprise/aistor-object-store/installation/kubernetes/install/)
        * [ Deploy AIStor on Kubernetes ](https://docs.min.io/enterprise/aistor-object-store/installation/kubernetes/install/deploy-aistor-on-kubernetes/)
        * [ Deploy AIStor on Red Hat OpenShift ](https://docs.min.io/enterprise/aistor-object-store/installation/kubernetes/install/deploy-aistor-on-openshift/)
      * [ Enable Network Encryption ](https://docs.min.io/enterprise/aistor-object-store/installation/kubernetes/network-encryption/)
      * [ Enable Server Side Encryption ](https://docs.min.io/enterprise/aistor-object-store/installation/kubernetes/server-side-encryption/)
        * [ Server Side Encryption with AIStor Key Manager ](https://docs.min.io/enterprise/aistor-object-store/installation/kubernetes/server-side-encryption/aistor-keymanager/)
        * [ Server Side Encryption with KES ](https://docs.min.io/enterprise/aistor-object-store/installation/kubernetes/server-side-encryption/minio-key-encryption-service/)
    * [ Linux ](https://docs.min.io/enterprise/aistor-object-store/installation/linux/)
      * [ Install AIStor ](https://docs.min.io/enterprise/aistor-object-store/installation/linux/install/)
        * [ Install AIStor on Ubuntu Server ](https://docs.min.io/enterprise/aistor-object-store/installation/linux/install/deploy-aistor-on-ubuntu-server/)
        * [ Install AIStor on Red Hat Enterprise Linux ](https://docs.min.io/enterprise/aistor-object-store/installation/linux/install/deploy-aistor-on-red-hat-linux/)
      * [ Enable Network Encryption ](https://docs.min.io/enterprise/aistor-object-store/installation/linux/network-encryption/)
      * [ Enable Server Side Encryption ](https://docs.min.io/enterprise/aistor-object-store/installation/linux/server-side-encryption/)
        * [ Server Side Encryption with AIStor Key Manager ](https://docs.min.io/enterprise/aistor-object-store/installation/linux/server-side-encryption/aistor-keymanager/)
        * [ Server Side Encryption with KES ](https://docs.min.io/enterprise/aistor-object-store/installation/linux/server-side-encryption/minio-key-encryption-service/)
    * [ macOS ](https://docs.min.io/enterprise/aistor-object-store/installation/macos/)
      * [ Deploy AIStor on macOS ](https://docs.min.io/enterprise/aistor-object-store/installation/macos/install/)
    * [ Container ](https://docs.min.io/enterprise/aistor-object-store/installation/container/)
      * [ Deploy AIStor as a Container ](https://docs.min.io/enterprise/aistor-object-store/installation/container/install/)
    * [ Windows ](https://docs.min.io/enterprise/aistor-object-store/installation/windows/)
      * [ Install AIStor ](https://docs.min.io/enterprise/aistor-object-store/installation/windows/install/)
        * [ Deploy AIStor Server on Windows ](https://docs.min.io/enterprise/aistor-object-store/installation/windows/install/deploy-aistor-on-windows/)
    * [ Checklists ](https://docs.min.io/enterprise/aistor-object-store/installation/checklists/)
      * [ Security checklist ](https://docs.min.io/enterprise/aistor-object-store/installation/checklists/security/)
      * [ Software checklist ](https://docs.min.io/enterprise/aistor-object-store/installation/checklists/software/)
  * [ Administration ](https://docs.min.io/enterprise/aistor-object-store/administration/)
    * [ Objects and Versioning ](https://docs.min.io/enterprise/aistor-object-store/administration/objects-and-versioning/)
      * [ Batch Job Expiration ](https://docs.min.io/enterprise/aistor-object-store/administration/objects-and-versioning/batch-expiration/)
      * [ Object Versioning ](https://docs.min.io/enterprise/aistor-object-store/administration/objects-and-versioning/versioning/)
        * [ Enable Object Versioning ](https://docs.min.io/enterprise/aistor-object-store/administration/objects-and-versioning/versioning/enable-versioning/)
        * [ Exclude Prefixes ](https://docs.min.io/enterprise/aistor-object-store/administration/objects-and-versioning/versioning/exclude-prefixes/)
        * [ Ignore Folder Objects ](https://docs.min.io/enterprise/aistor-object-store/administration/objects-and-versioning/versioning/ignore-folders/)
        * [ Suspend Versioning ](https://docs.min.io/enterprise/aistor-object-store/administration/objects-and-versioning/versioning/suspend-versioning/)
      * [ Data Compression ](https://docs.min.io/enterprise/aistor-object-store/administration/objects-and-versioning/data-compression/)
      * [ Object Deletion ](https://docs.min.io/enterprise/aistor-object-store/administration/objects-and-versioning/object-delete/)
    * [ Object Locking and Immutability ](https://docs.min.io/enterprise/aistor-object-store/administration/object-locking-and-immutability/)
    * [ Object Lifecycle Management ](https://docs.min.io/enterprise/aistor-object-store/administration/object-lifecycle-management/)
      * [ Automatic Object Expiration ](https://docs.min.io/enterprise/aistor-object-store/administration/object-lifecycle-management/create-lifecycle-management-expiration-rule/)
      * [ Object Tiering ](https://docs.min.io/enterprise/aistor-object-store/administration/object-lifecycle-management/object-tiering/)
        * [ Transition objects from AIStor to a remote AIStor deployment ](https://docs.min.io/enterprise/aistor-object-store/administration/object-lifecycle-management/object-tiering/transition-objects-to-minio/)
        * [ Transition objects from AIStor to Azure ](https://docs.min.io/enterprise/aistor-object-store/administration/object-lifecycle-management/object-tiering/transition-objects-to-azure/)
        * [ Transition objects from AIStor to Google Cloud Storage ](https://docs.min.io/enterprise/aistor-object-store/administration/object-lifecycle-management/object-tiering/transition-objects-to-gcs/)
        * [ Transition objects from AIStor to S3 ](https://docs.min.io/enterprise/aistor-object-store/administration/object-lifecycle-management/object-tiering/transition-objects-to-s3/)
    * [ Identity and Access Management ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/)
      * [ AIStor Identity Management ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/identity/)
        * [ Active Directory/LDAP Identity Management ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/identity/ldap-identity/)
        * [ OpenID Connect Identity Management ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/identity/oidc-identity/)
        * [ Identity Management Plugin ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/identity/plugin-identity/)
      * [ Access Control with Policy Management ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/)
        * [ Active Directory/LDAP Access Management ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/ldap-access/)
        * [ OpenID Connect Access Management ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/oidc-access/)
        * [ Authorization Plugin ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/plugin-access/)
    * [ Replication ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/)
      * [ Site Replication ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/site-replication/)
        * [ Initialize Site Replication ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/site-replication/create-initial-site-replication/)
        * [ Add New Peer Site ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/site-replication/site-replication-add-peer-site/)
        * [ Remove existing Peer Site ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/site-replication/remove-site-from-replication/)
        * [ Restore site replication service account ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/site-replication/restore-service-account/)
      * [ Bucket Replication ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/bucket-replication/)
        * [ Bucket Replication Requirements ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/bucket-replication/bucket-replication-requirements/)
        * [ Enable Active-Passive Bucket Replication ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/bucket-replication/enable-server-side-one-way-bucket-replication/)
        * [ Enable Active-Active Bucket Replication ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/bucket-replication/enable-server-side-two-way-bucket-replication/)
        * [ Resynchronize Remote after Data Loss ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/bucket-replication/server-side-replication-resynchronize-remote/)
      * [ Batch Replication ](https://docs.min.io/enterprise/aistor-object-store/administration/replication/batch-replication/)
    * [ Bucket Notifications ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/)
      * [ Publish Events to AMQP (RabbitMQ) ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/publish-events-to-amqp/)
      * [ Publish Events to Elasticsearch ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/publish-events-to-elasticsearch/)
      * [ Publish Events to Kafka ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/publish-events-to-kafka/)
      * [ Publish Events to MQTT ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/publish-events-to-mqtt/)
      * [ Publish Events to MySQL ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/publish-events-to-mysql/)
      * [ Publish Events to NATS ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/publish-events-to-nats/)
      * [ Publish Events to NSQ ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/publish-events-to-nsq/)
      * [ Publish Events to PostgreSQL ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/publish-events-to-postgresql/)
      * [ Publish Events to Redis ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/publish-events-to-redis/)
      * [ Publish Events to Webhook ](https://docs.min.io/enterprise/aistor-object-store/administration/bucket-notifications/publish-events-to-webhook/)
    * [ Batch Framework ](https://docs.min.io/enterprise/aistor-object-store/administration/batch-framework/)
    * [ AIStor Console ](https://docs.min.io/enterprise/aistor-object-store/administration/console/)
      * [ Managing a Deployment ](https://docs.min.io/enterprise/aistor-object-store/administration/console/managing-deployment/)
      * [ Managing Objects ](https://docs.min.io/enterprise/aistor-object-store/administration/console/managing-objects/)
      * [ Managing Security and Access ](https://docs.min.io/enterprise/aistor-object-store/administration/console/security-and-access/)
  * Operations
    * [ Core Concepts ](https://docs.min.io/enterprise/aistor-object-store/operations/core-concepts/)
      * [ Erasure Coding ](https://docs.min.io/enterprise/aistor-object-store/operations/core-concepts/erasure-coding/)
      * [ Healing ](https://docs.min.io/enterprise/aistor-object-store/operations/core-concepts/healing/)
    * [ Metrics and Logging ](https://docs.min.io/enterprise/aistor-object-store/operations/monitoring/)
      * [ Metrics ](https://docs.min.io/enterprise/aistor-object-store/operations/monitoring/metrics-and-alerts/)
        * [ Metrics version 2 ](https://docs.min.io/enterprise/aistor-object-store/operations/monitoring/metrics-and-alerts/metrics-v2/)
      * [ Server and Audit Logs ](https://docs.min.io/enterprise/aistor-object-store/operations/monitoring/minio-logging/)
      * [ Healthcheck Probes ](https://docs.min.io/enterprise/aistor-object-store/operations/monitoring/healthcheck-probe/)
    * [ Scaling ](https://docs.min.io/enterprise/aistor-object-store/operations/scaling/)
      * [ Expand Available Storage ](https://docs.min.io/enterprise/aistor-object-store/operations/scaling/expansion/)
        * [ Expand a Baremetal Deployment ](https://docs.min.io/enterprise/aistor-object-store/operations/scaling/expansion/baremetal/)
        * [ Expand an AIStor Deployment on Kubernetes ](https://docs.min.io/enterprise/aistor-object-store/operations/scaling/expansion/expand-aistor-kubernetes/)
        * [ Expand an AIStor Deployment on Linux ](https://docs.min.io/enterprise/aistor-object-store/operations/scaling/expansion/expand-aistor-linux/)
      * [ Decommision Aged Hardware ](https://docs.min.io/enterprise/aistor-object-store/operations/scaling/decommission/)
        * [ Decommission AIStor Server Pool on Linux ](https://docs.min.io/enterprise/aistor-object-store/operations/scaling/decommission/decommission-aistor-linux/)
    * [ Data Recovery ](https://docs.min.io/enterprise/aistor-object-store/operations/failure-and-recovery/)
      * [ Recover after drive failure ](https://docs.min.io/enterprise/aistor-object-store/operations/failure-and-recovery/recover-after-drive-failure/)
      * [ Recover after node failure ](https://docs.min.io/enterprise/aistor-object-store/operations/failure-and-recovery/recover-after-node-failure/)
      * [ Recover after site failure ](https://docs.min.io/enterprise/aistor-object-store/operations/failure-and-recovery/recover-after-site-failure/)
    * [ Licenses ](https://docs.min.io/enterprise/aistor-object-store/operations/licenses/)
  * [ Upgrade ](https://docs.min.io/enterprise/aistor-object-store/upgrade-aistor-server/)
    * [ Upgrade AIStor on Linux ](https://docs.min.io/enterprise/aistor-object-store/upgrade-aistor-server/upgrade-aistor-linux/)
    * [ Upgrade AIStor on Kubernetes (Helm) ](https://docs.min.io/enterprise/aistor-object-store/upgrade-aistor-server/upgrade-aistor-kubernetes-helm/)
  * Reference
    * [ AIStor Client ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/)
      * [ AIStor Client Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/aistor-client-settings/)
      * [ Admin Commands ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/)
        * [ mc admin accesskey ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-accesskey/)
          * [ mc admin accesskey create ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-accesskey/mc-admin-accesskey-create/)
          * [ mc admin accesskey disable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-accesskey/mc-admin-accesskey-disable/)
          * [ mc admin accesskey edit ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-accesskey/mc-admin-accesskey-edit/)
          * [ mc admin accesskey enable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-accesskey/mc-admin-accesskey-enable/)
          * [ mc admin accesskey info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-accesskey/mc-admin-accesskey-info/)
          * [ mc admin accesskey list ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-accesskey/mc-admin-accesskey-list/)
          * [ mc admin accesskey rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-accesskey/mc-admin-accesskey-remove/)
          * [ mc admin accesskey sts-revoke ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-accesskey/mc-admin-accesskey-sts-revoke/)
        * [ mc admin cluster bucket ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-cluster-bucket/)
          * [ mc admin cluster bucket export ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-cluster-bucket/mc-admin-cluster-bucket-export/)
          * [ mc admin cluster bucket import ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-cluster-bucket/mc-admin-cluster-bucket-import/)
        * [ mc admin cluster iam ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-cluster-iam/)
          * [ mc admin cluster iam export ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-cluster-iam/mc-admin-cluster-iam-export/)
          * [ mc admin cluster iam import ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-cluster-iam/mc-admin-cluster-iam-import/)
        * [ mc admin config ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-config/)
        * [ mc admin decommission ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-decommission/)
        * [ mc admin group ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-group/)
        * [ mc admin heal ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-heal/)
        * [ mc admin info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-info/)
        * [ mc admin kms key ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-kms-key/)
        * [ mc admin logs ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-logs/)
        * [ mc admin object info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-object-info/)
        * [ mc admin policy ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-policy/)
          * [ mc admin policy attach ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-policy/mc-admin-policy-attach/)
          * [ mc admin policy create ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-policy/mc-admin-policy-create/)
          * [ mc admin policy detach ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-policy/mc-admin-policy-detach/)
          * [ mc admin policy entities ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-policy/mc-admin-policy-entities/)
          * [ mc admin policy info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-policy/mc-admin-policy-info/)
          * [ mc admin policy ls ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-policy/mc-admin-policy-list/)
          * [ mc admin policy rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-policy/mc-admin-policy-remove/)
        * [ mc admin prometheus ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-prometheus/)
          * [ mc admin prometheus generate ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-prometheus/mc-admin-prometheus-generate/)
          * [ mc admin prometheus metrics ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-prometheus/mc-admin-prometheus-metrics/)
        * [ mc admin rebalance ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-rebalance/)
        * [ mc admin replicate ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-replicate/)
        * [ mc admin scanner ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-scanner/)
          * [ mc admin scanner status ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-scanner/mc-admin-scanner-status/)
          * [ mc admin scanner trace ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-scanner/mc-admin-scanner-trace/)
        * [ mc admin service ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-service/)
        * [ mc admin trace ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-trace/)
        * [ mc admin update ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-update/)
        * [ mc admin user ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-user/)
          * [ mc admin user add ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-user/mc-admin-user-add/)
          * [ mc admin user disable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-user/mc-admin-user-disable/)
          * [ mc admin user enable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-user/mc-admin-user-enable/)
          * [ mc admin user info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-user/mc-admin-user-info/)
          * [ mc admin user ls ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-user/mc-admin-user-list/)
          * [ mc admin user rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-user/mc-admin-user-remove/)
          * [ mc admin user sts info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-user/mc-admin-user-sts-info/)
      * [ mc alias ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-alias/)
        * [ mc alias export ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-alias/mc-alias-export/)
        * [ mc alias import ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-alias/mc-alias-import/)
        * [ mc alias list ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-alias/mc-alias-list/)
        * [ mc alias remove ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-alias/mc-alias-remove/)
        * [ mc alias set ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-alias/mc-alias-set/)
      * [ mc anonymous ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-anonymous/)
        * [ mc anonymous get ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-anonymous/mc-anonymous-get/)
        * [ mc anonymous get-json ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-anonymous/mc-anonymous-get-json/)
        * [ mc anonymous links ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-anonymous/mc-anonymous-links/)
        * [ mc anonymous list ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-anonymous/mc-anonymous-list/)
        * [ mc anonymous set ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-anonymous/mc-anonymous-set/)
        * [ mc anonymous set-json ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-anonymous/mc-anonymous-set-json/)
      * [ mc batch ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-batch/)
        * [ mc batch cancel ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-batch/mc-batch-cancel/)
        * [ mc batch describe ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-batch/mc-batch-describe/)
        * [ mc batch generate ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-batch/mc-batch-generate/)
        * [ mc batch list ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-batch/mc-batch-list/)
        * [ mc batch start ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-batch/mc-batch-start/)
        * [ mc batch status ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-batch/mc-batch-status/)
      * [ mc cat ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-cat/)
      * [ mc cp ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-cp/)
      * [ mc diff ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-diff/)
      * [ mc du ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-du/)
      * [ mc encrypt ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-encrypt/)
        * [ mc encrypt clear ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-encrypt/mc-encrypt-clear/)
        * [ mc encrypt info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-encrypt/mc-encrypt-info/)
        * [ mc encrypt set ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-encrypt/mc-encrypt-set/)
      * [ mc event ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-event/)
        * [ mc event add ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-event/mc-event-add/)
        * [ mc event list ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-event/mc-event-list/)
        * [ mc event remove ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-event/mc-event-remove/)
      * [ mc find ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-find/)
      * [ mc get ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-get/)
      * [ mc head ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-head/)
      * [ mc idp ldap ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap/)
        * [ mc idp ldap add ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap/mc-idp-ldap-add/)
        * [ mc idp ldap disable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap/mc-idp-ldap-disable/)
        * [ mc idp ldap enable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap/mc-idp-ldap-enable/)
        * [ mc idp ldap info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap/mc-idp-ldap-info/)
        * [ mc idp ldap ls ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap/mc-idp-ldap-ls/)
        * [ mc idp ldap rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap/mc-idp-ldap-rm/)
        * [ mc idp ldap update ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap/mc-idp-ldap-update/)
      * [ mc idp ldap accesskey ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-accesskey/)
        * [ mc idp ldap accesskey create ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-accesskey/mc-idp-ldap-accesskey-create/)
        * [ mc idp ldap accesskey create-with-login ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-accesskey/mc-idp-ldap-accesskey-create-with-login/)
        * [ mc idp ldap accesskey disable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-accesskey/mc-idp-ldap-accesskey-disable/)
        * [ mc idp ldap accesskey edit ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-accesskey/mc-idp-ldap-accesskey-edit/)
        * [ mc idp ldap accesskey enable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-accesskey/mc-idp-ldap-accesskey-enable/)
        * [ mc idp ldap accesskey info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-accesskey/mc-idp-ldap-accesskey-info/)
        * [ mc idp ldap accesskey ls ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-accesskey/mc-idp-ldap-accesskey-ls/)
        * [ mc idp ldap accesskey rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-accesskey/mc-idp-ldap-accesskey-rm/)
        * [ mc idp ldap accesskey sts-revoke ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-accesskey/mc-idp-ldap-accesskey-sts-revoke/)
      * [ mc idp ldap policy ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-policy/)
        * [ mc idp ldap policy attach ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-policy/mc-idp-ldap-policy-attach/)
        * [ mc idp ldap policy detach ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-policy/mc-idp-ldap-policy-detach/)
        * [ mc idp ldap policy entities ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-ldap-policy/mc-idp-ldap-policy-entities/)
      * [ mc idp openid ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid/)
        * [ mc idp openid add ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid/mc-idp-openid-add/)
        * [ mc idp openid disable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid/mc-idp-openid-disable/)
        * [ mc idp openid enable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid/mc-idp-openid-enable/)
        * [ mc idp openid info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid/mc-idp-openid-info/)
        * [ mc idp openid ls ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid/mc-idp-openid-ls/)
        * [ mc idp openid rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid/mc-idp-openid-rm/)
        * [ mc idp openid update ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid/mc-idp-openid-update/)
      * [ mc idp openid accesskey ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid-accesskey/)
        * [ mc idp openid accesskey disable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid-accesskey/mc-idp-openid-accesskey-disable/)
        * [ mc idp openid accesskey edit ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid-accesskey/mc-idp-openid-accesskey-edit/)
        * [ mc idp openid accesskey enable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid-accesskey/mc-idp-openid-accesskey-enable/)
        * [ mc idp openid accesskey info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid-accesskey/mc-idp-openid-accesskey-info/)
        * [ mc idp openid accesskey ls ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid-accesskey/mc-idp-openid-accesskey-ls/)
        * [ mc idp openid accesskey rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-idp-openid-accesskey/mc-idp-openid-accesskey-rm/)
      * [ mc ilm rule ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-rule/)
        * [ mc ilm rule add ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-rule/mc-ilm-rule-add/)
        * [ mc ilm rule edit ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-rule/mc-ilm-rule-edit/)
        * [ mc ilm rule export ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-rule/mc-ilm-rule-export/)
        * [ mc ilm rule import ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-rule/mc-ilm-rule-import/)
        * [ mc ilm rule ls ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-rule/mc-ilm-rule-ls/)
        * [ mc ilm rule rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-rule/mc-ilm-rule-rm/)
      * [ mc ilm tier ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-tier/)
        * [ mc ilm restore ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-tier/mc-ilm-restore/)
        * [ mc ilm tier add ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-tier/mc-ilm-tier-add/)
        * [ mc ilm tier check ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-tier/mc-ilm-tier-check/)
        * [ mc ilm tier info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-tier/mc-ilm-tier-info/)
        * [ mc ilm tier ls ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-tier/mc-ilm-tier-ls/)
        * [ mc ilm tier rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-tier/mc-ilm-tier-rm/)
        * [ mc ilm tier update ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ilm-tier/mc-ilm-tier-update/)
      * [ mc legalhold ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-legalhold/)
        * [ mc legalhold clear ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-legalhold/mc-legalhold-clear/)
        * [ mc legalhold info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-legalhold/mc-legalhold-info/)
        * [ mc legalhold set ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-legalhold/mc-legalhold-set/)
      * [ mc license ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-license/)
        * [ mc license info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-license/mc-license-info/)
        * [ mc license register ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-license/mc-license-register/)
      * [ mc ls ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ls/)
      * [ mc mb ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-mb/)
      * [ mc mirror ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-mirror/)
      * [ mc mv ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-mv/)
      * [ mc od ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-od/)
      * [ mc ping ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ping/)
      * [ mc pipe ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-pipe/)
      * [ mc put ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-put/)
      * [ mc rb ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-rb/)
      * [ mc ready ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-ready/)
      * [ mc replicate ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-replicate/)
        * [ mc replicate add ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-replicate/mc-replicate-add/)
        * [ mc replicate backlog ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-replicate/mc-replicate-backlog/)
        * [ mc replicate export ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-replicate/mc-replicate-export/)
        * [ mc replicate import ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-replicate/mc-replicate-import/)
        * [ mc replicate ls ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-replicate/mc-replicate-ls/)
        * [ mc replicate resync ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-replicate/mc-replicate-resync/)
        * [ mc replicate rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-replicate/mc-replicate-rm/)
        * [ mc replicate status ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-replicate/mc-replicate-status/)
        * [ mc replicate update ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-replicate/mc-replicate-update/)
      * [ mc retention ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-retention/)
        * [ mc retention clear ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-retention/mc-retention-clear/)
        * [ mc retention info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-retention/mc-retention-info/)
        * [ mc retention set ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-retention/mc-retention-set/)
      * [ mc rm ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-rm/)
      * [ mc share ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-share/)
        * [ mc share download ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-share/mc-share-download/)
        * [ mc share list ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-share/mc-share-list/)
        * [ mc share upload ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-share/mc-share-upload/)
      * [ mc sql ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-sql/)
      * [ mc stat ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-stat/)
      * [ mc support ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/)
        * [ mc support callhome ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-callhome/)
        * [ mc support diag ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-diag/)
        * [ mc support inspect ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-inspect/)
        * [ mc support perf ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-perf/)
        * [ mc support profile ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-profile/)
        * [ mc support proxy ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-proxy/)
        * [ mc support telemetry ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-telemetry/)
          * [ mc support telemetry proxy ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-telemetry/mc-support-telemetry-proxy/)
          * [ mc support telemetry record ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-telemetry/mc-support-telemetry-record/)
          * [ mc support telemetry replay ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-telemetry/mc-support-telemetry-replay/)
        * [ mc support upload ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support/mc-support-upload/)
      * [ mc support top ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support-top/)
        * [ mc support top api ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support-top/mc-support-top-api/)
        * [ mc support top disk ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support-top/mc-support-top-disk/)
        * [ mc support top locks ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support-top/mc-support-top-locks/)
        * [ mc support top net ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support-top/mc-support-top-net/)
        * [ mc support top rpc ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-support-top/mc-support-top-rpc/)
      * [ mc tag ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-tag/)
        * [ mc tag list ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-tag/mc-tag-list/)
        * [ mc tag remove ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-tag/mc-tag-remove/)
        * [ mc tag set ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-tag/mc-tag-set/)
      * [ mc tree ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-tree/)
      * [ mc undo ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-undo/)
      * [ mc update ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-update/)
      * [ mc version ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-version/)
        * [ mc version enable ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-version/mc-version-enable/)
        * [ mc version info ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-version/mc-version-info/)
        * [ mc version suspend ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-version/mc-version-suspend/)
      * [ mc watch ](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-watch/)
    * [ AIStor Server ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/)
      * [ Settings and Configurations ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/)
        * [ Core Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/core/)
        * [ Batch Job Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/batch/)
        * [ Console Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/console/)
        * [ Erasure Code Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/storage-class/)
        * [ ILM Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/ilm/)
        * [ Metrics and Logging Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/metrics-and-logging/)
          * [ Server Logs ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/metrics-and-logging/server-logs/)
          * [ Webhook Audit Logs ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/metrics-and-logging/webhook-audit-logs/)
          * [ Kafka Audit Logs ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/metrics-and-logging/kafka-audit-logs/)
          * [ Telemetry Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/metrics-and-logging/telemetry/)
        * [ Notifications Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/)
          * [ AMQP Notification Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/amqp/)
          * [ Elasticsearch Notification Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/elasticsearch/)
          * [ Kafka Notification Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/kafka/)
          * [ MQTT Notification Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/mqtt/)
          * [ MySQL Notification Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/mysql/)
          * [ NATS Notification Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/nats/)
          * [ NSQ Notification Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/nsq/)
          * [ PostgreSQL Notification Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/postgresql/)
          * [ Redis Notification Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/redis/)
          * [ Webhook Notification Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/notifications/webhook-service/)
        * [ Object Lambda Function Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/object-lambda/)
        * [ Root Access Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/root-credentials/)
        * [ Server Side Encryption Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/server-side-encryption/)
        * [ IAM Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/iam/)
          * [ LDAP Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/iam/ldap/)
          * [ OpenID Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/iam/openid/)
          * [ Access Management Plugin Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/iam/minio-access-plugin/)
          * [ Identity Management Plugin Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/iam/minio-identity-plugin/)
        * [ Deprecated Settings ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/settings/deprecated/)
      * [ Thresholds and Limits ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/thresholds/)
      * [ Scanner ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/scanner/)
      * [ Hardware Tests ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/recommended-hardware-tests/)
      * [ S3 Express Mode ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/s3-express/)
      * [ Hardware and System Requirements ](https://docs.min.io/enterprise/aistor-object-store/reference/aistor-server/hardware-and-system-requirements/)
    * [ AIStor Operator ](https://docs.min.io/enterprise/aistor-object-store/reference/kubernetes/)
      * [ AIStor Server Helm Chart ](https://docs.min.io/enterprise/aistor-object-store/reference/kubernetes/object-store-helm-chart/)
      * [ AIStor Custom Resource Definition (Stable) ](https://docs.min.io/enterprise/aistor-object-store/reference/kubernetes/aistor-crd-v1/)
      * [ AIStor Custom Resource Definition (Alpha) ](https://docs.min.io/enterprise/aistor-object-store/reference/kubernetes/aistor-crd-alpha1/)
    * [ Release Notes ](https://docs.min.io/enterprise/aistor-object-store/reference/release-notes/)
      * [ AIStor Server Release Notes ](https://docs.min.io/enterprise/aistor-object-store/reference/release-notes/aistor-server/)
      * [ AIStor Operator Release Notes ](https://docs.min.io/enterprise/aistor-object-store/reference/release-notes/kubernetes/)
        * [ Archived Releases ](https://docs.min.io/enterprise/aistor-object-store/reference/release-notes/kubernetes/archived-releases/)
  * Developers
    * [ Security Token Service ](https://docs.min.io/enterprise/aistor-object-store/developers/security-token-service/)
      * [ AssumeRoleWithCustomToken ](https://docs.min.io/enterprise/aistor-object-store/developers/security-token-service/assumerolewithcustomtoken/)
      * [ AssumeRoleWithLDAPIdentity ](https://docs.min.io/enterprise/aistor-object-store/developers/security-token-service/assumerolewithldapidentity/)
      * [ AssumeRoleWithWebIdentity ](https://docs.min.io/enterprise/aistor-object-store/developers/security-token-service/assumerolewithwebidentity/)
    * [ Software Development Kits (SDKs) ](https://docs.min.io/enterprise/aistor-object-store/developers/minio-drivers/)
    * [ File Transfer Protocol ](https://docs.min.io/enterprise/aistor-object-store/developers/file-transfer-protocol/)
    * [ Transforms with Object Lambda ](https://docs.min.io/enterprise/aistor-object-store/developers/transforms-with-object-lambda/)
    * [ S3 API Compatibility ](https://docs.min.io/enterprise/aistor-object-store/developers/s3-api-compatibility/)


Close Sidebar
  1. [ ](https://docs.min.io/)
  2. [Administration](https://docs.min.io/enterprise/aistor-object-store/administration/)
  3. [Identity and Access Management](https://docs.min.io/enterprise/aistor-object-store/administration/iam/)
  4. [Access Control with Policy Management](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/)


# Access Control with Policy Management
AIStor uses Policy-Based Access Control (PBAC) to define the authorized actions and resources to which an authenticated user has access. Each policy describes one or more actions and conditions that outline the permissions of a user or group of users.
AIStor PBAC is built for compatibility with AWS IAM policy syntax, structure, and behavior. This documentation makes a best-effort to cover IAM-specific behavior and functionality. Refer also to the AWS IAM documentation for more information.
The `mc admin policy` command supports creation and management of policies on the AIStor deployment. See the command reference for examples of usage.
##  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#built-in-policies "Anchor to: Built-In Policies") Built-In Policies 
AIStor provides the following built-in policies for assigning to users or groups:
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#consoleadmin "Anchor to: consoleAdmin") consoleAdmin 
Grants complete access to all S3 and administrative API operations against all resources on the AIStor deployment. Equivalent to the following set of actions:
  * `s3:*`
  * `admin:*`


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#readonly "Anchor to: readonly") readonly 
Grants read-only permissions on any object on the AIStor deployment. The GET action _must_ apply to a specific object without requiring any listing. Equivalent to the following set of actions:
  * `s3:GetBucketLocation`
  * `s3:GetObject`


For example, this policy supports GET operations on objects at a specific path (e.g. `GET play/mybucket/object.file`), such as:
  * `mc cp`
  * `mc stat`
  * `mc head`
  * `mc cat`


The exclusion of listing permissions is intentional, as typical use cases do not intend for a “read-only” role to have complete discoverability (listing all buckets and objects) on the object storage resource.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#readwrite "Anchor to: readwrite") readwrite 
Grants read and write permissions for all buckets and objects on the AIStor server. Equivalent to `s3:*`.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#diagnostics "Anchor to: diagnostics") diagnostics 
Grants permission to perform diagnostic actions on the AIStor deployment. Includes the following actions:
  * `admin:ServerTrace`
  * `admin:Profiling`
  * `admin:ConsoleLog`
  * `admin:ServerInfo`
  * `admin:TopLocksInfo`
  * `admin:OBDInfo`
  * `admin:BandwidthMonitor`
  * `admin:Prometheus`


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#writeonly "Anchor to: writeonly") writeonly 
Grants write-only permissions to any namespace in the AIStor deployment (bucket and path to object). The PUT action _must_ apply to a specific object location without requiring any listing. Equivalent to the `s3:PutObject` action.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#assign-default-policies "Anchor to: Assign default policies") Assign default policies 
Run [`mc admin policy attach`](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-policy/mc-admin-policy-attach/) to assign a policy to a user or group.
The following table illustrates how you might group users and assign built-in policies on specified resources.
User Group | Policy | Operations  
---|---|---  
Operations | `readwrite` on finance bucket`readonly` on audit bucket | `PUT` and `GET` on finance bucket.`GET` on audit bucket  
Auditing | `readonly` on audit bucket | `GET` on audit bucket  
Admin | admin:* | All `mc admin` commands.  
Each user can access only the resources and operations which are _explicitly_ assigned to the built-in role. AIStor denies access to any other resource or action.
Deny Overrides Allow
AIStor follows AWS IAM policy evaluation rules, where a `Deny` rule overrides an `Allow` rule on the same action or resource. For example, if a policy that includes an `Allow` rule for a specified resource is assigned to a user, but a policy with a `Deny` rule for the same resource is assigned to a group the user is a member of, only the `Deny` rule is applied.
For more information, see the AWS documentation on [Policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html).
##  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#policy-document-structure "Anchor to: Policy Document Structure") Policy Document Structure 
AIStor policy documents use the same schema as AWS IAM Policy documents.
The following sample document provides a template for creating custom policies for use with an AIStor deployment. For more complete documentation on IAM policy elements, see the IAM JSON Policy Elements Reference.
The maximum size for any single policy document is 20KiB. There is no limit to the number of policy documents that can be attached to a user or group.
```
{ 
   "Version" : "2012-10-17",
   "Statement" : [ 
      { "Effect" : "Allow",
       "Action" : [ "s3:<ActionName>", ... ],
       "Resource" : "arn:aws:s3:::*",
       "Condition" : { ... } 
      },
      { "Effect" : "Deny",
       "Action" : [ "s3:<ActionName>", ... ],
       "Resource" : "arn:aws:s3:::*",
       "Condition" : { ... }
      }
   ]
}

```

  * For the `Statement.Action` array, specify one or more supported S3 API operations.
  * For the `Statement.Resource` key, specify the bucket or bucket prefix to which to restrict the policy. You can use `*` and `?` wildcard characters as per the S3 Resource Spec.
The `*` wildcard may result in unintended application of a policy to multiple buckets or prefixes based on the pattern match. For example, `arn:aws:s3:::data*` would match the buckets `data`, `data_private`, and `data_internal`. Specifying only `*` as the resource key applies the policy to all buckets and prefixes on the deployment.
  * For the `Statement.Condition` key, you can specify one or more supported Conditions.


##  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#supported-s3-policy-actions "Anchor to: Supported S3 policy actions") Supported S3 policy actions 
AIStor policy documents support a subset of IAM S3 Action keys. AIStor also supports condition keys for some actions beyond the common set of supported keys.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#common-s3-operations "Anchor to: Common S3 operations") Common S3 operations 
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3 "Anchor to: s3:*") s3:* 
Selector for all AIStor S3 operations. Applying this action to a given resource allows the user to perform any S3 operation against that resource.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3createbucket "Anchor to: s3:CreateBucket") s3:CreateBucket 
Controls access to the [CreateBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html) S3 API
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3deletebucket "Anchor to: s3:DeleteBucket") s3:DeleteBucket 
Controls access to the [DeleteBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3forcedeletebucket "Anchor to: s3:ForceDeleteBucket") s3:ForceDeleteBucket 
Controls access to the DeleteBucket S3 API operation for operations with the `x-AIStor-force-delete` flag. Required for removing non-empty buckets.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getbucketlocation "Anchor to: s3:GetBucketLocation") s3:GetBucketLocation 
Controls access to the [GetBucketLocation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLocation.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3listallmybuckets "Anchor to: s3:ListAllMyBuckets") s3:ListAllMyBuckets 
Controls access to the [ListBuckets](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html) s3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3deleteobject "Anchor to: s3:DeleteObject") s3:DeleteObject 
Controls access to the [DeleteObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getobject "Anchor to: s3:GetObject") s3:GetObject 
Controls access to the [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html) S3 API operation.
Supports the following additional condition keys:
  * s3:x-amz-server-side-encryption
  * s3:x-amz-server-side-encryption-customer-algorithm
  * s3:ExistingObjectTag/
  * s3:versionid


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getobjectattributes "Anchor to: s3:GetObjectAttributes") s3:GetObjectAttributes 
Controls access to the [GetObjectAttributes](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getobjectversionattributes "Anchor to: s3:GetObjectVersionAttributes") s3:GetObjectVersionAttributes 
Controls access to the [`GetObjectAttributes`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html) S3 API operations on versioned objects.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3restoreobject "Anchor to: s3:RestoreObject") s3:RestoreObject 
Controls access to the [RestoreObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3listbucket "Anchor to: s3:ListBucket") s3:ListBucket 
Controls access to the [ListObjectsV2](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html) S3 API operation.
Supports the following additional condition keys:
  * s3:prefix
  * s3:delimiter
  * s3:max-keys


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putobject "Anchor to: s3:PutObject") s3:PutObject 
Controls access to the [PutObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html) S3 API operation.
Supports the following additional condition keys:
  * s3:x-amz-copy-source
  * s3:x-amz-server-side-encryption
  * s3:x-amz-server-side-encryption-customer-algorithm
  * s3:x-amz-metadata-directive
  * s3:x-amz-storage-class
  * s3:versionid
  * s3:object-lock-retain-until-date
  * s3:object-lock-mode
  * s3:object-lock-legal-hold
  * s3:RequestObjectTagKeys
  * s3:RequestObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putobjecttagging "Anchor to: s3:PutObjectTagging") s3:PutObjectTagging 
Controls access to the [PutObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectTagging.html) S3 API operation.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/
  * s3:RequestObjectTagKeys
  * s3:RequestObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getobjecttagging "Anchor to: s3:GetObjectTagging") s3:GetObjectTagging 
Controls access to the [GetObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html) S3 API operation.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3deleteobjecttagging "Anchor to: s3:DeleteObjectTagging") s3:DeleteObjectTagging 
Controls access to the [DeleteObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectTagging.html) S3 API operation.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/


###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#bucket-configuration "Anchor to: Bucket Configuration") Bucket Configuration 
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getbucketpolicy "Anchor to: s3:GetBucketPolicy") s3:GetBucketPolicy 
Controls access to the [GetBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicy.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putbucketpolicy "Anchor to: s3:PutBucketPolicy") s3:PutBucketPolicy 
Controls access to the [PutBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketPolicy.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3deletebucketpolicy "Anchor to: s3:DeleteBucketPolicy") s3:DeleteBucketPolicy 
Controls access to the [DeleteBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketPolicy.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getbuckettagging "Anchor to: s3:GetBucketTagging") s3:GetBucketTagging 
Controls access to the [GetBucketTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putbuckettagging "Anchor to: s3:PutBucketTagging") s3:PutBucketTagging 
Controls access to the [PutBucketTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html) S3 API operation.
Supports the following additional condition keys:
  * s3:RequestObjectTagKeys
  * s3:RequestObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getbucketpolicystatus "Anchor to: s3:GetBucketPolicyStatus") s3:GetBucketPolicyStatus 
Controls access to the [GetBucketPolicyStatus](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicyStatus.html) S3 API operation.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#multipart-upload "Anchor to: Multipart Upload") Multipart Upload 
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3abortmultipartupload "Anchor to: s3:AbortMultipartUpload") s3:AbortMultipartUpload 
Controls access to the [AbortMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3listmultipartuploadparts "Anchor to: s3:ListMultipartUploadParts") s3:ListMultipartUploadParts 
Controls access to the [ListParts](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3listbucketmultipartuploads "Anchor to: s3:ListBucketMultipartUploads") s3:ListBucketMultipartUploads 
Controls access to the [ListMultipartUploads](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html) S3 API operation.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#versioning-and-retention "Anchor to: Versioning and Retention") Versioning and Retention 
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putbucketversioning "Anchor to: s3:PutBucketVersioning") s3:PutBucketVersioning 
Controls access to the [PutBucketVersioning](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketVersioning.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getbucketversioning "Anchor to: s3:GetBucketVersioning") s3:GetBucketVersioning 
Controls access to the [GetBucketVersioning](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3deleteobjectversion "Anchor to: s3:DeleteObjectVersion") s3:DeleteObjectVersion 
Controls access to the [DeleteObjectVersion](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectVersion.html) S3 API operation.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3listbucketversions "Anchor to: s3:ListBucketVersions") s3:ListBucketVersions 
Controls access to the [ListBucketVersions](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketVersions.html) S3 API operation.
Supports the following additional condition keys:
  * s3:prefix
  * s3:delimiter
  * s3:max-keys


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putobjectversiontagging "Anchor to: s3:PutObjectVersionTagging") s3:PutObjectVersionTagging 
Controls access to the [PutObjectVersionTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectVersionTagging.html) S3 API operation.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/
  * s3:RequestObjectTagKeys
  * s3:RequestObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getobjectversiontagging "Anchor to: s3:GetObjectVersionTagging") s3:GetObjectVersionTagging 
Controls access to the [GetObjectVersionTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectVersionTagging.html) S3 API operation.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3deleteobjectversiontagging "Anchor to: s3:DeleteObjectVersionTagging") s3:DeleteObjectVersionTagging 
Controls access to the [DeleteObjectVersionTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectVersionTagging.html) S3 API operation.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getobjectversion "Anchor to: s3:GetObjectVersion") s3:GetObjectVersion 
Controls access to the [GetObjectVersion](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectVersion.html) S3 API operation.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3bypassgovernanceretention "Anchor to: s3:BypassGovernanceRetention") s3:BypassGovernanceRetention 
Controls access to the following S3 API operations on objects locked under [`GOVERNANCE`](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-retention/mc-retention-set/#MODE) retention mode:
  * `s3:PutObjectRetention`
  * `s3:PutObject`
  * `s3:DeleteObject`


See the S3 documentation on [s3:BypassGovernanceRetention](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-managing.html#object-lock-managing-bypass) for more information.
Supports the following additional condition keys:
  * s3:versionid
  * s3:object-lock-remaining-retention-days
  * s3:object-lock-retain-until-date
  * s3:object-lock-mode
  * s3:object-lock-legal-hold
  * s3:RequestObjectTagKeys
  * s3:RequestObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putobjectretention "Anchor to: s3:PutObjectRetention") s3:PutObjectRetention 
Controls access to the [PutObjectRetention](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectRetention.html) S3 API operation.
Required for any `PutObject` operation that specifies [retention metadata](https://docs.min.io/enterprise/aistor-object-store/administration/object-locking-and-immutability/).
Supports the following additional condition keys:
  * s3:x-amz-server-side-encryption
  * s3:x-amz-server-side-encryption-customer-algorithm
  * s3:x-amz-object-lock-remaining-retention-days
  * s3:x-amz-object-lock-retain-until-date
  * s3:x-amz-object-lock-mode
  * s3:versionid


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getobjectretention "Anchor to: s3:GetObjectRetention") s3:GetObjectRetention 
Controls access to the [GetObjectRetention](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectRetention.html) S3 API operation.
Required for including [object locking metadata](https://docs.min.io/enterprise/aistor-object-store/administration/object-locking-and-immutability/) as part of the response to a `GetObject` or `HeadObject` operation.
Supports the following additional condition keys:
  * s3:x-amz-server-side-encryption
  * s3:x-amz-server-side-encryption-customer-algorithm
  * s3:versionid


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getobjectlegalhold "Anchor to: s3:GetObjectLegalHold") s3:GetObjectLegalHold 
Controls access to the [GetObjectLegalHold](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLegalHold.html) S3 API operation.
Required for including [legal hold metadata](https://docs.min.io/enterprise/aistor-object-store/administration/object-locking-and-immutability/) as part of the response to a `GetObject` or `HeadObject` operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putobjectlegalhold "Anchor to: s3:PutObjectLegalHold") s3:PutObjectLegalHold 
Controls access to the [PutObjectLegalHold](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectLegalHold.html) S3 API operation.
Required for any `PutObject` operation that specifies [legal hold metadata](https://docs.min.io/enterprise/aistor-object-store/administration/object-locking-and-immutability/).
Supports the following additional condition keys:
  * s3:x-amz-server-side-encryption
  * s3:x-amz-server-side-encryption-customer-algorithm
  * s3:object-lock-legal-hold
  * s3:versionid


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getbucketobjectlockconfiguration "Anchor to: s3:GetBucketObjectLockConfiguration") s3:GetBucketObjectLockConfiguration 
Controls access to the [GetObjectLockConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLockConfiguration.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putbucketobjectlockconfiguration "Anchor to: s3:PutBucketObjectLockConfiguration") s3:PutBucketObjectLockConfiguration 
Controls access to the [PutObjectLockConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectLockConfiguration.html) S3 API operation.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#bucket-notifications "Anchor to: Bucket Notifications") Bucket Notifications 
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getbucketnotification "Anchor to: s3:GetBucketNotification") s3:GetBucketNotification 
Controls access to the [GetBucketNotification](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketNotification.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putbucketnotification "Anchor to: s3:PutBucketNotification") s3:PutBucketNotification 
Controls access to the [PutBucketNotification](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketNotification.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3listennotification "Anchor to: s3:ListenNotification") s3:ListenNotification 
Controls API operations related to AIStor bucket notifications.
This action is **not** intended for use with other S3-compatible services.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3listenbucketnotification "Anchor to: s3:ListenBucketNotification") s3:ListenBucketNotification 
Controls API operations related to AIStor bucket notifications.
This action is **not** intended for use with other S3-compatible services.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#object-lifecycle-management "Anchor to: Object Lifecycle Management") Object Lifecycle Management 
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putlifecycleconfiguration "Anchor to: s3:PutLifecycleConfiguration") s3:PutLifecycleConfiguration 
Controls access to the [PutLifecycleConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getlifecycleconfiguration "Anchor to: s3:GetLifecycleConfiguration") s3:GetLifecycleConfiguration 
Controls access to the [GetLifecycleConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html) S3 API operation.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#object-encryption "Anchor to: Object Encryption") Object Encryption 
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putencryptionconfiguration "Anchor to: s3:PutEncryptionConfiguration") s3:PutEncryptionConfiguration 
Controls access to the [PutEncryptionConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketEncryption.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getencryptionconfiguration "Anchor to: s3:GetEncryptionConfiguration") s3:GetEncryptionConfiguration 
Controls access to the [GetEncryptionConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketEncryption.html) S3 API operation.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#bucket-replication "Anchor to: Bucket Replication") Bucket Replication 
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getreplicationconfiguration "Anchor to: s3:GetReplicationConfiguration") s3:GetReplicationConfiguration 
Controls access to the [GetBucketReplication](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3putreplicationconfiguration "Anchor to: s3:PutReplicationConfiguration") s3:PutReplicationConfiguration 
Controls access to the [PutBucketReplication](https://docs.aws.amazon.com/AmazonS3/latest/API/PutBucketReplication.html) S3 API operation.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3replicateobject "Anchor to: s3:ReplicateObject") s3:ReplicateObject 
Controls API operations related to AIStor server-side bucket replication.
Required for AIStor server-side replication.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3replicatedelete "Anchor to: s3:ReplicateDelete") s3:ReplicateDelete 
Controls API operations related to AIStor server-side bucket replication.
Required for synchronizing delete operations as part of AIStor server-side replication.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3replicatetags "Anchor to: s3:ReplicateTags") s3:ReplicateTags 
Controls API operations related to AIStor server-side bucket replication.
Required for AIStor server-side replication.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/


####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#s3getobjectversionforreplication "Anchor to: s3:GetObjectVersionForReplication") s3:GetObjectVersionForReplication 
Controls API operations related to AIStor server-side bucket replication.
Required for AIStor server-side replication.
Supports the following additional condition keys:
  * s3:versionid
  * s3:ExistingObjectTag/


##  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#supported-s3-policy-condition-keys "Anchor to: Supported S3 Policy Condition Keys") Supported S3 Policy Condition Keys 
AIStor policy documents support IAM conditional statements.
Each condition element consists of [operators](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html) and condition keys. AIStor supports a subset of IAM condition keys. For complete information on any listed condition key, see the [IAM Condition Element Documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html)
AIStor supports the following condition keys for all supported actions:
  * `aws:Referer`
  * `aws:SourceIp`
  * `aws:UserAgent`
  * `aws:SecureTransport`
  * `aws:CurrentTime`
  * `aws:EpochTime`
  * `aws:PrincipalType`
  * `aws:userid`
  * `aws:username`
  * `x-amz-content-sha256`
  * `s3:signatureAge`


Warning
The `aws:Referer`, `aws:SourceIp`, and `aws.UserAgent` keys may be easily spoofed and therefore pose a potential security risk. AIStor recommends using these condition keys only to _deny_ access as a secondary security measure.
**Never** use these three keys to grant access by themselves.
For additional keys supported by a specific S3 action, see the reference documentation for that action.
##  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#aistor-extended-condition-keys "Anchor to: AIStor Extended Condition Keys") AIStor Extended Condition Keys 
AIStor extends the S3 standard condition keys with the following extended key:
`sts:DurationSeconds`
Specify a time in seconds to limit the duration of _all_ Security Token Service credentials generated by `AssumeRoleWithWebIdentity`.
This value overrides the `DurationSeconds` field specified to the client.
For example:
```
{
   "Version": "2012-10-17",
   "Statement": [
   {
      "Effect": "Allow",
      "Action": [
         "sts:AssumeRoleWithWebIdentity"
      ],
      "Condition": {
         "NumericLessThanEquals": {
         "sts:DurationSeconds": "300"
         }
      }
   }
  ]
}

```

##  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#kms-policy-action-keys "Anchor to: KMS policy action keys") KMS policy action keys 
AIStor supports restricting key management service (KMS) actions by policy.
You can restrict KMS activities in a policy with any of the following KMS actions:
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#kmsstatus "Anchor to: kms:Status") kms:Status 
Check the status of KMS.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#kmsmetrics "Anchor to: kms:Metrics") kms:Metrics 
Obtain Prometheus-formatted metrics.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#kmsapi "Anchor to: kms:API") kms:API 
List supported API endpoints.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#kmsversion "Anchor to: kms:Version") kms:Version 
Retrieve the KMS version.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#kmscreatekey "Anchor to: kms:CreateKey") kms:CreateKey 
Create a new KMS key.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#kmslistkeys "Anchor to: kms:ListKeys") kms:ListKeys 
Retrieve a list of existing KMS keys.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#kmskeystatus "Anchor to: kms:KeyStatus") kms:KeyStatus 
Retrieve the status of a specified KMS key.
To select all of the available kms policy actions, use `kms:*`.
KMS actions can be restricted by resource or a resource prefix. The wildcard character `*` can be used to apply the KMS action policy to all resources that match the prefix.
For example, the following policy document allows a user to list keys, create new keys, and check the status of keys for any resource that begins with `keys-abc-` or `myuser-`.
```
{
   "Version": "2012-10-17",
   "Statement": [
      {
      "Effect": "Allow",
      "Action": [
         "kms:CreateKey",
         "kms:KeyStatus",
         "kms:ListKeys"
      ],
      "Resource": [
         "arn:AIStor:kms:::keys-abc-*",
         "arn:AIStor:kms:::myuser-*"
      ]
      }
   ]
}

```

##  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#mc-admin-policy-action-keys "Anchor to: mc admin Policy Action Keys") `mc admin` Policy Action Keys 
AIStor supports the following actions for use with defining policies for `mc admin` operations. These actions are _only_ valid for AIStor deployments and are _not_ intended for use with other S3-compatible services.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admin "Anchor to: admin:*") admin:* 
Selector for all admin action keys.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminheal "Anchor to: admin:Heal") admin:Heal 
Allows heal command
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminstorageinfo "Anchor to: admin:StorageInfo") admin:StorageInfo 
Allows listing server info
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admindatausageinfo "Anchor to: admin:DataUsageInfo") admin:DataUsageInfo 
Allows listing data usage info
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admintoplocksinfo "Anchor to: admin:TopLocksInfo") admin:TopLocksInfo 
Allows listing top locks
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminprofiling "Anchor to: admin:Profiling") admin:Profiling 
Allows profiling
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminservertrace "Anchor to: admin:ServerTrace") admin:ServerTrace 
Allows listing server trace
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminconsolelog "Anchor to: admin:ConsoleLog") admin:ConsoleLog 
Allows listing console logs on terminal
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminkmscreatekey "Anchor to: admin:KMSCreateKey") admin:KMSCreateKey 
Allows creating a new KMS master key
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminkmskeystatus "Anchor to: admin:KMSKeyStatus") admin:KMSKeyStatus 
Allows getting KMS key status
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminserverinfo "Anchor to: admin:ServerInfo") admin:ServerInfo 
Allows listing server info
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminobdinfo "Anchor to: admin:OBDInfo") admin:OBDInfo 
Allows obtaining cluster on-board diagnostics
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminserverupdate "Anchor to: admin:ServerUpdate") admin:ServerUpdate 
Allows AIStor binary update
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminservicerestart "Anchor to: admin:ServiceRestart") admin:ServiceRestart 
Allows restart of AIStor service.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminservicestop "Anchor to: admin:ServiceStop") admin:ServiceStop 
Allows stopping AIStor service.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminconfigupdate "Anchor to: admin:ConfigUpdate") admin:ConfigUpdate 
Allows AIStor config management
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admincreateuser "Anchor to: admin:CreateUser") admin:CreateUser 
Allows creating AIStor user
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admindeleteuser "Anchor to: admin:DeleteUser") admin:DeleteUser 
Allows deleting AIStor user
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminlistusers "Anchor to: admin:ListUsers") admin:ListUsers 
Allows list users permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminenableuser "Anchor to: admin:EnableUser") admin:EnableUser 
Allows enable user permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admindisableuser "Anchor to: admin:DisableUser") admin:DisableUser 
Allows disable user permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admingetuser "Anchor to: admin:GetUser") admin:GetUser 
Allows GET permission on user info
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminaddusertogroup "Anchor to: admin:AddUserToGroup") admin:AddUserToGroup 
Allows adding user to group permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminremoveuserfromgroup "Anchor to: admin:RemoveUserFromGroup") admin:RemoveUserFromGroup 
Allows removing user to group permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admingetgroup "Anchor to: admin:GetGroup") admin:GetGroup 
Allows getting group info
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminlistgroups "Anchor to: admin:ListGroups") admin:ListGroups 
Allows list groups permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminenablegroup "Anchor to: admin:EnableGroup") admin:EnableGroup 
Allows enable group permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admindisablegroup "Anchor to: admin:DisableGroup") admin:DisableGroup 
Allows disable group permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admincreatepolicy "Anchor to: admin:CreatePolicy") admin:CreatePolicy 
Allows create policy permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admindeletepolicy "Anchor to: admin:DeletePolicy") admin:DeletePolicy 
Allows delete policy permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admingetpolicy "Anchor to: admin:GetPolicy") admin:GetPolicy 
Allows get policy permission
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminattachuserorgrouppolicy "Anchor to: admin:AttachUserOrGroupPolicy") admin:AttachUserOrGroupPolicy 
Allows attaching a policy to a user/group
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminlistuserpolicies "Anchor to: admin:ListUserPolicies") admin:ListUserPolicies 
Allows listing user policies
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admincreateserviceaccount "Anchor to: admin:CreateServiceAccount") admin:CreateServiceAccount 
Allows creating AIStor Access Key
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminupdateserviceaccount "Anchor to: admin:UpdateServiceAccount") admin:UpdateServiceAccount 
Allows updating AIStor Access Key
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminremoveserviceaccount "Anchor to: admin:RemoveServiceAccount") admin:RemoveServiceAccount 
Allows deleting AIStor Access Key
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminlistserviceaccounts "Anchor to: admin:ListServiceAccounts") admin:ListServiceAccounts 
Allows listing AIStor Access Key
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminsetbucketquota "Anchor to: admin:SetBucketQuota") admin:SetBucketQuota 
Allows setting bucket quota
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admingetbucketquota "Anchor to: admin:GetBucketQuota") admin:GetBucketQuota 
Allows getting bucket quota
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminsetbuckettarget "Anchor to: admin:SetBucketTarget") admin:SetBucketTarget 
Allows setting bucket target
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admingetbuckettarget "Anchor to: admin:GetBucketTarget") admin:GetBucketTarget 
Allows getting bucket targets
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminsettier "Anchor to: admin:SetTier") admin:SetTier 
Allows creating and modifying remote storage tiers using the `mc ilm tier` commands.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminlisttier "Anchor to: admin:ListTier") admin:ListTier 
Allows listing configured remote storage tiers using the `mc ilm tier` commands.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminbandwidthmonitor "Anchor to: admin:BandwidthMonitor") admin:BandwidthMonitor 
Allows retrieving metrics related to current bandwidth consumption.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminprometheus "Anchor to: admin:Prometheus") admin:Prometheus 
Allows access to AIStor metrics. Only required if AIStor requires authentication for scraping metrics.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminlistbatchjobs "Anchor to: admin:ListBatchJobs") admin:ListBatchJobs 
Allows access to list the active batch jobs.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admindescribebatchjob "Anchor to: admin:DescribeBatchJob") admin:DescribeBatchJob 
Allows access to the see the definition details of a running batch job.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminstartbatchjob "Anchor to: admin:StartBatchJob") admin:StartBatchJob 
Allows user to begin a batch job run.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#admincancelbatchjob "Anchor to: admin:CancelBatchJob") admin:CancelBatchJob 
Allows user to stop a batch job currently in process.
####  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#adminrebalance "Anchor to: admin:Rebalance") admin:Rebalance 
Allows access to start, query, or stop a rebalancing of objects across pools with varying free storage space.
##  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#mc-admin-policy-condition-keys "Anchor to: mc admin Policy Condition Keys") `mc admin` Policy Condition Keys 
AIStor supports the following conditions for use with defining policies for `mc admin` actions.
  * `aws:Referer`
  * `aws:SourceIp`
  * `aws:UserAgent`
  * `aws:SecureTransport`
  * `aws:CurrentTime`
  * `aws:EpochTime`


For complete information on any listed condition key, see the IAM Condition Element Documentation.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#tag-based-policy-conditions "Anchor to: Tag-Based Policy Conditions") Tag-Based Policy Conditions 
Policies can use conditions to limit a user’s access only to objects with a specific tag. AIStor supports tag-based conditionals for policies for selected actions. Use the `s3:ExistingObjectTag/<key>` in the `Condition` statement of the policy.
##  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#policy-variables "Anchor to: Policy Variables") Policy Variables 
AIStor supports using policy variables for automatically substituting context from the authenticated user and/or the operation into the user’s assigned policy or policies. Use the `${POLICYVARIABLE}` format to specify the variable to the policy as part of the `Condition` or `Resource` definition. AIStor policy variables function similarly to AWS IAM policy elements: Variables and tags.
Each AIStor identity provider supports its own set of policy variables:
  * AIStor Policy Variables
  * OpenID Policy Variables
  * Active Directory / LDAP Policy Variables


###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#aistor-policy-variables "Anchor to: AIStor Policy Variables") AIStor Policy Variables 
The following table contains a list of recommended policy variables for use in authorizing AIStor-managed users:
Variable | Description  
---|---  
[aws:referrer](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-referer) | The referrer in the HTTP header for the authenticated API call.  
[aws:SourceIp](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-sourceip) | The source IP in the HTTP header for the authenticated API call.  
[aws:username](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-username) | The name of the user associated with the authenticated API call.  
For example, the following policy uses variables to substitute the authenticated user’s username as part of the `Resource` field such that the user can access only the prefixes that match their username:
```
{ 
   "Version": "2012-10-17", 
   "Statement": [ 
      { 
      "Action": ["s3:ListBucket"],
      "Effect": "Allow", 
      "Resource": ["arn:aws:s3:::mybucket"],
      "Condition": {"StringLike": {"s3:prefix": ["${aws:username}/**"]}}
      },
      {
      "Action": [
      "s3:GetObject",
      "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": ["arn:aws:s3:::mybucket/${aws:username}/**"]
      }
   ]
}

```

AIStor replaces the `${aws:username}` variable in the `Resource` field with the username. AIStor then evaluates the policy and allows or denies access to the requested resource and action.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#openid-policy-variables "Anchor to: OpenID Policy Variables") OpenID Policy Variables 
The following table contains a list of supported policy variables for use in authorizing OIDC-managed users.
Each variable corresponds to a claim returned as part of the authenticated user’s JWT token:
Variable | Description  
---|---  
`jwt:sub` | Returns the `sub` claim for the user.  
`jwt:iss` | Returns the Issuer Identifier claim from the ID token.  
`jwt:aud` | Returns the Audience claim from the ID token.  
`jwt:jti` | Returns the JWT ID claim from the client authentication information.  
`jwt:upn` | Returns the User Principal Name claim from the client authentication information.  
`jwt:name` | Returns the `name` claim for the user.  
`jwt:groups` | Returns the `groups` claim for the user.  
`jwt:given_name` | Returns the `given_name` claim for the user.  
`jwt:family_name` | Returns the `family_name` claim for the user.  
`jwt:middle_name` | Returns the `middle_name` claim for the user.  
`jwt:nickname ` | Returns the `nickname` claim for the user.  
`jwt:preferred_username` | Returns the `preferred_username` claim for the user.  
`jwt:profile` | Returns the `profile` claim for the user.  
`jwt:picture ` | Returns the `picture` claim for the user.  
`jwt:website` | Returns the `website` claim for the user.  
`jwt:email ` | Returns the `email` claim for the user.  
`jwt:gender` | Returns the `gender` claim for the user.  
`jwt:birthdate` | Returns the `birthdate` claim for the user.  
`jwt:phone_number` | Returns the `phone_number` claim for the user.  
`jwt:address` | Returns the `address` claim for the user.  
`jwt:scope` | Returns the `scope` claim for the user.  
`jwt:client_id` | Returns the `client_id` claim for the user.  
See the [OpenID Connect Core 1.0 specification](https://openid.net/specs/openid-connect-core-1_0.html) for more information on these scopes. Your OIDC provider of choice may have more specific documentation.
For example, the following policy uses variables to substitute the authenticated user’s `preferred_username` as part of the `Resource` field such that the user can only access those prefixes which match their username:
```
{ 
   "Version": "2012-10-17",
   "Statement": [ 
      { 
         "Action": ["s3:ListBucket"],
         "Effect": "Allow",
         "Resource": ["arn:aws:s3:::mybucket"],
         "Condition": {"StringLike": {"s3:prefix": ["${jwt:preferred_username}/*"]}}
      },
      {
         "Action": [
            "s3:GetObject",
            "s3:PutObject"
         ],
         "Effect": "Allow",
         "Resource": ["arn:aws:s3:::mybucket/${jwt:preferred_username}/*"]
      } 
   ] 
}

```

AIStor replaces the `${jwt:preferred_username}` variable in the `Resource` field with the value of the `preferred_username` in the JWT token. AIStor then evaluates the policy and grants or revokes access to the requested API and resource.
###  [ ](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#active-directory--ldap-policy-variables "Anchor to: Active Directory / LDAP Policy Variables") Active Directory / LDAP Policy Variables 
The following table contains a list of supported policy variables for use in authorizing AD/LDAP users:
Variable | Description  
---|---  
`ldap:username` | The simple username (`name`) for the authenticated user.This is distinct from the user’s DistinguishedName or CommonName.  
`ldap:user` | The Distinguished Name used by the authenticated user.  
`ldap:groups` | The Group Distinguished Name for the authenticated user.  
For example, the following policy uses variables to substitute the authenticated user’s `name` as part of the `Resource` field such that the user can only access those prefixes which match their name:
```
{ 
   "Version": "2012-10-17",
   "Statement": [ 
      { 
         "Action": ["s3:ListBucket"], 
         "Effect": "Allow", 
         "Resource": ["arn:aws:s3:::mybucket"], 
         "Condition": {"StringLike": {"s3:prefix": ["${ldap:username}/*"]}}
      },
      {
         "Action": [
         "s3:GetObject",
         "s3:PutObject"
         ],
         "Effect": "Allow",
         "Resource": ["arn:aws:s3:::mybucket/${ldap:username}/*"] 
      }
   ] 
}

```

AIStor replaces the `${ldap:username}` variable in the `Resource` field with the value of the authenticated user’s `name`. AIStor then evaluates the policy and grants or revokes access to the requested API and resource.
All rights reserved 2024-Present, MinIO, Inc.
On this page
  * [Built-In Policies](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#built-in-policies)
    * [Assign default policies](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#assign-default-policies)
  * [Policy Document Structure](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#policy-document-structure)
  * [Supported S3 policy actions](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#supported-s3-policy-actions)
    * [Common S3 operations](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#common-s3-operations)
    * [Bucket Configuration](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#bucket-configuration)
    * [Multipart Upload](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#multipart-upload)
    * [Versioning and Retention](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#versioning-and-retention)
    * [Bucket Notifications](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#bucket-notifications)
    * [Object Lifecycle Management](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#object-lifecycle-management)
    * [Object Encryption](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#object-encryption)
    * [Bucket Replication](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#bucket-replication)
  * [Supported S3 Policy Condition Keys](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#supported-s3-policy-condition-keys)
  * [AIStor Extended Condition Keys](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#aistor-extended-condition-keys)
  * [KMS policy action keys](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#kms-policy-action-keys)
  * [`mc admin` Policy Action Keys](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#mc-admin-policy-action-keys)
  * [`mc admin` Policy Condition Keys](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#mc-admin-policy-condition-keys)
    * [Tag-Based Policy Conditions](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#tag-based-policy-conditions)
  * [Policy Variables](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#policy-variables)
    * [AIStor Policy Variables](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#aistor-policy-variables)
    * [OpenID Policy Variables](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#openid-policy-variables)
    * [Active Directory / LDAP Policy Variables](https://docs.min.io/enterprise/aistor-object-store/administration/iam/access/#active-directory--ldap-policy-variables)




## Source Information
- URL: https://min.io/docs/minio/linux/administration/identity-access-management/policy-based-access-control.html
- Harvested: 2025-08-19T22:09:00.395028
- Profile: deep_research
- Agent: architect
