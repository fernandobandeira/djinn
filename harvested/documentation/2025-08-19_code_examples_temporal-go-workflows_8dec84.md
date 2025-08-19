---
agent_context: unknown
confidence: 0.95
harvested_at: '2025-08-19T01:20:14.630875'
profile: code_examples
source: https://docs.temporal.io/develop/go/temporal-clients
topic: temporal-go-workflows
---

# temporal-go-workflows

![](https://static.scarf.sh/a.png?x-pxid=6fb132d3-92f6-455f-bf17-eb3d6937bdea)
[Skip to main content](https://docs.temporal.io/develop/go/temporal-clients#__docusaurus_skipToContent_fallback)
[![Temporal logo](https://docs.temporal.io/img/assets/temporal-logo-dark.svg)](https://temporal.io)[Home](https://docs.temporal.io/)[Start learning](https://learn.temporal.io/getting_started/)[Start building](https://docs.temporal.io/develop)[Code Exchange](https://temporal.io/code-exchange)[Temporal Cloud](https://docs.temporal.io/cloud)
Search`K`
  * [Documentation Home](https://docs.temporal.io/)
  * [Quickstarts](https://docs.temporal.io/quickstarts)
  * [Evaluate](https://docs.temporal.io/evaluate/)
    * [Why Temporal](https://docs.temporal.io/evaluate/why-temporal)
    * [Understanding Temporal](https://docs.temporal.io/evaluate/understanding-temporal)
    * [Features](https://docs.temporal.io/evaluate/development-production-features/)
    * [Temporal Cloud](https://docs.temporal.io/cloud/introduction)
    * [Security](https://docs.temporal.io/security)
    * [Use Cases](https://docs.temporal.io/evaluate/use-cases-design-patterns)
  * [Get started](https://learn.temporal.io/getting_started/)
  * [Develop](https://docs.temporal.io/develop/)
    * [Go SDK](https://docs.temporal.io/develop/go/)
      * [Quickstart - Setup](https://docs.temporal.io/develop/go/set-up-your-local-go)
      * [Core application](https://docs.temporal.io/develop/go/core-application)
      * [Temporal Client](https://docs.temporal.io/develop/go/temporal-clients)
      * [Multithreading](https://docs.temporal.io/develop/go/go-sdk-multithreading)
      * [Namespaces](https://docs.temporal.io/develop/go/namespaces)
      * [Testing](https://docs.temporal.io/develop/go/testing-suite)
      * [Failure detection](https://docs.temporal.io/develop/go/failure-detection)
      * [Messages](https://docs.temporal.io/develop/go/message-passing)
      * [Interrupt a Workflow Execution](https://docs.temporal.io/develop/go/cancellation)
      * [Asynchronous Activity Completion](https://docs.temporal.io/develop/go/asynchronous-activity-completion)
      * [Versioning](https://docs.temporal.io/develop/go/versioning)
      * [Observability](https://docs.temporal.io/develop/go/observability)
      * [Debugging](https://docs.temporal.io/develop/go/debugging)
      * [Schedules](https://docs.temporal.io/develop/go/schedules)
      * [Converters and encryption](https://docs.temporal.io/develop/go/converters-and-encryption)
      * [Durable Timers](https://docs.temporal.io/develop/go/timers)
      * [Temporal Nexus](https://docs.temporal.io/develop/go/nexus)
      * [Child Workflows](https://docs.temporal.io/develop/go/child-workflows)
      * [Continue-As-New](https://docs.temporal.io/develop/go/continue-as-new)
      * [Side Effects](https://docs.temporal.io/develop/go/side-effects)
      * [Selectors](https://docs.temporal.io/develop/go/selectors)
      * [Sessions](https://docs.temporal.io/develop/go/sessions)
    * [Java SDK](https://docs.temporal.io/develop/java/)
    * [PHP SDK](https://docs.temporal.io/develop/php/)
    * [Python SDK](https://docs.temporal.io/develop/python/)
    * [TypeScript SDK](https://docs.temporal.io/develop/typescript/)
    * [.NET SDK](https://docs.temporal.io/develop/dotnet/)
    * [Ruby SDK](https://docs.temporal.io/develop/ruby/)
    * [Activity retry simulator](https://docs.temporal.io/develop/activity-retry-simulator)
    * [Worker performance](https://docs.temporal.io/develop/worker-performance)
    * [Safe deployments](https://docs.temporal.io/develop/safe-deployments)
  * [Deploy to production](https://docs.temporal.io/production-deployment)
    * [Temporal Cloud](https://docs.temporal.io/cloud)
    * [Codecs and Encryption](https://docs.temporal.io/production-deployment/data-encryption)
    * [Migrate to Cloud](https://docs.temporal.io/production-deployments/migration)
    * [Self-host](https://docs.temporal.io/self-hosted-guide)
    * [Worker deployments](https://docs.temporal.io/production-deployment/worker-deployments)
  * [CLI (temporal)](https://docs.temporal.io/cli)
  * [References](https://docs.temporal.io/references/)
  * [Troubleshooting](https://docs.temporal.io/troubleshooting/)
  * [Encyclopedia](https://docs.temporal.io/encyclopedia/)
  * [Glossary](https://docs.temporal.io/glossary)


  * [](https://docs.temporal.io/)
  * [Develop](https://docs.temporal.io/develop/)
  * [Go SDK](https://docs.temporal.io/develop/go/)
  * Temporal Client


On this page
# Temporal Client - Go SDK
The pages shows how to do the following:
  * [Connect to a local development Temporal Service](https://docs.temporal.io/develop/go/temporal-clients#connect-to-development-service)
  * [Connect to Temporal Cloud](https://docs.temporal.io/develop/go/temporal-clients#connect-to-temporal-cloud)
  * [Start a Workflow Execution](https://docs.temporal.io/develop/go/temporal-clients#start-workflow-execution)
  * [How to start a Workflow Execution](https://docs.temporal.io/develop/go/temporal-clients#start-workflow-execution)


## Connect to development Temporal Service[​](https://docs.temporal.io/develop/go/temporal-clients#connect-to-development-service "Direct link to Connect to development Temporal Service")
**How to connect to the local Temporal CLI development Temporal Service using the Go SDK**
A [Temporal Client](https://docs.temporal.io/encyclopedia/temporal-sdks#temporal-client) enables you to communicate with the [Temporal Service](https://docs.temporal.io/temporal-service). Communication with a Temporal Service includes, but isn't limited to, the following:
  * Starting Workflow Executions.
  * Sending Signals to Workflow Executions.
  * Sending Queries to Workflow Executions.
  * Getting the results of a Workflow Execution.
  * Providing an Activity Task Token.


caution
A Temporal Client cannot be initialized and used inside a Workflow. However, it is acceptable and common to use a Temporal Client inside an Activity to communicate with a Temporal Service.
When you are running a Temporal Service locally (such as the [Temporal CLI](https://docs.temporal.io/cli/server#start-dev)), the number of connection options you must provide is minimal. Many SDKs default to the local host or IP address and port that Temporalite and [Docker Compose](https://github.com/temporalio/docker-compose) serve (`127.0.0.1:7233`).
Use the [`Dial()`](https://pkg.go.dev/go.temporal.io/sdk/client#Dial) API available in the [`go.temporal.io/sdk/client`](https://pkg.go.dev/go.temporal.io/sdk/client) package to create a [`Client`](https://pkg.go.dev/go.temporal.io/sdk/client#Client).
If you don't provide [`HostPort`](https://pkg.go.dev/go.temporal.io/sdk/internal#ClientOptions), the Client defaults the address and port number to `127.0.0.1:7233`, which is the port of the development Temporal Service.
If you don't set a custom Namespace name in the Namespace field, the client connects to the default Namespace.
[sample-apps/go/yourapp/gateway/main.go](https://github.com/temporalio/documentation/blob/main/sample-apps/go/yourapp/gateway/main.go)
```
package mainimport("context""encoding/json""log""net/http""documentation-samples-go/yourapp""go.temporal.io/sdk/client")funcmain(){// Create a Temporal Client to communicate with the Temporal Service.// A Temporal Client is a heavyweight object that should be created just once per process.	temporalClient, err := client.Dial(client.Options{		HostPort: client.DefaultHostPort,})if err !=nil{		log.Fatalln("Unable to create Temporal Client", err)}defer temporalClient.Close()// ...}
```

## Connect to Temporal Cloud[​](https://docs.temporal.io/develop/go/temporal-clients#connect-to-temporal-cloud "Direct link to Connect to Temporal Cloud")
### How to connect to Temporal Cloud using an API Key[​](https://docs.temporal.io/develop/go/temporal-clients#connect-to-temporal-cloud-api-key "Direct link to How to connect to Temporal Cloud using an API Key")
To use an [API key](https://docs.temporal.io/cloud/api-keys) with the Temporal Go SDK, you will need to provide additional connection options:
  * Your _API Key_ value
  * Your _Namespace and Account id_ combination, which follows the format `<namespace_id>.<account_id>`.
  * The _endpoint_ may vary. The most common endpoint used is the gRPC regional endpoint, which follows the format: `<region>.<cloud_provider>.api.temporal.io:7233`.
  * For Namespaces with High Availability features with API key authentication enabled, use the gRPC Namespace endpoint: `<namespace>.<account>.tmprl.cloud:7233`. This allows automated failover without needing to switch endpoints.


You can find the Namespace and Account ID, as well as the endpoint, on the Namespaces tab:
![The Namespace and Account ID combination on the left, and the regional endpoint on the right](https://docs.temporal.io/assets/images/namespaces-and-regional-endpoints-5d0328eb623fc5e3307226a01a5f35b1.png)
Now, when instantiating a Temporal `client` in your Temporal Go SDK code, provide the API key with the following `clientOptions`:
**v1.33.0+**
```
clientOptions := client.Options{  HostPort:<endpoint>,  Namespace:<namespace_id>.<account_id>,  ConnectionOptions: client.ConnectionOptions{TLS:&tls.Config{}},  Credentials: client.NewAPIKeyStaticCredentials(apiKey),}c, err := client.Dial(clientOptions)
```

To update an API key, use the Go `context` object:
```
// Assuming client Credentials created withvar myKey stringcreds := client.NewAPIKeyDynamicCredentials(func(context.Context)(string,error){return myKey,nil})// Update by replacingmyKey = myKeyUpdated
```

**v1.26.0 to v1.33.0 (exclusive)**
Create an initial connection:
```
clientOptions := client.Options{  HostPort:<endpoint>,  Namespace:<namespace_id>.<account_id>,  ConnectionOptions: client.ConnectionOptions{    TLS:&tls.Config{},    DialOptions:[]grpc.DialOption{      grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req any, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption)error{returninvoker(            metadata.AppendToOutgoingContext(ctx,"temporal-namespace",<namespace_id>.<account_id>),            method,            req,            reply,            cc,            opts...,)},),},},  Credentials: client.NewAPIKeyStaticCredentials(apiKey),}c, err := client.Dial(clientOptions)if err !=nil{ log.Fatalf("error creating temporal client: %v", err)}
```

Update an API key:
```
// Assuming client Credentials created withvar myKey stringcreds := client.NewAPIKeyDynamicCredentials(func(context.Context)(string,error){return myKey,nil})// Just update by replacingmyKey = myKeyUpdated
```

**pre v1.26.0**
Create an initial connection:
```
// Create headers providertype APIKeyProvider struct{  APIKey string  Namespace string}func(a *APIKeyProvider)GetHeaders(context.Context)(map[string]string,error){returnmap[string]string{"Authorization":"Bearer "+ a.APIKey,"temporal-namespace": a.Namespace},nil}// Use headers providerapiKeyProvider :=&APIKeyProvider{APIKey:<APIKey>, Namespace:<namespace_id>.<account_id>}c, err := client.Dial(client.Options{  HostPort:<endpoint>,  Namespace:<namespace_id>.<account_id>,  HeadersProvider: apiKeyProvider,  ConnectionOptions: client.ConnectionOptions{TLS:&tls.Config{}},})
```

Update an API key:
```
apiKeyProvider.APIKey = myKeyUpdated
```

### How to connect to Temporal Cloud using mTLS[​](https://docs.temporal.io/develop/go/temporal-clients#connect-to-temporal-cloud-tls "Direct link to How to connect to Temporal Cloud using mTLS")
When you connect to [Temporal Cloud](https://docs.temporal.io/cloud) with mTLS, you need to provide additional connection and client options that include the following:
  * The [Temporal Cloud Namespace Id](https://docs.temporal.io/cloud/namespaces#temporal-cloud-namespace-id).
  * The [Namespace's gRPC endpoint](https://docs.temporal.io/cloud/namespaces#temporal-cloud-grpc-endpoint). An endpoint listing is available at the [Temporal Cloud Website](https://cloud.temporal.io/namespaces) on each Namespace detail page. The endpoint contains the Namespace Id and port.
  * mTLS CA certificate.
  * mTLS private key.


For more information about managing and generating client certificates for Temporal Cloud, see [How to manage certificates in Temporal Cloud](https://docs.temporal.io/cloud/certificates).
For more information about configuring TLS to secure inter- and intra-network communication for a Temporal Service, see [Temporal Customization Samples](https://github.com/temporalio/samples-server).
To connect to and run Workflows through Temporal Cloud, you need the following:
  * A compatible mTLS CA certificate and mTLS private key that has been added to your Namespace. See [certificate requirements](https://docs.temporal.io/cloud/certificates#certificate-requirements).
  * Your [Temporal Cloud Namespace Id](https://docs.temporal.io/cloud/namespaces#temporal-cloud-namespace-id), which includes your [Temporal Cloud Namespace Name](https://docs.temporal.io/cloud/namespaces#temporal-cloud-namespace-name) and the unique five- or six-digit [Temporal Cloud Account Id](https://docs.temporal.io/cloud/namespaces#temporal-cloud-account-id) that is appended to it. This information can be found in the URL of your Namespace; for example, `https://cloud.temporal.io/namespaces/yournamespace.a2fx6/`. Remember that the Namespace Id must include the Account Id: `yournamespace.a2fx6`.


For more information about managing and generating client certificates for Temporal Cloud, see [How to manage certificates in Temporal Cloud](https://docs.temporal.io/cloud/certificates).
For more information about configuring TLS to secure inter- and intra-network communication for a Temporal Service, see [Temporal Customization Samples](https://github.com/temporalio/samples-server).
[sample-apps/go/cloud/client/main.go](https://github.com/temporalio/documentation/blob/main/sample-apps/go/cloud/client/main.go)
```
package mainimport("context""crypto/tls""encoding/json""log""documentation-samples-go/cloud""go.temporal.io/sdk/client")funcmain(){// Get the key and cert from your env or local machine	clientKeyPath :="./secrets/yourkey.key"	clientCertPath :="./secrets/yourcert.pem"// Specify the host and port of your Temporal Cloud Namespace// Host and port format: namespace.unique_id.tmprl.cloud:port	hostPort :="<yournamespace>.<id>.tmprl.cloud:7233"	namespace :="<yournamespace>.<id>"// Use the crypto/tls package to create a cert object	cert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)if err !=nil{		log.Fatalln("Unable to load cert and key pair.", err)}// Add the cert to the tls certificates in the ConnectionOptions of the Client	temporalClient, err := client.Dial(client.Options{		HostPort: hostPort,		Namespace: namespace,		ConnectionOptions: client.ConnectionOptions{			TLS:&tls.Config{Certificates:[]tls.Certificate{cert}},},})if err !=nil{		log.Fatalln("Unable to connect to Temporal Cloud.", err)}defer temporalClient.Close()// ...}
```

## Start Workflow Execution[​](https://docs.temporal.io/develop/go/temporal-clients#start-workflow-execution "Direct link to Start Workflow Execution")
**How to start a Workflow Execution using the Go SDK**
[Workflow Execution](https://docs.temporal.io/workflow-execution) semantics rely on several parameters—that is, to start a Workflow Execution you must supply a Task Queue that will be used for the Tasks (one that a Worker is polling), the Workflow Type, language-specific contextual data, and Workflow Function parameters.
In the examples below, all Workflow Executions are started using a Temporal Client. To spawn Workflow Executions from within another Workflow Execution, use either the [Child Workflow](https://docs.temporal.io/develop/go/child-workflows) or External Workflow APIs.
See the [Customize Workflow Type](https://docs.temporal.io/develop/go/core-application#customize-workflow-type) section to see how to customize the name of the Workflow Type.
A request to spawn a Workflow Execution causes the Temporal Service to create the first Event ([WorkflowExecutionStarted](https://docs.temporal.io/references/events#workflowexecutionstarted)) in the Workflow Execution Event History. The Temporal Service then creates the first Workflow Task, resulting in the first [WorkflowTaskScheduled](https://docs.temporal.io/references/events#workflowtaskscheduled) Event.
To spawn a [Workflow Execution](https://docs.temporal.io/workflow-execution), use the `ExecuteWorkflow()` method on the Go SDK [`Client`](https://pkg.go.dev/go.temporal.io/sdk/client#Client).
The `ExecuteWorkflow()` API call requires an instance of [`context.Context`](https://pkg.go.dev/context#Context), an instance of [`StartWorkflowOptions`](https://pkg.go.dev/go.temporal.io/sdk/client#StartWorkflowOptions), a Workflow Type name, and all variables to be passed to the Workflow Execution. The `ExecuteWorkflow()` call returns a Future, which can be used to get the result of the Workflow Execution.
```
package mainimport(// ..."go.temporal.io/sdk/client")funcmain(){ temporalClient, err := client.Dial(client.Options{})if err !=nil{// ...}defer temporalClient.Close()// ... workflowOptions := client.StartWorkflowOptions{// ...} workflowRun, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition, param)if err !=nil{// ...}// ...}funcYourWorkflowDefinition(ctx workflow.Context, param YourWorkflowParam)(YourWorkflowResponse,error){// ...}
```

If the invocation process has access to the function directly, then the Workflow Type name parameter can be passed as if the function name were a variable, without quotations.
```
workflowRun, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition, param)
```

If the invocation process does not have direct access to the statically defined Workflow Definition, for example, if the Workflow Definition is in an un-importable package, or it is written in a completely different language, then the Workflow Type can be provided as a `string`.
```
workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions,"YourWorkflowDefinition", param)
```

### Set Workflow Task Queue[​](https://docs.temporal.io/develop/go/temporal-clients#set-task-queue "Direct link to Set Workflow Task Queue")
**How to set a Workflow's Task Queue using the Go SDK**
In most SDKs, the only Workflow Option that must be set is the name of the [Task Queue](https://docs.temporal.io/task-queue).
For any code to execute, a Worker Process must be running that contains a Worker Entity that is polling the same Task Queue name.
Create an instance of [`StartWorkflowOptions`](https://pkg.go.dev/go.temporal.io/sdk@v1.10.0/client#StartWorkflowOptions) from the `go.temporal.io/sdk/client` package, set the `TaskQueue` field, and pass the instance to the `ExecuteWorkflow` call.
  * Type: `string`
  * Default: None, this is a required field to be set by the developer


```
workflowOptions := client.StartWorkflowOptions{// ... TaskQueue:"your-task-queue",// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

You can configure Task Queues that are host-specific, Worker-specific or Workflow-specific to distribute your application load. For more information, refer to [Task Queues Processing Tuning](https://docs.temporal.io/develop/worker-performance#task-queues-processing-tuning) and [Worker Versioning](https://docs.temporal.io/worker-versioning).
### Set custom Workflow Id[​](https://docs.temporal.io/develop/go/temporal-clients#workflow-id "Direct link to Set custom Workflow Id")
**How to set a custom Workflow Id using the Go SDK**
Although it is not required, we recommend providing your own [Workflow Id](https://docs.temporal.io/workflow-execution/workflowid-runid#workflow-id)that maps to a business process or business entity identifier, such as an order identifier or customer identifier.
Create an instance of [`StartWorkflowOptions`](https://pkg.go.dev/go.temporal.io/sdk@v1.10.0/client#StartWorkflowOptions) from the `go.temporal.io/sdk/client` package, set the `ID` field, and pass the instance to the `ExecuteWorkflow` call.
  * Type: `string`
  * Default: System generated UUID


```
workflowOptions := client.StartWorkflowOptions{// ... ID:"Your-Custom-Workflow-Id",// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

### Go StartWorkflowOptions reference[​](https://docs.temporal.io/develop/go/temporal-clients#workflow-options-reference "Direct link to Go StartWorkflowOptions reference")
Create an instance of [`StartWorkflowOptions`](https://pkg.go.dev/go.temporal.io/sdk@v1.10.0/client#StartWorkflowOptions) from the `go.temporal.io/sdk/client` package, and pass the instance to the `ExecuteWorkflow` call.
The following fields are available:
Field| Required| Type  
---|---|---  
[`ID`](https://docs.temporal.io/develop/go/temporal-clients#id)| No| `string`  
[`TaskQueue`](https://docs.temporal.io/develop/go/temporal-clients#taskqueue)| **Yes**| `string`  
[`WorkflowExecutionTimeout`](https://docs.temporal.io/develop/go/temporal-clients#workflowexecutiontimeout)|  No| `time.Duration`  
[`WorkflowRunTimeout`](https://docs.temporal.io/develop/go/temporal-clients#workflowruntimeout)| No| `time.Duration`  
[`WorkflowTaskTimeout`](https://docs.temporal.io/develop/go/temporal-clients#workflowtasktimeout)| No| `time.Duration`  
[`WorkflowIDReusePolicy`](https://docs.temporal.io/develop/go/temporal-clients#workflowidreusepolicy)| No| [`WorkflowIdReusePolicy`](https://pkg.go.dev/go.temporal.io/api/enums/v1#WorkflowIdReusePolicy)  
[`WorkflowExecutionErrorWhenAlreadyStarted`](https://docs.temporal.io/develop/go/temporal-clients#workflowexecutionerrorwhenalreadystarted)| No| `bool`  
[`RetryPolicy`](https://docs.temporal.io/develop/go/temporal-clients#retrypolicy)| No| [`RetryPolicy`](https://pkg.go.dev/go.temporal.io/sdk/temporal#RetryPolicy)  
[`CronSchedule`](https://docs.temporal.io/develop/go/temporal-clients#cronschedule)| No| `string`  
[`Memo`](https://docs.temporal.io/develop/go/temporal-clients#memo)| No| `map[string]interface{}`  
[`SearchAttributes`](https://docs.temporal.io/develop/go/temporal-clients#searchattributes)| No| `map[string]interface{}`  
#### ID[​](https://docs.temporal.io/develop/go/temporal-clients#id "Direct link to ID")
Although it is not required, we recommend providing your own [Workflow Id](https://docs.temporal.io/workflow-execution/workflowid-runid#workflow-id)that maps to a business process or business entity identifier, such as an order identifier or customer identifier.
Create an instance of [StartWorkflowOptions](https://pkg.go.dev/go.temporal.io/sdk@v1.10.0/client#StartWorkflowOptions) from the `go.temporal.io/sdk/client` package, set the `ID` field, and pass the instance to the `ExecuteWorkflow` call.
  * Type: `string`
  * Default: System generated UUID


```
workflowOptions := client.StartWorkflowOptions{// ... ID:"Your-Custom-Workflow-Id",// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

#### TaskQueue[​](https://docs.temporal.io/develop/go/temporal-clients#taskqueue "Direct link to TaskQueue")
Create an instance of [StartWorkflowOptions](https://pkg.go.dev/go.temporal.io/sdk@v1.10.0/client#StartWorkflowOptions) from the `go.temporal.io/sdk/client` package, set the `TaskQueue` field, and pass the instance to the `ExecuteWorkflow` call.
  * Type: `string`
  * Default: None; this is a required field to be set by the developer


```
workflowOptions := client.StartWorkflowOptions{// ... TaskQueue:"your-task-queue",// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

#### WorkflowExecutionTimeout[​](https://docs.temporal.io/develop/go/temporal-clients#workflowexecutiontimeout "Direct link to WorkflowExecutionTimeout")
Create an instance of [StartWorkflowOptions](https://pkg.go.dev/go.temporal.io/sdk/client#StartWorkflowOptions) from the `go.temporal.io/sdk/client` package, set the `WorkflowExecutionTimeout` field, and pass the instance to the `ExecuteWorkflow` call.
  * Type: `time.Duration`
  * Default: Unlimited


```
workflowOptions := client.StartWorkflowOptions{// ... WorkflowExecutionTimeout: time.Hours *24*365*10,// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

#### WorkflowRunTimeout[​](https://docs.temporal.io/develop/go/temporal-clients#workflowruntimeout "Direct link to WorkflowRunTimeout")
Create an instance of [StartWorkflowOptions](https://pkg.go.dev/go.temporal.io/sdk/client#StartWorkflowOptions) from the `go.temporal.io/sdk/client` package, set the `WorkflowRunTimeout` field, and pass the instance to the `ExecuteWorkflow` call.
  * Type: `time.Duration`
  * Default: Same as [`WorkflowExecutionTimeout`](https://docs.temporal.io/develop/go/temporal-clients#workflowexecutiontimeout)


```
workflowOptions := client.StartWorkflowOptions{ WorkflowRunTimeout: time.Hours *24*365*10,// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

#### WorkflowTaskTimeout[​](https://docs.temporal.io/develop/go/temporal-clients#workflowtasktimeout "Direct link to WorkflowTaskTimeout")
Create an instance of [StartWorkflowOptions](https://pkg.go.dev/go.temporal.io/sdk/client#StartWorkflowOptions) from the `go.temporal.io/sdk/client` package, set the `WorkflowTaskTimeout` field, and pass the instance to the `ExecuteWorkflow` call.
  * Type: `time.Duration`
  * Default: `time.Seconds * 10`


```
workflowOptions := client.StartWorkflowOptions{ WorkflowTaskTimeout: time.Second *10,//...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

#### WorkflowIDReusePolicy[​](https://docs.temporal.io/develop/go/temporal-clients#workflowidreusepolicy "Direct link to WorkflowIDReusePolicy")
  * Type: [WorkflowIdReusePolicy](https://pkg.go.dev/go.temporal.io/api/enums/v1#WorkflowIdReusePolicy)
  * Default: `enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE`


Set a value from the `go.temporal.io/api/enums/v1` package.
```
workflowOptions := client.StartWorkflowOptions{ WorkflowIdReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

#### WorkflowExecutionErrorWhenAlreadyStarted[​](https://docs.temporal.io/develop/go/temporal-clients#workflowexecutionerrorwhenalreadystarted "Direct link to WorkflowExecutionErrorWhenAlreadyStarted")
  * Type: `bool`
  * Default: `false`


```
workflowOptions := client.StartWorkflowOptions{ WorkflowExecutionErrorWhenAlreadyStarted:false,// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

#### RetryPolicy[​](https://docs.temporal.io/develop/go/temporal-clients#retrypolicy "Direct link to RetryPolicy")
Create an instance of a [RetryPolicy](https://pkg.go.dev/go.temporal.io/sdk/temporal#RetryPolicy) from the `go.temporal.io/sdk/temporal` package and provide it as the value to the `RetryPolicy` field of the instance of `StartWorkflowOptions`.
  * Type: [RetryPolicy](https://pkg.go.dev/go.temporal.io/sdk/temporal#RetryPolicy)
  * Default: None


```
retrypolicy :=&temporal.RetryPolicy{ InitialInterval:  time.Second, BackoffCoefficient:2.0, MaximumInterval:  time.Second *100,}workflowOptions := client.StartWorkflowOptions{ RetryPolicy: retrypolicy,// ...}workflowRun, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

#### CronSchedule[​](https://docs.temporal.io/develop/go/temporal-clients#cronschedule "Direct link to CronSchedule")
  * Type: `string`
  * Default: None


```
workflowOptions := client.StartWorkflowOptions{ CronSchedule:"15 8 * * *",// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

[Sample](https://github.com/temporalio/samples-go/tree/master/cron)
#### Memo[​](https://docs.temporal.io/develop/go/temporal-clients#memo "Direct link to Memo")
  * Type: `map[string]interface{}`
  * Default: Empty


```
workflowOptions := client.StartWorkflowOptions{ Memo:map[string]interface{}{"description":"Test search attributes workflow",},// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

#### SearchAttributes[​](https://docs.temporal.io/develop/go/temporal-clients#searchattributes "Direct link to SearchAttributes")
**How to set Workflow Execution Search Attributes in Go**
  * Type: `map[string]interface{}`
  * Default: Empty.


These are the corresponding [Search Attribute value types](https://docs.temporal.io/search-attribute#supported-types) in Go:
  * Keyword = string
  * Int = int64
  * Double = float64
  * Bool = bool
  * Datetime = time.Time
  * Text = string


```
searchAttributes :=map[string]interface{}{"CustomIntField":1,"MiscData":"yellow",}workflowOptions := client.StartWorkflowOptions{ SearchAttributes: searchAttributes,// ...}workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition)if err !=nil{// ...}
```

### Get Workflow results[​](https://docs.temporal.io/develop/go/temporal-clients#get-workflow-results "Direct link to Get Workflow results")
**How to get the results of a Workflow Execution using the Go SDK**
If the call to start a Workflow Execution is successful, you will gain access to the Workflow Execution's Run Id.
The Workflow Id, Run Id, and Namespace may be used to uniquely identify a Workflow Execution in the system and get its result.
It's possible to both block progress on the result (synchronous execution) or get the result at some other point in time (asynchronous execution).
In the Temporal Platform, it's also acceptable to use Queries as the preferred method for accessing the state and results of Workflow Executions.
The `ExecuteWorkflow` call returns an instance of [`WorkflowRun`](https://pkg.go.dev/go.temporal.io/sdk/client#WorkflowRun), which is the `workflowRun` variable in the following line.
```
 workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, app.YourWorkflowDefinition, param)if err !=nil{// ...}// ...}
```

The instance of `WorkflowRun` has the following three methods:
  * `GetWorkflowID()`: Returns the Workflow Id of the invoked Workflow Execution.
  * `GetRunID()`: Always returns the Run Id of the initial Run (See [Continue As New](https://docs.temporal.io/develop/go/temporal-clients)) in the series of Runs that make up the full Workflow Execution.
  * `Get`: Takes a pointer as a parameter and populates the associated variable with the Workflow Execution result.


To wait on the result of Workflow Execution in the same process that invoked it, call `Get()` on the instance of `WorkflowRun` that is returned by the `ExecuteWorkflow()` call.
```
 workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowDefinition, param)if err !=nil{// ...}var result YourWorkflowResponse err = workflowRun.Get(context.Background(),&result)if err !=nil{// ...}// ...}
```

However, the result of a Workflow Execution can be obtained from a completely different process. All that is needed is the [Workflow Id](https://docs.temporal.io/develop/go/temporal-clients). (A [Run Id](https://docs.temporal.io/develop/go/temporal-clients) is optional if more than one closed Workflow Execution has the same Workflow Id.) The result of the Workflow Execution is available for as long as the Workflow Execution Event History remains in the system.
Call the `GetWorkflow()` method on an instance of the Go SDK Client and pass it the Workflow Id used to spawn the Workflow Execution. Then call the `Get()` method on the instance of `WorkflowRun` that is returned, passing it a pointer to populate the result.
```
// ... workflowID :="Your-Custom-Workflow-Id" workflowRun := c.GetWorkflow(context.Background, workflowID)var result YourWorkflowResponse err = workflowRun.Get(context.Background(),&result)if err !=nil{// ...}// ...
```

**Get last completion result**
In the case of a [Temporal Cron Job](https://docs.temporal.io/cron-job), you might need to get the result of the previous Workflow Run and use it in the current Workflow Run.
To do this, use the [`HasLastCompletionResult`](https://pkg.go.dev/go.temporal.io/sdk/workflow#HasLastCompletionResult) and [`GetLastCompletionResult`](https://pkg.go.dev/go.temporal.io/sdk/workflow#GetLastCompletionResult) APIs, available from the [`go.temporal.io/sdk/workflow`](https://pkg.go.dev/go.temporal.io/sdk/workflow) package, directly in your Workflow code.
```
type CronResult struct{ Count int}funcYourCronWorkflowDefinition(ctx workflow.Context)(CronResult,error){ count :=1if workflow.HasLastCompletionResult(ctx){var lastResult CronResultif err := workflow.GetLastCompletionResult(ctx,&lastResult); err ==nil{   count = count + lastResult.Count}} newResult := CronResult {  Count: count,}return newResult,nil}
```

This will work even if one of the cron Workflow Runs fails. The next Workflow Run gets the result of the last successfully Completed Workflow Run.
**Tags:**
  * [Temporal Client](https://docs.temporal.io/tags/temporal-client)
  * [Go SDK](https://docs.temporal.io/tags/go-sdk)
  * [Temporal SDKs](https://docs.temporal.io/tags/temporal-sd-ks)
  * [Certificates](https://docs.temporal.io/tags/certificates)


Help us make Temporal better. Contribute to our [documentation](https://github.com/temporalio/documentation).
[PreviousCore application](https://docs.temporal.io/develop/go/core-application)[NextMultithreading](https://docs.temporal.io/develop/go/go-sdk-multithreading)
  * [Connect to development Temporal Service](https://docs.temporal.io/develop/go/temporal-clients#connect-to-development-service)
  * [Connect to Temporal Cloud](https://docs.temporal.io/develop/go/temporal-clients#connect-to-temporal-cloud)
    * [How to connect to Temporal Cloud using an API Key](https://docs.temporal.io/develop/go/temporal-clients#connect-to-temporal-cloud-api-key)
    * [How to connect to Temporal Cloud using mTLS](https://docs.temporal.io/develop/go/temporal-clients#connect-to-temporal-cloud-tls)
  * [Start Workflow Execution](https://docs.temporal.io/develop/go/temporal-clients#start-workflow-execution)
    * [Set Workflow Task Queue](https://docs.temporal.io/develop/go/temporal-clients#set-task-queue)
    * [Set custom Workflow Id](https://docs.temporal.io/develop/go/temporal-clients#workflow-id)
    * [Go StartWorkflowOptions reference](https://docs.temporal.io/develop/go/temporal-clients#workflow-options-reference)
      * [ID](https://docs.temporal.io/develop/go/temporal-clients#id)
      * [TaskQueue](https://docs.temporal.io/develop/go/temporal-clients#taskqueue)
      * [WorkflowExecutionTimeout](https://docs.temporal.io/develop/go/temporal-clients#workflowexecutiontimeout)
      * [WorkflowRunTimeout](https://docs.temporal.io/develop/go/temporal-clients#workflowruntimeout)
      * [WorkflowTaskTimeout](https://docs.temporal.io/develop/go/temporal-clients#workflowtasktimeout)
      * [WorkflowIDReusePolicy](https://docs.temporal.io/develop/go/temporal-clients#workflowidreusepolicy)
      * [WorkflowExecutionErrorWhenAlreadyStarted](https://docs.temporal.io/develop/go/temporal-clients#workflowexecutionerrorwhenalreadystarted)
      * [RetryPolicy](https://docs.temporal.io/develop/go/temporal-clients#retrypolicy)
      * [CronSchedule](https://docs.temporal.io/develop/go/temporal-clients#cronschedule)
      * [Memo](https://docs.temporal.io/develop/go/temporal-clients#memo)
      * [SearchAttributes](https://docs.temporal.io/develop/go/temporal-clients#searchattributes)
    * [Get Workflow results](https://docs.temporal.io/develop/go/temporal-clients#get-workflow-results)


  * [Github](https://github.com/temporalio)
  * [Twitter](https://twitter.com/temporalio)
  * [YouTube](https://www.youtube.com/c/Temporalio)
  * [About the docs](https://github.com/temporalio/documentation/blob/master/README.md)


  * [Temporal Cloud](https://temporal.io/cloud)
  * [Meetups](https://temporal.io/community#events)
  * [Workshops](https://temporal.io/community#workshops)
  * [Support forum](https://community.temporal.io/)
  * [Ask an expert](https://pages.temporal.io/ask-an-expert)


  * [Learn Temporal](https://learn.temporal.io)
  * [Blog](https://temporal.io/blog)
  * [Use cases](https://temporal.io/use-cases)
  * [Newsletter signup](https://pages.temporal.io/newsletter-subscribe)


  * [Security](https://docs.temporal.io/security)
  * [Privacy policy](https://temporal.io/global-privacy-policy)
  * [Terms of service](https://docs.temporal.io/pdf/temporal-tos-2021-07-24.pdf)
  * [We're hiring](https://temporal.io/careers)


[![Temporal logo](https://docs.temporal.io/img/favicon.png)](https://temporal.io)
Copyright © 2025 Temporal Technologies Inc.
Feedback


## Source Information
- URL: https://docs.temporal.io/develop/go/temporal-clients
- Harvested: 2025-08-19T01:20:14.630875
- Profile: code_examples
- Agent: unknown
