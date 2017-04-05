
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.Scanner;


public class Network {
	private static final String FILENAME = "./Dane.txt";
	private static MGraph graph = new MGraph();

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