import org.jgrapht.alg.ConnectivityInspector;
import org.jgrapht.alg.DijkstraShortestPath;
import org.jgrapht.graph.DefaultWeightedEdge;
import org.jgrapht.graph.SimpleWeightedGraph;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

class MGraph {
    private static final int SIZEOFPACKET = 1500;
    private SimpleWeightedGraph<Integer, DefaultWeightedEdge> graph = new SimpleWeightedGraph<>(DefaultWeightedEdge.class);
    private List<MEdge> edges;
    private List<Integer> vertices;
    private List<Integer> packets;

    MGraph() {
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

    /*
    capacity is maximum value of bits per second
     */
    void addEdge(int vertex1, int vertex2, int capacity) {
        int v1 = vertex1;
        int v2 = vertex2;
        if (v2 < v1) {
            int tmp = v1;
            v1 = v2;
            v2 = tmp;
        }
        MEdge edge = new MEdge(v1, v2, capacity);
        if (edges.contains(edge)) {
            return;
        }
        edges.add(edge);
        graph.addEdge(v1, v2);
    }

    public String toString() {
        return graph.toString();
    }

    boolean sendPacket(int begining, int destination, int numberOfpackets) {
//        if(numberOfpackets==0) return;

        int v1 = begining, v2 = destination;
//        System.out.println("Wysylam "+numberOfpackets+" pakietow z "+ v1+" do "+v2);
        if (v2 < v1) {
            int tmp = v1;
            v1 = v2;
            v2 = tmp;
        }
//        if(v2==v1) {
//            return;
//        }
        int amount = numberOfpackets * SIZEOFPACKET;
        List path = DijkstraShortestPath.findPathBetween(graph, v1, v2);
//        System.out.print("Sciezka  ");
        for (Object i : path) {
            int vertice1 = Character.getNumericValue(i.toString().charAt(1));
            if (Character.getNumericValue(i.toString().charAt(2)) >= 0 && Character.getNumericValue(i.toString().charAt(2)) <= 9) {
                vertice1 = vertice1 * 10 + Character.getNumericValue(i.toString().charAt(1));
            }
            int vertice2 = Character.getNumericValue(i.toString().charAt(5));
            if (Character.getNumericValue(i.toString().charAt(6)) >= 0 && Character.getNumericValue(i.toString().charAt(6)) <= 9) {
                vertice2 = vertice2 * 10 + Character.getNumericValue(i.toString().charAt(6));
            }
            MEdge e = null;
//            System.out.print(vertice1+" - "+vertice2+ " | ");
            for (MEdge edge : edges) {
                if (edge.getV1() == vertice1 && edge.getV2() == vertice2) {
                    e = edge;
                }
            }
            e.setNumberOfPackets(numberOfpackets);
            if (!e.setFlow(amount)) {
                return false;
            }
        }
        packets.add(numberOfpackets);
        return true;
//        System.out.print("\n");
    }

    double avgPacketDeley(){
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

    private void deleteEdge(int index){
        MEdge edge = edges.get(index);
        int vertice1 = edge.getV1();
        int vertice2 = edge.getV2();
//        System.out.println("Usuwam krawedz "+vertice1+" "+vertice2);
        edges.remove(edge);
        graph.removeEdge(vertice1,vertice2);
    }

    void test(double p) {
        for(int i = 0; i < edges.size(); i++) {
            Random r = new Random();
            double q = (r.nextInt(10001) * 1.0) / (10000);
            if(p < q) {
                deleteEdge(i);
            }
        }
    }

    boolean isConsistent(){
        ConnectivityInspector inspector = new ConnectivityInspector(graph);
        if(!inspector.isGraphConnected()) {
            return false;
        } else return true;
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
