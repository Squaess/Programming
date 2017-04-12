
class MEdge {
    private int v1;
    private int v2;
    private int capacity;
    private int flow;
    private int numberOfPackets;

    MEdge(int vertex1, int vertex2, int capacity) {
        this.v1 = vertex1;
        this.v2 = vertex2;
        this.capacity = capacity;
        this.flow = 0;
        this.numberOfPackets = 0;
    }

    int getFlow() {
        return this.flow;
    }

    boolean setFlow(int value) {
        this.flow += value;
        if(flow >= capacity) {
            return false;
        }
        return true;
    }

    int getV1(){
        return this.v1;
    }

    int getV2() {
        return this.v2;
    }

    int getCapacity() {
        return this.capacity;
    }

    int getNumberOfPackets() {
        return this.numberOfPackets;
    }

    void setNumberOfPackets(int numberOfPackets) {
        this.numberOfPackets += numberOfPackets;
    }
}
