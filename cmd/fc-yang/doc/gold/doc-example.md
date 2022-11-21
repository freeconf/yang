
# example## <a name=""></a>






<details>
 <summary><a name="doc-example"></a><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:doc-example</b></code> </summary>

#### GET Response Data / PUT, POST Request Data
````
{
  "bird":[{
     "name":"",
     "family":"",
     "wingSpan":n
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



#### Data Details

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

#### Responses
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

#### HTTP response codes
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

#### Examples
````
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
 <summary><a name="bird"></a><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:bird</b></code> </summary>

#### GET Response Data / PUT, POST Request Data
````
{"bird":[
  "name":"",
  "family":"",
  "wingSpan":n},...]}
````



#### Data Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | name | string  |   |  |
> | family | identityref  |   |  |
> | wingSpan | int32  |  in cm | Default: 64 |

#### Responses
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

#### HTTP response codes
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

#### Examples
````
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
 <summary><a name="bird={name}"></a><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:bird={name}</b></code> </summary>

#### GET Response Data / PUT, POST Request Data
````
{
  "name":"",
  "family":"",
  "wingSpan":n}
````



#### Data Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | name | string  |   |  |
> | family | identityref  |   |  |
> | wingSpan | int32  |  in cm | Default: 64 |

#### Responses
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

#### HTTP response codes
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

#### Examples
````
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
 <summary><a name="audobon"></a><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:audobon</b></code> </summary>

#### GET Response Data / PUT, POST Request Data
````
{
  "page":""}
````



#### Data Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | page | string  |   |  |

#### Responses
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

#### HTTP response codes
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

#### Examples
````
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
 <summary><a name="peterson"></a><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:peterson</b></code> </summary>

#### GET Response Data / PUT, POST Request Data
````
{
  "link":""}
````



#### Data Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | link | string  |   |  |

#### Responses
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

#### HTTP response codes
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

#### Examples
````
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
 <summary><a name="fly"></a><code>[POST]</code> <code><b>restconf/data/acc:fly</b></code> </summary>

##### Request Body
    
      
````
{
  "vector":{
     "x":""
  },
  "x":"",
  "originalWeight":n
}
````

#### Request Body Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | vector.x | string  |   |  |
> | x | string  |   |  |
> | originalWeight | decimal64  |   |  |
    

##### Response Body
    
      
````
{
  "log":[{
     "length":n
  }],
  "length":n,
  "speed":n
}
````

#### Response Body Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | log.length | int32  |   |  |
> | length | int32  |   |  |
> | speed | decimal64  |   |  |
    

  <details><summary>more</summary>

#### HTTP response codes
> | http code |  reason for code |
> |-----------|------------------|
> | 200       | success          |
> | 401       | not authorized   |
> | 400       | invalid request  |
> | 404       | data does not exist |
> | 500       | internal error   |

#### Examples
````
# create new data
curl -X POST -d @request.json https://server/restconf/data/acc:fly
````
  </details>

</details>
  



  <details>
 <summary><a name="migration"></a><code>[GET]</code> <code><b>restconf/data/acc:migration</b></code> </summary>

##### Response Stream [SSE Format](https://html.spec.whatwg.org/multipage/server-sent-events.html#server-sent-events)

````
data: {
  "logEntry":"",
  "status":{
     "energyLevel":n
  },
  "energyLevel":n,
  "choice1":"",
  "choice2":""}\n
\n
data: `{  ... next message with same format ... }`
````

#### Response Body Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | logEntry | string  |   |  |
> | status.energyLevel | int64  |   |  |
> | energyLevel | int64  |   |  |
> | choice1 | string  |   | choice: notifChoice, case: choice1 |
> | choice2 | string  |   | choice: notifChoice, case: choice2 |

</details>
  

