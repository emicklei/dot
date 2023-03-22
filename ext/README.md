## dot/ext package

This package contains utilities to create graphs on top of the `emicklei/dot package`.

### Subsystem

The `Subsystem` type can be used to create abstraction hierarchies like but not the same as clustering. 

Let's examine this diagram.

![](../doc/TestExampleSubsystemSameGraph.png)

On the most right, you find a node called `subsystem` which is a Subsystem with 2 inputs and 1 output edge.

On the most left, you find a cluster of the same `subsystem` with both inputs and an output (point shaped with label).

The `subsystem` cluster contains other nodes, 2 regular nodes (`subcomponent 1` and `subcomponent 2`) and another Subsystem labelled `subsystem2`.

So, `subsystem` is a composition of 3 components and 1 of these components is itself a composition of a component (`subcomponent 3`).

### external option

If you create a Subsystem using the `ExternalGraph` kind then its graph can be exported separatedly from the containing graph. If you visualize such a graph using `SVG` then you can **nagivate into** the subsystems.

![](../doc/TestExampleSubsystemExternalGraph.svg)

See `subsystem_test.go` for the code of these examples.