---
agent_context: unknown
confidence: 0.95
harvested_at: '2025-08-19T01:21:02.957429'
profile: code_examples
source: https://riverpod.dev/docs/introduction/getting_started
topic: riverpod-state-management
---

# riverpod-state-management

[Skip to main content](https://riverpod.dev/docs/introduction/getting_started#__docusaurus_skipToContent_fallback)
[![Riverpod](https://riverpod.dev/img/logo.png)**Riverpod**](https://riverpod.dev/)
[Docs](https://riverpod.dev/docs/introduction/why_riverpod)
[English](https://riverpod.dev/docs/introduction/getting_started)
  * [English](https://riverpod.dev/docs/introduction/getting_started)
  * [Français](https://riverpod.dev/fr/docs/introduction/getting_started)
  * [한국어](https://riverpod.dev/ko/docs/introduction/getting_started)
  * [日本語](https://riverpod.dev/ja/docs/introduction/getting_started)
  * [Español](https://riverpod.dev/es/docs/introduction/getting_started)
  * [বাংলা](https://riverpod.dev/bn/docs/introduction/getting_started)
  * [Deutsch](https://riverpod.dev/de/docs/introduction/getting_started)
  * [Italiano](https://riverpod.dev/it/docs/introduction/getting_started)
  * [Русский](https://riverpod.dev/ru/docs/introduction/getting_started)
  * [Türkçe](https://riverpod.dev/tr/docs/introduction/getting_started)
  * [简体中文](https://riverpod.dev/zh-Hans/docs/introduction/getting_started)


[GitHub](https://github.com/rrousselGit/riverpod)
Search`K`
  * [What's new in Riverpod 3.0](https://riverpod.dev/docs/whats_new)
  * [Migrating from 2.0 to 3.0](https://riverpod.dev/docs/3.0_migration)
  * Introduction
    * [Why Riverpod?](https://riverpod.dev/docs/introduction/why_riverpod)
    * [Getting started](https://riverpod.dev/docs/introduction/getting_started)
  * [Riverpod for Provider Users](https://riverpod.dev/docs/introduction/getting_started)
  * References
    * [All Providers](https://riverpod.dev/docs/introduction/getting_started)
    * [Containers/Scopes](https://riverpod.dev/docs/introduction/getting_started)
    * [Refs](https://riverpod.dev/docs/introduction/getting_started)
    * [Consumers](https://riverpod.dev/docs/introduction/getting_started)
    * [Offline persistence (experimental)](https://riverpod.dev/docs/introduction/getting_started)
    * [core](https://riverpod.dev/docs/introduction/getting_started)
    * [misc](https://pub.dev/documentation/hooks_riverpod/3.0.0-dev.16/misc/)
  * Essentials
    * [Make your first provider/network request](https://riverpod.dev/docs/essentials/first_request)
    * [Performing side effects](https://riverpod.dev/docs/essentials/side_effects)
    * [Passing arguments to your requests](https://riverpod.dev/docs/essentials/passing_args)
    * [Websockets and synchronous execution](https://riverpod.dev/docs/essentials/websockets_sync)
    * [Combining requests](https://riverpod.dev/docs/essentials/combining_requests)
    * [Clearing cache and reacting to state disposal](https://riverpod.dev/docs/essentials/auto_dispose)
    * [Eager initialization of providers](https://riverpod.dev/docs/essentials/eager_initialization)
    * [Testing your providers](https://riverpod.dev/docs/essentials/testing)
    * [Logging and error reporting](https://riverpod.dev/docs/essentials/provider_observer)
    * [FAQ](https://riverpod.dev/docs/essentials/faq)
    * [DO/DON'T](https://riverpod.dev/docs/essentials/do_dont)
  * Case studies
    * [Pull to refresh](https://riverpod.dev/docs/case_studies/pull_to_refresh)
    * [Debouncing/Cancelling network requests](https://riverpod.dev/docs/case_studies/cancel)
  * Advanced topics
    * [Optimizing performance](https://riverpod.dev/docs/advanced/select)
  * Concepts
    * [About code generation](https://riverpod.dev/docs/concepts/about_code_generation)
    * [About hooks](https://riverpod.dev/docs/concepts/about_hooks)
  * Migration guides
    * [From `StateNotifier`](https://riverpod.dev/docs/migration/from_state_notifier)
    * [From `ChangeNotifier`](https://riverpod.dev/docs/migration/from_change_notifier)
    * [^0.14.0 to ^1.0.0](https://riverpod.dev/docs/migration/0.14.0_to_1.0.0)
    * [^0.13.0 to ^0.14.0](https://riverpod.dev/docs/migration/0.13.0_to_0.14.0)
  * [Official examples](https://riverpod.dev/docs/introduction/getting_started)
  * [Third party examples](https://riverpod.dev/docs/introduction/getting_started)
  * [API reference](https://pub.dev/documentation/hooks_riverpod/latest/hooks_riverpod/hooks_riverpod-library.html)
  * [Concepts 🚧](https://riverpod.dev/docs/introduction/getting_started)
  * [All Providers 🚧](https://riverpod.dev/docs/introduction/getting_started)
  * [Guides 🚧](https://riverpod.dev/docs/introduction/getting_started)


  * [](https://riverpod.dev/)
  * Introduction
  * Getting started


On this page
# Getting started
## Try Riverpod online[​](https://riverpod.dev/docs/introduction/getting_started#try-riverpod-online "Direct link to Try Riverpod online")
To get a feel of Riverpod, try it online on [Dartpad](https://dartpad.dev/?null_safety=true&id=ef06ab3ce0b822e6cc5db0575248e6e2) or on [Zapp](https://zapp.run/new):
## Installing the package[​](https://riverpod.dev/docs/introduction/getting_started#installing-the-package "Direct link to Installing the package")
Riverpod comes as a main "riverpod" package that’s self-sufficient, complemented by optional packages for using code generation ([About code generation](https://riverpod.dev/docs/concepts/about_code_generation)) and hooks ([About hooks](https://riverpod.dev/docs/concepts/about_hooks)).
Once you know what package(s) you want to install, proceed to add the dependency to your app in a single line like this:
  * Flutter
  * Dart only


  * riverpod_generator
  * riverpod_generator + flutter_hooks
  * riverpod
  * riverpod + flutter_hooks


```
flutter pub add flutter_riverpodflutter pub add riverpod_annotationflutter pub add dev:riverpod_generatorflutter pub add dev:build_runnerflutter pub add dev:custom_lintflutter pub add dev:riverpod_lint
```

```
flutter pub add hooks_riverpodflutter pub add flutter_hooksflutter pub add riverpod_annotationflutter pub add dev:riverpod_generatorflutter pub add dev:build_runnerflutter pub add dev:custom_lintflutter pub add dev:riverpod_lint
```

```
flutter pub add flutter_riverpodflutter pub add dev:custom_lintflutter pub add dev:riverpod_lint
```

```
flutter pub add hooks_riverpodflutter pub add flutter_hooksflutter pub add dev:custom_lintflutter pub add dev:riverpod_lint
```

  * riverpod_generator
  * riverpod


```
dart pub add riverpoddart pub add riverpod_annotationdart pub add dev:riverpod_generatordart pub add dev:build_runnerdart pub add dev:custom_lintdart pub add dev:riverpod_lint
```

```
dart pub add riverpoddart pub add dev:custom_lintdart pub add dev:riverpod_lint
```

Alternatively, you can manually add the dependency to your app from within your `pubspec.yaml`:
  * Flutter
  * Dart only


  * riverpod_generator
  * riverpod_generator + flutter_hooks
  * riverpod
  * riverpod + flutter_hooks


pubspec.yaml
```
name: my_app_nameenvironment:sdk:">=3.0.0 <4.0.0"flutter:">=3.0.0"dependencies:flutter:sdk: flutterflutter_riverpod: ^3.0.0-dev.17riverpod_annotation: ^3.0.0-dev.17dev_dependencies:build_runner:custom_lint:riverpod_generator: ^3.0.0-dev.17riverpod_lint: ^3.0.0-dev.17
```

pubspec.yaml
```
name: my_app_nameenvironment:sdk:">=3.0.0 <4.0.0"flutter:">=3.0.0"dependencies:flutter:sdk: flutterhooks_riverpod: ^3.0.0-dev.17flutter_hooks:riverpod_annotation: ^3.0.0-dev.17dev_dependencies:build_runner:custom_lint:riverpod_generator: ^3.0.0-dev.17riverpod_lint: ^3.0.0-dev.17
```

pubspec.yaml
```
name: my_app_nameenvironment:sdk:">=3.0.0 <4.0.0"flutter:">=3.0.0"dependencies:flutter:sdk: flutterflutter_riverpod: ^3.0.0-dev.17dev_dependencies:custom_lint:riverpod_lint: ^3.0.0-dev.17
```

pubspec.yaml
```
name: my_app_nameenvironment:sdk:">=3.0.0 <4.0.0"flutter:">=3.0.0"dependencies:flutter:sdk: flutterhooks_riverpod: ^3.0.0-dev.17flutter_hooks:dev_dependencies:custom_lint:riverpod_lint: ^3.0.0-dev.17
```

Then, install packages with `flutter pub get`.
  * riverpod_generator
  * riverpod


pubspec.yaml
```
name: my_app_nameenvironment:sdk:">=3.0.0 <4.0.0"dependencies:riverpod: ^3.0.0-dev.17riverpod_annotation: ^3.0.0-dev.17dev_dependencies:build_runner:custom_lint:riverpod_generator: ^3.0.0-dev.17riverpod_lint: ^3.0.0-dev.17
```

pubspec.yaml
```
name: my_app_nameenvironment:sdk:">=3.0.0 <4.0.0"dependencies:riverpod: ^3.0.0-dev.17dev_dependencies:custom_lint:riverpod_lint: ^3.0.0-dev.17
```

Then, install packages with `dart pub get`.
info
If using code-generation, you can now run the code-generator with:
```
dart run build_runner watch -d
```

That's it. You've added [Riverpod](https://github.com/rrousselgit/riverpod) to your app!
## Enabling riverpod_lint/custom_lint[​](https://riverpod.dev/docs/introduction/getting_started#enabling-riverpod_lintcustom_lint "Direct link to Enabling riverpod_lint/custom_lint")
Riverpod comes with an optional [riverpod_lint](https://pub.dev/packages/riverpod_lint) package that provides lint rules to help you write better code, and provide custom refactoring options.
The package should already be installed if you followed the previous steps, but a separate step is necessary to enable it.
To enable [riverpod_lint](https://pub.dev/packages/riverpod_lint), you need add an `analysis_options.yaml` placed next to your `pubspec.yaml` and include the following:
analysis_options.yaml
```
analyzer:plugins:- custom_lint
```

You should now see warnings in your IDE if you made mistakes when using Riverpod in your codebase.
To see the full list of warnings and refactorings, head to the [riverpod_lint](https://pub.dev/packages/riverpod_lint) page.
note
Those warnings will not show-up in the `dart analyze` command. If you want to check those warnings in the CI/terminal, you can run the following:
```
dart run custom_lint
```

## Usage example: Hello world[​](https://riverpod.dev/docs/introduction/getting_started#usage-example-hello-world "Direct link to Usage example: Hello world")
Now that we have installed [Riverpod](https://github.com/rrousselgit/riverpod), we can start using it.
The following snippets showcase how to use our new dependency to make a "Hello world":
  * Flutter
  * Dart only


  * riverpod_generator
  * riverpod_generator + flutter_hooks
  * riverpod
  * riverpod + flutter_hooks


lib/main.dart
```
import'package:flutter/material.dart';import'package:flutter_riverpod/flutter_riverpod.dart';import'package:riverpod_annotation/riverpod_annotation.dart';part'main.g.dart';// We create a "provider", which will store a value (here "Hello world").// By using a provider, this allows us to mock/override the value exposed.@riverpodStringhelloWorld(Ref ref){return'Hello world';}voidmain(){runApp(// For widgets to be able to read providers, we need to wrap the entire// application in a "ProviderScope" widget.// This is where the state of our providers will be stored.ProviderScope(   child:MyApp(),),);}// Extend ConsumerWidget instead of StatelessWidget, which is exposed by RiverpodclassMyAppextendsConsumerWidget{@overrideWidgetbuild(BuildContext context,WidgetRef ref){finalString value = ref.watch(helloWorldProvider);returnMaterialApp(   home:Scaffold(    appBar:AppBar(title:constText('Example')),    body:Center(     child:Text(value),),),);}}
```

lib/main.dart
```
import'package:flutter/material.dart';import'package:flutter_hooks/flutter_hooks.dart';import'package:hooks_riverpod/hooks_riverpod.dart';import'package:riverpod_annotation/riverpod_annotation.dart';part'main.g.dart';// We create a "provider", which will store a value (here "Hello world").// By using a provider, this allows us to mock/override the value exposed.@riverpodStringhelloWorld(Ref ref){return'Hello world';}voidmain(){runApp(// For widgets to be able to read providers, we need to wrap the entire// application in a "ProviderScope" widget.// This is where the state of our providers will be stored.ProviderScope(   child:MyApp(),),);}// Extend HookConsumerWidget instead of HookWidget, which is exposed by RiverpodclassMyAppextendsHookConsumerWidget{@overrideWidgetbuild(BuildContext context,WidgetRef ref){// We can use hooks inside HookConsumerWidgetfinal counter =useState(0);finalString value = ref.watch(helloWorldProvider);returnMaterialApp(   home:Scaffold(    appBar:AppBar(title:constText('Example')),    body:Center(     child:Text('$value${counter.value}'),),),);}}
```

lib/main.dart
```
import'package:flutter/material.dart';import'package:flutter_riverpod/flutter_riverpod.dart';// We create a "provider", which will store a value (here "Hello world").// By using a provider, this allows us to mock/override the value exposed.final helloWorldProvider =Provider((_)=>'Hello world');voidmain(){runApp(// For widgets to be able to read providers, we need to wrap the entire// application in a "ProviderScope" widget.// This is where the state of our providers will be stored.ProviderScope(   child:MyApp(),),);}// Extend ConsumerWidget instead of StatelessWidget, which is exposed by RiverpodclassMyAppextendsConsumerWidget{@overrideWidgetbuild(BuildContext context,WidgetRef ref){finalString value = ref.watch(helloWorldProvider);returnMaterialApp(   home:Scaffold(    appBar:AppBar(title:constText('Example')),    body:Center(     child:Text(value),),),);}}
```

lib/main.dart
```
import'package:flutter/material.dart';import'package:flutter_hooks/flutter_hooks.dart';import'package:hooks_riverpod/hooks_riverpod.dart';// We create a "provider", which will store a value (here "Hello world").// By using a provider, this allows us to mock/override the value exposed.final helloWorldProvider =Provider((_)=>'Hello world');voidmain(){runApp(// For widgets to be able to read providers, we need to wrap the entire// application in a "ProviderScope" widget.// This is where the state of our providers will be stored.ProviderScope(   child:MyApp(),),);}// Extend HookConsumerWidget instead of HookWidget, which is exposed by RiverpodclassMyAppextendsHookConsumerWidget{@overrideWidgetbuild(BuildContext context,WidgetRef ref){// We can use hooks inside HookConsumerWidgetfinal counter =useState(0);finalString value = ref.watch(helloWorldProvider);returnMaterialApp(   home:Scaffold(    appBar:AppBar(title:constText('Example')),    body:Center(     child:Text('$value${counter.value}'),),),);}}
```

Then, start the application with `flutter run`. This will render "Hello world" on your device.
  * riverpod_generator
  * riverpod


lib/main.dart
```
import'package:riverpod_annotation/riverpod_annotation.dart';part'main.g.dart';// We create a "provider", which will store a value (here "Hello world").// By using a provider, this allows us to mock/override the value exposed.@riverpodStringhelloWorld(Ref ref){return'Hello world';}voidmain(){// This object is where the state of our providers will be stored.final container =ProviderContainer();// Thanks to "container", we can read our provider.final value = container.read(helloWorldProvider);print(value);// Hello world}
```

lib/main.dart
```
import'package:riverpod/riverpod.dart';// We create a "provider", which will store a value (here "Hello world").// By using a provider, this allows us to mock/override the value exposed.final helloWorldProvider =Provider((_)=>'Hello world');voidmain(){// This object is where the state of our providers will be stored.final container =ProviderContainer();// Thanks to "container", we can read our provider.final value = container.read(helloWorldProvider);print(value);// Hello world}
```

Then, start the application with `dart lib/main.dart`. This will print "Hello world" in the console.
## Going further: Installing code snippets[​](https://riverpod.dev/docs/introduction/getting_started#going-further-installing-code-snippets "Direct link to Going further: Installing code snippets")
If you are using `Flutter` and `VS Code` , consider using [Flutter Riverpod Snippets](https://marketplace.visualstudio.com/items?itemName=robert-brunhage.flutter-riverpod-snippets)
If you are using `Flutter` and `Android Studio` or `IntelliJ`, consider using [Flutter Riverpod Snippets](https://plugins.jetbrains.com/plugin/14641-flutter-riverpod-snippets)
![img](https://riverpod.dev/assets/images/greetingProvider-47179931ef18184e7ab68f4e701ca916.gif)
## Choose your next step[​](https://riverpod.dev/docs/introduction/getting_started#choose-your-next-step "Direct link to Choose your next step")
Learn some basic concepts:
  * [Providers](https://riverpod.dev/docs/concepts/providers)


Follow a cookbook:
  * [Testing](https://riverpod.dev/docs/cookbooks/testing)


[Edit this page](https://github.com/rrousselGit/riverpod/edit/master/website/docs/introduction/getting_started.mdx)
[PreviousWhy Riverpod?](https://riverpod.dev/docs/introduction/why_riverpod)[NextMake your first provider/network request](https://riverpod.dev/docs/essentials/first_request)
  * [Try Riverpod online](https://riverpod.dev/docs/introduction/getting_started#try-riverpod-online)
  * [Installing the package](https://riverpod.dev/docs/introduction/getting_started#installing-the-package)
  * [Enabling riverpod_lint/custom_lint](https://riverpod.dev/docs/introduction/getting_started#enabling-riverpod_lintcustom_lint)
  * [Usage example: Hello world](https://riverpod.dev/docs/introduction/getting_started#usage-example-hello-world)
  * [Going further: Installing code snippets](https://riverpod.dev/docs/introduction/getting_started#going-further-installing-code-snippets)
  * [Choose your next step](https://riverpod.dev/docs/introduction/getting_started#choose-your-next-step)


Docs
  * [Why Riverpod?](https://riverpod.dev/docs/introduction/why_riverpod)
  * [Getting started](https://riverpod.dev/docs/introduction/getting_started)


Community
  * [Discord](https://discord.gg/hUUQkd9v)
  * [GitHub](https://github.com/rrousselgit/riverpod)
  * [Stack Overflow](https://stackoverflow.com/questions/tagged/riverpod)
  * [Twitter](https://twitter.com/remi_rousselet)
  * [Code of conduct](https://github.com/rrousselGit/riverpod/blob/master/CODE_OF_CONDUCT.md)
  * [Contributing guide](https://github.com/rrousselGit/riverpod/blob/master/CONTRIBUTING.md)


Sponsors
  * [![Deploys by Netlify](https://www.netlify.com/img/global/badges/netlify-color-bg.svg)](https://www.netlify.com)


![Riverpod](https://riverpod.dev/img/full_logo.svg)
Copyright © 2025 Remi Rousselet. Built with Docusaurus.


## Source Information
- URL: https://riverpod.dev/docs/introduction/getting_started
- Harvested: 2025-08-19T01:21:02.957429
- Profile: code_examples
- Agent: unknown
