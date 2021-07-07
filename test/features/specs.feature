Feature: specs.smi-spec.io
  In order to test the specs.smi-spec.io CRDs
  As a developer
  I need to ensure the specification is accepted by the server

  @specs @HTTPRouteGroup
  Scenario: Apply alpha1 HTTPRouteGroup
    Given the server is running
    When I create the following resource
    ```
      apiVersion: specs.smi-spec.io/v1alpha1
      kind: HTTPRouteGroup
      metadata:
        name: the-routes
      matches:
      - name: metrics
        pathRegex: "/metrics"
        methods:
        - GET
      - name: health
        pathRegex: "/ping"
        methods: ["*"]
    ```
    Then I expect "UpsertHTTPRouteGroup" to be called 1 time
  
  @specs @HTTPRouteGroup
  Scenario: Apply alpha2 HTTPRouteGroup
    Given the server is running
    When I create the following resource
    ```
      apiVersion: specs.smi-spec.io/v1alpha2
      kind: HTTPRouteGroup
      metadata:
        name: the-routes
      matches:
      - name: metrics
        pathRegex: "/metrics"
        methods:
        - GET
      - name: health
        pathRegex: "/ping"
        methods: ["*"]
    ```
    Then I expect "UpsertHTTPRouteGroup" to be called 1 time
    
  @specs @HTTPRouteGroup
  Scenario: Apply alpha3 HTTPRouteGroup
    Given the server is running
    When I create the following resource
    ```
      apiVersion: specs.smi-spec.io/v1alpha3
      kind: HTTPRouteGroup
      metadata:
        name: the-routes
      spec:
        matches:
        - name: metrics
          pathRegex: "/metrics"
          methods:
          - GET
        - name: health
          pathRegex: "/ping"
          methods: ["*"]
    ```
    Then I expect "UpsertHTTPRouteGroup" to be called 1 time
  
  @specs @HTTPRouteGroup
  Scenario: Apply alpha4 HTTPRouteGroup
    Given the server is running
    When I create the following resource
    ```
      apiVersion: specs.smi-spec.io/v1alpha4
      kind: HTTPRouteGroup
      metadata:
        name: the-routes
      spec:
        matches:
        - name: metrics
          pathRegex: "/metrics"
          methods:
          - GET
          headers:
            x-debug: "1"
        - name: health
          pathRegex: "/ping"
          methods: ["*"]
    ```
    Then I expect "UpsertHTTPRouteGroup" to be called 1 time


  @specs @TPCRoute
  Scenario: Apply alpha1 TCPRoute
    Given the server is running
    When I create the following resource
    ```
      apiVersion: specs.smi-spec.io/v1alpha1
      kind: TCPRoute
      metadata:
        name: tcp-route
    ```
    Then I expect "UpsertTCPRoute" to be called 1 time
  
  @specs @TPCRoute
  Scenario: Apply alpha2 TCPRoute
    Given the server is running
    When I create the following resource
    ```
      apiVersion: specs.smi-spec.io/v1alpha2
      kind: TCPRoute
      metadata:
        name: tcp-route
    ```
    Then I expect "UpsertTCPRoute" to be called 1 time
  
  @specs @TPCRoute
  Scenario: Apply alpha3 TCPRoute
    Given the server is running
    When I create the following resource
    ```
      apiVersion: specs.smi-spec.io/v1alpha3
      kind: TCPRoute
      metadata:
        name: tcp-route
      spec: {}
    ```
    Then I expect "UpsertTCPRoute" to be called 1 time
  
  @specs @TPCRoute
  Scenario: Apply alpha4 TCPRoute
    Given the server is running
    When I create the following resource
    ```
      apiVersion: specs.smi-spec.io/v1alpha4
      kind: TCPRoute
      metadata:
        name: the-routes
      spec:
        matches:
          ports:
          - 3306
          - 6446
    ```
    Then I expect "UpsertTCPRoute" to be called 1 time
  
  
  @specs @UDPRoute
  Scenario: Apply alpha4 UDPRoute
    Given the server is running
    When I create the following resource
    ```
      apiVersion: specs.smi-spec.io/v1alpha4
      kind: UDPRoute
      metadata:
        name: the-routes
      spec:
        matches:
          ports:
          - 989
          - 990
    ```
    Then I expect "UpsertUDPRoute" to be called 1 time