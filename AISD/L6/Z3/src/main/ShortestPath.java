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
    // ilosc wierzchołków
    static int V;

    private int minDistance(double dist[], Boolean sptSet[]) {
        // zwracamy indeks wierzcholka do którego jeszcze nie jest skonczone obliczanie
        // a dystans jest najmniejszy
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

    // Wypisywanie samych distance
    void printSolution(double dist[]) {
        System.out.println("Vertex   Odległość od source");
        for (int i = 0; i < V; i++)
            System.out.println(i+" \t\t "+dist[i]);
    }

    void dijkstra(double graph[][], int src) {
        V = graph.length;
        double dist[] = new double[V];

        // jeśli wierzchołek już ma znaleziony min dist to rue
        Boolean sptSet[] = new Boolean[V];

        // Inicjalizacja
        for (int i = 0; i < V; i++)
        {
            dist[i] = Double.MAX_VALUE;
            sptSet[i] = false;
        }

        dist[src] = 0;

        // lista przetrzymująca trase dla danego wierzchołka
        ArrayList<Integer>[] parents = new ArrayList[V];
        for (int i = 0; i<V; i++){
            parents[i] = new ArrayList<>();
        }

        // Find shortest path for all vertices
        for (int count = 0; count < V-1; count++)
        {
            // Wybieramy minimalna odleglosc z wierchołków jeszcze nie zatwierdzonych
            int u = minDistance(dist, sptSet);

            // Zatwierdzamy wierzcholek
            sptSet[u] = true;

            // Update na dist dla wierchołków łączących sie z u
            for (int v = 0; v < V; v++)

                // Robimy update tylko jęsli jeszcze nie jest w sptSet[v]
                // i istnieje krawęðz (u,v) i jeśli now wartosc ma byc mniejsza od starej dla v
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
