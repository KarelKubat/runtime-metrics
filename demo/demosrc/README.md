# demosrc
--
    import "github.com/KarelKubat/runtime-metrics/demo/demosrc"

Package demosrc holds demos of a server that holds metrics and publishes them,
and of the functionalities of the reporting client.

## Usage

#### func  CheckErr

```go
func CheckErr(err error)
```
CheckErr checks whether its argument is a non-nil error and panics if so. It's
included to simplify the examples.

#### func  ClientAllNamesDemo

```go
func ClientAllNamesDemo()
```
ClientAllNamesDemo runs a reporter client, fetches all metric names, and
displays them.

#### func  ClientFullDumpDemo

```go
func ClientFullDumpDemo()
```
ClientFullDumpDemo runs a reporter client, obtains a full metrics dump from the
server, and displays all metrics.

#### func  ClientScrapeDemo

```go
func ClientScrapeDemo()
```
ClientScrapeDemo runs a reporter client that fetches server metrics in an
endless loop (or until the server exits). The metrics are displayed on stdout.

#### func  PublishingProgramDemo

```go
func PublishingProgramDemo()
```
PublishingProgramDemo shows how a program can inspect its own metrics, so that
they can be e.g. pushed to a monitoring service. This is a different approach
than having the metrics in a server, and having a client scrape them and process
them further.

Advantage: All-in-one, no need for TCP ports that can already be take (because
the server/client force you to pick a por);

Disadvantage: Forces your program to be midful of what could go wrong when
sending metrics to a remote service, how to recover when out of quota, etc..

#### func  ServerDemo

```go
func ServerDemo()
```
ServerDemo starts a reporter, creates some metrics, and then in a loop changes
their values - so that clients may scrape them.
