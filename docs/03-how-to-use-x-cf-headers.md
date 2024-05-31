---
title: How To Use X-CF Headers
expires_at: never
tags: [routing-release]
---

# How To Use X-CF Headers

| Header | If a client provides this header will that affect routing decisions? | Value | More info |
| -- | -- | -- | -- |
| X-CF-ApplicationID | No | App guid for the app that gorouter proxied the request to. | [This header is always set/overwritten on the request by gorouter when sending traffic to an app.](https://github.com/cloudfoundry/gorouter/blob/46ad6582a5c07ae191004c3b84142a0445f47595/proxy/round_tripper/proxy_round_tripper.go#L239-L241) |
| X-CF-InstanceID | No | Instance guid for the app that gorouter proxied the request to. | [This header is always set/overwritten on the request by gorouter when sending traffic to an app.](https://github.com/cloudfoundry/gorouter/blob/46ad6582a5c07ae191004c3b84142a0445f47595/proxy/round_tripper/proxy_round_tripper.go#L239-L241) |
| X-Cf-Instanceindex | No | Instance index for the app that gorouter proxied the request to. | [This header is always set/overwritten on the request by gorouter when sending traffic to an app.](https://github.com/cloudfoundry/gorouter/blob/46ad6582a5c07ae191004c3b84142a0445f47595/proxy/round_tripper/proxy_round_tripper.go#L239-L241) |
| X-Cf-App-Instance | Yes | Instance index for the app that gorouter proxied the request to. | This is a header that can be provided by the user for routing a specific index. Gorouter does not change this header. You set it in the format: "APP-GUID:APP-INSTANCE-INDEX". If you provide an invalid header (bad format, nonexistent guid/index) then the gorouter will return a 400. See [here](https://docs.cloudfoundry.org/concepts/http-routing.html#app-instance-routing) for more docs. |
| X-Cf-RouterError | No | See possible values [here](https://docs.cloudfoundry.org/adminguide/troubleshooting-router-error-responses.html#gorouter-specific-response-headers). | Gorouter adds this header to the response when a routing error has occured. | 
