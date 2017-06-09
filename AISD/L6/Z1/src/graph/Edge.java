package graph;

/**
 * Created by Bartosz on 25.05.2017
 */

public class Edge implements Comparable<Edge> {
    public int src, dest;
    public double weight;

     public Edge(int src, int dest, double weight){
         this.dest = dest;
         this.src = src;
         this.weight = weight;
     }

     Edge(){

     }

    @Override
    public int compareTo(Edge compareEdge) {
        if (this.weight == compareEdge.weight) return 0;
        else if (this.weight > compareEdge.weight) return 1;
        else return -1;
    }
}