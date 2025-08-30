## 路由配置

### 1.跨域设置
gateway.globalcors.corsConfigurations 跨域配置
gateway:  # 跨域
      globalcors:
        corsConfigurations:
          '[/**]':
            allowedHeaders: "*"
            allowedOrigins: "*"
            allowedMethods:
              - GET
                POST
                DELETE
                PUT
                OPTION
### 路由配置 [id 是路由标记， lb://表示负载均衡，service-name为服务发现配置的服务名
gateway.routes
	- id: service-name-route
	  uri: lb://service-name
	  predicates: 【匹配规则】
		- Path=/api/v1/reportService/**

## 常见问题

### 问题情况1 503 SERVICE_UNAVAILABLE
Resolved [NotFoundException: 503 SERVICE_UNAVAILABLE "Unable to find instance for athena-web"] for HTTP GET 

在Spring Cloud Gateway中，当使用服务发现（如Nacos、Eureka等）时，需要负载均衡器来解析服务名称并路由到具体的服务实例。从Spring Cloud 2020.x（即Ilford版本）开始，Ribbon被移除，
取而代之的是Spring Cloud LoadBalancer，但它不再是默认依赖。