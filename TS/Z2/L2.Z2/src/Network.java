import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;



public class Network {
	private static final String FILENAME = "./Dane.txt";
	private static final String GRAF_TXT = "./graf.txt";
	private static final String KONFIG = "./config.txt";

    static boolean startTraffic(MyGraph g) {
        try {
            BufferedReader input = new BufferedReader(new FileReader(FILENAME));
            int size = Integer.parseInt(input.readLine());
            String line;
            for(int i = 0; i < size; i++) {

                line = input.readLine();

                String[] parts = line.split(" ");
                int v1 = i+1;
                int v2 = 1;
                for(String s : parts) {
                    if(v1 != v2) {
                        if(!g.sendPacket(v1,v2,Integer.parseInt(s))) {
                            return false;
                        }
                    }
                    v2++;
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        return true;
	}

    private static MyGraph generateGraph() {
	    MyGraph g = new MyGraph();
        try {
            BufferedReader input = new BufferedReader(new FileReader(GRAF_TXT));
            int number_of_vertex  = Integer.parseInt(input.readLine());
            for(int i = 1; i <= number_of_vertex ; i++) {
                g.addVertex(i);
            }
            String line;
            while((line = input.readLine()) != null) {
                String[] edge = line.split(" ");
                edge[0] = edge[0].substring(1);
                edge[1] = edge[1].substring(0,edge[1].length()-1);
                int beg = Integer.parseInt(edge[0]);
                int end = Integer.parseInt(edge[1]);
                int capacity = Integer.parseInt(edge[2]);
                g.addEdge(beg, end, capacity);

            }
        } catch (IOException e) {
            e.printStackTrace();
        }
	    return g;
    }

	public static void main(String[] args) {
        double p = 0;
        double t_max = 0;
        try {
            BufferedReader input = new BufferedReader(new FileReader(KONFIG));
            String line = input.readLine();
            String[] parts = line.split(" ");
            p = Double.parseDouble(parts[0]);
            t_max = Double.parseDouble(parts[1]);

        } catch (IOException e) {
            e.printStackTrace();
        }
        int succes = 0;
        int failures = 0;
        int overload = 0;
        int przeciazenie = 0;

        for (int i = 0; i < 5000; i++) {
            MyGraph graph = generateGraph();
            graph.simulateApocalypse(p);

            if (graph.isConnected()) {
                if (startTraffic(graph)) {
//                    System.out.println(graph.avgPacketDelay());
                    if (graph.avgPacketDelay() < t_max) {
                        succes++;
                    } else overload++;
                } else {
                    przeciazenie++;
//                    graph.printFlow();
                }
            } else failures++;
        }
        System.out.println("Ilosc testow: 5000" );
        System.out.println("Ilosc sukcesow: "+succes);
        System.out.println("Ilosc rozspojnien: "+failures);
        System.out.println("Ilosc przekroczonych opoznien: "+overload);
        System.out.println("Ilosc przekroczonych przeciazen krawedzi "+ przeciazenie);

        double reliability = (5000.0 - (failures + overload + przeciazenie)*1.0)/(5000.0) * 100.0;

        System.out.println("Testowana niezawodnosc sieci wynosi: "+ reliability+" %");

    }

}