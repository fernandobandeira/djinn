---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:21:33.011063'
profile: code_examples
source: https://pub.dev/packages/mockito
topic: Mockito vs Mocktail Comparison Analysis
---

# Mockito vs Mocktail Comparison Analysis

This site uses cookies from Google to deliver and enhance the quality of its services and to analyze traffic.[Learn more](https://policies.google.com/technologies/cookies?hl=en)[OK, got it](https://pub.dev/packages/mockito)
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
# mockito 5.5.0 ![copy "mockito: ^5.5.0" to clipboard](https://pub.dev/static/hash-0mq6dpb8/img/content-copy-icon.svg)
mockito: ^5.5.0 copied to clipboard
Published [29 days ago](https://pub.dev/packages/mockito "Jul 21, 2025") • [![verified publisher](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-verified.svg)dart.dev](https://pub.dev/publishers/dart.dev)Dart 3 compatible
SDK[Dart](https://pub.dev/packages?q=sdk%3Adart "Packages compatible with Dart SDK")[Flutter](https://pub.dev/packages?q=sdk%3Aflutter "Packages compatible with Flutter SDK")
Platform[Android](https://pub.dev/packages?q=platform%3Aandroid "Packages compatible with Android platform")[iOS](https://pub.dev/packages?q=platform%3Aios "Packages compatible with iOS platform")[Linux](https://pub.dev/packages?q=platform%3Alinux "Packages compatible with Linux platform")[macOS](https://pub.dev/packages?q=platform%3Amacos "Packages compatible with macOS platform")[web](https://pub.dev/packages?q=platform%3Aweb "Packages compatible with Web platform")[Windows](https://pub.dev/packages?q=platform%3Awindows "Packages compatible with Windows platform")
![liked status: inactive](https://pub.dev/static/hash-0mq6dpb8/img/like-inactive.svg)![liked status: active](https://pub.dev/static/hash-0mq6dpb8/img/like-active.svg)1.4k
→
### Metadata
A mock framework inspired by Mockito with APIs for Fakes, Mocks, behavior verification, and stubbing.
More...
  * Readme
  * [Changelog](https://pub.dev/packages/mockito/changelog)
  * [Example](https://pub.dev/packages/mockito/example)
  * [Installing](https://pub.dev/packages/mockito/install)
  * [Versions](https://pub.dev/packages/mockito/versions)
  * [Scores](https://pub.dev/packages/mockito/score)


[![Dart CI](https://github.com/dart-lang/mockito/actions/workflows/test-package.yml/badge.svg)](https://github.com/dart-lang/mockito/actions/workflows/test-package.yml) [![Pub](https://img.shields.io/pub/v/mockito.svg)](https://pub.dev/packages/mockito) [![package publisher](https://img.shields.io/pub/publisher/mockito.svg)](https://pub.dev/packages/mockito/publisher)
Mock library for Dart inspired by [Mockito](https://github.com/mockito/mockito).
## Let's create mocks [#](https://pub.dev/packages/mockito#lets-create-mocks)
Mockito 5.0.0 supports Dart's new **null safety** language feature in Dart 2.12, primarily with code generation.
To use Mockito's generated mock classes, add a `build_runner` dependency in your package's `pubspec.yaml` file, under `dev_dependencies`; something like `build_runner: ^1.11.0`.
For alternatives to the code generation API, see the [NULL_SAFETY_README](https://github.com/dart-lang/mockito/blob/master/NULL_SAFETY_README.md).
Let's start with a Dart library, `cat.dart`:
```
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
// Annotation which generates the cat.mocks.dart library and the MockCat class.
@GenerateNiceMocks([MockSpec<Cat>()])
import 'cat.mocks.dart';
// Real class
class Cat {
 String sound() => "Meow";
 bool eatFood(String food, {bool? hungry}) => true;
 Future<void> chew() async => print("Chewing...");
 int walk(List<String> places) => 7;
 void sleep() {}
 void hunt(String place, String prey) {}
 int lives = 9;
}
void main() {
 // Create mock object.
 var cat = MockCat();
}

```

copied to clipboard
By annotating the import of a `.mocks.dart` library with `@GenerateNiceMocks`, you are directing Mockito's code generation to write a mock class for each "real" class listed, in a new library.
The next step is to run `build_runner` in order to generate this new library:
```
flutter pub run build_runner build
# OR
dart run build_runner build

```

copied to clipboard
`build_runner` will generate a file with a name based on the file containing the `@GenerateNiceMocks` annotation. In the above `cat.dart` example, we import the generated library as `cat.mocks.dart`.
**NOTE** : by default only annotations in files under `test/` are processed, if you want to add Mockito annotations in other places, you will need to add a `build.yaml` file to your project, see [this SO answer](https://stackoverflow.com/questions/68275811/is-there-a-way-to-let-mockito-generate-mocks-for-integration-tests-in-a-flutter).
The generated mock class, `MockCat`, extends Mockito's Mock class and implements the Cat class, giving us a class which supports stubbing and verifying.
## Let's verify some behavior! [#](https://pub.dev/packages/mockito#lets-verify-some-behavior)
```
// Interact with the mock object.
cat.sound();
// Verify the interaction.
verify(cat.sound());

```

copied to clipboard
Once created, the mock instance will remember all interactions. Then you can selectively [`verify`](https://pub.dev/documentation/mockito/latest/mockito/verify.html) (or [`verifyInOrder`](https://pub.dev/documentation/mockito/latest/mockito/verifyInOrder.html), or [`verifyNever`](https://pub.dev/documentation/mockito/latest/mockito/verifyNever.html)) the interactions you are interested in.
## How about some stubbing? [#](https://pub.dev/packages/mockito#how-about-some-stubbing)
```
// Stub a mock method before interacting.
when(cat.sound()).thenReturn("Purr");
expect(cat.sound(), "Purr");
// You can call it again.
expect(cat.sound(), "Purr");
// Let's change the stub.
when(cat.sound()).thenReturn("Meow");
expect(cat.sound(), "Meow");
// You can stub getters.
when(cat.lives).thenReturn(9);
expect(cat.lives, 9);
// You can stub a method to throw.
when(cat.lives).thenThrow(RangeError('Boo'));
expect(() => cat.lives, throwsRangeError);
// We can calculate a response at call time.
var responses = ["Purr", "Meow"];
when(cat.sound()).thenAnswer((_) => responses.removeAt(0));
expect(cat.sound(), "Purr");
expect(cat.sound(), "Meow");
// We can stub a method with multiple calls that happened in a particular order.
when(cat.sound()).thenReturnInOrder(["Purr", "Meow"]);
expect(cat.sound(), "Purr");
expect(cat.sound(), "Meow");
expect(() => cat.sound(), throwsA(isA<StateError>()));

```

copied to clipboard
The [`when`](https://pub.dev/documentation/mockito/latest/mockito/when.html), [`thenReturn`](https://pub.dev/documentation/mockito/latest/mockito/PostExpectation/thenReturn.html), [`thenAnswer`](https://pub.dev/documentation/mockito/latest/mockito/PostExpectation/thenAnswer.html), and [`thenThrow`](https://pub.dev/documentation/mockito/latest/mockito/PostExpectation/thenThrow.html) APIs provide a stubbing mechanism to override this behavior. Once stubbed, the method will always return stubbed value regardless of how many times it is called. If a method invocation matches multiple stubs, the one which was declared last will be used. It is worth noting that stubbing and verifying only works on methods of a mocked class; in this case, an instance of `MockCat` must be used, not an instance of `Cat`.
### A quick word on async stubbing [#](https://pub.dev/packages/mockito#a-quick-word-on-async-stubbing)
**Using[`thenReturn`](https://pub.dev/documentation/mockito/latest/mockito/PostExpectation/thenReturn.html) to return a `Future` or `Stream` will throw an `ArgumentError`.** This is because it can lead to unexpected behaviors. For example:
  * If the method is stubbed in a different zone than the zone that consumes the `Future`, unexpected behavior could occur.
  * If the method is stubbed to return a failed `Future` or `Stream` and it doesn't get consumed in the same run loop, it might get consumed by the global exception handler instead of an exception handler the consumer applies.


Instead, use `thenAnswer` to stub methods that return a `Future` or `Stream`.
```
// BAD
when(mock.methodThatReturnsAFuture())
  .thenReturn(Future.value('Stub'));
when(mock.methodThatReturnsAStream())
  .thenReturn(Stream.fromIterable(['Stub']));
// GOOD
when(mock.methodThatReturnsAFuture())
  .thenAnswer((_) async => 'Stub');
when(mock.methodThatReturnsAStream())
  .thenAnswer((_) => Stream.fromIterable(['Stub']));

```

copied to clipboard
If, for some reason, you desire the behavior of `thenReturn`, you can return a pre-defined instance.
```
// Use the above method unless you're sure you want to create the Future ahead
// of time.
final future = Future.value('Stub');
when(mock.methodThatReturnsAFuture()).thenAnswer((_) => future);

```

copied to clipboard
## Argument matchers [#](https://pub.dev/packages/mockito#argument-matchers)
Mockito provides the concept of the "argument matcher" (using the class ArgMatcher) to capture arguments and to track how named arguments are passed. In most cases, both plain arguments and argument matchers can be passed into mock methods:
```
// You can use `any`
when(cat.eatFood(any)).thenReturn(false);
// ... or plain arguments themselves
when(cat.eatFood("fish")).thenReturn(true);
// ... including collections
when(cat.walk(["roof","tree"])).thenReturn(2);
// ... or matchers
when(cat.eatFood(argThat(startsWith("dry")))).thenReturn(false);
// ... or mix arguments with matchers
when(cat.eatFood(argThat(startsWith("dry")), hungry: true)).thenReturn(true);
expect(cat.eatFood("fish"), isTrue);
expect(cat.walk(["roof","tree"]), equals(2));
expect(cat.eatFood("dry food"), isFalse);
expect(cat.eatFood("dry food", hungry: true), isTrue);
// You can also verify using an argument matcher.
verify(cat.eatFood("fish"));
verify(cat.walk(["roof","tree"]));
verify(cat.eatFood(argThat(contains("food"))));
// You can verify setters.
cat.lives = 9;
verify(cat.lives=9);

```

copied to clipboard
If an argument other than an ArgMatcher (like [`any`](https://pub.dev/documentation/mockito/latest/mockito/any.html), [`anyNamed`](https://pub.dev/documentation/mockito/latest/mockito/anyNamed.html), [`argThat`](https://pub.dev/documentation/mockito/latest/mockito/argThat.html), [`captureThat`](https://pub.dev/documentation/mockito/latest/mockito/captureThat.html), etc.) is passed to a mock method, then the [`equals`](https://pub.dev/documentation/matcher/latest/matcher/equals.html) matcher is used for argument matching. If you need more strict matching, consider using `argThat(identical(arg))`.
However, note that `null` cannot be used as an argument adjacent to ArgMatcher arguments, nor as an un-wrapped value passed as a named argument. For example:
```
verify(cat.hunt("backyard", null)); // OK: no arg matchers.
verify(cat.hunt(argThat(contains("yard")), null)); // BAD: adjacent null.
verify(cat.hunt(argThat(contains("yard")), argThat(isNull))); // OK: wrapped in an arg matcher.
verify(cat.eatFood("Milk", hungry: null)); // BAD: null as a named argument.
verify(cat.eatFood("Milk", hungry: argThat(isNull))); // BAD: null as a named argument.

```

copied to clipboard
## Named arguments [#](https://pub.dev/packages/mockito#named-arguments)
Mockito currently has an awkward nuisance to its syntax: named arguments and argument matchers require more specification than you might think: you must declare the name of the argument in the argument matcher. This is because we can't rely on the position of a named argument, and the language doesn't provide a mechanism to answer "Is this element being used as a named element?"
```
// GOOD: argument matchers include their names.
when(cat.eatFood(any, hungry: anyNamed('hungry'))).thenReturn(true);
when(cat.eatFood(any, hungry: argThat(isNotNull, named: 'hungry'))).thenReturn(false);
when(cat.eatFood(any, hungry: captureAnyNamed('hungry'))).thenReturn(false);
when(cat.eatFood(any, hungry: captureThat(isNotNull, named: 'hungry'))).thenReturn(true);
// BAD: argument matchers do not include their names.
when(cat.eatFood(any, hungry: any)).thenReturn(true);
when(cat.eatFood(any, hungry: argThat(isNotNull))).thenReturn(false);
when(cat.eatFood(any, hungry: captureAny)).thenReturn(false);
when(cat.eatFood(any, hungry: captureThat(isNotNull))).thenReturn(true);

```

copied to clipboard
## Verifying exact number of invocations / at least x / never [#](https://pub.dev/packages/mockito#verifying-exact-number-of-invocations--at-least-x--never)
Use [`verify`](https://pub.dev/documentation/mockito/latest/mockito/verify.html) or [`verifyNever`](https://pub.dev/documentation/mockito/latest/mockito/verifyNever.html):
```
cat.sound();
cat.sound();
// Exact number of invocations
verify(cat.sound()).called(2);
// Or using matcher
verify(cat.sound()).called(greaterThan(1));
// Or never called
verifyNever(cat.eatFood(any));

```

copied to clipboard
## Verification in order [#](https://pub.dev/packages/mockito#verification-in-order)
Use [`verifyInOrder`](https://pub.dev/documentation/mockito/latest/mockito/verifyInOrder.html):
```
cat.eatFood("Milk");
cat.sound();
cat.eatFood("Fish");
verifyInOrder([
 cat.eatFood("Milk"),
 cat.sound(),
 cat.eatFood("Fish")
]);

```

copied to clipboard
Verification in order is flexible - you don't have to verify all interactions one-by-one but only those that you are interested in testing in order.
## Making sure interaction(s) never happened on mock [#](https://pub.dev/packages/mockito#making-sure-interactions-never-happened-on-mock)
Use [`verifyZeroInteractions`](https://pub.dev/documentation/mockito/latest/mockito/verifyZeroInteractions.html):
```
verifyZeroInteractions(cat);

```

copied to clipboard
## Finding redundant invocations [#](https://pub.dev/packages/mockito#finding-redundant-invocations)
Use [`verifyNoMoreInteractions`](https://pub.dev/documentation/mockito/latest/mockito/verifyNoMoreInteractions.html):
```
cat.sound();
verify(cat.sound());
verifyNoMoreInteractions(cat);

```

copied to clipboard
## Capturing arguments for further assertions [#](https://pub.dev/packages/mockito#capturing-arguments-for-further-assertions)
Use the [`captureAny`](https://pub.dev/documentation/mockito/latest/mockito/captureAny.html), [`captureThat`](https://pub.dev/documentation/mockito/latest/mockito/captureThat.html), and [`captureAnyNamed`](https://pub.dev/documentation/mockito/latest/mockito/captureAnyNamed.html) argument matchers:
```
// Simple capture
cat.eatFood("Fish");
expect(verify(cat.eatFood(captureAny)).captured.single, "Fish");
// Capture multiple calls
cat.eatFood("Milk");
cat.eatFood("Fish");
expect(verify(cat.eatFood(captureAny)).captured, ["Milk", "Fish"]);
// Conditional capture
cat.eatFood("Milk");
cat.eatFood("Fish");
expect(verify(cat.eatFood(captureThat(startsWith("F")))).captured, ["Fish"]);

```

copied to clipboard
## Waiting for an interaction [#](https://pub.dev/packages/mockito#waiting-for-an-interaction)
Use [`untilCalled`](https://pub.dev/documentation/mockito/latest/mockito/untilCalled.html):
```
// Waiting for a call.
cat.eatFood("Fish");
await untilCalled(cat.chew()); // Completes when cat.chew() is called.
// Waiting for a call that has already happened.
cat.eatFood("Fish");
await untilCalled(cat.eatFood(any)); // Completes immediately.

```

copied to clipboard
## Nice mocks vs classic mocks [#](https://pub.dev/packages/mockito#nice-mocks-vs-classic-mocks)
Mockito provides two APIs for generating mocks, the `@GenerateNiceMocks` annotation and the `@GenerateMocks` annotation. **The recommended API is`@GenerateNiceMocks`.** The difference between these two APIs is in the behavior of a generated mock class when a method is called and no stub could be found. For example:
```
void main() {
 var cat = MockCat();
 cat.sound();
}

```

copied to clipboard
The `Cat.sound` method returns a non-nullable String, but no stub has been made with `when(cat.sound())`, so what should the code do? What is the "missing stub" behavior?
  * The "missing stub" behavior of a mock class generated with `@GenerateMocks` is to throw an exception.
  * The "missing stub" behavior of a mock class generated with `@GenerateNiceMocks` is to return a "simple" legal value (for example, a non-`null` value for a non-nullable return type). The value should not be used in any way; it is returned solely to avoid a runtime type exception.


## Mocking a Function type [#](https://pub.dev/packages/mockito#mocking-a-function-type)
To create mocks for Function objects, write an `abstract class` with a method for each function type signature that needs to be mocked. The methods can be torn off and individually stubbed and verified.
```
@GenerateMocks([Cat, Callbacks])
import 'cat_test.mocks.dart'
abstract class Callbacks {
 Cat findCat(String name);
}
void main() {
 var mockCat = MockCat();
 var findCatCallback = MockCallbacks().findCat;
 when(findCatCallback('Pete')).thenReturn(mockCat);
}

```

copied to clipboard
## Writing a fake [#](https://pub.dev/packages/mockito#writing-a-fake)
You can also write a simple fake class that implements a real class, by extending [Fake](https://pub.dev/documentation/mockito/latest/mockito/Fake-class.html). Fake allows your subclass to satisfy the implementation of your real class, without overriding the methods that aren't used in your test; the Fake class implements the default behavior of throwing [UnimplementedError](https://api.dartlang.org/stable/dart-core/UnimplementedError-class.html) (which you can override in your fake class):
```
// Fake class
class FakeCat extends Fake implements Cat {
 @override
 bool eatFood(String food, {bool? hungry}) {
  print('Fake eat $food');
  return true;
 }
}
void main() {
 // Create a new fake Cat at runtime.
 var cat = FakeCat();
 cat.eatFood("Milk"); // Prints 'Fake eat Milk'.
 cat.sleep(); // Throws.
}

```

copied to clipboard
## Resetting mocks [#](https://pub.dev/packages/mockito#resetting-mocks)
Use [`reset`](https://pub.dev/documentation/mockito/latest/mockito/reset.html):
```
// Clearing collected interactions:
cat.eatFood("Fish");
clearInteractions(cat);
cat.eatFood("Fish");
verify(cat.eatFood("Fish")).called(1);
// Resetting stubs and collected interactions:
when(cat.eatFood("Fish")).thenReturn(true);
cat.eatFood("Fish");
reset(cat);
when(cat.eatFood(any)).thenReturn(false);
expect(cat.eatFood("Fish"), false);

```

copied to clipboard
## Debugging [#](https://pub.dev/packages/mockito#debugging)
Use [`logInvocations`](https://pub.dev/documentation/mockito/latest/mockito/logInvocations.html) and [`throwOnMissingStub`](https://pub.dev/documentation/mockito/latest/mockito/throwOnMissingStub.html):
```
// Print all collected invocations of any mock methods of a list of mock objects:
logInvocations([catOne, catTwo]);
// Throw every time that a mock method is called without a stub being matched:
throwOnMissingStub(cat);

```

copied to clipboard
## Best Practices [#](https://pub.dev/packages/mockito#best-practices)
Testing with real objects is preferred over testing with mocks - if you can construct a real instance for your tests, you should! If there are no calls to [`verify`](https://pub.dev/documentation/mockito/latest/mockito/verify.html) in your test, it is a strong signal that you may not need mocks at all, though it's also OK to use a `Mock` like a stub. Data models never need to be mocked if they can be constructed with stubbed data. When it's not possible to use the real object, a tested implementation of a fake is the next best thing; it's more likely to behave similarly to the real class than responses stubbed out in tests. Finally an object which `extends Fake` using manually overridden methods is preferred over an object which `extends Mock` used as either a stub or a mock.
A class which `extends Mock` should _never_ stub out its own responses with `when` in its constructor or anywhere else. Stubbed responses should be defined in the tests where they are used. For responses controlled outside of the test use `@override` methods for either the entire interface, or with `extends Fake` to skip some parts of the interface.
Similarly, a class which `extends Mock` should _never_ have any implementation. It should not define any `@override` methods, and it should not mixin any implementations. Actual member definitions can't be stubbed by tests and can't be tracked and verified by Mockito. A mix of test defined stubbed responses and mock defined overrides will lead to confusion. It is OK to define _static_ utilities on a class which `extends Mock` if it helps with code structure.
## Frequently asked questions [#](https://pub.dev/packages/mockito#frequently-asked-questions)
Read more information about this package in the [FAQ](https://github.com/dart-lang/mockito/blob/master/FAQ.md).
[1.49klikes160points1.84Mdownloads](https://pub.dev/packages/mockito/score)
### Publisher
[![verified publisher](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-verified.svg)dart.dev](https://pub.dev/publishers/dart.dev)
### Weekly Downloads
2024.09.21 - 2025.08.16
### Metadata
A mock framework inspired by Mockito with APIs for Fakes, Mocks, behavior verification, and stubbing.
[Repository (GitHub)](https://github.com/dart-lang/mockito)[Contributing](https://github.com/dart-lang/mockito/blob/master/CONTRIBUTING.md)
### Topics
[#testing](https://pub.dev/packages?q=topic%3Atesting "Packages that facilitate testing.") [#mocking](https://pub.dev/packages?q=topic%3Amocking)
### Documentation
[API reference](https://pub.dev/documentation/mockito/latest/)
### License
![](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-balance.svg)Apache-2.0 ([license](https://pub.dev/packages/mockito/license))
### Dependencies
[analyzer](https://pub.dev/packages/analyzer ">=7.5.5 <8.0.0"), [build](https://pub.dev/packages/build "^3.0.0"), [code_builder](https://pub.dev/packages/code_builder "^4.5.0"), [collection](https://pub.dev/packages/collection "^1.19.0"), [dart_style](https://pub.dev/packages/dart_style ">=2.3.7 <4.0.0"), [matcher](https://pub.dev/packages/matcher "^0.12.16"), [meta](https://pub.dev/packages/meta "^1.15.0"), [path](https://pub.dev/packages/path "^1.9.0"), [source_gen](https://pub.dev/packages/source_gen "^3.0.0"), [test_api](https://pub.dev/packages/test_api ">=0.6.1 <0.8.0")
### More
[Packages that depend on mockito](https://pub.dev/packages?q=dependency%3Amockito)
### ← Metadata
[1.49klikes160points1.84Mdownloads](https://pub.dev/packages/mockito/score)
### Publisher
[![verified publisher](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-verified.svg)dart.dev](https://pub.dev/publishers/dart.dev)
### Weekly Downloads
2024.09.21 - 2025.08.16
### Metadata
A mock framework inspired by Mockito with APIs for Fakes, Mocks, behavior verification, and stubbing.
[Repository (GitHub)](https://github.com/dart-lang/mockito)[Contributing](https://github.com/dart-lang/mockito/blob/master/CONTRIBUTING.md)
### Topics
[#testing](https://pub.dev/packages?q=topic%3Atesting "Packages that facilitate testing.") [#mocking](https://pub.dev/packages?q=topic%3Amocking)
### Documentation
[API reference](https://pub.dev/documentation/mockito/latest/)
### License
![](https://pub.dev/static/hash-0mq6dpb8/img/material-icon-balance.svg)Apache-2.0 ([license](https://pub.dev/packages/mockito/license))
### Dependencies
[analyzer](https://pub.dev/packages/analyzer ">=7.5.5 <8.0.0"), [build](https://pub.dev/packages/build "^3.0.0"), [code_builder](https://pub.dev/packages/code_builder "^4.5.0"), [collection](https://pub.dev/packages/collection "^1.19.0"), [dart_style](https://pub.dev/packages/dart_style ">=2.3.7 <4.0.0"), [matcher](https://pub.dev/packages/matcher "^0.12.16"), [meta](https://pub.dev/packages/meta "^1.15.0"), [path](https://pub.dev/packages/path "^1.9.0"), [source_gen](https://pub.dev/packages/source_gen "^3.0.0"), [test_api](https://pub.dev/packages/test_api ">=0.6.1 <0.8.0")
### More
[Packages that depend on mockito](https://pub.dev/packages?q=dependency%3Amockito)
Back
![previous](https://pub.dev/static/hash-0mq6dpb8/img/keyboard_arrow_left.svg)![next](https://pub.dev/static/hash-0mq6dpb8/img/keyboard_arrow_right.svg)
[Dart language](https://dart.dev/)[Report package](https://pub.dev/report?subject=package%3Amockito&url=https%3A%2F%2Fpub.dev%2Fpackages%2Fmockito)[Policy](https://pub.dev/policy)[Terms](https://www.google.com/intl/en/policies/terms/)[API Terms](https://developers.google.com/terms/)[Security](https://pub.dev/security)[Privacy](https://www.google.com/intl/en/policies/privacy/)[Help](https://pub.dev/help)[![RSS](https://pub.dev/static/hash-0mq6dpb8/img/rss-feed-icon.svg)](https://pub.dev/feed.atom)[![bug report](https://pub.dev/static/hash-0mq6dpb8/img/bug-report-white-96px.png)](https://github.com/dart-lang/pub-dev/issues/new?body=URL%3A+https%3A%2F%2Fpub.dev%2Fpackages%2Fmockito%0A%0A%3CDescribe+your+issue+or+suggestion+here%3E&title=%3CSummarize+your+issues+here%3E&labels=Area%3A+site+feedback)


## Source Information
- URL: https://pub.dev/packages/mockito
- Harvested: 2025-08-19T23:21:33.011063
- Profile: code_examples
- Agent: architect
