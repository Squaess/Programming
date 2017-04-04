import org.jgrapht.graph.DefaultWeightedEdge;
import org.jgrapht.graph.SimpleWeightedGraph;

import java.util.ArrayList;
import java.util.List;

class MGraph {
    private SimpleWeightedGraph<Integer, DefaultWeightedEdge> graph = new SimpleWeightedGraph<Integer, DefaultWeightedEdge>(DefaultWeightedEdge.class);
    private List<MEdge> edges;
    private List<Integer> vertices;

    MGraph() {
        edges = new ArrayList<MEdge>();
        vertices = new ArrayList<Integer>();
    }

    void addVertex(int vertex) {
        graph.addVertex(vertex);
        if(!(vertices.contains(vertex))) {
            vertices.add(vertex);
        }
    }

    void addEdge(int vertex1, int vertex2, int capacity) {
        MEdge edge = new MEdge(vertex1, vertex2, capacity);
        if(edges.contains(edge)) {
            return;
        }
        edges.add(edge);
        graph.addEdge(vertex1,vertex2);
    }

    public String toString(){
        return graph.toString();
    }
}
