---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T20:33:47.656794'
profile: deep_research
source: https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html
topic: OWASP File Upload Security Best Practices
---

# OWASP File Upload Security Best Practices

[ Skip to content ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-upload-cheat-sheet)
[ ![logo](https://cheatsheetseries.owasp.org/assets/OWASP_Logo.svg) ](https://cheatsheetseries.owasp.org/index.html "OWASP Cheat Sheet Series")
OWASP Cheat Sheet Series 
File Upload 
Type to start searching
[ OWASP/CheatSheetSeries 
  * 30.1k
  * 4.2k

](https://github.com/OWASP/CheatSheetSeries "Go to repository")
[ ![logo](https://cheatsheetseries.owasp.org/assets/OWASP_Logo.svg) ](https://cheatsheetseries.owasp.org/index.html "OWASP Cheat Sheet Series") OWASP Cheat Sheet Series 
[ OWASP/CheatSheetSeries 
  * 30.1k
  * 4.2k

](https://github.com/OWASP/CheatSheetSeries "Go to repository")
  * [ Introduction  ](https://cheatsheetseries.owasp.org/index.html)
  * [ Index Alphabetical  ](https://cheatsheetseries.owasp.org/Glossary.html)
  * [ Index ASVS  ](https://cheatsheetseries.owasp.org/IndexASVS.html)
  * [ Index MASVS  ](https://cheatsheetseries.owasp.org/IndexMASVS.html)
  * [ Index Proactive Controls  ](https://cheatsheetseries.owasp.org/IndexProactiveControls.html)
  * [ Index Top 10  ](https://cheatsheetseries.owasp.org/IndexTopTen.html)
  * Cheatsheets  Cheatsheets 
    * [ AJAX Security  ](https://cheatsheetseries.owasp.org/cheatsheets/AJAX_Security_Cheat_Sheet.html)
    * [ Abuse Case  ](https://cheatsheetseries.owasp.org/cheatsheets/Abuse_Case_Cheat_Sheet.html)
    * [ Access Control  ](https://cheatsheetseries.owasp.org/cheatsheets/Access_Control_Cheat_Sheet.html)
    * [ Attack Surface Analysis  ](https://cheatsheetseries.owasp.org/cheatsheets/Attack_Surface_Analysis_Cheat_Sheet.html)
    * [ Authentication  ](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)
    * [ Authorization  ](https://cheatsheetseries.owasp.org/cheatsheets/Authorization_Cheat_Sheet.html)
    * [ Authorization Testing Automation  ](https://cheatsheetseries.owasp.org/cheatsheets/Authorization_Testing_Automation_Cheat_Sheet.html)
    * [ Automotive Security.md  ](https://cheatsheetseries.owasp.org/cheatsheets/Automotive_Security.html)
    * [ Bean Validation  ](https://cheatsheetseries.owasp.org/cheatsheets/Bean_Validation_Cheat_Sheet.html)
    * [ Browser Extension Vulnerabilities  ](https://cheatsheetseries.owasp.org/cheatsheets/Browser_Extension_Vulnerabilities_Cheat_Sheet.html)
    * [ C-Based Toolchain Hardening  ](https://cheatsheetseries.owasp.org/cheatsheets/C-Based_Toolchain_Hardening_Cheat_Sheet.html)
    * [ CI CD Security  ](https://cheatsheetseries.owasp.org/cheatsheets/CI_CD_Security_Cheat_Sheet.html)
    * [ Choosing and Using Security Questions  ](https://cheatsheetseries.owasp.org/cheatsheets/Choosing_and_Using_Security_Questions_Cheat_Sheet.html)
    * [ Clickjacking Defense  ](https://cheatsheetseries.owasp.org/cheatsheets/Clickjacking_Defense_Cheat_Sheet.html)
    * [ Content Security Policy  ](https://cheatsheetseries.owasp.org/cheatsheets/Content_Security_Policy_Cheat_Sheet.html)
    * [ Cookie Theft Mitigation  ](https://cheatsheetseries.owasp.org/cheatsheets/Cookie_Theft_Mitigation_Cheat_Sheet.html)
    * [ Credential Stuffing Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/Credential_Stuffing_Prevention_Cheat_Sheet.html)
    * [ Cross-Site Request Forgery Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html)
    * [ Cross Site Scripting Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/Cross_Site_Scripting_Prevention_Cheat_Sheet.html)
    * [ Cryptographic Storage  ](https://cheatsheetseries.owasp.org/cheatsheets/Cryptographic_Storage_Cheat_Sheet.html)
    * [ DOM Clobbering Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/DOM_Clobbering_Prevention_Cheat_Sheet.html)
    * [ DOM based XSS Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/DOM_based_XSS_Prevention_Cheat_Sheet.html)
    * [ Database Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Database_Security_Cheat_Sheet.html)
    * [ Denial of Service  ](https://cheatsheetseries.owasp.org/cheatsheets/Denial_of_Service_Cheat_Sheet.html)
    * [ Deserialization  ](https://cheatsheetseries.owasp.org/cheatsheets/Deserialization_Cheat_Sheet.html)
    * [ Django REST Framework  ](https://cheatsheetseries.owasp.org/cheatsheets/Django_REST_Framework_Cheat_Sheet.html)
    * [ Django Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Django_Security_Cheat_Sheet.html)
    * [ Docker Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Docker_Security_Cheat_Sheet.html)
    * [ DotNet Security  ](https://cheatsheetseries.owasp.org/cheatsheets/DotNet_Security_Cheat_Sheet.html)
    * [ Drone Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Drone_Security_Cheat_Sheet.html)
    * [ Error Handling  ](https://cheatsheetseries.owasp.org/cheatsheets/Error_Handling_Cheat_Sheet.html)
    * File Upload  [ File Upload  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html) Table of contents 
      * [ Introduction  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#introduction)
      * [ File Upload Threats  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-upload-threats)
        * [ Malicious Files  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#malicious-files)
        * [ Public File Retrieval  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#public-file-retrieval)
      * [ File Upload Protection  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-upload-protection)
        * [ Extension Validation  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#extension-validation)
          * [ List Allowed Extensions  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#list-allowed-extensions)
          * [ Block Extensions  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#block-extensions)
        * [ Content-Type Validation  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#content-type-validation)
        * [ File Signature Validation  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-signature-validation)
        * [ Filename Safety  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#filename-safety)
        * [ File Content Validation  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-content-validation)
        * [ File Storage Location  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-storage-location)
        * [ User Permissions  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#user-permissions)
        * [ Filesystem Permissions  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#filesystem-permissions)
        * [ Upload and Download Limits  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#upload-and-download-limits)
      * [ Java Code Snippets  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#java-code-snippets)
    * [ Forgot Password  ](https://cheatsheetseries.owasp.org/cheatsheets/Forgot_Password_Cheat_Sheet.html)
    * [ GraphQL  ](https://cheatsheetseries.owasp.org/cheatsheets/GraphQL_Cheat_Sheet.html)
    * [ HTML5 Security  ](https://cheatsheetseries.owasp.org/cheatsheets/HTML5_Security_Cheat_Sheet.html)
    * [ HTTP Headers  ](https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html)
    * [ HTTP Strict Transport Security  ](https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Strict_Transport_Security_Cheat_Sheet.html)
    * [ Infrastructure as Code Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Infrastructure_as_Code_Security_Cheat_Sheet.html)
    * [ Injection Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/Injection_Prevention_Cheat_Sheet.html)
    * [ Injection Prevention in Java  ](https://cheatsheetseries.owasp.org/cheatsheets/Injection_Prevention_in_Java_Cheat_Sheet.html)
    * [ Input Validation  ](https://cheatsheetseries.owasp.org/cheatsheets/Input_Validation_Cheat_Sheet.html)
    * [ Insecure Direct Object Reference Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/Insecure_Direct_Object_Reference_Prevention_Cheat_Sheet.html)
    * [ JAAS  ](https://cheatsheetseries.owasp.org/cheatsheets/JAAS_Cheat_Sheet.html)
    * [ JSON Web Token for Java  ](https://cheatsheetseries.owasp.org/cheatsheets/JSON_Web_Token_for_Java_Cheat_Sheet.html)
    * [ Java Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Java_Security_Cheat_Sheet.html)
    * [ Key Management  ](https://cheatsheetseries.owasp.org/cheatsheets/Key_Management_Cheat_Sheet.html)
    * [ Kubernetes Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Kubernetes_Security_Cheat_Sheet.html)
    * [ LDAP Injection Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/LDAP_Injection_Prevention_Cheat_Sheet.html)
    * [ LLM Prompt Injection Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/LLM_Prompt_Injection_Prevention_Cheat_Sheet.html)
    * [ Laravel  ](https://cheatsheetseries.owasp.org/cheatsheets/Laravel_Cheat_Sheet.html)
    * [ Legacy Application Management  ](https://cheatsheetseries.owasp.org/cheatsheets/Legacy_Application_Management_Cheat_Sheet.html)
    * [ Logging  ](https://cheatsheetseries.owasp.org/cheatsheets/Logging_Cheat_Sheet.html)
    * [ Logging Vocabulary  ](https://cheatsheetseries.owasp.org/cheatsheets/Logging_Vocabulary_Cheat_Sheet.html)
    * [ Mass Assignment  ](https://cheatsheetseries.owasp.org/cheatsheets/Mass_Assignment_Cheat_Sheet.html)
    * [ Microservices Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Microservices_Security_Cheat_Sheet.html)
    * [ Microservices based Security Arch Doc  ](https://cheatsheetseries.owasp.org/cheatsheets/Microservices_based_Security_Arch_Doc_Cheat_Sheet.html)
    * [ Mobile Application Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Mobile_Application_Security_Cheat_Sheet.html)
    * [ Multifactor Authentication  ](https://cheatsheetseries.owasp.org/cheatsheets/Multifactor_Authentication_Cheat_Sheet.html)
    * [ NPM Security  ](https://cheatsheetseries.owasp.org/cheatsheets/NPM_Security_Cheat_Sheet.html)
    * [ Network Segmentation  ](https://cheatsheetseries.owasp.org/cheatsheets/Network_Segmentation_Cheat_Sheet.html)
    * [ NodeJS Docker  ](https://cheatsheetseries.owasp.org/cheatsheets/NodeJS_Docker_Cheat_Sheet.html)
    * [ Nodejs Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Nodejs_Security_Cheat_Sheet.html)
    * [ OAuth2  ](https://cheatsheetseries.owasp.org/cheatsheets/OAuth2_Cheat_Sheet.html)
    * [ OS Command Injection Defense  ](https://cheatsheetseries.owasp.org/cheatsheets/OS_Command_Injection_Defense_Cheat_Sheet.html)
    * [ PHP Configuration  ](https://cheatsheetseries.owasp.org/cheatsheets/PHP_Configuration_Cheat_Sheet.html)
    * [ Password Storage  ](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
    * [ Pinning  ](https://cheatsheetseries.owasp.org/cheatsheets/Pinning_Cheat_Sheet.html)
    * [ Prototype Pollution Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/Prototype_Pollution_Prevention_Cheat_Sheet.html)
    * [ Query Parameterization  ](https://cheatsheetseries.owasp.org/cheatsheets/Query_Parameterization_Cheat_Sheet.html)
    * [ REST Assessment  ](https://cheatsheetseries.owasp.org/cheatsheets/REST_Assessment_Cheat_Sheet.html)
    * [ REST Security  ](https://cheatsheetseries.owasp.org/cheatsheets/REST_Security_Cheat_Sheet.html)
    * [ Ruby on Rails  ](https://cheatsheetseries.owasp.org/cheatsheets/Ruby_on_Rails_Cheat_Sheet.html)
    * [ SAML Security  ](https://cheatsheetseries.owasp.org/cheatsheets/SAML_Security_Cheat_Sheet.html)
    * [ SQL Injection Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/SQL_Injection_Prevention_Cheat_Sheet.html)
    * [ Secrets Management  ](https://cheatsheetseries.owasp.org/cheatsheets/Secrets_Management_Cheat_Sheet.html)
    * [ Secure AI Model Ops  ](https://cheatsheetseries.owasp.org/cheatsheets/Secure_AI_Model_Ops_Cheat_Sheet.html)
    * [ Secure Cloud Architecture  ](https://cheatsheetseries.owasp.org/cheatsheets/Secure_Cloud_Architecture_Cheat_Sheet.html)
    * [ Secure Product Design  ](https://cheatsheetseries.owasp.org/cheatsheets/Secure_Product_Design_Cheat_Sheet.html)
    * [ Securing Cascading Style Sheets  ](https://cheatsheetseries.owasp.org/cheatsheets/Securing_Cascading_Style_Sheets_Cheat_Sheet.html)
    * [ Server Side Request Forgery Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/Server_Side_Request_Forgery_Prevention_Cheat_Sheet.html)
    * [ Session Management  ](https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html)
    * [ Software Supply Chain Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Software_Supply_Chain_Security_Cheat_Sheet.html)
    * [ Symfony  ](https://cheatsheetseries.owasp.org/cheatsheets/Symfony_Cheat_Sheet.html)
    * [ TLS Cipher String  ](https://cheatsheetseries.owasp.org/cheatsheets/TLS_Cipher_String_Cheat_Sheet.html)
    * [ Third Party Javascript Management  ](https://cheatsheetseries.owasp.org/cheatsheets/Third_Party_Javascript_Management_Cheat_Sheet.html)
    * [ Threat Modeling  ](https://cheatsheetseries.owasp.org/cheatsheets/Threat_Modeling_Cheat_Sheet.html)
    * [ Transaction Authorization  ](https://cheatsheetseries.owasp.org/cheatsheets/Transaction_Authorization_Cheat_Sheet.html)
    * [ Transport Layer Protection  ](https://cheatsheetseries.owasp.org/cheatsheets/Transport_Layer_Protection_Cheat_Sheet.html)
    * [ Transport Layer Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Transport_Layer_Security_Cheat_Sheet.html)
    * [ Unvalidated Redirects and Forwards  ](https://cheatsheetseries.owasp.org/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.html)
    * [ User Privacy Protection  ](https://cheatsheetseries.owasp.org/cheatsheets/User_Privacy_Protection_Cheat_Sheet.html)
    * [ Virtual Patching  ](https://cheatsheetseries.owasp.org/cheatsheets/Virtual_Patching_Cheat_Sheet.html)
    * [ Vulnerability Disclosure  ](https://cheatsheetseries.owasp.org/cheatsheets/Vulnerability_Disclosure_Cheat_Sheet.html)
    * [ Vulnerable Dependency Management  ](https://cheatsheetseries.owasp.org/cheatsheets/Vulnerable_Dependency_Management_Cheat_Sheet.html)
    * [ Web Service Security  ](https://cheatsheetseries.owasp.org/cheatsheets/Web_Service_Security_Cheat_Sheet.html)
    * [ XML External Entity Prevention  ](https://cheatsheetseries.owasp.org/cheatsheets/XML_External_Entity_Prevention_Cheat_Sheet.html)
    * [ XML Security  ](https://cheatsheetseries.owasp.org/cheatsheets/XML_Security_Cheat_Sheet.html)
    * [ XSS Filter Evasion  ](https://cheatsheetseries.owasp.org/cheatsheets/XSS_Filter_Evasion_Cheat_Sheet.html)
    * [ XS Leaks  ](https://cheatsheetseries.owasp.org/cheatsheets/XS_Leaks_Cheat_Sheet.html)
    * [ Zero Trust Architecture  ](https://cheatsheetseries.owasp.org/cheatsheets/Zero_Trust_Architecture_Cheat_Sheet.html)
    * [ gRPC Security  ](https://cheatsheetseries.owasp.org/cheatsheets/gRPC_Security_Cheat_Sheet.html)


Table of contents 
  * [ Introduction  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#introduction)
  * [ File Upload Threats  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-upload-threats)
    * [ Malicious Files  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#malicious-files)
    * [ Public File Retrieval  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#public-file-retrieval)
  * [ File Upload Protection  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-upload-protection)
    * [ Extension Validation  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#extension-validation)
      * [ List Allowed Extensions  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#list-allowed-extensions)
      * [ Block Extensions  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#block-extensions)
    * [ Content-Type Validation  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#content-type-validation)
    * [ File Signature Validation  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-signature-validation)
    * [ Filename Safety  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#filename-safety)
    * [ File Content Validation  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-content-validation)
    * [ File Storage Location  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-storage-location)
    * [ User Permissions  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#user-permissions)
    * [ Filesystem Permissions  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#filesystem-permissions)
    * [ Upload and Download Limits  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#upload-and-download-limits)
  * [ Java Code Snippets  ](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#java-code-snippets)


# File Upload Cheat Sheet[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-upload-cheat-sheet "Permanent link")
## Introduction[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#introduction "Permanent link")
File upload is becoming a more and more essential part of any application, where the user is able to upload their photo, their CV, or a video showcasing a project they are working on. The application should be able to fend off bogus and malicious files in a way to keep the application and the users safe.
In short, the following principles should be followed to reach a secure file upload implementation:
  * **List allowed extensions. Only allow safe and critical extensions for business functionality**
    * **Ensure that[input validation](https://cheatsheetseries.owasp.org/cheatsheets/Input_Validation_Cheat_Sheet.html#file-upload-validation) is applied before validating the extensions.**
  * **Validate the file type, don't trust the[Content-Type header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type) as it can be spoofed**
  * **Change the filename to something generated by the application**
  * **Set a filename length limit. Restrict the allowed characters if possible**
  * **Set a file size limit**
  * **Only allow authorized users to upload files**
  * **Store the files on a different server. If that's not possible, store them outside of the webroot**
    * **In the case of public access to the files, use a handler that gets mapped to filenames inside the application (someid - > file.ext)**
  * **Run the file through an antivirus or a sandbox if available to validate that it doesn't contain malicious data**
  * **Run the file through CDR (Content Disarm & Reconstruct) if applicable type (PDF, DOCX, etc...)**
  * **Ensure that any libraries used are securely configured and kept up to date**
  * **Protect the file upload from[CSRF](https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html) attacks**


## File Upload Threats[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-upload-threats "Permanent link")
In order to assess and know exactly what controls to implement, knowing what you're facing is essential to protect your assets. The following sections will hopefully showcase the risks accompanying the file upload functionality.
### Malicious Files[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#malicious-files "Permanent link")
The attacker delivers a file for malicious intent, such as:
  1. Exploit vulnerabilities in the file parser or processing module (_e.g._ [ImageTrick Exploit](https://imagetragick.com/), [XXE](https://owasp.org/www-community/vulnerabilities/XML_External_Entity_%28XXE%29_Processing))
  2. Use the file for phishing (_e.g._ careers form)
  3. Send ZIP bombs, XML bombs (otherwise known as billion laughs attack), or simply huge files in a way to fill the server storage which hinders and damages the server's availability
  4. Overwrite an existing file on the system
  5. Client-side active content (XSS, CSRF, etc.) that could endanger other users if the files are publicly retrievable.


### Public File Retrieval[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#public-file-retrieval "Permanent link")
If the file uploaded is publicly retrievable, additional threats can be addressed:
  1. Public disclosure of other files
  2. Initiate a DoS attack by requesting lots of files. Requests are small, yet responses are much larger
  3. File content that could be deemed as illegal, offensive, or dangerous (_e.g._ personal data, copyrighted data, etc.) which will make you a host for such malicious files.


## File Upload Protection[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-upload-protection "Permanent link")
There is no silver bullet in validating user content. Implementing a defense in depth approach is key to make the upload process harder and more locked down to the needs and requirements for the service. Implementing multiple techniques is key and recommended, as no one technique is enough to secure the service.
### Extension Validation[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#extension-validation "Permanent link")
Ensure that the validation occurs after decoding the filename, and that a proper filter is set in place in order to avoid certain known bypasses, such as the following:
  * Double extensions, _e.g._ `.jpg.php`, where it circumvents easily the regex `\.jpg`
  * Null bytes, _e.g._ `.php%00.jpg`, where `.jpg` gets truncated and `.php` becomes the new extension
  * Generic bad regex that isn't properly tested and well reviewed. Refrain from building your own logic unless you have enough knowledge on this topic.


Refer to the [Input Validation CS](https://cheatsheetseries.owasp.org/cheatsheets/Input_Validation_Cheat_Sheet.html) to properly parse and process the extension.
#### List Allowed Extensions[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#list-allowed-extensions "Permanent link")
Ensure the usage of _business-critical_ extensions only, without allowing any type of _non-required_ extensions. For example if the system requires:
  * image upload, allow one type that is agreed upon to fit the business requirement;
  * cv upload, allow `docx` and `pdf` extensions.


Based on the needs of the application, ensure the **least harmful** and the **lowest risk** file types to be used.
#### Block Extensions[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#block-extensions "Permanent link")
Identify potentially harmful file types and block extensions that you regard harmful to your service.
Please be aware that blocking specific extensions is a weak protection method on its own. The [Unrestricted File Upload vulnerability](https://owasp.org/www-community/vulnerabilities/Unrestricted_File_Upload) article describes how attackers may attempt to bypass such a check.
### Content-Type Validation[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#content-type-validation "Permanent link")
_The Content-Type for uploaded files is provided by the user, and as such cannot be trusted, as it is trivial to spoof. Although it should not be relied upon for security, it provides a quick check to prevent users from unintentionally uploading files with the incorrect type._
Other than defining the extension of the uploaded file, its MIME-type can be checked for a quick protection against simple file upload attacks.
This can be done preferably in an allowlist approach; otherwise, this can be done in a denylist approach.
### File Signature Validation[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-signature-validation "Permanent link")
In conjunction with [content-type validation](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#content-type-validation), validating the file's signature can be checked and verified against the expected file that should be received.
> This should not be used on its own, as bypassing it is pretty common and easy.
### Filename Safety[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#filename-safety "Permanent link")
Filenames can endanger the system in multiple ways, either by using non acceptable characters, or by using special and restricted filenames. For Windows, refer to the following [MSDN guide](https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file?redirectedfrom=MSDN#naming-conventions). For a wider overview on different filesystems and how they treat files, refer to [Wikipedia's Filename page](https://en.wikipedia.org/wiki/Filename).
In order to avoid the above mentioned threat, creating a **random string** as a filename, such as generating a UUID/GUID, is essential. If the filename is required by the business needs, proper input validation should be done for client-side (_e.g._ active content that results in XSS and CSRF attacks) and back-end side (_e.g._ special files overwrite or creation) attack vectors. Filename length limits should be taken into consideration based on the system storing the files, as each system has its own filename length limit. If user filenames are required, consider implementing the following:
  * Implement a maximum length
  * Restrict characters to an allowed subset specifically, such as alphanumeric characters, hyphen, spaces, and periods
    * Consider telling the user what an acceptable filename is.
    * Restrict use of leading periods (hidden files) and sequential periods (directory traversal).
    * Restrict the use of a leading hyphen or spaces to make it safer to use shell scripts to process files.
    * If this is not possible, block-list dangerous characters that could endanger the framework and system that is storing and using the files.


### File Content Validation[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-content-validation "Permanent link")
As mentioned in the [Public File Retrieval](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#public-file-retrieval) section, file content can contain malicious, inappropriate, or illegal data.
Based on the expected type, special file content validation can be applied:
  * For **images** , applying image rewriting techniques destroys any kind of malicious content injected in an image; this could be done through [randomization](https://security.stackexchange.com/a/8625/118367).
  * For **Microsoft documents** , the usage of [Apache POI](https://poi.apache.org/) helps validating the uploaded documents.
  * **ZIP files** are not recommended since they can contain all types of files, and the attack vectors pertaining to them are numerous.


The File Upload service should allow users to report illegal content, and copyright owners to report abuse.
If there are enough resources, manual file review should be conducted in a sandboxed environment before releasing the files to the public.
Adding some automation to the review could be helpful, which is a harsh process and should be well studied before its usage. Some services (_e.g._ Virus Total) provide APIs to scan files against well known malicious file hashes. Some frameworks can check and validate the raw content type and validating it against predefined file types, such as in [ASP.NET Drawing Library](https://docs.microsoft.com/en-us/dotnet/api/system.drawing.imaging.imageformat). Beware of data leakage threats and information gathering by public services.
### File Storage Location[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#file-storage-location "Permanent link")
The location where the files should be stored must be chosen based on security and business requirements. The following points are set by security priority, and are inclusive:
  1. Store the files on a **different host** , which allows for complete segregation of duties between the application serving the user, and the host handling file uploads and their storage.
  2. Store the files **outside the webroot** , where only administrative access is allowed.
  3. Store the files **inside the webroot** , and set them in write permissions only. - If read access is required, setting proper controls is a must (_e.g._ internal IP, authorized user, etc.)


Storing files in a studied manner in databases is one additional technique. This is sometimes used for automatic backup processes, non file-system attacks, and permissions issues. In return, this opens up the door to performance issues (in some cases), storage considerations for the database and its backups, and this opens up the door to SQLi attack. This is advised only when a DBA is on the team and that this process shows to be an improvement on storing them on the file-system.
> Some files are emailed or processed once they are uploaded, and are not stored on the server. It is essential to conduct the security measures discussed in this sheet before doing any actions on them.
### User Permissions[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#user-permissions "Permanent link")
Before any file upload service is accessed, proper validation should occur on two levels for the user uploading a file:
  * Authentication level
    * The user should be a registered user, or an identifiable user, in order to set restrictions and limitations for their upload capabilities
  * Authorization level
    * The user should have appropriate permissions to access or modify the files


### Filesystem Permissions[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#filesystem-permissions "Permanent link")
> Set the files permissions on the principle of least privilege.
Files should be stored in a way that ensures:
  * Allowed system users are the only ones capable of reading the files
  * Required modes only are set for the file
    * If execution is required, scanning the file before running it is required as a security best practice, to ensure that no macros or hidden scripts are available.


### Upload and Download Limits[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#upload-and-download-limits "Permanent link")
The application should set proper size limits for the upload service in order to protect the file storage capacity. If the system is going to extract the files or process them, the file size limit should be considered after file decompression is conducted and by using secure methods to calculate zip files size. For more on this, see how to [Safely extract files from ZipInputStream](https://wiki.sei.cmu.edu/confluence/display/java/IDS04-J.+Safely+extract+files+from+ZipInputStream), Java's input stream to handle ZIP files.
The application should set proper request limits as well for the download service if available to protect the server from DoS attacks.
## Java Code Snippets[¶](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html#java-code-snippets "Permanent link")
[Document Upload Protection](https://github.com/righettod/document-upload-protection) repository written by Dominique for certain document types in Java.
©Copyright 2025 - Cheat Sheets Series Team - This work is licensed under [Creative Commons Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/). 
Made with [ Material for MkDocs ](https://squidfunk.github.io/mkdocs-material/)


## Source Information
- URL: https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html
- Harvested: 2025-08-19T20:33:47.656794
- Profile: deep_research
- Agent: architect
