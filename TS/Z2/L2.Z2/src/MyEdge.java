
public class MyEdge {
	private String vertex1;
	private String vertex2;
	
	private double capacity;
	
	public MyEdge(String v1, String v2, double c)
	{
		vertex1 = v1;
		vertex2 = v2;
		
		capacity = c;
	}
	
	public String v1()
	{
		return vertex1;
	}
	
	public String v2()
	{
		return vertex2;
	}
	
	public double getCapacity()
	{
		return capacity;
	}
}
