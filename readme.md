### Creq
Build a request from the command line. The current state of the default request is displayed as seen below. 
```
Current Request
GET http://localhost:3001
```
Modify the state of the request e.g. 
```
$: -j {"user":"1"} -P /users/id 
```

See updated request.
```
Current Request
GET http://localhost:3001/users/id
Content-Type : [application/json]
{{"user":"1"}}
```
Send the request as many times as required and/or continue to modify via the command line.

## Request Options
| Options                         | Flag | Short | Param Required| Description                                               |
|:--------|:------|:-------|:---------------|:-----------------------------------------------------------|
| JSON                            | `json` | `j` | true          | Adds a JSON payload to the HTTP request body. Changes request method to POST by default          |
| Data                            | `data` | `d` | true          | Adds key-value pairs as query params to the request URL e.g. -d foo=bar|
| URL                             | `url`  |      | true         | Sets the URL for the request.                         |
| User                            | `user` | `u` | true          | Sets the user and password for basic url authentication. Example: -u user:password                        |
| Port                            | `port` | `P` | true          | Sets the port for the request URL.                        |
| Path                            | `path` | `p` | true          | Sets the path for the request URL.                    |
| Scheme                          | `scheme` | `s` | true        | Sets the protocol for the request. http/https    |
| Header                          | `header` | `h` | true        | Applies the given header to the request e.g. foo:bar       |
| Header: Content-Type            | `c` | `c` | true             | Sets the 'Content-Type' header for the HTTP request.       |
| Cookie                          | `cookie` | `C` | true        | Sets a cookie for the HTTP request.                        |
| Get Request                     | `get` | `g` | false          | Sets the request method to GET.                               |
| Put Request                     | `put` |      | false         | Sets the request method to PUT.                               |
| Post Request                    | `post` |      | false        | Sets the request method to POST.                              |
| Delete Request                  | `delete` |      | false      | Sets the request method to DELETE.                            |
| Head Request                    | `head` |      | false        | Sets the request method to HEAD.                              |
| Patch Request                   | `patch` |     | false        | Sets the request method to PATCH.                             |
| Trace Request                   | `trace` |     | false        | Sets the request method to TRACE.                             |
| Reset                           | `reset` |  `x`   | false     | Returns the current state of the request to the default settings |
| Send                            | `send` |  `s`   | false     | Sends the current request |
| Dump Header                            | `dump-header` |  `D`   | true     | Writes the response headers to the specified file.  |


Creq takes an optional URL as the first argument or use the --url flag to explicity change the url from the default. 

