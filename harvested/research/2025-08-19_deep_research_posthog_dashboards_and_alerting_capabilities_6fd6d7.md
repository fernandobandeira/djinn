---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T22:09:31.984805'
profile: deep_research
source: https://posthog.com/docs/product-analytics/dashboards
topic: PostHog Dashboards and Alerting Capabilities
---

# PostHog Dashboards and Alerting Capabilities

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


Dashboards
  * Product analytics


  * [Overview](https://posthog.com/docs/product-analytics)


  * [Installation](https://posthog.com/docs/product-analytics/installation)


  * [Capturing events](https://posthog.com/docs/product-analytics/capture-events)


  * [Creating insights](https://posthog.com/docs/product-analytics/insights)


  * [Identifying users](https://posthog.com/docs/product-analytics/identify)


  * [Setting person properties](https://posthog.com/docs/product-analytics/person-properties)


  * [Group analytics](https://posthog.com/docs/product-analytics/group-analytics)


  * [Best practices](https://posthog.com/docs/product-analytics/best-practices)


  * [Troubleshooting and FAQs](https://posthog.com/docs/product-analytics/troubleshooting)


  * [Tutorials and guides](https://posthog.com/docs/product-analytics/tutorials)


  * [Cutting costs](https://posthog.com/docs/product-analytics/cutting-costs)


  * Analysis views


  * [Trends](https://posthog.com/docs/product-analytics/trends/overview)
    * [Getting started](https://posthog.com/docs/product-analytics/trends/overview)
    * [Charts](https://posthog.com/docs/product-analytics/trends/charts)
    * [Filters](https://posthog.com/docs/product-analytics/trends/filters)
    * [Aggregations](https://posthog.com/docs/product-analytics/trends/aggregations)
    * [Breakdowns](https://posthog.com/docs/product-analytics/trends/breakdowns)
    * [Formulas](https://posthog.com/docs/product-analytics/trends/formulas)
    * [Statistical analysis](https://posthog.com/docs/product-analytics/trends/statistical-analysis)
    * [Tips](https://posthog.com/docs/product-analytics/trends/tips)


  * [Funnels](https://posthog.com/docs/product-analytics/funnels)


  * [Dashboards](https://posthog.com/docs/product-analytics/dashboards)


  * [User paths](https://posthog.com/docs/product-analytics/paths)


  * [Stickiness](https://posthog.com/docs/product-analytics/stickiness)


  * [Correlation analysis](https://posthog.com/docs/product-analytics/correlation)


  * [Retention](https://posthog.com/docs/product-analytics/retention)


  * [Lifecycle](https://posthog.com/docs/product-analytics/lifecycle)


  * [SQL](https://posthog.com/docs/data-warehouse/sql)


  * [Calendar heatmap](https://posthog.com/docs/product-analytics/calendar-heatmap)


  * Tools


  * [LLM insights Beta](https://posthog.com/docs/ai-engineering/llms)


  * [Autocapture](https://posthog.com/docs/product-analytics/autocapture)


  * [Privacy controls](https://posthog.com/docs/product-analytics/privacy)


  * [Data management](https://posthog.com/docs/data)


  * [Sharing & embedding](https://posthog.com/docs/product-analytics/sharing)


  * [Subscriptions](https://posthog.com/docs/product-analytics/subscriptions)


  * [Alerts](https://posthog.com/docs/alerts)


  * [Sampling Beta](https://posthog.com/docs/product-analytics/sampling)


  * [Color themes Beta](https://posthog.com/docs/product-analytics/color-themes)


# Dashboards
Last updated: Jul 11, 2025
|[Edit this page](https://github.com/PostHog/posthog.com/tree/master/contents/docs/product-analytics/dashboards.mdx)|
Copy page
On this page
  * Creating a new dashboard
  * Dates and filters
  * Date range overrides
  * Dashboard filters
  * Dashboard options
  * Editing the layout
  * Adding text cards
  * Sharing a dashboard
  * Auto refresh


Dashboards are the easiest way to track all your most important product and performance metrics.
Unlike [notebooks](https://posthog.com/docs/notebooks), which are ideal of adhoc analysis of specific issues, dashboards are designed for tracking common metrics over time. 
You can create a new dashboard from scratch, but we also offer numerous [dashboard templates](https://posthog.com/templates) for tracking things like [website metrics](https://posthog.com/templates/website-dashboard), [product health metrics](https://posthog.com/templates/health-dashboard), and [metrics for large language models](https://posthog.com/docs/ai-engineering/llm-insights).
![Example of a dashboard](https://res.cloudinary.com/dmukukwp6/image/upload/dashboard_light_61b3bab3b6.png)![Example of a dashboard](https://res.cloudinary.com/dmukukwp6/image/upload/dashboard_dark_5f2002f750.png)
## Creating a new dashboard
  1. Click on **Dashboards** in the left hand navigation and then **New Dashboard**.
  2. You can create a dashboard from a list of [available templates](https://posthog.com/templates), or select **Blank Dashboard** to start from scratch.
  3. Name your dashboard, add some optional details if desired.
  4. Your dashboard will be empty – click **+New insight** to create an insight to add. Dashboards support all the core [product analytics insights](https://posthog.com/docs/product-analytics/insights).
  5. Click **Save & add to dashboard** when you've created your insight.


If you already have some insights set up, an alternative approach is to click **Add to dashboard** in any insight. 
> **Important:** Insights can appear on multiple dashboards at the same time, so you don't have to create multiple copies of the same insight.
## Dates and filters
### Date range overrides
New dashboards are set to 'No date range overide' by default. This means all insights will use the date range applied in their configuration. 
Changing the date range on a dashboard forces all the insights to use the same date range, but this doesn't impact the date range that's shown when viewing an individual insight.
### Dashboard filters
Dashboards also support most of the same [filters](https://posthog.com/docs/product-analytics/trends/filters) as individual insights, including:
  * **Event properties:** Properties stored on event, such as the `Current URL` when the event was triggered.
  * **Person properties:** Properties of individual users, such as `company_name`.
  * **Feature flags:** Filtering for users with a specific feature flag enabled.
  * **Group properties:** Account-level properties, like `organization_id`, which are only available if you have the [group analytics](https://posthog.com/docs/product-analytics/group-analytics) add-on.
  * **Cohorts:** Filtering by [cohorts](https://posthog.com/docs/data/cohorts) of users you've already created.


Using filters on dashboards is a useful way to compare usage between different types of users without recreating insights over and over.
You could, for example, duplicate a dashboard – click on the '...' menu and click 'Duplicate' – and apply different filters to each version of the dashboard.
## Dashboard options
### Editing the layout
You can move and resize all the insights on a dashboard by entering the 'Edit layout' mode. 
You can do this in three ways:
  * Tap 'E' on your keyboard
  * Click on the '...' icon at the top of the dashboard and click 'Edit layout'
  * Click on the '...' icon on any insight card and click 'Edit layout'


Tapping 'F' enables fullscreen mode.
> **Note:** Dashboards have separate layouts for mobile and desktop. Shrink your viewport until insights are a single column to edit the mobile one.
### Adding text cards
Clicking the dropdown on the 'Add insight button' reveals the option to add a text card to your dashboards.
Text cards support Markdown formatting for text. You can also drag and drop images.
You can use text cards to annotate your dashboard – useful for adding context for other users of the dashboard. For more in-depth analysis, however, we recommend [creating a notebook](https://posthog.com/docs/notebooks).
### Sharing a dashboard
By clicking 'Share' in the top right corner you can:
  * Restrict edit access to certain members within a project. Dashboards can be shared either by members with administrator privileges or by the dashboard creator ([platforms add-ons](https://posthog.com/platform-addons) required).
  * Create a link to share your dashboard publicly, or embed your dashboard on a website. Read more about this in our [sharing and embedding docs](https://posthog.com/docs/product-analytics/sharing).


### Auto refresh
You can manually refresh a dashboard at any time, but you can also set it to automatically refresh.
To do so, click the dropdown arrow next to refresh button, turn on auto refresh and choose your refresh interval.
This is useful if you have a PostHog dashboard on a TV screen for others to see.
**Important:** Auto refresh only works when the browser tab is active.
### Questions? Ask Max AI.
It's easier than reading through **718 pages of documentation**
Chat with Max
### Community questions
  * [**Carlos**](https://posthog.com/community/profiles/32521) 3 months ago
### [Enable detailed table below the chart](https://posthog.com/questions/enable-detailed-table-below-the-chart)
Would be possible to show the detailed table of the chart within the dashboard widget?
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1728482424/ben_h_8fd5ee8b8b.png)**Ben**(He / Him)](https://posthog.com/community/profiles/31540)3 months ago
Hey Carlos! I raised a [feature request](https://github.com/PostHog/posthog/issues/33025) yesterday which is similar to what you are looking for here. Please feel free to subscribe to it and add any further detail you think is relevant.
We can't promise that we'll build every feature request, as we have a lot on our roadmap, but we'll certainly give it serious thought!
10
Reply
  * [**Pieter**](https://posthog.com/community/profiles/29826) 6 months ago
### [Filter Group for Dashboard](https://posthog.com/questions/filter-group-for-dashboard)
Hi, it looks like it's not possible at the moment to apply a filter group to a dashboard? Is this on the roadmap? Saves us from applying the filter groups to each insight individually.
Reply
  * [**Jake**](https://posthog.com/community/profiles/32585) 7 months ago
### [Resizing](https://posthog.com/questions/resizing)
I'm not able to resize any of the insights on my dashboard. I'm in Edit Layout mode and I see the handles for resizing but no matter what I do, the component won't resize. I've tried different widths for my browser in case it was layout specific but no luck.
    * [**Guillaume**](https://posthog.com/community/profiles/32589) Edited 7 months ago
Same here. Tried clearing cache and force reloading, also tried with a brand new dashboard. I can drag and move, but cannot "pick" the handles and resize, which worked fine a few weeks back when I created my dashboards.
00
Reply
  * [**Juliette**](https://posthog.com/community/profiles/32179) 9 months ago
### [Rearrange charts ?](https://posthog.com/questions/rearrange-charts)
Is it possible to change the order of the charts on the dashboard ?
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1695024176/andy_86a7232754.png)**Andy**](https://posthog.com/community/profiles/30208) 8 months agoSolution
Yes. Click on the '...' menu and then 'Edit layout' on any dashboard.
Alternatively, pressing 'E' key will enabled the layout mode as well.
00
Reply
  * [**Josh**](https://posthog.com/community/profiles/31048) 10 months ago
### [Filtering SQL insights directly from data warehouse](https://posthog.com/questions/filtering-sql-insights-directly-from-data-warehouse)
Consider a dashboard where we are trying to get a holistic view of a certain company of users in our application. We send this company data as a group to posthog, which makes it really easy to filter our event based dashboards.
But now we want to add company level data from our MySQL Data Warehouse. How can we dynamically apply the same filters? Is it even possible? For example, something like
```

select count(*)
from pending_users 
where company_uuid ={filter.group.company.id}

```

After a bit of digging, I think at a minimum we would need the **extended group properties** feature, like is available for person properties/joins but even then, probably not quite as straightforward to achieve what I had in mind. Any suggestions?
    * [**Martin**](https://posthog.com/community/profiles/32842) Edited 6 months ago
I second that. As an example we have 15 Sql Insights on a Dashboard and if we duplicate the dashboard with the Insights we need to edit those 15 Insights to adjust for the name of our specific event.
```

select count()
from s3_impulse_prod_humans
where EventId='nameofevent'

```

We would love to be able to set a $EventId variable at the Dashboard level so we can deploy new Dashboard super fast without manually editing all SQL requests.
00
Reply
  * [**Jonathan**](https://posthog.com/community/profiles/31757) 10 months ago
### [Adding recent recording to dashboard](https://posthog.com/questions/adding-recent-recording-to-dashboard)
Hi, is there a way to add "recent recordings" (like the widget on the home screen) to a dashboard? Thanks!
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1688575125/paul_64ee2de98e.png)**Paul**(he/him)](https://posthog.com/community/profiles/30173)10 months agoSolution
hey, there is, using hogql!
you can add a sql insight like
```

select person.properties.email, min_first_timestamp as start,recording_button(session_id)
from raw_session_replay_events
where min_first_timestamp >=now()- interval 1 day
and min_first_timestamp <=now()
order by min_first_timestamp desc
limit 10

```

00
    * [**Jonathan**](https://posthog.com/community/profiles/31757) Author10 months ago
Thank you Paul! Just a small correction to your query - in order to save as view, recording_button(session_id) must have an alias (or at least this is what to button tooltip says)
00
Reply
  * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1726827779/avatar_square_08f851ae2d.png)**Alex**](https://posthog.com/community/profiles/31609) a year ago
### [Breakdown filter for dashboard?](https://posthog.com/questions/breakdown-filter-for-dashboard)
Is there a way to make a dashboard level variable with the flag name to do breakdown in graphs by it.
For context: I would like to have one dashboard for all AB tests of landing pages with the same set of graphs. A new flag is created for each AB test, so I have to change the breakdown by feature flag in each graph manually for each new AB test.
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1695024176/andy_86a7232754.png)**Andy**](https://posthog.com/community/profiles/30208) 8 months ago
No. Breakdowns are specific to insights.
00
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1726827779/avatar_square_08f851ae2d.png)**Alex**](https://posthog.com/community/profiles/31609) Author2 months agoSolution
Just noticed you added dashboard-level breakdowns, that's awesome, thanks! ![Screenshot 2025-06-13 at 15.09.45.png](https://res.cloudinary.com/dmukukwp6/image/upload/v1749820340/Screenshot_2025_06_13_at_15_09_45_f4ee1b13b9.png)
00
Reply
  * [![](https://www.gravatar.com/avatar/400c8736de118a45ac9275ee970f9c29)**Marc**](https://posthog.com/community/profiles/29759) 2 years ago
### [Share with external users?](https://posthog.com/questions/share-with-external-users)
Is it possible to share a single dashboard with external users who do not have access to the posthog org without publishing it?
    * ![](https://res.cloudinary.com/dmukukwp6/image/upload/v1698054319/marcus_cb55867b99.png)
![](https://www.gravatar.com/avatar/400c8736de118a45ac9275ee970f9c29)
View 4 other replies
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1698054319/marcus_cb55867b99.png)**Marcus**](https://posthog.com/community/profiles/30211) 2 years agoSolution
Hey Mohamed, publicly shared dashboards are accessible to anyone who has the correct link.
00
Reply


Ask a questionLogin
### Was this page useful?
HelpfulCould be better
Next article
### User paths
User paths are a type of insight that enable you to follow users along their journey through your product and determine where the biggest drop-offs are. You can learn the following from paths: Where users are getting confused or stuck. Which parts of your app people are actually using. Why users aren't discovering new features. Where new users are landing into your marketing website. How to create a paths insight Click Product Analytics on the left sidebar Click the + New insight button…
[Read next article](https://posthog.com/docs/product-analytics/paths)
### Contributors
### [Edwin LimTechnical Content Marketer![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_200/edwin_36a928c30e)](https://posthog.com/community/profiles/33938)### [Zach WaterfieldGrowth Engineer![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_200/Zach_c7f6a5a292)](https://posthog.com/community/profiles/30086)### [Andy VandervellContent Marketer![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_200/andy_86a7232754)](https://posthog.com/community/profiles/30208)+3 more
### Feature availability
Free / Open-source![Partially available](https://posthog.com/docs/product-analytics/dashboards)
Self-serve![Partially available](https://posthog.com/docs/product-analytics/dashboards)
Enterprise![Available](https://posthog.com/docs/product-analytics/dashboards)
### Related articles
  * [How to embed a shared Dashboard within a web page](https://posthog.com/tutorials/how-to-embed-shared-dashboard)


#### Jump to:
  * Creating a new dashboard
  * Dates and filters
  * Date range overrides
  * Dashboard filters
  * Dashboard options
  * Editing the layout
  * Adding text cards
  * Sharing a dashboard
  * Auto refresh
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
- URL: https://posthog.com/docs/product-analytics/dashboards
- Harvested: 2025-08-19T22:09:31.984805
- Profile: deep_research
- Agent: architect
