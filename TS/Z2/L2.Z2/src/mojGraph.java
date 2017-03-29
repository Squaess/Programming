import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Random;

import org.jgraph.graph.DefaultEdge;
import org.jgrapht.alg.ConnectivityInspector;
import org.jgrapht.alg.DijkstraShortestPath;
import org.jgrapht.graph.DefaultWeightedEdge;
import org.jgrapht.graph.SimpleDirectedWeightedGraph;
import org.jgrapht.graph.SimpleGraph;

public class mojGraph {
	private SimpleDirectedWeightedGraph<String, DefaultWeightedEdge>  graph = new SimpleDirectedWeightedGraph<String, DefaultWeightedEdge>(DefaultWeightedEdge.class);
	private List<String> vertices = new ArrayList<String>();
	private List<MyEdge> edges = new ArrayList<MyEdge>();
	
	private int sumOfPacketsSize = 0;
	private int numberOfPackets = 0;
	
	private List<List<Double>> N = new ArrayList<List<Double>>();

	public void addVertex(String vertexName)
	{
		vertices.add(vertexName);
		graph.addVertex(vertexName);
	}
	
	public void addEdge(String vertex1, String vertex2, double c)
	{
		if(graph.containsEdge(vertex1, vertex2) || graph.containsEdge(vertex2, vertex1))
		{
			return;
		}
			
		MyEdge edge = new MyEdge(vertex1, vertex2, c);
		edges.add(edge);
		
		DefaultWeightedEdge e = graph.addEdge(vertex1, vertex2); 
        graph.setEdgeWeight(e, 0); 
	}
	
	public void deleteEdge(String vertex1, String vertex2)
	{	
		if(graph.containsEdge(vertex1, vertex2))
		graph.removeEdge(vertex1, vertex2);
		else if(graph.containsEdge(vertex2, vertex1))
		graph.removeEdge(vertex2, vertex1);
		
		for(int i=0; i<edges.size(); i++)
		{
			MyEdge edge = edges.get(i);
			String v1 = edge.v1();
			String v2 = edge.v2();
			
			if(v1.contentEquals(vertex1) || v1.contentEquals(vertex2))
			{
				if(v2.contentEquals(vertex1) || v2.contentEquals(vertex2))
				{
					edges.remove(i);
					return;
				}
			}
		}
	}
	
	public void clearConnections()
	{
		N.clear();
		
		for(int i=1; i<= vertices.size(); i++)
		{
			List<Double> L = new ArrayList<Double>();
			
			for(int j=1; j<= vertices.size(); j++)
			{
				L.add(0.0);
			}
			
			N.add(L);
		}

		sumOfPacketsSize = 0;
		numberOfPackets = 0;
		//System.out.println(Arrays.deepToString(N.toArray()));
	}
	
	public void displayGraph()
	{
		System.out.println(graph.toString());
	}
	
	public void sendPacket(String vertex1, String vertex2, int packetSize)
	{
		if(vertex1.contentEquals(vertex2))
		{
			return;
		}
		
		List path = DijkstraShortestPath.findPathBetween(graph, vertex1, vertex2);
		
		for(int i=0; i<path.size(); i++)
		{
			String S = path.get(i).toString();
			S = S.replace("(", "");
			S = S.replace(")", "");
			
			String[] parts = S.split(" : ");
			
			//Szukanie ID wierzcho³ków
			int ID1 = 0, ID2 = 0;
			for(int j=0; j<vertices.size(); j++)
			{
				if(vertices.get(j).contentEquals(parts[0]))
					ID1 = j;
				else if(vertices.get(j).contentEquals(parts[1]))
					ID2 = j;
			}
		
			//Wpisywanie do macierzy natê¿eñ
			List<Double> L = N.get(ID1);
			double value = L.get(ID2);
			value = value + packetSize;
			L.set(ID2, value);
			N.set(ID1, L);
		}
		
		sumOfPacketsSize += packetSize;
		numberOfPackets++;
		
		checkCongestion();
	}
	
