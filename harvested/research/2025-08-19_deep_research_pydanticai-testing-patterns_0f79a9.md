---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T01:10:47.155475'
profile: deep_research
source: https://ai.pydantic.dev/testing/
topic: PydanticAI-testing-patterns
---

# PydanticAI-testing-patterns

[ Skip to content ](https://ai.pydantic.dev/testing/#unit-testing)
[ ![logo](https://ai.pydantic.dev/img/logo-white.svg) ](https://ai.pydantic.dev/ "Pydantic AI")
Pydantic AI 
Unit testing 
Type to start searching
[ pydantic/pydantic-ai 
  * v0.7.2
  * 11.7k
  * 1.1k

](https://github.com/pydantic/pydantic-ai "Go to repository")
[ ![logo](https://ai.pydantic.dev/img/logo-white.svg) ](https://ai.pydantic.dev/ "Pydantic AI") Pydantic AI 
[ pydantic/pydantic-ai 
  * v0.7.2
  * 11.7k
  * 1.1k

](https://github.com/pydantic/pydantic-ai "Go to repository")
  * [ Introduction  ](https://ai.pydantic.dev/)
  * [ Installation  ](https://ai.pydantic.dev/install/)
  * [ Getting Help  ](https://ai.pydantic.dev/help/)
  * [ Contributing  ](https://ai.pydantic.dev/contributing/)
  * [ Troubleshooting  ](https://ai.pydantic.dev/troubleshooting/)
  * [ Upgrade Guide  ](https://ai.pydantic.dev/changelog/)
  * Documentation  Documentation 
    * [ Agents  ](https://ai.pydantic.dev/agents/)
    * [ Models  ](https://ai.pydantic.dev/models/)
Models 
      * [ OpenAI  ](https://ai.pydantic.dev/models/openai/)
      * [ Anthropic  ](https://ai.pydantic.dev/models/anthropic/)
      * [ Google  ](https://ai.pydantic.dev/models/google/)
      * [ Bedrock  ](https://ai.pydantic.dev/models/bedrock/)
      * [ Cohere  ](https://ai.pydantic.dev/models/cohere/)
      * [ Groq  ](https://ai.pydantic.dev/models/groq/)
      * [ Mistral  ](https://ai.pydantic.dev/models/mistral/)
      * [ Hugging Face  ](https://ai.pydantic.dev/models/huggingface/)
    * [ Dependencies  ](https://ai.pydantic.dev/dependencies/)
    * [ Function Tools  ](https://ai.pydantic.dev/tools/)
    * [ Toolsets  ](https://ai.pydantic.dev/toolsets/)
    * [ Output  ](https://ai.pydantic.dev/output/)
    * [ Messages and chat history  ](https://ai.pydantic.dev/message-history/)
    * Unit testing  [ Unit testing  ](https://ai.pydantic.dev/testing/) Table of contents 
      * [ Unit testing with TestModel  ](https://ai.pydantic.dev/testing/#unit-testing-with-testmodel)
      * [ Unit testing with FunctionModel  ](https://ai.pydantic.dev/testing/#unit-testing-with-functionmodel)
      * [ Overriding model via pytest fixtures  ](https://ai.pydantic.dev/testing/#overriding-model-via-pytest-fixtures)
    * [ Pydantic Logfire Debugging and Monitoring  ](https://ai.pydantic.dev/logfire/)
    * [ Multi-agent Applications  ](https://ai.pydantic.dev/multi-agent-applications/)
    * [ Graphs  ](https://ai.pydantic.dev/graph/)
    * [ Evals  ](https://ai.pydantic.dev/evals/)
    * [ Image, Audio, Video & Document Input  ](https://ai.pydantic.dev/input/)
    * [ Thinking  ](https://ai.pydantic.dev/thinking/)
    * [ Direct Model Requests  ](https://ai.pydantic.dev/direct/)
    * [ Builtin Tools  ](https://ai.pydantic.dev/builtin-tools/)
    * [ Common Tools  ](https://ai.pydantic.dev/common-tools/)
    * [ HTTP Request Retries  ](https://ai.pydantic.dev/retries/)
    * [ Durable Execution with Temporal  ](https://ai.pydantic.dev/temporal/)
    * [ MCP  ](https://ai.pydantic.dev/mcp/)
MCP 
      * [ Client  ](https://ai.pydantic.dev/mcp/client/)
      * [ Server  ](https://ai.pydantic.dev/mcp/server/)
      * [ MCP Run Python  ](https://ai.pydantic.dev/mcp/run-python/)
    * [ A2A  ](https://ai.pydantic.dev/a2a/)
    * [ AG-UI  ](https://ai.pydantic.dev/ag-ui/)
    * [ Command Line Interface (CLI)  ](https://ai.pydantic.dev/cli/)
  * [ Examples  ](https://ai.pydantic.dev/examples/)
Examples 
    * [ Agent User Interaction (AG-UI)  ](https://ai.pydantic.dev/examples/ag-ui/)
    * [ Pydantic Model  ](https://ai.pydantic.dev/examples/pydantic-model/)
    * [ Weather agent  ](https://ai.pydantic.dev/examples/weather-agent/)
    * [ Bank support  ](https://ai.pydantic.dev/examples/bank-support/)
    * [ SQL Generation  ](https://ai.pydantic.dev/examples/sql-gen/)
    * [ Flight booking  ](https://ai.pydantic.dev/examples/flight-booking/)
    * [ RAG  ](https://ai.pydantic.dev/examples/rag/)
    * [ Stream markdown  ](https://ai.pydantic.dev/examples/stream-markdown/)
    * [ Stream whales  ](https://ai.pydantic.dev/examples/stream-whales/)
    * [ Chat App with FastAPI  ](https://ai.pydantic.dev/examples/chat-app/)
    * [ Question Graph  ](https://ai.pydantic.dev/examples/question-graph/)
    * [ Slack Lead Qualifier with Modal  ](https://ai.pydantic.dev/examples/slack-lead-qualifier/)
    * [ Data Analyst  ](https://ai.pydantic.dev/examples/data-analyst/)
  * API Reference  API Reference 
    * [ pydantic_ai.ag_ui  ](https://ai.pydantic.dev/api/ag_ui/)
    * [ pydantic_ai.agent  ](https://ai.pydantic.dev/api/agent/)
    * [ pydantic_ai.tools  ](https://ai.pydantic.dev/api/tools/)
    * [ pydantic_ai.toolsets  ](https://ai.pydantic.dev/api/toolsets/)
    * [ pydantic_ai.builtin_tools  ](https://ai.pydantic.dev/api/builtin_tools/)
    * [ pydantic_ai.common_tools  ](https://ai.pydantic.dev/api/common_tools/)
    * [ pydantic_ai.durable_exec  ](https://ai.pydantic.dev/api/durable_exec/)
    * [ pydantic_ai.output  ](https://ai.pydantic.dev/api/output/)
    * [ pydantic_ai.result  ](https://ai.pydantic.dev/api/result/)
    * [ pydantic_ai.messages  ](https://ai.pydantic.dev/api/messages/)
    * [ pydantic_ai.exceptions  ](https://ai.pydantic.dev/api/exceptions/)
    * [ pydantic_ai.settings  ](https://ai.pydantic.dev/api/settings/)
    * [ pydantic_ai.usage  ](https://ai.pydantic.dev/api/usage/)
    * [ pydantic_ai.mcp  ](https://ai.pydantic.dev/api/mcp/)
    * [ pydantic_ai.format_prompt  ](https://ai.pydantic.dev/api/format_prompt/)
    * [ pydantic_ai.direct  ](https://ai.pydantic.dev/api/direct/)
    * [ pydantic_ai.ext  ](https://ai.pydantic.dev/api/ext/)
    * [ pydantic_ai.models  ](https://ai.pydantic.dev/api/models/base/)
    * [ pydantic_ai.models.openai  ](https://ai.pydantic.dev/api/models/openai/)
    * [ pydantic_ai.models.anthropic  ](https://ai.pydantic.dev/api/models/anthropic/)
    * [ pydantic_ai.models.bedrock  ](https://ai.pydantic.dev/api/models/bedrock/)
    * [ pydantic_ai.models.cohere  ](https://ai.pydantic.dev/api/models/cohere/)
    * [ pydantic_ai.models.google  ](https://ai.pydantic.dev/api/models/google/)
    * [ pydantic_ai.models.groq  ](https://ai.pydantic.dev/api/models/groq/)
    * [ pydantic_ai.models.huggingface  ](https://ai.pydantic.dev/api/models/huggingface/)
    * [ pydantic_ai.models.instrumented  ](https://ai.pydantic.dev/api/models/instrumented/)
    * [ pydantic_ai.models.mistral  ](https://ai.pydantic.dev/api/models/mistral/)
    * [ pydantic_ai.models.test  ](https://ai.pydantic.dev/api/models/test/)
    * [ pydantic_ai.models.function  ](https://ai.pydantic.dev/api/models/function/)
    * [ pydantic_ai.models.fallback  ](https://ai.pydantic.dev/api/models/fallback/)
    * [ pydantic_ai.models.wrapper  ](https://ai.pydantic.dev/api/models/wrapper/)
    * [ pydantic_ai.models.mcp_sampling  ](https://ai.pydantic.dev/api/models/mcp-sampling/)
    * [ pydantic_ai.profiles  ](https://ai.pydantic.dev/api/profiles/)
    * [ pydantic_ai.providers  ](https://ai.pydantic.dev/api/providers/)
    * [ pydantic_ai.retries  ](https://ai.pydantic.dev/api/retries/)
    * [ pydantic_graph  ](https://ai.pydantic.dev/api/pydantic_graph/graph/)
    * [ pydantic_graph.nodes  ](https://ai.pydantic.dev/api/pydantic_graph/nodes/)
    * [ pydantic_graph.persistence  ](https://ai.pydantic.dev/api/pydantic_graph/persistence/)
    * [ pydantic_graph.mermaid  ](https://ai.pydantic.dev/api/pydantic_graph/mermaid/)
    * [ pydantic_graph.exceptions  ](https://ai.pydantic.dev/api/pydantic_graph/exceptions/)
    * [ pydantic_evals.dataset  ](https://ai.pydantic.dev/api/pydantic_evals/dataset/)
    * [ pydantic_evals.evaluators  ](https://ai.pydantic.dev/api/pydantic_evals/evaluators/)
    * [ pydantic_evals.reporting  ](https://ai.pydantic.dev/api/pydantic_evals/reporting/)
    * [ pydantic_evals.otel  ](https://ai.pydantic.dev/api/pydantic_evals/otel/)
    * [ pydantic_evals.generation  ](https://ai.pydantic.dev/api/pydantic_evals/generation/)
    * [ fasta2a  ](https://ai.pydantic.dev/api/fasta2a/)


Table of contents 
  * [ Unit testing with TestModel  ](https://ai.pydantic.dev/testing/#unit-testing-with-testmodel)
  * [ Unit testing with FunctionModel  ](https://ai.pydantic.dev/testing/#unit-testing-with-functionmodel)
  * [ Overriding model via pytest fixtures  ](https://ai.pydantic.dev/testing/#overriding-model-via-pytest-fixtures)


# Unit testing
Writing unit tests for Pydantic AI code is just like unit tests for any other Python code.
Because for the most part they're nothing new, we have pretty well established tools and patterns for writing and running these kinds of tests.
Unless you're really sure you know better, you'll probably want to follow roughly this strategy:
  * Use [`pytest`](https://docs.pytest.org/en/stable/) as your test harness
  * If you find yourself typing out long assertions, use [inline-snapshot](https://15r10nk.github.io/inline-snapshot/latest/)
  * Similarly, [dirty-equals](https://dirty-equals.helpmanual.io/latest/) can be useful for comparing large data structures
  * Use [`TestModel`](https://ai.pydantic.dev/api/models/test/#pydantic_ai.models.test.TestModel) or [`FunctionModel`](https://ai.pydantic.dev/api/models/function/#pydantic_ai.models.function.FunctionModel) in place of your actual model to avoid the usage, latency and variability of real LLM calls
  * Use [`Agent.override`](https://ai.pydantic.dev/api/agent/#pydantic_ai.agent.Agent.override) to replace an agent's model, dependencies, or toolsets inside your application logic
  * Set [`ALLOW_MODEL_REQUESTS=False`](https://ai.pydantic.dev/api/models/base/#pydantic_ai.models.ALLOW_MODEL_REQUESTS) globally to block any requests from being made to non-test models accidentally


### Unit testing with `TestModel`
The simplest and fastest way to exercise most of your application code is using [`TestModel`](https://ai.pydantic.dev/api/models/test/#pydantic_ai.models.test.TestModel), this will (by default) call all tools in the agent, then return either plain text or a structured response depending on the return type of the agent.
`TestModel` is not magic
The "clever" (but not too clever) part of `TestModel` is that it will attempt to generate valid structured data for [function tools](https://ai.pydantic.dev/tools/) and [output types](https://ai.pydantic.dev/output/#structured-output) based on the schema of the registered tools.
There's no ML or AI in `TestModel`, it's just plain old procedural Python code that tries to generate data that satisfies the JSON schema of a tool.
The resulting data won't look pretty or relevant, but it should pass Pydantic's validation in most cases. If you want something more sophisticated, use [`FunctionModel`](https://ai.pydantic.dev/api/models/function/#pydantic_ai.models.function.FunctionModel) and write your own data generation logic.
Let's write unit tests for the following application code:
weather_app.py```
importasyncio
fromdatetimeimport date
frompydantic_aiimport Agent, RunContext
fromfake_databaseimport DatabaseConn [](https://ai.pydantic.dev/testing/#__code_0_annotation_1)
fromweather_serviceimport WeatherService [](https://ai.pydantic.dev/testing/#__code_0_annotation_2)
weather_agent = Agent(
  'openai:gpt-4o',
  deps_type=WeatherService,
  system_prompt='Providing a weather forecast at the locations the user provides.',
)

@weather_agent.tool
defweather_forecast(
  ctx: RunContext[WeatherService], location: str, forecast_date: date
) -> str:
  if forecast_date < date.today(): [](https://ai.pydantic.dev/testing/#__code_0_annotation_3)
    return ctx.deps.get_historic_weather(location, forecast_date)
  else:
    return ctx.deps.get_forecast(location, forecast_date)

async defrun_weather_forecast( [](https://ai.pydantic.dev/testing/#__code_0_annotation_4)
  user_prompts: list[tuple[str, int]], conn: DatabaseConn
):
"""Run weather forecast for a list of user prompts and save."""
  async with WeatherService() as weather_service:
    async defrun_forecast(prompt: str, user_id: int):
      result = await weather_agent.run(prompt, deps=weather_service)
      await conn.store_forecast(user_id, result.output)
    # run all prompts in parallel
    await asyncio.gather(
      *(run_forecast(prompt, user_id) for (prompt, user_id) in user_prompts)
    )

```

Here we have a function that takes a list of `(user_prompt, user_id)` tuples, gets a weather forecast for each prompt, and stores the result in the database.
**We want to test this code without having to mock certain objects or modify our code so we can pass test objects in.**
Here's how we would write tests using [`TestModel`](https://ai.pydantic.dev/api/models/test/#pydantic_ai.models.test.TestModel):
test_weather_app.py```
fromdatetimeimport timezone
importpytest
fromdirty_equalsimport IsNow, IsStr
frompydantic_aiimport models, capture_run_messages
frompydantic_ai.models.testimport TestModel
frompydantic_ai.messagesimport (
  ModelResponse,
  SystemPromptPart,
  TextPart,
  ToolCallPart,
  ToolReturnPart,
  UserPromptPart,
  ModelRequest,
)
frompydantic_ai.usageimport Usage
fromfake_databaseimport DatabaseConn
fromweather_appimport run_weather_forecast, weather_agent
pytestmark = pytest.mark.anyio [](https://ai.pydantic.dev/testing/#__code_1_annotation_1)
models.ALLOW_MODEL_REQUESTS = False [](https://ai.pydantic.dev/testing/#__code_1_annotation_2)

async deftest_forecast():
  conn = DatabaseConn()
  user_id = 1
  with capture_run_messages() as messages:
    with weather_agent.override(model=TestModel()): [](https://ai.pydantic.dev/testing/#__code_1_annotation_3)
      prompt = 'What will the weather be like in London on 2024-11-28?'
      await run_weather_forecast([(prompt, user_id)], conn) [](https://ai.pydantic.dev/testing/#__code_1_annotation_4)
  forecast = await conn.get_forecast(user_id)
  assert forecast == '{"weather_forecast":"Sunny with a chance of rain"}' [](https://ai.pydantic.dev/testing/#__code_1_annotation_5)
  assert messages == [ [](https://ai.pydantic.dev/testing/#__code_1_annotation_6)
    ModelRequest(
      parts=[
        SystemPromptPart(
          content='Providing a weather forecast at the locations the user provides.',
          timestamp=IsNow(tz=timezone.utc),
        ),
        UserPromptPart(
          content='What will the weather be like in London on 2024-11-28?',
          timestamp=IsNow(tz=timezone.utc), [](https://ai.pydantic.dev/testing/#__code_1_annotation_7)
        ),
      ]
    ),
    ModelResponse(
      parts=[
        ToolCallPart(
          tool_name='weather_forecast',
          args={
            'location': 'a',
            'forecast_date': '2024-01-01', [](https://ai.pydantic.dev/testing/#__code_1_annotation_8)
          },
          tool_call_id=IsStr(),
        )
      ],
      usage=Usage(
        requests=1,
        request_tokens=71,
        response_tokens=7,
        total_tokens=78,
        details=None,
      ),
      model_name='test',
      timestamp=IsNow(tz=timezone.utc),
    ),
    ModelRequest(
      parts=[
        ToolReturnPart(
          tool_name='weather_forecast',
          content='Sunny with a chance of rain',
          tool_call_id=IsStr(),
          timestamp=IsNow(tz=timezone.utc),
        ),
      ],
    ),
    ModelResponse(
      parts=[
        TextPart(
          content='{"weather_forecast":"Sunny with a chance of rain"}',
        )
      ],
      usage=Usage(
        requests=1,
        request_tokens=77,
        response_tokens=16,
        total_tokens=93,
        details=None,
      ),
      model_name='test',
      timestamp=IsNow(tz=timezone.utc),
    ),
  ]

```

### Unit testing with `FunctionModel`
The above tests are a great start, but careful readers will notice that the `WeatherService.get_forecast` is never called since `TestModel` calls `weather_forecast` with a date in the past.
To fully exercise `weather_forecast`, we need to use [`FunctionModel`](https://ai.pydantic.dev/api/models/function/#pydantic_ai.models.function.FunctionModel) to customise how the tools is called.
Here's an example of using `FunctionModel` to test the `weather_forecast` tool with custom inputs
test_weather_app2.py```
importre
importpytest
frompydantic_aiimport models
frompydantic_ai.messagesimport (
  ModelMessage,
  ModelResponse,
  TextPart,
  ToolCallPart,
)
frompydantic_ai.models.functionimport AgentInfo, FunctionModel
fromfake_databaseimport DatabaseConn
fromweather_appimport run_weather_forecast, weather_agent
pytestmark = pytest.mark.anyio
models.ALLOW_MODEL_REQUESTS = False

defcall_weather_forecast( [](https://ai.pydantic.dev/testing/#__code_2_annotation_1)
  messages: list[ModelMessage], info: AgentInfo
) -> ModelResponse:
  if len(messages) == 1:
    # first call, call the weather forecast tool
    user_prompt = messages[0].parts[-1]
    m = re.search(r'\d{4}-\d{2}-\d{2}', user_prompt.content)
    assert m is not None
    args = {'location': 'London', 'forecast_date': m.group()} [](https://ai.pydantic.dev/testing/#__code_2_annotation_2)
    return ModelResponse(parts=[ToolCallPart('weather_forecast', args)])
  else:
    # second call, return the forecast
    msg = messages[-1].parts[0]
    assert msg.part_kind == 'tool-return'
    return ModelResponse(parts=[TextPart(f'The forecast is: {msg.content}')])

async deftest_forecast_future():
  conn = DatabaseConn()
  user_id = 1
  with weather_agent.override(model=FunctionModel(call_weather_forecast)): [](https://ai.pydantic.dev/testing/#__code_2_annotation_3)
    prompt = 'What will the weather be like in London on 2032-01-01?'
    await run_weather_forecast([(prompt, user_id)], conn)
  forecast = await conn.get_forecast(user_id)
  assert forecast == 'The forecast is: Rainy with a chance of sun'

```

### Overriding model via pytest fixtures
If you're writing lots of tests that all require model to be overridden, you can use [pytest fixtures](https://docs.pytest.org/en/6.2.x/fixture.html) to override the model with [`TestModel`](https://ai.pydantic.dev/api/models/test/#pydantic_ai.models.test.TestModel) or [`FunctionModel`](https://ai.pydantic.dev/api/models/function/#pydantic_ai.models.function.FunctionModel) in a reusable way.
Here's an example of a fixture that overrides the model with `TestModel`:
test_agent.py```
importpytest
fromweather_appimport weather_agent
frompydantic_ai.models.testimport TestModel

@pytest.fixture
defoverride_weather_agent():
  with weather_agent.override(model=TestModel()):
    yield

async deftest_forecast(override_weather_agent: None):
  ...
  # test code here

```

© Pydantic Services Inc. 2024 to present 


## Source Information
- URL: https://ai.pydantic.dev/testing/
- Harvested: 2025-08-19T01:10:47.155475
- Profile: deep_research
- Agent: architect
