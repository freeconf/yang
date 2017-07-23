

# YANG in YANG


## <a name=""></a>/



  
* **[module](#/module)** - . 







## <a name="/module"></a>/module/



  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **namespace** `string` - . 

  
* **prefix** `string` - . 

  
* **[revision](#/module/revision)** - . 

  
* **[groupings[…]](#/module/groupings)** - . 

  
* **[typedefs[…]](#/module/groupings/typedefs)** - . 

  
* **[definitions[…]](#/module/groupings/definitions)** - . 







## <a name="/module/revision"></a>/module/revision/



  
* **rev-date** `string` - . 

  
* **description** `string` - . 







## <a name="/module/groupings"></a>/module/groupings={ident}/



  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **[groupings[…]](#/module/groupings)** - . 

  
* **[typedefs[…]](#/module/groupings/typedefs)** - . 

  
* **[definitions[…]](#/module/groupings/definitions)** - . 







## <a name="/module/groupings/typedefs"></a>/module/groupings={ident}/typedefs={ident}/



  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **[type](#/module/groupings/typedefs/type)** - . 







## <a name="/module/groupings/typedefs/type"></a>/module/groupings={ident}/typedefs={ident}/type/



  
* **ident** `string` - . 

  
* **range** `string` - . 

  
* **enumeration** `string[]` - . 

  
* **path** `string` - . 

  
* **minLength** `int32` - . 

  
* **maxLength** `int32` - . 







## <a name="/module/groupings/definitions"></a>/module/groupings={ident}/definitions={ident}/



  
* **ident** `string` - . 

  
* **[container](#/module/groupings/definitions/container)** - . 

  
* **[list](#/module/groupings/definitions/list)** - . 

  
* **[leaf](#/module/groupings/definitions/leaf)** - . 

  
* **[anyxml](#/module/groupings/definitions/anyxml)** - . 

  
* **[leaf-list](#/module/groupings/definitions/leaf-list)** - . 

  
* **[uses](#/module/groupings/definitions/uses)** - . 

  
* **[choice](#/module/groupings/definitions/choice)** - . 

  
* **[notification](#/module/groupings/definitions/notification)** - . 

  
* **[action](#/module/groupings/definitions/action)** - . 







## <a name="/module/groupings/definitions/container"></a>/module/groupings={ident}/definitions={ident}/container/



  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **config** `boolean` - . 

  
* **mandatory** `boolean` - . 

  
* **[groupings[…]](#/module/groupings)** - . 

  
* **[typedefs[…]](#/module/groupings/typedefs)** - . 

  
* **[definitions[…]](#/module/groupings/definitions)** - . 







## <a name="/module/groupings/definitions/list"></a>/module/groupings={ident}/definitions={ident}/list/



  
* **key** `string[]` - . 

  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **config** `boolean` - . 

  
* **mandatory** `boolean` - . 

  
* **[groupings[…]](#/module/groupings)** - . 

  
* **[typedefs[…]](#/module/groupings/typedefs)** - . 

  
* **[definitions[…]](#/module/groupings/definitions)** - . 







## <a name="/module/groupings/definitions/leaf"></a>/module/groupings={ident}/definitions={ident}/leaf/



  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **config** `boolean` - . 

  
* **mandatory** `boolean` - . 

  
* **[type](#/module/groupings/typedefs/type)** - . 







## <a name="/module/groupings/definitions/anyxml"></a>/module/groupings={ident}/definitions={ident}/anyxml/



  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **config** `boolean` - . 

  
* **mandatory** `boolean` - . 

  
* **[type](#/module/groupings/typedefs/type)** - . 







## <a name="/module/groupings/definitions/leaf-list"></a>/module/groupings={ident}/definitions={ident}/leaf-list/



  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **config** `boolean` - . 

  
* **mandatory** `boolean` - . 

  
* **[type](#/module/groupings/typedefs/type)** - . 







## <a name="/module/groupings/definitions/uses"></a>/module/groupings={ident}/definitions={ident}/uses/



  
* **ident** `string` - . 

  
* **description** `string` - . 







## <a name="/module/groupings/definitions/choice"></a>/module/groupings={ident}/definitions={ident}/choice/



  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **[cases[…]](#/module/groupings/definitions/choice/cases)** - . 







## <a name="/module/groupings/definitions/choice/cases"></a>/module/groupings={ident}/definitions={ident}/choice/cases={ident}/



  
* **ident** `string` - . 

  
* **[definitions[…]](#/module/groupings/definitions)** - . 







## <a name="/module/groupings/definitions/notification"></a>/module/groupings={ident}/definitions={ident}/notification/



  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **[groupings[…]](#/module/groupings)** - . 

  
* **[typedefs[…]](#/module/groupings/typedefs)** - . 

  
* **[definitions[…]](#/module/groupings/definitions)** - . 







## <a name="/module/groupings/definitions/action"></a>/module/groupings={ident}/definitions={ident}/action/



  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **ident** `string` - . 

  
* **description** `string` - . 

  
* **[input](#/module/groupings/definitions/action/input)** - . 

  
* **[output](#/module/groupings/definitions/action/output)** - . 







## <a name="/module/groupings/definitions/action/input"></a>/module/groupings={ident}/definitions={ident}/action/input/



  
* **[groupings[…]](#/module/groupings)** - . 

  
* **[typedefs[…]](#/module/groupings/typedefs)** - . 

  
* **[definitions[…]](#/module/groupings/definitions)** - . 







## <a name="/module/groupings/definitions/action/output"></a>/module/groupings={ident}/definitions={ident}/action/output/



  
* **[groupings[…]](#/module/groupings)** - . 

  
* **[typedefs[…]](#/module/groupings/typedefs)** - . 

  
* **[definitions[…]](#/module/groupings/definitions)** - . 







