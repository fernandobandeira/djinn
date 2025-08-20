---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:21:39.504083'
profile: code_examples
source: https://docs.flutter.dev/cookbook/testing/widget/golden-tests
topic: Golden Testing for UI Regression in Flutter
---

# Golden Testing for UI Regression in Flutter

docs.flutter.dev uses cookies from Google to deliver and enhance the quality of its services and to analyze traffic. [Learn more](https://policies.google.com/technologies/cookies).
OK, got it
[ Skip to main content ](https://docs.flutter.dev/cookbook/testing/widget/golden-tests#document-title)[ ![Flutter logo](https://docs.flutter.dev/assets/images/branding/flutter/logo/default.svg) Flutter Docs ](https://docs.flutter.dev/ "Go to the Flutter docs homepage")
[ search ](https://docs.flutter.dev/search "Navigate to the docs search page.")
routine
  * light_mode Light
  * dark_mode Dark
  * night_sight_auto Automatic


apps
  * [ ![Flutter logo](https://docs.flutter.dev/assets/images/branding/flutter/logo/default.svg) Flutter ](https://flutter.dev "Flutter homepage")
  * [ ![Flutter logo](https://docs.flutter.dev/assets/images/branding/flutter/logo/default.svg) Flutter Docs ](https://docs.flutter.dev/ "Flutter docs homepage")
  * [ ![Flutter logo](https://docs.flutter.dev/assets/images/branding/flutter/logo/default.svg) Flutter API ](https://api.flutter.dev "Flutter API reference")
  * [ ![Dart logo](https://docs.flutter.dev/assets/images/branding/dart/logo.svg) Dart ](https://dart.dev "Dart homepage")
  * [ ![Dart logo](https://docs.flutter.dev/assets/images/branding/dart/logo.svg) DartPad ](https://dartpad.dev "DartPad playground")
  * [ ![Dart logo](https://docs.flutter.dev/assets/images/branding/dart/logo.svg) pub.dev ](https://pub.dev "pub.dev homepage")


[Get started](https://docs.flutter.dev/get-started/install/) menu close
  * Get started
  * [Set up Flutter](https://docs.flutter.dev/get-started/install)
  * Install Flutter expand_more
    * [Overview](https://docs.flutter.dev/install)
    * [Install with VS Code](https://docs.flutter.dev/install/with-vs-code)
    * [Install manually](https://docs.flutter.dev/install/manual)
    * [Upgrade SDK](https://docs.flutter.dev/install/upgrade)
    * [SDK archive](https://docs.flutter.dev/install/archive)
    * [Add to path](https://docs.flutter.dev/install/add-to-path)
    * [Troubleshoot](https://docs.flutter.dev/install/troubleshoot)
    * [Uninstall SDK](https://docs.flutter.dev/install/uninstall)
  * Learn Flutter expand_more
    * [Introduction](https://docs.flutter.dev/get-started/learn-flutter)
    * [Write your first app](https://docs.flutter.dev/get-started/codelab)
    * Learn the fundamentals expand_more
      * [Introduction](https://docs.flutter.dev/get-started/fundamentals)
      * [Intro to Dart](https://docs.flutter.dev/get-started/fundamentals/dart)
      * [Widgets](https://docs.flutter.dev/get-started/fundamentals/widgets)
      * [Layout](https://docs.flutter.dev/get-started/fundamentals/layout)
      * [State management](https://docs.flutter.dev/get-started/fundamentals/state-management)
      * [Handling user input](https://docs.flutter.dev/get-started/fundamentals/user-input)
      * [Networking and data](https://docs.flutter.dev/get-started/fundamentals/networking)
      * [Local data and caching](https://docs.flutter.dev/get-started/fundamentals/local-caching)
    * From another platform? expand_more
      * [Flutter for Android devs](https://docs.flutter.dev/get-started/flutter-for/android-devs)
      * [Flutter for Jetpack Compose devs](https://docs.flutter.dev/get-started/flutter-for/compose-devs)
      * [Flutter for SwiftUI devs](https://docs.flutter.dev/get-started/flutter-for/swiftui-devs)
      * [Flutter for UIKit devs](https://docs.flutter.dev/get-started/flutter-for/uikit-devs)
      * [Flutter for React Native devs](https://docs.flutter.dev/get-started/flutter-for/react-native-devs)
      * [Flutter for web devs](https://docs.flutter.dev/get-started/flutter-for/web-devs)
      * [Flutter for Xamarin.Forms devs](https://docs.flutter.dev/get-started/flutter-for/xamarin-forms-devs)
      * [Introduction to declarative UI](https://docs.flutter.dev/get-started/flutter-for/declarative)
      * [Flutter versus Swift concurrency](https://docs.flutter.dev/get-started/flutter-for/dart-swift-concurrency)
    * [Samples & tutorials](https://docs.flutter.dev/reference/learning-resources)
  * Stay up to date expand_more
    * [Release notes](https://docs.flutter.dev/release/release-notes)
    * [Breaking changes](https://docs.flutter.dev/release/breaking-changes)
    * [Compatibility policy](https://docs.flutter.dev/release/compatibility-policy)
    * [Medium publicationopen_in_new](https://medium.com/flutter)
    * [What's new in the docs](https://docs.flutter.dev/release/whats-new)
  * App solutions expand_more
    * Develop with Firebase expand_more
      * [Overview](https://docs.flutter.dev/data-and-backend/firebase)
      * [Discover Firebase for Flutteropen_in_new](https://firebase.google.com/docs/flutter)
      * [Get to know Firebase for Flutteropen_in_new](https://www.youtube.com/watch?v=wUSkeTaBonA)
      * [Add a user authentication flow to a Flutter app using FirebaseUIopen_in_new](https://firebase.google.com/codelabs/firebase-auth-in-flutter-apps)
      * [Get to know Firebase for webopen_in_new](https://firebase.google.com/codelabs/firebase-get-to-know-web)
    * Build multi-platform games expand_more
      * [Overview](https://docs.flutter.dev/resources/games-toolkit)
      * [Add achievements and leaderboards](https://docs.flutter.dev/cookbook/games/achievements-leaderboard)
      * [Build leaderboards with Firestoreopen_in_new](https://firebase.google.com/codelabs/build-leaderboards-with-firestore#0)
      * [Add advertising](https://docs.flutter.dev/cookbook/plugins/google-mobile-ads)
      * [Add multiplayer support](https://docs.flutter.dev/cookbook/games/firestore-multiplayer)
      * [Add in-app purchasesopen_in_new](https://codelabs.developers.google.com/codelabs/flutter-in-app-purchases)
      * [Add user authenticationopen_in_new](https://firebase.google.com/codelabs/firebase-auth-in-flutter-apps)
      * [Debug using Crashlyticsopen_in_new](https://firebase.google.com/docs/crashlytics/get-started?platform=flutter)
      * [Intro to Flame with Flutteropen_in_new](https://codelabs.developers.google.com/codelabs/flutter-flame-brick-breaker)
    * Monetize your app expand_more
      * Integrate ads expand_more
        * [Ads overview](https://docs.flutter.dev/resources/ads-overview)
        * [Add advertising](https://docs.flutter.dev/cookbook/plugins/google-mobile-ads)
        * [Add AdMob ads to your Flutter appopen_in_new](https://codelabs.developers.google.com/codelabs/admob-ads-in-flutter)
        * [Add an AdMob banner and native inline adsopen_in_new](https://codelabs.developers.google.com/codelabs/admob-inline-ads-in-flutter)
        * [Integrate multimedia ads (video)open_in_new](https://www.youtube.com/watch?v=U8x5n6RwZOo)
        * [Google AdMob mediationopen_in_new](https://developers.google.com/admob/flutter/mediation)
        * [Interactive Media Ads SDKopen_in_new](https://pub.dev/packages/interactive_media_ads)
      * Support payments expand_more
        * [Payments overview](https://docs.flutter.dev/resources/payments-overview)
        * [Google pay packageopen_in_new](https://pub.dev/packages/pay)
      * [Add in-app purchasesopen_in_new](https://codelabs.developers.google.com/codelabs/flutter-in-app-purchases)
    * Integrate maps expand_more
      * [Add Google maps to a Flutter appopen_in_new](https://codelabs.developers.google.com/codelabs/google-maps-in-flutter)
      * [Google Maps packageopen_in_new](https://developers.google.com/maps/flutter-package)
    * [Build a news app](https://docs.flutter.dev/resources/news-toolkit)
  * Create with AI expand_more
    * [Overview](https://docs.flutter.dev/ai/create-with-ai)
    * AI Toolkit expand_more
      * [Overview](https://docs.flutter.dev/ai-toolkit)
      * [User experience](https://docs.flutter.dev/ai-toolkit/user-experience)
      * [Feature integration](https://docs.flutter.dev/ai-toolkit/feature-integration)
      * [Custom LLM providers](https://docs.flutter.dev/ai-toolkit/custom-llm-providers)
      * [Chat client sample](https://docs.flutter.dev/ai-toolkit/chat-client-sample)
    * [Vertex AI in Firebaseopen_in_new](https://firebase.google.com/docs/vertex-ai/get-started?platform=flutter)
  * User interface
  * [Introduction](https://docs.flutter.dev/ui)
  * [Widget catalog](https://docs.flutter.dev/ui/widgets)
  * Layout expand_more
    * [Introduction](https://docs.flutter.dev/ui/layout)
    * [Build a layout](https://docs.flutter.dev/ui/layout/tutorial)
    * Lists & grids expand_more
      * [Create and use lists](https://docs.flutter.dev/cookbook/lists/basic-list)
      * [Create a horizontal list](https://docs.flutter.dev/cookbook/lists/horizontal-list)
      * [Create a grid view](https://docs.flutter.dev/cookbook/lists/grid-lists)
      * [Create lists with different types of items](https://docs.flutter.dev/cookbook/lists/mixed-list)
      * [Create lists with spaced items](https://docs.flutter.dev/cookbook/lists/spaced-items)
      * [Work with long lists](https://docs.flutter.dev/cookbook/lists/long-lists)
    * Scrolling expand_more
      * [Overview](https://docs.flutter.dev/ui/layout/scrolling)
      * [Use slivers to achieve fancy scrolling](https://docs.flutter.dev/ui/layout/scrolling/slivers)
      * [Place a floating app bar above a list](https://docs.flutter.dev/cookbook/lists/floating-app-bar)
      * [Create a scrolling parallax effect](https://docs.flutter.dev/cookbook/effects/parallax-scrolling)
  * Adaptive & responsive design expand_more
    * [Overview](https://docs.flutter.dev/ui/adaptive-responsive)
    * [General approach](https://docs.flutter.dev/ui/adaptive-responsive/general)
    * [SafeArea & MediaQuery](https://docs.flutter.dev/ui/adaptive-responsive/safearea-mediaquery)
    * [Large screens & foldables](https://docs.flutter.dev/ui/adaptive-responsive/large-screens)
    * [User input & accessibility](https://docs.flutter.dev/ui/adaptive-responsive/input)
    * [Capabilities & policies](https://docs.flutter.dev/ui/adaptive-responsive/capabilities)
    * [Automatic platform adaptations](https://docs.flutter.dev/ui/adaptive-responsive/platform-adaptations)
    * [Best practices](https://docs.flutter.dev/ui/adaptive-responsive/best-practices)
    * [Additional resources](https://docs.flutter.dev/ui/adaptive-responsive/more-info)
  * Design & theming expand_more
    * [Share styles with themes](https://docs.flutter.dev/cookbook/design/themes)
    * [Material design](https://docs.flutter.dev/ui/design/material)
    * [Migrate to Material 3](https://docs.flutter.dev/release/breaking-changes/material-3-migration)
    * Text expand_more
      * [Fonts & typography](https://docs.flutter.dev/ui/design/text/typography)
      * [Use a custom font](https://docs.flutter.dev/cookbook/design/fonts)
      * [Export fonts from a package](https://docs.flutter.dev/cookbook/design/package-fonts)
      * [Google Fonts packageopen_in_new](https://pub.dev/packages/google_fonts)
    * Custom graphics expand_more
      * [Use custom fragment shaders](https://docs.flutter.dev/ui/design/graphics/fragment-shaders)
  * Interactivity expand_more
    * [Add interactivity to your app](https://docs.flutter.dev/ui/interactivity)
    * Gestures expand_more
      * [Introduction](https://docs.flutter.dev/ui/interactivity/gestures)
      * [Handle taps](https://docs.flutter.dev/cookbook/gestures/handling-taps)
      * [Drag an object outside an app](https://docs.flutter.dev/ui/interactivity/gestures/drag-outside)
      * [Drag a UI element within an app](https://docs.flutter.dev/cookbook/effects/drag-a-widget)
      * [Add Material touch ripples](https://docs.flutter.dev/cookbook/gestures/ripples)
      * [Implement swipe to dismiss](https://docs.flutter.dev/cookbook/gestures/dismissible)
    * Input & forms expand_more
      * [Create and style a text field](https://docs.flutter.dev/cookbook/forms/text-input)
      * [Retrieve the value of a text field](https://docs.flutter.dev/cookbook/forms/retrieve-input)
      * [Handle changes to a text field](https://docs.flutter.dev/cookbook/forms/text-field-changes)
      * [Manage focus in text fields](https://docs.flutter.dev/cookbook/forms/focus)
      * [Build a form with validation](https://docs.flutter.dev/cookbook/forms/validation)
    * [Display a snackbar](https://docs.flutter.dev/cookbook/design/snackbars)
    * [Implement actions & shortcuts](https://docs.flutter.dev/ui/interactivity/actions-and-shortcuts)
    * [Manage keyboard focus](https://docs.flutter.dev/ui/interactivity/focus)
  * Assets & media expand_more
    * [Add assets and images](https://docs.flutter.dev/ui/assets/assets-and-images)
    * [Display images from the internet](https://docs.flutter.dev/cookbook/images/network-image)
    * [Fade in images with a placeholder](https://docs.flutter.dev/cookbook/images/fading-in-images)
    * [Play and pause a video](https://docs.flutter.dev/cookbook/plugins/play-video)
    * [Transform assets at build time](https://docs.flutter.dev/ui/assets/asset-transformation)
  * Navigation & routing expand_more
    * [Overview](https://docs.flutter.dev/ui/navigation)
    * [Add tabs to your app](https://docs.flutter.dev/cookbook/design/tabs)
    * [Navigate to a new screen and back](https://docs.flutter.dev/cookbook/navigation/navigation-basics)
    * [Send data to a new screen](https://docs.flutter.dev/cookbook/navigation/passing-data)
    * [Return data from a screen](https://docs.flutter.dev/cookbook/navigation/returning-data)
    * [Add a drawer to a screen](https://docs.flutter.dev/cookbook/design/drawer)
    * [Set up deep linking](https://docs.flutter.dev/ui/navigation/deep-linking)
    * [Set up app links for Android](https://docs.flutter.dev/cookbook/navigation/set-up-app-links)
    * [Set up universal links for iOS](https://docs.flutter.dev/cookbook/navigation/set-up-universal-links)
    * [Configure web URL strategies](https://docs.flutter.dev/ui/navigation/url-strategies)
  * Animations & transitions expand_more
    * [Introduction](https://docs.flutter.dev/ui/animations)
    * [Tutorial](https://docs.flutter.dev/ui/animations/tutorial)
    * [Implicit animations](https://docs.flutter.dev/ui/animations/implicit-animations)
    * [Animate the properties of a container](https://docs.flutter.dev/cookbook/animation/animated-container)
    * [Fade a widget in and out](https://docs.flutter.dev/cookbook/animation/opacity-animation)
    * [Hero animations](https://docs.flutter.dev/ui/animations/hero-animations)
    * [Animate a page route transition](https://docs.flutter.dev/cookbook/animation/page-route-animation)
    * [Animate using a physics simulation](https://docs.flutter.dev/cookbook/animation/physics-simulation)
    * [Staggered animations](https://docs.flutter.dev/ui/animations/staggered-animations)
    * [Create a staggered menu animation](https://docs.flutter.dev/cookbook/effects/staggered-menu-animation)
    * [API overview](https://docs.flutter.dev/ui/animations/overview)
  * Accessibility & internationalization expand_more
    * [Accessibility](https://docs.flutter.dev/ui/accessibility-and-internationalization/accessibility)
    * [Internationalization](https://docs.flutter.dev/ui/accessibility-and-internationalization/internationalization)
  * Beyond UI
  * Data & backend expand_more
    * State management expand_more
      * [Introduction](https://docs.flutter.dev/data-and-backend/state-mgmt/intro)
      * [Think declaratively](https://docs.flutter.dev/data-and-backend/state-mgmt/declarative)
      * [Ephemeral vs app state](https://docs.flutter.dev/data-and-backend/state-mgmt/ephemeral-vs-app)
      * [Simple app state management](https://docs.flutter.dev/data-and-backend/state-mgmt/simple)
      * [Options](https://docs.flutter.dev/data-and-backend/state-mgmt/options)
    * Networking & http expand_more
      * [Overview](https://docs.flutter.dev/data-and-backend/networking)
      * [Fetch data from the internet](https://docs.flutter.dev/cookbook/networking/fetch-data)
      * [Make authenticated requests](https://docs.flutter.dev/cookbook/networking/authenticated-requests)
      * [Send data to the internet](https://docs.flutter.dev/cookbook/networking/send-data)
      * [Update data over the internet](https://docs.flutter.dev/cookbook/networking/update-data)
      * [Delete data on the internet](https://docs.flutter.dev/cookbook/networking/delete-data)
      * [Communicate with WebSockets](https://docs.flutter.dev/cookbook/networking/web-sockets)
    * Serialization expand_more
      * [JSON serialization](https://docs.flutter.dev/data-and-backend/serialization/json)
      * [Parse JSON in the background](https://docs.flutter.dev/cookbook/networking/background-parsing)
    * Persistence expand_more
      * [Store key-value data on disk](https://docs.flutter.dev/cookbook/persistence/key-value)
      * [Read and write files](https://docs.flutter.dev/cookbook/persistence/reading-writing-files)
      * [Persist data with SQLite](https://docs.flutter.dev/cookbook/persistence/sqlite)
    * Firebase expand_more
      * [Overview](https://docs.flutter.dev/data-and-backend/firebase)
      * [Add Firebase to your Flutter appopen_in_new](https://firebase.google.com/docs/flutter/setup)
    * [Google APIs](https://docs.flutter.dev/data-and-backend/google-apis)
  * App architecture expand_more
    * [Introduction](https://docs.flutter.dev/app-architecture)
    * [Architecture concepts](https://docs.flutter.dev/app-architecture/concepts)
    * [Guide to app architecture](https://docs.flutter.dev/app-architecture/guide)
    * Architecture case study expand_more
      * [Overview](https://docs.flutter.dev/app-architecture/case-study)
      * [UI layer](https://docs.flutter.dev/app-architecture/case-study/ui-layer)
      * [Data layer](https://docs.flutter.dev/app-architecture/case-study/data-layer)
      * [Dependency injection](https://docs.flutter.dev/app-architecture/case-study/dependency-injection)
      * [Testing each layer](https://docs.flutter.dev/app-architecture/case-study/testing)
    * [Recommendations](https://docs.flutter.dev/app-architecture/recommendations)
    * [Design patterns](https://docs.flutter.dev/app-architecture/design-patterns)
  * Platform integration expand_more
    * [Supported platforms](https://docs.flutter.dev/reference/supported-platforms)
    * [Build desktop apps with Flutter](https://docs.flutter.dev/platform-integration/desktop)
    * [Write platform-specific code](https://docs.flutter.dev/platform-integration/platform-channels)
    * Android expand_more
      * [Set up Android development](https://docs.flutter.dev/platform-integration/android/setup)
      * [Add a splash screen](https://docs.flutter.dev/platform-integration/android/splash-screen)
      * [Add predictive back](https://docs.flutter.dev/platform-integration/android/predictive-back)
      * [Bind to native code](https://docs.flutter.dev/platform-integration/android/c-interop)
      * [Host a native Android view](https://docs.flutter.dev/platform-integration/android/platform-views)
      * [Calling JetPack APIs](https://docs.flutter.dev/platform-integration/android/call-jetpack-apis)
      * [Launch a Jetpack Compose activity](https://docs.flutter.dev/platform-integration/android/compose-activity)
      * [Restore state on Android](https://docs.flutter.dev/platform-integration/android/restore-state-android)
      * [Target ChromeOS with Android](https://docs.flutter.dev/platform-integration/android/chromeos)
      * [Protect your app's sensitive content](https://docs.flutter.dev/platform-integration/android/sensitive-content)
    * iOS expand_more
      * [Set up iOS development](https://docs.flutter.dev/platform-integration/ios/setup)
      * [Flutter on latest iOS](https://docs.flutter.dev/platform-integration/ios/ios-latest)
      * [Leverage Apple's system libraries](https://docs.flutter.dev/platform-integration/ios/apple-frameworks)
      * [Add a launch screen](https://docs.flutter.dev/platform-integration/ios/launch-screen)
      * [Add iOS App Clip support](https://docs.flutter.dev/platform-integration/ios/ios-app-clip)
      * [Add iOS app extensions](https://docs.flutter.dev/platform-integration/ios/app-extensions)
      * [Bind to native code](https://docs.flutter.dev/platform-integration/ios/c-interop)
      * [Host a native iOS view](https://docs.flutter.dev/platform-integration/ios/platform-views)
      * [Enable debugging on iOS](https://docs.flutter.dev/platform-integration/ios/ios-debugging)
      * [Restore state on iOS](https://docs.flutter.dev/platform-integration/ios/restore-state-ios)
    * Linux expand_more
      * [Set up Linux development](https://docs.flutter.dev/platform-integration/linux/setup)
      * [Build a Linux app](https://docs.flutter.dev/platform-integration/linux/building)
    * macOS expand_more
      * [Set up macOS development](https://docs.flutter.dev/platform-integration/macos/setup)
      * [Build a macOS app](https://docs.flutter.dev/platform-integration/macos/building)
      * [Bind to native code](https://docs.flutter.dev/platform-integration/macos/c-interop)
      * [Host a native macOS view](https://docs.flutter.dev/platform-integration/macos/platform-views)
    * Web expand_more
      * [Web support in Flutter](https://docs.flutter.dev/platform-integration/web)
      * [Set up web development](https://docs.flutter.dev/platform-integration/web/setup)
      * [Build a web app](https://docs.flutter.dev/platform-integration/web/building)
      * [Compile to WebAssembly](https://docs.flutter.dev/platform-integration/web/wasm)
      * [Customize app initialization](https://docs.flutter.dev/platform-integration/web/initialization)
      * [Add Flutter to any web app](https://docs.flutter.dev/platform-integration/web/embedding-flutter-web)
      * [Web content in Flutter](https://docs.flutter.dev/platform-integration/web/web-content-in-flutter)
      * [Web renderers](https://docs.flutter.dev/platform-integration/web/renderers)
      * [Display images on the web](https://docs.flutter.dev/platform-integration/web/web-images)
      * [Web FAQ](https://docs.flutter.dev/platform-integration/web/faq)
    * Windows expand_more
      * [Set up Windows development](https://docs.flutter.dev/platform-integration/windows/setup)
      * [Build a Windows app](https://docs.flutter.dev/platform-integration/windows/building)
  * Packages & plugins expand_more
    * [Use packages & plugins](https://docs.flutter.dev/packages-and-plugins/using-packages)
    * [Develop packages & plugins](https://docs.flutter.dev/packages-and-plugins/developing-packages)
    * Swift Package Manager expand_more
      * [For app developers](https://docs.flutter.dev/packages-and-plugins/swift-package-manager/for-app-developers)
      * [For plugin authors](https://docs.flutter.dev/packages-and-plugins/swift-package-manager/for-plugin-authors)
    * [Flutter Favorites](https://docs.flutter.dev/packages-and-plugins/favorites)
    * [Package repositoryopen_in_new](https://pub.dev/flutter)
  * Testing & debugging expand_more
    * Testing
    * [Overview](https://docs.flutter.dev/testing/overview)
    * Unit testing expand_more
      * [Introduction](https://docs.flutter.dev/cookbook/testing/unit/introduction)
      * [Mock dependencies](https://docs.flutter.dev/cookbook/testing/unit/mocking)
    * Widget testing expand_more
      * [Introduction](https://docs.flutter.dev/cookbook/testing/widget/introduction)
      * [Find widgets](https://docs.flutter.dev/cookbook/testing/widget/finders)
      * [Simulate scrolling](https://docs.flutter.dev/cookbook/testing/widget/scrolling)
      * [Simulate user interaction](https://docs.flutter.dev/cookbook/testing/widget/tap-drag)
    * Integration testing expand_more
      * [Introduction](https://docs.flutter.dev/cookbook/testing/integration/introduction)
      * [Write and run an integration test](https://docs.flutter.dev/testing/integration-tests)
      * [Profile an integration test](https://docs.flutter.dev/cookbook/testing/integration/profiling)
    * [Test a plugin](https://docs.flutter.dev/testing/testing-plugins)
    * [Handle plugin code in tests](https://docs.flutter.dev/testing/plugins-in-tests)
    * Debugging
    * [Debugging tools](https://docs.flutter.dev/testing/debugging)
    * [Debug your app programmatically](https://docs.flutter.dev/testing/code-debugging)
    * [Use a native language debugger](https://docs.flutter.dev/testing/native-debugging)
    * [Common Flutter errors](https://docs.flutter.dev/testing/common-errors)
    * [Handle errors](https://docs.flutter.dev/testing/errors)
    * [Report errors to a service](https://docs.flutter.dev/cookbook/maintenance/error-reporting)
  * Performance & optimization expand_more
    * [Overview](https://docs.flutter.dev/perf)
    * [Impeller](https://docs.flutter.dev/perf/impeller)
    * [Performance best practices](https://docs.flutter.dev/perf/best-practices)
    * [App size](https://docs.flutter.dev/perf/app-size)
    * [Deferred components](https://docs.flutter.dev/perf/deferred-components)
    * [Rendering performance](https://docs.flutter.dev/perf/rendering-performance)
    * [Performance profiling](https://docs.flutter.dev/perf/ui-performance)
    * [Performance profiling for web](https://docs.flutter.dev/perf/web-performance)
    * [Performance metrics](https://docs.flutter.dev/perf/metrics)
    * [Concurrency and isolates](https://docs.flutter.dev/perf/isolates)
    * [Performance FAQ](https://docs.flutter.dev/perf/faq)
    * [Appendix](https://docs.flutter.dev/perf/appendix)
  * Deployment expand_more
    * [Obfuscate Dart code](https://docs.flutter.dev/deployment/obfuscate)
    * [Create app flavors for Android](https://docs.flutter.dev/deployment/flavors)
    * [Create app flavors for iOS and macOS](https://docs.flutter.dev/deployment/flavors-ios)
    * [Build and release an Android app](https://docs.flutter.dev/deployment/android)
    * [Build and release an iOS app](https://docs.flutter.dev/deployment/ios)
    * [Build and release a macOS app](https://docs.flutter.dev/deployment/macos)
    * [Build and release a Linux app](https://docs.flutter.dev/deployment/linux)
    * [Build and release a Windows app](https://docs.flutter.dev/deployment/windows)
    * [Build and release a web app](https://docs.flutter.dev/deployment/web)
    * [Set up continuous deployment](https://docs.flutter.dev/deployment/cd)
  * Add to an existing app expand_more
    * [Introduction](https://docs.flutter.dev/add-to-app)
    * Add to an Android app expand_more
      * [Set up Android project](https://docs.flutter.dev/add-to-app/android/project-setup)
      * [Add a single Flutter screen](https://docs.flutter.dev/add-to-app/android/add-flutter-screen)
      * [Add a Flutter Fragment](https://docs.flutter.dev/add-to-app/android/add-flutter-fragment)
      * [Add a Flutter View](https://docs.flutter.dev/add-to-app/android/add-flutter-view)
      * [Use a Flutter plugin](https://docs.flutter.dev/add-to-app/android/plugin-setup)
    * Add to an iOS app expand_more
      * [Set up iOS project](https://docs.flutter.dev/add-to-app/ios/project-setup)
      * [Add a single Flutter screen](https://docs.flutter.dev/add-to-app/ios/add-flutter-screen)
    * [Add to a web app](https://docs.flutter.dev/platform-integration/web/embedding-flutter-web)
    * [Debug embedded Flutter module](https://docs.flutter.dev/add-to-app/debugging)
    * [Add multiple Flutter instances](https://docs.flutter.dev/add-to-app/multiple-flutters)
    * [Loading sequence and performance](https://docs.flutter.dev/add-to-app/performance)
  * Tools & editors expand_more
    * [Android Studio & IntelliJ](https://docs.flutter.dev/tools/android-studio)
    * [Visual Studio Code](https://docs.flutter.dev/tools/vs-code)
    * DevTools expand_more
      * [Overview](https://docs.flutter.dev/tools/devtools)
      * [Run from Android Studio & IntelliJ](https://docs.flutter.dev/tools/devtools/android-studio)
      * [Run from VS Code](https://docs.flutter.dev/tools/devtools/vscode)
      * [Run from command line](https://docs.flutter.dev/tools/devtools/cli)
      * [Flutter inspector](https://docs.flutter.dev/tools/devtools/inspector)
      * [Legacy Flutter inspector](https://docs.flutter.dev/tools/devtools/legacy-inspector)
      * [Performance view](https://docs.flutter.dev/tools/devtools/performance)
      * [CPU Profiler view](https://docs.flutter.dev/tools/devtools/cpu-profiler)
      * [Memory view](https://docs.flutter.dev/tools/devtools/memory)
      * [Debug console view](https://docs.flutter.dev/tools/devtools/console)
      * [Network view](https://docs.flutter.dev/tools/devtools/network)
      * [Debugger](https://docs.flutter.dev/tools/devtools/debugger)
      * [Logging view](https://docs.flutter.dev/tools/devtools/logging)
      * [App size tool](https://docs.flutter.dev/tools/devtools/app-size)
      * [DevTools extensions](https://docs.flutter.dev/tools/devtools/extensions)
      * [Validate deep links](https://docs.flutter.dev/tools/devtools/deep-links)
      * [Release notes](https://docs.flutter.dev/tools/devtools/release-notes)
    * [Flutter Widget Previewer](https://docs.flutter.dev/tools/widget-previewer)
    * [Flutter Property Editor](https://docs.flutter.dev/tools/property-editor)
    * [SDK overview](https://docs.flutter.dev/tools/sdk)
    * [Flutter's pubspec options](https://docs.flutter.dev/tools/pubspec)
    * [Automated fixes](https://docs.flutter.dev/tools/flutter-fix)
    * [Code formatting](https://docs.flutter.dev/tools/formatting)
  * Flutter concepts expand_more
    * [Architectural overview](https://docs.flutter.dev/resources/architectural-overview)
    * [Inside Flutter](https://docs.flutter.dev/resources/inside-flutter)
    * [Understanding constraints](https://docs.flutter.dev/ui/layout/constraints)
    * [Flutter's build modes](https://docs.flutter.dev/testing/build-modes)
    * [Hot reload](https://docs.flutter.dev/tools/hot-reload)
  * Resources expand_more
    * [FAQ](https://docs.flutter.dev/resources/faq)
    * [Books](https://docs.flutter.dev/resources/books)
    * [Videos](https://docs.flutter.dev/resources/videos)
    * [Courses](https://docs.flutter.dev/resources/courses)
    * [Learn Dart](https://docs.flutter.dev/resources/bootstrap-into-dart)
    * [Get support](https://docs.flutter.dev/resources/support)
    * Contribute expand_more
      * [Contribute to Flutter](https://docs.flutter.dev/contribute)
      * [Create useful bug reports](https://docs.flutter.dev/resources/bug-reports)
      * [Discover proposed features](https://docs.flutter.dev/resources/design-docs)
    * Reference expand_more
      * [Who is Dash?](https://docs.flutter.dev/dash)
      * [Widget index](https://docs.flutter.dev/reference/widgets)
      * [Create a new app](https://docs.flutter.dev/reference/create-new-app)
      * [flutter CLI](https://docs.flutter.dev/reference/flutter-cli)
      * [API docsopen_in_new](https://api.flutter.dev)


# Sorry, we couldn't find that page...
![Dash pointing you in the right direction](https://docs.flutter.dev/assets/images/404/dash_404.png)
But Dash is here to help - maybe one of these will point you in the right direction?
  * [Homepage](https://flutter.dev/)
  * [Package site](https://pub.dev/)
  * [API reference](https://api.flutter.dev/)
  * [Documentation](https://docs.flutter.dev/)
  * [Community](https://flutter.dev/community)
  * [Medium](https://medium.com/flutter)
  * [X (Twitter)](https://twitter.com/FlutterDev/)
  * [FAQ](https://docs.flutter.dev/resources/faq)


Was this page's content helpful?
thumb_up thumb_down
Thank you for your feedback!
[ feedback Provide details ](https://github.com/flutter/website/issues/new?template=1_page_issue.yml&&page-url=https://docs.flutter.dev/404.html&page-source=https://github.com/flutter/website/tree/main/src/content/404.html)
Thank you for your feedback! Please let us know what we can do to improve.
[ bug_report Provide details ](https://github.com/flutter/website/issues/new?template=1_page_issue.yml&&page-url=https://docs.flutter.dev/404.html&page-source=https://github.com/flutter/website/tree/main/src/content/404.html)
Unless stated otherwise, the documentation on this site reflects the latest stable version of Flutter. Page last updated on 2025-05-30. [View source](https://github.com/flutter/website/tree/main/src/content/404.html) or [report an issue](https://github.com/flutter/website/issues/new?template=1_page_issue.yml&&page-url=https://docs.flutter.dev/404.html&page-source=https://github.com/flutter/website/tree/main/src/content/404.html "Report an issue with this page").
[ ![Flutter logo](https://docs.flutter.dev/assets/images/branding/flutter/logo+text/horizontal/white.svg) ](https://flutter.dev)
[ ](https://medium.com/flutter "Flutter's Medium blog")[ ](https://youtube.com/@flutterdev "Flutter's YouTube")[ ](https://github.com/flutter "Flutter's GitHub")[ ](https://bsky.app/profile/flutter.dev "Flutter's Bluesky")[ ](https://twitter.com/FlutterDev "Flutter's X \(Twitter\)")
Except as otherwise noted, this site is licensed under a [Creative Commons Attribution 4.0 International License](https://creativecommons.org/licenses/by/4.0/), and code samples are licensed under the [3-Clause BSD License](https://opensource.org/licenses/BSD-3-Clause).
  * [Terms](https://docs.flutter.dev/tos "Terms of use")
  * [Brand](https://docs.flutter.dev/brand "Brand usage guidelines")
  * [Privacy](https://policies.google.com/privacy "Privacy policy")
  * [Security](https://docs.flutter.dev/security "Security philosophy and practices")




## Source Information
- URL: https://docs.flutter.dev/cookbook/testing/widget/golden-tests
- Harvested: 2025-08-19T23:21:39.504083
- Profile: code_examples
- Agent: architect
