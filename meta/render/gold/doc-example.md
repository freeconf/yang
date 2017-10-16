

# example


## <a name=""></a>/



  
* **[bird[…]](#/bird)** - . 

  
* **level** `enumeration` - .  *Allowed Values: casual,hobbiest,birdNerd* 



### Actions:

* <a name="/fly"></a>**/fly** - 
 
  
#### Input:
> * **vector** - 
>     * **x** -  
> * **originalWeight** `decimal64` - 


  
#### Output:
> * **log[…]** - 
>     * **length** -  
> * **speed** `decimal64` - 





### Events:

* <a name="/migration"></a>**/migration** - 

 	
> * **logEntry** `string` - 
> * **status** - 
>     * **energyLevel** -  





## <a name="/bird"></a>/bird={name}/



  
* **name** `string` - . 

  
* **wingSpan** `int32` - in cm.  *Default: 64* 







