---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:20:49.764333'
profile: code_examples
source: https://pub.dev/packages/mocktail
topic: Mocktail vs Mockito Comparison for Flutter Testing
---

# Mocktail vs Mockito Comparison for Flutter Testing

This site uses cookies from Google to deliver and enhance the quality of its services and to analyze traffic.[Learn more](https://policies.google.com/technologies/cookies?hl=en)[OK, got it](https://pub.dev/packages/mocktail)
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
# mocktail 1.0.4 ![copy "mocktail: ^1.0.4" to clipboard](https://pub.dev/static/hash-0mq6dpb8/img/content-copy-icon.svg)
mocktail: ^1.0.4 copied to clipboard
Published [14 months ago](https://pub.dev/packages/mocktail "Jun 12, 2024") • [![verified publisher](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-verified.svg)felangel.dev](https://pub.dev/publishers/felangel.dev)Dart 3 compatible
SDK[Dart](https://pub.dev/packages?q=sdk%3Adart "Packages compatible with Dart SDK")[Flutter](https://pub.dev/packages?q=sdk%3Aflutter "Packages compatible with Flutter SDK")
Platform[Android](https://pub.dev/packages?q=platform%3Aandroid "Packages compatible with Android platform")[iOS](https://pub.dev/packages?q=platform%3Aios "Packages compatible with iOS platform")[Linux](https://pub.dev/packages?q=platform%3Alinux "Packages compatible with Linux platform")[macOS](https://pub.dev/packages?q=platform%3Amacos "Packages compatible with macOS platform")[web](https://pub.dev/packages?q=platform%3Aweb "Packages compatible with Web platform")[Windows](https://pub.dev/packages?q=platform%3Awindows "Packages compatible with Windows platform")
![liked status: inactive](https://pub.dev/static/hash-0mq6dpb8/img/like-inactive.svg)![liked status: active](https://pub.dev/static/hash-0mq6dpb8/img/like-active.svg)1.1k
→
### Metadata
A Dart mock library which simplifies mocking with null safety support and no manual mocks or code generation.
More...
  * Readme
  * [Changelog](https://pub.dev/packages/mocktail/changelog)
  * [Example](https://pub.dev/packages/mocktail/example)
  * [Installing](https://pub.dev/packages/mocktail/install)
  * [Versions](https://pub.dev/packages/mocktail/versions)
  * [Scores](https://pub.dev/packages/mocktail/score)


# 🍹 mocktail [#](https://pub.dev/packages/mocktail#-mocktail)
[![Pub](https://img.shields.io/pub/v/mocktail.svg)](https://pub.dev/packages/mocktail) [![mocktail](https://github.com/felangel/mocktail/actions/workflows/mocktail.yaml/badge.svg)](https://github.com/felangel/mocktail/actions) [![coverage](https://raw.githubusercontent.com/felangel/mocktail/main/coverage_badge.svg)](https://github.com/felangel/mocktail/actions) [![License: MIT](https://img.shields.io/badge/license-MIT-purple.svg)](https://opensource.org/licenses/MIT)
Mock library for Dart inspired by [mockito](https://pub.dev/packages/mockito).
Mocktail focuses on providing a familiar, simple API for creating mocks in Dart (with null-safety) without the need for manual mocks or code generation.
## Creating a Mock [#](https://pub.dev/packages/mocktail#creating-a-mock)
```
import 'package:mocktail/mocktail.dart';
// A Real Cat class
class Cat {
 String sound() => 'meow!';
 bool likes(String food, {bool isHungry = false}) => false;
 final int lives = 9;
}
// A Mock Cat class
class MockCat extends Mock implements Cat {}
void main() {
 // Create a Mock Cat instance
 final cat = MockCat();
}

```

copied to clipboard
## Stub and Verify Behavior [#](https://pub.dev/packages/mocktail#stub-and-verify-behavior)
The `MockCat` instance can then be used to stub and verify calls.
```
// Stub the `sound` method.
when(() => cat.sound()).thenReturn('meow');
// Verify no interactions have occurred.
verifyNever(() => cat.sound());
// Interact with the mock cat instance.
cat.sound();
// Verify the interaction occurred.
verify(() => cat.sound()).called(1);
// When mocktail verifies an invocation, it is then excluded from further verifications.
verifyNever(() => cat.sound());
// Interact with the mock instance again.
cat.sound();
// The verification count is 1 since there was only 1 invocation since the last verification.
verify(() => cat.sound()).called(1);

```

copied to clipboard
## Additional Usage [#](https://pub.dev/packages/mocktail#additional-usage)
```
// Stub a method before interacting with the mock.
when(() => cat.sound()).thenReturn('purrr!');
expect(cat.sound(), 'purrr!');
// You can interact with the mock multiple times.
expect(cat.sound(), 'purrr!');
// You can change the stub.
when(() => cat.sound()).thenReturn('meow!');
expect(cat.sound(), 'meow');
// You can stub getters.
when(() => cat.lives).thenReturn(10);
expect(cat.lives, 10);
// You can stub a method for specific arguments.
when(() => cat.likes('fish', isHungry: false)).thenReturn(true);
expect(cat.likes('fish', isHungry: false), isTrue);
// You can verify the interaction for specific arguments.
verify(() => cat.likes('fish', isHungry: false)).called(1);
// Or alternatively use any(that: ...) to use a matcher.
verify(() => cat.likes(any(that: isA<String>().having((food) => food, 'name', 'fish')))).called(1);
// You can stub a method using argument matcher: `any`.
// When stubbing a positional argument, use `any()`.
// When stubbing a named argument, use `any(named: '<argName>`)`.
// A custom matcher can be provided using `any(that: customMatcher)`.
when(() => cat.likes(any(), isHungry: any(named: 'isHungry', that: isFalse)).thenReturn(true);
expect(cat.likes('fish', isHungry: false), isTrue);
// You can stub a method to throw.
when(() => cat.sound()).thenThrow(Exception('oops'));
expect(() => cat.sound(), throwsA(isA<Exception>()));
// You can calculate stubs dynamically.
final sounds = ['purrr', 'meow'];
when(() => cat.sound()).thenAnswer((_) => sounds.removeAt(0));
expect(cat.sound(), 'purrr');
expect(cat.sound(), 'meow');
// You can capture any argument.
when(() => cat.likes('fish')).thenReturn(true);
expect(cat.likes('fish'), isTrue);
final captured = verify(() => cat.likes(captureAny())).captured;
expect(captured.last, equals(['fish']));
// You can capture a specific argument based on a matcher.
when(() => cat.likes(any())).thenReturn(true);
expect(cat.likes('fish'), isTrue);
expect(cat.likes('dog food'), isTrue);
final captured = verify(() => cat.likes(captureAny(that: startsWith('d')))).captured;
expect(captured.last, equals(['dog food']));

```

copied to clipboard
## Resetting Mocks [#](https://pub.dev/packages/mocktail#resetting-mocks)
```
reset(cat); // Reset stubs and interactions

```

copied to clipboard
## How it works [#](https://pub.dev/packages/mocktail#how-it-works)
Mocktail uses closures to handle catching `TypeError` instances which would otherwise propagate and cause test failures when stubbing/verifying non-nullable return types. Check out [#24](https://github.com/felangel/mocktail/issues/24) for more information.
In order to support argument matchers such as `any` and `captureAny` mocktail has to register default fallback values to return when the argument matchers are used. Out of the box, it automatically handles all primitive types, however, when using argument matchers in place of custom types developers must use `registerFallbackValue` to provide a default return value. It is only required to call `registerFallbackValue` once per type so it is recommended to place all `registerFallbackValue` calls within `setUpAll`.
```
class Food {...}
class Cat {
 bool likes(Food food) {...}
}
...
class MockCat extends Mock implements Cat {}
class FakeFood extends Fake implements Food {}
void main() {
 setUpAll(() {
  registerFallbackValue(FakeFood());
 });
 test('...', () {
  final cat = MockCat();
  when(() => cat.likes(any()).thenReturn(true);
  ...
 });
}

```

copied to clipboard
## FAQs [#](https://pub.dev/packages/mocktail#faqs)
#### Why am I getting an invalid_implementation_override error when trying to Fake certain classes like ThemeData and ColorScheme?
[Relevant Issue](https://github.com/felangel/mocktail/issues/59)
This is likely due to differences in the function signature of `toString` for the class and can be resolved using a mixin as demonstrated below:
```
mixin DiagnosticableToStringMixin on Object {
 @override
 String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
  return super.toString();
 }
}
class FakeThemeData extends Fake
 with DiagnosticableToStringMixin
 implements ThemeData {}

```

copied to clipboard
#### Why can't I stub/verify extension methods properly?
[Relevant Issue](https://github.com/felangel/mocktail/issues/58)
Extension methods cannot be stubbed/verified as they are treated like static methods. This means that calls go directly to the extension method without caring about the instance. As a result, stubs and verify calls to extensions always result in an invocation of the real extension method.
Instead of stubbing/verifying extension methods directly, prefer to stub/verify public members on the instance with which the extension methods interact.
#### Why am I seeing error: type 'Null' is not a subtype of type 'Future
[Relevant Issue](https://github.com/felangel/mocktail/issues/78)
By default when a class extends `Mock` any unstubbed methods return `null`.
For example, take the following class:
```
class Person {
 Future<void> sleep() {
  await Future<void>.delayed(Duration(hours: 8));
 }
}

```

copied to clipboard
We can create a `MockPerson` like:
```
class MockPerson extends Mock implements Person {}

```

copied to clipboard
If we have code that invokes `sleep` on `MockPerson` we will get a `TypeError`:
```
type 'Null' is not a subtype of type 'Future<void>'

```

copied to clipboard
This is because we did not stub `sleep` so when `sleep` is called on an instance of `MockPerson`, `mocktail` returns `null` which is not compatible with `Future<void>`.
To address this, we must explicitly stub `sleep` like:
```
final person = MockPerson();
when(() => person.sleep()).thenAnswer((_) async {});

```

copied to clipboard
#### Why is my method throwing a `TypeError` when stubbing using `any()`?
[Relevant Issue](https://github.com/felangel/mocktail/issues/162)
By default when a class extends `Mock` any unstubbed methods return `null`. When stubbing using `any()` the type must be inferable. However, when a method has a generic type argument it might not be able to infer the type and as a result, the generic would fallback to `dynamic` causing the method to act as if it was unstubbed.
For example, take the following class and its method:
```
class Cache {
 bool set<T>(String key, T value) {
  return true;
 }
}

```

copied to clipboard
The following stub will be equivalent to calling `set<dynamic>(...)`:
```
// The type `T` of `any<T>()` is inferred to be `dynamic`.
when(() => cache.set(any(), any())).thenReturn((_) => true);

```

copied to clipboard
To address this, we must explicitly stub `set` with a type:
```
final cache = MockCache();
when(() => cache.set<int>(any(), any())).thenReturn((_) => true);
cache.set<int>('key', 1);
verify(() => cache.set<int>(any(), any())).called(1);

```

copied to clipboard
The type doesn't need to be applied to `set<T>()`, any explicit type that allows `any<T>()` infer its type will allow the method to be stubbed for that type:
```
when(() => cache.set(any(), any<int>())).thenReturn((_) => true);

```

copied to clipboard
[1.19klikes160points2.65Mdownloads](https://pub.dev/packages/mocktail/score)
### Publisher
[![verified publisher](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-verified.svg)felangel.dev](https://pub.dev/publishers/felangel.dev)
### Weekly Downloads
2024.09.21 - 2025.08.16
### Metadata
A Dart mock library which simplifies mocking with null safety support and no manual mocks or code generation.
[Homepage](https://github.com/felangel/mocktail/tree/main/packages/mocktail)[Repository (GitHub)](https://github.com/felangel/mocktail)
### Topics
[#mock](https://pub.dev/packages?q=topic%3Amock) [#testing](https://pub.dev/packages?q=topic%3Atesting "Packages that facilitate testing.")
### Documentation
[API reference](https://pub.dev/documentation/mocktail/latest/)
### License
![](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-balance.svg)MIT ([license](https://pub.dev/packages/mocktail/license))
### Dependencies
[collection](https://pub.dev/packages/collection "^1.15.0"), [matcher](https://pub.dev/packages/matcher "^0.12.15"), [test_api](https://pub.dev/packages/test_api ">=0.2.1 <0.8.0")
### More
[Packages that depend on mocktail](https://pub.dev/packages?q=dependency%3Amocktail)
### ← Metadata
[1.19klikes160points2.65Mdownloads](https://pub.dev/packages/mocktail/score)
### Publisher
[![verified publisher](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-verified.svg)felangel.dev](https://pub.dev/publishers/felangel.dev)
### Weekly Downloads
2024.09.21 - 2025.08.16
### Metadata
A Dart mock library which simplifies mocking with null safety support and no manual mocks or code generation.
[Homepage](https://github.com/felangel/mocktail/tree/main/packages/mocktail)[Repository (GitHub)](https://github.com/felangel/mocktail)
### Topics
[#mock](https://pub.dev/packages?q=topic%3Amock) [#testing](https://pub.dev/packages?q=topic%3Atesting "Packages that facilitate testing.")
### Documentation
[API reference](https://pub.dev/documentation/mocktail/latest/)
### License
![](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-balance.svg)MIT ([license](https://pub.dev/packages/mocktail/license))
### Dependencies
[collection](https://pub.dev/packages/collection "^1.15.0"), [matcher](https://pub.dev/packages/matcher "^0.12.15"), [test_api](https://pub.dev/packages/test_api ">=0.2.1 <0.8.0")
### More
[Packages that depend on mocktail](https://pub.dev/packages?q=dependency%3Amocktail)
Back
![previous](https://pub.dev/static/hash-0mq6dpb8/img/keyboard_arrow_left.svg)![next](https://pub.dev/static/hash-0mq6dpb8/img/keyboard_arrow_right.svg)
[Dart language](https://dart.dev/)[Report package](https://pub.dev/report?subject=package%3Amocktail&url=https%3A%2F%2Fpub.dev%2Fpackages%2Fmocktail)[Policy](https://pub.dev/policy)[Terms](https://www.google.com/intl/en/policies/terms/)[API Terms](https://developers.google.com/terms/)[Security](https://pub.dev/security)[Privacy](https://www.google.com/intl/en/policies/privacy/)[Help](https://pub.dev/help)[![RSS](https://pub.dev/static/hash-0mq6dpb8/img/rss-feed-icon.svg)](https://pub.dev/feed.atom)[![bug report](https://pub.dev/static/hash-0mq6dpb8/img/bug-report-white-96px.png)](https://github.com/dart-lang/pub-dev/issues/new?body=URL%3A+https%3A%2F%2Fpub.dev%2Fpackages%2Fmocktail%0A%0A%3CDescribe+your+issue+or+suggestion+here%3E&title=%3CSummarize+your+issues+here%3E&labels=Area%3A+site+feedback)


## Source Information
- URL: https://pub.dev/packages/mocktail
- Harvested: 2025-08-19T23:20:49.764333
- Profile: code_examples
- Agent: architect
