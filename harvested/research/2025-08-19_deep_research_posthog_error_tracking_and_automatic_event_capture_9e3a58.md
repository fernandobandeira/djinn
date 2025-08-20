---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T22:10:02.664992'
profile: deep_research
source: https://posthog.com/docs/product-analytics/autocapture
topic: PostHog Error Tracking and Automatic Event Capture
---

# PostHog Error Tracking and Automatic Event Capture

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


Autocapture
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


# Autocapture
Last updated: Aug 19, 2025
|[Edit this page](https://github.com/PostHog/posthog.com/tree/master/contents/docs/product-analytics/autocapture.mdx)|
Copy page
On this page
  * Types of autocaptured events
  * Supported SDKs
  * Interaction autocapture
  * Web interaction autocapture
  * iOS interaction autocapture
  * React Native interaction autocapture
  * Navigation and lifecycle event autocapture
  * Web navigation autocapture
  * iOS navigation and lifecycle autocapture
  * React Native navigation and lifecycle autocapture
  * Clipboard autocapture
  * Heatmap autocapture
  * Dead clicks autocapture
  * Analyzing autocaptured events and properties
  * Captured properties


PostHog can automatically capture a variety of events in your app without specific tracking code. This page covers the different types of events that PostHog can capture and how to configure them.
## Types of autocaptured events
PostHog can automatically capture these types of data without specific tracking code:
Type| Description  
---|---  
[Interaction](https://posthog.com/docs/product-analytics/autocapture#interaction-autocapture)| Captures clicks, taps, and other user interactions  
[Navigation](https://posthog.com/docs/product-analytics/autocapture#navigation-autocapture)| Captures pageviews, pageleaves, and screen views  
[Clipboard](https://posthog.com/docs/product-analytics/autocapture#clipboard-autocapture)| Captures copy and paste actions  
[Heatmap](https://posthog.com/docs/product-analytics/autocapture#heatmap-autocapture)| Shows where users interact with your product the most  
[Dead clicks](https://posthog.com/docs/product-analytics/autocapture#dead-clicks-autocapture)| Captures clicks that don't trigger a change to the page  
[Exception](https://posthog.com/docs/error-tracking/start-here)| Captures errors and crashes. See the [error tracking](https://posthog.com/docs/error-tracking/start-here) docs for more information.  
[Session](https://posthog.com/docs/session-replay/installation)| Records real user behavior for playback. See the [session replay](https://posthog.com/docs/session-replay/installation) docs for more information.  
[Web vitals](https://posthog.com/docs/web-analytics/web-vitals)| Captures largest contentful paint, first input delay, cumulative layout shift, and first contentful paint  
[Lifecycle](https://posthog.com/docs/product-analytics/autocapture#navigation-and-lifecycle-event-autocapture)| Captures app launches, backgrounds, and updates.  
## Supported SDKs
Autocapture is available in the following SDKs, each with a different set of captured events: 
  * [JavaScript Web](https://posthog.com/docs/libraries/js) - enabled by default
  * [React](https://posthog.com/docs/libraries/react) - enabled by default
  * [React Native](https://posthog.com/docs/libraries/react-native) - disabled by default
  * [iOS](https://posthog.com/docs/libraries/ios) - disabled by default


**Looking for session and exception autocapture?**
You can also autocapture session and exception events. See [session replay](https://posthog.com/docs/session-replay/installation) and [error tracking](https://posthog.com/docs/error-tracking/start-here) for more information.
## Interaction autocapture
  * Web
  * iOS
  * React Native


### Web interaction autocapture
The [JavaScript web SDK](https://posthog.com/docs/libraries/js) captures the following events by default:
  * Interactions like clicks with a tag like `a`, `button`, `form`, `input`, `select`, `textarea`, `label`
  * Form submissions and form changes
  * Changes on content with `contenteditable="true"`. 
  * Copies from clipboard


Autocaptures are [displayed](https://app.posthog.com/activity/explore) with names like `clicked span with text "Delete"`. You can filter for **Autocapture** events to see all interactions.
#### Configuring interaction autocapture
You can configure `posthog-js` to autocapture information that users copy or cut on your page with the `capture_copied_text` config option. 
You can configure the following options:
Option| Description  
---|---  
`url_allowlist`**Type:** `(string \| RegExp)[]`| List of URLs to enable autocapture on. Can be strings to match or regexes (e.g., `['https://example.com', 'test.com/.*']`). Useful when you want to autocapture on specific pages only. If both `url_allowlist` and `url_ignorelist` are set, the allowlist is checked first, then the ignorelist (which can override the allowlist).  
`url_ignorelist`**Type:** `(string \| RegExp)[]`| List of URLs to not enable autocapture on. Can be strings to match or regexes (e.g., `['https://example.com', 'test.com/.*']`). Useful when you want to autocapture on most pages but not some specific ones.  
`dom_event_allowlist`**Type:** `DomAutocaptureEvents[]`| List of DOM events to enable autocapture on (e.g., `['click', 'change', 'submit']`).  
`element_allowlist`**Type:** `AutocaptureCompatibleElement[]`| List of DOM elements to enable autocapture on (e.g., `['a', 'button', 'form', 'input', 'select', 'textarea', 'label']`). We consider the element tree from root to target, so if `button` is in the allowlist, clicks on `button` or its children (like `svg`) are captured, but not clicks on parent `div` elements.  
`css_selector_allowlist`**Type:** `string[]`| List of CSS selectors to enable autocapture on (e.g., `['[ph-capture]']`). We consider the element tree from root to target, so if `['[id]']` is in the allowlist, clicks on elements with IDs or their parents with IDs are captured. Everything is enabled when there's no allowlist.  
`element_attribute_ignorelist`**Type:** `string[]`| Exclude certain element attributes from autocapture (e.g., `['aria-label']` or `['data-attr-pii']`).  
`capture_copied_text`**Type:** `boolean`| When set to `true`, autocapture will capture the text of any element that is cut or copied.  
Here's a full example of how to configure autocapture:
Web
```

posthog.init('<ph_project_api_key>',{
api_host:'https://us.i.posthog.com',
defaults:'2025-05-24',
autocapture:{
// URL filtering
url_allowlist:['https://example.com','test.com/.*'],
url_ignorelist:['https://example.com/admin','test.com/private/.*'],
// Event filtering
dom_event_allowlist:['click','change','submit','input'],
// Element filtering
element_allowlist:['a','button','form','input','select','textarea','label'],
css_selector_allowlist:['[ph-capture]','[data-track]'],
// Attribute filtering
element_attribute_ignorelist:['aria-label','data-attr-pii','data-sensitive'],
// Copy/cut capture
capture_copied_text:true,
},
})

```

#### Disabling interaction autocapture
You can disable autocapture in two ways:
  * In your [project settings](https://us.posthog.com/project/settings)
  * By setting `autocapture: false` in the [config](https://posthog.com/docs/libraries/js/config)


Web
```

// Before initialization
posthog.init('<ph_project_api_key>',{
api_host:'https://us.i.posthog.com',
defaults:'2025-05-24',
autocapture:false,
})
// Or after initialization
posthog.set_config({autocapture:false})

```

> **Note:** Disabling autocapture in your project settings isn't instant. PostHog will continue to capture events where it is initialized. Instead, autocaptured events will taper off as users trigger reinitialization of PostHog (like when they reload your site). 
> Disabling autocapture does not disable capture of page views or page leaves. See [navigation autocapture](https://posthog.com/docs/product-analytics/autocapture#navigation-autocapture) for more information.
#### Capturing additional properties in autocapture events
If you add a data attribute onto an element in the format `data-ph-capture-attribute-some-key={someValue}`, then any autocapture event from that element or one of its children will have the property `some-key: 'someValue'` added to it. This can be useful when you want to add additional information to autocapture events. 
As an example, say you have a notification bell with a value like this:
![a notification bell showing 1 unread notification](https://res.cloudinary.com/dmukukwp6/image/upload/v1710055416/posthog.com/contents/images/docs/data/autocapture/notification-bell.png)
You can include the unread count in the autocapture event by adding the `data-ph-capture-attribute` class like this:
HTML
```

<div
onClick={toggleNotificationsPopover}
data-ph-capture-attribute-unread-notifications-count={unreadCount}
>

```

The autocapture event for clicks on the bell will include the unread count as an `unread-notifications-count` property.
#### Tracking metadata
You can also attach metadata to autocapture events by adding data attributes to the element that triggers the event. This helps you track something like a customer performing a transaction (adding an item to a cart or completing a purchase).
The below ecommerce example helps you understand _what_ users are interested in, even if they don't complete a transaction. It can also reveal which products users are interested in when correlated with information like marketing campaigns, regionality, or device type.
HTML
```

<button
data-ph-capture-attribute-product-id={productId}
data-ph-capture-attribute-product-name={productName}
data-ph-capture-attribute-product-price={productPrice}
data-ph-capture-attribute-product-quantity={productQuantity}
>
  Add to cart
</button>

```

Replace the `{productXx}` values with the relevant information available on the webpage. Now when the _Add to cart_ button is clicked, the autocapture event will include the product information in the event's properties, like:
JSON
```

properties:{
"product-id":"12345678",
"product-name":"Red t-shirt",
"product-price":"30",
"product-quantity":"1"
}

```

#### Sending custom properties with autocaptured form submissions
To prevent accidental sensitive data capture, we do not automatically capture form values from form submissions. To add custom properties to form submissions, you can use the `data-ph-capture-attribute` attribute on the form element.
In this example, the `product-id` property will be sent with the form submission event. The `product-name` input will not be captured.
HTML
```

<form
data-ph-capture-attribute-product-id="12345678"
>
<inputname="product-name"placeholder="this input will not be autocaptured"/>
<buttontype="submit">Submit</button>
</form>

```

The following example shows how you can use the `data-ph-capture-attribute` attribute dynamically in a React component. The `product-name` property will be sent with the form submission event.
TSX
```

import{ useState }from'react'
constMyForm=()=>{
const[productName, setProductName]=useState('')
return(
<form
data-ph-capture-attribute-name={productName}
>
<inputname="product-name"onChange={(e)=>setProductName(e.target.value)}value={productName}/>
<buttontype="submit">Submit</button>
</form>
)
}

```

#### Rage clicks autocapture
A rage click is when a user clicks a button multiple times in quick succession, specifically more than 3 clicks in 1 second.
The event is captured as a `$rageclick` event. You can use this event to identify opportunities to improve your UI, as it shows where users are frustrated with your product.
## Navigation and lifecycle event autocapture
  * Web
  * iOS
  * React Native


### Web navigation autocapture
PostHog automatically captures navigation events as `$pageview` and `$pageleave` events. 
#### Configuring page navigation autocapture
You can configure the following options:
Option| Description  
---|---  
`capture_pageview`| **Type:** `boolean`, `'history_change'` (default). When set to `history_change`, autocapture will capture pageviews. This also works on single-page apps.  
`capturePageleaves`| **Type:** `boolean`, `'if_capture_pageview'` (default). When set to `true`, autocapture will capture pageleaves. If set to `'if_capture_pageview'`, it only captures pageleaves if `capturePageviews` is also set to `true` or `'history_change'`.  
For example, your initialization code might look like this:
Web
```

posthog.init('<ph_project_api_key>',{
api_host:'https://us.i.posthog.com',
defaults:'2025-05-24',
capture_pageview:'history_change',
capture_pageleave:'if_capture_pageview'
})

```

#### Disabling page navigation autocapture
You can disable `$pageview` and `$pageleave` autocapture in two ways:
  * In your [project settings](https://us.posthog.com/project/settings)
  * By setting `capturePageviews: false` and `capture_pageleave: false` in the [config](https://posthog.com/docs/libraries/js/config)


For example:
Web
```

posthog.init('<ph_project_api_key>',{
api_host:'https://us.i.posthog.com',
defaults:'2025-05-24',
capture_pageview:false,
  capture capture_pageleave ageleaves:false,
})

```

## Clipboard autocapture
> **Note:** This is only available for the JavaScript web SDK and framework libraries for web like React and Next.js.
When enabled with `capture_copied_text` set to true, we capture the copied or cut text as a **Clipboard autocapture** event. You can then use the `$selected_content` property in analysis or use the [activity page](https://posthog.com/docs/activity) to view the copied content in context.
![The activity view showing the copied content highlighted in context](https://res.cloudinary.com/dmukukwp6/image/upload/posthog.com/contents/images/docs/autocapture/clipboard-activity-light.png)![The activity view showing the copied content highlighted in context](https://res.cloudinary.com/dmukukwp6/image/upload/posthog.com/contents/images/docs/autocapture/clipboard-activity-dark.png)
Clipboard autocapture respects other privacy settings. For example, won't capture content from a password field.
> **Note:** Browsers don't directly allow access to copied data for privacy reasons so when `posthog-js` sees a clipboard event, we capture any text currently selected in the browser.import { useSelector } from 'react-redux'
## Heatmap autocapture
> **Note:** This is only available for the JavaScript web SDK and framework libraries for web like React and Next.js.
If you use our JavaScript libraries and enable heatmap autocapture in your [project settings](https://app.posthog.com/settings/project-autocapture#heatmaps), we can capture general clicks, mouse movements, and scrolling to create [heatmaps](https://posthog.com/docs/toolbar/heatmaps). No additional events are created.
Whereas autocapture creates events whenever it can uniquely identify an interacted element, heatmaps are generated based on overall mouse or touch positions and are useful for understanding more general user behavior across your site.
## Dead clicks autocapture
A dead click (or slow click) is a click which isn't followed by a change to the page. 
Dead clicks are a great way to identify opportunities to improve your UI, showing you where your users expect to be able to interact with the page but cannot.
You can collect dead clicks with the Web SDK by enabling them in [your project settings](https://app.posthog.com/settings/project-autocapture#dead-clicks-autocapture). 
Or by setting your config:
Web
```

posthog.init('<ph_project_api_key>',{
api_host:'https://us.i.posthog.com',
defaults:'2025-05-24',
capture_dead_clicks:true,
// any other config
})

```

> **Note:** The PostHog heatmap captures dead clicks for free, collecting only the coordinates of dead clicks to display in heatmaps. Enabling the autocapture of dead clicks here allows for deeper analysis and is priced as a standard product analytics event.
## Analyzing autocaptured events and properties
Autocapture events and properties can be used like any other [event type](https://posthog.com/docs/data/events). You can use them in trends, funnels, cohorts, surveys, and more. Beyond this, they come with some special features:
  * When using the autocapture event series, you can filter by the autocaptured element's tag name, text, `href` target, and/or CSS selector.


![Trends using autocapture properties](https://res.cloudinary.com/dmukukwp6/image/upload/auto_light_b669bff067.png)![Trends using autocapture properties](https://res.cloudinary.com/dmukukwp6/image/upload/auto_dark_c52206511f.png)
  * Autocapture events can be organized and renamed using [actions](https://posthog.com/docs/data/actions).
  * You can query autocapture `elements_chain` using [SQL](https://posthog.com/tutorials/hogql-autocapture).


### Captured properties
  * Web
  * iOS
  * React Native


Autocaptured events (and client-side custom events) have many default properties. These are distinguished by `$` prefix in their name, the PostHog logo next to them in the activity tab, and the verified event logo. You can find them in [PostHog](https://app.posthog.com/data-management/properties) or in [the references](https://posthog.com/docs/references/posthog-js/types/Properties).
Name| Key| Example value  
---|---|---  
Timestamp| `$timestamp`| `2024-05-29T17:32:07.202Z`  
OS| `$os`| `Mac OS X`  
OS Version| `$os_version`| `10.15.7`  
Browser| `$browser`| `Chrome`  
Browser Version| `$browser_version`| `125`  
Device Type| `$device_type`| `Desktop`  
Current URL| `$current_url`| `https://example.com/page`  
Host| `$host`| `example.com`  
Path Name| `$pathname`| `/page`  
Screen Height| `$screen_height`| `1080`  
Screen Width| `$screen_width`| `1920`  
Viewport Height| `$viewport_height`| `950`  
Viewport Width| `$viewport_width`| `1903`  
Library| `$lib`| `web`  
Library Version| `$lib_version`| `1.31.0`  
Referrer URL| `$referrer`| `https://google.com`  
Referring Domain| `$referring_domain`| `www.google.com`  
Active Feature Flags| `$active_feature_flags`| `['beta_feature']`  
Event Type| `$event_type`| `click`  
UTM Source| `$utm_source`| `newsletter`  
UTM Medium| `$utm_medium`| `email`  
UTM Campaign| `$utm_campaign`| `product_launch`  
UTM Term| `$utm_term`| `new+product`  
UTM Content| `$utm_content`| `logolink`  
Google Click ID| `$gclid`| `TeSter-123`  
Google Ads Source| `$gad_source`| `google_ads`  
Google Search Ads 360 Click| `$gclsrc`| `dsa`  
Google DoubleClick Click ID| `$dclid`| `testDclid123`  
Google Web-to-app Measure| `$wbraid`| `testWbraid123`  
Google App-to-web Measure| `$gbraid`| `testGbraid123`  
Facebook Click ID| `$fbclid`| `testFbclid123`  
Microsoft Click ID| `$msclkid`| `testMsclkid123`  
Twitter Click ID| `$twclid`| `testTwclid123`  
LinkedIn Ad Tracking ID| `$la_fat_id`| `testLaFatId123`  
Mailchimp Campaign ID| `$mc_cid`| `testMcCid123`  
Instagram Share Id| `$igshid`| `testIgshid123`  
TikTok Click ID| `$ttclid`| `testTtclid123`  
IP Address| `$ip`| `192.168.1.1`  
#### Notes:
  * If enabled, [GeoIP data](https://posthog.com/docs/cdp/geoip-enrichment) is added also as properties at ingestion.
  * Many of these are also captured as [session properties](https://posthog.com/docs/data/sessions).
  * These properties can be hidden in activity by checking the **Hide PostHog properties** box.


### Questions? Ask Max AI.
It's easier than reading through **718 pages of documentation**
Chat with Max
### Community questions
  * [![](https://www.gravatar.com/avatar/4e0ddbc80aa955505e2d8d75016b697f)**Vipul**](https://posthog.com/community/profiles/34621) 2 months ago
### [How to set limited autocapture event with css_selector_allowlist?](https://posthog.com/questions/how-to-set-limited-autocapture-event-with-css-selector-allowlist)
I set the css_selector_allowlist and added the attribute to the button to check if autocapture works only where I want, but no event is being triggered. This is provider setup:
```

"use client";
importposthogfrom"posthog-js";
import{PostHogProvider}from"posthog-js/react";
exportfunctionPHProvider({ children }){
if(typeofwindow!=="undefined"){
    posthog.init(process.env.NEXT_PUBLIC_POSTHOG_KEY,{
api_host: process.env.NEXT_PUBLIC_POSTHOG_HOST,
capture_pageview:false,
autocapture:{
css_selector_allowlist:["[ph-autocapture]"]
},
disable_web_experiments:false,
loaded:(posthog)=>{
// development mode
if(process.env.NODE_ENV==="development") posthog.debug();
}
});
}
return<PostHogProvider client={posthog}>{children}</PostHogProvider>;
}

```

This is how attribute passed to the button to allow autocapture on click:
```

<button ph-autocapture>Click me</button>

```

    * ![](https://www.gravatar.com/avatar/4e0ddbc80aa955505e2d8d75016b697f)
View 3 other replies
    * [![](https://www.gravatar.com/avatar/4e0ddbc80aa955505e2d8d75016b697f)**Vipul**](https://posthog.com/community/profiles/34621) Author2 months ago
PH Provider code:
```

"use client";
importposthogfrom"posthog-js";
import{PostHogProvider}from"posthog-js/react";
exportfunctionPHProvider({ children }){
if(typeofwindow!=="undefined"){
    posthog.init(process.env.NEXT_PUBLIC_POSTHOG_KEY,{
api_host: process.env.NEXT_PUBLIC_POSTHOG_HOST,
capture_pageview:false,
autocapture:{
css_selector_allowlist:["[ph-autocapture]"],
dom_event_allowlist:["click"],
element_allowlist:["button"]
},
disable_web_experiments:false,
loaded:(posthog)=>{
// development mode
if(process.env.NODE_ENV==="development") posthog.debug();
}
});
}
return<PostHogProvider client={posthog}>{children}</PostHogProvider>;
}

```

After removing css_selector_allowlist, we were able to autocapture onClick events, but it started capturing every button click. We also tested using url_allowlist, which worked fine. However, with css_selector_allowlist, events are not being captured regardless of the parameters we pass.
00
Reply
  * [**Tal**](https://posthog.com/community/profiles/33047) 6 months ago
### [How to capture element Id?](https://posthog.com/questions/how-to-capture-element-id)
I don't see element Id being captured in events? Why is this? How do I capture them with minimal changes to my app? I do set id="foo" for example on some things.
Reply
  * [**Don**](https://posthog.com/community/profiles/32347) 7 months ago
### [Does disabling auto-capture disable pageView and pageLeave?](https://posthog.com/questions/does-disabling-auto-capture-disable-page-view-and-page-leave)
I want to disable auto-capture but keep pageView and pageLeave, would this config do the trick
```

PostHogProvider
    apiKey={import.meta.env.VITE_PUBLIC_POSTHOG_KEY}
    options={{
api_host:import.meta.env.VITE_PUBLIC_POSTHOG_HOST,
person_profiles:'identified_only',
autocapture:false,
rageclick:false,
capture_pageview:true,
capture_pageleave:true,
disable_session_recording:true,
}}
>
{children}
</PostHogProvider>

```

    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1728482366/Ben_L_a3ac90960f.png)**Ben**](https://posthog.com/community/profiles/31567) 7 months ago
This looks good to me Don!
Reach out if you run into any trouble.
00
    * [**Tal**](https://posthog.com/community/profiles/33047) 6 months ago
Hi, not Don, but how do I capture element ID? It seems to not be working for me and several other people in this comments section.
00
Reply
  * [**Ahmed**](https://posthog.com/community/profiles/32389) 8 months ago
### [How can I resolve this CORS issue with dead-clicks-autocapture.js](https://posthog.com/questions/how-can-i-resolve-this-cors-issue-with-dead-clicks-autocapture-js)
I am trying to use autocapture for my nextJS website. When the page loads, the dead-clicks-autocapture.js fails due to CORS. How can I fix this?
Reply
  * [**Jason**](https://posthog.com/community/profiles/32279) 9 months ago
### [Autocapture events not available in product analytics](https://posthog.com/questions/autocapture-events-not-available-in-product-analytics)
We recently enable autocapture with PostHog. I see the autocapture events and can filter them as expected in the Activity>Explore page. However, when I try to make an insight chart with Product Analytics (e.g. trend, funnel), I have no option to select/filter the autocapture events.
Any suggestions what might be missing?
Reply
  * [**Shavit**](https://posthog.com/community/profiles/32219) 9 months ago
### [Autocapture by element id](https://posthog.com/questions/autocapture-by-element-id)
I autocapture click events in my app. I would like to view the metrics results by element ID, BUT at the moment I can set the column configuration by element text. How can I set the autocapture to capture element ID? and how can the element id be added to the element properties options?
Reply
  * [**Shavit**](https://posthog.com/community/profiles/32219) 9 months ago
### [Autocapture - element ID](https://posthog.com/questions/autocapture-element-id)
Hi , I use autocapture to collect click events. at the moment I have data-attr configured in settings - autocapture I can only filter the element clicks by element text.. What do I need to do in order to add the element id option? My app elements have the id= attribute. Do I need to change the snippet?
Reply
  * [**Aleksei**](https://posthog.com/community/profiles/31839) 10 months ago
### [Autocapture of internet connection types](https://posthog.com/questions/autocapture-of-internet-connection-types)
Is it possible to collect info about internet connection of user? (wi-fi, 5G /4G / LTE / 3G etc.)
Reply
  * [**Dhaval**](https://posthog.com/community/profiles/31574) a year ago
### [Intercepting the captureNativeAppLifecycleEvents](https://posthog.com/questions/intercepting-the-capture-native-app-lifecycle-events)
Is there anyway to intercept the captureNativeAppLifecycleEvents? I am specifically interested to intercept the Application Installed event to push this event to GTM(Google Tag Manager).
Reply
  * [**landry**](https://posthog.com/community/profiles/31451) a year ago
### [Autocapture + posthog.capture](https://posthog.com/questions/autocapture-posthog-capture)
Is Autocapture automatically disabled when we add a `posthog.capture("my_event")` in the code or can both live alongside ?
Reply
  * [**Snorre Lothar von Gohren**](https://posthog.com/community/profiles/31408) a year ago
### [What element is clicked with an Autocapture event?](https://posthog.com/questions/what-element-is-clicked-with-an-autocapture-event)
Im trying to figure out what element is clicked with an autocapture event. I have a name but I dont know really what element this is. What is the algorithm for setting a name, so I can tell my product manager on how to understand the autocapture events
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1698054319/marcus_cb55867b99.png)**Marcus**](https://posthog.com/community/profiles/30211) a year agoSolution
Hey Snorre, you can access the element for each auto-capture event using the following HogQL expression: `arrayElement(splitByChar('.', elements_chain), 1)`. Check out [this guide](https://posthog.com/tutorials/hogql-autocapture) for more details.
00
Reply
  * [**Muhammad**](https://posthog.com/community/profiles/31156) a year ago
### [Element IDs are not always captured](https://posthog.com/questions/element-i-ds-are-not-always-captured)
I have some click events that are capture but without ids for ex: I have a button with ID that contains and SVG image. what I get is svg clicked without the button id which makes displaying dashboards very hard. Any fix for that?
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1711641190/Steven_fdb952b5e0.png)**Steven**(he/him)](https://posthog.com/community/profiles/28949)a year ago
Hey Muhammad,
Yes, for details that are not autocaptured, check out our docs on how to [capture custom events](https://posthog.com/docs/getting-started/send-events#2-capture-custom-events).
00
    * [**Patryk**](https://posthog.com/community/profiles/32569) Edited 7 months ago
You can style it certain way for it to show up on auto capture
If your button has an aria-label it can be identified, otherwise you may need use "capture custom event" to identify it or show more data.
When you have svg wrapped in a button that's when you get "clicked svg" Here is how you resolve it:
```

importIconButtonfrom'@mui/material/IconButton';
importstyledfrom'@emotion/styled';
constStyledIconButton=styled(
IconButton
)`
> *{
pointer-events: none;
`;

```

That will stop the "clicked svg"
00
Reply
  * [**Zaeem**](https://posthog.com/community/profiles/30878) a year ago
### [Session Event](https://posthog.com/questions/session-event-1)
When a user comes to the product, is there a way to capture/filter session start and session end events with timestamps and duration between them?
Reply
  * [![](https://www.gravatar.com/avatar/955fc06f3f55492e97534859bfcc536b)**Andrey**](https://posthog.com/community/profiles/30827) a year ago
### [How to disable Autocaptured properties?](https://posthog.com/questions/how-to-disable-autocaptured-properties)
Hey, I am working on extension, and don't want accidentally sneak and see what they are viewing, or on what pages they use my extension. Is there any way to receive only custom properties in posthog.capture()?
    * ![](https://www.gravatar.com/avatar/955fc06f3f55492e97534859bfcc536b)
![](https://res.cloudinary.com/dmukukwp6/image/upload/v1698054319/marcus_cb55867b99.png)
![](https://www.gravatar.com/avatar/c5a259c0dca7c0cedc6d1717a7ed0701)
View 7 other replies
    * [![](https://www.gravatar.com/avatar/c5a259c0dca7c0cedc6d1717a7ed0701)**Federico**](https://posthog.com/community/profiles/32694) 7 months ago
Hey Marcus I checked the link you sent to remove event properties, can you give me more help with this? I didn't achieve it with the link
00
Reply
  * [**Daniel**](https://posthog.com/community/profiles/29894) a year ago
### [edit autocaptured events?](https://posthog.com/questions/edit-autocaptured-events)
Is there a way to easily see all the things autocapture is tagging as an event? (I mean the event categories, not measured user events), and then edit them, etc?
    * ![](https://res.cloudinary.com/dmukukwp6/image/upload/v1698054319/marcus_cb55867b99.png)
View 3 other replies
    * [**Daniel**](https://posthog.com/community/profiles/29894) Authora year ago
Ok, I was just wondering if there was a way to make autocapture useful, otherwise I will create customs events.
00
Reply
  * [**Alexander**](https://posthog.com/community/profiles/29776) 2 years ago
### [Autocapture and Heat maps](https://posthog.com/questions/autocapture-and-heat-maps)
If I disable the autocaptured events, will the clicks still show on the Heat map? Or I will be able to see only the custom event clicks?
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1698054319/marcus_cb55867b99.png)**Marcus**](https://posthog.com/community/profiles/30211) 2 years agoSolution
Hey Alexander, that won't work, since the heatmap is using existing auto-captured events under the hood. It's not possible to achieve this with custom events.
00
Reply
  * [**Uri**](https://posthog.com/community/profiles/29567) 2 years ago
### [what is the event_name?](https://posthog.com/questions/what-is-the-event-name)
I'm trying to optimize autocapture so my team won't have to add .capture() programmatically.
    * What's the `event_name` generated by autocapture? I'm guessing it's the "clicked div" events I see in PostHog
    * How to give posthog id for autocapture? e.g. ... so it will report "buy now" instead of "div".
    * Can I adjust the template for event_name?
    * What happens when clicking on a nested element? e.g. cart icon. Will the report say "clicked on svg", and not "clicked on the important 'buy now' button"?
    * Can I produce nesting? ie. report "clicked on product page > buy now" for this:
    * ![](https://res.cloudinary.com/dmukukwp6/image/upload/v1698054319/marcus_cb55867b99.png)
View 3 other replies
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1698054319/marcus_cb55867b99.png)**Marcus**](https://posthog.com/community/profiles/30211) a year ago
Hey Max, you could tag those SVG elements using `data-ph-capture-attribute-some-key={"edit button"}`. Check out [this page](https://posthog.com/docs/product-analytics/autocapture#capturing-additional-properties-in-autocapture-events) for more details.
00
Reply
  * [![](https://www.gravatar.com/avatar/560ae9b0e00499217800ef3c19f0f7ca)**Alireza**](https://posthog.com/community/profiles/28825) 2 years ago
### [Effect of autocapture on heatmaps](https://posthog.com/questions/effect-of-autocapture-on-heatmaps)
Would it affect heatmap if we disable autocapture or limit it for certain elements?
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1688575125/paul_64ee2de98e.png)**Paul**(he/him)](https://posthog.com/community/profiles/30173)2 years agoSolution
Hey 👋,
Heatmaps are built from autocapture information. So, if you disable autocapture that's the same as disabling heatmaps, exclude certain elements and they won't be represented in the heat map.
Thanks
00
Reply
  * [![](https://www.gravatar.com/avatar/560ae9b0e00499217800ef3c19f0f7ca)**Alireza**](https://posthog.com/community/profiles/28825) 2 years ago
### [Is it possible to intercept the autocapture events and add properties to them?](https://posthog.com/questions/is-it-possible-to-intercept-the-autocapture-events-and-add-properties-to-them)
I want to add properties to autocapture events, but no with element attributes. Our usecase is an e-commerce website in which a user can select/change their store (physical store who ships the product). We want to add the currently selected store name to the autocapture events to be able to filter and analyze. For example to know how many clicks we get for different store users. Let me know if you need more details.
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1688575173/simon_bb4af1b047.png)**Simon**(he/him)](https://posthog.com/community/profiles/28895)2 years agoSolution
Hey Alireza - [Super Properties](https://posthog.com/docs/libraries/js#super-properties) will help here. Set one when you know the store name and it will be auto-added to subsequent autocapture events.
00
    * [![](https://www.gravatar.com/avatar/560ae9b0e00499217800ef3c19f0f7ca)**Alireza**](https://posthog.com/community/profiles/28825) Author2 years ago
Thank you! Stupid of me to not read all the docs!
00
Reply
  * [**Robert**](https://posthog.com/community/profiles/29678) 2 years ago
### [Reducing autocaptured properties](https://posthog.com/questions/reducing-autocaptured-properties)
The "Reducing events" section gave examples for DOM events and CSS selectors, but what about reducing autocaptured properties such as `$plugins_deferred`, `$host`, etc?
    * [![](https://res.cloudinary.com/dmukukwp6/image/upload/v1698054319/marcus_cb55867b99.png)**Marcus**](https://posthog.com/community/profiles/30211) 2 years ago
Hey Robert, it's not possible to remove those properties. Since you are not being charged by the amount of properties, what is your intention of reducing those?
00
Reply


Ask a questionLogin
### Was this page useful?
HelpfulCould be better
Next article
### Privacy controls
PostHog offers a range of controls to limit what data is captured by product analytics. They are listed below in order of least to most restrictive. EU-cloud hosting PostHog offers hosting on EU cloud. To use this, sign up at eu.posthog.com . If you've already created a US cloud instance and wish to migrate ticket, you must raise a support ticket in-app with the Data pipelines topic for the PostHog team to do this migration for you. This option is only available to customers on the boost…
[Read next article](https://posthog.com/docs/product-analytics/privacy)
### Contributors
### [Vincent GeTechnical Content Marketer![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_200/image_removebg_preview_28_0aedb42bc3)](https://posthog.com/community/profiles/34100)### [Joshua OrdehiSupport Engineer![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_200/Joshua_O_129141412f)](https://posthog.com/community/profiles/32997)### [Ian VanagasTechnical Content Marketer![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_200/ian_31bf16ef7d_c3339bc255)](https://posthog.com/community/profiles/29296)+8 more
### Feature availability
Free / Open-source![Available](https://posthog.com/docs/product-analytics/autocapture)
Self-serve![Available](https://posthog.com/docs/product-analytics/autocapture)
Enterprise![Available](https://posthog.com/docs/product-analytics/autocapture)
#### Jump to:
  * Types of autocaptured events
  * Supported SDKs
  * Interaction autocapture
  * Web interaction autocapture
  * iOS interaction autocapture
  * React Native interaction autocapture
  * Navigation and lifecycle event autocapture
  * Web navigation autocapture
  * iOS navigation and lifecycle autocapture
  * React Native navigation and lifecycle autocapture
  * Clipboard autocapture
  * Heatmap autocapture
  * Dead clicks autocapture
  * Analyzing autocaptured events and properties
  * Captured properties
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
- URL: https://posthog.com/docs/product-analytics/autocapture
- Harvested: 2025-08-19T22:10:02.664992
- Profile: deep_research
- Agent: architect
