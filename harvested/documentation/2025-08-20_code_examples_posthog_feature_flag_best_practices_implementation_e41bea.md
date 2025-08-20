---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:38:02.969925'
profile: code_examples
source: https://posthog.com/docs/feature-flags/best-practices
topic: PostHog Feature Flag Best Practices Implementation
---

# PostHog Feature Flag Best Practices Implementation

[](https://posthog.com/)
  * [Why PostHog?](https://posthog.com/why)
  * [Products](https://posthog.com/products)
  * [Pricing](https://posthog.com/pricing)
  * [Docs](https://posthog.com/docs)
  * [Community](https://posthog.com/community)
  * [Company](https://posthog.com/about)


[Get started](https://us.posthog.com/signup)
  * [Product OS](https://posthog.com/docs/product-os)
  * [Product analytics](https://posthog.com/docs/product-analytics)
  * [Web analytics](https://posthog.com/docs/web-analytics)
  * [Session replay](https://posthog.com/docs/session-replay)
  * [Feature flags](https://posthog.com/docs/feature-flags)
  * [Experiments](https://posthog.com/docs/experiments)
  * [Error tracking](https://posthog.com/docs/error-tracking)
  * [Surveys](https://posthog.com/docs/surveys)
  * [Data pipelines](https://posthog.com/docs/cdp)
  * [Data warehouse](https://posthog.com/docs/data-warehouse)
  * [LLM observability](https://posthog.com/docs/ai-engineering)
  * [Max AI](https://posthog.com/docs/max-ai)


Best practices
  * Feature flags


  * [Overview](https://posthog.com/docs/feature-flags)


  * [Installation](https://posthog.com/docs/feature-flags/installation)


  * [Creating feature flags](https://posthog.com/docs/feature-flags/creating-feature-flags)


  * [Adding your code](https://posthog.com/docs/feature-flags/adding-feature-flag-code)


  * [Testing your flag](https://posthog.com/docs/feature-flags/testing)


  * [Best practices](https://posthog.com/docs/feature-flags/best-practices)


  * [Cutting costs](https://posthog.com/docs/feature-flags/cutting-costs)


  * [Troubleshooting and FAQs](https://posthog.com/docs/feature-flags/common-questions)


  * [Tutorials and guides](https://posthog.com/docs/feature-flags/tutorials)


  * Settings


  * [Project-wide settings](https://posthog.com/docs/feature-flags/project-wide-settings)


  * Features


  * [Server-side local evaluation](https://posthog.com/docs/feature-flags/local-evaluation)


  * [Client-side bootstrapping](https://posthog.com/docs/feature-flags/bootstrapping)


  * [Remote config](https://posthog.com/docs/feature-flags/remote-config)


  * [Early access feature management](https://posthog.com/docs/feature-flags/early-access-feature-management)


  * [Multi-project feature flags](https://posthog.com/docs/feature-flags/multi-project-feature-flags)


  * [Scheduled flag changes](https://posthog.com/docs/feature-flags/scheduled-flag-changes)


# Feature flag best practices
Last updated: Jun 26, 2025
|[Edit this page](https://github.com/PostHog/posthog.com/tree/master/contents/docs/feature-flags/best-practices.mdx)|
Copy page
On this page
  * 1. Use a reverse proxy
  * 2. Call your flag in as few places as possible
  * 3. Identify users
  * 4. Use server-side local evaluation for faster flags
  * 5. Bootstrap flags on the client to make them available immediately
  * 6. Naming tips
  * 7. Roll out progressively
  * 8. Clean up after yourself
  * 9. Fallback to working code
  * 10. Reducing your bill


## 1. Use a reverse proxy
Ad blockers have the potential to disable your feature flags, which can lead to bad experiences, such as users seeing the wrong version of your app, or missing a new feature rollout.
To avoid this, deploy a reverse proxy, which enables you to make requests and send events to PostHog Cloud using your own domain.
This means that requests are less likely to be intercepted by tracking blockers, and your feature flags are more likely to work as intended. You'll also capture more usage data.
PostHog customers on one of our [platforms add-ons](https://posthog.com/platform-addons) can use our [managed reverse proxy](https://posthog.com/docs/advanced/proxy/managed-reverse-proxy), but there are numerous methods to run your own. See our [reverse proxy docs](https://posthog.com/docs/advanced/proxy/managed-reverse-proxy) for more.
## 2. Call your flag in as few places as possible
It should be easy to understand how feature flags affect your code. The more locations a flag is in, the more likely it is to cause problems. For example, a developer could remove the flag in one place but forget to remove it in another.
If you expect to use a feature flag in multiple places, it's a good idea to wrap the flag in a single function or method. For example:
JavaScript
```

functionuseBetaFeature(){
return posthog.isFeatureEnabled('beta-feature')
}

```

## 3. Identify users
Because PostHog evaluates flags based on the user's distinct ID, having different IDs can cause the same user to receive different flag values across different sessions, devices, and platforms. By [identifying](https://posthog.com/docs/getting-started/identify-users) them, you can ensure they consistent flag values.
The same applies to identifying [groups](https://posthog.com/docs/product-analytics/group-analytics) for group-level flags.
## 4. Use server-side local evaluation for faster flags
Evaluating feature flags requires making a request to PostHog for each flag. However, you can improve performance by evaluating flags locally. Instead of making a request for each flag, PostHog will periodically request and store feature flag definitions locally, enabling you to evaluate flags without making additional requests.
Evaluate flags locally when possible, since this enables you to resolve flags faster and with fewer API calls. See our docs on [local evaluation](https://posthog.com/docs/feature-flags/local-evaluation) for more details.
## 5. Bootstrap flags on the client to make them available immediately
Since there is a delay between initializing PostHog and fetching feature flags, feature flags are not always available immediately. This makes them unusable if you want to do something like redirecting a user to a different page based on a feature flag.
To have your feature flags available immediately, you can initialize PostHog with precomputed values until it has had a chance to fetch them. This is called bootstrapping.
See our docs on [bootstrapping](https://posthog.com/docs/feature-flags/bootstrapping) for more details on how to do this.
## 6. Naming tips
Good naming conventions for your flags makes them easier to understand and maintain. Below are tips for naming your flags:
  * **Use descriptive names.** For example, `is_v2_billing_dashboard_enabled` is much clearer than `is_dashboard_enabled`.
  * **Use name "types".** This helps organize them and makes their purpose clear. Types might include experiments, releases, and permissions. For example, instead of `new-billing`, they would be `new-billing-experiment` or `new-billing-release`.
  * **Name flags to reflect their return type.** For example, `is_premium_user` for a boolean, `enabled_integrations` for an array, or `selected_theme` for a single string.
  * **Use positive language for boolean flags.** For example, `is_premium_user` instead of `is_not_premium_user`. This helps avoid double negatives when checking the flag value (e.g. `if !is_not_premium_user` is confusing).


## 7. Roll out progressively
When testing a change behind a feature flag, it is best to roll it out to a small group of users and increase that group over time. This is also known as a [phased rollout](https://posthog.com/tutorials/phased-rollout). It enables you to identify any potential issues ahead of the full release.
For example, at PostHog we often roll out the flag to just the responsible developer. It then moves on to the internal team, then beta users, and finally into a full rollout. This enables us to [test in production](https://posthog.com/product-engineers/testing-in-production), get multiple rounds of feedback, identify issues, and polish the feature before the full release.
## 8. Clean up after yourself
Leaving flags in your code for too long can confuse future developers and create technical debt, especially if it's already rolled out and integrated. Be sure to remove stale flags once they are completely rolled out or no longer needed.
## 9. Fallback to working code
It's possible that a feature flag will return an [unexpected value](https://posthog.com/docs/feature-flags/common-questions#my-feature-flag-called-events-show-none-empty-string-or-false-instead-of-my-variant-names). For example, if the flag is disabled or failed to load due to a network error. 
In this case, its best to check that the feature flag returns a valid expected value before using it. If it isn't, fallback to working code.
## 10. Reducing your bill
We aim to be significantly cheaper than our competitors. To help you reduce your bill, we've created a [dedicated guide](https://posthog.com/docs/feature-flags/cutting-costs) to estimating and reducing your feature flag costs.
### Questions? Ask Max AI.
It's easier than reading through **718 pages of documentation**
Chat with Max
### Community questions
Ask a questionLogin
### Was this page useful?
HelpfulCould be better
Next article
### Cutting feature flag costs
We aim to be significantly cheaper than our competitors. In addition to our pay-as-you-go pricing , below are tips to reduce your feature flag costs. Understanding your bill Feature flags are charged based on requests made to our /flags endpoint, which is used to evaluate feature flags for a given user. This endpoint is called in several scenarios: When explicitly requesting feature flags via server-side SDK methods like getFeatureFlag() , getAllFlags() , or isFeatureEnabled() When using…
[Read next article](https://posthog.com/docs/feature-flags/cutting-costs)
### Contributors
### [kyleswankContributor![](https://avatars.githubusercontent.com/u/89221808?v=4)](https://github.com/kyleswank)### [Vincent GeTechnical Content Marketer![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_200/image_removebg_preview_28_0aedb42bc3)](https://posthog.com/community/profiles/34100)### [Zach WaterfieldGrowth Engineer![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_200/Zach_c7f6a5a292)](https://posthog.com/community/profiles/30086)+3 more
#### Jump to:
  * 1. Use a reverse proxy
  * 2. Call your flag in as few places as possible
  * 3. Identify users
  * 4. Use server-side local evaluation for faster flags
  * 5. Bootstrap flags on the client to make them available immediately
  * 6. Naming tips
  * 7. Roll out progressively
  * 8. Clean up after yourself
  * 9. Fallback to working code
  * 10. Reducing your bill
  * Questions?


[](https://posthog.com/)
  * #### [Products](https://posthog.com/products)
  * [All products](https://posthog.com/products)
  * [Max AI](https://posthog.com/max)
  * [Product analytics](https://posthog.com/product-analytics)
  * [Web analytics](https://posthog.com/web-analytics)
  * [Session replay](https://posthog.com/session-replay)
  * [Feature flags](https://posthog.com/feature-flags)
  * [Error tracking](https://posthog.com/error-tracking)
  * [Experiments](https://posthog.com/experiments)
  * [Surveys](https://posthog.com/surveys)
  * [Product OS](https://posthog.com/product-os)
  * [Data connections](https://posthog.com/cdp)
  * [Customer stories](https://posthog.com/customers)
  * [PostHog vs...](https://posthog.com/blog/tags/comparisons)
  * [For startups](https://posthog.com/startups)
  * [Pricing](https://posthog.com/pricing)
  * [How we do "sales"](https://posthog.com/sales)
  * [Founder stack](https://posthog.com/founder-stack)


  * #### [Product OS](https://posthog.com/docs/product-os)
  * [New? Start here.](https://posthog.com/docs/getting-started/install)
  * [SDKs](https://posthog.com/docs/libraries/js)
  * [Framework guides](https://posthog.com/docs/frameworks)
  * [Data management](https://posthog.com/docs/data)
  * [SQL access](https://posthog.com/docs/sql)
  * [Toolbar](https://posthog.com/docs/toolbar)
  * [API](https://posthog.com/docs/api)


  * #### [Docs](https://posthog.com/docs)
  * [Product analytics](https://posthog.com/docs/product-analytics)
  * [Session replay](https://posthog.com/docs/session-replay)
  * [Feature flags](https://posthog.com/docs/feature-flags)
  * [Error tracking](https://posthog.com/docs/error-tracking)
  * [Experiments](https://posthog.com/docs/experiments)
  * [Surveys](https://posthog.com/docs/surveys)
  * [CDP](https://posthog.com/docs/cdp)
  * [Data warehouse](https://posthog.com/docs/data-warehouse)
  * [Migrate](https://posthog.com/docs/migrate)


  * #### [Community](https://posthog.com/questions)
  * [Questions?](https://posthog.com/questions)
  * [Guides](https://posthog.com/tutorials)
  * [Integrations](https://posthog.com/cdp)
  * [Dashboard templates](https://posthog.com/templates)
  * [Founders](https://posthog.com/founders/all)
  * [Product engineers](https://posthog.com/product-engineers/all)
  * [Tracks](https://posthog.com/tracks)
  * [Merch](https://posthog.com/merch)
  * [Contributors](https://posthog.com/contributors)
  * [Newsletter](https://newsletter.posthog.com)
  * [PostHog FM](https://open.spotify.com/playlist/7A2H2J3WhpJmMEwAhKahWH?si=47418915a8d0447b)
  * [PostHog on GitHub](https://github.com/PostHog/posthog)
  * [Cool tech jobs](https://posthog.com/cool-tech-jobs)


  * #### [Handbook](https://posthog.com/handbook)
  * [Why we're here](https://posthog.com/handbook/why-does-posthog-exist)
  * [Our story](https://posthog.com/handbook/story)
  * [How we work](https://posthog.com/handbook/company/culture)
  * [Values](https://posthog.com/handbook/values)
  * [Tips for working here](https://posthog.com/handbook/help)
  * [Team structure](https://posthog.com/handbook/team-structure)
  * [Engineering](https://posthog.com/handbook/engineering/developing-locally)
  * [Brand](https://posthog.com/handbook/brand/philosophy)
  * [Marketing](https://posthog.com/handbook/growth/marketing)


  * #### [Company](https://posthog.com/about)
  * [About](https://posthog.com/about)
  * [Roadmap](https://posthog.com/roadmap)
  * [Changelog](https://posthog.com/changelog/2025)
  * [People](https://posthog.com/people)
  * [Small teams](https://posthog.com/teams)
  * [Blog](https://posthog.com/blog/all)
  * [Investors](https://posthog.com/handbook/strategy/investors)
  * [Press](https://posthog.com/media)
  * [FAQ](https://posthog.com/faq)
  * [Security](https://posthog.com/handbook/company/security)
  * [Support](https://posthog.com/questions)
  * [Careers](https://posthog.com/careers)


  * [](https://x.com/posthog)
  * [](https://www.linkedin.com/company/posthog)
  * [](https://www.youtube.com/channel/UCn4mJ4kK5KVSvozJre645LA)
  * [](https://github.com/PostHog)


© 2025 PostHog, Inc.
  * [System status](https://status.posthog.com)
  * 👉[Generate a DPA](https://posthog.com/dpa)👈(It's guaranteed fun!)
  * [SOC 2](https://posthog.com/docs/privacy/soc2)
  * [HIPAA](https://posthog.com/docs/privacy/hipaa-compliance)
  * [Privacy policy](https://posthog.com/privacy)
  * [Terms](https://posthog.com/terms)


**PostHog.com doesn't use third party cookies** - only a single in-house cookie.
No data is sent to a third party.
AcceptDecline
![Ursula von der Leyen, President of the European Commission](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_250/v1/posthog.com/src/components/EU/images/ursula)
  * [Product OS](https://posthog.com/docs/product-os)
  * [Product analytics](https://posthog.com/docs/product-analytics)
  * [Web analytics](https://posthog.com/docs/web-analytics)
  * [Session replay](https://posthog.com/docs/session-replay)
  * [Feature flags](https://posthog.com/docs/feature-flags)
  * [Experiments](https://posthog.com/docs/experiments)
  * [Error tracking](https://posthog.com/docs/error-tracking)
  * [Surveys](https://posthog.com/docs/surveys)
  * [Data pipelines](https://posthog.com/docs/cdp)
  * [Data warehouse](https://posthog.com/docs/data-warehouse)
  * [LLM observability](https://posthog.com/docs/ai-engineering)
  * [Max AI](https://posthog.com/docs/max-ai)


  * [Why PostHog?](https://posthog.com/why)
  * [Products](https://posthog.com/products)
  * [Pricing](https://posthog.com/pricing)
  * [Docs](https://posthog.com/docs)
  * [Community](https://posthog.com/community)
  * [Company](https://posthog.com/about)




## Source Information
- URL: https://posthog.com/docs/feature-flags/best-practices
- Harvested: 2025-08-20T00:38:02.969925
- Profile: code_examples
- Agent: architect
