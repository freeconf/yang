

# RESTConf


## <a name="/restconf"></a>/
service that implements RESTCONF RFC8040 device protocol


  
* **notifyKeepaliveTimeoutMs** `int32` - close the connection after N milliseconds of no pings or activity.  *Default: 30000* 

  
* **debug** `boolean` - enable debug log messages.  *Default: false* 

  
* **streamCount** `int32` - number of open sessions. each session have have many subscriptions. 

  
* **subscriptionCount** `int32` - number of subscriptions across all sessions. 

  
* **[web{…}](#/web)** - web service used by restconf server. 

  
* **[callHome{…}](#/callHome)** - . 







## <a name="/web"></a>/web
web service used by restconf server


  
* **port** `string` - required port number.  Examples :8010  192.168.1.10:8080. 

  
* **readTimeout** `int32` - timeout in milliseconds to wait for reading data from client.  *Default: 10000* 

  
* **writeTimeout** `int32` - timeout in milliseconds for sending data from client.  *Default: 10000* 

  
* **[tls{…}](#/web/tls)** - required for secure transport. 







## <a name="/web/tls"></a>/web/tls
required for secure transport


  
* **serverName** `string` - Name identified in certificate for this server. 

  
* **[cert{…}](#/web/tls/cert)** - . 

  
* **[ca{…}](#/web/tls/ca)** - . 







## <a name="/web/tls/cert"></a>/web/tls/cert



  
* **certFile** `string` - PEM encoded certification. 

  
* **keyFile** `string` - PEM encoded private key used to build certificate. 







## <a name="/web/tls/ca"></a>/web/tls/ca



  
* **certFile** `string` - PEM encoded certificate of certificate authority used to sign certificate. 







## <a name="/callHome"></a>/callHome



  
* **registered** `boolean` - Success registration. 

  
* **deviceId** `string` - Unique device id within your infrastructure for this device.. 

  
* **address** `string` - Hostname or IP address of application management system.. 

  
* **endpoint** `string` - endpoint appended to address to device to register to. e.g. /restconf. 

  
* **localAddress** `string` - When client is initiating connection to a registration server, this is the network address. 

  
* **retryRateMs** `int32` - If registration fails, try again after given ms..  *Default: 10000* 



### Actions:

* <a name="/callHome/register"></a>**/callHome/register** - 
 
  


  


* <a name="/callHome/unregister"></a>**/callHome/unregister** - 
 
  


  





### Events:

* <a name="/callHome/update"></a>**/callHome/update** - Change in registration status.

  
>
* **registered** `boolean` - 
* **err** `string` - Last registration error if there was one





