---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:54:13.777874'
profile: code_examples
source: https://aquasecurity.github.io/trivy/
topic: Trivy Security Scanner for Go Applications
---

# Trivy Security Scanner for Go Applications

You're not viewing the latest version of the documentation. [ **Click here to go to latest.** ](https://trivy.dev/)
[ ![logo](https://trivy.dev/latest/imgs/logo-white.svg) ](https://trivy.dev/latest/ "Trivy")
Trivy 
v0.65
  * [dev](https://trivy.dev/dev/)
  * [v0.65](https://trivy.dev/v0.65/)
  * [v0.64](https://trivy.dev/v0.64/)
  * [v0.63](https://trivy.dev/v0.63/)
  * [v0.62](https://trivy.dev/v0.62/)
  * [v0.61](https://trivy.dev/v0.61/)
  * [v0.60](https://trivy.dev/v0.60/)
  * [v0.59](https://trivy.dev/v0.59/)
  * [v0.58](https://trivy.dev/v0.58/)
  * [v0.57](https://trivy.dev/v0.57/)
  * [v0.56](https://trivy.dev/v0.56/)
  * [v0.55](https://trivy.dev/v0.55/)
  * [v0.54](https://trivy.dev/v0.54/)
  * [v0.53](https://trivy.dev/v0.53/)
  * [v0.52](https://trivy.dev/v0.52/)
  * [v0.51](https://trivy.dev/v0.51/)
  * [v0.50](https://trivy.dev/v0.50/)
  * [v0.49](https://trivy.dev/v0.49/)
  * [v0.48](https://trivy.dev/v0.48/)
  * [v0.47](https://trivy.dev/v0.47/)
  * [v0.46](https://trivy.dev/v0.46/)
  * [v0.45](https://trivy.dev/v0.45/)
  * [v0.44](https://trivy.dev/v0.44/)
  * [v0.43](https://trivy.dev/v0.43/)
  * [v0.42](https://trivy.dev/v0.42/)
  * [v0.41](https://trivy.dev/v0.41/)
  * [v0.40](https://trivy.dev/v0.40/)
  * [v0.39](https://trivy.dev/v0.39/)
  * [v0.38](https://trivy.dev/v0.38/)
  * [v0.37](https://trivy.dev/v0.37/)
  * [v0.36](https://trivy.dev/v0.36/)
  * [v0.35](https://trivy.dev/v0.35/)
  * [v0.34](https://trivy.dev/v0.34/)
  * [v0.33](https://trivy.dev/v0.33/)
  * [v0.32](https://trivy.dev/v0.32/)
  * [v0.31](https://trivy.dev/v0.31.3/)
  * [v0.30](https://trivy.dev/v0.30.4/)
  * [v0.29](https://trivy.dev/v0.29.2/)
  * [v0.28](https://trivy.dev/v0.28.1/)
  * [v0.27](https://trivy.dev/v0.27.1/)
  * [v0.26](https://trivy.dev/v0.26.0/)
  * [v0.25](https://trivy.dev/v0.25.4/)
  * [v0.24](https://trivy.dev/v0.24.4/)
  * [v0.23](https://trivy.dev/v0.23.0/)
  * [v0.22](https://trivy.dev/v0.22.0/)
  * [v0.21](https://trivy.dev/v0.21.3/)
  * [v0.20](https://trivy.dev/v0.20.2/)
  * [v0.19](https://trivy.dev/v0.19.2/)
  * [v0.18](https://trivy.dev/v0.18.3/)
  * [v0.17](https://trivy.dev/v0.17.2/)
  * [v0.16](https://trivy.dev/v0.16.0/)
  * [v0.15](https://trivy.dev/v0.15.0/)


Home 
Type to start searching
[ GitHub 
  * v0.65.0
  * 28k
  * 2.7k

](https://github.com/aquasecurity/trivy "Go to repository")
  * [ Home ](https://trivy.dev/latest/)
  * [ Getting Started ](https://trivy.dev/latest/getting-started/)
  * [ Tutorials ](https://trivy.dev/latest/tutorials/overview/)
  * [ Docs ](https://trivy.dev/latest/docs/)
  * [ Ecosystem ](https://trivy.dev/latest/ecosystem/)
  * [ Contributing ](https://trivy.dev/latest/community/principles/)
  * [ Enterprise ](https://trivy.dev/latest/commercial/compare/)
  * [ Partners ](https://trivy.dev/latest/partners/)


![Trivy Logo](https://trivy.dev/latest/assets/images/trivy_logo_horizontal_white.svg)
#  The all-in-one open source security scanner 
##  Use Trivy to find vulnerabilities (CVE) & misconfigurations (IaC) across code repositories, binary artifacts, container images, Kubernetes clusters, and more. All in one tool! 
[Get started](https://trivy.dev/latest/getting-started) [Read the Docs](https://trivy.dev/latest/docs)
It's all about the community!
Trivy is praised by professionals worldwide. Are you a Trivy fan as well? We'd love to hear from you! 
[Share your story](https://github.com/aquasecurity/trivy/discussions/categories/adopters) [Discuss on GitHub](https://github.com/aquasecurity/trivy/discussions/)
Cristiano Corrado, Wise
[wise-engineering](https://medium.com/wise-engineering/our-application-security-journey-part-1-fb7d449a7126)
"The discovery process led us to evaluate multiple open-source tools ... The final decision was to use Trivy"
Saim Safdar
[@cloudnativeboy](https://x.com/cloudnativeboy/status/1538840080873291777)
"my favorite @Docker extension. (Trivy)"
Jonathan Gonzalez V., CloudNativePG
[@sxd](https://x.com/sxd/status/1555688296092672000)
"Thanks @AquaSecTeam for creating Trivy and help us to improve @CloudNativePg security"
Andy Roberts, Vista
[@andyr8939](https://x.com/andyr8939/status/1523402643397562368)
"I've tried a few and all have pros/cons but I found that Trivy is the best of them"
Ali Faraj, Antigen Security
[@andyr8939](https://www.linkedin.com/posts/activity-6932020213307625472-zWs8/)
"Trivy from Aqua Security is my new favorite tool... It's such a powerful tool with the ability to generate SBOMs, find vulnerabilities, misconfigurations, and secrets!"
Sam White, GitLab
[Case Study](https://www.aquasec.com/blog/trivy-scanner-gitlab-case-study/)
"Trivy was a clear leader in the market as far as features, functionality, and capabilities"
Ariadne Conill, Alpine Security
@ariadneconill
...the tl;dr is basically Aqua's Trivy is the best one, all of the other ones are a waste of time 
Harbor Team
[Harbor blog](https://goharbor.io/blog/harbor-2.0/)
"Trivy takes container image scanning to higher levels of usability and performance."
Milind Gadre, Mirantis
[Mirantis](https://www.aquasec.com/news/trivy-vulnerability-scanner-aqua-security-adopted-by-leading-cloud-native-platforms/)
"After evaluating several leading options for open source vulnerability scanning, Trivy really stood out"
Jerry Gambli
@JGamblin
The way the @AquaSecTeam team has turned Trivy into the best open-source vulnerability scanner in such a short time is really amazing.
Yaney
[y4ney](https://github.com/aquasecurity/trivy/discussions/4738#discussioncomment-6335342)
"Trivy is, by far, the best open-source tool for cloud-native security that I have ever used"
Ulises Galeano, MasterCard
[ulises-galeano](https://www.linkedin.com/posts/ulises-galeano_aqua-securitys-trivy-adds-cspm-capabilities-activity-6965797718019502080-8ah_?utm_source=share&utm_medium=member_desktop)
"This tool just keeps getting better and better..."
Damian Naprawa
[@DamianNaprawa](https://x.com/DamianNaprawa/status/1469229068419842049)
"So happy to see collaboration between @Azure and @AquaSecTeam on scanning container images in Azure Container Registry CI/CD workflows using such a great tool - Trivy."
Mustafa Akin, Resmo
[mustafaakin](https://github.com/aquasecurity/trivy/discussions/5898#discussioncomment-8257890)
"I love how Trivy democratized dependency scanning to the masses as a free and extremely easy to use tool, with also a permissive license. This used to be a gated community with predatory security vendors charging premium, and they were not half as good as Trivy."
dynaptik, Deutsche Bahn
[@dynaptik](https://x.com/dynaptik/status/1760063329941291098?s=20)
"Trivy easily for what it brings to the table (secret scanning, vuln scan, license check) and little effort required. Bang for the buck it's pretty amazing"
Cristiano Corrado, Wise
[wise-engineering](https://medium.com/wise-engineering/our-application-security-journey-part-1-fb7d449a7126)
"The discovery process led us to evaluate multiple open-source tools ... The final decision was to use Trivy"
Saim Safdar
[@cloudnativeboy](https://x.com/cloudnativeboy/status/1538840080873291777)
"my favorite @Docker extension. (Trivy)"
Jonathan Gonzalez V., CloudNativePG
[@sxd](https://x.com/sxd/status/1555688296092672000)
"Thanks @AquaSecTeam for creating Trivy and help us to improve @CloudNativePg security"
Andy Roberts, Vista
[@andyr8939](https://x.com/andyr8939/status/1523402643397562368)
"I've tried a few and all have pros/cons but I found that Trivy is the best of them"
Ali Faraj, Antigen Security
[@andyr8939](https://www.linkedin.com/posts/activity-6932020213307625472-zWs8/)
"Trivy from Aqua Security is my new favorite tool... It's such a powerful tool with the ability to generate SBOMs, find vulnerabilities, misconfigurations, and secrets!"
Sam White, GitLab
[Case Study](https://www.aquasec.com/blog/trivy-scanner-gitlab-case-study/)
"Trivy was a clear leader in the market as far as features, functionality, and capabilities"
Ariadne Conill, Alpine Security
@ariadneconill
...the tl;dr is basically Aqua's Trivy is the best one, all of the other ones are a waste of time 
Harbor Team
[Harbor blog](https://goharbor.io/blog/harbor-2.0/)
"Trivy takes container image scanning to higher levels of usability and performance."
Milind Gadre, Mirantis
[Mirantis](https://www.aquasec.com/news/trivy-vulnerability-scanner-aqua-security-adopted-by-leading-cloud-native-platforms/)
"After evaluating several leading options for open source vulnerability scanning, Trivy really stood out"
Jerry Gambli
@JGamblin
The way the @AquaSecTeam team has turned Trivy into the best open-source vulnerability scanner in such a short time is really amazing.
Yaney
[y4ney](https://github.com/aquasecurity/trivy/discussions/4738#discussioncomment-6335342)
"Trivy is, by far, the best open-source tool for cloud-native security that I have ever used"
Ulises Galeano, MasterCard
[ulises-galeano](https://www.linkedin.com/posts/ulises-galeano_aqua-securitys-trivy-adds-cspm-capabilities-activity-6965797718019502080-8ah_?utm_source=share&utm_medium=member_desktop)
"This tool just keeps getting better and better..."
Damian Naprawa
[@DamianNaprawa](https://x.com/DamianNaprawa/status/1469229068419842049)
"So happy to see collaboration between @Azure and @AquaSecTeam on scanning container images in Azure Container Registry CI/CD workflows using such a great tool - Trivy."
Mustafa Akin, Resmo
[mustafaakin](https://github.com/aquasecurity/trivy/discussions/5898#discussioncomment-8257890)
"I love how Trivy democratized dependency scanning to the masses as a free and extremely easy to use tool, with also a permissive license. This used to be a gated community with predatory security vendors charging premium, and they were not half as good as Trivy."
dynaptik, Deutsche Bahn
[@dynaptik](https://x.com/dynaptik/status/1760063329941291098?s=20)
"Trivy easily for what it brings to the table (secret scanning, vuln scan, license check) and little effort required. Bang for the buck it's pretty amazing"
Cristiano Corrado, Wise
[wise-engineering](https://medium.com/wise-engineering/our-application-security-journey-part-1-fb7d449a7126)
"The discovery process led us to evaluate multiple open-source tools ... The final decision was to use Trivy"
Saim Safdar
[@cloudnativeboy](https://x.com/cloudnativeboy/status/1538840080873291777)
"my favorite @Docker extension. (Trivy)"
Jonathan Gonzalez V., CloudNativePG
[@sxd](https://x.com/sxd/status/1555688296092672000)
"Thanks @AquaSecTeam for creating Trivy and help us to improve @CloudNativePg security"
Andy Roberts, Vista
[@andyr8939](https://x.com/andyr8939/status/1523402643397562368)
"I've tried a few and all have pros/cons but I found that Trivy is the best of them"
Ali Faraj, Antigen Security
[@andyr8939](https://www.linkedin.com/posts/activity-6932020213307625472-zWs8/)
"Trivy from Aqua Security is my new favorite tool... It's such a powerful tool with the ability to generate SBOMs, find vulnerabilities, misconfigurations, and secrets!"
  * 1
  * 2
  * 3


[ ![logo](https://trivy.dev/latest/imgs/logo-white.svg) ](https://trivy.dev/latest/ "Trivy") Trivy 
[ GitHub 
  * v0.65.0
  * 28k
  * 2.7k

](https://github.com/aquasecurity/trivy "Go to repository")
  * [ Home  ](https://trivy.dev/latest/)
  * Getting Started  Getting Started 
    * [ First steps  ](https://trivy.dev/latest/getting-started/)
    * [ Installation  ](https://trivy.dev/latest/getting-started/installation/)
    * [ Signature Verification  ](https://trivy.dev/latest/getting-started/signature-verification/)
    * [ FAQ  ](https://trivy.dev/latest/getting-started/faq/)
  * Tutorials  Tutorials 
    * [ Overview  ](https://trivy.dev/latest/tutorials/overview/)
    * CI/CD  CI/CD 
      * [ Overview  ](https://trivy.dev/latest/tutorials/integrations/)
      * [ GitHub Actions  ](https://trivy.dev/latest/tutorials/integrations/github-actions/)
      * [ CircleCI  ](https://trivy.dev/latest/tutorials/integrations/circleci/)
      * [ Travis CI  ](https://trivy.dev/latest/tutorials/integrations/travis-ci/)
      * [ GitLab CI  ](https://trivy.dev/latest/tutorials/integrations/gitlab-ci/)
      * [ Bitbucket Pipelines  ](https://trivy.dev/latest/tutorials/integrations/bitbucket/)
      * [ AWS CodePipeline  ](https://trivy.dev/latest/tutorials/integrations/aws-codepipeline/)
      * [ AWS Security Hub  ](https://trivy.dev/latest/tutorials/integrations/aws-security-hub/)
      * [ Azure  ](https://trivy.dev/latest/tutorials/integrations/azure-devops/)
    * Kubernetes  Kubernetes 
      * [ Cluster Scanning  ](https://trivy.dev/latest/tutorials/kubernetes/cluster-scanning/)
      * [ Kyverno  ](https://trivy.dev/latest/tutorials/kubernetes/kyverno/)
      * [ GitOps  ](https://trivy.dev/latest/tutorials/kubernetes/gitops/)
    * Misconfiguration  Misconfiguration 
      * [ Terraform scanning  ](https://trivy.dev/latest/tutorials/misconfiguration/terraform/)
      * [ Custom Checks with Rego  ](https://trivy.dev/latest/tutorials/misconfiguration/custom-checks/)
    * Signing  Signing 
      * [ Vulnerability Scan Record Attestation  ](https://trivy.dev/latest/tutorials/signing/vuln-attestation/)
    * Shell  Shell 
      * [ Completion  ](https://trivy.dev/latest/tutorials/shell/shell-completion/)
    * Additional Resources  Additional Resources 
      * [ Additional Resources  ](https://trivy.dev/latest/tutorials/additional-resources/references/)
      * [ Community References  ](https://trivy.dev/latest/tutorials/additional-resources/community/)
      * [ CKS Reference  ](https://trivy.dev/latest/tutorials/additional-resources/cks/)
  * Docs  Docs 
    * [ Overview  ](https://trivy.dev/latest/docs/)
    * Target  Target 
      * [ Container Image  ](https://trivy.dev/latest/docs/target/container_image/)
      * [ Filesystem  ](https://trivy.dev/latest/docs/target/filesystem/)
      * [ Rootfs  ](https://trivy.dev/latest/docs/target/rootfs/)
      * [ Code Repository  ](https://trivy.dev/latest/docs/target/repository/)
      * [ Virtual Machine Image  ](https://trivy.dev/latest/docs/target/vm/)
      * [ Kubernetes  ](https://trivy.dev/latest/docs/target/kubernetes/)
      * [ SBOM  ](https://trivy.dev/latest/docs/target/sbom/)
    * Scanner  Scanner 
      * [ Vulnerability  ](https://trivy.dev/latest/docs/scanner/vulnerability/)
      * Misconfiguration  Misconfiguration 
        * [ Overview  ](https://trivy.dev/latest/docs/scanner/misconfiguration/)
        * [ Configuration  ](https://trivy.dev/latest/docs/scanner/misconfiguration/config/config/)
        * Policy  Policy 
          * [ Built-in Checks  ](https://trivy.dev/latest/docs/scanner/misconfiguration/check/builtin/)
        * Custom Checks  Custom Checks 
          * [ Overview  ](https://trivy.dev/latest/docs/scanner/misconfiguration/custom/)
          * [ Data  ](https://trivy.dev/latest/docs/scanner/misconfiguration/custom/data/)
          * [ Combine  ](https://trivy.dev/latest/docs/scanner/misconfiguration/custom/combine/)
          * [ Selectors  ](https://trivy.dev/latest/docs/scanner/misconfiguration/custom/selectors/)
          * [ Schemas  ](https://trivy.dev/latest/docs/scanner/misconfiguration/custom/schema/)
          * [ Testing  ](https://trivy.dev/latest/docs/scanner/misconfiguration/custom/testing/)
          * [ Debugging Policies  ](https://trivy.dev/latest/docs/scanner/misconfiguration/custom/debug/)
          * [ Contribute Checks  ](https://trivy.dev/latest/docs/scanner/misconfiguration/custom/contribute-checks/)
      * [ Secret  ](https://trivy.dev/latest/docs/scanner/secret/)
      * [ License  ](https://trivy.dev/latest/docs/scanner/license/)
    * Coverage  Coverage 
      * [ Overview  ](https://trivy.dev/latest/docs/coverage/)
      * OS  OS 
        * [ Overview  ](https://trivy.dev/latest/docs/coverage/os/)
        * [ AlmaLinux  ](https://trivy.dev/latest/docs/coverage/os/alma/)
        * [ Alpine Linux  ](https://trivy.dev/latest/docs/coverage/os/alpine/)
        * [ Amazon Linux  ](https://trivy.dev/latest/docs/coverage/os/amazon/)
        * [ Azure Linux (CBL-Mariner)  ](https://trivy.dev/latest/docs/coverage/os/azure/)
        * [ Bottlerocket  ](https://trivy.dev/latest/docs/coverage/os/bottlerocket/)
        * [ CentOS  ](https://trivy.dev/latest/docs/coverage/os/centos/)
        * [ Chainguard  ](https://trivy.dev/latest/docs/coverage/os/chainguard/)
        * [ Debian  ](https://trivy.dev/latest/docs/coverage/os/debian/)
        * [ Echo  ](https://trivy.dev/latest/docs/coverage/os/echo/)
        * [ MinimOS  ](https://trivy.dev/latest/docs/coverage/os/minimos/)
        * [ Oracle Linux  ](https://trivy.dev/latest/docs/coverage/os/oracle/)
        * [ Photon OS  ](https://trivy.dev/latest/docs/coverage/os/photon/)
        * [ Red Hat  ](https://trivy.dev/latest/docs/coverage/os/rhel/)
        * [ Rocky Linux  ](https://trivy.dev/latest/docs/coverage/os/rocky/)
        * [ SUSE  ](https://trivy.dev/latest/docs/coverage/os/suse/)
        * [ Ubuntu  ](https://trivy.dev/latest/docs/coverage/os/ubuntu/)
        * [ Wolfi  ](https://trivy.dev/latest/docs/coverage/os/wolfi/)
        * [ Google Distroless (Images)  ](https://trivy.dev/latest/docs/coverage/os/google-distroless/)
      * Language  Language 
        * [ Overview  ](https://trivy.dev/latest/docs/coverage/language/)
        * [ C/C++  ](https://trivy.dev/latest/docs/coverage/language/c/)
        * [ Dart  ](https://trivy.dev/latest/docs/coverage/language/dart/)
        * [ .NET  ](https://trivy.dev/latest/docs/coverage/language/dotnet/)
        * [ Elixir  ](https://trivy.dev/latest/docs/coverage/language/elixir/)
        * [ Go  ](https://trivy.dev/latest/docs/coverage/language/golang/)
        * [ Java  ](https://trivy.dev/latest/docs/coverage/language/java/)
        * [ Node.js  ](https://trivy.dev/latest/docs/coverage/language/nodejs/)
        * [ PHP  ](https://trivy.dev/latest/docs/coverage/language/php/)
        * [ Python  ](https://trivy.dev/latest/docs/coverage/language/python/)
        * [ Ruby  ](https://trivy.dev/latest/docs/coverage/language/ruby/)
        * [ Rust  ](https://trivy.dev/latest/docs/coverage/language/rust/)
        * [ Swift  ](https://trivy.dev/latest/docs/coverage/language/swift/)
        * [ Julia  ](https://trivy.dev/latest/docs/coverage/language/julia/)
      * IaC  IaC 
        * [ Overview  ](https://trivy.dev/latest/docs/coverage/iac/)
        * [ Azure ARM Template  ](https://trivy.dev/latest/docs/coverage/iac/azure-arm/)
        * [ CloudFormation  ](https://trivy.dev/latest/docs/coverage/iac/cloudformation/)
        * [ Docker  ](https://trivy.dev/latest/docs/coverage/iac/docker/)
        * [ Helm  ](https://trivy.dev/latest/docs/coverage/iac/helm/)
        * [ Kubernetes  ](https://trivy.dev/latest/docs/coverage/iac/kubernetes/)
        * [ Terraform  ](https://trivy.dev/latest/docs/coverage/iac/terraform/)
      * Others  Others 
        * [ Overview  ](https://trivy.dev/latest/docs/coverage/others/)
        * [ Bitnami Images  ](https://trivy.dev/latest/docs/coverage/others/bitnami/)
        * [ Conda  ](https://trivy.dev/latest/docs/coverage/others/conda/)
        * [ Root.io Images  ](https://trivy.dev/latest/docs/coverage/others/rootio/)
        * [ RPM Archives  ](https://trivy.dev/latest/docs/coverage/others/rpm/)
      * [ Kubernetes  ](https://trivy.dev/latest/docs/coverage/kubernetes/)
    * Configuration  Configuration 
      * [ Overview  ](https://trivy.dev/latest/docs/configuration/)
      * [ Filtering  ](https://trivy.dev/latest/docs/configuration/filtering/)
      * [ Selecting Files  ](https://trivy.dev/latest/docs/configuration/skipping/)
      * [ Reporting  ](https://trivy.dev/latest/docs/configuration/reporting/)
      * [ Cache  ](https://trivy.dev/latest/docs/configuration/cache/)
      * [ Databases  ](https://trivy.dev/latest/docs/configuration/db/)
      * [ Others  ](https://trivy.dev/latest/docs/configuration/others/)
    * Supply Chain  Supply Chain 
      * [ SBOM  ](https://trivy.dev/latest/docs/supply-chain/sbom/)
      * Attestation  Attestation 
        * [ SBOM  ](https://trivy.dev/latest/docs/supply-chain/attestation/sbom/)
        * [ Cosign Vulnerability Scan Record  ](https://trivy.dev/latest/docs/supply-chain/attestation/vuln/)
        * [ SBOM Attestation in Rekor  ](https://trivy.dev/latest/docs/supply-chain/attestation/rekor/)
      * VEX  VEX 
        * [ Overview  ](https://trivy.dev/latest/docs/supply-chain/vex/)
        * [ VEX Repository  ](https://trivy.dev/latest/docs/supply-chain/vex/repo/)
        * [ Local VEX Files  ](https://trivy.dev/latest/docs/supply-chain/vex/file/)
        * [ VEX SBOM Reference  ](https://trivy.dev/latest/docs/supply-chain/vex/sbom-ref/)
        * [ VEX Attestation  ](https://trivy.dev/latest/docs/supply-chain/vex/oci/)
    * Compliance  Compliance 
      * [ Built-in Compliance  ](https://trivy.dev/latest/docs/compliance/compliance/)
      * [ Custom Compliance  ](https://trivy.dev/latest/docs/compliance/contrib-compliance/)
    * Plugins  Plugins 
      * [ Overview  ](https://trivy.dev/latest/docs/plugin/)
      * [ User guide  ](https://trivy.dev/latest/docs/plugin/user-guide/)
      * [ Developer guide  ](https://trivy.dev/latest/docs/plugin/developer-guide/)
    * Advanced  Advanced 
      * [ Modules  ](https://trivy.dev/latest/docs/advanced/modules/)
      * [ Connectivity and Network considerations  ](https://trivy.dev/latest/docs/advanced/air-gap/)
      * [ Self-Hosting Trivy's Databases  ](https://trivy.dev/latest/docs/advanced/self-hosting/)
      * Container Image  Container Image 
        * [ Embed in Dockerfile  ](https://trivy.dev/latest/docs/advanced/container/embed-in-dockerfile/)
        * [ Unpacked container image filesystem  ](https://trivy.dev/latest/docs/advanced/container/unpacked-filesystem/)
        * Private Docker Registries  Private Docker Registries 
          * [ Overview  ](https://trivy.dev/latest/docs/advanced/private-registries/)
          * [ Docker Hub  ](https://trivy.dev/latest/docs/advanced/private-registries/docker-hub/)
          * [ AWS ECR (Elastic Container Registry)  ](https://trivy.dev/latest/docs/advanced/private-registries/ecr/)
          * [ GCR (Google Container Registry)  ](https://trivy.dev/latest/docs/advanced/private-registries/gcr/)
          * [ ACR (Azure Container Registry)  ](https://trivy.dev/latest/docs/advanced/private-registries/acr/)
          * [ Self-Hosted  ](https://trivy.dev/latest/docs/advanced/private-registries/self/)
      * [ Usage Telemetry  ](https://trivy.dev/latest/docs/advanced/telemetry/)
    * References  References 
      * Configuration  Configuration 
        * CLI  CLI 
          * [ Overview  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy/)
          * [ Clean  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_clean/)
          * [ Config  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_config/)
          * [ Convert  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_convert/)
          * [ Filesystem  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_filesystem/)
          * [ Image  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_image/)
          * [ Kubernetes  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_kubernetes/)
          * Module  Module 
            * [ Module  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_module/)
            * [ Module Install  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_module_install/)
            * [ Module Uninstall  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_module_uninstall/)
          * Plugin  Plugin 
            * [ Plugin  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_plugin/)
            * [ Plugin Info  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_plugin_info/)
            * [ Plugin Install  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_plugin_install/)
            * [ Plugin List  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_plugin_list/)
            * [ Plugin Run  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_plugin_run/)
            * [ Plugin Uninstall  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_plugin_uninstall/)
            * [ Plugin Update  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_plugin_update/)
            * [ Plugin Upgrade  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_plugin_upgrade/)
            * [ Plugin Search  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_plugin_search/)
          * Registry  Registry 
            * [ Registry  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_registry/)
            * [ Registry Login  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_registry_login/)
            * [ Registry Logout  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_registry_logout/)
          * [ Repository  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_repository/)
          * [ Rootfs  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_rootfs/)
          * [ SBOM  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_sbom/)
          * [ Server  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_server/)
          * [ Version  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_version/)
          * VEX  VEX 
            * [ VEX  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_vex/)
            * [ VEX Download  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_vex_repo_download/)
            * [ VEX Init  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_vex_repo_init/)
            * [ VEX List  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_vex_repo_list/)
            * [ VEX Repo  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_vex_repo/)
          * [ VM  ](https://trivy.dev/latest/docs/references/configuration/cli/trivy_vm/)
        * [ Config file  ](https://trivy.dev/latest/docs/references/configuration/config-file/)
      * Modes  Modes 
        * [ Standalone  ](https://trivy.dev/latest/docs/references/modes/standalone/)
        * [ Client/Server  ](https://trivy.dev/latest/docs/references/modes/client-server/)
      * [ Troubleshooting  ](https://trivy.dev/latest/docs/references/troubleshooting/)
      * [ Terminology  ](https://trivy.dev/latest/docs/references/terminology/)
      * [ Abbreviations  ](https://trivy.dev/latest/docs/references/abbreviations/)
  * Ecosystem  Ecosystem 
    * [ Overview  ](https://trivy.dev/latest/ecosystem/)
    * [ CI/CD  ](https://trivy.dev/latest/ecosystem/cicd/)
    * [ IDE and Dev tools  ](https://trivy.dev/latest/ecosystem/ide/)
    * [ Production and Clouds  ](https://trivy.dev/latest/ecosystem/prod/)
    * [ Reporting  ](https://trivy.dev/latest/ecosystem/reporting/)
  * Contributing  Contributing 
    * [ Principles  ](https://trivy.dev/latest/community/principles/)
    * How to contribute  How to contribute 
      * [ Issues  ](https://trivy.dev/latest/community/contribute/issue/)
      * [ Discussions  ](https://trivy.dev/latest/community/contribute/discussion/)
      * [ Pull Requests  ](https://trivy.dev/latest/community/contribute/pr/)
    * Contribute Rego Checks  Contribute Rego Checks 
      * [ Overview  ](https://trivy.dev/latest/community/contribute/checks/overview/)
      * [ Add Service Support  ](https://trivy.dev/latest/community/contribute/checks/service-support/)
    * Maintainer  Maintainer 
      * [ PR Review  ](https://trivy.dev/latest/community/maintainer/pr-review/)
      * [ Release Flow  ](https://trivy.dev/latest/community/maintainer/release-flow/)
      * [ Backporting  ](https://trivy.dev/latest/community/maintainer/backporting/)
      * [ Help Wanted  ](https://trivy.dev/latest/community/maintainer/help-wanted/)
      * [ Triage  ](https://trivy.dev/latest/community/maintainer/triage/)
  * Enterprise  Enterprise 
    * [ Comparison  ](https://trivy.dev/latest/commercial/compare/)
    * [ Contact Us  ](https://trivy.dev/latest/commercial/contact/)
  * [ Partners  ](https://trivy.dev/latest/partners/)


[ Next  First steps  ](https://trivy.dev/latest/getting-started/)
[ ](https://twitter.com/AquaTrivy "twitter.com") [ ](https://github.com/aquasecurity/trivy "github.com")


## Source Information
- URL: https://aquasecurity.github.io/trivy/
- Harvested: 2025-08-20T00:54:13.777874
- Profile: code_examples
- Agent: architect
