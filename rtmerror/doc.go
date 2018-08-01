/*
Package rtmerror implements the error type for runtime-metrics.

Errors in runtime-metrics are a custom error type, rtmerror.Error. You
will normally not care, except when checking whether an error is
retryable. This only occurs during network calls, and normally the
backoff policy of the reporting client takes care of retries anyway.

*/
package rtmerror
