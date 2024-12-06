Choose to implement (poorly, most likely) a DAG. This is the first time I've ever used any graphing in programming.  

Pages are Nodes  
Each page number (within the update) is a node in the graph. Pages 47, 53, 97... the nodes are also 47, 53, 97...  

Edges are page - to page lines  
An edge exists between two nodes (pages) if there's a rule connecting them. The rule 47|53 means there is a *directed edge* from node (page) 47 to node (page) 53 (47 → 53). This edge implies that page 47 must appear before page 53 in any valid order. Directed edges are representative of the rule's being applied.  

Acyclic  
The graph **must not** contain cycles because a cycle in this puzzle is a circular dependency that cannot be resolved:  47 → 53 and 53 → 47 cannot both be true  

Why a DAG  
I wanted/needed to learn graphs, even if it's not the ideal or maybe even a good approach to today's problem.  
The DAG helps:  
  - Validate Order: For a given update (list of pages), check if the pages appear in a sequence consistent with the edges in the graph. Consistency means the pages are in correct order already.  
  - page moving: If an update is not in the correct order, a topological sort of the DAG determines the page order with the rules in place  

Wait, Topological Sorting?  
  - Definition (rando internet blog copy/paste here but I liked it): A topological sort of a DAG is a linear ordering of its nodes such that for every directed edge u → v, node u comes before node v in the ordering.  

  - for the puzzle: for each update, construct a subgraph of the DAG that includes only the pages in the update (is this efficient? Who knows!). Perform a topological sort on this subgraph to reorder the pages correctly. Happens in reorderInvalidUpdate func. "visited" indexes which pages are done. tempStack is a guard against invalid cycle (again a cycle should be impossible for this problem). visitNode (ugh, names) handles the node (page) ordering itself  

To determine if an update is in the correct order:  
- loop the graph and for each x|y in the update, check if the order respects the directed edges (the rule). If it doesnt the update is invalid.

If the update is invalid:  
  - find the subgraph containing only the pages in the update.  
  - Perform another topological sort on this subgraph (Expensive? Almost certainly. Not getting into Google with this trash).  
  - Replace the update with the fixed one 
