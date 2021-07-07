Feature: access.smi-spec.io
  In order to test the TrafficTarget
  As a developer
  I need to ensure the specification is accepted by the server

  Scenario: Apply alpha3 TrafficTarget
    Given the server is running
    When I create the following resource
    ```
      apiVersion: access.smi-spec.io/v1alpha3
      kind: TrafficTarget
      metadata:
        name: path-specific
        namespace: default
      spec:
        destination:
          kind: ServiceAccount
          name: service-a
          namespace: default
        rules:
        - kind: TCPRoute
          name: the-routes
        - kind: HTTPRouteGroup
          name: the-routes
          matches:
          - metrics
        sources:
        - kind: ServiceAccount
          name: prometheus
          namespace: default
    ```
    Then I expect "UpsertTrafficTarget" to be called 1 time
  
  Scenario: Apply alpha2 TrafficSplitter
    Given the server is running
    When I create the following resource
    ```
      apiVersion: access.smi-spec.io/v1alpha2
      kind: TrafficTarget
      metadata:
        name: path-specific
        namespace: default
      spec:
        destination:
          kind: ServiceAccount
          name: service-a
          namespace: default
          port: 8080
        rules:
        - kind: HTTPRouteGroup
          name: the-routes
          matches:
          - metrics
        sources:
        - kind: ServiceAccount
          name: prometheus
          namespace: default
    ```
    Then I expect "UpsertTrafficTarget" to be called 1 time
 # 
  Scenario: Apply alpha1 TrafficSplitter
    Given the server is running
    When I create the following resource
    ```
      apiVersion: access.smi-spec.io/v1alpha1
      kind: TrafficTarget
      metadata:
        name: path-specific
        namespace: default
      destination:
        kind: ServiceAccount
        name: service-a
        namespace: default
        port: 8080
      specs:
      - kind: HTTPRouteGroup
        name: the-routes
        matches:
        - metrics
      sources:
      - kind: ServiceAccount
        name: prometheus
        namespace: default
    ```
    Then I expect "UpsertTrafficTarget" to be called 1 time

  