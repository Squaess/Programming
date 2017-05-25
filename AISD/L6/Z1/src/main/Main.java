package main;

import graph.Edge;
import graph.Graph;

import java.io.BufferedReader;
import java.io.FileReader;

/**
 * Created by Bartosz on 25.05.2017
 */
public class Main {

    private final static String PATH = "./data.txt";

    public static void main(String[] args) throws Exception {

        FileReader fr = new FileReader(PATH);
        BufferedReader input = new BufferedReader(fr);

        String number_of_verticles = input.readLine();
        String number_of_edges = input.readLine();
        String line;

        int V = Integer.parseInt(number_of_verticles);
        int E = Integer.parseInt(number_of_edges);

        Graph graph = new Graph(V, E);
        for(int i = 0; i<E; i++){
            line = input.readLine();
            String[] parts = line.split(" ");
            graph.edges[i].src = Integer.parseInt(parts[0]);
            graph.edges[i].dest = Integer.parseInt(parts[1]);
            graph.edges[i].weight = Double.parseDouble(parts[2]);
        }


        graph.KruskaMST();
        graph.printSumWeight();

    }
}
