

# example


## <a name=""></a>/



  
* **[bird[…]](#/bird)** - . 

  
* **level** `enumeration` - .  *Allowed Values: casual,hobbiest,birdNerd* 

  
* **country** `string` - .  *choice: origin, case: case0* 

  
* **planet** `string` - .  *choice: origin, case: case1* 

  
* **moon** `string` - .  *choice: origin, case: case1* 

  
* **[audobon](#/audobon)** - .  *choice: record, case: audobon* 

  
* **[peterson](#/peterson)** - .  *choice: record, case: peterson* 



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







## <a name="/audobon"></a>/audobon/



  
* **page** `string` - . 







## <a name="/peterson"></a>/peterson/



  
* **link** `string` - . 







