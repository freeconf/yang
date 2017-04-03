var notify = {};

//
// Supports websocket instance from web browser or ws node package.
//
// For Node
//  var ws = require('ws');
//  var driver = new ws('ws://host:99/path','', {origin:'host:99'});
//  var n = new notify.handler(driver);
//
// For web browser
//  var driver = new WebSocket('ws://host:99/path');
//  var n = new notify.handler(driver);
//
notify.handler = function(driver) {
	this.listeners = {};
	this.driver = driver;
	this.lastErr = null;
	this.isConnected = false;
	this.onClose = null;
	this.onDisconnect = null;
	this.OnError = null;

	var self  = this;

	this.decode = function(data) {
		var	packet = JSON.parse(data);
	    if ('payload' in packet) {
			packet.payload = self.decodePayload(packet);
	    }
		return packet;
	};

	this.decodePayload = function(packet) {		
		var s;
		if (typeof Buffer == 'function') {
			// nodejs
			s = new Buffer(packet.payload, 'base64').toString("utf8");
		} else {
			// browser
			s = atob(packet.payload);
		}

		if (packet.type === "error") {
			return s;
		}
		return JSON.parse(s);
	};

	this.subscribe = function(listener) {
        var packet = {
            op:'+',
            path:listener.path,
            group:listener.group
        }
        var s = JSON.stringify(packet);
		self.driver.send(s, self.onDriverErr);
	};

	this.resubscribe = function() {
	    for (prop in self.listeners) {
	        if (self.listeners.hasOwnProperty(prop)) {
                self.subscribe(self.listeners[prop]);
	        }
	    }
	};

	this.close = function() {
		if ('terminate' in self.driver) {
			// I had to hack ws library to get this quick timeout to work.  I wouldn't have done this except
			// this really only affects test usage.  web browser will not need to terminate quickly. If
			// your tests do not exit quickly, it's simply because you don't have my edit to the ws lib.
			// file: ws/lib/WebSocket.js
			//
			// Original:
            //    WebSocket.prototype.terminate = function terminate() {
			// Edit:
            //    WebSocket.prototype.terminate = function terminate(closeTimeout) {
			//
			self.driver.terminate(1);
		}
	};

	this.fire = function(packet) {
	    var id = this.buildId(packet.group, packet.path, packet.module, packet.device);
	    var listener = self.listeners[id];
        if (listener) {
            var error, data;
            if (packet.type === "error")  {
                listener.f(null, packet.payload);
            } else {
                listener.f(packet.payload, null);
            }
        }
	};

	this.onDriverErr =  function(err) {
		if (err) {
			console.log("error", arguments);
			self.lastErr = err;
		}
	};

	this.onDriverOpen = function(conn) {
		self.isConnected = true;
		self.resubscribe();
	};

	this.onDriverClose = function() {
		self.isConnected = false;
	};

	this.onDriverMessage = function(msg) {
		var packet = null;
		if (typeof msg === 'string') {
			packet = self.decode(msg);
		} else {
			packet = self.decode(msg.data);
		}

        self.fire(packet);
	};

	if (driver.on) {
		// Node - ws
		driver.on('error', this.onDriverErr);
		driver.on('open', this.onDriverOpen);
		driver.on('close', this.onDriverClose);
		driver.on('message', this.onDriverMessage);
	} else {
		// web browser
		driver.onerror = this.onDriverErr;
		driver.onopen = this.onDriverOpen;
		driver.onclose = this.onDriverClose;
		driver.onmessage = this.onDriverMessage;
	}

	this.buildId = function(group, path, moduleName, device) {
		var id = group + '|' + path + '|' + moduleName;
		if (typeof device != "undefined") {
			id += '|' + device;
		}
		return id;
	}

	this.on = function(group, path, moduleName, f, device) {
       	var listener = {
       		group : group,
       		path: path,
			module: moduleName,
			device: device,			
       		id : this.buildId(group, path, moduleName, device),
       		f : f
       	};
		self.listeners[listener.id] = listener;
		if (self.isConnected) {
			self.subscribe(listener);
		}
       	return listener;
    };

    this.off = function(listener) {
	    var listener = self.listeners[listener.id];
        if (listener) {
            delete self.listeners[listener.id];
            if (self.isConnected) {
                var packet = {
                    op:'-',
                    path:listener.path,
					module:listener.module,
					device:listener.device,					
                    group:listener.group
                }
                self.driver.send(JSON.stringify(packet), self.onDriverErr);
            }
    	}
    };

	return this;
};

// For node CommonJS compatibility, ignored in web browser
if (typeof module != 'undefined') {
	module.exports = {
		handler : notify.handler
	};
}