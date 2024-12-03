# Transfer Tool

generate a file `credentials.json` from google cloud project

[scopes](https://developers.google.com/identity/protocols/oauth2/scopes)

setup domain delegation scopes. currently this is just using lazy scopes (global read/write) - better security is to use `user:read` but for the transfer you'll need the full scope.

https://www.googleapis.com/auth/admin.datatransfer
https://www.googleapis.com/auth/admin.directory.user.readonly

to use the more secure scope modify `service.go` to use `AdminDirectoryUserReadonlyScope`
