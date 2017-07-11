

# Map


## <a name=""></a>/



  
* **[device[…]](#/device)** - . 



### Actions:

* <a name="/register"></a>**/register** - 
 
  
#### Input:
> * **deviceId** `string` - Id that is unique to this device in the infrastructure pool
> * **address** `string` - Optional.  Will use incoming address of request.  Hint: If you use the text
                  phrase &#39;{REG_ADD}&#39; anywhere in the address, it will be replaced by the IP address found
                  in the registration request. This does not include the port number because often that
                  is not typically the port used when registering.  Example  https://{REG_ADDR}:8090/restconf


  





### Events:

* <a name="/update"></a>**/update** - 

 	
> * **deviceId** `string` - 	
> * **address** `string` - 
> * **module[…]** - 
>     * **name** -  
>     * **revision** -  	
> * **change** `enumeration` - 





## <a name="/device"></a>/device={deviceId}/



  
* **deviceId** `string` - . 

  
* **address** `string` - . 

  
* **[module[…]](#/device/module)** - . 







## <a name="/device/module"></a>/device={deviceId}/module={name}/



  
* **name** `string` - . 

  
* **revision** `string` - . 







