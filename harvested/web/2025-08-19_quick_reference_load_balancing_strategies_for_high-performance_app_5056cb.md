---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:12:47.462127'
profile: quick_reference
source: https://nginx.org/en/docs/http/load_balancing.html
topic: Load Balancing Strategies for High-Performance Applications
---

# Load Balancing Strategies for High-Performance Applications

Join us on the new [NGINX Community Forum](https://community.nginx.org) to connect with users, discover the latest community activity, and troubleshoot issues together. 
# [![NGINX](https://nginx.org/img/nginx_logo.png)![NGINX](https://nginx.org/img/nginx_logo_dark.png)](https://nginx.org/)
  * english
  * русский
  * [news](https://nginx.org/news.html)
  * [about](https://nginx.org/en/)
  * [download](https://nginx.org/en/download.html)
  * [security](https://nginx.org/en/security_advisories.html)
  * [documentation](https://nginx.org/en/docs/)
  * [faq](https://nginx.org/en/docs/faq.html)
  * [books](https://nginx.org/en/books.html)
  * [community](https://nginx.org/en/community.html)
  * [enterprise](https://nginx.org/en/enterprise.html)
  * [community forum (new)](https://community.nginx.org)
  * [x.com](https://x.com/nginxorg)
  * [blog](https://blog.nginx.org/)
  * [unit](https://unit.nginx.org/)
  * [njs](https://nginx.org/en/docs/njs/)


## Using nginx as HTTP load balancer
[Load balancing methods](https://nginx.org/en/docs/http/load_balancing.html#nginx_load_balancing_methods)[Default load balancing configuration](https://nginx.org/en/docs/http/load_balancing.html#nginx_load_balancing_configuration)[Least connected load balancing](https://nginx.org/en/docs/http/load_balancing.html#nginx_load_balancing_with_least_connected)[Session persistence](https://nginx.org/en/docs/http/load_balancing.html#nginx_load_balancing_with_ip_hash)[Weighted load balancing](https://nginx.org/en/docs/http/load_balancing.html#nginx_weighted_load_balancing)[Health checks](https://nginx.org/en/docs/http/load_balancing.html#nginx_load_balancing_health_checks)[Further reading](https://nginx.org/en/docs/http/load_balancing.html#nginx_load_balancing_additional_information)  
---  
#### Introduction
Load balancing across multiple application instances is a commonly used technique for optimizing resource utilization, maximizing throughput, reducing latency, and ensuring fault-tolerant configurations. 
It is possible to use nginx as a very efficient HTTP load balancer to distribute traffic to several application servers and to improve performance, scalability and reliability of web applications with nginx. 
#### Load balancing methods
The following load balancing mechanisms (or methods) are supported in nginx: 
  * round-robin — requests to the application servers are distributed in a round-robin fashion, 
  * least-connected — next request is assigned to the server with the least number of active connections, 
  * ip-hash — a hash-function is used to determine what server should be selected for the next request (based on the client’s IP address). 


#### Default load balancing configuration
The simplest configuration for load balancing with nginx may look like the following: 
> ```
http {
  upstream myapp1 {
    server srv1.example.com;
    server srv2.example.com;
    server srv3.example.com;
  }
  server {
    listen 80;
    location / {
      proxy_pass http://myapp1;
    }
  }
}

```

In the example above, there are 3 instances of the same application running on srv1-srv3. When the load balancing method is not specifically configured, it defaults to round-robin. All requests are [ proxied](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass) to the server group myapp1, and nginx applies HTTP load balancing to distribute the requests. 
Reverse proxy implementation in nginx includes load balancing for HTTP, HTTPS, FastCGI, uwsgi, SCGI, memcached, and gRPC. 
To configure load balancing for HTTPS instead of HTTP, just use “https” as the protocol. 
When setting up load balancing for FastCGI, uwsgi, SCGI, memcached, or gRPC, use [fastcgi_pass](https://nginx.org/en/docs/http/ngx_http_fastcgi_module.html#fastcgi_pass), [uwsgi_pass](https://nginx.org/en/docs/http/ngx_http_uwsgi_module.html#uwsgi_pass), [scgi_pass](https://nginx.org/en/docs/http/ngx_http_scgi_module.html#scgi_pass), [memcached_pass](https://nginx.org/en/docs/http/ngx_http_memcached_module.html#memcached_pass), and [grpc_pass](https://nginx.org/en/docs/http/ngx_http_grpc_module.html#grpc_pass) directives respectively. 
#### Least connected load balancing
Another load balancing discipline is least-connected. Least-connected allows controlling the load on application instances more fairly in a situation when some of the requests take longer to complete. 
With the least-connected load balancing, nginx will try not to overload a busy application server with excessive requests, distributing the new requests to a less busy server instead. 
Least-connected load balancing in nginx is activated when the [ least_conn](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#least_conn) directive is used as part of the server group configuration: 
> ```
  upstream myapp1 {
    least_conn;
    server srv1.example.com;
    server srv2.example.com;
    server srv3.example.com;
  }

```

#### Session persistence
Please note that with round-robin or least-connected load balancing, each subsequent client’s request can be potentially distributed to a different server. There is no guarantee that the same client will be always directed to the same server. 
If there is the need to tie a client to a particular application server — in other words, make the client’s session “sticky” or “persistent” in terms of always trying to select a particular server — the ip-hash load balancing mechanism can be used. 
With ip-hash, the client’s IP address is used as a hashing key to determine what server in a server group should be selected for the client’s requests. This method ensures that the requests from the same client will always be directed to the same server except when this server is unavailable. 
To configure ip-hash load balancing, just add the [ip_hash](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#ip_hash) directive to the server (upstream) group configuration: 
> ```
upstream myapp1 {
  ip_hash;
  server srv1.example.com;
  server srv2.example.com;
  server srv3.example.com;
}

```

#### Weighted load balancing
It is also possible to influence nginx load balancing algorithms even further by using server weights. 
In the examples above, the server weights are not configured which means that all specified servers are treated as equally qualified for a particular load balancing method. 
With the round-robin in particular it also means a more or less equal distribution of requests across the servers — provided there are enough requests, and when the requests are processed in a uniform manner and completed fast enough. 
When the [weight](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server) parameter is specified for a server, the weight is accounted as part of the load balancing decision. 
> ```
  upstream myapp1 {
    server srv1.example.com weight=3;
    server srv2.example.com;
    server srv3.example.com;
  }

```

With this configuration, every 5 new requests will be distributed across the application instances as the following: 3 requests will be directed to srv1, one request will go to srv2, and another one — to srv3. 
It is similarly possible to use weights with the least-connected and ip-hash load balancing in the recent versions of nginx. 
#### Health checks
Reverse proxy implementation in nginx includes in-band (or passive) server health checks. If the response from a particular server fails with an error, nginx will mark this server as failed, and will try to avoid selecting this server for subsequent inbound requests for a while. 
The [max_fails](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server) directive sets the number of consecutive unsuccessful attempts to communicate with the server that should happen during [fail_timeout](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server). By default, [max_fails](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server) is set to 1. When it is set to 0, health checks are disabled for this server. The [fail_timeout](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server) parameter also defines how long the server will be marked as failed. After [fail_timeout](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server) interval following the server failure, nginx will start to gracefully probe the server with the live client’s requests. If the probes have been successful, the server is marked as a live one. 
#### Further reading
In addition, there are more directives and parameters that control server load balancing in nginx, e.g. [proxy_next_upstream](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_next_upstream), [backup](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server), [down](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#server), and [keepalive](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#keepalive). For more information please check our [reference documentation](https://nginx.org/en/docs/). 
Last but not least, application load balancing, application health checks, activity monitoring and on-the-fly reconfiguration of server groups are available as part of our paid NGINX Plus subscriptions. 


## Source Information
- URL: https://nginx.org/en/docs/http/load_balancing.html
- Harvested: 2025-08-19T23:12:47.462127
- Profile: quick_reference
- Agent: architect
