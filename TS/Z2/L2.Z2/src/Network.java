
import org.jgrapht.graph.DefaultWeightedEdge;
import org.jgrapht.graph.SimpleWeightedGraph;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class Network {
	private static final String FILENAME = ".\\Dane.txt";
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

	private static void gen() {
		for (int i = 1; i <= 10; i++) {
		 	graph.addVertex(i);
		 }

		for (int i = 1; i < 10; i++) graph.addEdge( i, i + 1, 1000);
		for (int i = 1; i <= 3; i++) {
			graph.addEdge( i,i+2, 1000);
			graph.addEdge( i, i+5, 1000);
		}
	}
	public static void main(String[] args) {
//		System.out.println("T = " + gen().getDelay());
//		testNetwork(50000, 0.9, -0.001);
		try {
			BufferedReader input = new BufferedReader(new FileReader(FILENAME));
			String line;
			while((line = input.readLine()) != null) {
				System.out.println(line);
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
		gen();
		System.out.println(graph.toString());
	}
}