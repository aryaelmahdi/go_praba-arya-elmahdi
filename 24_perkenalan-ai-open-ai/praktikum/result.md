Sure! Here's an example of how you can create a RESTful API with an OpenAPI connection in Golang using the popular Gorilla Mux router and swaggo for generating the OpenAPI specification.

First, make sure you have Go installed on your system.

Step 1: Set up the project structure

Create a new folder for your project and navigate into it. Then, create the following folders inside it:

```
- cmd
- internal
- pkg
- api
- docs
```

Step 2: Install the necessary packages

In your project folder, open a terminal and run the following command to install the Gorilla Mux router:

```bash
go get -u github.com/gorilla/mux
```

Next, install swaggo for generating the OpenAPI specification:

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

Step 3: Set up the API

Inside the `api` folder, create a file `user.go` with the following content:

```go
package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Username string `