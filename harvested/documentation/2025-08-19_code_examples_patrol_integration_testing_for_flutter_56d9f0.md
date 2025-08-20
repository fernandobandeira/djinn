---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:20:43.582445'
profile: code_examples
source: https://pub.dev/packages/patrol
topic: Patrol Integration Testing for Flutter
---

# Patrol Integration Testing for Flutter

This site uses cookies from Google to deliver and enhance the quality of its services and to analyze traffic.[Learn more](https://policies.google.com/technologies/cookies?hl=en)[OK, got it](https://pub.dev/packages/patrol)
[![](https://pub.dev/static/hash-0mq6dpb8/img/pub-dev-logo.svg)](https://pub.dev/)
Sign in
Help
### pub.dev
[Searching for packages](https://pub.dev/help/search)[Package scoring and pub points](https://pub.dev/help/scoring)
### Flutter
[Using packages](https://flutter.dev/using-packages/)[Developing packages and plugins](https://flutter.dev/developing-packages/)[Publishing a package](https://dart.dev/tools/pub/publishing)
### Dart
[Using packages](https://dart.dev/guides/packages)[Publishing a package](https://dart.dev/tools/pub/publishing)
### pub.dev ![toggle folding of the section](https://pub.dev/static/hash-0mq6dpb8/img/nav-mobile-foldable-icon.svg)
[Searching for packages](https://pub.dev/help/search)[Package scoring and pub points](https://pub.dev/help/scoring)
### Flutter ![toggle folding of the section](https://pub.dev/static/hash-0mq6dpb8/img/nav-mobile-foldable-icon.svg)
[Using packages](https://flutter.dev/using-packages/)[Developing packages and plugins](https://flutter.dev/developing-packages/)[Publishing a package](https://dart.dev/tools/pub/publishing)
### Dart ![toggle folding of the section](https://pub.dev/static/hash-0mq6dpb8/img/nav-mobile-foldable-icon.svg)
[Using packages](https://dart.dev/guides/packages)[Publishing a package](https://dart.dev/tools/pub/publishing)
# patrol 3.19.0 ![copy "patrol: ^3.19.0" to clipboard](https://pub.dev/static/hash-0mq6dpb8/img/content-copy-icon.svg)
patrol: ^3.19.0 copied to clipboard
Published [14 days ago](https://pub.dev/packages/patrol "Aug 5, 2025") • [![verified publisher](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-verified.svg)leancode.co](https://pub.dev/publishers/leancode.co)Dart 3 compatible
SDK[Flutter](https://pub.dev/packages?q=sdk%3Aflutter "Packages compatible with Flutter SDK")
Platform[Android](https://pub.dev/packages?q=platform%3Aandroid "Packages compatible with Android platform")[iOS](https://pub.dev/packages?q=platform%3Aios "Packages compatible with iOS platform")[macOS](https://pub.dev/packages?q=platform%3Amacos "Packages compatible with macOS platform")
![liked status: inactive](https://pub.dev/static/hash-0mq6dpb8/img/like-inactive.svg)![liked status: active](https://pub.dev/static/hash-0mq6dpb8/img/like-active.svg)627
→
### Metadata
Powerful Flutter-native UI testing framework overcoming limitations of existing Flutter testing tools. Ready for action! 
More...
  * Readme
  * [Changelog](https://pub.dev/packages/patrol/changelog)
  * [Example](https://pub.dev/packages/patrol/example)
  * [Installing](https://pub.dev/packages/patrol/install)
  * [Versions](https://pub.dev/packages/patrol/versions)
  * [Scores](https://pub.dev/packages/patrol/score)


# patrol [#](https://pub.dev/packages/patrol#patrol)
[![patrol on pub.dev](https://img.shields.io/pub/v/patrol.svg)](https://pub.dartlang.org/packages/patrol) [![codestyle](https://img.shields.io/badge/style-leancode__lint-black)](https://pub.dartlang.org/packages/leancode_lint)
`patrol` package builds on top of `flutter_test` and `integration_test` to make it easy to control the native UI from Dart test code. Created and supported by [LeanCode](https://leancode.co).
It must be used together with [patrol_cli](https://pub.dev/packages/patrol_cli).
It also provides a new custom finder system to make Flutter widget tests more concise and understandable, and writing them – faster and more fun. It you want to only use custom finders, check out [patrol_finders](https://pub.dev/packages/patrol_finders).
## Installation [#](https://pub.dev/packages/patrol#installation)
```
$ dart pub add patrol --dev

```

copied to clipboard
## Usage [#](https://pub.dev/packages/patrol#usage)
Patrol has 2 main features – [native automation](https://patrol.leancode.co/native/overview) and [custom finders](https://patrol.leancode.co/finders/overview).
[Read our docs](https://patrol.leancode.co) to learn more about them!
### Accessing native platform features [#](https://pub.dev/packages/patrol#accessing-native-platform-features)
```
// example/integration_test/example_test.dart
import 'package:example/main.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:patrol/patrol.dart';
void main() {
 patrolTest(
  'counter state is the same after going to home and going back',
  ($) async {
   await $.pumpWidgetAndSettle(const MyApp());
   await $(FloatingActionButton).tap();
   expect($(#counterText).text, '1');
   await $.native.pressHome();
   await $.native.pressDoubleRecentApps();
   expect($(#counterText).text, '1');
   await $(FloatingActionButton).tap();
   expect($(#counterText).text, '2');
   await $.native.openNotifications();
   await $.native.pressBack();
  },
 );
}

```

copied to clipboard
### Custom finders [#](https://pub.dev/packages/patrol#custom-finders)
```
import 'package:example/main.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:patrol/patrol.dart';
void main() {
 patrolTest(
  'logs in successfully',
  ($) async {
   await $.pumpWidgetAndSettle(const MyApp());
   await $(#emailInput).enterText('user@leancode.co');
   await $(#passwordInput).enterText('ny4ncat');
   // Finds all widgets with text 'Log in' which are descendants of widgets with key
   // box1, which are descendants of a Scaffold widget and tap on the first one.
   await $(Scaffold).$(#box1).$('Log in').tap();
   // Finds all Scrollables which have Text descendant
   $(Scrollable).containing(Text);
   // Finds all Scrollables which have a Button descendant which has a Text descendant
   $(Scrollable).containing($(Button).containing(Text));
   // Finds all Scrollables which have a Button descendant and a Text descendant
   $(Scrollable).containing(Button).containing(Text);
  },
 );
}

```

copied to clipboard
[627likes160points141kdownloads](https://pub.dev/packages/patrol/score)
![screenshot](https://pub.dev/packages/patrol/versions/3.19.0/gen-res/gen/190x190/logo.webp)
![](https://pub.dev/static/hash-0mq6dpb8/img/collections_white_24dp.svg)
### Publisher
[![verified publisher](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-verified.svg)leancode.co](https://pub.dev/publishers/leancode.co)
### Weekly Downloads
2024.09.21 - 2025.08.16
### Metadata
Powerful Flutter-native UI testing framework overcoming limitations of existing Flutter testing tools. Ready for action! 
[Homepage](https://patrol.leancode.co)[Repository (GitHub)](https://github.com/leancodepl/patrol/tree/master/packages/patrol)[View/report issues](https://github.com/leancodepl/patrol/issues)[Contributing](https://github.com/leancodepl/patrol/blob/master/CONTRIBUTING.md)
### Topics
[#testing](https://pub.dev/packages?q=topic%3Atesting "Packages that facilitate testing.") [#integration-test](https://pub.dev/packages?q=topic%3Aintegration-test) [#ui-test](https://pub.dev/packages?q=topic%3Aui-test) [#patrol](https://pub.dev/packages?q=topic%3Apatrol)
### Documentation
[Documentation](https://patrol.leancode.co)[API reference](https://pub.dev/documentation/patrol/latest/)
### License
![](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-balance.svg)Apache-2.0 ([license](https://pub.dev/packages/patrol/license))
### Dependencies
[boolean_selector](https://pub.dev/packages/boolean_selector "^2.1.1"), [equatable](https://pub.dev/packages/equatable "^2.0.5"), [flutter](https://api.flutter.dev/), [flutter_test](https://api.flutter.dev/flutter/flutter_test/flutter_test-library.html), [http](https://pub.dev/packages/http "^1.1.0"), [json_annotation](https://pub.dev/packages/json_annotation "^4.8.1"), [meta](https://pub.dev/packages/meta "^1.10.0"), [patrol_finders](https://pub.dev/packages/patrol_finders "^2.9.0"), [patrol_log](https://pub.dev/packages/patrol_log "^0.5.0"), [shelf](https://pub.dev/packages/shelf "^1.4.1"), [test_api](https://pub.dev/packages/test_api "^0.7.0")
### More
[Packages that depend on patrol](https://pub.dev/packages?q=dependency%3Apatrol)
### ← Metadata
[627likes160points141kdownloads](https://pub.dev/packages/patrol/score)
![screenshot](https://pub.dev/packages/patrol/versions/3.19.0/gen-res/gen/190x190/logo.webp)
![](https://pub.dev/static/hash-0mq6dpb8/img/collections_white_24dp.svg)
### Publisher
[![verified publisher](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-verified.svg)leancode.co](https://pub.dev/publishers/leancode.co)
### Weekly Downloads
2024.09.21 - 2025.08.16
### Metadata
Powerful Flutter-native UI testing framework overcoming limitations of existing Flutter testing tools. Ready for action! 
[Homepage](https://patrol.leancode.co)[Repository (GitHub)](https://github.com/leancodepl/patrol/tree/master/packages/patrol)[View/report issues](https://github.com/leancodepl/patrol/issues)[Contributing](https://github.com/leancodepl/patrol/blob/master/CONTRIBUTING.md)
### Topics
[#testing](https://pub.dev/packages?q=topic%3Atesting "Packages that facilitate testing.") [#integration-test](https://pub.dev/packages?q=topic%3Aintegration-test) [#ui-test](https://pub.dev/packages?q=topic%3Aui-test) [#patrol](https://pub.dev/packages?q=topic%3Apatrol)
### Documentation
[Documentation](https://patrol.leancode.co)[API reference](https://pub.dev/documentation/patrol/latest/)
### License
![](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-balance.svg)Apache-2.0 ([license](https://pub.dev/packages/patrol/license))
### Dependencies
[boolean_selector](https://pub.dev/packages/boolean_selector "^2.1.1"), [equatable](https://pub.dev/packages/equatable "^2.0.5"), [flutter](https://api.flutter.dev/), [flutter_test](https://api.flutter.dev/flutter/flutter_test/flutter_test-library.html), [http](https://pub.dev/packages/http "^1.1.0"), [json_annotation](https://pub.dev/packages/json_annotation "^4.8.1"), [meta](https://pub.dev/packages/meta "^1.10.0"), [patrol_finders](https://pub.dev/packages/patrol_finders "^2.9.0"), [patrol_log](https://pub.dev/packages/patrol_log "^0.5.0"), [shelf](https://pub.dev/packages/shelf "^1.4.1"), [test_api](https://pub.dev/packages/test_api "^0.7.0")
### More
[Packages that depend on patrol](https://pub.dev/packages?q=dependency%3Apatrol)
Back
![previous](https://pub.dev/static/hash-0mq6dpb8/img/keyboard_arrow_left.svg)![next](https://pub.dev/static/hash-0mq6dpb8/img/keyboard_arrow_right.svg)
[Dart language](https://dart.dev/)[Report package](https://pub.dev/report?subject=package%3Apatrol&url=https%3A%2F%2Fpub.dev%2Fpackages%2Fpatrol)[Policy](https://pub.dev/policy)[Terms](https://www.google.com/intl/en/policies/terms/)[API Terms](https://developers.google.com/terms/)[Security](https://pub.dev/security)[Privacy](https://www.google.com/intl/en/policies/privacy/)[Help](https://pub.dev/help)[![RSS](https://pub.dev/static/hash-0mq6dpb8/img/rss-feed-icon.svg)](https://pub.dev/feed.atom)[![bug report](https://pub.dev/static/hash-0mq6dpb8/img/bug-report-white-96px.png)](https://github.com/dart-lang/pub-dev/issues/new?body=URL%3A+https%3A%2F%2Fpub.dev%2Fpackages%2Fpatrol%0A%0A%3CDescribe+your+issue+or+suggestion+here%3E&title=%3CSummarize+your+issues+here%3E&labels=Area%3A+site+feedback)


## Source Information
- URL: https://pub.dev/packages/patrol
- Harvested: 2025-08-19T23:20:43.582445
- Profile: code_examples
- Agent: architect
