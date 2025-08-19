---
agent_context: test
confidence: 0.95
harvested_at: '2025-08-19T01:05:40.036929'
profile: quick_reference
source: https://docs.python.org/3/library/asyncio.html
topic: Python Async Programming
---

# Python Async Programming

[ ![Python logo](https://docs.python.org/3/_static/py.svg) ](https://www.python.org/) dev (3.15)pre (3.14)3.13.73.123.113.103.93.83.73.63.53.43.33.23.13.02.72.6
Greek | ΕλληνικάEnglishSpanish | españolFrench | françaisItalian | italianoJapanese | 日本語Korean | 한국어Polish | polskiBrazilian Portuguese | Português brasileiroTurkish | TürkçeSimplified Chinese | 简体中文Traditional Chinese | 繁體中文
Theme  Auto Light Dark
#### Previous topic
[Networking and Interprocess Communication](https://docs.python.org/3/library/ipc.html "previous chapter")
#### Next topic
[Runners](https://docs.python.org/3/library/asyncio-runner.html "next chapter")
### This page
  * [Report a bug](https://docs.python.org/3/bugs.html)
  * [Show source ](https://github.com/python/cpython/blob/main/Doc/library/asyncio.rst?plain=1)


### Navigation
  * [index](https://docs.python.org/3/genindex.html "General Index")
  * [modules](https://docs.python.org/3/py-modindex.html "Python Module Index") |
  * [next](https://docs.python.org/3/library/asyncio-runner.html "Runners") |
  * [previous](https://docs.python.org/3/library/ipc.html "Networking and Interprocess Communication") |
  * ![Python logo](https://docs.python.org/3/_static/py.svg)
  * [Python](https://www.python.org/) »
  * Greek | ΕλληνικάEnglishSpanish | españolFrench | françaisItalian | italianoJapanese | 日本語Korean | 한국어Polish | polskiBrazilian Portuguese | Português brasileiroTurkish | TürkçeSimplified Chinese | 简体中文Traditional Chinese | 繁體中文
dev (3.15)pre (3.14)3.13.73.123.113.103.93.83.73.63.53.43.33.23.13.02.72.6
  * [3.13.7 Documentation](https://docs.python.org/3/index.html) » 
  * [The Python Standard Library](https://docs.python.org/3/library/index.html) »
  * [Networking and Interprocess Communication](https://docs.python.org/3/library/ipc.html) »
  * [`asyncio` — Asynchronous I/O](https://docs.python.org/3/library/asyncio.html)
  * | 
  * Theme  Auto Light Dark |


# `asyncio` — Asynchronous I/O[¶](https://docs.python.org/3/library/asyncio.html#module-asyncio "Link to this heading")
Hello World!
Copy```
importasyncio
async defmain():
  print('Hello ...')
  await asyncio.sleep(1)
  print('... World!')
asyncio.run(main())

```

asyncio is a library to write **concurrent** code using the **async/await** syntax.
asyncio is used as a foundation for multiple Python asynchronous frameworks that provide high-performance network and web-servers, database connection libraries, distributed task queues, etc.
asyncio is often a perfect fit for IO-bound and high-level **structured** network code.
See also 

[A Conceptual Overview of asyncio](https://docs.python.org/3/howto/a-conceptual-overview-of-asyncio.html#a-conceptual-overview-of-asyncio)
    
Explanation of the fundamentals of asyncio.
asyncio provides a set of **high-level** APIs to:
  * [run Python coroutines](https://docs.python.org/3/library/asyncio-task.html#coroutine) concurrently and have full control over their execution;
  * perform [network IO and IPC](https://docs.python.org/3/library/asyncio-stream.html#asyncio-streams);
  * control [subprocesses](https://docs.python.org/3/library/asyncio-subprocess.html#asyncio-subprocess);
  * distribute tasks via [queues](https://docs.python.org/3/library/asyncio-queue.html#asyncio-queues);
  * [synchronize](https://docs.python.org/3/library/asyncio-sync.html#asyncio-sync) concurrent code;


Additionally, there are **low-level** APIs for _library and framework developers_ to:
  * create and manage [event loops](https://docs.python.org/3/library/asyncio-eventloop.html#asyncio-event-loop), which provide asynchronous APIs for [networking](https://docs.python.org/3/library/asyncio-eventloop.html#loop-create-server), running [subprocesses](https://docs.python.org/3/library/asyncio-eventloop.html#loop-subprocess-exec), handling [OS signals](https://docs.python.org/3/library/asyncio-eventloop.html#loop-add-signal-handler), etc;
  * implement efficient protocols using [transports](https://docs.python.org/3/library/asyncio-protocol.html#asyncio-transports-protocols);
  * [bridge](https://docs.python.org/3/library/asyncio-future.html#asyncio-futures) callback-based libraries and code with async/await syntax.


[Availability](https://docs.python.org/3/library/intro.html#availability): not WASI.
This module does not work or is not available on WebAssembly. See [WebAssembly platforms](https://docs.python.org/3/library/intro.html#wasm-availability) for more information.
asyncio REPL
You can experiment with an `asyncio` concurrent context in the [REPL](https://docs.python.org/3/glossary.html#term-REPL):
Copy```
$ python -m asyncio
asyncio REPL ...
Use "await" directly instead of "asyncio.run()".
Type "help", "copyright", "credits" or "license" for more information.
>>> importasyncio
>>> await asyncio.sleep(10, result='hello')
'hello'

```

Raises an [auditing event](https://docs.python.org/3/library/sys.html#auditing) `cpython.run_stdin` with no arguments.
Changed in version 3.12.5: (also 3.11.10, 3.10.15, 3.9.20, and 3.8.20) Emits audit events.
Changed in version 3.13: Uses PyREPL if possible, in which case [`PYTHONSTARTUP`](https://docs.python.org/3/using/cmdline.html#envvar-PYTHONSTARTUP) is also executed. Emits audit events.
Reference
High-level APIs
  * [Runners](https://docs.python.org/3/library/asyncio-runner.html)
  * [Coroutines and Tasks](https://docs.python.org/3/library/asyncio-task.html)
  * [Streams](https://docs.python.org/3/library/asyncio-stream.html)
  * [Synchronization Primitives](https://docs.python.org/3/library/asyncio-sync.html)
  * [Subprocesses](https://docs.python.org/3/library/asyncio-subprocess.html)
  * [Queues](https://docs.python.org/3/library/asyncio-queue.html)
  * [Exceptions](https://docs.python.org/3/library/asyncio-exceptions.html)


Low-level APIs
  * [Event Loop](https://docs.python.org/3/library/asyncio-eventloop.html)
  * [Futures](https://docs.python.org/3/library/asyncio-future.html)
  * [Transports and Protocols](https://docs.python.org/3/library/asyncio-protocol.html)
  * [Policies](https://docs.python.org/3/library/asyncio-policy.html)
  * [Platform Support](https://docs.python.org/3/library/asyncio-platforms.html)
  * [Extending](https://docs.python.org/3/library/asyncio-extending.html)


Guides and Tutorials
  * [High-level API Index](https://docs.python.org/3/library/asyncio-api-index.html)
  * [Low-level API Index](https://docs.python.org/3/library/asyncio-llapi-index.html)
  * [Developing with asyncio](https://docs.python.org/3/library/asyncio-dev.html)


Note
The source code for asyncio can be found in [Lib/asyncio/](https://github.com/python/cpython/tree/3.13/Lib/asyncio/).
#### Previous topic
[Networking and Interprocess Communication](https://docs.python.org/3/library/ipc.html "previous chapter")
#### Next topic
[Runners](https://docs.python.org/3/library/asyncio-runner.html "next chapter")
### This page
  * [Report a bug](https://docs.python.org/3/bugs.html)
  * [Show source ](https://github.com/python/cpython/blob/main/Doc/library/asyncio.rst?plain=1)


«
### Navigation
  * [index](https://docs.python.org/3/genindex.html "General Index")
  * [modules](https://docs.python.org/3/py-modindex.html "Python Module Index") |
  * [next](https://docs.python.org/3/library/asyncio-runner.html "Runners") |
  * [previous](https://docs.python.org/3/library/ipc.html "Networking and Interprocess Communication") |
  * ![Python logo](https://docs.python.org/3/_static/py.svg)
  * [Python](https://www.python.org/) »
  * Greek | ΕλληνικάEnglishSpanish | españolFrench | françaisItalian | italianoJapanese | 日本語Korean | 한국어Polish | polskiBrazilian Portuguese | Português brasileiroTurkish | TürkçeSimplified Chinese | 简体中文Traditional Chinese | 繁體中文
dev (3.15)pre (3.14)3.13.73.123.113.103.93.83.73.63.53.43.33.23.13.02.72.6
  * [3.13.7 Documentation](https://docs.python.org/3/index.html) » 
  * [The Python Standard Library](https://docs.python.org/3/library/index.html) »
  * [Networking and Interprocess Communication](https://docs.python.org/3/library/ipc.html) »
  * [`asyncio` — Asynchronous I/O](https://docs.python.org/3/library/asyncio.html)
  * | 
  * Theme  Auto Light Dark |


© [ Copyright ](https://docs.python.org/3/copyright.html) 2001-2025, Python Software Foundation. This page is licensed under the Python Software Foundation License Version 2. Examples, recipes, and other code in the documentation are additionally licensed under the Zero Clause BSD License. See [History and License](https://docs.python.org/license.html) for more information. The Python Software Foundation is a non-profit corporation. [Please donate.](https://www.python.org/psf/donations/) Last updated on Aug 18, 2025 (21:29 UTC). [Found a bug](https://docs.python.org/bugs.html)? Created using [Sphinx](https://www.sphinx-doc.org/) 8.2.3. 


## Source Information
- URL: https://docs.python.org/3/library/asyncio.html
- Harvested: 2025-08-19T01:05:40.036929
- Profile: quick_reference
- Agent: test
