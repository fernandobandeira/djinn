---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:38:36.077171'
profile: code_examples
source: https://posthog.com/tutorials/bootstrap-feature-flags-react
topic: PostHog React Feature Flag Bootstrapping Implementation
---

# PostHog React Feature Flag Bootstrapping Implementation

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
  * Feature flags


# How to bootstrap feature flags in React and Express
Mar 06, 2025
Posted by
  * [![](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_50/ian_31bf16ef7d_c3339bc255)Ian Vanagas](https://posthog.com/community/profiles/29296)


On this page
  * Create a React app with Vite and add PostHog
  * Create and setup a feature flag
  * Set up the React app for server-side rendering
  * Set up our server-rendering Express app
  * Bootstrapping the feature flags on the client
  * Further reading


Bootstrapping feature flags makes them available as soon as React and PostHog load on the client side. This enables use cases like routing to different pages on load, all feature flagged content being available on first load, and visual consistency.
To show you how you can bootstrap feature flags, we are going to build a React app with Vite, add PostHog, set up an Express server to server-side render our app, and finally bootstrap our flags from the server to the client.
> Already have an app set up? [Skip straight to the feature flag bootstrapping implementation](https://posthog.com/tutorials/bootstrap-feature-flags-react#handle-feature-flags-on-the-backend).
## Create a React app with Vite and add PostHog
Make sure you have [Node installed](https://nodejs.dev/en/learn/how-to-install-nodejs/), then create a new React app with Vite:
Terminal
```

npm create vite@latest client -- --template react

```

Once created, go into the new `client` folder and install the packages as well as `posthog-js` and its React wrapper:
Terminal
```

cd client
npminstall
npm i posthog-js

```

Next, get your PostHog project API key and instance address from the getting started flow or [your project settings](https://app.posthog.com/project/settings) and set up environment variables to store them. You can do this by creating a `.env.local` file in your project root:
Terminal
```

# .env.local
VITE_PUBLIC_POSTHOG_KEY=<ph_project_api_key>
VITE_PUBLIC_POSTHOG_HOST=https://us.i.posthog.com

```

Next, create your entry point for client-side rendering in `src/entry-client.jsx`:
JSX
```

// src/entry-client.jsx
import{StrictMode}from'react'
import{ createRoot }from'react-dom/client'
import'./index.css'
importAppfrom'./App.jsx'
import{PostHogProvider}from'posthog-js/react'
const options ={
api_host:import.meta.env.VITE_PUBLIC_POSTHOG_HOST,
}
createRoot(document.getElementById('root')).render(
<StrictMode>
<PostHogProviderapiKey={import.meta.env.VITE_PUBLIC_POSTHOG_KEY}options={options}>
<App/>
</PostHogProvider>
</StrictMode>,
)

```

Update your `index.html` to point to this file:
HTML
```

<!DOCTYPEhtml>
<htmllang="en">
<head>
<metacharset="UTF-8"/>
<linkrel="icon"type="image/svg+xml"href="/vite.svg"/>
<metaname="viewport"content="width=device-width, initial-scale=1.0"/>
<title>Vite + React</title>
</head>
<body>
<divid="root"></div>
<scripttype="module"src="/src/entry-client.jsx"></script>
</body>
</html>

```

If you want, you can run `npm run dev` to see the app in action.
## Create and setup a feature flag
If we want to bootstrap a feature flag, we first need to create it in PostHog. To do this, go to the [feature flag tab](https://us.posthog.com/feature_flags?tab=overview), create a new flag, set a key (I chose `test-flag`), set the rollout to 100% of users, and save it. 
Once done, you can evaluate your flag in the `loaded()` method on initialization like this:
JavaScript
```

const options ={
api_host:import.meta.env.VITE_PUBLIC_POSTHOG_HOST,
loaded(ph){
console.log(ph.isFeatureEnabled('test-flag'))
}
}

```

This shows us bootstrapping is valuable. On the first load of the site (before the flag is set in cookies), you see `undefined` in the console even though the flag should return `true`. This is because the flag isn't loaded yet when you check it, and the flag might not show the right code on the initial load for that user.
![Undefined](https://res.cloudinary.com/dmukukwp6/image/upload/Clean_Shot_2025_03_06_at_11_06_19_2x_f9bda663bb.png)
Bootstrapping flags solves this.
## Set up the React app for server-side rendering
To bootstrap our flags, we fetch the feature flag data on the backend and pass it to the frontend before it loads. This requires server-side rendering our React app. 
To do this with Vite, we need:
  1. A server entry point for rendering React on the server
  2. A client entry point for hydrating the app in the browser
  3. An Express server to get feature flags from PostHog and serve the React app


We'll start with the server entry point by creating `src/entry-server.jsx`:
JSX
```

// src/entry-server.jsx
importReactfrom'react'
import{ renderToString }from'react-dom/server'
importAppfrom'./App'
exportfunctionrender(){
const html =renderToString(
<React.StrictMode>
<App/>
</React.StrictMode>
)
return html
}

```

Next, modify your client entry point to support hydration in `src/entry-client.jsx`:
JSX
```

// src/entry-client.jsx
import{StrictMode}from'react'
import{ hydrateRoot }from'react-dom/client'
import'./index.css'
importAppfrom'./App'
import{PostHogProvider}from'posthog-js/react'
const options ={
api_host:import.meta.env.VITE_PUBLIC_POSTHOG_HOST,
loaded(ph){
console.log(ph.isFeatureEnabled('test-flag'))
}
}
// Use hydrateRoot instead of createRoot for SSR
hydrateRoot(
document.getElementById('root'),
<StrictMode>
<PostHogProviderapiKey={import.meta.env.VITE_PUBLIC_POSTHOG_KEY}options={options}>
<App/>
</PostHogProvider>
</StrictMode>
)

```

With this done, we can move on to setting up our server-rendering Express app.
## Set up our server-rendering Express app
To get the feature flags data on the backend and pass it to the frontend, we need to set up an Express server that:
  1. Gets or creates a distinct ID for PostHog
  2. Uses it to get the feature flags from PostHog
  3. Injects the feature flags data into the HTML
  4. Sends the HTML back to the client


This starts by installing the necessary packages:
Terminal
```

npminstall express cookie-parser posthog-node uuid dotenv
npminstall --save-dev nodemon

```

Next, create a server directory in the root of your project and a `index.js` file inside it:
Terminal
```

mkdir server
touch server/index.js

```

In this file, start by importing everything we need, setting up the environment variables, and initializing the PostHog client:
JavaScript
```

// server/index.js
importexpressfrom'express';
importfsfrom'fs';
importpathfrom'path';
import{ fileURLToPath }from'url';
import{ createServer as createViteServer }from'vite';
importcookieParserfrom'cookie-parser';
import{ v4 as uuidv4 }from'uuid';
import{PostHog}from'posthog-node';
// Import environment variables
importdotenvfrom'dotenv';
dotenv.config();
const __dirname = path.dirname(fileURLToPath(import.meta.url));
// Initialize PostHog client
const client =newPostHog(
 process.env.VITE_PUBLIC_POSTHOG_KEY,
{
host: process.env.VITE_PUBLIC_POSTHOG_HOST,
personalApiKey: process.env.POSTHOG_PERSONAL_API_KEY// This one is server-only
}
);

```

Next, create a function to create and start the server:
JavaScript
```

// ... existing code
asyncfunctioncreateServer(){
const app =express();
// Use cookie parser middleware
 app.use(cookieParser());
// Create Vite server in middleware mode
const vite =awaitcreateViteServer({
server:{middlewareMode:true},
appType:'custom'
});
// Use vite's connect instance as middleware
 app.use(vite.middlewares);
 app.use('*',async(req, res, next)=>{
const url = req.originalUrl;
try{
// More code here soon...
}catch(e){
// If an error is caught, let Vite fix the stack trace for better debugging
   vite.ssrFixStacktrace(e);
console.error(e);
next(e);
}
});
 app.listen(3000,()=>{
console.log('Server started at http://localhost:3000');
});
}
createServer();

```

In the route's `try` block, we'll get or create a distinct ID and use it to get the feature flags:
JavaScript
```

try{
// Get or create distinct ID
let distinctId =null;
const phCookie = req.cookies[`ph_${process.env.VITE_PUBLIC_POSTHOG_KEY}_posthog`];
if(phCookie){
  distinctId =JSON.parse(phCookie)['distinct_id'];
}
if(!distinctId){
  distinctId =uuidv4();
}
// Get all feature flags for this user
const flags =await client.getAllFlags(distinctId);
// More code here soon...

```

Once we have them, we'll inject them into the HTML and send it back to the client.
JavaScript
```

// ... existing code
// 1. Read index.html
let template = fs.readFileSync(
  path.resolve(__dirname,'../index.html'),
'utf-8'
);
// 2. Apply Vite HTML transforms
 template =await vite.transformIndexHtml(url, template);
// 3. Load the server entry
const{ render }=await vite.ssrLoadModule('/src/entry-server.jsx');
// 4. Render the app HTML
const appHtml =awaitrender(url);
// 5. Inject the app-rendered HTML and feature flag data into the template
const serializedFlags =JSON.stringify(flags);
const serializedDistinctId =JSON.stringify(distinctId);
const scriptTag =`<script>window.__FLAG_DATA__ = ${serializedFlags}; window.__PH_DISTINCT_ID__ = ${serializedDistinctId};</script>`;
const html = template
.replace(`<div id="root"></div>`,`<div id="root">${appHtml}</div>`)
.replace('</head>',`${scriptTag}</head>`);
// 6. Send the rendered HTML back
 res.status(200).set({'Content-Type':'text/html'}).end(html);
}catch(e){
// ... existing code

```

See the full index.js file
JavaScript
```

// server/index.js
importexpressfrom'express';
importfsfrom'fs';
importpathfrom'path';
import{ fileURLToPath }from'url';
import{ createServer as createViteServer }from'vite';
importcookieParserfrom'cookie-parser';
import{ v4 as uuidv4 }from'uuid';
import{PostHog}from'posthog-node';
// Import environment variables
importdotenvfrom'dotenv';
dotenv.config();
const __dirname = path.dirname(fileURLToPath(import.meta.url));
// Initialize PostHog client
const client =newPostHog(
 process.env.VITE_PUBLIC_POSTHOG_KEY,
{
host: process.env.VITE_PUBLIC_POSTHOG_HOST,
personalApiKey: process.env.POSTHOG_PERSONAL_API_KEY// This one is server-only
}
);
asyncfunctioncreateServer(){
const app =express();
// Use cookie parser middleware
 app.use(cookieParser());
// Create Vite server in middleware mode
const vite =awaitcreateViteServer({
server:{middlewareMode:true},
appType:'custom'
});
// Use vite's connect instance as middleware
 app.use(vite.middlewares);
 app.use('*',async(req, res, next)=>{
const url = req.originalUrl;
try{
// Get or create distinct ID
let distinctId =null;
const phCookie = req.cookies[`ph_${process.env.VITE_PUBLIC_POSTHOG_KEY}_posthog`];
if(phCookie){
    distinctId =JSON.parse(phCookie)['distinct_id'];
}
if(!distinctId){
    distinctId =uuidv4();
}
// Get all feature flags for this user
const flags =await client.getAllFlags(distinctId);
// 1. Read index.html
let template = fs.readFileSync(
    path.resolve(__dirname,'../index.html'),
'utf-8'
);
// 2. Apply Vite HTML transforms
   template =await vite.transformIndexHtml(url, template);
// 3. Load the server entry
const{ render }=await vite.ssrLoadModule('/src/entry-server.jsx');
// 4. Render the app HTML
const appHtml =awaitrender(url);
// 5. Inject the app-rendered HTML and feature flag data into the template
const serializedFlags =JSON.stringify(flags);
const serializedDistinctId =JSON.stringify(distinctId);
const scriptTag =`<script>window.__FLAG_DATA__ = ${serializedFlags}; window.__PH_DISTINCT_ID__ = ${serializedDistinctId};</script>`;
const html = template
.replace(`<div id="root"></div>`,`<div id="root">${appHtml}</div>`)
.replace('</head>',`${scriptTag}</head>`);
// 6. Send the rendered HTML back
   res.status(200).set({'Content-Type':'text/html'}).end(html);
}catch(e){
// If an error is caught, let Vite fix the stack trace for better debugging
   vite.ssrFixStacktrace(e);
console.error(e);
next(e);
}
});
 app.listen(3000,()=>{
console.log('Server started at http://localhost:3000');
});
}
createServer();

```

Once you got this all set up, you need a PostHog personal API key. To get one, go to [your user settings](https://us.posthog.com/settings/user-api-keys), click **Personal API keys** , then **Create personal API key** , select **All access** , and then select the **Local feature flag evaluation** preset.
![Creating a personal API key in PostHog](https://res.cloudinary.com/dmukukwp6/image/upload/Clean_Shot_2025_03_06_at_11_15_57_2x_6ef86dd370.png)![Creating a personal API key in PostHog](https://res.cloudinary.com/dmukukwp6/image/upload/Clean_Shot_2025_03_06_at_11_15_40_2x_1e5809b616.png)
Add it to your `.env.local` file:
Terminal
```

# .env.local
# ... rest of your environment variables
POSTHOG_PERSONAL_API_KEY=phx_your-personal-api-key

```

Your React app will now be server-side rendered with the feature flags data injected into the HTML.
## Bootstrapping the feature flags on the client
The last thing we need to do is bootstrap the feature flags on the client. To do this, we'll update our client entry point to use the bootstrapped data:
JSX
```

// src/entry-client.jsx
import{StrictMode}from'react'
import{ hydrateRoot }from'react-dom/client'
import'./index.css'
importAppfrom'./App'
import{PostHogProvider}from'posthog-js/react'
// Get bootstrapped data from window
const flagData =window.__FLAG_DATA__;
const distinctId =window.__PH_DISTINCT_ID__;
const options ={
api_host:import.meta.env.VITE_PUBLIC_POSTHOG_HOST,
bootstrap:{
distinctID: distinctId,
featureFlags: flagData,
},
loaded(ph){
console.log(ph.isFeatureEnabled('test-flag'))
}
}
// rest of your code...

```

Once this is done, we can run the server:
Terminal
```

nodemon --watch server --watch src/entry-server.jsx server/index.js

```

When you visit `http://localhost:3000`, you should see that feature flags are loaded immediately on the first page load. Open up the site on an incognito or guest window, and you'll see that the flag returns `true` on the first load without any delay.
![It's working](https://res.cloudinary.com/dmukukwp6/image/upload/Clean_Shot_2025_03_06_at_11_36_43_2x_476a2f9c96.png)
This is feature flag bootstrapping working successfully. From here, you can make the flag redirect to specific pages, control session recordings, or run an A/B test on your home page call to action.
## Further reading
  * [How to add popups to your React app with feature flags](https://posthog.com/tutorials/react-popups)
  * [Testing frontend feature flags with React, Jest, and PostHog](https://posthog.com/tutorials/test-frontend-feature-flags)
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
  * Create a React app with Vite and add PostHog
  * Create and setup a feature flag
  * Set up the React app for server-side rendering
  * Set up our server-rendering Express app
  * Bootstrapping the feature flags on the client
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


**PostHog.com doesn't use third party cookies** - only a single in-house cookie.
No data is sent to a third party.
AcceptDecline
![Ursula von der Leyen, President of the European Commission](https://res.cloudinary.com/dmukukwp6/image/upload/c_scale,w_250/v1/posthog.com/src/components/EU/images/ursula)
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
- URL: https://posthog.com/tutorials/bootstrap-feature-flags-react
- Harvested: 2025-08-20T00:38:36.077171
- Profile: code_examples
- Agent: architect
