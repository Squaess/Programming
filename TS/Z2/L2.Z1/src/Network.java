import java.util.Random;

import org.jgrapht.alg.ConnectivityInspector;
import org.jgrapht.graph.DefaultWeightedEdge;
import org.jgrapht.graph.SimpleWeightedGraph;

public class Network {
	/*
	 * liczba wierzcho�k�w w grafie
	 */
	private static int N = 20;
	private static int liczbaTestow = 100000;
	
	private Network() {}
	
	/*
	 * generator pierwszej sieci
	 */
	public static SimpleWeightedGraph<String, DefaultWeightedEdge> generate(int numberOfVertices) {
		SimpleWeightedGraph<String, DefaultWeightedEdge> graph = new SimpleWeightedGraph<String, DefaultWeightedEdge>(DefaultWeightedEdge.class);
		for(int i = 1; i <= numberOfVertices; i++) {
			String v = "v" + Integer.toString(i);
			graph.addVertex(v);
		}
		
		for(int i = 1; i< numberOfVertices; i++) {
			String v1 = "v" + Integer.toString(i);
			String v2 = "v" + Integer.toString(i+1);
			double dur = 0.95;
			
			DefaultWeightedEdge edge = graph.addEdge(v1, v2);
			graph.setEdgeWeight(edge, dur);
		}
		
		return graph;
	
	}
	/*
	 * zastosowanie metody MonteCarlo
	 */
	public static void test1() {
		
		int zle = 0;
		
		for(int i = 0; i < liczbaTestow; i++) {
			if(!testNetwork(generate(N))) zle++; 
		}
		double p = (liczbaTestow - zle)/(liczbaTestow * 1.0) * 100.0;
		System.out.println("Testowali�my Siec = "+generate(N).toString());
		System.out.println("Szacowana niezawodno�� wynosi = "+ p + " %");
	}
	
	/*
	 * generowanie drugiej sieci z dodan� jedn� kraw�dzi� mi�dzy v1 a v20
	 */
			
	public static SimpleWeightedGraph<String, DefaultWeightedEdge> gen2() {
		SimpleWeightedGraph<String, DefaultWeightedEdge> graph = generate(N);
		//dodanie krawedzi
		DefaultWeightedEdge edge = graph.addEdge("v1", "v20"); 
		graph.setEdgeWeight(edge, 0.95);
		
		return graph;
		
	}
	/*
	 * testowanie sieci nr2
	 */
	public static void test2() {
		int zle = 0;

		for(int i = 0; i < liczbaTestow; i++) {
			if(!testNetwork(gen2())) zle++; 
		}
		double p = (liczbaTestow - zle)/(liczbaTestow * 1.0) * 100.0;
		System.out.println("Testowali�my siec nr 2 (siec pierwsza z dodana krawedzia {v1,v20})");
		System.out.println("Szacowana niezawodno�� wynosi = "+ p + " %");
	}
	/*
	 * generowanie sieci nr 3 (taka sama jak 2 + krawedzie 1-10 i 5-15 z odpowiednimi wartosciami)
	 */
	public static SimpleWeightedGraph<String, DefaultWeightedEdge> gen3() {
		SimpleWeightedGraph<String, DefaultWeightedEdge> graph = gen2();
		
		DefaultWeightedEdge edge = graph.addEdge("v1", "v10"); 
		graph.setEdgeWeight(edge, 0.8);
		
		DefaultWeightedEdge edge2 = graph.addEdge("v5", "v15"); 
		graph.setEdgeWeight(edge2, 0.7);
		
		return graph;
	}
	/*
	 * testowanie sieci nr3
	 */
	public static void test3(){
		int zle = 0;

		for(int i = 0; i < liczbaTestow; i++) {
			if(!testNetwork(gen3())) zle++; 
		}
		double p = (liczbaTestow - zle)/(liczbaTestow * 1.0) * 100.0;
		System.out.println("Testowali�my siec nr 3 (siec druga z dodanymi krawedziami {v1,v10}, {v5,v15} o odpowiednich wartosciach 0.8 0.7 )");
		System.out.println("Szacowana niezawodno�� wynosi = "+ p + " %");
	}
	/*
	 * generowanie sieci nr 4 (taka sama jak siec nr 3 + 4 losowe krawedzie)
	 */
	
	public static SimpleWeightedGraph<String, DefaultWeightedEdge> gen4() {
		SimpleWeightedGraph<String, DefaultWeightedEdge> graph = gen3();
		
		for(int i = 0; i < 4; i++) {
			int p = 0;
			int d = 0;
			//warunek zapewnia ze nie trafimy na dwie takie same liczby
			while(p == d){
			
				Random r = new Random();
				p = r.nextInt(N) + 1;
				d = r.nextInt(N) + 1;
			}
			
			DefaultWeightedEdge edge = graph.addEdge("v"+Integer.toString(p), "v"+Integer.toString(d));
			graph.setEdgeWeight(graph.getEdge("v"+Integer.toString(p), "v"+Integer.toString(d)), 0.4);
		}
		return graph;
	}
	/*
	 * testowanie sieci nr 4
	 */
	public static void test4() {
		int zle = 0;

		for(int i = 0; i < liczbaTestow; i++) {
			if(!testNetwork(gen4())) zle++; 
		}
		double p = (liczbaTestow - zle)/(liczbaTestow * 1.0) * 100.0;
		System.out.println("Testowali�my siec nr 4 (siec trzecia z czterama losowo dodanymi krawedziami )");
		System.out.println("Szacowana niezawodno�� wynosi = "+ p + " %");
	}
	/*
	 * Dla kazdej krawedzi losujemy liczbe z przedzialu [0,1], jesli wylosowana liczba
	 * jest wieksza od wagi na naszej krawedzi to usuwamy ja. Po wszystkim spradzamy spojno�c naszej sieci
	 */
	public static boolean testNetwork(SimpleWeightedGraph<String, DefaultWeightedEdge> graph){
		for(int i=1; i<=N; i++)
        {	
    		for(int j=1; j<=N; j++)
    		{
    			if(i < j)
    			{
    				String vertex1 = "v" + Integer.toString(i);
        			String vertex2 = "v" + Integer.toString(j);
        				
        			if(graph.containsEdge(vertex1, vertex2))
        			{
        				Random generator = new Random();
        				double p = generator.nextInt(10000+1)/10000.0;
        				if(p > graph.getEdgeWeight(graph.getEdge(vertex1, vertex2)))
        				{
        					graph.removeEdge(vertex1, vertex2);
        				}
    				}
    			}
        	}
    	
    		ConnectivityInspector inspector = new ConnectivityInspector(graph);
        	if(!inspector.isGraphConnected())
    		{
    			return false;
    		}
        }

    	return true;
	}
	
	public static void main (String[] args) {
		test1();
		test2();
		test3();
		test4();
	}
}



