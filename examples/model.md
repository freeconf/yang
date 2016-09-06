# Module

top-most container of your module.  Acts like a "container"
```
module my-module {
```
Useful when combining YANG from multiple vendors to avoid name clashing
```
  namespace "";
  prefix "";
```
  
API versioning
```
  revision 20160905;
```

effectively copies the entire module into this file
```
  import my-other-module;
```  



YANG | Description
--- | ---
`module my-module {` | top-most container of your module.  Acts like a "container"
`namespace "";` |
`prefix "";` | Useful when combining YANG from multiple vendors to avoid name clashing


  
API versioning
```
  revision 20160905;
```

effectively copies the entire module into this file
```
  import my-other-module;
```  

# Leaf/Leaf-list

```
leaf {
  /* 
    Options: 
       string
       int32
       int64
  */
  type string;
}



```

# Typedef
```


```