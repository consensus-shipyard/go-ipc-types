# Go IPC types

This repository contains primitive and low level types used in IPC's FVM
actors state and params representation.

This is intended for seamless interaction with the state of
[subnet-actor](https://github.com/consensus-shipyard/ipc-subnet-actor)
and [ipc-gateway](https://github.com/consensus-shipyard/ipc-gateway).


## Versioning

Blockchain state versioning does not naturally align with common semantic versioning conventions.

Any change in behaviour, including repairing any error that may have affected blockchain evaluation,
must be released in a major version change. We intend that to be a rare event for the contents of 
this repository.

## License
This repository is dual-licensed under Apache 2.0 and MIT terms.

Copyright 2020. Protocol Labs, Inc.
