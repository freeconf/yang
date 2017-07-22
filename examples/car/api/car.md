

# Car


## <a name=""></a>/
Vehicle of sorts


  
* **[tire[…]](#/tire)** - rubber circular part that makes contact with road. 

  
* **miles** `int64` - . 

  
* **lastRotation** `int64` - . 

  
* **running** `boolean` - . 

  
* **speed** `int32` - number of millisecs it takes to travel one mile.  *Default: 1000* 

  
* **[engine](#/engine)** - . 



### Actions:

* <a name="/rotateTires"></a>**/rotateTires** - rotate tires for optimal wear
 
  


  


* <a name="/replaceTires"></a>**/replaceTires** - replace all tires
 
  


  





### Events:

* <a name="/update"></a>**/update** - important state information about your car

 
> * **tire[…]** - rubber circular part that makes contact with road
>     * **pos** -  
>     * **size** -  Default: 15
>     * **worn** -  
>     * **wear** -  
>     * **flat** -  	
> * **miles** `int64` - 	
> * **lastRotation** `int64` - 	
> * **running** `boolean` - 	
> * **speed** `int32` - number of millisecs it takes to travel one mile





## <a name="/tire"></a>/tire={pos}/
rubber circular part that makes contact with road


  
* **pos** `int32` - . 

  
* **size** `string` - .  *Default: 15* 

  
* **worn** `boolean` - . 

  
* **wear** `decimal64` - . 

  
* **flat** `boolean` - . 







## <a name="/engine"></a>/engine/



  
* **[specs](#/engine/specs)** - . 







## <a name="/engine/specs"></a>/engine/specs/



  
* **horsepower** `int32` - . 







