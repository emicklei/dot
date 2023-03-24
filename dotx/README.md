## dotx package (dot extensions)

This package contains utilities to create graphs on top of the `emicklei/dot package`.

### Subsystem

The `Subsystem` type can be used to create composition hierarchies like clustering. 

Let's examine this diagram.

![](../doc/TestExampleSubsystemSameGraph.png)

On the most right, you find a node called `subsystem` which is a Subsystem with 2 inputs and 1 output edge.

On the most left, you see the contents of the same `subsystem` with both inputs and an output (point shaped with label).

The `subsystem` contains other nodes, 2 regular nodes (`subcomponent 1` and `subcomponent 2`) and another Subsystem labeled `subsystem2`.

So, `subsystem` is a composition of 3 components and 1 of these components is itself a composition of a component (`subcomponent 3`).

### external option

If you create a Subsystem using the `ExternalGraph` kind then its graph can be exported separatedly from the containing graph. If you visualize such a graph using `SVG` then you can **nagivate into** the subsystems.

![](../doc/TestExampleSubsystemExternalGraph.svg)

And clicking on `subsystem`, your browse will show:

![](../doc/subsystem.svg)

And clicking on `subsystem2`, your browse will show:

![](../doc/subsystem2.svg)

See `subsystem_test.go` for the code of these examples.

### usage pattern

    import (
        "github.com/emicklei/dot"
        "github.com/emicklei/dot/dotx"
    )

    func YourService(parent *dot.Graph) *dotx.Subsystem {
        // external means it exports its own DOT file
        sub := dotx.NewSubsystem("Your Service", parent, dotx.ExternalGraph)

        return sub.Export(func(g *dot.Graph) {
            
            // build the inner graph of the Subsystem
            myComp := g.Node("myComp")
            
            // connect any inputs,outputs
            sub.Input("in", myComp)
        })
    }