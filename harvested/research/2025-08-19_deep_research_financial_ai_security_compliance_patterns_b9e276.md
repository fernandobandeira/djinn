---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T12:02:44.239579'
profile: deep_research
source: https://owasp.org/Top10/A03_2021-Injection/
topic: Financial AI security compliance patterns
---

# Financial AI security compliance patterns

[ Skip to content ](https://owasp.org/Top10/A03_2021-Injection/#a032021-injection)
[ ![logo](https://owasp.org/Top10/assets/OWASP_Logo_Transp.png) ](https://owasp.org/Top10/ "OWASP Top 10:2021")
OWASP Top 10:2021 
A03 Injection 
  * [ en - English ](https://owasp.org/Top10/A03_2021-Injection/)
  * [ ar - ﺎﻠﻋﺮﺒﻳﺓ ](https://owasp.org/Top10/ar/A03_2021-Injection/)
  * [ de - Deutsch ](https://owasp.org/Top10/de/A03_2021-Injection/)
  * [ es - Español ](https://owasp.org/Top10/es/A03_2021-Injection/)
  * [ fr - Français ](https://owasp.org/Top10/fr/A03_2021-Injection/)
  * [ id - Indonesian ](https://owasp.org/Top10/id/A03_2021-Injection/)
  * [ it - Italiano ](https://owasp.org/Top10/it/A03_2021-Injection/)
  * [ ja - 日本語 ](https://owasp.org/Top10/ja/A03_2021-Injection/)
  * [ pt-BR - Português (Brasil) ](https://owasp.org/Top10/pt-BR/A03_2021-Injection/)
  * [ zh-Hant - 简体中文 ](https://owasp.org/Top10/zh-Hant/A03_2021-Injection/)
  * [ zh-TW - 繁體中文 ](https://owasp.org/Top10/zh-TW/A03_2021-Injection/)


Type to start searching
[ OWASP/Top10 
  * 4.7k
  * 909

](https://github.com/OWASP/Top10 "Go to repository")
[ ![logo](https://owasp.org/Top10/assets/OWASP_Logo_Transp.png) ](https://owasp.org/Top10/ "OWASP Top 10:2021") OWASP Top 10:2021 
[ OWASP/Top10 
  * 4.7k
  * 909

](https://github.com/OWASP/Top10 "Go to repository")
  * [ Home  ](https://owasp.org/Top10/)
  * [ Notice  ](https://owasp.org/Top10/0x00_2021-notice/)
  * [ Introduction  ](https://owasp.org/Top10/A00_2021_Introduction/)
  * [ How to use the OWASP Top 10 as a standard  ](https://owasp.org/Top10/A00_2021_How_to_use_the_OWASP_Top_10_as_a_standard/)
  * [ How to start an AppSec program with the OWASP Top 10  ](https://owasp.org/Top10/A00_2021-How_to_start_an_AppSec_program_with_the_OWASP_Top_10/)
  * [ About OWASP  ](https://owasp.org/Top10/0x01_2021-about-owasp/)
  * Top 10:2021 List  Top 10:2021 List 
    * [ A01 Broken Access Control  ](https://owasp.org/Top10/A01_2021-Broken_Access_Control/)
    * [ A02 Cryptographic Failures  ](https://owasp.org/Top10/A02_2021-Cryptographic_Failures/)
    * A03 Injection  [ A03 Injection  ](https://owasp.org/Top10/A03_2021-Injection/) Table of contents 
      * [ Factors  ](https://owasp.org/Top10/A03_2021-Injection/#factors)
      * [ Overview  ](https://owasp.org/Top10/A03_2021-Injection/#overview)
      * [ Description  ](https://owasp.org/Top10/A03_2021-Injection/#description)
      * [ How to Prevent  ](https://owasp.org/Top10/A03_2021-Injection/#how-to-prevent)
      * [ Example Attack Scenarios  ](https://owasp.org/Top10/A03_2021-Injection/#example-attack-scenarios)
      * [ References  ](https://owasp.org/Top10/A03_2021-Injection/#references)
      * [ List of Mapped CWEs  ](https://owasp.org/Top10/A03_2021-Injection/#list-of-mapped-cwes)
    * [ A04 Insecure Design  ](https://owasp.org/Top10/A04_2021-Insecure_Design/)
    * [ A05 Security Misconfiguration  ](https://owasp.org/Top10/A05_2021-Security_Misconfiguration/)
    * [ A06 Vulnerable and Outdated Components  ](https://owasp.org/Top10/A06_2021-Vulnerable_and_Outdated_Components/)
    * [ A07 Identification and Authentication Failures  ](https://owasp.org/Top10/A07_2021-Identification_and_Authentication_Failures/)
    * [ A08 Software and Data Integrity Failures  ](https://owasp.org/Top10/A08_2021-Software_and_Data_Integrity_Failures/)
    * [ A09 Security Logging and Monitoring Failures  ](https://owasp.org/Top10/A09_2021-Security_Logging_and_Monitoring_Failures/)
    * [ A10 Server Side Request Forgery (SSRF)  ](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)
  * [ Next Steps  ](https://owasp.org/Top10/A11_2021-Next_Steps/)


Table of contents 
  * [ Factors  ](https://owasp.org/Top10/A03_2021-Injection/#factors)
  * [ Overview  ](https://owasp.org/Top10/A03_2021-Injection/#overview)
  * [ Description  ](https://owasp.org/Top10/A03_2021-Injection/#description)
  * [ How to Prevent  ](https://owasp.org/Top10/A03_2021-Injection/#how-to-prevent)
  * [ Example Attack Scenarios  ](https://owasp.org/Top10/A03_2021-Injection/#example-attack-scenarios)
  * [ References  ](https://owasp.org/Top10/A03_2021-Injection/#references)
  * [ List of Mapped CWEs  ](https://owasp.org/Top10/A03_2021-Injection/#list-of-mapped-cwes)


# A03:2021 – Injection ![icon](https://owasp.org/Top10/assets/TOP_10_Icons_Final_Injection.png)
## Factors
CWEs Mapped | Max Incidence Rate | Avg Incidence Rate | Avg Weighted Exploit | Avg Weighted Impact | Max Coverage | Avg Coverage | Total Occurrences | Total CVEs  
---|---|---|---|---|---|---|---|---  
33 | 19.09% | 3.37% | 7.25 | 7.15 | 94.04% | 47.90% | 274,228 | 32,078  
## Overview
Injection slides down to the third position. 94% of the applications were tested for some form of injection with a max incidence rate of 19%, an average incidence rate of 3%, and 274k occurrences. Notable Common Weakness Enumerations (CWEs) included are _CWE-79: Cross-site Scripting_ , _CWE-89: SQL Injection_ , and _CWE-73: External Control of File Name or Path_.
## Description
An application is vulnerable to attack when:
  * User-supplied data is not validated, filtered, or sanitized by the application.
  * Dynamic queries or non-parameterized calls without context-aware escaping are used directly in the interpreter.
  * Hostile data is used within object-relational mapping (ORM) search parameters to extract additional, sensitive records.
  * Hostile data is directly used or concatenated. The SQL or command contains the structure and malicious data in dynamic queries, commands, or stored procedures.


Some of the more common injections are SQL, NoSQL, OS command, Object Relational Mapping (ORM), LDAP, and Expression Language (EL) or Object Graph Navigation Library (OGNL) injection. The concept is identical among all interpreters. Source code review is the best method of detecting if applications are vulnerable to injections. Automated testing of all parameters, headers, URL, cookies, JSON, SOAP, and XML data inputs is strongly encouraged. Organizations can include static (SAST), dynamic (DAST), and interactive (IAST) application security testing tools into the CI/CD pipeline to identify introduced injection flaws before production deployment.
## How to Prevent
Preventing injection requires keeping data separate from commands and queries:
  * The preferred option is to use a safe API, which avoids using the interpreter entirely, provides a parameterized interface, or migrates to Object Relational Mapping Tools (ORMs). **Note:** Even when parameterized, stored procedures can still introduce SQL injection if PL/SQL or T-SQL concatenates queries and data or executes hostile data with EXECUTE IMMEDIATE or exec().
  * Use positive server-side input validation. This is not a complete defense as many applications require special characters, such as text areas or APIs for mobile applications.
  * For any residual dynamic queries, escape special characters using the specific escape syntax for that interpreter. **Note:** SQL structures such as table names, column names, and so on cannot be escaped, and thus user-supplied structure names are dangerous. This is a common issue in report-writing software.


## Example Attack Scenarios
**Scenario #1:** An application uses untrusted data in the construction of the following vulnerable SQL call:
```
String query = "SELECT \* FROM accounts WHERE custID='" + request.getParameter("id") + "'";

```

**Scenario #2:** Similarly, an application’s blind trust in frameworks may result in queries that are still vulnerable, (e.g., Hibernate Query Language (HQL)):
```
 Query HQLQuery = session.createQuery("FROM accounts WHERE custID='" + request.getParameter("id") + "'");

```

In both cases, the attacker modifies the ‘id’ parameter value in their browser to send: ' UNION SLEEP(10);--. For example:
```
 http://example.com/app/accountView?id=' UNION SELECT SLEEP(10);--

```

This changes the meaning of both queries to return all the records from the accounts table. More dangerous attacks could modify or delete data or even invoke stored procedures.
## References
  * [OWASP Proactive Controls: Secure Database Access](https://owasp.org/www-project-proactive-controls/v3/en/c3-secure-database)
  * [OWASP ASVS: V5 Input Validation and Encoding](https://owasp.org/www-project-application-security-verification-standard)
  * [OWASP Testing Guide: SQL Injection,](https://owasp.org/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/07-Input_Validation_Testing/05-Testing_for_SQL_Injection) [Command Injection](https://owasp.org/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/07-Input_Validation_Testing/12-Testing_for_Command_Injection), and [ORM Injection](https://owasp.org/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/07-Input_Validation_Testing/05.7-Testing_for_ORM_Injection)
  * [OWASP Cheat Sheet: Injection Prevention](https://cheatsheetseries.owasp.org/cheatsheets/Injection_Prevention_Cheat_Sheet.html)
  * [OWASP Cheat Sheet: SQL Injection Prevention](https://cheatsheetseries.owasp.org/cheatsheets/SQL_Injection_Prevention_Cheat_Sheet.html)
  * [OWASP Cheat Sheet: Injection Prevention in Java](https://cheatsheetseries.owasp.org/cheatsheets/Injection_Prevention_Cheat_Sheet_in_Java.html)
  * [OWASP Cheat Sheet: Query Parameterization](https://cheatsheetseries.owasp.org/cheatsheets/Query_Parameterization_Cheat_Sheet.html)
  * [OWASP Automated Threats to Web Applications – OAT-014](https://owasp.org/www-project-automated-threats-to-web-applications/)
  * [PortSwigger: Server-side template injection](https://portswigger.net/kb/issues/00101080_serversidetemplateinjection)


## List of Mapped CWEs
[CWE-20 Improper Input Validation](https://cwe.mitre.org/data/definitions/20.html)
[CWE-74 Improper Neutralization of Special Elements in Output Used by a Downstream Component ('Injection')](https://cwe.mitre.org/data/definitions/74.html)
[CWE-75 Failure to Sanitize Special Elements into a Different Plane (Special Element Injection)](https://cwe.mitre.org/data/definitions/75.html)
[CWE-77 Improper Neutralization of Special Elements used in a Command ('Command Injection')](https://cwe.mitre.org/data/definitions/77.html)
[CWE-78 Improper Neutralization of Special Elements used in an OS Command ('OS Command Injection')](https://cwe.mitre.org/data/definitions/78.html)
[CWE-79 Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')](https://cwe.mitre.org/data/definitions/79.html)
[CWE-80 Improper Neutralization of Script-Related HTML Tags in a Web Page (Basic XSS)](https://cwe.mitre.org/data/definitions/80.html)
[CWE-83 Improper Neutralization of Script in Attributes in a Web Page](https://cwe.mitre.org/data/definitions/83.html)
[CWE-87 Improper Neutralization of Alternate XSS Syntax](https://cwe.mitre.org/data/definitions/87.html)
[CWE-88 Improper Neutralization of Argument Delimiters in a Command ('Argument Injection')](https://cwe.mitre.org/data/definitions/88.html)
[CWE-89 Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')](https://cwe.mitre.org/data/definitions/89.html)
[CWE-90 Improper Neutralization of Special Elements used in an LDAP Query ('LDAP Injection')](https://cwe.mitre.org/data/definitions/90.html)
[CWE-91 XML Injection (aka Blind XPath Injection)](https://cwe.mitre.org/data/definitions/91.html)
[CWE-93 Improper Neutralization of CRLF Sequences ('CRLF Injection')](https://cwe.mitre.org/data/definitions/93.html)
[CWE-94 Improper Control of Generation of Code ('Code Injection')](https://cwe.mitre.org/data/definitions/94.html)
[CWE-95 Improper Neutralization of Directives in Dynamically Evaluated Code ('Eval Injection')](https://cwe.mitre.org/data/definitions/95.html)
[CWE-96 Improper Neutralization of Directives in Statically Saved Code ('Static Code Injection')](https://cwe.mitre.org/data/definitions/96.html)
[CWE-97 Improper Neutralization of Server-Side Includes (SSI) Within a Web Page](https://cwe.mitre.org/data/definitions/97.html)
[CWE-98 Improper Control of Filename for Include/Require Statement in PHP Program ('PHP Remote File Inclusion')](https://cwe.mitre.org/data/definitions/98.html)
[CWE-99 Improper Control of Resource Identifiers ('Resource Injection')](https://cwe.mitre.org/data/definitions/99.html)
[CWE-100 Deprecated: Was catch-all for input validation issues](https://cwe.mitre.org/data/definitions/100.html)
[CWE-113 Improper Neutralization of CRLF Sequences in HTTP Headers ('HTTP Response Splitting')](https://cwe.mitre.org/data/definitions/113.html)
[CWE-116 Improper Encoding or Escaping of Output](https://cwe.mitre.org/data/definitions/116.html)
[CWE-138 Improper Neutralization of Special Elements](https://cwe.mitre.org/data/definitions/138.html)
[CWE-184 Incomplete List of Disallowed Inputs](https://cwe.mitre.org/data/definitions/184.html)
[CWE-470 Use of Externally-Controlled Input to Select Classes or Code ('Unsafe Reflection')](https://cwe.mitre.org/data/definitions/470.html)
[CWE-471 Modification of Assumed-Immutable Data (MAID)](https://cwe.mitre.org/data/definitions/471.html)
[CWE-564 SQL Injection: Hibernate](https://cwe.mitre.org/data/definitions/564.html)
[CWE-610 Externally Controlled Reference to a Resource in Another Sphere](https://cwe.mitre.org/data/definitions/610.html)
[CWE-643 Improper Neutralization of Data within XPath Expressions ('XPath Injection')](https://cwe.mitre.org/data/definitions/643.html)
[CWE-644 Improper Neutralization of HTTP Headers for Scripting Syntax](https://cwe.mitre.org/data/definitions/644.html)
[CWE-652 Improper Neutralization of Data within XQuery Expressions ('XQuery Injection')](https://cwe.mitre.org/data/definitions/652.html)
[CWE-917 Improper Neutralization of Special Elements used in an Expression Language Statement ('Expression Language Injection')](https://cwe.mitre.org/data/definitions/917.html)
© Copyright 2021-2025 - OWASP Top 10 Team - This work is licensed under a [Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/deed.en_US). 
Made with [ Material for MkDocs ](https://squidfunk.github.io/mkdocs-material/)


## Source Information
- URL: https://owasp.org/Top10/A03_2021-Injection/
- Harvested: 2025-08-19T12:02:44.239579
- Profile: deep_research
- Agent: architect
