---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:15:37.005697'
profile: deep_research
source: https://stryker-mutator.io/docs/
topic: Mutation Testing Patterns and Best Practices
---

# Mutation Testing Patterns and Best Practices

[Skip to main content](https://stryker-mutator.io/docs/#__docusaurus_skipToContent_fallback)
📣 Discover the ["Who's testing the tests? Mutation testing with StrykerJS"](https://fosdem.org/2024/schedule/event/fosdem-2024-1683-who-s-testing-the-tests-mutation-testing-with-strykerjs/) talk on FOSDEM 2024.
[![Stryker Mutator Logo](https://stryker-mutator.io/images/stryker.svg)**Stryker Mutator**](https://stryker-mutator.io/)[Blog](https://stryker-mutator.io/blog/)[For JavaScript](https://stryker-mutator.io/docs/stryker-js/introduction/)[For C#](https://stryker-mutator.io/docs/stryker-net/introduction/)[For Scala](https://stryker-mutator.io/docs/stryker4s/getting-started/)[An example](https://stryker-mutator.io/docs/General/example/)[Playground](https://stryker-mutator.io/stryker-playground/)
[Dashboard](https://dashboard.stryker-mutator.io)
[](https://stryker-mutator.io/docs/)
  * [StrykerJS (JS & TS)](https://github.com/stryker-mutator/stryker-js)
  * [Stryker.NET (C#)](https://github.com/stryker-mutator/stryker-net)
  * [Stryker4s (Scala)](https://github.com/stryker-mutator/stryker4s)
  * [Mutation Testing Elements](https://github.com/stryker-mutator/mutation-testing-elements)
  * [Stryker Dashboard](https://github.com/stryker-mutator/stryker-dashboard)
  * [This website](https://github.com/stryker-mutator/stryker-mutator.github.io)


Search`K`
  * [General](https://stryker-mutator.io/docs/)
    * [Introduction](https://stryker-mutator.io/docs/)
    * [An example](https://stryker-mutator.io/docs/General/example/)
    * [Stryker dashboard](https://stryker-mutator.io/docs/General/dashboard/)
    * [FAQ](https://stryker-mutator.io/docs/General/faq/)
  * [Mutation Testing](https://stryker-mutator.io/docs/)
  * [StrykerJS](https://stryker-mutator.io/docs/)
  * [Stryker.NET](https://stryker-mutator.io/docs/)
  * [Stryker4s](https://stryker-mutator.io/docs/)


  * [](https://stryker-mutator.io/)
  * General
  * Introduction


On this page
# What is mutation testing?
**TL; DR: Mutation testing introduces changes to your code, then runs your unit tests against the changed code. It is expected that your unit tests will now fail. If they don't fail, it might indicate your tests do not sufficiently cover the code.**
Bugs, or _mutants_ , are automatically inserted into your production code. Your tests are run for each mutant. If your tests _fail_ then the mutant is _killed_. If your tests passed, the mutant _survived_. The higher the percentage of mutants killed, the more _effective_ your tests are.
It's that simple.
Are you still confused? Why not take a look at our [example page](https://stryker-mutator.io/docs/General/example/) and try it out yourself?
## But wait, what about code coverage?[​](https://stryker-mutator.io/docs/#but-wait-what-about-code-coverage "Direct link to But wait, what about code coverage?")
Well... code coverage doesn't tell you everything about the effectiveness of your tests. Think about it, when was the last time you saw a test without an assertion, purely to increase the code coverage?
Imagine a sandwich covered with paste. Code coverage would tell you the bread is 80% covered with paste. Mutation testing, on the other hand, would tell you it is actually _chocolate_ paste and not... well... something else.
## Meet: Stryker[​](https://stryker-mutator.io/docs/#meet-stryker "Direct link to Meet: Stryker")
Sounds complicated? Don't worry! Stryker has your back. It uses one design mentality to implement mutation testing on three platforms. It's _easy to use_ and _fast to run_. Stryker will only mutate _your source code_ , making sure there are no false positives.
## An example[​](https://stryker-mutator.io/docs/#an-example "Direct link to An example")
Say you're building an online casino. Users can only enter the casino when they're over 18. So you write the following piece of code to check if someone is allowed to use the website:
```
functionisUserOldEnough(user){return user.age>=18;}
```

Stryker will find the return statement and decide to change it in several ways:
```
/* 1 */return user.age>18;/* 2 */return user.age<18;/* 3 */returnfalse;/* 4 */returntrue;
```

We call those modifications mutants. After the mutants have been found, they are applied one by one, and your tests are executed against them. If at least one of your tests fail, we say the mutant is _killed_. That's what we want! If no test fails, it _survived_. The better your tests, the fewer mutants survive.
Stryker will output the results in various different formats. One of the easiest to read is the clear text reporter:
```
 Mutant killed: /yourPath/yourFile.js: line 10:27 Mutator: BinaryOperator -         return user.age >= 18; +         return user.age > 18; Mutant survived: /yourPath/yourFile.js: line 10:27 Mutator: RemoveConditionals -         return user.age >= 18; +         return true;
```

The clear text reporter will output how exactly your code was modified and which mutator was used. It will then tell us if a mutant was killed, meaning that at least one test failed, or if it survived. The second mutation in this example is marked as a survivor. This means there is probably a test missing that explicitly tests for age lower than 18.
> For someone who hates mutants... you certainly keep some strange company.
_- Professor X_
> Oh, they serve their purpose... as long as they can be controlled.
_- William Stryker_
[Edit this page](https://github.com/stryker-mutator/stryker-mutator.github.io/edit/develop/docs/General/introduction.mdx)
[NextAn example](https://stryker-mutator.io/docs/General/example/)
  * [But wait, what about code coverage?](https://stryker-mutator.io/docs/#but-wait-what-about-code-coverage)
  * [Meet: Stryker](https://stryker-mutator.io/docs/#meet-stryker)
  * [An example](https://stryker-mutator.io/docs/#an-example)


Docs
  * [FAQ](https://stryker-mutator.io/docs/General/faq/)


Community
  * [Slack](https://join.slack.com/t/stryker-mutator/shared_invite/enQtOTUyMTYyNTg1NDQ0LTU4ODNmZDlmN2I3MmEyMTVhYjZlYmJkOThlNTY3NTM1M2QxYmM5YTM3ODQxYmJjY2YyYzllM2RkMmM1NjNjZjM)
  * [Twitter](https://twitter.com/stryker_mutator)


More
  * [Blog](https://stryker-mutator.io/blog/)
  * [GitHub](https://github.com/stryker-mutator/)


[![Info Support Logo](https://stryker-mutator.io/images/info-support.svg)](https://infosupport.com)
Powered by [Info Support](https://www.infosupport.com/open-source/).Stryker is released under the Apache 2.0 license. Site by the [Stryker team](https://github.com/orgs/stryker-mutator/people). Logo by Selina van den Top.


## Source Information
- URL: https://stryker-mutator.io/docs/
- Harvested: 2025-08-19T23:15:37.005697
- Profile: deep_research
- Agent: architect
