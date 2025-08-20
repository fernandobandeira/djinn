---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T22:08:32.216508'
profile: deep_research
source: https://min.io/docs/minio/linux/operations/server-side-encryption.html
topic: MinIO Security Features and Server-Side Encryption
---

# MinIO Security Features and Server-Side Encryption

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
  2. [Operations](https://docs.min.io/enterprise/aistor-object-store/operations/)
  3. [Data Encryption](https://docs.min.io/enterprise/aistor-object-store/operations/server-side-encryption/baremetal/)
  4. [AIStor Key Manager on Baremetal](https://docs.min.io/enterprise/aistor-object-store/operations/server-side-encryption/baremetal/)


# AIStor Key Manager on Baremetal
This page explains how to enable Server-Side Encryption (SSE) with [AIStor Key Manager](https://docs.min.io/enterprise/aistor-key-manager/?ref=docs-internal) as the Key Management Service (KMS).
Important
Enabling SSE on an AIStor deployment automatically encrypts the backend data for the deployment using the specified default encryption key.
AIStor requires access to the Key Manager to decrypt the backend and start normally. You cannot disable or reset encryption of the backend.
##  [ ](https://docs.min.io/enterprise/aistor-object-store/operations/server-side-encryption/baremetal/#prerequisites "Anchor to: Prerequisites") Prerequisites 
  * [`mc`](https://docs.min.io/enterprise/aistor-object-store/reference/cli/) installed with network access to the cluster
  * A configured [`alias`](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-alias/) for AIStor
  * A Key Manager deployment with a pre-configured identity for use with AIStor. See Key Manager [Installation and Management](https://docs.min.io/enterprise/aistor-key-manager/installation/linux/#procedure?ref=docs-internal).


##  [ ](https://docs.min.io/enterprise/aistor-object-store/operations/server-side-encryption/baremetal/#how-to-do-it "Anchor to: How to do it") How to do it 
  1. Retrieve credentials for Key Manager
AIStor Key Manager requires authorization for all access and operations. The [Key Manager installation guide](https://docs.min.io/enterprise/aistor-key-manager/installation/linux/?ref=docs-internal) includes steps for generating an enclave and API Key for use by AIStor.
The enclave should be unique to the object store and named to facilitate easy identification among other configured enclaves.
The key resembles the following:
```
k1:2COl4dS3G-cjHa3Q-9fUmOrq8yL0Q7a12HH_Yi0oiLw

```

  2. Modify the environment file at `/etc/default/minio` by adding the following environment key-value pairs. Modify the example values to match your deployment configuration.
```
# AIStor Key Manager settings
# Provide a comma-separated array of hostnames for the AIStor Key Manager.
# If you use a dedicated load balancer or similar network control plane to manage connections
# to the Key Manager, provide that hostname here instead.
#
# All AIStor nodes **must** have bidirectional access to the specified Key Manager hosts
# to ensure normal cryptographic operations.
MINIO_KMS_SERVER="https://kms-1.example.net,https://kms-2.example.net,https://kms3-example.net"
# Specify the name for the encryption key AIStor uses for backend and default bucket encryption.
# Consider specifying a unique key name to facilitate easy identification among other stored keys.
#
# Do not modify the MINIO_KMS_SSE_KEY value after setup.
# AIStor requires this key to start successfully.
MINIO_KMS_SSE_KEY="object-store-primary-default-key"
# Specify the Key Manager enclave to use
MINIO_KMS_ENCLAVE="object-store-primary"
# Specify the AIStor Key Manager API key to use for authenticating operations.
# The API key must have permission to access and perform operations in the Key Manager enclave
MINIO_KMS_API_KEY="k1:2COl4dS3G-cjHa3Q-9fUmOrq8yL0Q7a12HH_Yi0oiLw"

```



Important
You must modify **all** AIStor nodes to have matching environment files. Repeat this step for each AIStor node in the deployment. 
  1. Restart the deployment and load the new settings with [`mc admin service restart`](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-alias/).
Alternatively, use an orchestration service, script, or other utility to perform the equivalent of `systemctl restart minio` on all nodes simultaneously.
Run `journalctl -u minio` to monitor the restart progress and check that the AIStor starts successfully.
  2. You can now configure automatic SSE-KMS for all objects written to a given bucket.
Run [`mc admin kms key create`](https://docs.min.io/enterprise/aistor-object-store/reference/cli/admin/mc-admin-kms-key/) to create a new data encryption key for use with a bucket.
You can then run [`mc encrypt set`](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-encrypt/mc-encrypt-set/) to configure bucket default encryption using SSE-KMS.
The following example assumes an [`alias`](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-alias/) of `primary` associated to the AIStor.
```
mc admin kms key create primary bucket-data-encryption-key
mc mb primary/data
mc encrypt set sse-kms bucket-data-encryption-key primary/data

```

If you create an object in the bucket and run [`mc stat`](https://docs.min.io/enterprise/aistor-object-store/reference/cli/mc-stat/) against the object, the output includes encryption information as part of the object metadata.
Disabling, deleting, or otherwise removing access to the key renders the encrypted object as unreadable.


All rights reserved 2024-Present, MinIO, Inc.
On this page
  * [Prerequisites](https://docs.min.io/enterprise/aistor-object-store/operations/server-side-encryption/baremetal/#prerequisites)
  * [How to do it](https://docs.min.io/enterprise/aistor-object-store/operations/server-side-encryption/baremetal/#how-to-do-it)




## Source Information
- URL: https://min.io/docs/minio/linux/operations/server-side-encryption.html
- Harvested: 2025-08-19T22:08:32.216508
- Profile: deep_research
- Agent: architect
