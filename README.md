# modb

Module used as part of an example to show how dependency management works - Module Bmodb

This module exposes 3 APIs in 3 different packages:

- an API that calls the API exposed by modc that calls an API exposed by modd
- an API that calls the API exposed by modc that remains internal to modc
- an API that does not call any other external API
