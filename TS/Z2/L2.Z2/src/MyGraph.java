import org.jgrapht.alg.ConnectivityInspector;
import org.jgrapht.alg.DijkstraShortestPath;
import org.jgrapht.graph.DefaultWeightedEdge;
import org.jgrapht.graph.SimpleWeightedGraph;
import java.util.Random;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by Bartosz on 25.04.2017
 */

public class MyGraph {

    private static final int SIZEOFPACKET = 32;
    private SimpleWeightedGraph<Integer, DefaultWeightedEdge> graph = new SimpleWeightedGraph<Integer, DefaultWeightedEdge>(DefaultWeightedEdge.class);
    private List<Integer> vertices;
    private List<MEdge> edges;
    private List<Integer> packets;

    MyGraph() {
        edges = new ArrayList<>();
        vertices = new ArrayList<>();
        packets = new ArrayList<>();
    }

    void addVertex(int vertex) {
        graph.addVertex(vertex);
        if (!(vertices.contains(vertex))) {
            vertices.add(vertex);
        }
    }

    void addEdge(int vertex1, int vertex2, int capacity) {

        if(graph.containsEdge(vertex1, vertex2) || graph.containsEdge(vertex2, vertex1))
        {
            return;
        }
        int v1 = vertex1;
        int v2 = vertex2;
        MEdge edge = new MEdge(v1, v2, capacity);
        if (edges.contains(edge)) {
            return;
        }
        edges.add(edge);
        DefaultWeightedEdge e = graph.addEdge(v1, v2);
        graph.setEdgeWeight(e, 0.9);
    }

    boolean sendPacket(int v1, int v2, int numberOfPackets) {

        int amount = numberOfPackets * SIZEOFPACKET;
        List path = DijkstraShortestPath.findPathBetween(graph, v1, v2);

        for (Object aPath : path) {
            String[] parts = aPath.toString().split(" ");
            int vertice1 = Integer.parseInt(parts[0].substring(1));
            int vertice2 = Integer.parseInt(parts[2].substring(0, parts[2].length() - 1));
            MEdge e = null;
            for (MEdge edge : edges) {
                if (edge.getV1() == vertice1 && edge.getV2() == vertice2) {
                    e = edge;
                }
            }
            e.setNumberOfPackets(numberOfPackets);
            if (!e.checkOverload(amount)) {
                return false;
            }
        }
        packets.add(numberOfPackets);
        return true;
    }

    void deleteEdge(int i) {
        MEdge e = edges.get(i);
        int v1 = e.getV1();
        int v2 = e.getV2();

        graph.removeEdge(v1,v2);
        edges.remove(i);

    }

    boolean isConnected() {
        ConnectivityInspector<Integer, DefaultWeightedEdge> connectivityInspector
                = new ConnectivityInspector<>(graph);
        return connectivityInspector.isGraphConnected();
    }

    void simulateApocalypse(double p) {
        for(int i = 1; i < edges.size(); i++) {
            Random r = new Random();
            double q = (r.nextInt(10001) * 1.0) / (10000);
            if(p < q) {
                deleteEdge(i);
            }
        }
    }

    double avgPacketDelay(){
        int G = 0;
        double sum = 0;

        for(MEdge edge : edges) {
            sum += edge.getNumberOfPackets() *1.0 /(edge.getCapacity() * 1.0 / SIZEOFPACKET * 1.0 - edge.getNumberOfPackets());
        }

        for(int i : packets) {
            G += i;
        }

        return (1.0/G*1.0) * sum;
    }

    void printFlow() {
        System.out.println("Wyswietlenie wynikow dla krawedzi");
        for (MEdge edge : edges) {
            System.out.print(edge.getV1()+" - "+edge.getV2()+" Przplyw: "+edge.getFlow()+ " c(v) = " + edge.getCapacity()+ " " + "a(v) = " + edge.getNumberOfPackets()+" ");
            if(edge.getFlow() >= edge.getCapacity()) {
                System.out.print("Przeciacenie krawedzi !!!!");
            }
            System.out.print("\n");
        }
    }
}
