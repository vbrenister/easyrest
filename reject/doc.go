// This package provides helper functions to reject the incoming request.
// In order to facilitate of the usage of this package, it is required to create a new instance of the LoggerSupport struct and provide an instance of slog.Logger.
// The response format is JSON and the error message is provided in the "error" field.
// Example:
// { "error": "Not Found" }
package reject
