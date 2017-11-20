

# IETF YANG Library


## <a name=""></a>/
This module contains monitoring information about the YANG
modules and submodules that are used within a YANG-based
server.

Copyright (c) 2016 IETF Trust and the persons identified as
authors of the code.  All rights reserved.

Redistribution and use in source and binary forms, with or
without modification, is permitted pursuant to, and subject
to the license terms contained in, the Simplified BSD License
set forth in Section 4.c of the IETF Trust&#39;s Legal Provisions
Relating to IETF Documents
(http://trustee.ietf.org/license-info).

This version of this YANG module is part of RFC 7895; see
the RFC itself for full legal notices.

NOTE: This file has been modified to be compatible with freeconf&#39;s 
YANG parser. Intention is to replace this file with the original 
version once the freeconf parser is in compliance.


  
* **[modules-state](#/modules-state)** - The module data structure is represented as a grouping
so it can be reused in configuration or another monitoring
data structure.. 





### Events:

* <a name="/yang-library-change"></a>**/yang-library-change** - Generated when the set of modules and submodules supported
by the server has changed.

 	
> * **module-set-id** `leafref` - Contains the module-set-id value representing the
set of modules and submodules supported at the server at
the time the notification is generated.





## <a name="/modules-state"></a>/modules-state/
The module data structure is represented as a grouping
so it can be reused in configuration or another monitoring
data structure.


  
* **module-set-id** `string` - Contains a server-specific identifier representing
the current set of modules and submodules.  The
server MUST change the value of this leaf if the
information represented by the &#39;module&#39; list instances
has changed.. 

  
* **[module[…]](#/modules-state/module)** - Each entry represents one revision of one module
currently supported by the server.. 







## <a name="/modules-state/module"></a>/modules-state/module={name}/
Each entry represents one revision of one module
currently supported by the server.


  
* **name** `string` - The YANG module or submodule name.. 

  
* **revision** `string` - The YANG module or submodule revision date.
A zero-length string is used if no revision statement
is present in the YANG module or submodule.. 

  
* **schema** `string` - Contains a URL that represents the YANG schema
resource for this module or submodule.

This leaf will only be present if there is a URL
available for retrieval of the schema for this entry.. 

  
* **namespace** `string` - The XML namespace identifier for this module.. 

  
* **feature** `string[]` - List of YANG feature names from this module that are
supported by the server, regardless of whether they are
defined in the module or any included submodule.. 

  
* **[deviation[…]](#/modules-state/module/deviation)** - List of YANG deviation module names and revisions
used by this server to modify the conformance of
the module associated with this entry.  Note that
the same module can be used for deviations for
multiple modules, so the same entry MAY appear
within multiple &#39;module&#39; entries.

The deviation module MUST be present in the &#39;module&#39;
list, with the same name and revision values.
The &#39;conformance-type&#39; value will be &#39;implement&#39; for
the deviation module.. 

  
* **conformance-type** `enumeration` - Indicates the type of conformance the server is claiming
for the YANG module identified by this entry..  *Allowed Values: implement,import* 

  
* **[submodule[…]](#/modules-state/module/submodule)** - Each entry represents one submodule within the
parent module.. 







## <a name="/modules-state/module/deviation"></a>/modules-state/module={name}/deviation={name,revision}/
List of YANG deviation module names and revisions
used by this server to modify the conformance of
the module associated with this entry.  Note that
the same module can be used for deviations for
multiple modules, so the same entry MAY appear
within multiple &#39;module&#39; entries.

The deviation module MUST be present in the &#39;module&#39;
list, with the same name and revision values.
The &#39;conformance-type&#39; value will be &#39;implement&#39; for
the deviation module.


  
* **name** `string` - The YANG module or submodule name.. 

  
* **revision** `string` - The YANG module or submodule revision date.
A zero-length string is used if no revision statement
is present in the YANG module or submodule.. 







## <a name="/modules-state/module/submodule"></a>/modules-state/module={name}/submodule={name,revision}/
Each entry represents one submodule within the
parent module.


  
* **name** `string` - The YANG module or submodule name.. 

  
* **revision** `string` - The YANG module or submodule revision date.
A zero-length string is used if no revision statement
is present in the YANG module or submodule.. 

  
* **schema** `string` - Contains a URL that represents the YANG schema
resource for this module or submodule.

This leaf will only be present if there is a URL
available for retrieval of the schema for this entry.. 







