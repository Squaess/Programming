
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.Scanner;


public class Network {
	private static final String FILENAME = "./Dane.txt";
	private static MGraph graph = new MGraph();
//private static mojGraph gen() {
//		mojGraph graph = new mojGraph();
//
//		graph.addVertex("v1");
//		graph.addVertex("v2");
//		graph.addVertex("v3");
//		graph.addVertex("v4");
//		graph.addVertex("v5");
//		graph.addVertex("v6");
//		graph.addVertex("v7");
//		graph.addVertex("v8");
//		graph.addVertex("v9");
//		graph.addVertex("v10");
//
//		graph.addEdge("v1", "v2", 1000);
//		graph.addEdge("v2", "v3", 1000);
//		graph.addEdge("v3", "v4", 1000);
//		graph.addEdge("v4", "v5", 1000);
//		graph.addEdge("v5", "v6", 1000);
//		graph.addEdge("v6", "v7", 1000);
//		graph.addEdge("v7", "v8", 1000);
//		graph.addEdge("v8", "v9", 1000);
//		graph.addEdge("v9", "v10", 1000);
//		graph.addEdge("v10", "v1", 1000);
//		graph.addEdge("v1", "v9", 4000);
//		graph.addEdge("v2", "v9", 4000);
//		graph.addEdge("v3", "v9", 4000);
//		graph.addEdge("v3", "v8", 1000);
//		graph.addEdge("v3", "v7", 4000);
//		graph.addEdge("v4", "v7", 4000);
//		graph.addEdge("v5", "v7", 4000);
//
//		graph.clearConnections();
//		graph.sendPacket("v1", "v10", 500);
//		graph.sendPacket("v4", "v8", 300);
//		graph.sendPacket("v7", "v10", 200);
//		graph.sendPacket("v2", "v8", 100);
//		graph.sendPacket("v4", "v8", 300);
//
//		return graph;
//	}
//
//	private static void testNetwork(int N, double p, double Tmax)
//	{
//		int success = 0;
//		int failure = 0;
//		int disconnect = 0;
//
//		for(int n=0; n<=N; n++)
//		{
//			mojGraph testGraph = gen();
//
//			testGraph.test(p);
//			if(testGraph.isItConnected())
//			{
//				if(testGraph.getDelay() < Tmax)
//				success++;
//				else
//				failure++;
//			}
//			else
//			{
//				disconnect++;
//			}
//		}
//
//		System.out.println("Ilo�c pr�b: " + N);
//		System.out.println("Ilo�c sukces�w: " + success);
//		System.out.println("Ilo�c pora�ek: " + failure);
//		System.out.println("Ilo�c rozsp�jnie�: " + disconnect);
//
//		double reliability = (N*1.0 - (failure + disconnect)*1.0)/(N * 1.0) * 100.0;
//
//		System.out.println("Niezawodno�c: " + reliability + "%");
//	}
    //Funkcja generujaca siec
	private static void gen() {
	    //Dodajemy 10 wierzcholkow
		for (int i = 1; i <= 10; i++) {
		 	graph.addVertex(i);
		 }

		for (int i = 1; i < 10; i++) graph.addEdge( i, i + 1, 170000);
            graph.addEdge(1,10,190000);
		for (int i = 1; i <= 5; i=i+2) {
			graph.addEdge( i,i+2, 117000);
		}
		graph.addEdge(2,4,170000);
        graph.addEdge(4,6,170000);
        graph.addEdge(6,8,170000);
        graph.addEdge(1,8,170000);
	}

	private static void genNumberOfPackets() {
	    //Czytamy z pliku zapisanym pod FILENAME macierz nastezen NxN
        try {
            BufferedReader input = new BufferedReader(new FileReader(FILENAME));
            int size = Integer.parseInt(input.readLine());
//            System.out.println(size);
            String line;
            for(int i = 0; i < size; i++) {
                line = input.readLine();
//                System.out.println(line);
                String[] parts = line.split(" ");
                int v1 = i+1;
                int v2 = 1;
                for(String s : parts) {
                    if(v1 != v2) {
                        graph.sendPacket(v1,v2,Integer.parseInt(s));
                    }
                    v2++;
                }
            }

        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private static void testGenTraffic(MGraph test) {
        try {
            BufferedReader input = new BufferedReader(new FileReader(FILENAME));
            int size = Integer.parseInt(input.readLine());
//            System.out.println(size);
            String line;
            for(int i = 0; i < size; i++) {
                line = input.readLine();
//                System.out.println(line);
                String[] parts = line.split(" ");
                int v1 = i+1;
                int v2 = 1;
                for(String s : parts) {
                    if(v1 != v2) {
                        test.sendPacket(v1,v2,Integer.parseInt(s));
                    }
                    v2++;
                }
            }

        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private static MGraph genTestGraph(){
	    MGraph ret = new MGraph();
        for (int i = 1; i <= 10; i++) {
            ret.addVertex(i);
        }

        for (int i = 1; i < 10; i++) ret.addEdge( i, i + 1, 170000);
        ret.addEdge(1,10,190000);
        for (int i = 1; i <= 5; i=i+2) {
            ret.addEdge( i,i+2, 117000);
        }
        ret.addEdge(2,4,170000);
        ret.addEdge(4,6,170000);
        ret.addEdge(6,8,170000);
        ret.addEdge(1,8,170000);

        return ret;
    }

    private static void test(double p, double t_max) {
        int succes = 0;
        int failures = 0;
        int overload = 0;

        for(int i = 0; i < 1000; i++) {
            MGraph test = genTestGraph();
            testGenTraffic(test);
            test.test(p);
            if(test.isConsistent()) {
                if(test.avgPacketDeley()< t_max) {
                    succes++;
                } else overload++;
            } else failures++;
        }

        System.out.println("Ilosc sukcesow: "+succes);
        System.out.println("Ilosc rozspojnien: "+failures);
        System.out.println("Ilosc przekroczonych opoznien: "+overload);

        double reliability = (1000.0 - (failures + overload)*1.0)/(1000.0) * 100.0;

        System.out.println("Niezawodnosc sieci wynosi: "+ reliability+" %");
    }

	public static void main(String[] args) {
//		System.out.println("T = " + gen().getDelay());
//		testNetwork(50000, 0.9, -0.001);
		Scanner input = new Scanner(System.in);
		System.out.println("Podaj prawdopodobienstwo p: ");
		double p = input.nextDouble();
		System.out.println("Podaj maksymalne opoznienie: ");
		double t_max = input.nextDouble();
        gen();
		System.out.println("Nasz graf to "+graph.toString());
        genNumberOfPackets();
        graph.printFlow();
        System.out.println("Srednie opoznienie pakietu : " + graph.avgPacketDeley());
        System.out.println("Badanie niezawodności sieci ...");
        test(p,t_max);

	}
}