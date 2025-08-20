---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:38:29.108130'
profile: code_examples
source: https://posthog.com/tutorials/canary-release
topic: PostHog Canary Release and Gradual Rollout Patterns
---

# PostHog Canary Release and Gradual Rollout Patterns

[](https://posthog.com/)
  * [Why PostHog?](https://posthog.com/why)
  * [Products](https://posthog.com/products)
  * [Pricing](https://posthog.com/pricing)
  * [Docs](https://posthog.com/docs)
  * [Community](https://posthog.com/community)
  * [Company](https://posthog.com/about)


[Get started](https://us.posthog.com/signup)
  * [News](https://posthog.com/community)
  * [Posts](https://posthog.com/posts)
  * [Questions](https://posthog.com/questions)
  * [Guides](https://posthog.com/tutorials)
  * [Templates](https://posthog.com/templates)
  * [DeskHog](https://posthog.com/deskhog)
  * [Cool tech jobs](https://posthog.com/cool-tech-jobs)
  * [Newsletter](https://posthog.com/newsletter)
  * [Merch](https://posthog.com/merch)


Guides & tutorials
Back
The latest from the PostHog community
Sign in
Guides & tutorials
  * [Posts](https://posthog.com/posts)
  * [Tutorials](https://posthog.com/tutorials)
  * Feature flags, Persons


# How to do a canary release with feature flags in PostHog
Sep 12, 2023
Posted by
  * [![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_50/ian_31bf16ef7d_c3339bc255)Ian Vanagas](https://posthog.com/community/profiles/29296)


On this page
  * What is a canary release?
  * Prerequisites for canary releases
  * Setting up feature flags
  * The steps to a canary release
  * Monitoring a canary release
  * Further reading


Few things are worse than shipping a new feature, having it unexpectedly break, and then scrambling to fix it. To mitigate problems like this, teams often roll out changes to a subset of users before releasing them to everyone. This is known as a canary release or deployment.
This tutorial explains what a canary release is, and how to set one up and monitor it in PostHog.
## What is a canary release?
A canary release or canary deployment is the process of rolling out a new feature to a subset of users before releasing it to a larger group. Developers check the new feature is working without issues on the limited group. They watch for issues and analyze usage metrics to confirm. Once satisfied with tests and analysis, the feature rolls out to a larger group (or everyone).
![Canary release](https://res.cloudinary.com/dmukukwp6/image/upload/v1710055416/posthog.com/contents/images/tutorials/canary-release/canary.png)
The name comes from the phrase “canary in a coal mine” which alludes to miners bringing a canary into mines with them. If there were toxic gases, the canary would die and stop chirping, giving an early warning to the miners to get out before they do the same. Testing on a small group of users acts as a canary for issues with the feature, preventing those issues from affecting a larger group.
## Prerequisites for canary releases
To set up a canary release in PostHog, we use feature flags. To get those working, we need to set up user identification.
Feature flags check the distinct ID of the user to decide if they should return `true` or `false`. This means users need a distinct ID unless the feature flag rolls out to everyone. Our snippet and JavaScript Web SDK automatically create anonymous ones, but [identifying users](https://posthog.com/docs/integrate/identifying-users) with a value you set (like email) is a better option.
Beyond a distinct ID, users also need properties or groups if you want to use them to target a release. For example, to canary release to a specific organization, you need to call [`group()` or set a group property on an event](https://posthog.com/manual/group-analytics).
## Setting up feature flags
Once you set up PostHog and user identification, you can create the feature flags for the canary release. Go to the [feature flags tab](https://app.posthog.com/feature_flags) in your PostHog instance, add a key, a release condition (like only yourself), and any other details you want.
Add this flag to your code around the feature you want to canary release. Test that it works for you, and that turning it off works as well. Once this goes well, you can expand the release to your users.
## The steps to a canary release
Beyond testing the flag yourself, here is a recommended step-by-step release with conditions for each:
  1. **Just yourself:** set email to equal to your own, like `email equals ian@posthog.com`. Test it yourself to make sure the feature flag works and the feature works as expected.
  2. **Internal team:** set email to contain your domain, like `email contains posthog.com` or `email equals <insert multiple team member emails>`. This enables [testing in production](https://posthog.com/product-engineers/testing-in-production). Make sure to communicate with your team about what is being released so they know to test and look out for issues.
  3. **Beta users, organizations:** [use early access management](https://posthog.com/docs/feature-flags/early-access-feature-management), set email to contain a company domain, or set the group name to equal theirs, like `email contains twitter.com` or `organization_name equals twitter`. To ensure you are aware of issues, communicate with them and monitor a related usage dashboard.
  4. **Expanded beta:** set release to a percentage of users or to match a popular property. Monitor insights and sessions related to the feature compared to those without. The overall metrics are more important here.
  5. **Full release:** set release to 100% of users, check metrics again, delete the flag, and announce the feature. 


At PostHog, we do all of these. Our feature flag page contains features at different stages of rollout. For example, when users have specific issues, we might canary release a fix to them before releasing it to others.
![PostHog's feature flags](https://res.cloudinary.com/dmukukwp6/image/upload/v1710055416/posthog.com/contents/images/tutorials/canary-release/posthog-flags.png)
## Monitoring a canary release
The key to a canary release is hearing from the “canary.” Getting feedback and hearing issues is critical to the process whether it is from your team, a limited group of external users, or everyone. Without communication channels, issues go unreported and cause more problems than needed.
> **Tip:** Combine a canary release with an [in-app survey](https://posthog.com/docs/surveys) to get qualitative feedback about your change.
On top of hearing about issues from users, you can monitor issues in PostHog. This may look like key metrics or events decreasing, or the number of errors increasing.
You can filter your insights or dashboards by a feature flag to see how the release is progressing. For example, if you are releasing a change to the signup page, it is useful to know if it improves conversion. You can set up a funnel for the conversion, then breakdown by the feature flag name.
![Funnel](https://res.cloudinary.com/dmukukwp6/image/upload/v1710055416/posthog.com/contents/images/tutorials/canary-release/funnel.png)
You can also get session recordings related to feature flags. On the session recording list, filter for sessions including your feature flag. This enables you to get a better idea of the details of how users are interacting with your new feature.
> **Tip:** You can go from funnels to session recordings by clicking the completed or dropped off value in the visualization. This will give you a list of users with session recordings related to that funnel and release.
Canary releases ensure higher quality features get shipped and fewer issues impact users. A release process and monitoring are critical for them to work properly. PostHog’s feature flags, analytics, and session recordings provide all the tools to help this happen.
## Further reading
  * [Targeting feature flags on groups, pages, machines, and more](https://posthog.com/tutorials/group-page-machine-flags)
  * [How to bootstrap feature flags in React and Express](https://posthog.com/tutorials/bootstrap-feature-flags-react)
  * [How to evaluate and update feature flags with the PostHog API](https://posthog.com/tutorials/api-feature-flags)


![](https://res.cloudinary.com/dmukukwp6/image/upload/engineer_47d6638eae)
![](https://res.cloudinary.com/dmukukwp6/image/upload/engineer_47d6638eae)
Subscribe to our newsletter
#### Product for Engineers
Read by 60,000+ founders and builders
Subscribe
We'll share your email with Substack
### Questions? Ask Max AI.
It's easier than reading through **718 pages of documentation**
Chat with Max
### Comments
Leave a commentLogin
Posted by
### [Ian VanagasTechnical Content Marketer![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_200/ian_31bf16ef7d_c3339bc255)](https://posthog.com/community/profiles/29296)
On this page
  * What is a canary release?
  * Prerequisites for canary releases
  * Setting up feature flags
  * The steps to a canary release
  * Monitoring a canary release
  * Further reading


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


  * [News](https://posthog.com/community)
  * [Posts](https://posthog.com/posts)
  * [Questions](https://posthog.com/questions)
  * [Guides](https://posthog.com/tutorials)
  * [Templates](https://posthog.com/templates)
  * [DeskHog](https://posthog.com/deskhog)
  * [Cool tech jobs](https://posthog.com/cool-tech-jobs)
  * [Newsletter](https://posthog.com/newsletter)
  * [Merch](https://posthog.com/merch)


  * [Why PostHog?](https://posthog.com/why)
  * [Products](https://posthog.com/products)
  * [Pricing](https://posthog.com/pricing)
  * [Docs](https://posthog.com/docs)
  * [Community](https://posthog.com/community)
  * [Company](https://posthog.com/about)




## Source Information
- URL: https://posthog.com/tutorials/canary-release
- Harvested: 2025-08-20T00:38:29.108130
- Profile: code_examples
- Agent: architect
