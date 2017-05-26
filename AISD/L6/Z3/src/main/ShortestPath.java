package main;

import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by Bartosz on 26.05.2017
 */

public class ShortestPath {
    // A utility function to find the vertex with minimum distance value,
    // from the set of vertices not yet included in shortest path tree
    static int V;

    int minDistance(double dist[], Boolean sptSet[]) {
        // Initialize min value
        double min = Double.MAX_VALUE;
        int min_index=-1;

        for (int v = 0; v < V; v++)
            if (sptSet[v] == false && dist[v] <= min)
            {
                min = dist[v];
                min_index = v;
            }

        return min_index;
    }

    // A utility function to print the constructed distance array
    void printSolution(double dist[]) {
        System.out.println("Vertex   Distance from Source");
        for (int i = 0; i < V; i++)
            System.out.println(i+" \t\t "+dist[i]);
    }

    // Funtion that implements Dijkstra's single source shortest path
    // algorithm for a graph represented using adjacency matrix
    // representation
    void dijkstra(double graph[][], int src) {
        V = graph.length;
        double dist[] = new double[V]; // The output array. dist[i] will hold
        // the shortest distance from src to i

        // sptSet[i] will true if vertex i is included in shortest
        // path tree or shortest distance from src to i is finalized
        Boolean sptSet[] = new Boolean[V];

        // Initialize all distances as INFINITE and stpSet[] as false
        for (int i = 0; i < V; i++)
        {
            dist[i] = Double.MAX_VALUE;
            sptSet[i] = false;
        }

        // Distance of source vertex from itself is always 0
        dist[src] = 0;

        ArrayList<Integer>[] parents = new ArrayList[V];
        for (int i = 0; i<V; i++){
            parents[i] = new ArrayList<>();
        }

        // Find shortest path for all vertices
        for (int count = 0; count < V-1; count++)
        {
            // Pick the minimum distance vertex from the set of vertices
            // not yet processed. u is always equal to src in first
            // iteration.
            int u = minDistance(dist, sptSet);

            // Mark the picked vertex as processed
            sptSet[u] = true;

            // Update dist value of the adjacent vertices of the
            // picked vertex.
            for (int v = 0; v < V; v++)

                // Update dist[v] only if is not in sptSet, there is an
                // edge from u to v, and total weight of path from src to
                // v through u is smaller than current value of dist[v]
                if (!sptSet[v] && graph[u][v]!=0 &&
                        dist[u] != Integer.MAX_VALUE &&
                        dist[u]+graph[u][v] < dist[v]) {
                    dist[v] = dist[u] + graph[u][v];
                    parents[v].clear();
                    parents[v].addAll(parents[u]);
                    parents[v].add(u);
                }
        }

        // print the constructed distance array
        printSolution(dist);
        printPaths(parents);
    }

    private void printPaths(ArrayList<Integer>[] parents) {
        System.out.println("All shortest paths:");
        for(int j = 0; j < parents.length; j++){
            for (int i:parents[j]){
                System.out.print(i+" ");
            }
            System.out.print(j);
            System.out.print("\n");
        }
    }

    // Driver method
    public static void main (String[] args) throws IOException {
        /* Let us create the example graph discussed above */
        FileReader fr = new FileReader("./data.txt");
        BufferedReader input = new BufferedReader(fr);

        int V = Integer.parseInt(input.readLine());
        int E = Integer.parseInt(input.readLine());

        double graph[][] = new double[V][V];
        for (int i = 0; i < V; i++){
            for ( int j = 0; j<V; j++){
                graph[i][j] = 0;
            }
        }
        String line;
        for (int i = 0; i < E; i++){
            line = input.readLine();
            String[] parts = line.split(" ");
            int v1 = Integer.parseInt(parts[0]);
            int v2 = Integer.parseInt(parts[1]);
            double weight = Double.parseDouble(parts[2]);
            graph[v1][v2] = weight;
        }
        ShortestPath t = new ShortestPath();
        t.dijkstra(graph, 0);
    }
}
