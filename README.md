# sigsci-syslog-client

## Description

This simple utility will take new lines in the SigSci Agent log and effectively forward it to a syslog server.

## Recommended SigSci agent.conf settings

I recommend the additional following settings in the /etc/sigsci/agent.conf. This generally will be the most useful data for sending to the syslog.

````
log-out = "/var/log/sigsci.log"
debug-log-web-inputs = 1
debug-log-web-outputs = 1
````

## Usage

### sigsci-syslog

**Example:** `./sigsci-syslog -P tcp -p 514 -s localhost -l /var/log/sigsci.log`

| flag | long flag | type | description|
| ---- | --------- | ---- | ---------- |
| -P | --port | string | Syslog port to connect to, default 514 (default "514") |
|  -p | --protocol | string | Syslog protocol (tcp/udp) (default "udp") |
|  -s | --server | string | Syslog server fqdn (eg localhost, 127.0.001) (default "localhost") |
|  -l | --sigsci-log | string | Path to the SigSci Agent log (default "/var/log/sigsci.log") |

### sigsciAgentLogSize.sh

**Example:** `./sigsciAgentLogSize.sh`

The SigSci Agent log can grow very fast if there is a high number of RPS (Requests Per Second). The shell script will blank out the file when it grows past a certain size. Since the log entries would have been forwarded to syslog there is no point to keep it in the original log file. 

Change the following to match your settings:

`file=/var/log/sigsci.log` (Path to sigsci agent log)
`maxSize=180385` (In bytes)

**Example Output:**

/var/log/syslog
````
Feb 28 19:52:27 hostname sigsci-agent[7099]: 2018/02/28 19:52:27 2018/02/28 19:31:12.630661 RPC.PreRequest Request: {"RequestID":"","ModuleVersion":"sigsci-module-apache 1.6.0","ServerVersion":"Apache/2.4.18 (Ubuntu) SVN/1.9.3 OpenSSL/1.0.2g mod_perl/2.0.9 Perl/v5.22.1","ServerFlavor":"prefork","ServerName":"","Timestamp":1519875072,"NowMillis":1519875072629,"RemoteAdhostname":"123.123.123.123","Method":"PROPFIND","Scheme":"https","URI":"/remote.php/dav/files/username/FOLDER","Protocol":"HTTP/1.1","TLSProtocol":"TLSv1.2","TLSCipher":"ECDHE-RSA-AES256-GCM-SHA384","ResponseCode":0,"ResponseMillis":0,"ResponseSize":0,"HeadersIn":[["Depth","0"],["Authorization","Basic TOKEN"],["User-Agent","Mozilla/5.0 (Windows) mirall/2.3.4 (build 8624)"],["Accept","*/*"],["Content-Type","text/xml; charset=utf-8"],["Cookie","oc_sessionPassphrase=passphrase; oc3a3f7b1dee=valuehere"],["Content-Length","105"],["Connection","Keep-Alive"],["Accept-Encoding","gzip, deflate"],["Accept-Language","en-US,*"],["Host","www.server.com"]],"PostBody":"\u003c?xml version=\"1.0\" ?\u003e\n\u003cd:propfind xmlns:d=\"DAV:\"\u003e\n  \u003cd:prop\u003e\n    \u003cd:getetag/\u003e\n  \u003c/d:prop\u003e\n\u003c/d:propfind\u003e\n"}
Feb 28 19:52:27 hostname sigsci-agent[7099]: 2018/02/28 19:52:27 2018/02/28 19:31:12.632145 RPC.PreRequest Response: {"WAFResponse":200}
Feb 28 20:57:48 hostname sigsci-agent[26267]: 2018/02/28 20:57:48 2018/02/28 20:57:48.914167 Started RPC listener on "unix:/var/run/sigsci.sock"
Feb 28 20:57:48 hostname sigsci-agent[26267]: 2018/02/28 20:57:48 2018/02/28 20:57:48.914512 Reading configuration file "/etc/sigsci/agent.conf"
Feb 28 20:57:48 hostname sigsci-agent[26267]: 2018/02/28 20:57:48 2018/02/28 20:57:48.915652 Started legacy RPC listener on "unix:/tmp/sigsci-lua"
Feb 28 20:57:48 hostname sigsci-agent[26267]: 2018/02/28 20:57:48 2018/02/28 20:57:48.916119 RPC server shutting down
Feb 28 20:57:48 hostname sigsci-agent[26267]: 2018/02/28 20:57:48 2018/02/28 20:57:48.916242 RPC NetOP error: use of closed network connection
Feb 28 20:57:48 hostname sigsci-agent[26267]: 2018/02/28 20:57:48 2018/02/28 20:57:48.916119 RPC server shutting down
Feb 28 20:57:48 hostname sigsci-agent[26267]: 2018/02/28 20:57:48 2018/02/28 20:57:48.916321 RPC NetOP error: use of closed network connection
Feb 28 20:57:48 hostname sigsci-agent[26267]: 2018/02/28 20:57:48 2018/02/28 20:57:48.916369 Starting RPC server with RPC-v0
Feb 28 20:57:48 hostname sigsci-agent[26267]: 2018/02/28 20:57:48 2018/02/28 20:57:48.916637 Started legacy RPC listener on "unix:/tmp/sigsci-lua"
Feb 28 20:57:48 hostname sigsci-agent[26267]: 2018/02/28 20:57:48 2018/02/28 20:57:48.916721 Started RPC listener on "unix:/var/run/sigsci.sock"
````
