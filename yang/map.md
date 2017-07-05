

# Map


## <a name="/map"></a>/



  
* **[device[…]](#/device)** - . 



### Actions:

* <a name="/map/register"></a>**/register** - 
 
  
#### Input:
>
* **deviceId** `string` - Id that is unique to this device in the infrastructure pool
* **address** `string` - Optional.  Will use incoming address of request.  Hint: If you use the text
                  phrase &#39;{REG_ADD}&#39; anywhere in the address, it will be replaced by the IP address found
                  in the registration request. This does not include the port number because often that
                  is not typically the port used when registering.  Example  https://{REG_ADDR}:8090/restconf


  





### Events:

* <a name="/map/update"></a>**/update** - 

  
>
* **deviceId** `string` - 
* **address** `string` - 
* **module[…]** `` - 
    * **name** -  
    * **revision** -  
* **change** `enumeration` - 





## <a name="/device"></a>/device



  
* **deviceId** `string` - . 

  
* **address** `string` - . 

  
* **[module[…]](#/device={deviceId}/module)** - . 







## <a name="/device={deviceId}/module"></a>/device={deviceId}/module



  
* **name** `string` - . 

  
* **revision** `string` - . 







