
# example



<details><summary>API Usage Notes:</summary>

#### General API Usage Notes
* `DELETE` implementation may be disallowed or ignored depending on the context
* Lists use `../path={key}/...` instead of `.../path/key/...` to avoid API name collision

#### `GET` Query Parameters

These parameters can be combined.

> | param                            | description | example |
> |----------------------------------|-------------|---------|
> | `content=[non-config\|config]` | Show only read-only fields or only read/write fields |   `.../path?content=config`|
> | `fields=field1;field2` | Return a portion of the data limited to fields listed | `.../path?fields=user%2faddress` |
> | `depth=n` | Return a portion of the data limited to depth of the hierarchy | `.../path?depth=1`
> | `fc.xfields=field1;fields` | Return a portion of the data excluding the fields listed | `.../path?fc.xfields=user%2faddress` |
> | `fc.range=field!{startRow}-[{endRow}]` | For lists, return only limited number of rows | `.../path?fc.range=user!10-20` 

</details>





<details>
 <summary><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:doc-example</b></code> </summary>

#### doc-example

**GET Response Data / PUT, POST Request Data**
````json
{
  "bird":[{
     "name":"",
     "family":"",
     "wingSpan":0
  }],
  "level":"",
  "country":"",
  "planet":"",
  "moon":"",
  "audobon":{
     "page":""
  },
  "peterson":{
     "link":""
  }}
````



**Data Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | bird.name | string  |   |  |
> | bird.family | identityref  |   |  |
> | bird.wingSpan | int32  |  in cm | Default: 64 |
> | level | enumeration  |   | Allowed Values: casual,hobbiest,birdNerd |
> | country | string  |   | choice: origin, case: case0 |
> | planet | string  |   | choice: origin, case: case1 |
> | moon | string  |   | choice: origin, case: case1 |
> | audobon.page | string  |   |  |
> | peterson.link | string  |   |  |

**Responses**
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

**HTTP response codes**
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

**Examples**
````bash
# retrieve data
curl https://server/restconf/data/acc:doc-example

# update existing data
curl -X PUT -d @data.json https://server/restconf/data/acc:doc-example

# create new data
curl -X POST -d @data.json https://server/restconf/data/acc:doc-example

# delete current data
curl -X DELETE https://server/restconf/data/acc:doc-example
````
</details>





<details>
 <summary><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:bird</b></code> </summary>

#### bird

**GET Response Data / PUT, POST Request Data**
````json
{"bird":[
  "name":"",
  "family":"",
  "wingSpan":0}, {"..."}]}
````



**Data Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | name | string  |   |  |
> | family | identityref  |   |  |
> | wingSpan | int32  |  in cm | Default: 64 |

**Responses**
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

**HTTP response codes**
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

**Examples**
````bash
# retrieve data
curl https://server/restconf/data/acc:bird

# update existing data
curl -X PUT -d @data.json https://server/restconf/data/acc:bird

# create new data
curl -X POST -d @data.json https://server/restconf/data/acc:bird

# delete current data
curl -X DELETE https://server/restconf/data/acc:bird
````
</details>




<details>
 <summary><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:bird={name}</b></code> </summary>

#### bird={name}

**GET Response Data / PUT, POST Request Data**
````json
{
  "name":"",
  "family":"",
  "wingSpan":0}
````



**Data Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | name | string  |   |  |
> | family | identityref  |   |  |
> | wingSpan | int32  |  in cm | Default: 64 |

**Responses**
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

**HTTP response codes**
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

**Examples**
````bash
# retrieve data
curl https://server/restconf/data/acc:bird={name}

# update existing data
curl -X PUT -d @data.json https://server/restconf/data/acc:bird={name}

# create new data
curl -X POST -d @data.json https://server/restconf/data/acc:bird={name}

# delete current data
curl -X DELETE https://server/restconf/data/acc:bird={name}
````
</details>





<details>
 <summary><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:audobon</b></code> </summary>

#### audobon

**GET Response Data / PUT, POST Request Data**
````json
{
  "page":""}
````



**Data Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | page | string  |   |  |

**Responses**
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

**HTTP response codes**
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

**Examples**
````bash
# retrieve data
curl https://server/restconf/data/acc:audobon

# update existing data
curl -X PUT -d @data.json https://server/restconf/data/acc:audobon

# create new data
curl -X POST -d @data.json https://server/restconf/data/acc:audobon

# delete current data
curl -X DELETE https://server/restconf/data/acc:audobon
````
</details>





<details>
 <summary><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:peterson</b></code> </summary>

#### peterson

**GET Response Data / PUT, POST Request Data**
````json
{
  "link":""}
````



**Data Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | link | string  |   |  |

**Responses**
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

**HTTP response codes**
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

**Examples**
````bash
# retrieve data
curl https://server/restconf/data/acc:peterson

# update existing data
curl -X PUT -d @data.json https://server/restconf/data/acc:peterson

# create new data
curl -X POST -d @data.json https://server/restconf/data/acc:peterson

# delete current data
curl -X DELETE https://server/restconf/data/acc:peterson
````
</details>




  <details>
 <summary><code>[POST]</code> <code><b>restconf/data/acc:fly</b></code> </summary>
 
#### fly

 **Request Body**
    
      
````json
{
  "vector":{
     "x":""
  },
  "x":"",
  "originalWeight":0
}
````

**Request Body Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | vector.x | string  |   |  |
> | x | string  |   |  |
> | originalWeight | decimal64  |   |  |
    

**Response Body**
    
      
````json
{
  "log":[{
     "length":0
  }],
  "length":0,
  "speed":0
}
````

**Response Body Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | log.length | int32  |   |  |
> | length | int32  |   |  |
> | speed | decimal64  |   |  |
    

**HTTP response codes**

> | http code |  reason for code |
> |-----------|------------------|
> | 200       | success          |
> | 401       | not authorized   |
> | 400       | invalid request  |
> | 404       | data does not exist |
> | 500       | internal error   |

**Examples**
````bash
# call function
curl -X POST -d @request.json] https://server/restconf/data/acc:fly
````
  </details>

  



  <details>
 <summary><code>[GET]</code> <code><b>restconf/data/acc:migration</b></code> </summary>

#### migration

**Response Stream** [SSE Format](https://html.spec.whatwg.org/multipage/server-sent-events.html#server-sent-events)

````
data: {first JSON message all on one line followed by 2 CRLFs}

data: {next JSON message with same format all on one line ...}
````

Each JSON message would have following data
````json
{
  "logEntry":"",
  "status":{
     "energyLevel":0
  },
  "energyLevel":0,
  "choice1":"",
  "choice2":""
}
````

**Response Body Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | logEntry | string  |   |  |
> | status.energyLevel | int64  |   |  |
> | energyLevel | int64  |   |  |
> | choice1 | string  |   | choice: notifChoice, case: choice1 |
> | choice2 | string  |   | choice: notifChoice, case: choice2 |

**Example**
````bash
# retrieve data stream, adjust timeout for slower streams
curl -N https://server/restconf/data/acc:migration
````

</details>
  

