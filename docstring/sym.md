---
title: "sym"
date: 2019-04-06T12:19:22+02:00
description: "converts a string into a symbol"
names: ["sym"]
usage: "(sym str)"
tags: ["symbol", "macro"]
---
Converts the provided string into a symbol. Both qualified and local symbols are accepted.

#### An Example

```clojure
(def hello-sym (sym "hello"))
(eq hello-sym 'hello)
```

This example will return _true_.