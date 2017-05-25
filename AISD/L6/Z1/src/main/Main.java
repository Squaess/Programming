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
//        int V = 5;
//        int E = 7;
//        Graph graph = new Graph(V, E);

//        graph.edges[0].src = 0;
//        graph.edges[0].dest = 1;
//        graph.edges[0].weight = 4;
//
//        graph.edges[1].src = 0;
//        graph.edges[1].dest = 2;
//        graph.edges[1].weight = 4;
//
//        graph.edges[2].src = 0;
//        graph.edges[2].dest = 4;
//        graph.edges[2].weight = 6;
//
//        graph.edges[3].src = 0;
//        graph.edges[3].dest = 3;
//        graph.edges[3].weight = 6;
//
//        graph.edges[4].src = 1;
//        graph.edges[4].dest = 2;
//        graph.edges[4].weight = 2;
//
//        graph.edges[5].src = 2;
//        graph.edges[5].dest = 3;
//        graph.edges[5].weight = 8;
//
//        graph.edges[6].src = 3;
//        graph.edges[6].dest = 4;
//        graph.edges[6].weight = 9;
//
//        graph.KruskaMST();

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


        graph.printSumWeight();
        graph.KruskaMST();

    }
}