	private void checkCongestion()
	{
		for(int i=0; i<edges.size(); i++)
		{
			MyEdge edge = edges.get(i);
			String vertex1 = edge.v1();
			String vertex2 = edge.v2();
			
			//Szukanie ID wierzcho³ków
			int ID1 = 0, ID2 = 0;
			for(int j=0; j<vertices.size(); j++)
			{
				if(vertices.get(j).contentEquals(vertex1))
					ID1 = j;
				else if(vertices.get(j).contentEquals(vertex2))
					ID2 = j;
			}
			
			//Sprawdzanie macierzy natê¿eñ
			double value = 0;
			List<Double> L = N.get(ID1);
			value += L.get(ID2);
			
			List<Double> L2 = N.get(ID2);
			value += L2.get(ID1);
			
			if(value > edge.getCapacity())
			{
				System.out.println("£¹cze " + vertex1 + "-" + vertex2 + " przeci¹¿one! (" + value + "/" + edge.getCapacity() + ")");	
			}
		}
		
	}
	
	public void displayInformation()
	{
		System.out.println("\nInformacje o ³¹czach: ");
		for(int i=0; i<edges.size(); i++)
		{
			MyEdge edge = edges.get(i);
			String vertex1 = edge.v1();
			String vertex2 = edge.v2();
			
			//Szukanie ID wierzcho³ków
			int ID1 = 0, ID2 = 0;
			for(int j=0; j<vertices.size(); j++)
			{
				if(vertices.get(j).contentEquals(vertex1))
					ID1 = j;
				else if(vertices.get(j).contentEquals(vertex2))
					ID2 = j;
			}
			
			//Sprawdzanie macierzy natê¿eñ
			double value = 0;
			List<Double> L = N.get(ID1);
			value += L.get(ID2);
			
			List<Double> L2 = N.get(ID2);
			value += L2.get(ID1);
			

			System.out.println("£¹cze " + vertex1 + "-" + vertex2 + ": natê¿enie pakietów " + value + ", przepustowoœc " + edge.getCapacity());
		}
		
		System.out.println("\nMacierz natê¿eñ: ");
		System.out.println(Arrays.deepToString(N.toArray()));
		
	}
	
	public double getDelay()
	{
		//G = suma elementów macierzy natê¿eñ
		double G = 0;
		for(int i=0; i<N.size(); i++)
		{
			List<Double> L = N.get(i);
			for(int j=0; j<L.size(); j++)
			{
				G += L.get(j);
			}
		}
		
		//m = œrednia wielkoœc pakietu
		double m = (sumOfPacketsSize * 1.0)/(numberOfPackets * 1.0);
		
		//Liczenie sumy
		double SUM = 0;
		for(int i=0; i<edges.size(); i++)
		{
			MyEdge edge = edges.get(i);
			String vertex1 = edge.v1();
			String vertex2 = edge.v2();
			
			int ID1 = 0, ID2 = 0;
			for(int j=0; j<vertices.size(); j++)
			{
				if(vertices.get(j).contentEquals(vertex1))
					ID1 = j;
				else if(vertices.get(j).contentEquals(vertex2))
					ID2 = j;
			}
			
			//Pobieranie a
			double A = 0;
			List<Double> L = N.get(ID1);
			A += L.get(ID2);
			List<Double> L2 = N.get(ID2);
			A += L2.get(ID1);
			
			//Pobieranie c
			double C = edge.getCapacity();
			
			//Dodawanie do sumy
			if(A > 0)
			{
				double licznik = A;
				double mianownik = (C/m) - A;
				
				SUM += licznik / mianownik;
			}
		}
		
		//Wynik
		return (1/G) * SUM;
	}
	
	public void test(double p)
	{
		for(int i=0; i<edges.size(); i++)
		{
			MyEdge edge = edges.get(i);
			String vertex1 = edge.v1();
			String vertex2 = edge.v2();
			
			Random generator = new Random();
			double q = (generator.nextInt(10001)*1.0)/(10000);
			if(q > p)
			{
				this.deleteEdge(vertex1, vertex2);
			}
		}
	}
	
	public boolean isItConnected()
	{
		ConnectivityInspector inspector = new ConnectivityInspector(graph);
    	if(!inspector.isGraphConnected())
		{
			return false;
		}
		return true;
	}
	
}
