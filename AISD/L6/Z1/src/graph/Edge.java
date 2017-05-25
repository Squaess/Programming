package graph;

/**
 * Created by Bartosz on 25.05.2017
 */

public class Edge {
    public int src, dest;
    public double weight;

     public Edge(int src, int dest, double weight){
         this.dest = dest;
         this.src = src;
         this.weight = weight;
     }

     Edge(){

     }
}
