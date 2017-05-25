package graph;

/**
 * Created by aedd on 5/25/17
 */

public class Vertex implements Comparable<Vertex> {

    int id;
    double weight;

    public Vertex(int id, double weight){
        this.id = id;
        this.weight = weight;
    }

    @Override
    public int compareTo(Vertex compareVertex) {
        if (this.weight == compareVertex.weight) return 0;
        else if (this.weight > compareVertex.weight) return 1;
        else return -1;
    }
}
