UNDER CONSTRUCTION
============

Experimental Server in support of the
[Idionautic](https://github.com/navicore/idionautic-agent/actions) application
performance monitoring project.

This repository holds the POC server that the [Idionautic Agent](https://github.com/navicore/idionautic-agent/actions) posts
observations to.  

Initially data is logged to an embedded sqlite DB.  The data
and an analysis of the data can be retrieved via API.

Eventually, the server will support forwarding to OpenTelemetry gateways as well
as an experimental wide-record analytics store designed to support columnar
storage tolerant of high-cardinality values in support of observability.

-------------
TODO: openapi spec
