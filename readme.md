## Request Options
| Options                 | Flag | Short | Param Required | Description                                               |
|-------------------------|------|-------|----------------|-----------------------------------------------------------|
| JSON                    | `json` | `j` | true           | Adds a JSON payload to the HTTP request body.              |
| Data                    | `data` | `d` | true           | Adds key-value pairs as query params to the request URL |
| URL                     | `url`  |      | true           | Sets the URL for the HTTP request.                         |
| Port                    | `port` | `P` | true           | Sets the port for the HTTP request.                        |
| Path                    | `path` | `p` | true           | Appends a path to the HTTP request URL. Default URL is localhost.com. If URL is given as the first command line argument, the path is appended. |
| Scheme                  | `scheme` | `s` | true         | Sets the protocol for the HTTP request.                    |
| Content-Type Header     | `c` | `c` | true             | Sets the 'Content-Type' header for the HTTP request.       |
| Cookie                  | `cookie` | `C` | true             | Sets a cookie for the HTTP request.                        |
| GET                     | `get` | `g` | false          | Sets the HTTP method to GET.                               |
| PUT                     | `put` |      | false          | Sets the HTTP method to PUT.                               |
| POST                    | `post` |      | false          | Sets the HTTP method to POST.                              |
| DELETE                  | `delete` |      | false        | Sets the HTTP method to DELETE.                            |
| HEAD                    | `head` |      | false          | Sets the HTTP method to HEAD.                              |
| PATCH                   | `patch` |     | false          | Sets the HTTP method to PATCH.                             |
| TRACE                   | `trace` |     | false          | Sets the HTTP method to TRACE.                             |


## Client Options
| Options                 | Flag | Short | Param Required | Description                                               |
|-------------------------|------|-------|----------------|-----------------------------------------------------------|
| Connect Timeout        | `connect-timeout` | | true        | Sets a limit in seconds for a connection request timeout. |
