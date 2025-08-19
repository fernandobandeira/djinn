---
agent_context: test
confidence: 0.95
harvested_at: '2025-08-19T01:08:39.569800'
profile: quick_reference
source: https://docs.python.org/3/library/typing.html
topic: Python Type Hints
---

# Python Type Hints

[ ![Python logo](https://docs.python.org/3/_static/py.svg) ](https://www.python.org/) dev (3.15)pre (3.14)3.13.73.123.113.103.93.83.73.63.53.43.33.23.13.02.72.6
Greek | ΕλληνικάEnglishSpanish | españolFrench | françaisItalian | italianoJapanese | 日本語Korean | 한국어Polish | polskiBrazilian Portuguese | Português brasileiroTurkish | TürkçeSimplified Chinese | 简体中文Traditional Chinese | 繁體中文
Theme  Auto Light Dark
### [Table of Contents](https://docs.python.org/3/contents.html)
  * [`typing` — Support for type hints](https://docs.python.org/3/library/typing.html)
    * [Specification for the Python Type System](https://docs.python.org/3/library/typing.html#specification-for-the-python-type-system)
    * [Type aliases](https://docs.python.org/3/library/typing.html#type-aliases)
    * [NewType](https://docs.python.org/3/library/typing.html#newtype)
    * [Annotating callable objects](https://docs.python.org/3/library/typing.html#annotating-callable-objects)
    * [Generics](https://docs.python.org/3/library/typing.html#generics)
    * [Annotating tuples](https://docs.python.org/3/library/typing.html#annotating-tuples)
    * [The type of class objects](https://docs.python.org/3/library/typing.html#the-type-of-class-objects)
    * [Annotating generators and coroutines](https://docs.python.org/3/library/typing.html#annotating-generators-and-coroutines)
    * [User-defined generic types](https://docs.python.org/3/library/typing.html#user-defined-generic-types)
    * [The `Any` type](https://docs.python.org/3/library/typing.html#the-any-type)
    * [Nominal vs structural subtyping](https://docs.python.org/3/library/typing.html#nominal-vs-structural-subtyping)
    * [Module contents](https://docs.python.org/3/library/typing.html#module-contents)
      * [Special typing primitives](https://docs.python.org/3/library/typing.html#special-typing-primitives)
        * [Special types](https://docs.python.org/3/library/typing.html#special-types)
        * [Special forms](https://docs.python.org/3/library/typing.html#special-forms)
        * [Building generic types and type aliases](https://docs.python.org/3/library/typing.html#building-generic-types-and-type-aliases)
        * [Other special directives](https://docs.python.org/3/library/typing.html#other-special-directives)
      * [Protocols](https://docs.python.org/3/library/typing.html#protocols)
      * [ABCs for working with IO](https://docs.python.org/3/library/typing.html#abcs-for-working-with-io)
      * [Functions and decorators](https://docs.python.org/3/library/typing.html#functions-and-decorators)
      * [Introspection helpers](https://docs.python.org/3/library/typing.html#introspection-helpers)
      * [Constant](https://docs.python.org/3/library/typing.html#constant)
      * [Deprecated aliases](https://docs.python.org/3/library/typing.html#deprecated-aliases)
        * [Aliases to built-in types](https://docs.python.org/3/library/typing.html#aliases-to-built-in-types)
        * [Aliases to types in `collections`](https://docs.python.org/3/library/typing.html#aliases-to-types-in-collections)
        * [Aliases to other concrete types](https://docs.python.org/3/library/typing.html#aliases-to-other-concrete-types)
        * [Aliases to container ABCs in `collections.abc`](https://docs.python.org/3/library/typing.html#aliases-to-container-abcs-in-collections-abc)
        * [Aliases to asynchronous ABCs in `collections.abc`](https://docs.python.org/3/library/typing.html#aliases-to-asynchronous-abcs-in-collections-abc)
        * [Aliases to other ABCs in `collections.abc`](https://docs.python.org/3/library/typing.html#aliases-to-other-abcs-in-collections-abc)
        * [Aliases to `contextlib` ABCs](https://docs.python.org/3/library/typing.html#aliases-to-contextlib-abcs)
    * [Deprecation Timeline of Major Features](https://docs.python.org/3/library/typing.html#deprecation-timeline-of-major-features)


#### Previous topic
[Development Tools](https://docs.python.org/3/library/development.html "previous chapter")
#### Next topic
[`pydoc` — Documentation generator and online help system](https://docs.python.org/3/library/pydoc.html "next chapter")
### This page
  * [Report a bug](https://docs.python.org/3/bugs.html)
  * [Show source ](https://github.com/python/cpython/blob/main/Doc/library/typing.rst?plain=1)


### Navigation
  * [index](https://docs.python.org/3/genindex.html "General Index")
  * [modules](https://docs.python.org/3/py-modindex.html "Python Module Index") |
  * [next](https://docs.python.org/3/library/pydoc.html "pydoc — Documentation generator and online help system") |
  * [previous](https://docs.python.org/3/library/development.html "Development Tools") |
  * ![Python logo](https://docs.python.org/3/_static/py.svg)
  * [Python](https://www.python.org/) »
  * Greek | ΕλληνικάEnglishSpanish | españolFrench | françaisItalian | italianoJapanese | 日本語Korean | 한국어Polish | polskiBrazilian Portuguese | Português brasileiroTurkish | TürkçeSimplified Chinese | 简体中文Traditional Chinese | 繁體中文
dev (3.15)pre (3.14)3.13.73.123.113.103.93.83.73.63.53.43.33.23.13.02.72.6
  * [3.13.7 Documentation](https://docs.python.org/3/index.html) » 
  * [The Python Standard Library](https://docs.python.org/3/library/index.html) »
  * [Development Tools](https://docs.python.org/3/library/development.html) »
  * [`typing` — Support for type hints](https://docs.python.org/3/library/typing.html)
  * | 
  * Theme  Auto Light Dark |


# `typing` — Support for type hints[¶](https://docs.python.org/3/library/typing.html#typing-support-for-type-hints "Link to this heading")
Added in version 3.5.
**Source code:** [Lib/typing.py](https://github.com/python/cpython/tree/3.13/Lib/typing.py)
Note
The Python runtime does not enforce function and variable type annotations. They can be used by third party tools such as [type checkers](https://docs.python.org/3/glossary.html#term-static-type-checker), IDEs, linters, etc.
This module provides runtime support for type hints.
Consider the function below:
Copy```
defsurface_area_of_cube(edge_length: float) -> str:
  return f"The surface area of the cube is {6*edge_length**2}."

```

The function `surface_area_of_cube` takes an argument expected to be an instance of [`float`](https://docs.python.org/3/library/functions.html#float "float"), as indicated by the [type hint](https://docs.python.org/3/glossary.html#term-type-hint) `edge_length: float`. The function is expected to return an instance of [`str`](https://docs.python.org/3/library/stdtypes.html#str "str"), as indicated by the `-> str` hint.
While type hints can be simple classes like [`float`](https://docs.python.org/3/library/functions.html#float "float") or [`str`](https://docs.python.org/3/library/stdtypes.html#str "str"), they can also be more complex. The [`typing`](https://docs.python.org/3/library/typing.html#module-typing "typing: Support for type hints \(see :pep:`484`\).") module provides a vocabulary of more advanced type hints.
New features are frequently added to the `typing` module. The [typing_extensions](https://pypi.org/project/typing_extensions/) package provides backports of these new features to older versions of Python.
See also 

[Typing cheat sheet](https://mypy.readthedocs.io/en/stable/cheat_sheet_py3.html)
    
A quick overview of type hints (hosted at the mypy docs) 

Type System Reference section of [the mypy docs](https://mypy.readthedocs.io/en/stable/index.html)
    
The Python typing system is standardised via PEPs, so this reference should broadly apply to most Python type checkers. (Some parts may still be specific to mypy.) 

[Static Typing with Python](https://typing.python.org/en/latest/)
    
Type-checker-agnostic documentation written by the community detailing type system features, useful typing related tools and typing best practices.
## Specification for the Python Type System[¶](https://docs.python.org/3/library/typing.html#specification-for-the-python-type-system "Link to this heading")
The canonical, up-to-date specification of the Python type system can be found at [Specification for the Python type system](https://typing.python.org/en/latest/spec/index.html).
## Type aliases[¶](https://docs.python.org/3/library/typing.html#type-aliases "Link to this heading")
A type alias is defined using the [`type`](https://docs.python.org/3/reference/simple_stmts.html#type) statement, which creates an instance of [`TypeAliasType`](https://docs.python.org/3/library/typing.html#typing.TypeAliasType "typing.TypeAliasType"). In this example, `Vector` and `list[float]` will be treated equivalently by static type checkers:
Copy```
type Vector = list[float]
defscale(scalar: float, vector: Vector) -> Vector:
  return [scalar * num for num in vector]
# passes type checking; a list of floats qualifies as a Vector.
new_vector = scale(2.0, [1.0, -4.2, 5.4])

```

Type aliases are useful for simplifying complex type signatures. For example:
Copy```
fromcollections.abcimport Sequence
type ConnectionOptions = dict[str, str]
type Address = tuple[str, int]
type Server = tuple[Address, ConnectionOptions]
defbroadcast_message(message: str, servers: Sequence[Server]) -> None:
  ...
# The static type checker will treat the previous type signature as
# being exactly equivalent to this one.
defbroadcast_message(
  message: str,
  servers: Sequence[tuple[tuple[str, int], dict[str, str]]]
) -> None:
  ...

```

The [`type`](https://docs.python.org/3/reference/simple_stmts.html#type) statement is new in Python 3.12. For backwards compatibility, type aliases can also be created through simple assignment:
Copy```
Vector = list[float]

```

Or marked with [`TypeAlias`](https://docs.python.org/3/library/typing.html#typing.TypeAlias "typing.TypeAlias") to make it explicit that this is a type alias, not a normal variable assignment:
Copy```
fromtypingimport TypeAlias
Vector: TypeAlias = list[float]

```

## NewType[¶](https://docs.python.org/3/library/typing.html#newtype "Link to this heading")
Use the [`NewType`](https://docs.python.org/3/library/typing.html#typing.NewType "typing.NewType") helper to create distinct types:
Copy```
fromtypingimport NewType
UserId = NewType('UserId', int)
some_id = UserId(524313)

```

The static type checker will treat the new type as if it were a subclass of the original type. This is useful in helping catch logical errors:
Copy```
defget_user_name(user_id: UserId) -> str:
  ...
# passes type checking
user_a = get_user_name(UserId(42351))
# fails type checking; an int is not a UserId
user_b = get_user_name(-1)

```

You may still perform all `int` operations on a variable of type `UserId`, but the result will always be of type `int`. This lets you pass in a `UserId` wherever an `int` might be expected, but will prevent you from accidentally creating a `UserId` in an invalid way:
Copy```
# 'output' is of type 'int', not 'UserId'
output = UserId(23413) + UserId(54341)

```

Note that these checks are enforced only by the static type checker. At runtime, the statement `Derived = NewType('Derived', Base)` will make `Derived` a callable that immediately returns whatever parameter you pass it. That means the expression `Derived(some_value)` does not create a new class or introduce much overhead beyond that of a regular function call.
More precisely, the expression `some_value is Derived(some_value)` is always true at runtime.
It is invalid to create a subtype of `Derived`:
Copy```
fromtypingimport NewType
UserId = NewType('UserId', int)
# Fails at runtime and does not pass type checking
classAdminUserId(UserId): pass

```

However, it is possible to create a [`NewType`](https://docs.python.org/3/library/typing.html#typing.NewType "typing.NewType") based on a ‘derived’ `NewType`:
Copy```
fromtypingimport NewType
UserId = NewType('UserId', int)
ProUserId = NewType('ProUserId', UserId)

```

and typechecking for `ProUserId` will work as expected.
See [**PEP 484**](https://peps.python.org/pep-0484/) for more details.
Note
Recall that the use of a type alias declares two types to be _equivalent_ to one another. Doing `type Alias = Original` will make the static type checker treat `Alias` as being _exactly equivalent_ to `Original` in all cases. This is useful when you want to simplify complex type signatures.
In contrast, `NewType` declares one type to be a _subtype_ of another. Doing `Derived = NewType('Derived', Original)` will make the static type checker treat `Derived` as a _subclass_ of `Original`, which means a value of type `Original` cannot be used in places where a value of type `Derived` is expected. This is useful when you want to prevent logic errors with minimal runtime cost.
Added in version 3.5.2.
Changed in version 3.10: `NewType` is now a class rather than a function. As a result, there is some additional runtime cost when calling `NewType` over a regular function.
Changed in version 3.11: The performance of calling `NewType` has been restored to its level in Python 3.9.
## Annotating callable objects[¶](https://docs.python.org/3/library/typing.html#annotating-callable-objects "Link to this heading")
Functions – or other [callable](https://docs.python.org/3/glossary.html#term-callable) objects – can be annotated using [`collections.abc.Callable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Callable "collections.abc.Callable") or deprecated [`typing.Callable`](https://docs.python.org/3/library/typing.html#typing.Callable "typing.Callable"). `Callable[[int], str]` signifies a function that takes a single parameter of type [`int`](https://docs.python.org/3/library/functions.html#int "int") and returns a [`str`](https://docs.python.org/3/library/stdtypes.html#str "str").
For example:
Copy```
fromcollections.abcimport Callable, Awaitable
deffeeder(get_next_item: Callable[[], str]) -> None:
  ... # Body
defasync_query(on_success: Callable[[int], None],
        on_error: Callable[[int, Exception], None]) -> None:
  ... # Body
async defon_update(value: str) -> None:
  ... # Body
callback: Callable[[str], Awaitable[None]] = on_update

```

The subscription syntax must always be used with exactly two values: the argument list and the return type. The argument list must be a list of types, a [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec"), [`Concatenate`](https://docs.python.org/3/library/typing.html#typing.Concatenate "typing.Concatenate"), or an ellipsis. The return type must be a single type.
If a literal ellipsis `...` is given as the argument list, it indicates that a callable with any arbitrary parameter list would be acceptable:
Copy```
defconcat(x: str, y: str) -> str:
  return x + y
x: Callable[..., str]
x = str   # OK
x = concat # Also OK

```

`Callable` cannot express complex signatures such as functions that take a variadic number of arguments, [overloaded functions](https://docs.python.org/3/library/typing.html#overload), or functions that have keyword-only parameters. However, these signatures can be expressed by defining a [`Protocol`](https://docs.python.org/3/library/typing.html#typing.Protocol "typing.Protocol") class with a [`__call__()`](https://docs.python.org/3/reference/datamodel.html#object.__call__ "object.__call__") method:
Copy```
fromcollections.abcimport Iterable
fromtypingimport Protocol
classCombiner(Protocol):
  def__call__(self, *vals: bytes, maxlen: int | None = None) -> list[bytes]: ...
defbatch_proc(data: Iterable[bytes], cb_results: Combiner) -> bytes:
  for item in data:
    ...
defgood_cb(*vals: bytes, maxlen: int | None = None) -> list[bytes]:
  ...
defbad_cb(*vals: bytes, maxitems: int | None) -> list[bytes]:
  ...
batch_proc([], good_cb) # OK
batch_proc([], bad_cb)  # Error! Argument 2 has incompatible type because of
             # different name and kind in the callback

```

Callables which take other callables as arguments may indicate that their parameter types are dependent on each other using [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec"). Additionally, if that callable adds or removes arguments from other callables, the [`Concatenate`](https://docs.python.org/3/library/typing.html#typing.Concatenate "typing.Concatenate") operator may be used. They take the form `Callable[ParamSpecVariable, ReturnType]` and `Callable[Concatenate[Arg1Type, Arg2Type, ..., ParamSpecVariable], ReturnType]` respectively.
Changed in version 3.10: `Callable` now supports [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec") and [`Concatenate`](https://docs.python.org/3/library/typing.html#typing.Concatenate "typing.Concatenate"). See [**PEP 612**](https://peps.python.org/pep-0612/) for more details.
See also
The documentation for [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec") and [`Concatenate`](https://docs.python.org/3/library/typing.html#typing.Concatenate "typing.Concatenate") provides examples of usage in `Callable`.
## Generics[¶](https://docs.python.org/3/library/typing.html#generics "Link to this heading")
Since type information about objects kept in containers cannot be statically inferred in a generic way, many container classes in the standard library support subscription to denote the expected types of container elements.
Copy```
fromcollections.abcimport Mapping, Sequence
classEmployee: ...
# Sequence[Employee] indicates that all elements in the sequence
# must be instances of "Employee".
# Mapping[str, str] indicates that all keys and all values in the mapping
# must be strings.
defnotify_by_email(employees: Sequence[Employee],
          overrides: Mapping[str, str]) -> None: ...

```

Generic functions and classes can be parameterized by using [type parameter syntax](https://docs.python.org/3/reference/compound_stmts.html#type-params):
Copy```
fromcollections.abcimport Sequence
deffirst[T](l: Sequence[T]) -> T: # Function is generic over the TypeVar "T"
  return l[0]

```

Or by using the [`TypeVar`](https://docs.python.org/3/library/typing.html#typing.TypeVar "typing.TypeVar") factory directly:
Copy```
fromcollections.abcimport Sequence
fromtypingimport TypeVar
U = TypeVar('U')         # Declare type variable "U"
defsecond(l: Sequence[U]) -> U: # Function is generic over the TypeVar "U"
  return l[1]

```

Changed in version 3.12: Syntactic support for generics is new in Python 3.12.
## Annotating tuples[¶](https://docs.python.org/3/library/typing.html#annotating-tuples "Link to this heading")
For most containers in Python, the typing system assumes that all elements in the container will be of the same type. For example:
Copy```
fromcollections.abcimport Mapping
# Type checker will infer that all elements in ``x`` are meant to be ints
x: list[int] = []
# Type checker error: ``list`` only accepts a single type argument:
y: list[int, str] = [1, 'foo']
# Type checker will infer that all keys in ``z`` are meant to be strings,
# and that all values in ``z`` are meant to be either strings or ints
z: Mapping[str, str | int] = {}

```

[`list`](https://docs.python.org/3/library/stdtypes.html#list "list") only accepts one type argument, so a type checker would emit an error on the `y` assignment above. Similarly, [`Mapping`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Mapping "collections.abc.Mapping") only accepts two type arguments: the first indicates the type of the keys, and the second indicates the type of the values.
Unlike most other Python containers, however, it is common in idiomatic Python code for tuples to have elements which are not all of the same type. For this reason, tuples are special-cased in Python’s typing system. [`tuple`](https://docs.python.org/3/library/stdtypes.html#tuple "tuple") accepts _any number_ of type arguments:
Copy```
# OK: ``x`` is assigned to a tuple of length 1 where the sole element is an int
x: tuple[int] = (5,)
# OK: ``y`` is assigned to a tuple of length 2;
# element 1 is an int, element 2 is a str
y: tuple[int, str] = (5, "foo")
# Error: the type annotation indicates a tuple of length 1,
# but ``z`` has been assigned to a tuple of length 3
z: tuple[int] = (1, 2, 3)

```

To denote a tuple which could be of _any_ length, and in which all elements are of the same type `T`, use `tuple[T, ...]`. To denote an empty tuple, use `tuple[()]`. Using plain `tuple` as an annotation is equivalent to using `tuple[Any, ...]`:
Copy```
x: tuple[int, ...] = (1, 2)
# These reassignments are OK: ``tuple[int, ...]`` indicates x can be of any length
x = (1, 2, 3)
x = ()
# This reassignment is an error: all elements in ``x`` must be ints
x = ("foo", "bar")
# ``y`` can only ever be assigned to an empty tuple
y: tuple[()] = ()
z: tuple = ("foo", "bar")
# These reassignments are OK: plain ``tuple`` is equivalent to ``tuple[Any, ...]``
z = (1, 2, 3)
z = ()

```

## The type of class objects[¶](https://docs.python.org/3/library/typing.html#the-type-of-class-objects "Link to this heading")
A variable annotated with `C` may accept a value of type `C`. In contrast, a variable annotated with `type[C]` (or deprecated [`typing.Type[C]`](https://docs.python.org/3/library/typing.html#typing.Type "typing.Type")) may accept values that are classes themselves – specifically, it will accept the _class object_ of `C`. For example:
Copy```
a = 3     # Has type ``int``
b = int    # Has type ``type[int]``
c = type(a)  # Also has type ``type[int]``

```

Note that `type[C]` is covariant:
Copy```
classUser: ...
classProUser(User): ...
classTeamUser(User): ...
defmake_new_user(user_class: type[User]) -> User:
  # ...
  return user_class()
make_new_user(User)   # OK
make_new_user(ProUser)  # Also OK: ``type[ProUser]`` is a subtype of ``type[User]``
make_new_user(TeamUser) # Still fine
make_new_user(User())  # Error: expected ``type[User]`` but got ``User``
make_new_user(int)    # Error: ``type[int]`` is not a subtype of ``type[User]``

```

The only legal parameters for [`type`](https://docs.python.org/3/library/functions.html#type "type") are classes, [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any"), [type variables](https://docs.python.org/3/library/typing.html#generics), and unions of any of these types. For example:
Copy```
defnew_non_team_user(user_class: type[BasicUser | ProUser]): ...
new_non_team_user(BasicUser) # OK
new_non_team_user(ProUser)  # OK
new_non_team_user(TeamUser)  # Error: ``type[TeamUser]`` is not a subtype
               # of ``type[BasicUser | ProUser]``
new_non_team_user(User)    # Also an error

```

`type[Any]` is equivalent to [`type`](https://docs.python.org/3/library/functions.html#type "type"), which is the root of Python’s [metaclass hierarchy](https://docs.python.org/3/reference/datamodel.html#metaclasses).
## Annotating generators and coroutines[¶](https://docs.python.org/3/library/typing.html#annotating-generators-and-coroutines "Link to this heading")
A generator can be annotated using the generic type [`Generator[YieldType, SendType, ReturnType]`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Generator "collections.abc.Generator"). For example:
Copy```
defecho_round() -> Generator[int, float, str]:
  sent = yield 0
  while sent >= 0:
    sent = yield round(sent)
  return 'Done'

```

Note that unlike many other generic classes in the standard library, the `SendType` of [`Generator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Generator "collections.abc.Generator") behaves contravariantly, not covariantly or invariantly.
The `SendType` and `ReturnType` parameters default to `None`:
Copy```
definfinite_stream(start: int) -> Generator[int]:
  while True:
    yield start
    start += 1

```

It is also possible to set these types explicitly:
Copy```
definfinite_stream(start: int) -> Generator[int, None, None]:
  while True:
    yield start
    start += 1

```

Simple generators that only ever yield values can also be annotated as having a return type of either [`Iterable[YieldType]`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Iterable "collections.abc.Iterable") or [`Iterator[YieldType]`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Iterator "collections.abc.Iterator"):
Copy```
definfinite_stream(start: int) -> Iterator[int]:
  while True:
    yield start
    start += 1

```

Async generators are handled in a similar fashion, but don’t expect a `ReturnType` type argument ([`AsyncGenerator[YieldType, SendType]`](https://docs.python.org/3/library/collections.abc.html#collections.abc.AsyncGenerator "collections.abc.AsyncGenerator")). The `SendType` argument defaults to `None`, so the following definitions are equivalent:
Copy```
async definfinite_stream(start: int) -> AsyncGenerator[int]:
  while True:
    yield start
    start = await increment(start)
async definfinite_stream(start: int) -> AsyncGenerator[int, None]:
  while True:
    yield start
    start = await increment(start)

```

As in the synchronous case, [`AsyncIterable[YieldType]`](https://docs.python.org/3/library/collections.abc.html#collections.abc.AsyncIterable "collections.abc.AsyncIterable") and [`AsyncIterator[YieldType]`](https://docs.python.org/3/library/collections.abc.html#collections.abc.AsyncIterator "collections.abc.AsyncIterator") are available as well:
Copy```
async definfinite_stream(start: int) -> AsyncIterator[int]:
  while True:
    yield start
    start = await increment(start)

```

Coroutines can be annotated using [`Coroutine[YieldType, SendType, ReturnType]`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Coroutine "collections.abc.Coroutine"). Generic arguments correspond to those of [`Generator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Generator "collections.abc.Generator"), for example:
Copy```
fromcollections.abcimport Coroutine
c: Coroutine[list[str], str, int] # Some coroutine defined elsewhere
x = c.send('hi')          # Inferred type of 'x' is list[str]
async defbar() -> None:
  y = await c          # Inferred type of 'y' is int

```

## User-defined generic types[¶](https://docs.python.org/3/library/typing.html#user-defined-generic-types "Link to this heading")
A user-defined class can be defined as a generic class.
Copy```
fromloggingimport Logger
classLoggedVar[T]:
  def__init__(self, value: T, name: str, logger: Logger) -> None:
    self.name = name
    self.logger = logger
    self.value = value
  defset(self, new: T) -> None:
    self.log('Set ' + repr(self.value))
    self.value = new
  defget(self) -> T:
    self.log('Get ' + repr(self.value))
    return self.value
  deflog(self, message: str) -> None:
    self.logger.info('%s: %s', self.name, message)

```

This syntax indicates that the class `LoggedVar` is parameterised around a single [type variable](https://docs.python.org/3/library/typing.html#typevar) `T` . This also makes `T` valid as a type within the class body.
Generic classes implicitly inherit from [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic"). For compatibility with Python 3.11 and lower, it is also possible to inherit explicitly from [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic") to indicate a generic class:
Copy```
fromtypingimport TypeVar, Generic
T = TypeVar('T')
classLoggedVar(Generic[T]):
  ...

```

Generic classes have [`__class_getitem__()`](https://docs.python.org/3/reference/datamodel.html#object.__class_getitem__ "object.__class_getitem__") methods, meaning they can be parameterised at runtime (e.g. `LoggedVar[int]` below):
Copy```
fromcollections.abcimport Iterable
defzero_all_vars(vars: Iterable[LoggedVar[int]]) -> None:
  for var in vars:
    var.set(0)

```

A generic type can have any number of type variables. All varieties of [`TypeVar`](https://docs.python.org/3/library/typing.html#typing.TypeVar "typing.TypeVar") are permissible as parameters for a generic type:
Copy```
fromtypingimport TypeVar, Generic, Sequence
classWeirdTrio[T, B: Sequence[bytes], S: (int, str)]:
  ...
OldT = TypeVar('OldT', contravariant=True)
OldB = TypeVar('OldB', bound=Sequence[bytes], covariant=True)
OldS = TypeVar('OldS', int, str)
classOldWeirdTrio(Generic[OldT, OldB, OldS]):
  ...

```

Each type variable argument to [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic") must be distinct. This is thus invalid:
Copy```
fromtypingimport TypeVar, Generic
...
classPair[M, M]: # SyntaxError
  ...
T = TypeVar('T')
classPair(Generic[T, T]):  # INVALID
  ...

```

Generic classes can also inherit from other classes:
Copy```
fromcollections.abcimport Sized
classLinkedList[T](Sized):
  ...

```

When inheriting from generic classes, some type parameters could be fixed:
Copy```
fromcollections.abcimport Mapping
classMyDict[T](Mapping[str, T]):
  ...

```

In this case `MyDict` has a single parameter, `T`.
Using a generic class without specifying type parameters assumes [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") for each position. In the following example, `MyIterable` is not generic but implicitly inherits from `Iterable[Any]`:
Copy```
fromcollections.abcimport Iterable
classMyIterable(Iterable): # Same as Iterable[Any]
  ...

```

User-defined generic type aliases are also supported. Examples:
Copy```
fromcollections.abcimport Iterable
type Response[S] = Iterable[S] | int
# Return type here is same as Iterable[str] | int
defresponse(query: str) -> Response[str]:
  ...
type Vec[T] = Iterable[tuple[T, T]]
definproduct[T: (int, float, complex)](v: Vec[T]) -> T: # Same as Iterable[tuple[T, T]]
  return sum(x*y for x, y in v)

```

For backward compatibility, generic type aliases can also be created through a simple assignment:
Copy```
fromcollections.abcimport Iterable
fromtypingimport TypeVar
S = TypeVar("S")
Response = Iterable[S] | int

```

Changed in version 3.7: [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic") no longer has a custom metaclass.
Changed in version 3.12: Syntactic support for generics and type aliases is new in version 3.12. Previously, generic classes had to explicitly inherit from [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic") or contain a type variable in one of their bases.
User-defined generics for parameter expressions are also supported via parameter specification variables in the form `[**P]`. The behavior is consistent with type variables’ described above as parameter specification variables are treated by the `typing` module as a specialized type variable. The one exception to this is that a list of types can be used to substitute a [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec"):
Copy```
>>> classZ[T, **P]: ... # T is a TypeVar; P is a ParamSpec
...
>>> Z[int, [dict, float]]
__main__.Z[int, [dict, float]]

```

Classes generic over a [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec") can also be created using explicit inheritance from [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic"). In this case, `**` is not used:
Copy```
fromtypingimport ParamSpec, Generic
P = ParamSpec('P')
classZ(Generic[P]):
  ...

```

Another difference between [`TypeVar`](https://docs.python.org/3/library/typing.html#typing.TypeVar "typing.TypeVar") and [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec") is that a generic with only one parameter specification variable will accept parameter lists in the forms `X[[Type1, Type2, ...]]` and also `X[Type1, Type2, ...]` for aesthetic reasons. Internally, the latter is converted to the former, so the following are equivalent:
Copy```
>>> classX[**P]: ...
...
>>> X[int, str]
__main__.X[[int, str]]
>>> X[[int, str]]
__main__.X[[int, str]]

```

Note that generics with [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec") may not have correct `__parameters__` after substitution in some cases because they are intended primarily for static type checking.
Changed in version 3.10: [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic") can now be parameterized over parameter expressions. See [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec") and [**PEP 612**](https://peps.python.org/pep-0612/) for more details.
A user-defined generic class can have ABCs as base classes without a metaclass conflict. Generic metaclasses are not supported. The outcome of parameterizing generics is cached, and most types in the `typing` module are [hashable](https://docs.python.org/3/glossary.html#term-hashable) and comparable for equality.
## The [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") type[¶](https://docs.python.org/3/library/typing.html#the-any-type "Link to this heading")
A special kind of type is [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any"). A static type checker will treat every type as being compatible with [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") and [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") as being compatible with every type.
This means that it is possible to perform any operation or method call on a value of type [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") and assign it to any variable:
Copy```
fromtypingimport Any
a: Any = None
a = []     # OK
a = 2      # OK
s: str = ''
s = a      # OK
deffoo(item: Any) -> int:
  # Passes type checking; 'item' could be any type,
  # and that type might have a 'bar' method
  item.bar()
  ...

```

Notice that no type checking is performed when assigning a value of type [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") to a more precise type. For example, the static type checker did not report an error when assigning `a` to `s` even though `s` was declared to be of type [`str`](https://docs.python.org/3/library/stdtypes.html#str "str") and receives an [`int`](https://docs.python.org/3/library/functions.html#int "int") value at runtime!
Furthermore, all functions without a return type or parameter types will implicitly default to using [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any"):
Copy```
deflegacy_parser(text):
  ...
  return data
# A static type checker will treat the above
# as having the same signature as:
deflegacy_parser(text: Any) -> Any:
  ...
  return data

```

This behavior allows [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") to be used as an _escape hatch_ when you need to mix dynamically and statically typed code.
Contrast the behavior of [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") with the behavior of [`object`](https://docs.python.org/3/library/functions.html#object "object"). Similar to [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any"), every type is a subtype of [`object`](https://docs.python.org/3/library/functions.html#object "object"). However, unlike [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any"), the reverse is not true: [`object`](https://docs.python.org/3/library/functions.html#object "object") is _not_ a subtype of every other type.
That means when the type of a value is [`object`](https://docs.python.org/3/library/functions.html#object "object"), a type checker will reject almost all operations on it, and assigning it to a variable (or using it as a return value) of a more specialized type is a type error. For example:
Copy```
defhash_a(item: object) -> int:
  # Fails type checking; an object does not have a 'magic' method.
  item.magic()
  ...
defhash_b(item: Any) -> int:
  # Passes type checking
  item.magic()
  ...
# Passes type checking, since ints and strs are subclasses of object
hash_a(42)
hash_a("foo")
# Passes type checking, since Any is compatible with all types
hash_b(42)
hash_b("foo")

```

Use [`object`](https://docs.python.org/3/library/functions.html#object "object") to indicate that a value could be any type in a typesafe manner. Use [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") to indicate that a value is dynamically typed.
## Nominal vs structural subtyping[¶](https://docs.python.org/3/library/typing.html#nominal-vs-structural-subtyping "Link to this heading")
Initially [**PEP 484**](https://peps.python.org/pep-0484/) defined the Python static type system as using _nominal subtyping_. This means that a class `A` is allowed where a class `B` is expected if and only if `A` is a subclass of `B`.
This requirement previously also applied to abstract base classes, such as [`Iterable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Iterable "collections.abc.Iterable"). The problem with this approach is that a class had to be explicitly marked to support them, which is unpythonic and unlike what one would normally do in idiomatic dynamically typed Python code. For example, this conforms to [**PEP 484**](https://peps.python.org/pep-0484/):
Copy```
fromcollections.abcimport Sized, Iterable, Iterator
classBucket(Sized, Iterable[int]):
  ...
  def__len__(self) -> int: ...
  def__iter__(self) -> Iterator[int]: ...

```

[**PEP 544**](https://peps.python.org/pep-0544/) allows to solve this problem by allowing users to write the above code without explicit base classes in the class definition, allowing `Bucket` to be implicitly considered a subtype of both `Sized` and `Iterable[int]` by static type checkers. This is known as _structural subtyping_ (or static duck-typing):
Copy```
fromcollections.abcimport Iterator, Iterable
classBucket: # Note: no base classes
  ...
  def__len__(self) -> int: ...
  def__iter__(self) -> Iterator[int]: ...
defcollect(items: Iterable[int]) -> int: ...
result = collect(Bucket()) # Passes type check

```

Moreover, by subclassing a special class [`Protocol`](https://docs.python.org/3/library/typing.html#typing.Protocol "typing.Protocol"), a user can define new custom protocols to fully enjoy structural subtyping (see examples below).
## Module contents[¶](https://docs.python.org/3/library/typing.html#module-contents "Link to this heading")
The `typing` module defines the following classes, functions and decorators.
### Special typing primitives[¶](https://docs.python.org/3/library/typing.html#special-typing-primitives "Link to this heading")
#### Special types[¶](https://docs.python.org/3/library/typing.html#special-types "Link to this heading")
These can be used as types in annotations. They do not support subscription using `[]`. 

typing.Any[¶](https://docs.python.org/3/library/typing.html#typing.Any "Link to this definition")
    
Special type indicating an unconstrained type.
  * Every type is compatible with [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any").
  * [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") is compatible with every type.


Changed in version 3.11: [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") can now be used as a base class. This can be useful for avoiding type checker errors with classes that can duck type anywhere or are highly dynamic. 

typing.AnyStr[¶](https://docs.python.org/3/library/typing.html#typing.AnyStr "Link to this definition")
    
A [constrained type variable](https://docs.python.org/3/library/typing.html#typing-constrained-typevar).
Definition:
Copy```
AnyStr = TypeVar('AnyStr', str, bytes)

```

`AnyStr` is meant to be used for functions that may accept [`str`](https://docs.python.org/3/library/stdtypes.html#str "str") or [`bytes`](https://docs.python.org/3/library/stdtypes.html#bytes "bytes") arguments but cannot allow the two to mix.
For example:
Copy```
defconcat(a: AnyStr, b: AnyStr) -> AnyStr:
  return a + b
concat("foo", "bar")  # OK, output has type 'str'
concat(b"foo", b"bar") # OK, output has type 'bytes'
concat("foo", b"bar")  # Error, cannot mix str and bytes

```

Note that, despite its name, `AnyStr` has nothing to do with the [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any") type, nor does it mean “any string”. In particular, `AnyStr` and `str | bytes` are different from each other and have different use cases:
Copy```
# Invalid use of AnyStr:
# The type variable is used only once in the function signature,
# so cannot be "solved" by the type checker
defgreet_bad(cond: bool) -> AnyStr:
  return "hi there!" if cond else b"greetings!"
# The better way of annotating this function:
defgreet_proper(cond: bool) -> str | bytes:
  return "hi there!" if cond else b"greetings!"

```

Deprecated since version 3.13, will be removed in version 3.18: Deprecated in favor of the new [type parameter syntax](https://docs.python.org/3/reference/compound_stmts.html#type-params). Use `class A[T: (str, bytes)]: ...` instead of importing `AnyStr`. See [**PEP 695**](https://peps.python.org/pep-0695/) for more details.
In Python 3.16, `AnyStr` will be removed from `typing.__all__`, and deprecation warnings will be emitted at runtime when it is accessed or imported from `typing`. `AnyStr` will be removed from `typing` in Python 3.18. 

typing.LiteralString[¶](https://docs.python.org/3/library/typing.html#typing.LiteralString "Link to this definition")
    
Special type that includes only literal strings.
Any string literal is compatible with `LiteralString`, as is another `LiteralString`. However, an object typed as just `str` is not. A string created by composing `LiteralString`-typed objects is also acceptable as a `LiteralString`.
Example:
Copy```
defrun_query(sql: LiteralString) -> None:
  ...
defcaller(arbitrary_string: str, literal_string: LiteralString) -> None:
  run_query("SELECT * FROM students") # OK
  run_query(literal_string) # OK
  run_query("SELECT * FROM " + literal_string) # OK
  run_query(arbitrary_string) # type checker error
  run_query( # type checker error
    f"SELECT * FROM students WHERE name = {arbitrary_string}"
  )

```

`LiteralString` is useful for sensitive APIs where arbitrary user-generated strings could generate problems. For example, the two cases above that generate type checker errors could be vulnerable to an SQL injection attack.
See [**PEP 675**](https://peps.python.org/pep-0675/) for more details.
Added in version 3.11. 

typing.Never[¶](https://docs.python.org/3/library/typing.html#typing.Never "Link to this definition")


typing.NoReturn[¶](https://docs.python.org/3/library/typing.html#typing.NoReturn "Link to this definition")
    
`Never` and `NoReturn` represent the [bottom type](https://en.wikipedia.org/wiki/Bottom_type), a type that has no members.
They can be used to indicate that a function never returns, such as [`sys.exit()`](https://docs.python.org/3/library/sys.html#sys.exit "sys.exit"):
Copy```
fromtypingimport Never # or NoReturn
defstop() -> Never:
  raise RuntimeError('no way')

```

Or to define a function that should never be called, as there are no valid arguments, such as [`assert_never()`](https://docs.python.org/3/library/typing.html#typing.assert_never "typing.assert_never"):
Copy```
fromtypingimport Never # or NoReturn
defnever_call_me(arg: Never) -> None:
  pass
defint_or_str(arg: int | str) -> None:
  never_call_me(arg) # type checker error
  match arg:
    case int():
      print("It's an int")
    case str():
      print("It's a str")
    case_:
      never_call_me(arg) # OK, arg is of type Never (or NoReturn)

```

`Never` and `NoReturn` have the same meaning in the type system and static type checkers treat both equivalently.
Added in version 3.6.2: Added [`NoReturn`](https://docs.python.org/3/library/typing.html#typing.NoReturn "typing.NoReturn").
Added in version 3.11: Added [`Never`](https://docs.python.org/3/library/typing.html#typing.Never "typing.Never"). 

typing.Self[¶](https://docs.python.org/3/library/typing.html#typing.Self "Link to this definition")
    
Special type to represent the current enclosed class.
For example:
Copy```
fromtypingimport Self, reveal_type
classFoo:
  defreturn_self(self) -> Self:
    ...
    return self
classSubclassOfFoo(Foo): pass
reveal_type(Foo().return_self()) # Revealed type is "Foo"
reveal_type(SubclassOfFoo().return_self()) # Revealed type is "SubclassOfFoo"

```

This annotation is semantically equivalent to the following, albeit in a more succinct fashion:
Copy```
fromtypingimport TypeVar
Self = TypeVar("Self", bound="Foo")
classFoo:
  defreturn_self(self: Self) -> Self:
    ...
    return self

```

In general, if something returns `self`, as in the above examples, you should use `Self` as the return annotation. If `Foo.return_self` was annotated as returning `"Foo"`, then the type checker would infer the object returned from `SubclassOfFoo.return_self` as being of type `Foo` rather than `SubclassOfFoo`.
Other common use cases include:
  * [`classmethod`](https://docs.python.org/3/library/functions.html#classmethod "classmethod")s that are used as alternative constructors and return instances of the `cls` parameter.
  * Annotating an [`__enter__()`](https://docs.python.org/3/reference/datamodel.html#object.__enter__ "object.__enter__") method which returns self.


You should not use `Self` as the return annotation if the method is not guaranteed to return an instance of a subclass when the class is subclassed:
Copy```
classEggs:
  # Self would be an incorrect return annotation here,
  # as the object returned is always an instance of Eggs,
  # even in subclasses
  defreturns_eggs(self) -> "Eggs":
    return Eggs()

```

See [**PEP 673**](https://peps.python.org/pep-0673/) for more details.
Added in version 3.11. 

typing.TypeAlias[¶](https://docs.python.org/3/library/typing.html#typing.TypeAlias "Link to this definition")
    
Special annotation for explicitly declaring a [type alias](https://docs.python.org/3/library/typing.html#type-aliases).
For example:
Copy```
fromtypingimport TypeAlias
Factors: TypeAlias = list[int]

```

`TypeAlias` is particularly useful on older Python versions for annotating aliases that make use of forward references, as it can be hard for type checkers to distinguish these from normal variable assignments:
Copy```
fromtypingimport Generic, TypeAlias, TypeVar
T = TypeVar("T")
# "Box" does not exist yet,
# so we have to use quotes for the forward reference on Python <3.12.
# Using ``TypeAlias`` tells the type checker that this is a type alias declaration,
# not a variable assignment to a string.
BoxOfStrings: TypeAlias = "Box[str]"
classBox(Generic[T]):
  @classmethod
  defmake_box_of_strings(cls) -> BoxOfStrings: ...

```

See [**PEP 613**](https://peps.python.org/pep-0613/) for more details.
Added in version 3.10.
Deprecated since version 3.12: [`TypeAlias`](https://docs.python.org/3/library/typing.html#typing.TypeAlias "typing.TypeAlias") is deprecated in favor of the [`type`](https://docs.python.org/3/reference/simple_stmts.html#type) statement, which creates instances of [`TypeAliasType`](https://docs.python.org/3/library/typing.html#typing.TypeAliasType "typing.TypeAliasType") and which natively supports forward references. Note that while [`TypeAlias`](https://docs.python.org/3/library/typing.html#typing.TypeAlias "typing.TypeAlias") and [`TypeAliasType`](https://docs.python.org/3/library/typing.html#typing.TypeAliasType "typing.TypeAliasType") serve similar purposes and have similar names, they are distinct and the latter is not the type of the former. Removal of [`TypeAlias`](https://docs.python.org/3/library/typing.html#typing.TypeAlias "typing.TypeAlias") is not currently planned, but users are encouraged to migrate to [`type`](https://docs.python.org/3/reference/simple_stmts.html#type) statements.
#### Special forms[¶](https://docs.python.org/3/library/typing.html#special-forms "Link to this heading")
These can be used as types in annotations. They all support subscription using `[]`, but each has a unique syntax. 

typing.Union[¶](https://docs.python.org/3/library/typing.html#typing.Union "Link to this definition")
    
Union type; `Union[X, Y]` is equivalent to `X | Y` and means either X or Y.
To define a union, use e.g. `Union[int, str]` or the shorthand `int | str`. Using that shorthand is recommended. Details:
  * The arguments must be types and there must be at least one.
  * Unions of unions are flattened, e.g.:
Copy```
Union[Union[int, str], float] == Union[int, str, float]

```

However, this does not apply to unions referenced through a type alias, to avoid forcing evaluation of the underlying [`TypeAliasType`](https://docs.python.org/3/library/typing.html#typing.TypeAliasType "typing.TypeAliasType"):
Copy```
type A = Union[int, str]
Union[A, float] != Union[int, str, float]

```

  * Unions of a single argument vanish, e.g.:
Copy```
Union[int] == int # The constructor actually returns int

```

  * Redundant arguments are skipped, e.g.:
Copy```
Union[int, str, int] == Union[int, str] == int | str

```

  * When comparing unions, the argument order is ignored, e.g.:
Copy```
Union[int, str] == Union[str, int]

```

  * You cannot subclass or instantiate a `Union`.
  * You cannot write `Union[X][Y]`.


Changed in version 3.7: Don’t remove explicit subclasses from unions at runtime.
Changed in version 3.10: Unions can now be written as `X | Y`. See [union type expressions](https://docs.python.org/3/library/stdtypes.html#types-union). 

typing.Optional[¶](https://docs.python.org/3/library/typing.html#typing.Optional "Link to this definition")
    
`Optional[X]` is equivalent to `X | None` (or `Union[X, None]`).
Note that this is not the same concept as an optional argument, which is one that has a default. An optional argument with a default does not require the `Optional` qualifier on its type annotation just because it is optional. For example:
Copy```
deffoo(arg: int = 0) -> None:
  ...

```

On the other hand, if an explicit value of `None` is allowed, the use of `Optional` is appropriate, whether the argument is optional or not. For example:
Copy```
deffoo(arg: Optional[int] = None) -> None:
  ...

```

Changed in version 3.10: Optional can now be written as `X | None`. See [union type expressions](https://docs.python.org/3/library/stdtypes.html#types-union). 

typing.Concatenate[¶](https://docs.python.org/3/library/typing.html#typing.Concatenate "Link to this definition")
    
Special form for annotating higher-order functions.
`Concatenate` can be used in conjunction with [Callable](https://docs.python.org/3/library/typing.html#annotating-callables) and [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec") to annotate a higher-order callable which adds, removes, or transforms parameters of another callable. Usage is in the form `Concatenate[Arg1Type, Arg2Type, ..., ParamSpecVariable]`. `Concatenate` is currently only valid when used as the first argument to a [Callable](https://docs.python.org/3/library/typing.html#annotating-callables). The last parameter to `Concatenate` must be a [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec") or ellipsis (`...`).
For example, to annotate a decorator `with_lock` which provides a [`threading.Lock`](https://docs.python.org/3/library/threading.html#threading.Lock "threading.Lock") to the decorated function, `Concatenate` can be used to indicate that `with_lock` expects a callable which takes in a `Lock` as the first argument, and returns a callable with a different type signature. In this case, the [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec") indicates that the returned callable’s parameter types are dependent on the parameter types of the callable being passed in:
Copy```
fromcollections.abcimport Callable
fromthreadingimport Lock
fromtypingimport Concatenate
# Use this lock to ensure that only one thread is executing a function
# at any time.
my_lock = Lock()
defwith_lock[**P, R](f: Callable[Concatenate[Lock, P], R]) -> Callable[P, R]:
'''A type-safe decorator which provides a lock.'''
  definner(*args: P.args, **kwargs: P.kwargs) -> R:
    # Provide the lock as the first argument.
    return f(my_lock, *args, **kwargs)
  return inner
@with_lock
defsum_threadsafe(lock: Lock, numbers: list[float]) -> float:
'''Add a list of numbers together in a thread-safe manner.'''
  with lock:
    return sum(numbers)
# We don't need to pass in the lock ourselves thanks to the decorator.
sum_threadsafe([1.1, 2.2, 3.3])

```

Added in version 3.10.
See also
  * [**PEP 612**](https://peps.python.org/pep-0612/) – Parameter Specification Variables (the PEP which introduced `ParamSpec` and `Concatenate`)
  * [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec")
  * [Annotating callable objects](https://docs.python.org/3/library/typing.html#annotating-callables)



typing.Literal[¶](https://docs.python.org/3/library/typing.html#typing.Literal "Link to this definition")
    
Special typing form to define “literal types”.
`Literal` can be used to indicate to type checkers that the annotated object has a value equivalent to one of the provided literals.
For example:
Copy```
defvalidate_simple(data: Any) -> Literal[True]: # always returns True
  ...
type Mode = Literal['r', 'rb', 'w', 'wb']
defopen_helper(file: str, mode: Mode) -> str:
  ...
open_helper('/some/path', 'r')   # Passes type check
open_helper('/other/path', 'typo') # Error in type checker

```

`Literal[...]` cannot be subclassed. At runtime, an arbitrary value is allowed as type argument to `Literal[...]`, but type checkers may impose restrictions. See [**PEP 586**](https://peps.python.org/pep-0586/) for more details about literal types.
Additional details:
  * The arguments must be literal values and there must be at least one.
  * Nested `Literal` types are flattened, e.g.:
Copy```
assert Literal[Literal[1, 2], 3] == Literal[1, 2, 3]

```

However, this does not apply to `Literal` types referenced through a type alias, to avoid forcing evaluation of the underlying [`TypeAliasType`](https://docs.python.org/3/library/typing.html#typing.TypeAliasType "typing.TypeAliasType"):
Copy```
type A = Literal[1, 2]
assert Literal[A, 3] != Literal[1, 2, 3]

```

  * Redundant arguments are skipped, e.g.:
Copy```
assert Literal[1, 2, 1] == Literal[1, 2]

```

  * When comparing literals, the argument order is ignored, e.g.:
Copy```
assert Literal[1, 2] == Literal[2, 1]

```

  * You cannot subclass or instantiate a `Literal`.
  * You cannot write `Literal[X][Y]`.


Added in version 3.8.
Changed in version 3.9.1: `Literal` now de-duplicates parameters. Equality comparisons of `Literal` objects are no longer order dependent. `Literal` objects will now raise a [`TypeError`](https://docs.python.org/3/library/exceptions.html#TypeError "TypeError") exception during equality comparisons if one of their parameters are not [hashable](https://docs.python.org/3/glossary.html#term-hashable). 

typing.ClassVar[¶](https://docs.python.org/3/library/typing.html#typing.ClassVar "Link to this definition")
    
Special type construct to mark class variables.
As introduced in [**PEP 526**](https://peps.python.org/pep-0526/), a variable annotation wrapped in ClassVar indicates that a given attribute is intended to be used as a class variable and should not be set on instances of that class. Usage:
Copy```
classStarship:
  stats: ClassVar[dict[str, int]] = {} # class variable
  damage: int = 10           # instance variable

```

[`ClassVar`](https://docs.python.org/3/library/typing.html#typing.ClassVar "typing.ClassVar") accepts only types and cannot be further subscribed.
[`ClassVar`](https://docs.python.org/3/library/typing.html#typing.ClassVar "typing.ClassVar") is not a class itself, and should not be used with [`isinstance()`](https://docs.python.org/3/library/functions.html#isinstance "isinstance") or [`issubclass()`](https://docs.python.org/3/library/functions.html#issubclass "issubclass"). [`ClassVar`](https://docs.python.org/3/library/typing.html#typing.ClassVar "typing.ClassVar") does not change Python runtime behavior, but it can be used by third-party type checkers. For example, a type checker might flag the following code as an error:
Copy```
enterprise_d = Starship(3000)
enterprise_d.stats = {} # Error, setting class variable on instance
Starship.stats = {}   # This is OK

```

Added in version 3.5.3.
Changed in version 3.13: [`ClassVar`](https://docs.python.org/3/library/typing.html#typing.ClassVar "typing.ClassVar") can now be nested in [`Final`](https://docs.python.org/3/library/typing.html#typing.Final "typing.Final") and vice versa. 

typing.Final[¶](https://docs.python.org/3/library/typing.html#typing.Final "Link to this definition")
    
Special typing construct to indicate final names to type checkers.
Final names cannot be reassigned in any scope. Final names declared in class scopes cannot be overridden in subclasses.
For example:
Copy```
MAX_SIZE: Final = 9000
MAX_SIZE += 1 # Error reported by type checker
classConnection:
  TIMEOUT: Final[int] = 10
classFastConnector(Connection):
  TIMEOUT = 1 # Error reported by type checker

```

There is no runtime checking of these properties. See [**PEP 591**](https://peps.python.org/pep-0591/) for more details.
Added in version 3.8.
Changed in version 3.13: [`Final`](https://docs.python.org/3/library/typing.html#typing.Final "typing.Final") can now be nested in [`ClassVar`](https://docs.python.org/3/library/typing.html#typing.ClassVar "typing.ClassVar") and vice versa. 

typing.Required[¶](https://docs.python.org/3/library/typing.html#typing.Required "Link to this definition")
    
Special typing construct to mark a [`TypedDict`](https://docs.python.org/3/library/typing.html#typing.TypedDict "typing.TypedDict") key as required.
This is mainly useful for `total=False` TypedDicts. See [`TypedDict`](https://docs.python.org/3/library/typing.html#typing.TypedDict "typing.TypedDict") and [**PEP 655**](https://peps.python.org/pep-0655/) for more details.
Added in version 3.11. 

typing.NotRequired[¶](https://docs.python.org/3/library/typing.html#typing.NotRequired "Link to this definition")
    
Special typing construct to mark a [`TypedDict`](https://docs.python.org/3/library/typing.html#typing.TypedDict "typing.TypedDict") key as potentially missing.
See [`TypedDict`](https://docs.python.org/3/library/typing.html#typing.TypedDict "typing.TypedDict") and [**PEP 655**](https://peps.python.org/pep-0655/) for more details.
Added in version 3.11. 

typing.ReadOnly[¶](https://docs.python.org/3/library/typing.html#typing.ReadOnly "Link to this definition")
    
A special typing construct to mark an item of a [`TypedDict`](https://docs.python.org/3/library/typing.html#typing.TypedDict "typing.TypedDict") as read-only.
For example:
Copy```
classMovie(TypedDict):
  title: ReadOnly[str]
  year: int
defmutate_movie(m: Movie) -> None:
  m["year"] = 1999 # allowed
  m["title"] = "The Matrix" # typechecker error

```

There is no runtime checking for this property.
See [`TypedDict`](https://docs.python.org/3/library/typing.html#typing.TypedDict "typing.TypedDict") and [**PEP 705**](https://peps.python.org/pep-0705/) for more details.
Added in version 3.13. 

typing.Annotated[¶](https://docs.python.org/3/library/typing.html#typing.Annotated "Link to this definition")
    
Special typing form to add context-specific metadata to an annotation.
Add metadata `x` to a given type `T` by using the annotation `Annotated[T, x]`. Metadata added using `Annotated` can be used by static analysis tools or at runtime. At runtime, the metadata is stored in a `__metadata__` attribute.
If a library or tool encounters an annotation `Annotated[T, x]` and has no special logic for the metadata, it should ignore the metadata and simply treat the annotation as `T`. As such, `Annotated` can be useful for code that wants to use annotations for purposes outside Python’s static typing system.
Using `Annotated[T, x]` as an annotation still allows for static typechecking of `T`, as type checkers will simply ignore the metadata `x`. In this way, `Annotated` differs from the [`@no_type_check`](https://docs.python.org/3/library/typing.html#typing.no_type_check "typing.no_type_check") decorator, which can also be used for adding annotations outside the scope of the typing system, but completely disables typechecking for a function or class.
The responsibility of how to interpret the metadata lies with the tool or library encountering an `Annotated` annotation. A tool or library encountering an `Annotated` type can scan through the metadata elements to determine if they are of interest (e.g., using [`isinstance()`](https://docs.python.org/3/library/functions.html#isinstance "isinstance")). 

Annotated[<type>, <metadata>]

Here is an example of how you might use `Annotated` to add metadata to type annotations if you were doing range analysis:
Copy```
@dataclass
classValueRange:
  lo: int
  hi: int
T1 = Annotated[int, ValueRange(-10, 5)]
T2 = Annotated[T1, ValueRange(-20, 3)]

```

The first argument to `Annotated` must be a valid type. Multiple metadata elements can be supplied as `Annotated` supports variadic arguments. The order of the metadata elements is preserved and matters for equality checks:
Copy```
@dataclass
classctype:
   kind: str
a1 = Annotated[int, ValueRange(3, 10), ctype("char")]
a2 = Annotated[int, ctype("char"), ValueRange(3, 10)]
assert a1 != a2 # Order matters

```

It is up to the tool consuming the annotations to decide whether the client is allowed to add multiple metadata elements to one annotation and how to merge those annotations.
Nested `Annotated` types are flattened. The order of the metadata elements starts with the innermost annotation:
Copy```
assert Annotated[Annotated[int, ValueRange(3, 10)], ctype("char")] == Annotated[
  int, ValueRange(3, 10), ctype("char")
]

```

However, this does not apply to `Annotated` types referenced through a type alias, to avoid forcing evaluation of the underlying [`TypeAliasType`](https://docs.python.org/3/library/typing.html#typing.TypeAliasType "typing.TypeAliasType"):
Copy```
type From3To10[T] = Annotated[T, ValueRange(3, 10)]
assert Annotated[From3To10[int], ctype("char")] != Annotated[
  int, ValueRange(3, 10), ctype("char")
]

```

Duplicated metadata elements are not removed:
Copy```
assert Annotated[int, ValueRange(3, 10)] != Annotated[
  int, ValueRange(3, 10), ValueRange(3, 10)
]

```

`Annotated` can be used with nested and generic aliases:
> Copy```
@dataclass
classMaxLen:
  value: int
type Vec[T] = Annotated[list[tuple[T, T]], MaxLen(10)]
# When used in a type annotation, a type checker will treat "V" the same as
# ``Annotated[list[tuple[int, int]], MaxLen(10)]``:
type V = Vec[int]

```

`Annotated` cannot be used with an unpacked [`TypeVarTuple`](https://docs.python.org/3/library/typing.html#typing.TypeVarTuple "typing.TypeVarTuple"):
Copy```
type Variadic[*Ts] = Annotated[*Ts, Ann1] = Annotated[T1, T2, T3, ..., Ann1] # NOT valid

```

where `T1`, `T2`, … are [`TypeVars`](https://docs.python.org/3/library/typing.html#typing.TypeVar "typing.TypeVar"). This is invalid as only one type should be passed to Annotated.
By default, [`get_type_hints()`](https://docs.python.org/3/library/typing.html#typing.get_type_hints "typing.get_type_hints") strips the metadata from annotations. Pass `include_extras=True` to have the metadata preserved:
> Copy```
>>> fromtypingimport Annotated, get_type_hints
>>> deffunc(x: Annotated[int, "metadata"]) -> None: pass
...
>>> get_type_hints(func)
{'x': <class 'int'>, 'return': <class 'NoneType'>}
>>> get_type_hints(func, include_extras=True)
{'x': typing.Annotated[int, 'metadata'], 'return': <class 'NoneType'>}

```

At runtime, the metadata associated with an `Annotated` type can be retrieved via the `__metadata__` attribute:
> Copy```
>>> fromtypingimport Annotated
>>> X = Annotated[int, "very", "important", "metadata"]
>>> X
typing.Annotated[int, 'very', 'important', 'metadata']
>>> X.__metadata__
('very', 'important', 'metadata')

```

If you want to retrieve the original type wrapped by `Annotated`, use the `__origin__` attribute:
> Copy```
>>> fromtypingimport Annotated, get_origin
>>> Password = Annotated[str, "secret"]
>>> Password.__origin__
<class 'str'>

```

Note that using [`get_origin()`](https://docs.python.org/3/library/typing.html#typing.get_origin "typing.get_origin") will return `Annotated` itself:
> Copy```
>>> get_origin(Password)
typing.Annotated

```

See also 

[**PEP 593**](https://peps.python.org/pep-0593/) - Flexible function and variable annotations
    
The PEP introducing `Annotated` to the standard library.
Added in version 3.9. 

typing.TypeIs[¶](https://docs.python.org/3/library/typing.html#typing.TypeIs "Link to this definition")
    
Special typing construct for marking user-defined type predicate functions.
`TypeIs` can be used to annotate the return type of a user-defined type predicate function. `TypeIs` only accepts a single type argument. At runtime, functions marked this way should return a boolean and take at least one positional argument.
`TypeIs` aims to benefit _type narrowing_ – a technique used by static type checkers to determine a more precise type of an expression within a program’s code flow. Usually type narrowing is done by analyzing conditional code flow and applying the narrowing to a block of code. The conditional expression here is sometimes referred to as a “type predicate”:
Copy```
defis_str(val: str | float):
  # "isinstance" type predicate
  if isinstance(val, str):
    # Type of ``val`` is narrowed to ``str``
    ...
  else:
    # Else, type of ``val`` is narrowed to ``float``.
    ...

```

Sometimes it would be convenient to use a user-defined boolean function as a type predicate. Such a function should use `TypeIs[...]` or [`TypeGuard`](https://docs.python.org/3/library/typing.html#typing.TypeGuard "typing.TypeGuard") as its return type to alert static type checkers to this intention. `TypeIs` usually has more intuitive behavior than `TypeGuard`, but it cannot be used when the input and output types are incompatible (e.g., `list[object]` to `list[int]`) or when the function does not return `True` for all instances of the narrowed type.
Using `-> TypeIs[NarrowedType]` tells the static type checker that for a given function:
  1. The return value is a boolean.
  2. If the return value is `True`, the type of its argument is the intersection of the argument’s original type and `NarrowedType`.
  3. If the return value is `False`, the type of its argument is narrowed to exclude `NarrowedType`.


For example:
Copy```
fromtypingimport assert_type, final, TypeIs
classParent: pass
classChild(Parent): pass
@final
classUnrelated: pass
defis_parent(val: object) -> TypeIs[Parent]:
  return isinstance(val, Parent)
defrun(arg: Child | Unrelated):
  if is_parent(arg):
    # Type of ``arg`` is narrowed to the intersection
    # of ``Parent`` and ``Child``, which is equivalent to
    # ``Child``.
    assert_type(arg, Child)
  else:
    # Type of ``arg`` is narrowed to exclude ``Parent``,
    # so only ``Unrelated`` is left.
    assert_type(arg, Unrelated)

```

The type inside `TypeIs` must be consistent with the type of the function’s argument; if it is not, static type checkers will raise an error. An incorrectly written `TypeIs` function can lead to unsound behavior in the type system; it is the user’s responsibility to write such functions in a type-safe manner.
If a `TypeIs` function is a class or instance method, then the type in `TypeIs` maps to the type of the second parameter (after `cls` or `self`).
In short, the form `def foo(arg: TypeA) -> TypeIs[TypeB]: ...`, means that if `foo(arg)` returns `True`, then `arg` is an instance of `TypeB`, and if it returns `False`, it is not an instance of `TypeB`.
`TypeIs` also works with type variables. For more information, see [**PEP 742**](https://peps.python.org/pep-0742/) (Narrowing types with `TypeIs`).
Added in version 3.13. 

typing.TypeGuard[¶](https://docs.python.org/3/library/typing.html#typing.TypeGuard "Link to this definition")
    
Special typing construct for marking user-defined type predicate functions.
Type predicate functions are user-defined functions that return whether their argument is an instance of a particular type. `TypeGuard` works similarly to [`TypeIs`](https://docs.python.org/3/library/typing.html#typing.TypeIs "typing.TypeIs"), but has subtly different effects on type checking behavior (see below).
Using `-> TypeGuard` tells the static type checker that for a given function:
  1. The return value is a boolean.
  2. If the return value is `True`, the type of its argument is the type inside `TypeGuard`.


`TypeGuard` also works with type variables. See [**PEP 647**](https://peps.python.org/pep-0647/) for more details.
For example:
Copy```
defis_str_list(val: list[object]) -> TypeGuard[list[str]]:
'''Determines whether all objects in the list are strings'''
  return all(isinstance(x, str) for x in val)
deffunc1(val: list[object]):
  if is_str_list(val):
    # Type of ``val`` is narrowed to ``list[str]``.
    print(" ".join(val))
  else:
    # Type of ``val`` remains as ``list[object]``.
    print("Not a list of strings!")

```

`TypeIs` and `TypeGuard` differ in the following ways:
  * `TypeIs` requires the narrowed type to be a subtype of the input type, while `TypeGuard` does not. The main reason is to allow for things like narrowing `list[object]` to `list[str]` even though the latter is not a subtype of the former, since `list` is invariant.
  * When a `TypeGuard` function returns `True`, type checkers narrow the type of the variable to exactly the `TypeGuard` type. When a `TypeIs` function returns `True`, type checkers can infer a more precise type combining the previously known type of the variable with the `TypeIs` type. (Technically, this is known as an intersection type.)
  * When a `TypeGuard` function returns `False`, type checkers cannot narrow the type of the variable at all. When a `TypeIs` function returns `False`, type checkers can narrow the type of the variable to exclude the `TypeIs` type.


Added in version 3.10. 

typing.Unpack[¶](https://docs.python.org/3/library/typing.html#typing.Unpack "Link to this definition")
    
Typing operator to conceptually mark an object as having been unpacked.
For example, using the unpack operator `*` on a [type variable tuple](https://docs.python.org/3/library/typing.html#typevartuple) is equivalent to using `Unpack` to mark the type variable tuple as having been unpacked:
Copy```
Ts = TypeVarTuple('Ts')
tup: tuple[*Ts]
# Effectively does:
tup: tuple[Unpack[Ts]]

```

In fact, `Unpack` can be used interchangeably with `*` in the context of [`typing.TypeVarTuple`](https://docs.python.org/3/library/typing.html#typing.TypeVarTuple "typing.TypeVarTuple") and [`builtins.tuple`](https://docs.python.org/3/library/stdtypes.html#tuple "tuple") types. You might see `Unpack` being used explicitly in older versions of Python, where `*` couldn’t be used in certain places:
Copy```
# In older versions of Python, TypeVarTuple and Unpack
# are located in the `typing_extensions` backports package.
fromtyping_extensionsimport TypeVarTuple, Unpack
Ts = TypeVarTuple('Ts')
tup: tuple[*Ts]     # Syntax error on Python <= 3.10!
tup: tuple[Unpack[Ts]] # Semantically equivalent, and backwards-compatible

```

`Unpack` can also be used along with [`typing.TypedDict`](https://docs.python.org/3/library/typing.html#typing.TypedDict "typing.TypedDict") for typing `**kwargs` in a function signature:
Copy```
fromtypingimport TypedDict, Unpack
classMovie(TypedDict):
  name: str
  year: int
# This function expects two keyword arguments - `name` of type `str`
# and `year` of type `int`.
deffoo(**kwargs: Unpack[Movie]): ...

```

See [**PEP 692**](https://peps.python.org/pep-0692/) for more details on using `Unpack` for `**kwargs` typing.
Added in version 3.11.
#### Building generic types and type aliases[¶](https://docs.python.org/3/library/typing.html#building-generic-types-and-type-aliases "Link to this heading")
The following classes should not be used directly as annotations. Their intended purpose is to be building blocks for creating generic types and type aliases.
These objects can be created through special syntax ([type parameter lists](https://docs.python.org/3/reference/compound_stmts.html#type-params) and the [`type`](https://docs.python.org/3/reference/simple_stmts.html#type) statement). For compatibility with Python 3.11 and earlier, they can also be created without the dedicated syntax, as documented below. 

_class_ typing.Generic[¶](https://docs.python.org/3/library/typing.html#typing.Generic "Link to this definition")
    
Abstract base class for generic types.
A generic type is typically declared by adding a list of type parameters after the class name:
Copy```
classMapping[KT, VT]:
  def__getitem__(self, key: KT) -> VT:
    ...
    # Etc.

```

Such a class implicitly inherits from `Generic`. The runtime semantics of this syntax are discussed in the [Language Reference](https://docs.python.org/3/reference/compound_stmts.html#generic-classes).
This class can then be used as follows:
Copy```
deflookup_name[X, Y](mapping: Mapping[X, Y], key: X, default: Y) -> Y:
  try:
    return mapping[key]
  except KeyError:
    return default

```

Here the brackets after the function name indicate a [generic function](https://docs.python.org/3/reference/compound_stmts.html#generic-functions).
For backwards compatibility, generic classes can also be declared by explicitly inheriting from `Generic`. In this case, the type parameters must be declared separately:
Copy```
KT = TypeVar('KT')
VT = TypeVar('VT')
classMapping(Generic[KT, VT]):
  def__getitem__(self, key: KT) -> VT:
    ...
    # Etc.

```


_class_ typing.TypeVar(_name_ , _* constraints_, _bound =None_, _covariant =False_, _contravariant =False_, _infer_variance =False_, _default =typing.NoDefault_)[¶](https://docs.python.org/3/library/typing.html#typing.TypeVar "Link to this definition")
    
Type variable.
The preferred way to construct a type variable is via the dedicated syntax for [generic functions](https://docs.python.org/3/reference/compound_stmts.html#generic-functions), [generic classes](https://docs.python.org/3/reference/compound_stmts.html#generic-classes), and [generic type aliases](https://docs.python.org/3/reference/compound_stmts.html#generic-type-aliases):
Copy```
classSequence[T]: # T is a TypeVar
  ...

```

This syntax can also be used to create bounded and constrained type variables:
Copy```
classStrSequence[S: str]: # S is a TypeVar with a `str` upper bound;
  ...           # we can say that S is "bounded by `str`"

classStrOrBytesSequence[A: (str, bytes)]: # A is a TypeVar constrained to str or bytes
  ...

```

However, if desired, reusable type variables can also be constructed manually, like so:
Copy```
T = TypeVar('T') # Can be anything
S = TypeVar('S', bound=str) # Can be any subtype of str
A = TypeVar('A', str, bytes) # Must be exactly str or bytes

```

Type variables exist primarily for the benefit of static type checkers. They serve as the parameters for generic types as well as for generic function and type alias definitions. See [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic") for more information on generic types. Generic functions work as follows:
Copy```
defrepeat[T](x: T, n: int) -> Sequence[T]:
"""Return a list containing n references to x."""
  return [x]*n

defprint_capitalized[S: str](x: S) -> S:
"""Print x capitalized, and return x."""
  print(x.capitalize())
  return x

defconcatenate[A: (str, bytes)](x: A, y: A) -> A:
"""Add two strings or bytes objects together."""
  return x + y

```

Note that type variables can be _bounded_ , _constrained_ , or neither, but cannot be both bounded _and_ constrained.
The variance of type variables is inferred by type checkers when they are created through the [type parameter syntax](https://docs.python.org/3/reference/compound_stmts.html#type-params) or when `infer_variance=True` is passed. Manually created type variables may be explicitly marked covariant or contravariant by passing `covariant=True` or `contravariant=True`. By default, manually created type variables are invariant. See [**PEP 484**](https://peps.python.org/pep-0484/) and [**PEP 695**](https://peps.python.org/pep-0695/) for more details.
Bounded type variables and constrained type variables have different semantics in several important ways. Using a _bounded_ type variable means that the `TypeVar` will be solved using the most specific type possible:
Copy```
x = print_capitalized('a string')
reveal_type(x) # revealed type is str
classStringSubclass(str):
  pass
y = print_capitalized(StringSubclass('another string'))
reveal_type(y) # revealed type is StringSubclass
z = print_capitalized(45) # error: int is not a subtype of str

```

The upper bound of a type variable can be a concrete type, abstract type (ABC or Protocol), or even a union of types:
Copy```
# Can be anything with an __abs__ method
defprint_abs[T: SupportsAbs](arg: T) -> None:
  print("Absolute value:", abs(arg))
U = TypeVar('U', bound=str|bytes) # Can be any subtype of the union str|bytes
V = TypeVar('V', bound=SupportsAbs) # Can be anything with an __abs__ method

```

Using a _constrained_ type variable, however, means that the `TypeVar` can only ever be solved as being exactly one of the constraints given:
Copy```
a = concatenate('one', 'two')
reveal_type(a) # revealed type is str
b = concatenate(StringSubclass('one'), StringSubclass('two'))
reveal_type(b) # revealed type is str, despite StringSubclass being passed in
c = concatenate('one', b'two') # error: type variable 'A' can be either str or bytes in a function call, but not both

```

At runtime, `isinstance(x, T)` will raise [`TypeError`](https://docs.python.org/3/library/exceptions.html#TypeError "TypeError"). 

__name__[¶](https://docs.python.org/3/library/typing.html#typing.TypeVar.__name__ "Link to this definition")
    
The name of the type variable. 

__covariant__[¶](https://docs.python.org/3/library/typing.html#typing.TypeVar.__covariant__ "Link to this definition")
    
Whether the type var has been explicitly marked as covariant. 

__contravariant__[¶](https://docs.python.org/3/library/typing.html#typing.TypeVar.__contravariant__ "Link to this definition")
    
Whether the type var has been explicitly marked as contravariant. 

__infer_variance__[¶](https://docs.python.org/3/library/typing.html#typing.TypeVar.__infer_variance__ "Link to this definition")
    
Whether the type variable’s variance should be inferred by type checkers.
Added in version 3.12. 

__bound__[¶](https://docs.python.org/3/library/typing.html#typing.TypeVar.__bound__ "Link to this definition")
    
The upper bound of the type variable, if any.
Changed in version 3.12: For type variables created through [type parameter syntax](https://docs.python.org/3/reference/compound_stmts.html#type-params), the bound is evaluated only when the attribute is accessed, not when the type variable is created (see [Lazy evaluation](https://docs.python.org/3/reference/executionmodel.html#lazy-evaluation)). 

__constraints__[¶](https://docs.python.org/3/library/typing.html#typing.TypeVar.__constraints__ "Link to this definition")
    
A tuple containing the constraints of the type variable, if any.
Changed in version 3.12: For type variables created through [type parameter syntax](https://docs.python.org/3/reference/compound_stmts.html#type-params), the constraints are evaluated only when the attribute is accessed, not when the type variable is created (see [Lazy evaluation](https://docs.python.org/3/reference/executionmodel.html#lazy-evaluation)). 

__default__[¶](https://docs.python.org/3/library/typing.html#typing.TypeVar.__default__ "Link to this definition")
    
The default value of the type variable, or [`typing.NoDefault`](https://docs.python.org/3/library/typing.html#typing.NoDefault "typing.NoDefault") if it has no default.
Added in version 3.13. 

has_default()[¶](https://docs.python.org/3/library/typing.html#typing.TypeVar.has_default "Link to this definition")
    
Return whether or not the type variable has a default value. This is equivalent to checking whether [`__default__`](https://docs.python.org/3/library/typing.html#typing.TypeVar.__default__ "typing.TypeVar.__default__") is not the [`typing.NoDefault`](https://docs.python.org/3/library/typing.html#typing.NoDefault "typing.NoDefault") singleton, except that it does not force evaluation of the [lazily evaluated](https://docs.python.org/3/reference/executionmodel.html#lazy-evaluation) default value.
Added in version 3.13.
Changed in version 3.12: Type variables can now be declared using the [type parameter](https://docs.python.org/3/reference/compound_stmts.html#type-params) syntax introduced by [**PEP 695**](https://peps.python.org/pep-0695/). The `infer_variance` parameter was added.
Changed in version 3.13: Support for default values was added. 

_class_ typing.TypeVarTuple(_name_ , _*_ , _default =typing.NoDefault_)[¶](https://docs.python.org/3/library/typing.html#typing.TypeVarTuple "Link to this definition")
    
Type variable tuple. A specialized form of [type variable](https://docs.python.org/3/library/typing.html#typevar) that enables _variadic_ generics.
Type variable tuples can be declared in [type parameter lists](https://docs.python.org/3/reference/compound_stmts.html#type-params) using a single asterisk (`*`) before the name:
Copy```
defmove_first_element_to_last[T, *Ts](tup: tuple[T, *Ts]) -> tuple[*Ts, T]:
  return (*tup[1:], tup[0])

```

Or by explicitly invoking the `TypeVarTuple` constructor:
Copy```
T = TypeVar("T")
Ts = TypeVarTuple("Ts")
defmove_first_element_to_last(tup: tuple[T, *Ts]) -> tuple[*Ts, T]:
  return (*tup[1:], tup[0])

```

A normal type variable enables parameterization with a single type. A type variable tuple, in contrast, allows parameterization with an _arbitrary_ number of types by acting like an _arbitrary_ number of type variables wrapped in a tuple. For example:
Copy```
# T is bound to int, Ts is bound to ()
# Return value is (1,), which has type tuple[int]
move_first_element_to_last(tup=(1,))
# T is bound to int, Ts is bound to (str,)
# Return value is ('spam', 1), which has type tuple[str, int]
move_first_element_to_last(tup=(1, 'spam'))
# T is bound to int, Ts is bound to (str, float)
# Return value is ('spam', 3.0, 1), which has type tuple[str, float, int]
move_first_element_to_last(tup=(1, 'spam', 3.0))
# This fails to type check (and fails at runtime)
# because tuple[()] is not compatible with tuple[T, *Ts]
# (at least one element is required)
move_first_element_to_last(tup=())

```

Note the use of the unpacking operator `*` in `tuple[T, *Ts]`. Conceptually, you can think of `Ts` as a tuple of type variables `(T1, T2, ...)`. `tuple[T, *Ts]` would then become `tuple[T, *(T1, T2, ...)]`, which is equivalent to `tuple[T, T1, T2, ...]`. (Note that in older versions of Python, you might see this written using [`Unpack`](https://docs.python.org/3/library/typing.html#typing.Unpack "typing.Unpack") instead, as `Unpack[Ts]`.)
Type variable tuples must _always_ be unpacked. This helps distinguish type variable tuples from normal type variables:
Copy```
x: Ts     # Not valid
x: tuple[Ts]  # Not valid
x: tuple[*Ts] # The correct way to do it

```

Type variable tuples can be used in the same contexts as normal type variables. For example, in class definitions, arguments, and return types:
Copy```
classArray[*Shape]:
  def__getitem__(self, key: tuple[*Shape]) -> float: ...
  def__abs__(self) -> "Array[*Shape]": ...
  defget_shape(self) -> tuple[*Shape]: ...

```

Type variable tuples can be happily combined with normal type variables:
Copy```
classArray[DType, *Shape]: # This is fine
  pass
classArray2[*Shape, DType]: # This would also be fine
  pass
classHeight: ...
classWidth: ...
float_array_1d: Array[float, Height] = Array()   # Totally fine
int_array_2d: Array[int, Height, Width] = Array() # Yup, fine too

```

However, note that at most one type variable tuple may appear in a single list of type arguments or type parameters:
Copy```
x: tuple[*Ts, *Ts]      # Not valid
classArray[*Shape, *Shape]: # Not valid
  pass

```

Finally, an unpacked type variable tuple can be used as the type annotation of `*args`:
Copy```
defcall_soon[*Ts](
  callback: Callable[[*Ts], None],
  *args: *Ts
) -> None:
  ...
  callback(*args)

```

In contrast to non-unpacked annotations of `*args` - e.g. `*args: int`, which would specify that _all_ arguments are `int` - `*args: *Ts` enables reference to the types of the _individual_ arguments in `*args`. Here, this allows us to ensure the types of the `*args` passed to `call_soon` match the types of the (positional) arguments of `callback`.
See [**PEP 646**](https://peps.python.org/pep-0646/) for more details on type variable tuples. 

__name__[¶](https://docs.python.org/3/library/typing.html#typing.TypeVarTuple.__name__ "Link to this definition")
    
The name of the type variable tuple. 

__default__[¶](https://docs.python.org/3/library/typing.html#typing.TypeVarTuple.__default__ "Link to this definition")
    
The default value of the type variable tuple, or [`typing.NoDefault`](https://docs.python.org/3/library/typing.html#typing.NoDefault "typing.NoDefault") if it has no default.
Added in version 3.13. 

has_default()[¶](https://docs.python.org/3/library/typing.html#typing.TypeVarTuple.has_default "Link to this definition")
    
Return whether or not the type variable tuple has a default value. This is equivalent to checking whether [`__default__`](https://docs.python.org/3/library/typing.html#typing.TypeVarTuple.__default__ "typing.TypeVarTuple.__default__") is not the [`typing.NoDefault`](https://docs.python.org/3/library/typing.html#typing.NoDefault "typing.NoDefault") singleton, except that it does not force evaluation of the [lazily evaluated](https://docs.python.org/3/reference/executionmodel.html#lazy-evaluation) default value.
Added in version 3.13.
Added in version 3.11.
Changed in version 3.12: Type variable tuples can now be declared using the [type parameter](https://docs.python.org/3/reference/compound_stmts.html#type-params) syntax introduced by [**PEP 695**](https://peps.python.org/pep-0695/).
Changed in version 3.13: Support for default values was added. 

_class_ typing.ParamSpec(_name_ , _*_ , _bound =None_, _covariant =False_, _contravariant =False_, _default =typing.NoDefault_)[¶](https://docs.python.org/3/library/typing.html#typing.ParamSpec "Link to this definition")
    
Parameter specification variable. A specialized version of [type variables](https://docs.python.org/3/library/typing.html#typevar).
In [type parameter lists](https://docs.python.org/3/reference/compound_stmts.html#type-params), parameter specifications can be declared with two asterisks (`**`):
Copy```
type IntFunc[**P] = Callable[P, int]

```

For compatibility with Python 3.11 and earlier, `ParamSpec` objects can also be created as follows:
Copy```
P = ParamSpec('P')

```

Parameter specification variables exist primarily for the benefit of static type checkers. They are used to forward the parameter types of one callable to another callable – a pattern commonly found in higher order functions and decorators. They are only valid when used in `Concatenate`, or as the first argument to `Callable`, or as parameters for user-defined Generics. See [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic") for more information on generic types.
For example, to add basic logging to a function, one can create a decorator `add_logging` to log function calls. The parameter specification variable tells the type checker that the callable passed into the decorator and the new callable returned by it have inter-dependent type parameters:
Copy```
fromcollections.abcimport Callable
importlogging
defadd_logging[T, **P](f: Callable[P, T]) -> Callable[P, T]:
'''A type-safe decorator to add logging to a function.'''
  definner(*args: P.args, **kwargs: P.kwargs) -> T:
    logging.info(f'{f.__name__} was called')
    return f(*args, **kwargs)
  return inner
@add_logging
defadd_two(x: float, y: float) -> float:
'''Add two numbers together.'''
  return x + y

```

Without `ParamSpec`, the simplest way to annotate this previously was to use a [`TypeVar`](https://docs.python.org/3/library/typing.html#typing.TypeVar "typing.TypeVar") with upper bound `Callable[..., Any]`. However this causes two problems:
  1. The type checker can’t type check the `inner` function because `*args` and `**kwargs` have to be typed [`Any`](https://docs.python.org/3/library/typing.html#typing.Any "typing.Any").
  2. [`cast()`](https://docs.python.org/3/library/typing.html#typing.cast "typing.cast") may be required in the body of the `add_logging` decorator when returning the `inner` function, or the static type checker must be told to ignore the `return inner`.



args[¶](https://docs.python.org/3/library/typing.html#typing.ParamSpec.args "Link to this definition")


kwargs[¶](https://docs.python.org/3/library/typing.html#typing.ParamSpec.kwargs "Link to this definition")
    
Since `ParamSpec` captures both positional and keyword parameters, `P.args` and `P.kwargs` can be used to split a `ParamSpec` into its components. `P.args` represents the tuple of positional parameters in a given call and should only be used to annotate `*args`. `P.kwargs` represents the mapping of keyword parameters to their values in a given call, and should be only be used to annotate `**kwargs`. Both attributes require the annotated parameter to be in scope. At runtime, `P.args` and `P.kwargs` are instances respectively of [`ParamSpecArgs`](https://docs.python.org/3/library/typing.html#typing.ParamSpecArgs "typing.ParamSpecArgs") and [`ParamSpecKwargs`](https://docs.python.org/3/library/typing.html#typing.ParamSpecKwargs "typing.ParamSpecKwargs"). 

__name__[¶](https://docs.python.org/3/library/typing.html#typing.ParamSpec.__name__ "Link to this definition")
    
The name of the parameter specification. 

__default__[¶](https://docs.python.org/3/library/typing.html#typing.ParamSpec.__default__ "Link to this definition")
    
The default value of the parameter specification, or [`typing.NoDefault`](https://docs.python.org/3/library/typing.html#typing.NoDefault "typing.NoDefault") if it has no default.
Added in version 3.13. 

has_default()[¶](https://docs.python.org/3/library/typing.html#typing.ParamSpec.has_default "Link to this definition")
    
Return whether or not the parameter specification has a default value. This is equivalent to checking whether [`__default__`](https://docs.python.org/3/library/typing.html#typing.ParamSpec.__default__ "typing.ParamSpec.__default__") is not the [`typing.NoDefault`](https://docs.python.org/3/library/typing.html#typing.NoDefault "typing.NoDefault") singleton, except that it does not force evaluation of the [lazily evaluated](https://docs.python.org/3/reference/executionmodel.html#lazy-evaluation) default value.
Added in version 3.13.
Parameter specification variables created with `covariant=True` or `contravariant=True` can be used to declare covariant or contravariant generic types. The `bound` argument is also accepted, similar to [`TypeVar`](https://docs.python.org/3/library/typing.html#typing.TypeVar "typing.TypeVar"). However the actual semantics of these keywords are yet to be decided.
Added in version 3.10.
Changed in version 3.12: Parameter specifications can now be declared using the [type parameter](https://docs.python.org/3/reference/compound_stmts.html#type-params) syntax introduced by [**PEP 695**](https://peps.python.org/pep-0695/).
Changed in version 3.13: Support for default values was added.
Note
Only parameter specification variables defined in global scope can be pickled.
See also
  * [**PEP 612**](https://peps.python.org/pep-0612/) – Parameter Specification Variables (the PEP which introduced `ParamSpec` and `Concatenate`)
  * [`Concatenate`](https://docs.python.org/3/library/typing.html#typing.Concatenate "typing.Concatenate")
  * [Annotating callable objects](https://docs.python.org/3/library/typing.html#annotating-callables)



typing.ParamSpecArgs[¶](https://docs.python.org/3/library/typing.html#typing.ParamSpecArgs "Link to this definition")


typing.ParamSpecKwargs[¶](https://docs.python.org/3/library/typing.html#typing.ParamSpecKwargs "Link to this definition")
    
Arguments and keyword arguments attributes of a [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec"). The `P.args` attribute of a `ParamSpec` is an instance of `ParamSpecArgs`, and `P.kwargs` is an instance of `ParamSpecKwargs`. They are intended for runtime introspection and have no special meaning to static type checkers.
Calling [`get_origin()`](https://docs.python.org/3/library/typing.html#typing.get_origin "typing.get_origin") on either of these objects will return the original `ParamSpec`:
Copy```
>>> fromtypingimport ParamSpec, get_origin
>>> P = ParamSpec("P")
>>> get_origin(P.args) is P
True
>>> get_origin(P.kwargs) is P
True

```

Added in version 3.10. 

_class_ typing.TypeAliasType(_name_ , _value_ , _*_ , _type_params =()_)[¶](https://docs.python.org/3/library/typing.html#typing.TypeAliasType "Link to this definition")
    
The type of type aliases created through the [`type`](https://docs.python.org/3/reference/simple_stmts.html#type) statement.
Example:
Copy```
>>> type Alias = int
>>> type(Alias)
<class 'typing.TypeAliasType'>

```

Added in version 3.12. 

__name__[¶](https://docs.python.org/3/library/typing.html#typing.TypeAliasType.__name__ "Link to this definition")
    
The name of the type alias:
Copy```
>>> type Alias = int
>>> Alias.__name__
'Alias'

```


__module__[¶](https://docs.python.org/3/library/typing.html#typing.TypeAliasType.__module__ "Link to this definition")
    
The module in which the type alias was defined:
Copy```
>>> type Alias = int
>>> Alias.__module__
'__main__'

```


__type_params__[¶](https://docs.python.org/3/library/typing.html#typing.TypeAliasType.__type_params__ "Link to this definition")
    
The type parameters of the type alias, or an empty tuple if the alias is not generic:
Copy```
>>> type ListOrSet[T] = list[T] | set[T]
>>> ListOrSet.__type_params__
(T,)
>>> type NotGeneric = int
>>> NotGeneric.__type_params__
()

```


__value__[¶](https://docs.python.org/3/library/typing.html#typing.TypeAliasType.__value__ "Link to this definition")
    
The type alias’s value. This is [lazily evaluated](https://docs.python.org/3/reference/executionmodel.html#lazy-evaluation), so names used in the definition of the alias are not resolved until the `__value__` attribute is accessed:
Copy```
>>> type Mutually = Recursive
>>> type Recursive = Mutually
>>> Mutually
Mutually
>>> Recursive
Recursive
>>> Mutually.__value__
Recursive
>>> Recursive.__value__
Mutually

```

#### Other special directives[¶](https://docs.python.org/3/library/typing.html#other-special-directives "Link to this heading")
These functions and classes should not be used directly as annotations. Their intended purpose is to be building blocks for creating and declaring types. 

_class_ typing.NamedTuple[¶](https://docs.python.org/3/library/typing.html#typing.NamedTuple "Link to this definition")
    
Typed version of [`collections.namedtuple()`](https://docs.python.org/3/library/collections.html#collections.namedtuple "collections.namedtuple").
Usage:
Copy```
classEmployee(NamedTuple):
  name: str
  id: int

```

This is equivalent to:
Copy```
Employee = collections.namedtuple('Employee', ['name', 'id'])

```

To give a field a default value, you can assign to it in the class body:
Copy```
classEmployee(NamedTuple):
  name: str
  id: int = 3
employee = Employee('Guido')
assert employee.id == 3

```

Fields with a default value must come after any fields without a default.
The resulting class has an extra attribute `__annotations__` giving a dict that maps the field names to the field types. (The field names are in the `_fields` attribute and the default values are in the `_field_defaults` attribute, both of which are part of the [`namedtuple()`](https://docs.python.org/3/library/collections.html#collections.namedtuple "collections.namedtuple") API.)
`NamedTuple` subclasses can also have docstrings and methods:
Copy```
classEmployee(NamedTuple):
"""Represents an employee."""
  name: str
  id: int = 3
  def__repr__(self) -> str:
    return f'<Employee {self.name}, id={self.id}>'

```

`NamedTuple` subclasses can be generic:
Copy```
classGroup[T](NamedTuple):
  key: T
  group: list[T]

```

Backward-compatible usage:
Copy```
# For creating a generic NamedTuple on Python 3.11
T = TypeVar("T")
classGroup(NamedTuple, Generic[T]):
  key: T
  group: list[T]
# A functional syntax is also supported
Employee = NamedTuple('Employee', [('name', str), ('id', int)])

```

Changed in version 3.6: Added support for [**PEP 526**](https://peps.python.org/pep-0526/) variable annotation syntax.
Changed in version 3.6.1: Added support for default values, methods, and docstrings.
Changed in version 3.8: The `_field_types` and `__annotations__` attributes are now regular dictionaries instead of instances of `OrderedDict`.
Changed in version 3.9: Removed the `_field_types` attribute in favor of the more standard `__annotations__` attribute which has the same information.
Changed in version 3.11: Added support for generic namedtuples.
Deprecated since version 3.13, will be removed in version 3.15: The undocumented keyword argument syntax for creating NamedTuple classes (`NT = NamedTuple("NT", x=int)`) is deprecated, and will be disallowed in 3.15. Use the class-based syntax or the functional syntax instead.
Deprecated since version 3.13, will be removed in version 3.15: When using the functional syntax to create a NamedTuple class, failing to pass a value to the ‘fields’ parameter (`NT = NamedTuple("NT")`) is deprecated. Passing `None` to the ‘fields’ parameter (`NT = NamedTuple("NT", None)`) is also deprecated. Both will be disallowed in Python 3.15. To create a NamedTuple class with 0 fields, use `class NT(NamedTuple): pass` or `NT = NamedTuple("NT", [])`. 

_class_ typing.NewType(_name_ , _tp_)[¶](https://docs.python.org/3/library/typing.html#typing.NewType "Link to this definition")
    
Helper class to create low-overhead [distinct types](https://docs.python.org/3/library/typing.html#distinct).
A `NewType` is considered a distinct type by a typechecker. At runtime, however, calling a `NewType` returns its argument unchanged.
Usage:
Copy```
UserId = NewType('UserId', int) # Declare the NewType "UserId"
first_user = UserId(1) # "UserId" returns the argument unchanged at runtime

```


__module__[¶](https://docs.python.org/3/library/typing.html#typing.NewType.__module__ "Link to this definition")
    
The module in which the new type is defined. 

__name__[¶](https://docs.python.org/3/library/typing.html#typing.NewType.__name__ "Link to this definition")
    
The name of the new type. 

__supertype__[¶](https://docs.python.org/3/library/typing.html#typing.NewType.__supertype__ "Link to this definition")
    
The type that the new type is based on.
Added in version 3.5.2.
Changed in version 3.10: `NewType` is now a class rather than a function. 

_class_ typing.Protocol(_Generic_)[¶](https://docs.python.org/3/library/typing.html#typing.Protocol "Link to this definition")
    
Base class for protocol classes.
Protocol classes are defined like this:
Copy```
classProto(Protocol):
  defmeth(self) -> int:
    ...

```

Such classes are primarily used with static type checkers that recognize structural subtyping (static duck-typing), for example:
Copy```
classC:
  defmeth(self) -> int:
    return 0
deffunc(x: Proto) -> int:
  return x.meth()
func(C()) # Passes static type check

```

See [**PEP 544**](https://peps.python.org/pep-0544/) for more details. Protocol classes decorated with [`runtime_checkable()`](https://docs.python.org/3/library/typing.html#typing.runtime_checkable "typing.runtime_checkable") (described later) act as simple-minded runtime protocols that check only the presence of given attributes, ignoring their type signatures. Protocol classes without this decorator cannot be used as the second argument to [`isinstance()`](https://docs.python.org/3/library/functions.html#isinstance "isinstance") or [`issubclass()`](https://docs.python.org/3/library/functions.html#issubclass "issubclass").
Protocol classes can be generic, for example:
Copy```
classGenProto[T](Protocol):
  defmeth(self) -> T:
    ...

```

In code that needs to be compatible with Python 3.11 or older, generic Protocols can be written as follows:
Copy```
T = TypeVar("T")
classGenProto(Protocol[T]):
  defmeth(self) -> T:
    ...

```

Added in version 3.8. 

@typing.runtime_checkable[¶](https://docs.python.org/3/library/typing.html#typing.runtime_checkable "Link to this definition")
    
Mark a protocol class as a runtime protocol.
Such a protocol can be used with [`isinstance()`](https://docs.python.org/3/library/functions.html#isinstance "isinstance") and [`issubclass()`](https://docs.python.org/3/library/functions.html#issubclass "issubclass"). This allows a simple-minded structural check, very similar to “one trick ponies” in [`collections.abc`](https://docs.python.org/3/library/collections.abc.html#module-collections.abc "collections.abc: Abstract base classes for containers") such as [`Iterable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Iterable "collections.abc.Iterable"). For example:
Copy```
@runtime_checkable
classClosable(Protocol):
  defclose(self): ...
assert isinstance(open('/some/file'), Closable)
@runtime_checkable
classNamed(Protocol):
  name: str
importthreading
assert isinstance(threading.Thread(name='Bob'), Named)

```

This decorator raises [`TypeError`](https://docs.python.org/3/library/exceptions.html#TypeError "TypeError") when applied to a non-protocol class.
Note
`runtime_checkable()` will check only the presence of the required methods or attributes, not their type signatures or types. For example, [`ssl.SSLObject`](https://docs.python.org/3/library/ssl.html#ssl.SSLObject "ssl.SSLObject") is a class, therefore it passes an [`issubclass()`](https://docs.python.org/3/library/functions.html#issubclass "issubclass") check against [Callable](https://docs.python.org/3/library/typing.html#annotating-callables). However, the `ssl.SSLObject.__init__` method exists only to raise a [`TypeError`](https://docs.python.org/3/library/exceptions.html#TypeError "TypeError") with a more informative message, therefore making it impossible to call (instantiate) [`ssl.SSLObject`](https://docs.python.org/3/library/ssl.html#ssl.SSLObject "ssl.SSLObject").
Note
An [`isinstance()`](https://docs.python.org/3/library/functions.html#isinstance "isinstance") check against a runtime-checkable protocol can be surprisingly slow compared to an `isinstance()` check against a non-protocol class. Consider using alternative idioms such as [`hasattr()`](https://docs.python.org/3/library/functions.html#hasattr "hasattr") calls for structural checks in performance-sensitive code.
Added in version 3.8.
Changed in version 3.12: The internal implementation of [`isinstance()`](https://docs.python.org/3/library/functions.html#isinstance "isinstance") checks against runtime-checkable protocols now uses [`inspect.getattr_static()`](https://docs.python.org/3/library/inspect.html#inspect.getattr_static "inspect.getattr_static") to look up attributes (previously, [`hasattr()`](https://docs.python.org/3/library/functions.html#hasattr "hasattr") was used). As a result, some objects which used to be considered instances of a runtime-checkable protocol may no longer be considered instances of that protocol on Python 3.12+, and vice versa. Most users are unlikely to be affected by this change.
Changed in version 3.12: The members of a runtime-checkable protocol are now considered “frozen” at runtime as soon as the class has been created. Monkey-patching attributes onto a runtime-checkable protocol will still work, but will have no impact on [`isinstance()`](https://docs.python.org/3/library/functions.html#isinstance "isinstance") checks comparing objects to the protocol. See [What’s new in Python 3.12](https://docs.python.org/3/whatsnew/3.12.html#whatsnew-typing-py312) for more details. 

_class_ typing.TypedDict(_dict_)[¶](https://docs.python.org/3/library/typing.html#typing.TypedDict "Link to this definition")
    
Special construct to add type hints to a dictionary. At runtime it is a plain [`dict`](https://docs.python.org/3/library/stdtypes.html#dict "dict").
`TypedDict` declares a dictionary type that expects all of its instances to have a certain set of keys, where each key is associated with a value of a consistent type. This expectation is not checked at runtime but is only enforced by type checkers. Usage:
Copy```
classPoint2D(TypedDict):
  x: int
  y: int
  label: str
a: Point2D = {'x': 1, 'y': 2, 'label': 'good'} # OK
b: Point2D = {'z': 3, 'label': 'bad'}      # Fails type check
assert Point2D(x=1, y=2, label='first') == dict(x=1, y=2, label='first')

```

An alternative way to create a `TypedDict` is by using function-call syntax. The second argument must be a literal [`dict`](https://docs.python.org/3/library/stdtypes.html#dict "dict"):
Copy```
Point2D = TypedDict('Point2D', {'x': int, 'y': int, 'label': str})

```

This functional syntax allows defining keys which are not valid [identifiers](https://docs.python.org/3/reference/lexical_analysis.html#identifiers), for example because they are keywords or contain hyphens, or when key names must not be [mangled](https://docs.python.org/3/reference/expressions.html#private-name-mangling) like regular private names:
Copy```
# raises SyntaxError
classPoint2D(TypedDict):
  in: int # 'in' is a keyword
  x-y: int # name with hyphens
classDefinition(TypedDict):
  __schema: str # mangled to `_Definition__schema`
# OK, functional syntax
Point2D = TypedDict('Point2D', {'in': int, 'x-y': int})
Definition = TypedDict('Definition', {'__schema': str}) # not mangled

```

By default, all keys must be present in a `TypedDict`. It is possible to mark individual keys as non-required using [`NotRequired`](https://docs.python.org/3/library/typing.html#typing.NotRequired "typing.NotRequired"):
Copy```
classPoint2D(TypedDict):
  x: int
  y: int
  label: NotRequired[str]
# Alternative syntax
Point2D = TypedDict('Point2D', {'x': int, 'y': int, 'label': NotRequired[str]})

```

This means that a `Point2D` `TypedDict` can have the `label` key omitted.
It is also possible to mark all keys as non-required by default by specifying a totality of `False`:
Copy```
classPoint2D(TypedDict, total=False):
  x: int
  y: int
# Alternative syntax
Point2D = TypedDict('Point2D', {'x': int, 'y': int}, total=False)

```

This means that a `Point2D` `TypedDict` can have any of the keys omitted. A type checker is only expected to support a literal `False` or `True` as the value of the `total` argument. `True` is the default, and makes all items defined in the class body required.
Individual keys of a `total=False` `TypedDict` can be marked as required using [`Required`](https://docs.python.org/3/library/typing.html#typing.Required "typing.Required"):
Copy```
classPoint2D(TypedDict, total=False):
  x: Required[int]
  y: Required[int]
  label: str
# Alternative syntax
Point2D = TypedDict('Point2D', {
  'x': Required[int],
  'y': Required[int],
  'label': str
}, total=False)

```

It is possible for a `TypedDict` type to inherit from one or more other `TypedDict` types using the class-based syntax. Usage:
Copy```
classPoint3D(Point2D):
  z: int

```

`Point3D` has three items: `x`, `y` and `z`. It is equivalent to this definition:
Copy```
classPoint3D(TypedDict):
  x: int
  y: int
  z: int

```

A `TypedDict` cannot inherit from a non-`TypedDict` class, except for [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic"). For example:
Copy```
classX(TypedDict):
  x: int
classY(TypedDict):
  y: int
classZ(object): pass # A non-TypedDict class
classXY(X, Y): pass # OK
classXZ(X, Z): pass # raises TypeError

```

A `TypedDict` can be generic:
Copy```
classGroup[T](TypedDict):
  key: T
  group: list[T]

```

To create a generic `TypedDict` that is compatible with Python 3.11 or lower, inherit from [`Generic`](https://docs.python.org/3/library/typing.html#typing.Generic "typing.Generic") explicitly:
Copy```
T = TypeVar("T")
classGroup(TypedDict, Generic[T]):
  key: T
  group: list[T]

```

A `TypedDict` can be introspected via annotations dicts (see [Annotations Best Practices](https://docs.python.org/3/howto/annotations.html#annotations-howto) for more information on annotations best practices), [`__total__`](https://docs.python.org/3/library/typing.html#typing.TypedDict.__total__ "typing.TypedDict.__total__"), [`__required_keys__`](https://docs.python.org/3/library/typing.html#typing.TypedDict.__required_keys__ "typing.TypedDict.__required_keys__"), and [`__optional_keys__`](https://docs.python.org/3/library/typing.html#typing.TypedDict.__optional_keys__ "typing.TypedDict.__optional_keys__"). 

__total__[¶](https://docs.python.org/3/library/typing.html#typing.TypedDict.__total__ "Link to this definition")
    
`Point2D.__total__` gives the value of the `total` argument. Example:
Copy```
>>> fromtypingimport TypedDict
>>> classPoint2D(TypedDict): pass
>>> Point2D.__total__
True
>>> classPoint2D(TypedDict, total=False): pass
>>> Point2D.__total__
False
>>> classPoint3D(Point2D): pass
>>> Point3D.__total__
True

```

This attribute reflects _only_ the value of the `total` argument to the current `TypedDict` class, not whether the class is semantically total. For example, a `TypedDict` with `__total__` set to `True` may have keys marked with [`NotRequired`](https://docs.python.org/3/library/typing.html#typing.NotRequired "typing.NotRequired"), or it may inherit from another `TypedDict` with `total=False`. Therefore, it is generally better to use [`__required_keys__`](https://docs.python.org/3/library/typing.html#typing.TypedDict.__required_keys__ "typing.TypedDict.__required_keys__") and [`__optional_keys__`](https://docs.python.org/3/library/typing.html#typing.TypedDict.__optional_keys__ "typing.TypedDict.__optional_keys__") for introspection. 

__required_keys__[¶](https://docs.python.org/3/library/typing.html#typing.TypedDict.__required_keys__ "Link to this definition")
    
Added in version 3.9. 

__optional_keys__[¶](https://docs.python.org/3/library/typing.html#typing.TypedDict.__optional_keys__ "Link to this definition")
    
`Point2D.__required_keys__` and `Point2D.__optional_keys__` return [`frozenset`](https://docs.python.org/3/library/stdtypes.html#frozenset "frozenset") objects containing required and non-required keys, respectively.
Keys marked with [`Required`](https://docs.python.org/3/library/typing.html#typing.Required "typing.Required") will always appear in `__required_keys__` and keys marked with [`NotRequired`](https://docs.python.org/3/library/typing.html#typing.NotRequired "typing.NotRequired") will always appear in `__optional_keys__`.
For backwards compatibility with Python 3.10 and below, it is also possible to use inheritance to declare both required and non-required keys in the same `TypedDict` . This is done by declaring a `TypedDict` with one value for the `total` argument and then inheriting from it in another `TypedDict` with a different value for `total`:
Copy```
>>> classPoint2D(TypedDict, total=False):
...   x: int
...   y: int
...
>>> classPoint3D(Point2D):
...   z: int
...
>>> Point3D.__required_keys__ == frozenset({'z'})
True
>>> Point3D.__optional_keys__ == frozenset({'x', 'y'})
True

```

Added in version 3.9.
Note
If `from __future__ import annotations` is used or if annotations are given as strings, annotations are not evaluated when the `TypedDict` is defined. Therefore, the runtime introspection that `__required_keys__` and `__optional_keys__` rely on may not work properly, and the values of the attributes may be incorrect.
Support for [`ReadOnly`](https://docs.python.org/3/library/typing.html#typing.ReadOnly "typing.ReadOnly") is reflected in the following attributes: 

__readonly_keys__[¶](https://docs.python.org/3/library/typing.html#typing.TypedDict.__readonly_keys__ "Link to this definition")
    
A [`frozenset`](https://docs.python.org/3/library/stdtypes.html#frozenset "frozenset") containing the names of all read-only keys. Keys are read-only if they carry the [`ReadOnly`](https://docs.python.org/3/library/typing.html#typing.ReadOnly "typing.ReadOnly") qualifier.
Added in version 3.13. 

__mutable_keys__[¶](https://docs.python.org/3/library/typing.html#typing.TypedDict.__mutable_keys__ "Link to this definition")
    
A [`frozenset`](https://docs.python.org/3/library/stdtypes.html#frozenset "frozenset") containing the names of all mutable keys. Keys are mutable if they do not carry the [`ReadOnly`](https://docs.python.org/3/library/typing.html#typing.ReadOnly "typing.ReadOnly") qualifier.
Added in version 3.13.
See [**PEP 589**](https://peps.python.org/pep-0589/) for more examples and detailed rules of using `TypedDict`.
Added in version 3.8.
Changed in version 3.11: Added support for marking individual keys as [`Required`](https://docs.python.org/3/library/typing.html#typing.Required "typing.Required") or [`NotRequired`](https://docs.python.org/3/library/typing.html#typing.NotRequired "typing.NotRequired"). See [**PEP 655**](https://peps.python.org/pep-0655/).
Changed in version 3.11: Added support for generic `TypedDict`s.
Changed in version 3.13: Removed support for the keyword-argument method of creating `TypedDict`s.
Changed in version 3.13: Support for the [`ReadOnly`](https://docs.python.org/3/library/typing.html#typing.ReadOnly "typing.ReadOnly") qualifier was added.
Deprecated since version 3.13, will be removed in version 3.15: When using the functional syntax to create a TypedDict class, failing to pass a value to the ‘fields’ parameter (`TD = TypedDict("TD")`) is deprecated. Passing `None` to the ‘fields’ parameter (`TD = TypedDict("TD", None)`) is also deprecated. Both will be disallowed in Python 3.15. To create a TypedDict class with 0 fields, use `class TD(TypedDict): pass` or `TD = TypedDict("TD", {})`.
### Protocols[¶](https://docs.python.org/3/library/typing.html#protocols "Link to this heading")
The following protocols are provided by the `typing` module. All are decorated with [`@runtime_checkable`](https://docs.python.org/3/library/typing.html#typing.runtime_checkable "typing.runtime_checkable"). 

_class_ typing.SupportsAbs[¶](https://docs.python.org/3/library/typing.html#typing.SupportsAbs "Link to this definition")
    
An ABC with one abstract method `__abs__` that is covariant in its return type. 

_class_ typing.SupportsBytes[¶](https://docs.python.org/3/library/typing.html#typing.SupportsBytes "Link to this definition")
    
An ABC with one abstract method `__bytes__`. 

_class_ typing.SupportsComplex[¶](https://docs.python.org/3/library/typing.html#typing.SupportsComplex "Link to this definition")
    
An ABC with one abstract method `__complex__`. 

_class_ typing.SupportsFloat[¶](https://docs.python.org/3/library/typing.html#typing.SupportsFloat "Link to this definition")
    
An ABC with one abstract method `__float__`. 

_class_ typing.SupportsIndex[¶](https://docs.python.org/3/library/typing.html#typing.SupportsIndex "Link to this definition")
    
An ABC with one abstract method `__index__`.
Added in version 3.8. 

_class_ typing.SupportsInt[¶](https://docs.python.org/3/library/typing.html#typing.SupportsInt "Link to this definition")
    
An ABC with one abstract method `__int__`. 

_class_ typing.SupportsRound[¶](https://docs.python.org/3/library/typing.html#typing.SupportsRound "Link to this definition")
    
An ABC with one abstract method `__round__` that is covariant in its return type.
### ABCs for working with IO[¶](https://docs.python.org/3/library/typing.html#abcs-for-working-with-io "Link to this heading") 

_class_ typing.IO[¶](https://docs.python.org/3/library/typing.html#typing.IO "Link to this definition")


_class_ typing.TextIO[¶](https://docs.python.org/3/library/typing.html#typing.TextIO "Link to this definition")


_class_ typing.BinaryIO[¶](https://docs.python.org/3/library/typing.html#typing.BinaryIO "Link to this definition")
    
Generic type `IO[AnyStr]` and its subclasses `TextIO(IO[str])` and `BinaryIO(IO[bytes])` represent the types of I/O streams such as returned by [`open()`](https://docs.python.org/3/library/functions.html#open "open").
### Functions and decorators[¶](https://docs.python.org/3/library/typing.html#functions-and-decorators "Link to this heading") 

typing.cast(_typ_ , _val_)[¶](https://docs.python.org/3/library/typing.html#typing.cast "Link to this definition")
    
Cast a value to a type.
This returns the value unchanged. To the type checker this signals that the return value has the designated type, but at runtime we intentionally don’t check anything (we want this to be as fast as possible). 

typing.assert_type(_val_ , _typ_ , _/_)[¶](https://docs.python.org/3/library/typing.html#typing.assert_type "Link to this definition")
    
Ask a static type checker to confirm that _val_ has an inferred type of _typ_.
At runtime this does nothing: it returns the first argument unchanged with no checks or side effects, no matter the actual type of the argument.
When a static type checker encounters a call to `assert_type()`, it emits an error if the value is not of the specified type:
Copy```
defgreet(name: str) -> None:
  assert_type(name, str) # OK, inferred type of `name` is `str`
  assert_type(name, int) # type checker error

```

This function is useful for ensuring the type checker’s understanding of a script is in line with the developer’s intentions:
Copy```
defcomplex_function(arg: object):
  # Do some complex type-narrowing logic,
  # after which we hope the inferred type will be `int`
  ...
  # Test whether the type checker correctly understands our function
  assert_type(arg, int)

```

Added in version 3.11. 

typing.assert_never(_arg_ , _/_)[¶](https://docs.python.org/3/library/typing.html#typing.assert_never "Link to this definition")
    
Ask a static type checker to confirm that a line of code is unreachable.
Example:
Copy```
defint_or_str(arg: int | str) -> None:
  match arg:
    case int():
      print("It's an int")
    case str():
      print("It's a str")
    case_ as unreachable:
      assert_never(unreachable)

```

Here, the annotations allow the type checker to infer that the last case can never execute, because `arg` is either an [`int`](https://docs.python.org/3/library/functions.html#int "int") or a [`str`](https://docs.python.org/3/library/stdtypes.html#str "str"), and both options are covered by earlier cases.
If a type checker finds that a call to `assert_never()` is reachable, it will emit an error. For example, if the type annotation for `arg` was instead `int | str | float`, the type checker would emit an error pointing out that `unreachable` is of type [`float`](https://docs.python.org/3/library/functions.html#float "float"). For a call to `assert_never` to pass type checking, the inferred type of the argument passed in must be the bottom type, [`Never`](https://docs.python.org/3/library/typing.html#typing.Never "typing.Never"), and nothing else.
At runtime, this throws an exception when called.
See also
[Unreachable Code and Exhaustiveness Checking](https://typing.python.org/en/latest/guides/unreachable.html) has more information about exhaustiveness checking with static typing.
Added in version 3.11. 

typing.reveal_type(_obj_ , _/_)[¶](https://docs.python.org/3/library/typing.html#typing.reveal_type "Link to this definition")
    
Ask a static type checker to reveal the inferred type of an expression.
When a static type checker encounters a call to this function, it emits a diagnostic with the inferred type of the argument. For example:
Copy```
x: int = 1
reveal_type(x) # Revealed type is "builtins.int"

```

This can be useful when you want to debug how your type checker handles a particular piece of code.
At runtime, this function prints the runtime type of its argument to [`sys.stderr`](https://docs.python.org/3/library/sys.html#sys.stderr "sys.stderr") and returns the argument unchanged (allowing the call to be used within an expression):
Copy```
x = reveal_type(1) # prints "Runtime type is int"
print(x) # prints "1"

```

Note that the runtime type may be different from (more or less specific than) the type statically inferred by a type checker.
Most type checkers support `reveal_type()` anywhere, even if the name is not imported from `typing`. Importing the name from `typing`, however, allows your code to run without runtime errors and communicates intent more clearly.
Added in version 3.11. 

@typing.dataclass_transform(_*_ , _eq_default =True_, _order_default =False_, _kw_only_default =False_, _frozen_default =False_, _field_specifiers =()_, _** kwargs_)[¶](https://docs.python.org/3/library/typing.html#typing.dataclass_transform "Link to this definition")
    
Decorator to mark an object as providing [`dataclass`](https://docs.python.org/3/library/dataclasses.html#dataclasses.dataclass "dataclasses.dataclass")-like behavior.
`dataclass_transform` may be used to decorate a class, metaclass, or a function that is itself a decorator. The presence of `@dataclass_transform()` tells a static type checker that the decorated object performs runtime “magic” that transforms a class in a similar way to [`@dataclasses.dataclass`](https://docs.python.org/3/library/dataclasses.html#dataclasses.dataclass "dataclasses.dataclass").
Example usage with a decorator function:
Copy```
@dataclass_transform()
defcreate_model[T](cls: type[T]) -> type[T]:
  ...
  return cls
@create_model
classCustomerModel:
  id: int
  name: str

```

On a base class:
Copy```
@dataclass_transform()
classModelBase: ...
classCustomerModel(ModelBase):
  id: int
  name: str

```

On a metaclass:
Copy```
@dataclass_transform()
classModelMeta(type): ...
classModelBase(metaclass=ModelMeta): ...
classCustomerModel(ModelBase):
  id: int
  name: str

```

The `CustomerModel` classes defined above will be treated by type checkers similarly to classes created with [`@dataclasses.dataclass`](https://docs.python.org/3/library/dataclasses.html#dataclasses.dataclass "dataclasses.dataclass"). For example, type checkers will assume these classes have `__init__` methods that accept `id` and `name`.
The decorated class, metaclass, or function may accept the following bool arguments which type checkers will assume have the same effect as they would have on the [`@dataclasses.dataclass`](https://docs.python.org/3/library/dataclasses.html#dataclasses.dataclass "dataclasses.dataclass") decorator: `init`, `eq`, `order`, `unsafe_hash`, `frozen`, `match_args`, `kw_only`, and `slots`. It must be possible for the value of these arguments (`True` or `False`) to be statically evaluated.
The arguments to the `dataclass_transform` decorator can be used to customize the default behaviors of the decorated class, metaclass, or function: 

Parameters:
    
  * **eq_default** ([_bool_](https://docs.python.org/3/library/functions.html#bool "bool")) – Indicates whether the `eq` parameter is assumed to be `True` or `False` if it is omitted by the caller. Defaults to `True`.
  * **order_default** ([_bool_](https://docs.python.org/3/library/functions.html#bool "bool")) – Indicates whether the `order` parameter is assumed to be `True` or `False` if it is omitted by the caller. Defaults to `False`.
  * **kw_only_default** ([_bool_](https://docs.python.org/3/library/functions.html#bool "bool")) – Indicates whether the `kw_only` parameter is assumed to be `True` or `False` if it is omitted by the caller. Defaults to `False`.
  * **frozen_default** ([_bool_](https://docs.python.org/3/library/functions.html#bool "bool")) – 
Indicates whether the `frozen` parameter is assumed to be `True` or `False` if it is omitted by the caller. Defaults to `False`.
Added in version 3.12.
  * **field_specifiers** ([_tuple_](https://docs.python.org/3/library/stdtypes.html#tuple "tuple") _[_[_Callable_](https://docs.python.org/3/library/collections.abc.html#collections.abc.Callable "collections.abc.Callable") _[__...__,__Any_ _]__,__...__]_) – Specifies a static list of supported classes or functions that describe fields, similar to [`dataclasses.field()`](https://docs.python.org/3/library/dataclasses.html#dataclasses.field "dataclasses.field"). Defaults to `()`.
  * ****kwargs** (_Any_) – Arbitrary other keyword arguments are accepted in order to allow for possible future extensions.


Type checkers recognize the following optional parameters on field specifiers:
**Recognised parameters for field specifiers**[¶](https://docs.python.org/3/library/typing.html#id7 "Link to this table") Parameter name | Description  
---|---  
`init` | Indicates whether the field should be included in the synthesized `__init__` method. If unspecified, `init` defaults to `True`.  
`default` | Provides the default value for the field.  
`default_factory` | Provides a runtime callback that returns the default value for the field. If neither `default` nor `default_factory` are specified, the field is assumed to have no default value and must be provided a value when the class is instantiated.  
`factory` | An alias for the `default_factory` parameter on field specifiers.  
`kw_only` | Indicates whether the field should be marked as keyword-only. If `True`, the field will be keyword-only. If `False`, it will not be keyword-only. If unspecified, the value of the `kw_only` parameter on the object decorated with `dataclass_transform` will be used, or if that is unspecified, the value of `kw_only_default` on `dataclass_transform` will be used.  
`alias` | Provides an alternative name for the field. This alternative name is used in the synthesized `__init__` method.  
At runtime, this decorator records its arguments in the `__dataclass_transform__` attribute on the decorated object. It has no other runtime effect.
See [**PEP 681**](https://peps.python.org/pep-0681/) for more details.
Added in version 3.11. 

@typing.overload[¶](https://docs.python.org/3/library/typing.html#typing.overload "Link to this definition")
    
Decorator for creating overloaded functions and methods.
The `@overload` decorator allows describing functions and methods that support multiple different combinations of argument types. A series of `@overload`-decorated definitions must be followed by exactly one non-`@overload`-decorated definition (for the same function/method).
`@overload`-decorated definitions are for the benefit of the type checker only, since they will be overwritten by the non-`@overload`-decorated definition. The non-`@overload`-decorated definition, meanwhile, will be used at runtime but should be ignored by a type checker. At runtime, calling an `@overload`-decorated function directly will raise [`NotImplementedError`](https://docs.python.org/3/library/exceptions.html#NotImplementedError "NotImplementedError").
An example of overload that gives a more precise type than can be expressed using a union or a type variable:
Copy```
@overload
defprocess(response: None) -> None:
  ...
@overload
defprocess(response: int) -> tuple[int, str]:
  ...
@overload
defprocess(response: bytes) -> str:
  ...
defprocess(response):
  ... # actual implementation goes here

```

See [**PEP 484**](https://peps.python.org/pep-0484/) for more details and comparison with other typing semantics.
Changed in version 3.11: Overloaded functions can now be introspected at runtime using [`get_overloads()`](https://docs.python.org/3/library/typing.html#typing.get_overloads "typing.get_overloads"). 

typing.get_overloads(_func_)[¶](https://docs.python.org/3/library/typing.html#typing.get_overloads "Link to this definition")
    
Return a sequence of [`@overload`](https://docs.python.org/3/library/typing.html#typing.overload "typing.overload")-decorated definitions for _func_.
_func_ is the function object for the implementation of the overloaded function. For example, given the definition of `process` in the documentation for [`@overload`](https://docs.python.org/3/library/typing.html#typing.overload "typing.overload"), `get_overloads(process)` will return a sequence of three function objects for the three defined overloads. If called on a function with no overloads, `get_overloads()` returns an empty sequence.
`get_overloads()` can be used for introspecting an overloaded function at runtime.
Added in version 3.11. 

typing.clear_overloads()[¶](https://docs.python.org/3/library/typing.html#typing.clear_overloads "Link to this definition")
    
Clear all registered overloads in the internal registry.
This can be used to reclaim the memory used by the registry.
Added in version 3.11. 

@typing.final[¶](https://docs.python.org/3/library/typing.html#typing.final "Link to this definition")
    
Decorator to indicate final methods and final classes.
Decorating a method with `@final` indicates to a type checker that the method cannot be overridden in a subclass. Decorating a class with `@final` indicates that it cannot be subclassed.
For example:
Copy```
classBase:
  @final
  defdone(self) -> None:
    ...
classSub(Base):
  defdone(self) -> None: # Error reported by type checker
    ...
@final
classLeaf:
  ...
classOther(Leaf): # Error reported by type checker
  ...

```

There is no runtime checking of these properties. See [**PEP 591**](https://peps.python.org/pep-0591/) for more details.
Added in version 3.8.
Changed in version 3.11: The decorator will now attempt to set a `__final__` attribute to `True` on the decorated object. Thus, a check like `if getattr(obj, "__final__", False)` can be used at runtime to determine whether an object `obj` has been marked as final. If the decorated object does not support setting attributes, the decorator returns the object unchanged without raising an exception. 

@typing.no_type_check[¶](https://docs.python.org/3/library/typing.html#typing.no_type_check "Link to this definition")
    
Decorator to indicate that annotations are not type hints.
This works as a class or function [decorator](https://docs.python.org/3/glossary.html#term-decorator). With a class, it applies recursively to all methods and classes defined in that class (but not to methods defined in its superclasses or subclasses). Type checkers will ignore all annotations in a function or class with this decorator.
`@no_type_check` mutates the decorated object in place. 

@typing.no_type_check_decorator[¶](https://docs.python.org/3/library/typing.html#typing.no_type_check_decorator "Link to this definition")
    
Decorator to give another decorator the [`no_type_check()`](https://docs.python.org/3/library/typing.html#typing.no_type_check "typing.no_type_check") effect.
This wraps the decorator with something that wraps the decorated function in [`no_type_check()`](https://docs.python.org/3/library/typing.html#typing.no_type_check "typing.no_type_check").
Deprecated since version 3.13, will be removed in version 3.15: No type checker ever added support for `@no_type_check_decorator`. It is therefore deprecated, and will be removed in Python 3.15. 

@typing.override[¶](https://docs.python.org/3/library/typing.html#typing.override "Link to this definition")
    
Decorator to indicate that a method in a subclass is intended to override a method or attribute in a superclass.
Type checkers should emit an error if a method decorated with `@override` does not, in fact, override anything. This helps prevent bugs that may occur when a base class is changed without an equivalent change to a child class.
For example:
Copy```
classBase:
  deflog_status(self) -> None:
    ...
classSub(Base):
  @override
  deflog_status(self) -> None: # Okay: overrides Base.log_status
    ...
  @override
  defdone(self) -> None: # Error reported by type checker
    ...

```

There is no runtime checking of this property.
The decorator will attempt to set an `__override__` attribute to `True` on the decorated object. Thus, a check like `if getattr(obj, "__override__", False)` can be used at runtime to determine whether an object `obj` has been marked as an override. If the decorated object does not support setting attributes, the decorator returns the object unchanged without raising an exception.
See [**PEP 698**](https://peps.python.org/pep-0698/) for more details.
Added in version 3.12. 

@typing.type_check_only[¶](https://docs.python.org/3/library/typing.html#typing.type_check_only "Link to this definition")
    
Decorator to mark a class or function as unavailable at runtime.
This decorator is itself not available at runtime. It is mainly intended to mark classes that are defined in type stub files if an implementation returns an instance of a private class:
Copy```
@type_check_only
classResponse: # private or not available at runtime
  code: int
  defget_header(self, name: str) -> str: ...
deffetch_response() -> Response: ...

```

Note that returning instances of private classes is not recommended. It is usually preferable to make such classes public.
### Introspection helpers[¶](https://docs.python.org/3/library/typing.html#introspection-helpers "Link to this heading") 

typing.get_type_hints(_obj_ , _globalns =None_, _localns =None_, _include_extras =False_)[¶](https://docs.python.org/3/library/typing.html#typing.get_type_hints "Link to this definition")
    
Return a dictionary containing type hints for a function, method, module or class object.
This is often the same as `obj.__annotations__`, but this function makes the following changes to the annotations dictionary:
  * Forward references encoded as string literals or [`ForwardRef`](https://docs.python.org/3/library/typing.html#typing.ForwardRef "typing.ForwardRef") objects are handled by evaluating them in _globalns_ , _localns_ , and (where applicable) _obj_ ’s [type parameter](https://docs.python.org/3/reference/compound_stmts.html#type-params) namespace. If _globalns_ or _localns_ is not given, appropriate namespace dictionaries are inferred from _obj_.
  * `None` is replaced with [`types.NoneType`](https://docs.python.org/3/library/types.html#types.NoneType "types.NoneType").
  * If [`@no_type_check`](https://docs.python.org/3/library/typing.html#typing.no_type_check "typing.no_type_check") has been applied to _obj_ , an empty dictionary is returned.
  * If _obj_ is a class `C`, the function returns a dictionary that merges annotations from `C`’s base classes with those on `C` directly. This is done by traversing [`C.__mro__`](https://docs.python.org/3/reference/datamodel.html#type.__mro__ "type.__mro__") and iteratively combining `__annotations__` dictionaries. Annotations on classes appearing earlier in the [method resolution order](https://docs.python.org/3/glossary.html#term-method-resolution-order) always take precedence over annotations on classes appearing later in the method resolution order.
  * The function recursively replaces all occurrences of `Annotated[T, ...]` with `T`, unless _include_extras_ is set to `True` (see [`Annotated`](https://docs.python.org/3/library/typing.html#typing.Annotated "typing.Annotated") for more information).


See also [`inspect.get_annotations()`](https://docs.python.org/3/library/inspect.html#inspect.get_annotations "inspect.get_annotations"), a lower-level function that returns annotations more directly.
Note
If any forward references in the annotations of _obj_ are not resolvable or are not valid Python code, this function will raise an exception such as [`NameError`](https://docs.python.org/3/library/exceptions.html#NameError "NameError"). For example, this can happen with imported [type aliases](https://docs.python.org/3/library/typing.html#type-aliases) that include forward references, or with names imported under [`if TYPE_CHECKING`](https://docs.python.org/3/library/typing.html#typing.TYPE_CHECKING "typing.TYPE_CHECKING").
Changed in version 3.9: Added `include_extras` parameter as part of [**PEP 593**](https://peps.python.org/pep-0593/). See the documentation on [`Annotated`](https://docs.python.org/3/library/typing.html#typing.Annotated "typing.Annotated") for more information.
Changed in version 3.11: Previously, `Optional[t]` was added for function and method annotations if a default value equal to `None` was set. Now the annotation is returned unchanged. 

typing.get_origin(_tp_)[¶](https://docs.python.org/3/library/typing.html#typing.get_origin "Link to this definition")
    
Get the unsubscripted version of a type: for a typing object of the form `X[Y, Z, ...]` return `X`.
If `X` is a typing-module alias for a builtin or [`collections`](https://docs.python.org/3/library/collections.html#module-collections "collections: Container datatypes") class, it will be normalized to the original class. If `X` is an instance of [`ParamSpecArgs`](https://docs.python.org/3/library/typing.html#typing.ParamSpecArgs "typing.ParamSpecArgs") or [`ParamSpecKwargs`](https://docs.python.org/3/library/typing.html#typing.ParamSpecKwargs "typing.ParamSpecKwargs"), return the underlying [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec"). Return `None` for unsupported objects.
Examples:
Copy```
assert get_origin(str) is None
assert get_origin(Dict[str, int]) is dict
assert get_origin(Union[int, str]) is Union
assert get_origin(Annotated[str, "metadata"]) is Annotated
P = ParamSpec('P')
assert get_origin(P.args) is P
assert get_origin(P.kwargs) is P

```

Added in version 3.8. 

typing.get_args(_tp_)[¶](https://docs.python.org/3/library/typing.html#typing.get_args "Link to this definition")
    
Get type arguments with all substitutions performed: for a typing object of the form `X[Y, Z, ...]` return `(Y, Z, ...)`.
If `X` is a union or [`Literal`](https://docs.python.org/3/library/typing.html#typing.Literal "typing.Literal") contained in another generic type, the order of `(Y, Z, ...)` may be different from the order of the original arguments `[Y, Z, ...]` due to type caching. Return `()` for unsupported objects.
Examples:
Copy```
assert get_args(int) == ()
assert get_args(Dict[int, str]) == (int, str)
assert get_args(Union[int, str]) == (int, str)

```

Added in version 3.8. 

typing.get_protocol_members(_tp_)[¶](https://docs.python.org/3/library/typing.html#typing.get_protocol_members "Link to this definition")
    
Return the set of members defined in a [`Protocol`](https://docs.python.org/3/library/typing.html#typing.Protocol "typing.Protocol").
Copy```
>>> fromtypingimport Protocol, get_protocol_members
>>> classP(Protocol):
...   defa(self) -> str: ...
...   b: int
>>> get_protocol_members(P) == frozenset({'a', 'b'})
True

```

Raise [`TypeError`](https://docs.python.org/3/library/exceptions.html#TypeError "TypeError") for arguments that are not Protocols.
Added in version 3.13. 

typing.is_protocol(_tp_)[¶](https://docs.python.org/3/library/typing.html#typing.is_protocol "Link to this definition")
    
Determine if a type is a [`Protocol`](https://docs.python.org/3/library/typing.html#typing.Protocol "typing.Protocol").
For example:
Copy```
classP(Protocol):
  defa(self) -> str: ...
  b: int
is_protocol(P)  # => True
is_protocol(int) # => False

```

Added in version 3.13. 

typing.is_typeddict(_tp_)[¶](https://docs.python.org/3/library/typing.html#typing.is_typeddict "Link to this definition")
    
Check if a type is a [`TypedDict`](https://docs.python.org/3/library/typing.html#typing.TypedDict "typing.TypedDict").
For example:
Copy```
classFilm(TypedDict):
  title: str
  year: int
assert is_typeddict(Film)
assert not is_typeddict(list | str)
# TypedDict is a factory for creating typed dicts,
# not a typed dict itself
assert not is_typeddict(TypedDict)

```

Added in version 3.10. 

_class_ typing.ForwardRef[¶](https://docs.python.org/3/library/typing.html#typing.ForwardRef "Link to this definition")
    
Class used for internal typing representation of string forward references.
For example, `List["SomeClass"]` is implicitly transformed into `List[ForwardRef("SomeClass")]`. `ForwardRef` should not be instantiated by a user, but may be used by introspection tools.
Note
[**PEP 585**](https://peps.python.org/pep-0585/) generic types such as `list["SomeClass"]` will not be implicitly transformed into `list[ForwardRef("SomeClass")]` and thus will not automatically resolve to `list[SomeClass]`.
Added in version 3.7.4. 

typing.NoDefault[¶](https://docs.python.org/3/library/typing.html#typing.NoDefault "Link to this definition")
    
A sentinel object used to indicate that a type parameter has no default value. For example:
Copy```
>>> T = TypeVar("T")
>>> T.__default__ is typing.NoDefault
True
>>> S = TypeVar("S", default=None)
>>> S.__default__ is None
True

```

Added in version 3.13.
### Constant[¶](https://docs.python.org/3/library/typing.html#constant "Link to this heading") 

typing.TYPE_CHECKING[¶](https://docs.python.org/3/library/typing.html#typing.TYPE_CHECKING "Link to this definition")
    
A special constant that is assumed to be `True` by 3rd party static type checkers. It is `False` at runtime.
Usage:
Copy```
if TYPE_CHECKING:
  importexpensive_mod
deffun(arg: 'expensive_mod.SomeType') -> None:
  local_var: expensive_mod.AnotherType = other_fun()

```

The first type annotation must be enclosed in quotes, making it a “forward reference”, to hide the `expensive_mod` reference from the interpreter runtime. Type annotations for local variables are not evaluated, so the second annotation does not need to be enclosed in quotes.
Note
If `from __future__ import annotations` is used, annotations are not evaluated at function definition time. Instead, they are stored as strings in `__annotations__`. This makes it unnecessary to use quotes around the annotation (see [**PEP 563**](https://peps.python.org/pep-0563/)).
Added in version 3.5.2.
### Deprecated aliases[¶](https://docs.python.org/3/library/typing.html#deprecated-aliases "Link to this heading")
This module defines several deprecated aliases to pre-existing standard library classes. These were originally included in the `typing` module in order to support parameterizing these generic classes using `[]`. However, the aliases became redundant in Python 3.9 when the corresponding pre-existing classes were enhanced to support `[]` (see [**PEP 585**](https://peps.python.org/pep-0585/)).
The redundant types are deprecated as of Python 3.9. However, while the aliases may be removed at some point, removal of these aliases is not currently planned. As such, no deprecation warnings are currently issued by the interpreter for these aliases.
If at some point it is decided to remove these deprecated aliases, a deprecation warning will be issued by the interpreter for at least two releases prior to removal. The aliases are guaranteed to remain in the `typing` module without deprecation warnings until at least Python 3.14.
Type checkers are encouraged to flag uses of the deprecated types if the program they are checking targets a minimum Python version of 3.9 or newer.
#### Aliases to built-in types[¶](https://docs.python.org/3/library/typing.html#aliases-to-built-in-types "Link to this heading") 

_class_ typing.Dict(_dict, MutableMapping[KT, VT]_)[¶](https://docs.python.org/3/library/typing.html#typing.Dict "Link to this definition")
    
Deprecated alias to [`dict`](https://docs.python.org/3/library/stdtypes.html#dict "dict").
Note that to annotate arguments, it is preferred to use an abstract collection type such as [`Mapping`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Mapping "collections.abc.Mapping") rather than to use [`dict`](https://docs.python.org/3/library/stdtypes.html#dict "dict") or `typing.Dict`.
Deprecated since version 3.9: [`builtins.dict`](https://docs.python.org/3/library/stdtypes.html#dict "dict") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.List(_list, MutableSequence[T]_)[¶](https://docs.python.org/3/library/typing.html#typing.List "Link to this definition")
    
Deprecated alias to [`list`](https://docs.python.org/3/library/stdtypes.html#list "list").
Note that to annotate arguments, it is preferred to use an abstract collection type such as [`Sequence`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Sequence "collections.abc.Sequence") or [`Iterable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Iterable "collections.abc.Iterable") rather than to use [`list`](https://docs.python.org/3/library/stdtypes.html#list "list") or `typing.List`.
Deprecated since version 3.9: [`builtins.list`](https://docs.python.org/3/library/stdtypes.html#list "list") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Set(_set, MutableSet[T]_)[¶](https://docs.python.org/3/library/typing.html#typing.Set "Link to this definition")
    
Deprecated alias to [`builtins.set`](https://docs.python.org/3/library/stdtypes.html#set "set").
Note that to annotate arguments, it is preferred to use an abstract collection type such as [`collections.abc.Set`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Set "collections.abc.Set") rather than to use [`set`](https://docs.python.org/3/library/stdtypes.html#set "set") or [`typing.Set`](https://docs.python.org/3/library/typing.html#typing.Set "typing.Set").
Deprecated since version 3.9: [`builtins.set`](https://docs.python.org/3/library/stdtypes.html#set "set") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.FrozenSet(_frozenset, AbstractSet[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.FrozenSet "Link to this definition")
    
Deprecated alias to [`builtins.frozenset`](https://docs.python.org/3/library/stdtypes.html#frozenset "frozenset").
Deprecated since version 3.9: [`builtins.frozenset`](https://docs.python.org/3/library/stdtypes.html#frozenset "frozenset") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

typing.Tuple[¶](https://docs.python.org/3/library/typing.html#typing.Tuple "Link to this definition")
    
Deprecated alias for [`tuple`](https://docs.python.org/3/library/stdtypes.html#tuple "tuple").
[`tuple`](https://docs.python.org/3/library/stdtypes.html#tuple "tuple") and `Tuple` are special-cased in the type system; see [Annotating tuples](https://docs.python.org/3/library/typing.html#annotating-tuples) for more details.
Deprecated since version 3.9: [`builtins.tuple`](https://docs.python.org/3/library/stdtypes.html#tuple "tuple") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Type(_Generic[CT_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.Type "Link to this definition")
    
Deprecated alias to [`type`](https://docs.python.org/3/library/functions.html#type "type").
See [The type of class objects](https://docs.python.org/3/library/typing.html#type-of-class-objects) for details on using [`type`](https://docs.python.org/3/library/functions.html#type "type") or `typing.Type` in type annotations.
Added in version 3.5.2.
Deprecated since version 3.9: [`builtins.type`](https://docs.python.org/3/library/functions.html#type "type") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias).
#### Aliases to types in [`collections`](https://docs.python.org/3/library/collections.html#module-collections "collections: Container datatypes")[¶](https://docs.python.org/3/library/typing.html#aliases-to-types-in-collections "Link to this heading") 

_class_ typing.DefaultDict(_collections.defaultdict, MutableMapping[KT, VT]_)[¶](https://docs.python.org/3/library/typing.html#typing.DefaultDict "Link to this definition")
    
Deprecated alias to [`collections.defaultdict`](https://docs.python.org/3/library/collections.html#collections.defaultdict "collections.defaultdict").
Added in version 3.5.2.
Deprecated since version 3.9: [`collections.defaultdict`](https://docs.python.org/3/library/collections.html#collections.defaultdict "collections.defaultdict") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.OrderedDict(_collections.OrderedDict, MutableMapping[KT, VT]_)[¶](https://docs.python.org/3/library/typing.html#typing.OrderedDict "Link to this definition")
    
Deprecated alias to [`collections.OrderedDict`](https://docs.python.org/3/library/collections.html#collections.OrderedDict "collections.OrderedDict").
Added in version 3.7.2.
Deprecated since version 3.9: [`collections.OrderedDict`](https://docs.python.org/3/library/collections.html#collections.OrderedDict "collections.OrderedDict") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.ChainMap(_collections.ChainMap, MutableMapping[KT, VT]_)[¶](https://docs.python.org/3/library/typing.html#typing.ChainMap "Link to this definition")
    
Deprecated alias to [`collections.ChainMap`](https://docs.python.org/3/library/collections.html#collections.ChainMap "collections.ChainMap").
Added in version 3.6.1.
Deprecated since version 3.9: [`collections.ChainMap`](https://docs.python.org/3/library/collections.html#collections.ChainMap "collections.ChainMap") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Counter(_collections.Counter, Dict[T, int]_)[¶](https://docs.python.org/3/library/typing.html#typing.Counter "Link to this definition")
    
Deprecated alias to [`collections.Counter`](https://docs.python.org/3/library/collections.html#collections.Counter "collections.Counter").
Added in version 3.6.1.
Deprecated since version 3.9: [`collections.Counter`](https://docs.python.org/3/library/collections.html#collections.Counter "collections.Counter") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Deque(_deque, MutableSequence[T]_)[¶](https://docs.python.org/3/library/typing.html#typing.Deque "Link to this definition")
    
Deprecated alias to [`collections.deque`](https://docs.python.org/3/library/collections.html#collections.deque "collections.deque").
Added in version 3.6.1.
Deprecated since version 3.9: [`collections.deque`](https://docs.python.org/3/library/collections.html#collections.deque "collections.deque") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias).
#### Aliases to other concrete types[¶](https://docs.python.org/3/library/typing.html#aliases-to-other-concrete-types "Link to this heading") 

_class_ typing.Pattern[¶](https://docs.python.org/3/library/typing.html#typing.Pattern "Link to this definition")


_class_ typing.Match[¶](https://docs.python.org/3/library/typing.html#typing.Match "Link to this definition")
    
Deprecated aliases corresponding to the return types from [`re.compile()`](https://docs.python.org/3/library/re.html#re.compile "re.compile") and [`re.match()`](https://docs.python.org/3/library/re.html#re.match "re.match").
These types (and the corresponding functions) are generic over [`AnyStr`](https://docs.python.org/3/library/typing.html#typing.AnyStr "typing.AnyStr"). `Pattern` can be specialised as `Pattern[str]` or `Pattern[bytes]`; `Match` can be specialised as `Match[str]` or `Match[bytes]`.
Deprecated since version 3.9: Classes `Pattern` and `Match` from [`re`](https://docs.python.org/3/library/re.html#module-re "re: Regular expression operations.") now support `[]`. See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Text[¶](https://docs.python.org/3/library/typing.html#typing.Text "Link to this definition")
    
Deprecated alias for [`str`](https://docs.python.org/3/library/stdtypes.html#str "str").
`Text` is provided to supply a forward compatible path for Python 2 code: in Python 2, `Text` is an alias for `unicode`.
Use `Text` to indicate that a value must contain a unicode string in a manner that is compatible with both Python 2 and Python 3:
Copy```
defadd_unicode_checkmark(text: Text) -> Text:
  return text + u' \u2713'

```

Added in version 3.5.2.
Deprecated since version 3.11: Python 2 is no longer supported, and most type checkers also no longer support type checking Python 2 code. Removal of the alias is not currently planned, but users are encouraged to use [`str`](https://docs.python.org/3/library/stdtypes.html#str "str") instead of `Text`.
#### Aliases to container ABCs in [`collections.abc`](https://docs.python.org/3/library/collections.abc.html#module-collections.abc "collections.abc: Abstract base classes for containers")[¶](https://docs.python.org/3/library/typing.html#aliases-to-container-abcs-in-collections-abc "Link to this heading") 

_class_ typing.AbstractSet(_Collection[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.AbstractSet "Link to this definition")
    
Deprecated alias to [`collections.abc.Set`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Set "collections.abc.Set").
Deprecated since version 3.9: [`collections.abc.Set`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Set "collections.abc.Set") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.ByteString(_Sequence[int]_)[¶](https://docs.python.org/3/library/typing.html#typing.ByteString "Link to this definition")
    
This type represents the types [`bytes`](https://docs.python.org/3/library/stdtypes.html#bytes "bytes"), [`bytearray`](https://docs.python.org/3/library/stdtypes.html#bytearray "bytearray"), and [`memoryview`](https://docs.python.org/3/library/stdtypes.html#memoryview "memoryview") of byte sequences.
Deprecated since version 3.9, will be removed in version 3.14: Prefer [`collections.abc.Buffer`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Buffer "collections.abc.Buffer"), or a union like `bytes | bytearray | memoryview`. 

_class_ typing.Collection(_Sized, Iterable[T_co], Container[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.Collection "Link to this definition")
    
Deprecated alias to [`collections.abc.Collection`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Collection "collections.abc.Collection").
Added in version 3.6.
Deprecated since version 3.9: [`collections.abc.Collection`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Collection "collections.abc.Collection") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Container(_Generic[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.Container "Link to this definition")
    
Deprecated alias to [`collections.abc.Container`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Container "collections.abc.Container").
Deprecated since version 3.9: [`collections.abc.Container`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Container "collections.abc.Container") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.ItemsView(_MappingView, AbstractSet[tuple[KT_co, VT_co]]_)[¶](https://docs.python.org/3/library/typing.html#typing.ItemsView "Link to this definition")
    
Deprecated alias to [`collections.abc.ItemsView`](https://docs.python.org/3/library/collections.abc.html#collections.abc.ItemsView "collections.abc.ItemsView").
Deprecated since version 3.9: [`collections.abc.ItemsView`](https://docs.python.org/3/library/collections.abc.html#collections.abc.ItemsView "collections.abc.ItemsView") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.KeysView(_MappingView, AbstractSet[KT_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.KeysView "Link to this definition")
    
Deprecated alias to [`collections.abc.KeysView`](https://docs.python.org/3/library/collections.abc.html#collections.abc.KeysView "collections.abc.KeysView").
Deprecated since version 3.9: [`collections.abc.KeysView`](https://docs.python.org/3/library/collections.abc.html#collections.abc.KeysView "collections.abc.KeysView") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Mapping(_Collection[KT], Generic[KT, VT_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.Mapping "Link to this definition")
    
Deprecated alias to [`collections.abc.Mapping`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Mapping "collections.abc.Mapping").
Deprecated since version 3.9: [`collections.abc.Mapping`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Mapping "collections.abc.Mapping") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.MappingView(_Sized_)[¶](https://docs.python.org/3/library/typing.html#typing.MappingView "Link to this definition")
    
Deprecated alias to [`collections.abc.MappingView`](https://docs.python.org/3/library/collections.abc.html#collections.abc.MappingView "collections.abc.MappingView").
Deprecated since version 3.9: [`collections.abc.MappingView`](https://docs.python.org/3/library/collections.abc.html#collections.abc.MappingView "collections.abc.MappingView") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.MutableMapping(_Mapping[KT, VT]_)[¶](https://docs.python.org/3/library/typing.html#typing.MutableMapping "Link to this definition")
    
Deprecated alias to [`collections.abc.MutableMapping`](https://docs.python.org/3/library/collections.abc.html#collections.abc.MutableMapping "collections.abc.MutableMapping").
Deprecated since version 3.9: [`collections.abc.MutableMapping`](https://docs.python.org/3/library/collections.abc.html#collections.abc.MutableMapping "collections.abc.MutableMapping") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.MutableSequence(_Sequence[T]_)[¶](https://docs.python.org/3/library/typing.html#typing.MutableSequence "Link to this definition")
    
Deprecated alias to [`collections.abc.MutableSequence`](https://docs.python.org/3/library/collections.abc.html#collections.abc.MutableSequence "collections.abc.MutableSequence").
Deprecated since version 3.9: [`collections.abc.MutableSequence`](https://docs.python.org/3/library/collections.abc.html#collections.abc.MutableSequence "collections.abc.MutableSequence") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.MutableSet(_AbstractSet[T]_)[¶](https://docs.python.org/3/library/typing.html#typing.MutableSet "Link to this definition")
    
Deprecated alias to [`collections.abc.MutableSet`](https://docs.python.org/3/library/collections.abc.html#collections.abc.MutableSet "collections.abc.MutableSet").
Deprecated since version 3.9: [`collections.abc.MutableSet`](https://docs.python.org/3/library/collections.abc.html#collections.abc.MutableSet "collections.abc.MutableSet") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Sequence(_Reversible[T_co], Collection[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.Sequence "Link to this definition")
    
Deprecated alias to [`collections.abc.Sequence`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Sequence "collections.abc.Sequence").
Deprecated since version 3.9: [`collections.abc.Sequence`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Sequence "collections.abc.Sequence") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.ValuesView(_MappingView, Collection[_VT_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.ValuesView "Link to this definition")
    
Deprecated alias to [`collections.abc.ValuesView`](https://docs.python.org/3/library/collections.abc.html#collections.abc.ValuesView "collections.abc.ValuesView").
Deprecated since version 3.9: [`collections.abc.ValuesView`](https://docs.python.org/3/library/collections.abc.html#collections.abc.ValuesView "collections.abc.ValuesView") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias).
#### Aliases to asynchronous ABCs in [`collections.abc`](https://docs.python.org/3/library/collections.abc.html#module-collections.abc "collections.abc: Abstract base classes for containers")[¶](https://docs.python.org/3/library/typing.html#aliases-to-asynchronous-abcs-in-collections-abc "Link to this heading") 

_class_ typing.Coroutine(_Awaitable[ReturnType], Generic[YieldType, SendType, ReturnType]_)[¶](https://docs.python.org/3/library/typing.html#typing.Coroutine "Link to this definition")
    
Deprecated alias to [`collections.abc.Coroutine`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Coroutine "collections.abc.Coroutine").
See [Annotating generators and coroutines](https://docs.python.org/3/library/typing.html#annotating-generators-and-coroutines) for details on using [`collections.abc.Coroutine`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Coroutine "collections.abc.Coroutine") and `typing.Coroutine` in type annotations.
Added in version 3.5.3.
Deprecated since version 3.9: [`collections.abc.Coroutine`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Coroutine "collections.abc.Coroutine") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.AsyncGenerator(_AsyncIterator[YieldType], Generic[YieldType, SendType]_)[¶](https://docs.python.org/3/library/typing.html#typing.AsyncGenerator "Link to this definition")
    
Deprecated alias to [`collections.abc.AsyncGenerator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.AsyncGenerator "collections.abc.AsyncGenerator").
See [Annotating generators and coroutines](https://docs.python.org/3/library/typing.html#annotating-generators-and-coroutines) for details on using [`collections.abc.AsyncGenerator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.AsyncGenerator "collections.abc.AsyncGenerator") and `typing.AsyncGenerator` in type annotations.
Added in version 3.6.1.
Deprecated since version 3.9: [`collections.abc.AsyncGenerator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.AsyncGenerator "collections.abc.AsyncGenerator") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias).
Changed in version 3.13: The `SendType` parameter now has a default. 

_class_ typing.AsyncIterable(_Generic[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.AsyncIterable "Link to this definition")
    
Deprecated alias to [`collections.abc.AsyncIterable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.AsyncIterable "collections.abc.AsyncIterable").
Added in version 3.5.2.
Deprecated since version 3.9: [`collections.abc.AsyncIterable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.AsyncIterable "collections.abc.AsyncIterable") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.AsyncIterator(_AsyncIterable[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.AsyncIterator "Link to this definition")
    
Deprecated alias to [`collections.abc.AsyncIterator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.AsyncIterator "collections.abc.AsyncIterator").
Added in version 3.5.2.
Deprecated since version 3.9: [`collections.abc.AsyncIterator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.AsyncIterator "collections.abc.AsyncIterator") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Awaitable(_Generic[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.Awaitable "Link to this definition")
    
Deprecated alias to [`collections.abc.Awaitable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Awaitable "collections.abc.Awaitable").
Added in version 3.5.2.
Deprecated since version 3.9: [`collections.abc.Awaitable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Awaitable "collections.abc.Awaitable") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias).
#### Aliases to other ABCs in [`collections.abc`](https://docs.python.org/3/library/collections.abc.html#module-collections.abc "collections.abc: Abstract base classes for containers")[¶](https://docs.python.org/3/library/typing.html#aliases-to-other-abcs-in-collections-abc "Link to this heading") 

_class_ typing.Iterable(_Generic[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.Iterable "Link to this definition")
    
Deprecated alias to [`collections.abc.Iterable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Iterable "collections.abc.Iterable").
Deprecated since version 3.9: [`collections.abc.Iterable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Iterable "collections.abc.Iterable") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Iterator(_Iterable[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.Iterator "Link to this definition")
    
Deprecated alias to [`collections.abc.Iterator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Iterator "collections.abc.Iterator").
Deprecated since version 3.9: [`collections.abc.Iterator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Iterator "collections.abc.Iterator") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

typing.Callable[¶](https://docs.python.org/3/library/typing.html#typing.Callable "Link to this definition")
    
Deprecated alias to [`collections.abc.Callable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Callable "collections.abc.Callable").
See [Annotating callable objects](https://docs.python.org/3/library/typing.html#annotating-callables) for details on how to use [`collections.abc.Callable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Callable "collections.abc.Callable") and `typing.Callable` in type annotations.
Deprecated since version 3.9: [`collections.abc.Callable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Callable "collections.abc.Callable") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias).
Changed in version 3.10: `Callable` now supports [`ParamSpec`](https://docs.python.org/3/library/typing.html#typing.ParamSpec "typing.ParamSpec") and [`Concatenate`](https://docs.python.org/3/library/typing.html#typing.Concatenate "typing.Concatenate"). See [**PEP 612**](https://peps.python.org/pep-0612/) for more details. 

_class_ typing.Generator(_Iterator[YieldType], Generic[YieldType, SendType, ReturnType]_)[¶](https://docs.python.org/3/library/typing.html#typing.Generator "Link to this definition")
    
Deprecated alias to [`collections.abc.Generator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Generator "collections.abc.Generator").
See [Annotating generators and coroutines](https://docs.python.org/3/library/typing.html#annotating-generators-and-coroutines) for details on using [`collections.abc.Generator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Generator "collections.abc.Generator") and `typing.Generator` in type annotations.
Deprecated since version 3.9: [`collections.abc.Generator`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Generator "collections.abc.Generator") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias).
Changed in version 3.13: Default values for the send and return types were added. 

_class_ typing.Hashable[¶](https://docs.python.org/3/library/typing.html#typing.Hashable "Link to this definition")
    
Deprecated alias to [`collections.abc.Hashable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Hashable "collections.abc.Hashable").
Deprecated since version 3.12: Use [`collections.abc.Hashable`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Hashable "collections.abc.Hashable") directly instead. 

_class_ typing.Reversible(_Iterable[T_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.Reversible "Link to this definition")
    
Deprecated alias to [`collections.abc.Reversible`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Reversible "collections.abc.Reversible").
Deprecated since version 3.9: [`collections.abc.Reversible`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Reversible "collections.abc.Reversible") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias). 

_class_ typing.Sized[¶](https://docs.python.org/3/library/typing.html#typing.Sized "Link to this definition")
    
Deprecated alias to [`collections.abc.Sized`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Sized "collections.abc.Sized").
Deprecated since version 3.12: Use [`collections.abc.Sized`](https://docs.python.org/3/library/collections.abc.html#collections.abc.Sized "collections.abc.Sized") directly instead.
#### Aliases to [`contextlib`](https://docs.python.org/3/library/contextlib.html#module-contextlib "contextlib: Utilities for with-statement contexts.") ABCs[¶](https://docs.python.org/3/library/typing.html#aliases-to-contextlib-abcs "Link to this heading") 

_class_ typing.ContextManager(_Generic[T_co, ExitT_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.ContextManager "Link to this definition")
    
Deprecated alias to [`contextlib.AbstractContextManager`](https://docs.python.org/3/library/contextlib.html#contextlib.AbstractContextManager "contextlib.AbstractContextManager").
The first type parameter, `T_co`, represents the type returned by the [`__enter__()`](https://docs.python.org/3/reference/datamodel.html#object.__enter__ "object.__enter__") method. The optional second type parameter, `ExitT_co`, which defaults to `bool | None`, represents the type returned by the [`__exit__()`](https://docs.python.org/3/reference/datamodel.html#object.__exit__ "object.__exit__") method.
Added in version 3.5.4.
Deprecated since version 3.9: [`contextlib.AbstractContextManager`](https://docs.python.org/3/library/contextlib.html#contextlib.AbstractContextManager "contextlib.AbstractContextManager") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias).
Changed in version 3.13: Added the optional second type parameter, `ExitT_co`. 

_class_ typing.AsyncContextManager(_Generic[T_co, AExitT_co]_)[¶](https://docs.python.org/3/library/typing.html#typing.AsyncContextManager "Link to this definition")
    
Deprecated alias to [`contextlib.AbstractAsyncContextManager`](https://docs.python.org/3/library/contextlib.html#contextlib.AbstractAsyncContextManager "contextlib.AbstractAsyncContextManager").
The first type parameter, `T_co`, represents the type returned by the [`__aenter__()`](https://docs.python.org/3/reference/datamodel.html#object.__aenter__ "object.__aenter__") method. The optional second type parameter, `AExitT_co`, which defaults to `bool | None`, represents the type returned by the [`__aexit__()`](https://docs.python.org/3/reference/datamodel.html#object.__aexit__ "object.__aexit__") method.
Added in version 3.6.2.
Deprecated since version 3.9: [`contextlib.AbstractAsyncContextManager`](https://docs.python.org/3/library/contextlib.html#contextlib.AbstractAsyncContextManager "contextlib.AbstractAsyncContextManager") now supports subscripting (`[]`). See [**PEP 585**](https://peps.python.org/pep-0585/) and [Generic Alias Type](https://docs.python.org/3/library/stdtypes.html#types-genericalias).
Changed in version 3.13: Added the optional second type parameter, `AExitT_co`.
## Deprecation Timeline of Major Features[¶](https://docs.python.org/3/library/typing.html#deprecation-timeline-of-major-features "Link to this heading")
Certain features in `typing` are deprecated and may be removed in a future version of Python. The following table summarizes major deprecations for your convenience. This is subject to change, and not all deprecations are listed.
Feature | Deprecated in | Projected removal | PEP/issue  
---|---|---|---  
`typing` versions of standard collections | 3.9 | Undecided (see [Deprecated aliases](https://docs.python.org/3/library/typing.html#deprecated-aliases) for more information) | [**PEP 585**](https://peps.python.org/pep-0585/)  
[`typing.ByteString`](https://docs.python.org/3/library/typing.html#typing.ByteString "typing.ByteString") | 3.9 | 3.14 | [gh-91896](https://github.com/python/cpython/issues/91896)  
[`typing.Text`](https://docs.python.org/3/library/typing.html#typing.Text "typing.Text") | 3.11 | Undecided | [gh-92332](https://github.com/python/cpython/issues/92332)  
[`typing.Hashable`](https://docs.python.org/3/library/typing.html#typing.Hashable "typing.Hashable") and [`typing.Sized`](https://docs.python.org/3/library/typing.html#typing.Sized "typing.Sized") | 3.12 | Undecided | [gh-94309](https://github.com/python/cpython/issues/94309)  
[`typing.TypeAlias`](https://docs.python.org/3/library/typing.html#typing.TypeAlias "typing.TypeAlias") | 3.12 | Undecided | [**PEP 695**](https://peps.python.org/pep-0695/)  
[`@typing.no_type_check_decorator`](https://docs.python.org/3/library/typing.html#typing.no_type_check_decorator "typing.no_type_check_decorator") | 3.13 | 3.15 | [gh-106309](https://github.com/python/cpython/issues/106309)  
[`typing.AnyStr`](https://docs.python.org/3/library/typing.html#typing.AnyStr "typing.AnyStr") | 3.13 | 3.18 | [gh-105578](https://github.com/python/cpython/issues/105578)  
### [Table of Contents](https://docs.python.org/3/contents.html)
  * [`typing` — Support for type hints](https://docs.python.org/3/library/typing.html)
    * [Specification for the Python Type System](https://docs.python.org/3/library/typing.html#specification-for-the-python-type-system)
    * [Type aliases](https://docs.python.org/3/library/typing.html#type-aliases)
    * [NewType](https://docs.python.org/3/library/typing.html#newtype)
    * [Annotating callable objects](https://docs.python.org/3/library/typing.html#annotating-callable-objects)
    * [Generics](https://docs.python.org/3/library/typing.html#generics)
    * [Annotating tuples](https://docs.python.org/3/library/typing.html#annotating-tuples)
    * [The type of class objects](https://docs.python.org/3/library/typing.html#the-type-of-class-objects)
    * [Annotating generators and coroutines](https://docs.python.org/3/library/typing.html#annotating-generators-and-coroutines)
    * [User-defined generic types](https://docs.python.org/3/library/typing.html#user-defined-generic-types)
    * [The `Any` type](https://docs.python.org/3/library/typing.html#the-any-type)
    * [Nominal vs structural subtyping](https://docs.python.org/3/library/typing.html#nominal-vs-structural-subtyping)
    * [Module contents](https://docs.python.org/3/library/typing.html#module-contents)
      * [Special typing primitives](https://docs.python.org/3/library/typing.html#special-typing-primitives)
        * [Special types](https://docs.python.org/3/library/typing.html#special-types)
        * [Special forms](https://docs.python.org/3/library/typing.html#special-forms)
        * [Building generic types and type aliases](https://docs.python.org/3/library/typing.html#building-generic-types-and-type-aliases)
        * [Other special directives](https://docs.python.org/3/library/typing.html#other-special-directives)
      * [Protocols](https://docs.python.org/3/library/typing.html#protocols)
      * [ABCs for working with IO](https://docs.python.org/3/library/typing.html#abcs-for-working-with-io)
      * [Functions and decorators](https://docs.python.org/3/library/typing.html#functions-and-decorators)
      * [Introspection helpers](https://docs.python.org/3/library/typing.html#introspection-helpers)
      * [Constant](https://docs.python.org/3/library/typing.html#constant)
      * [Deprecated aliases](https://docs.python.org/3/library/typing.html#deprecated-aliases)
        * [Aliases to built-in types](https://docs.python.org/3/library/typing.html#aliases-to-built-in-types)
        * [Aliases to types in `collections`](https://docs.python.org/3/library/typing.html#aliases-to-types-in-collections)
        * [Aliases to other concrete types](https://docs.python.org/3/library/typing.html#aliases-to-other-concrete-types)
        * [Aliases to container ABCs in `collections.abc`](https://docs.python.org/3/library/typing.html#aliases-to-container-abcs-in-collections-abc)
        * [Aliases to asynchronous ABCs in `collections.abc`](https://docs.python.org/3/library/typing.html#aliases-to-asynchronous-abcs-in-collections-abc)
        * [Aliases to other ABCs in `collections.abc`](https://docs.python.org/3/library/typing.html#aliases-to-other-abcs-in-collections-abc)
        * [Aliases to `contextlib` ABCs](https://docs.python.org/3/library/typing.html#aliases-to-contextlib-abcs)
    * [Deprecation Timeline of Major Features](https://docs.python.org/3/library/typing.html#deprecation-timeline-of-major-features)


#### Previous topic
[Development Tools](https://docs.python.org/3/library/development.html "previous chapter")
#### Next topic
[`pydoc` — Documentation generator and online help system](https://docs.python.org/3/library/pydoc.html "next chapter")
### This page
  * [Report a bug](https://docs.python.org/3/bugs.html)
  * [Show source ](https://github.com/python/cpython/blob/main/Doc/library/typing.rst?plain=1)


«
### Navigation
  * [index](https://docs.python.org/3/genindex.html "General Index")
  * [modules](https://docs.python.org/3/py-modindex.html "Python Module Index") |
  * [next](https://docs.python.org/3/library/pydoc.html "pydoc — Documentation generator and online help system") |
  * [previous](https://docs.python.org/3/library/development.html "Development Tools") |
  * ![Python logo](https://docs.python.org/3/_static/py.svg)
  * [Python](https://www.python.org/) »
  * Greek | ΕλληνικάEnglishSpanish | españolFrench | françaisItalian | italianoJapanese | 日本語Korean | 한국어Polish | polskiBrazilian Portuguese | Português brasileiroTurkish | TürkçeSimplified Chinese | 简体中文Traditional Chinese | 繁體中文
dev (3.15)pre (3.14)3.13.73.123.113.103.93.83.73.63.53.43.33.23.13.02.72.6
  * [3.13.7 Documentation](https://docs.python.org/3/index.html) » 
  * [The Python Standard Library](https://docs.python.org/3/library/index.html) »
  * [Development Tools](https://docs.python.org/3/library/development.html) »
  * [`typing` — Support for type hints](https://docs.python.org/3/library/typing.html)
  * | 
  * Theme  Auto Light Dark |


© [ Copyright ](https://docs.python.org/3/copyright.html) 2001-2025, Python Software Foundation. This page is licensed under the Python Software Foundation License Version 2. Examples, recipes, and other code in the documentation are additionally licensed under the Zero Clause BSD License. See [History and License](https://docs.python.org/license.html) for more information. The Python Software Foundation is a non-profit corporation. [Please donate.](https://www.python.org/psf/donations/) Last updated on Aug 18, 2025 (21:29 UTC). [Found a bug](https://docs.python.org/bugs.html)? Created using [Sphinx](https://www.sphinx-doc.org/) 8.2.3. 
  *[*]: Keyword-only parameters separator (PEP 3102)
  *[/]: Positional-only parameter separator (PEP 570)


## Source Information
- URL: https://docs.python.org/3/library/typing.html
- Harvested: 2025-08-19T01:08:39.569800
- Profile: quick_reference
- Agent: test
