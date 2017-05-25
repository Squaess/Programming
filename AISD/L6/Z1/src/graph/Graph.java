package graph;

import main.MyQueue;

/**
 * Created by Bartosz on 25.05.2017
 */

public class Graph {

    int V, E;
    public Edge[] edges;

    public Graph(int V, int E){
        this.E = E;
        this.V = V;
        edges = new Edge[E];
        for (int i = 0; i < edges.length; i++){
            edges[i] = new Edge();
        }
    }

    class Subset{
        int parent;
    }

    int find(Subset[] subsets, int i){
        if(subsets[i].parent != i)
            subsets[i].parent = find(subsets,subsets[i].parent);

        return subsets[i].parent;
    }

    void union(Subset[] subsets, int x, int y){
        int xroot = find(subsets, x);
        int yroot = find(subsets, y);

        subsets[yroot].parent = xroot;
    }

    public void KruskaMST() throws Exception {
        Edge[] result = new Edge[V];
        int e = 0;
        int i = 0;

        for (i = 0; i<V; ++i){
            result[i] = new Edge();
        }

        MyQueue heap = new MyQueue(edges);

        Subset subsets[] = new Subset[V];
        for (i=0; i<subsets.length; i++)
            subsets[i] = new Subset();

        for (int v = 0; v < V; v++){
            subsets[v].parent = v;
        }

        i = 0;

        while( e < V -1){
            Edge next_edge = new Edge();
            next_edge = heap.heap_Extraxt_Max();

            int x = find(subsets, next_edge.src);
            int y = find(subsets, next_edge.dest);

            if(x != y){
                result[e++] = next_edge;
                union(subsets,x,y);
            }
        }
        System.out.println("Following are the edges in constructed MST");
        for (i = 0; i < e; ++i){
            System.out.println(result[i].src+" -- "+result[i].dest + " weight: "+result[i].weight);
        }
    }

    public void printSumWeight(){
        double suma = 0;
        for (Edge e:edges){
            suma += e.weight;
            System.out.print(e.weight+" ");
        }
        System.out.println("Suma wszystkich krawedzi: "+suma);
    }
}
